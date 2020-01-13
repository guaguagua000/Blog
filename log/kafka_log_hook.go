package log

import (
	"Blog/config"
	kafka2 "Blog/kafka"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

type KafkaLogHook struct {
}

/**
实现KafkaLogHook的io.writer接口方法, 使当前对象可以作为zap的hook使用
*/
func (this *KafkaLogHook) Write(p []byte) (n int, err error) {
	if !this.SendKafkaMessage(config.GlobalConfig.LogKafkaTopic, string(p)) {
		HookLogger.Error("写kafka日志异常")
		return 0, errors.New("写kafka日志异常")
	}

	return
}

/**
发送kafka日志
*/
func (this *KafkaLogHook) SendKafkaMessage(topic string, message string) bool {
	success := false
	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	err := kafka2.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		HookLogger.Error("发送kafka日志, 异常", zap.Any("topic", topic), zap.Any("message", message), zap.Error(err))
	} else {
		HookLogger.Info("发送kafka日志, 完成", zap.Any("topic", *m.TopicPartition.Topic), zap.Any("partition", m.TopicPartition.Partition), zap.Any("offset", m.TopicPartition.Offset))
		success = true
	}

	close(deliveryChan)

	return success
}
