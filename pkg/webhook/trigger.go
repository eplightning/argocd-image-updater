package webhook

import "time"

type UpdateTrigger struct {
	interval time.Duration
	notify   chan struct{}
}

func NewUpdateTrigger(interval time.Duration) *UpdateTrigger {
	return &UpdateTrigger{
		interval: interval,
		notify:   make(chan struct{}, 1),
	}
}

func (t *UpdateTrigger) Listen() <-chan struct{} {
	ch := make(chan struct{}, 1)

	go func() {
		ch <- struct{}{}
		ticker := time.NewTicker(t.interval)

		for {
			select {
			case <-t.notify:
			case <-ticker.C:
			}

			ch <- struct{}{}
		}
	}()

	return ch
}

func (t *UpdateTrigger) Trigger() {
	select {
	case t.notify <- struct{}{}:
	default:
	}
}
