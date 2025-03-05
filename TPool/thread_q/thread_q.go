package thread_q

import (
	//"runtime"
	//"sync/atomic"
	"errors"
	"os_prj/Mutex"
	//"time"
)

type Task struct {
	TaskNum   int
	Arrtime   int // Simulated execution time
	Bursttime int
}
type TaskQueue struct {
	Tasks  []Task
	Qmutex Mutex.Mymutex
}

//read from file and add the task id
// func (tq *TaskQueue) Set_id(num []int) () {

// 	tq.tasks = make([]Task, len(num))
// 	for i:= 0; i<len(num); i++ {
// 		tq.tasks[i].ID = num[i]
// 	}

// }
// //read from file and add the task burst time

// func (tq *TaskQueue)Set_burst(num []int) () {
// 	for i:= 0; i<len(num); i++ {
// 		tq.tasks[i].Bursttime = num[i]
// 	}

// }
// //read from file and add the task arriving time

// func (tq *TaskQueue) Set_arriving(num []int) () {
// 	for i:= 0; i<len(num); i++ {
// 		tq.tasks[i].ExecutionTime = num[i]
// 	}

// }
// Enqueue adds a task to the queue
func Newqueue() *TaskQueue {
	var Ntq TaskQueue
	//Ntq.Tasks = make([]Task, 0)
	Ntq.Qmutex.Reset()

	return &Ntq
}

func (tq *TaskQueue) Enqueue(task Task) {
	tq.Qmutex.Lock()
	defer tq.Qmutex.Unlock()
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
