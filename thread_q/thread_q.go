package thread_q


import (
	//"runtime"
	//"sync/atomic"
	"os_prj/Mutex"
	//"time"
)

type Task struct {
	ID           int       // Task ID
	ExecutionTime int // Simulated execution time
	Bursttime	int
}
type TaskQueue struct {

	tasks []Task
	mutex Mutex.Mymutex
}
func (tq *TaskQueue)Set_id(num []int) () {
	for i:= 0; i<len(num); i++ {
		tq.tasks[i].ID = num[i]
	}
	
}
func (tq *TaskQueue)Set_burst(num []int) () {
	for i:= 0; i<len(num); i++ {
		tq.tasks[i].Bursttime = num[i]
	}
	
}
func (tq *TaskQueue) Set_arriving(num []int) () {
	for i:= 0; i<len(num); i++ {
		tq.tasks[i].ExecutionTime = num[i]
	}
	
}
// Enqueue adds a task to the queue
func (tq *TaskQueue) Enqueue(task Task) {
	tq.mutex.Lock()
	defer tq.mutex.Unlock()
	tq.tasks = append(tq.tasks, task)
}

// Dequeue removes and returns a task from the queue
func (tq *TaskQueue) Dequeue() *Task {
	tq.mutex.Lock()
	defer tq.mutex.Unlock()
	if len(tq.tasks) == 0 {
		return nil // Queue is empty
	}
	task := tq.tasks[0]
	tq.tasks = tq.tasks[1:]
	return &task
}


