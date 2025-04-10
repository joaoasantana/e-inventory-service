package conn

import (
	"github.com/IBM/sarama"
	"github.com/joaoasantana/e-inventory-service/pkg/config"
)

func KafkaConsumer(config config.KafkaInfo) sarama.Consumer {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true

	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(config.Brokers, kafkaConfig)
	if err != nil {
		panic(err)
	}

	return consumer
}

func KafkaProducer(config config.KafkaInfo) sarama.SyncProducer {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true

	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(config.Brokers, kafkaConfig)
	if err != nil {
		panic(err)
	}

	return producer
}
