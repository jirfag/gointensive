package workers

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPoolNoWork(t *testing.T) {
	p := NewPool(1)
	p.Run()
	p.Stop()
}

func TestPoolOneTask(t *testing.T) {
	p := NewPool(1)
	p.Run()
	res := p.AddTaskSync(func() interface{} {
		time.Sleep(time.Millisecond * 100)
		return 7
	})
	assert.Equal(t, 7, res)
}

func TestPoolSize(t *testing.T) {
	p := NewPool(3)
	p.Run()
	wg := sync.WaitGroup{}
	for i := 0; i < p.Size(); i++ {
		wg.Add(1)
		go func(j int) {
			res, err := p.AddTaskSyncTimed(func() interface{} {
				wg.Done()
				time.Sleep(time.Second)
				return j
			}, time.Millisecond*100)

			assert.Nil(t, err)
			assert.Equal(t, j, res)
		}(i)
	}
	wg.Wait()

	res, err := p.AddTaskSyncTimed(func() interface{} {
		return 1
	}, time.Millisecond*100)

	assert.Equal(t, ErrJobTimedOut, err)
	assert.Nil(t, res)
}
