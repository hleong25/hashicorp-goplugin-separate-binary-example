package basepluginimpl

import (
	"time"

	"lahenry.com/library/baseplugin/greeter"
)

type greeterImpl struct {
	startTime time.Time
}

var greeterStartTime = time.Now()

func Greeter() greeter.IGreeter {
	return &greeterImpl{
		startTime: greeterStartTime,
	}
}

func (p *greeterImpl) Name() string {
	return "plugin2"
}

func (p *greeterImpl) Version() string {
	return "2.0.0.0"
}

func (p *greeterImpl) StartTime() int64 {
	return p.startTime.Unix()
}
