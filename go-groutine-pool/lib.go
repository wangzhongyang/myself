package go_groutine_pool

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const mutexLocked = 1 << iota

type Mutex struct {
	in sync.Mutex
}

func (m *Mutex) Lock() {
	m.in.Lock()
}
func (m *Mutex) Unlock() {
	m.in.Unlock()
}
func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.in)), 0, mutexLocked)
}

type Worker struct {
	Url  string
	work func(url string) error
	l    *Mutex
}

type Pool struct {
	freeSignal chan struct{}
	workers    []*Worker
}

func NewPool(urls []string, f func(url string) error) (*Pool, error) {
	p := new(Pool)
	if len(urls) == 0 {
		return nil, errors.New("url number failed")
	}
	p.freeSignal = make(chan struct{}, len(urls))
	for _, url := range urls {
		w := &Worker{
			Url: url,
			l:   new(Mutex),
		}
		w.work = func(url string) error {
			defer func() {
				w.l.Unlock()
				p.freeSignal <- struct{}{}
			}()

			return f(url)
		}
		p.workers = append(p.workers, w)
		p.freeSignal <- struct{}{}
	}
	return p, nil
}

func (p *Pool) Run() {
	for {
		select {
		case <-p.freeSignal:
			for _, w := range p.workers {
				if w.l.TryLock() {
					go w.work(w.Url)
					break
				}
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}
