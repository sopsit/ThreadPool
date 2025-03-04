

package Thread_pool

import (
	"fmt"
	"os_prj/thread_q"
	"sync"
)

type Worker struct {
	ID       int
	TaskChan chan thread_q.Task
}

// ThreadPool structure
type ThreadPool struct {
	workers      []Worker
	taskQueue    *thread_q.TaskQueue
	workerCount  int
	wg           sync.WaitGroup
	stopChan     chan struct{}
}

// Initialize the thread pool
func NewThreadPool(workerCount int, tq *thread_q.TaskQueue) *ThreadPool {
	tp := &ThreadPool{
		workerCount: workerCount,
		taskQueue:   tq,
		stopChan:    make(chan struct{}),
	}

	for i := 0; i < workerCount; i++ {
		worker := Worker{
			ID:       i,
			TaskChan: make(chan thread_q.Task),
		}
		tp.workers = append(tp.workers, worker)
	}

	return tp
}

// Start workers
func (tp *ThreadPool) Start() {
	for i := range tp.workers {
		go tp.workerRoutine(&tp.workers[i])
	}
}

// Worker routine
func (tp *ThreadPool) workerRoutine(worker *Worker) {
	for {
		select {
		case task := <-worker.TaskChan:
			fmt.Printf("Worker %d executing task %d\n", worker.ID, task.ID)
			// Simulate execution
		case <-tp.stopChan:
			return
		}
	}
}

// Assign tasks to workers
func (tp *ThreadPool) AssignTasks() {
	for {
		task := tp.taskQueue.Dequeue()
		if task != nil {
			workerID := task.ID % tp.workerCount
			tp.workers[workerID].TaskChan <- *task
		}
	}
}

// Stop the thread pool
func (tp *ThreadPool) Stop() {
	close(tp.stopChan)
	tp.wg.Wait()
}
