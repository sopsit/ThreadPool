package TreadPool

import (
	"fmt"
	"os_prj/Mutex"
	"os_prj/read"
	"os_prj/thread_q"
	"time"
)

type MyTreadPool struct {
	Tqueue     *thread_q.TaskQueue
	Workersnum int
	Taskdone   Mutex.Mymutex
	Stop       bool
	TaskCount  int
	Thread_work int
	Thread_work_lock Mutex.Mymutex
}

func NewTredpool(workerCount int) *MyTreadPool {

	var temptreadpool MyTreadPool
	temptreadpool.Tqueue = thread_q.Newqueue()
	temptreadpool.Workersnum = workerCount
	temptreadpool.Stop = false
	temptreadpool.TaskCount = 0
	temptreadpool.Thread_work = 0
	temptreadpool.Thread_work_lock.Reset()
	temptreadpool.Taskdone.Reset()

	return &temptreadpool

}

func (tp *MyTreadPool) ReadFromFiles(f1 string, f2 string, f3 string) ([]thread_q.Task, error) {

	var addtask []thread_q.Task
	Tasknumber, err := read.ReadFromFile(f1)
	if err != nil {
		return nil, err
	}
	artime, err1 := read.ReadFromFile(f2)
	if err1 != nil {
		return nil, err1
	}
	btime, err2 := read.ReadFromFile(f3)
	if err2 != nil {
		return nil, err2
	}
	for i := 0; i < len(Tasknumber); i++ {
		var temp thread_q.Task
		temp.TaskNum = Tasknumber[i]
		temp.Arrtime = artime[i]
		temp.Bursttime = btime[i]
		addtask = append(addtask, temp)

	}

	return addtask, nil
}

func (tp *MyTreadPool) AddToQueue(file1 string, file2 string, file3 string) error {

	treadtask, err := tp.ReadFromFiles(file1, file2, file3)
	if err != nil {
		return err
	}
	//fmt.Println(treadtask)
	count := 0
	ntime := time.Now()
	for _, tsk := range treadtask {
		tp.Taskdone.Lock()
		tp.TaskCount++
		tp.Taskdone.Unlock()
		dealy := time.Duration(tsk.Arrtime) * time.Second
		addtime := ntime.Add(dealy)
		time.Sleep(time.Until(addtime))
		tp.Tqueue.Enqueue(treadtask[count])
		fmt.Printf(" Task %v arrived \n", tsk.TaskNum)
		count++
	}
	return nil
}

func (tp *MyTreadPool) Work() {

	flag := false
	for {

		Tsk, err := tp.Tqueue.Dequeue()
		if err == nil {

			fmt.Printf(" Task %v started \n", Tsk.TaskNum)
			time.Sleep(time.Duration(Tsk.Bursttime) * time.Second)
			fmt.Printf(" Task %v finished \n", Tsk.TaskNum)
			if !flag { 

				tp.Thread_work_lock.Lock()
				tp.Thread_work ++
				tp.Thread_work_lock.Unlock()
				flag = true
			}
			tp.Taskdone.Lock()
			tp.TaskCount--
			tp.Taskdone.Unlock()
		} else {

			if tp.Stop {
				return
			}

			time.Sleep(5 * time.Second)
		}

	}
}

func (tp *MyTreadPool) Wait() {

	time.Sleep(10 * time.Second)

	for {
		tp.Taskdone.Lock()
		if tp.TaskCount == 0 {
			tp.Taskdone.Unlock()
			tp.Stop = true
			time.Sleep(10 * time.Second)
			break
		}
		tp.Taskdone.Unlock()
		time.Sleep(1 * time.Second)

	}
}

func (tp *MyTreadPool) PoolManager(F1 string, F2 string, F3 string) {

	go tp.AddToQueue(F1, F2, F3)

	for i := 0; i < tp.Workersnum; i++ {
		go tp.Work()
	}

	tp.Wait()
	fmt.Println("Done")
	fmt.Println(tp.Thread_work)

}

func (tp *MyTreadPool) Print() {
	// for i:= 0 ; i< len(tp.Tqueue.Tasks)  ; i++ {
	// 	fmt.Println(tp.Tqueue.Tasks[i]) }
	fmt.Println(tp.TaskCount)
}
