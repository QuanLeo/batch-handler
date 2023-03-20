package batchhandler

import (
	"time"
)

type ThrottHandler func([]any) error

type BatchHandler struct {
	c           chan []any
	terminate   bool
	duration    time.Duration //period time execute
	number      int           //number of items execute
	funcHandler ThrottHandler //function execute after period time or greater than number
	data        []any         //data will be execute
}

/** Init new batch handler
 * duration: period time execute
 * number: number of items execute
 * funcHandler: function handler events
 */
func New(duration time.Duration, number int, funcHandler ThrottHandler) *BatchHandler {
	if duration == 0 && number == 0 {
		panic("Duration or number must be greater 0")
	}

	instance := BatchHandler{duration: duration, number: number, funcHandler: funcHandler, c: make(chan []any), terminate: false}
	go instance.execute()
	return &instance
}

func (t *BatchHandler) Push(data any) {
	t.data = append(t.data, data)
}

func (t *BatchHandler) numTrigger() {
	for {
		if t.terminate {
			return
		}

		if len(t.data) >= t.number {
			t.trigger()
		}
	}
}

func (t *BatchHandler) trigger() {
	if t.data == nil {
		return
	}

	data := t.data
	t.data = nil
	go t.funcHandler(data)
}

func (t *BatchHandler) timeTrigger() {
	for {
		time.Sleep(t.duration)
		if t.terminate {
			return
		}

		t.trigger()
	}
}

func (t *BatchHandler) execute() {
	defer t.trigger()

	if t.number > 0 {
		go t.numTrigger()
	}

	if t.duration > 0 {
		go t.timeTrigger()
	}
}

func (t *BatchHandler) Terminate() {
	t.trigger()
	t.terminate = true
}
