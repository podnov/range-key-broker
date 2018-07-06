package pkg

import (
	"errors"
	"fmt"
	"github.com/go-redsync/redsync"
)

type broker struct {
	mutexes map[string]*redsync.Mutex
}

func (b *broker) CheckIn(rangeValue string) (err error) {
	mutex := b.mutexes[rangeValue]

	succeeded := mutex.Unlock()

	if !succeeded {
		message := fmt.Sprintf("Unlock failed for [%s]", rangeValue)
		err = errors.New(message)
	}

	return err
}

func (b *broker) CheckOut() (string, error) {
	var acquiredMutex *string
	var err error

	for acquiredMutex == nil {
		for name, mutex := range b.mutexes {
			err := mutex.Lock()

			if err == nil {
				acquiredMutex = &name
				break
			}
		}
	}

	return *acquiredMutex, err
}

func NewBroker(rangeValues []string) broker {
	mutexes := createMutexes(rangeValues)

	return broker{
		mutexes: mutexes,
	}
}
