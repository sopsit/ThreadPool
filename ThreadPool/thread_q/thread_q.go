package thread_q

import (
	"errors"
	"os_prj/Mutex"
	"time"
)

type Task struct {
	TaskNum   int
	Arrtime   int 
	Bursttime int
}
type TaskQueue struct {
	Tasks  []Task
	Qmutex Mutex.Mymutex
	Qsize  int
}

func Newqueue(size int) *TaskQueue {
	var Ntq TaskQueue
	Ntq.Tasks = make([]Task, 0)
	Ntq.Qmutex.Reset()
    Ntq.Qsize = size
	return &Ntq
}

func (tq *TaskQueue) Enqueue(task Task) {
	tq.Qmutex.Lock()
	defer tq.Qmutex.Unlock()
	for len(tq.Tasks) >= tq.Qsize {
		tq.Qmutex.Unlock()
		time.Sleep(100*time.Millisecond)
		tq.Qmutex.Lock()
	}
	tq.Tasks = append(tq.Tasks, task)
}

// Dequeue removes and returns a task from the queue
func (tq *TaskQueue) Dequeue() (*Task, error) {
	tq.Qmutex.Lock()
	defer tq.Qmutex.Unlock()
	if len(tq.Tasks) <= 0 {
		return nil, errors.New(" Queue is empty. ") // Queue is empty
	}
	task := tq.Tasks[0]
	tq.Tasks = tq.Tasks[1:]
	return &task, nil
}
