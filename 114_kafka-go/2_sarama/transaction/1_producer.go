package main

// SIGUSR1 toggle the pause/resume consumption
import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	brokers          = ""
	version          = ""
	group            = ""
	topics           = ""
	destinationTopic = ""
	oldest           = true
	verbose          = false
	assignor         = ""
)

func init() {
	flag.StringVar(&brokers, "brokers", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&group, "group", "", "Kafka consumer group definition")
	flag.StringVar(&version, "version", sarama.DefaultVersion.String(), "Kafka cluster version")
	flag.StringVar(&topics, "topics", "", "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&destinationTopic, "destination-topic", "", "Kafka topic where records will be copied from topics.")
	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	// 可以传入多个topics
	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(destinationTopic) == 0 {
		panic("no destination topics given to be consumed, please set the -destination-topics flag")
	}

	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

func main() {
	keepRunning := true
	log.Println("Starting a new Sarama consumer")

	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	config.Consumer.IsolationLevel = sarama.ReadCommitted
	// IsolationLevel 在 Kafka 中是用来控制消费者消费消息时的隔离级别，支持两种模式：
	//
	//ReadUncommitted：这是默认的隔离级别。在此模式下，消费者可以读取所有消息，包括那些可能属于已中止事务的消息。消费者会返回来自所有事务的消息，无论它们的提交状态如何，即使有些消息可能是未提交或已中止的。
	//
	//ReadCommitted：当设置为此模式时，消费者会隐藏属于已中止事务的消息。这意味着消费者将只返回已经提交的事务中的消息，而不会返回未提交或已中止的事务中的消息。这样可以确保消费者仅获取到被成功提交的消息。
	//
	//隔离级别的选择取决于应用程序的需求。如果需要读取所有消息并且不关心事务提交状态，则可以使用 ReadUncommitted 模式。但如果需要忽略未提交或已中止的事务的消息，只消费已经成功提交的消息，则可以选择 ReadCommitted 模式，这样可以避免读取未稳定的数据。
	// 未提交的消息指的是生产者发送到 Kafka 但尚未被确认提交的消息。这些消息可能包含在尚未完成的事务中，或者生产者尚未收到 Kafka Broker 的确认。
	//
	//已中止的消息是指属于已中止的事务的消息，这些事务可能由于某种原因（例如，提交事务之前发生了错误）而被中止。在 ReadCommitted 模式下，这些已中止的消息将被隐藏，消费者不会看到这些消息，从而确保只有已成功提交的消息被消费。
	config.Consumer.Offsets.AutoCommit.Enable = false

	producerProvider := newProducerProvider(strings.Split(brokers, ","), func() *sarama.Config {
		producerConfig := sarama.NewConfig()
		producerConfig.Version = version

		// producer.configs_max.in.flight.requests.per.connection 一个连接之前允许有多少个未完成的请求
		producerConfig.Net.MaxOpenRequests = 1
		// Net.MaxOpenRequests:
		//
		//默认情况下，该参数是允许同时进行的未完成请求的最大数量。设置为1确保了只有一个请求在进行中，这对于事务处理非常重要，因为它确保了请求的顺序性。
		//在使用事务时，需要确保请求按预期的顺序发送到 Kafka Broker，以便维护事务的完整性。
		//将 Net.MaxOpenRequests 设置为1可确保只有一个请求在进行中，这样可以避免并发请求的可能性，从而确保事务中的操作是有序的。
		// 设置大于1就确保不了请求的顺序性
		producerConfig.Producer.RequiredAcks = sarama.WaitForAll
		// Producer.RequiredAcks:
		//
		//设置为 sarama.WaitForAll 意味着生产者需要等待所有副本都已成功写入后才会收到来自服务器的成功响应。这对于事务完整性非常重要，因为它确保了所有副本都成功写入了消息。
		producerConfig.Producer.Idempotent = true
		// 在 Kafka 中，Idempotent 是生产者（Producer）的一个配置选项。当设置为 true 时，表示启用了幂等性（Idempotence）。启用幂等性后，生产者会确保每条消息在分区中的写入是“幂等”的，也就是确保每条消息最终只会被成功写入一次，无论是因为网络错误、生产者重新发送、或者其他原因。
		//
		// 确保消息的顺序性、幂等性和完整性
		//启用幂等性能够保证以下几点：
		//
		//消息发送具有原子性：同一条消息不会被多次重复写入同一分区，从而避免了重复消息的问题。
		//在网络或其他错误发生时，生产者重试消息发送不会导致重复消息的产生。
		//这个选项通常在需要确保消息不被重复写入的情况下使用，尤其是在对消息的唯一性要求较高的场景下。启用幂等性能够避免重复消息带来的问题，并且不会影响到 Kafka 的性能。
		producerConfig.Producer.Transaction.ID = "sarama"
		return producerConfig
	})

	/**
	 * Set up a new Sarama consumer group
	 */
	consumer := Consumer{
		groupId:          group,
		brokers:          strings.Split(brokers, ","),
		producerProvider: producerProvider,
		ready:            make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					// 什么时候会出现这个错误
					// 消费的时候遇到了ErrClosedConsumerGroup 错误，说明消费者组已经被关闭了，此时应该退出循环，结束消费
					log.Println("kafka consumer group has been closed")
					return
				}
				log.Panicf("kafka consume Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println("kafka consume context has been cancelled")
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)
	// syscall.SIGUSR1 是一个特殊的信号，它表示用户定义的信号1。在大多数 Unix 或类 Unix 操作系统上，SIGUSR1 是由用户自定义的，可以用于应用程序内部的信号处理。
	// 开发人员可以根据需要将 SIGUSR1 和 SIGUSR2 信号用作自定义信号，例如，用于触发特定操作、通知程序执行某些功能或执行特定的任务等。这两个信号并没有预先定义的行为，而是由程序开发人员根据需要自行定义和处理。

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	// syscall.SIGINT 代表中断信号，通常是由用户发出的终止程序的信号。在终端窗口中，按下 Ctrl + C 会发送 SIGINT 信号给当前运行的程序，以请求程序正常地终止执行。
	//
	//syscall.SIGTERM 是终止信号，类似于 SIGINT，但通常不由用户直接发送。它被用于请求程序正常终止执行。通常情况下，操作系统或其他程序通过发送 SIGTERM 来请求另一个程序安全地关闭和退出。
	//
	//在处理这两个信号时，程序可以选择捕获它们并执行特定的操作或清理工作，然后正常退出。如果程序未捕获这些信号，操作系统可能会以默认方式终止该程序的执行。

	// 然后一直监听,如果超时，如果收到了信号，就退出，如果自定义信号
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()

	producerProvider.clear()

	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready            chan bool
	groupId          string
	brokers          []string
	producerProvider *producerProvider
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L2
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			fmt.Println("message received", message)
			func() {
				producer := consumer.producerProvider.borrow(message.Topic, message.Partition)
				defer consumer.producerProvider.release(message.Topic, message.Partition, producer)

				startTime := time.Now()

				// BeginTxn must be called before any messages.
				// 开启事务
				err := producer.BeginTxn()
				if err != nil {
					log.Printf("Message consumer: unable to start transaction: %+v", err)
					return
				}
				// Produce current record in producer transaction.
				producer.Input() <- &sarama.ProducerMessage{
					Topic: destinationTopic,
					Key:   sarama.ByteEncoder(message.Key),
					Value: sarama.ByteEncoder(message.Value),
				}

				// You can add current message to this transaction
				err = producer.AddMessageToTxn(message, consumer.groupId, nil)
				if err != nil {
					log.Println("error on AddMessageToTxn")
					consumer.handleTxnError(producer, message, session, err, func() error {
						return producer.AddMessageToTxn(message, consumer.groupId, nil)
					})
					return
				}

				// Commit producer transaction.
				// 提交事务
				err = producer.CommitTxn()
				if err != nil {
					log.Println("error on CommitTxn")
					consumer.handleTxnError(producer, message, session, err, func() error {
						return producer.CommitTxn()
					})
					return
				}
				log.Printf("Message claimed [%s]: value = %s, timestamp = %v, topic = %s, partition = %d", time.Since(startTime), string(message.Value), message.Timestamp, message.Topic, message.Partition)
			}()
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

func (consumer *Consumer) handleTxnError(producer sarama.AsyncProducer, message *sarama.ConsumerMessage, session sarama.ConsumerGroupSession, err error, defaulthandler func() error) {
	log.Printf("Message consumer: unable to process transaction: %+v", err)
	for {
		if producer.TxnStatus()&sarama.ProducerTxnFlagFatalError != 0 {
			// 如果两个相应的位都为 1，则结果为 1。
			// 如果两个相应的位中至少有一个为 0，则结果为 0。
			// fatal error. need to recreate producer.
			log.Printf("kafka Message consumer: producer is in a fatal state, need to recreate it")
			// reset current consumer offset to retry consume this record.
			session.ResetOffset(message.Topic, message.Partition, message.Offset, "")
			return
		}
		if producer.TxnStatus()&sarama.ProducerTxnFlagAbortableError != 0 {
			err = producer.AbortTxn()
			if err != nil {
				log.Printf("Message consumer: unable to abort transaction: %+v", err)
				continue
			}
			// reset current consumer offset to retry consume this record.
			session.ResetOffset(message.Topic, message.Partition, message.Offset, "")
			return
		}
		// if not you can retry
		err = defaulthandler()
		if err == nil {
			return
		}
	}
}

type topicPartition struct {
	topic     string
	partition int32
}

type producerProvider struct {
	producersLock sync.Mutex
	producers     map[topicPartition][]sarama.AsyncProducer

	producerProvider func(topic string, partition int32) sarama.AsyncProducer
}

func newProducerProvider(brokers []string, producerConfigurationProvider func() *sarama.Config) *producerProvider {
	provider := &producerProvider{
		producers: make(map[topicPartition][]sarama.AsyncProducer),
	}
	provider.producerProvider = func(topic string, partition int32) sarama.AsyncProducer {
		config := producerConfigurationProvider()
		if config.Producer.Transaction.ID != "" {
			config.Producer.Transaction.ID = config.Producer.Transaction.ID + "-" + topic + "-" + fmt.Sprint(partition)
		}
		producer, err := sarama.NewAsyncProducer(brokers, config)
		if err != nil {
			return nil
		}
		return producer
	}
	return provider
}

func (p *producerProvider) borrow(topic string, partition int32) (producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	tp := topicPartition{topic: topic, partition: partition}
	// 相同值的结构体在map中的key是一样的
	if producers, ok := p.producers[tp]; !ok || len(producers) == 0 { // 如果没有对应的topic和分区 获取该topic下没有对应的消费者组
		for {
			producer = p.producerProvider(topic, partition)
			if producer != nil {
				return
			}
		}
	}

	index := len(p.producers[tp]) - 1
	producer = p.producers[tp][index]
	p.producers[tp] = p.producers[tp][:index]
	return
}

func (p *producerProvider) release(topic string, partition int32, producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	if producer.TxnStatus()&sarama.ProducerTxnFlagInError != 0 {
		// Try to close it
		_ = producer.Close()
		return
	}
	tp := topicPartition{topic: topic, partition: partition}
	p.producers[tp] = append(p.producers[tp], producer)
}

func (p *producerProvider) clear() {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	for _, producers := range p.producers {
		for _, producer := range producers {
			producer.Close()
		}
	}
	for _, producers := range p.producers {
		producers = producers[:0]
	}
}

// $ go run main.go -brokers="47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092" -topics="first" -destination-topic="destination-sarama" -group="example"
// "47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092"

// 一亿条数据 高峰期1w/s
// 1w5/s 处理日志速度
// 一条日志 (2k)
// 1w5/s * 2k = 30M/s x 2 = 60M/s
//
// 服务器台数 2 * (生产者峰值生产速率 * 服本数 / 100) + 1
// 2 * (60M/s * 2 / 100 ) + 1 = 3
// 3台服务器  磁盘大小 需要多少
// 内存大小  堆内存 (内部配置) + 页缓存 (服务器内存)
//
