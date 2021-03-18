package threadpool

import "errors"

var (
	// ErrPoolIsFull pool is full
	ErrPoolIsFull = errors.New("pool is full")
)

// threadPool thread pool
type threadPool struct {
	pool chan func()
}

// NewThreadPool init
func NewThreadPool(size int) ThreadPool {
	p := &threadPool{
		pool: make(chan func(), size),
	}
	for i := 0; i < size; i++ {
		go p.run()
	}
	return p
}

func (p *threadPool) run() {
	for {
		select {
		case f := <-p.pool:
			f()
		}
	}
}

// Go 执行函数，如果线程池已满则返回失败
func (p *threadPool) Go(f func()) error {
	select {
	case p.pool <- f:
		return nil
	default:
		return ErrPoolIsFull
	}
}

// Run 执行函数，如果线程池已满则等待
func (p *threadPool) Run(f func()) {
	p.pool <- f
}
