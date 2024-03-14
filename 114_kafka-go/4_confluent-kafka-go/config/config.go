package config

var (
	ChannelKafkaConfigSetting = &ChannelKafkaConfig{}
)

type (
	ChannelKafkaConfig struct {
		GaoDe       GaoDeConfig       `yaml:"GaoDe"`
		Tencent     TencentConfig     `yaml:"Tencent"`
		Test        TestConfig        `yaml:"Test"`
		Transaction TransactionConfig `yaml:"Transaction"`

		Consumer KafkaConsumerConfig `yaml:"Consumer"`
	}

	GaoDeConfig struct {

		// 连接配置
		BootstrapServers string `yaml:"BootstrapServers"`
		Topic            string `yaml:"Topic"`

		// 发送配置
		LingerMs                 int    `yaml:"LingerMs"`
		BatchSize                int    `yaml:"BatchSize"`
		QueueBuffingMaxKBytes    int    `yaml:"QueueBuffingMaxKBytes"`
		QueueBufferIngMaxMessage int    `yaml:"QueueBufferIngMaxMessage"`
		CompressionCodec         string `yaml:"CompressionCodec"`
		Acks                     string `yaml:"Acks"`
		Retries                  int    `yaml:"Retries"`
		RetryBackoffMs           int    `yaml:"RetryBackoffMs"`
	}

	TencentConfig struct {
		// 连接配置
		BootstrapServers string `yaml:"BootstrapServers"`
		Topic            string `yaml:"Topic"`

		// 发送配置
		LingerMs                 int    `yaml:"LingerMs"`
		BatchSize                int    `yaml:"BatchSize"`
		QueueBuffingMaxKBytes    int    `yaml:"QueueBuffingMaxKBytes"`
		QueueBufferIngMaxMessage int    `yaml:"QueueBufferIngMaxMessage"`
		CompressionCodec         string `yaml:"CompressionCodec"`
		Acks                     string `yaml:"Acks"`
		Retries                  int    `yaml:"Retries"`
		RetryBackoffMs           int    `yaml:"RetryBackoffMs"`
	}

	TestConfig struct {
		// 连接配置
		BootstrapServers string `yaml:"BootstrapServers"`
		Topic            string `yaml:"Topic"`
		Partition        int    `yaml:"Partition"`

		// 发送配置
		LingerMs                 int    `yaml:"LingerMs"`
		BatchSize                int    `yaml:"BatchSize"`
		QueueBuffingMaxKBytes    int    `yaml:"QueueBuffingMaxKBytes"`
		QueueBufferIngMaxMessage int    `yaml:"QueueBufferIngMaxMessage"`
		CompressionCodec         string `yaml:"CompressionCodec"`
		Acks                     string `yaml:"Acks"`
		Retries                  int    `yaml:"Retries"`
		RetryBackoffMs           int    `yaml:"RetryBackoffMs"`
	}

	TransactionConfig struct {
		// 连接配置
		BootstrapServers string `yaml:"BootstrapServers"`
		Topic            string `yaml:"Topic"`
		Partition        int    `yaml:"Partition"`
		TransactionId    string `yaml:"TransactionId"`

		// 发送配置
		LingerMs                 int    `yaml:"LingerMs"`
		BatchSize                int    `yaml:"BatchSize"`
		QueueBuffingMaxKBytes    int    `yaml:"QueueBuffingMaxKBytes"`
		QueueBufferIngMaxMessage int    `yaml:"QueueBufferIngMaxMessage"`
		CompressionCodec         string `yaml:"CompressionCodec"`
		Acks                     string `yaml:"Acks"`
		Retries                  int    `yaml:"Retries"`
		RetryBackoffMs           int    `yaml:"RetryBackoffMs"`
	}
)

type KafkaConfig struct {
	// 连接配置
	BootstrapServers string `yaml:"BootstrapServers"`
	Topic            string `yaml:"Topic"`
	Partition        int    `yaml:"Partition"`
	TransactionId    string `yaml:"TransactionId"`

	// 发送配置
	LingerMs                 int    `yaml:"LingerMs"`
	BatchSize                int    `yaml:"BatchSize"`
	QueueBuffingMaxKBytes    int    `yaml:"QueueBuffingMaxKBytes"`
	QueueBufferIngMaxMessage int    `yaml:"QueueBufferIngMaxMessage"`
	CompressionCodec         string `yaml:"CompressionCodec"`
	Acks                     string `yaml:"Acks"`
	Retries                  int    `yaml:"Retries"`
	RetryBackoffMs           int    `yaml:"RetryBackoffMs"`
}

// batch.size 只有数据积累到batch.size之后才会发送 默认是16k
// linger.ms 如果数据没有达到batch.size，那么在linger.ms后也会发送 默认是0ms,表示数据立即发送

// 多文件进行读取

type KafkaConsumerConfig struct {
	// 连接配置
	BootstrapServers string `yaml:"BootstrapServers"`

	// 消费配置
	Topic           []string `yaml:"Topic"`
	GroupId         string   `yaml:"GroupId"`
	AutoOffsetReset string   `yaml:"AutoOffsetReset"`
}
