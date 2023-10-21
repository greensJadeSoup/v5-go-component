package cp_mq_rmq

import (
	"v5-go-component/cp_mq"
)

func init() {
	cp_mq.ConsumerRegister("rocketmq", &RConsumer{})
	cp_mq.ProducerRegister("rocketmq", &RProducer{})
}
