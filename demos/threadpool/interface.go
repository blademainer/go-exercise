package threadpool

// ThreadPool 线程池
type ThreadPool interface {
	// Go 执行函数，如果线程池已满则返回失败
	Go(f func()) error

	// Run 执行函数，如果线程池已满则等待直到有新的线程被空出
	Run(f func())
}
