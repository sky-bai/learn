package config

var (
	GaoDeConfigSetting   = &GaoDeConfig{}
	TencentConfigSetting = &TencentConfig{}

	TestConfigSetting = &KafkaConfig{}
)

type GaoDeConfig struct {

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

type TencentConfig struct {
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

type KafkaConfig struct {
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

// batch.size 只有数据积累到batch.size之后才会发送 默认是16k
// linger.ms 如果数据没有达到batch.size，那么在linger.ms后也会发送 默认是0ms,表示数据立即发送
