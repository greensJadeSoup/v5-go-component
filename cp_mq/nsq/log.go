package cp_mq_nsq

import (
	"github.com/greensJadeSoup/v5-go-component/cp_mq"
)

type logger struct {
	ilog cp_mq.ILog
}

// 实现nsq日志接口
func (l *logger) Output(calldepth int, s string) error {
	if calldepth == 3 {
		l.ilog.Error("NSQ错误："+s, true)
	}

	return nil
}
