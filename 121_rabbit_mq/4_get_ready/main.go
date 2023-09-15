package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// RabbitMQ管理API的URL，替换为您的RabbitMQ实例的URL
	rabbitmqAPIURL := "http://localhost:15672/api/queues/pre_prod/queue_3"

	// RabbitMQ管理API的用户名和密码，替换为您的凭据
	username := "guest"
	password := "guest"

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", rabbitmqAPIURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置HTTP请求的基本认证头部
	req.SetBasicAuth(username, password)

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 解析JSON响应
	// 注意：此处需要使用适当的JSON解析库，例如encoding/json
	// 这里仅提供了简单的示例
	fmt.Println("Response:", string(body))

	var data AutoGenerated
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ready", data.MessagesReady)
}

type AutoGenerated struct {
	ConsumerDetails []any `json:"consumer_details"`
	Arguments       struct {
		XQueueType string `json:"x-queue-type"`
	} `json:"arguments"`
	AutoDelete         bool `json:"auto_delete"`
	BackingQueueStatus struct {
		AvgAckEgressRate  float64 `json:"avg_ack_egress_rate"`
		AvgAckIngressRate float64 `json:"avg_ack_ingress_rate"`
		AvgEgressRate     float64 `json:"avg_egress_rate"`
		AvgIngressRate    float64 `json:"avg_ingress_rate"`
		Delta             []any   `json:"delta"`
		Len               int     `json:"len"`
		Mode              string  `json:"mode"`
		NextDeliverSeqID  int     `json:"next_deliver_seq_id"`
		NextSeqID         int     `json:"next_seq_id"`
		NumPendingAcks    int     `json:"num_pending_acks"`
		NumUnconfirmed    int     `json:"num_unconfirmed"`
		Q1                int     `json:"q1"`
		Q2                int     `json:"q2"`
		Q3                int     `json:"q3"`
		Q4                int     `json:"q4"`
		TargetRAMCount    string  `json:"target_ram_count"`
		Version           int     `json:"version"`
	} `json:"backing_queue_status"`
	ConsumerCapacity          int   `json:"consumer_capacity"`
	ConsumerUtilisation       int   `json:"consumer_utilisation"`
	Consumers                 int   `json:"consumers"`
	Deliveries                []any `json:"deliveries"`
	Durable                   bool  `json:"durable"`
	EffectivePolicyDefinition struct {
	} `json:"effective_policy_definition"`
	Exclusive            bool `json:"exclusive"`
	ExclusiveConsumerTag any  `json:"exclusive_consumer_tag"`
	GarbageCollection    struct {
		FullsweepAfter  int `json:"fullsweep_after"`
		MaxHeapSize     int `json:"max_heap_size"`
		MinBinVheapSize int `json:"min_bin_vheap_size"`
		MinHeapSize     int `json:"min_heap_size"`
		MinorGcs        int `json:"minor_gcs"`
	} `json:"garbage_collection"`
	HeadMessageTimestamp       any       `json:"head_message_timestamp"`
	IdleSince                  time.Time `json:"idle_since"`
	Incoming                   []any     `json:"incoming"`
	Memory                     int       `json:"memory"`
	MessageBytes               int       `json:"message_bytes"`
	MessageBytesPagedOut       int       `json:"message_bytes_paged_out"`
	MessageBytesPersistent     int       `json:"message_bytes_persistent"`
	MessageBytesRAM            int       `json:"message_bytes_ram"`
	MessageBytesReady          int       `json:"message_bytes_ready"`
	MessageBytesUnacknowledged int       `json:"message_bytes_unacknowledged"`
	MessageStats               struct {
		Publish        int `json:"publish"`
		PublishDetails struct {
			Rate float64 `json:"rate"`
		} `json:"publish_details"`
	} `json:"message_stats"`
	Messages        int `json:"messages"`
	MessagesDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_details"`
	MessagesPagedOut     int `json:"messages_paged_out"`
	MessagesPersistent   int `json:"messages_persistent"`
	MessagesReady        int `json:"messages_ready"`
	MessagesReadyDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_ready_details"`
	MessagesReadyRAM              int `json:"messages_ready_ram"`
	MessagesUnacknowledged        int `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_unacknowledged_details"`
	MessagesUnacknowledgedRAM int    `json:"messages_unacknowledged_ram"`
	Name                      string `json:"name"`
	Node                      string `json:"node"`
	OperatorPolicy            any    `json:"operator_policy"`
	Policy                    any    `json:"policy"`
	RecoverableSlaves         any    `json:"recoverable_slaves"`
	Reductions                int    `json:"reductions"`
	ReductionsDetails         struct {
		Rate float64 `json:"rate"`
	} `json:"reductions_details"`
	SingleActiveConsumerTag any    `json:"single_active_consumer_tag"`
	State                   string `json:"state"`
	Type                    string `json:"type"`
	Vhost                   string `json:"vhost"`
}