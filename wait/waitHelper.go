package wait

import (
	"time"
)

type WaitHelper struct{}

func NewWaitHelper() *WaitHelper {
	return &WaitHelper{}
}

func (waitHelper *WaitHelper) WaitForMilliSeconds(millis int) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}
