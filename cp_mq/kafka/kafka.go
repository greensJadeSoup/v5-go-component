package cp_mq_kafka

import (
	"github.com/greensJadeSoup/v5-go-component/cp_mq"
)

func init() {
	cp_mq.ConsumerRegister("kafka", &KConsumer{})
	cp_mq.ProducerRegister("kafka", &KProducer{})
}
