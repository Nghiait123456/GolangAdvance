package main

//
//func init() {
//	// this is bad, dont use goroutine in init, init when have many logic and order involve, easy to freeze code for maintain
//	go doWork()
//}
//
//func doWork() {
//	// this is bad, dont stop go routine
//	for {
//		// ...
//	}
//}

//type Worker struct{ /* ... */ }
//
//func NewWorker(...) *Worker {
//	w := &Worker{
//		stop: make(chan struct{}),
//		done: make(chan struct{}),
//		// ...
//	}
//	go w.doWork()
//}
//
//func (w *Worker) doWork() {
//	defer close(w.done)
//	for {
//		// ...
//		case <-w.stop:
//		return
//	}
//}
//
//// Shutdown tells the worker to stop
//// and waits until it has finished.
//func (w *Worker) Shutdown() {
//	close(w.stop)
//	<-w.done
//}
