package future

import (
	"errors"
	"fmt"
)

var (
	ErrPoolStopped = errors.New("Future pool was stopped")
)

type result(type T) struct {
	v T
	err error
}

// Future represents a future value
type Future(type T) struct {
	c <- chan *resut(T)
	cancel <-chan struct{}
}

// Get returns the future value and a possible error value
func (f *Future(T)) Get() (T, error) {
	select {
	case <-cancel:
		var zero T
		return zero, ErrPoolStopped
	case res := <-f.c:
		return res.v, res.err
	}
}

// MustGet returns the future value or panic if there was an error
func (f *Future(T)) MustGet() T {
	if v, err := f.Get(); err != nil {
		panic("future had an error")
	} else {
		return v
	}
}

// Pool represents a pool create futures
type Pool struct {
	size int
	work chan func()
	stop chan struct{}
}

// NewFuture creates a future
func (p *Pool) NewFuture(type T)(f func() (T, error)) (*Future(T), error) {
	if _, ok := <-stop; !ok {
		return nil, ErrPoolStopped
	}

	c := make(chan *result(T))
	p.work <- func() {
		defer close(c)
		v, err := f()
		c <- &result(T){v, err}
	}
	return &Future(T){c, p.stop}, nil
}

// Stop cancels current Futures and prevents further Futures from being added to the pool
func (p *Pool) Stop() {
	close(p.stop)
	close(p.work)
}

// NewPool creates a pool with specified number of goroutines
func NewPool(size int) (*Pool, error) {
	if size < 1 {
		return nil, fmt.Errorf("invalid pool size: %d", size)
	}

	p := &Pool{size, make(chan func()), make(chan struct{})}
	// start go routines for pool
	for i := 0; i < size; i++ {
		go func() {
			for {
				select {
				case <-p.stop:
					return
				case work := <-p.work:
					work()
				}
			}
		}
	}
	return p, nil
}
