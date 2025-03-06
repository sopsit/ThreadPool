package TreadPool

import (
	"fmt"
	"os_prj/Mutex"
	"os_prj/read"
	"os_prj/thread_q"
	"time"
)



type MyTreadPool struct {


	Tqueue    *thread_q.TaskQueue
    Workersnum   int
	Taskdone     Mutex.Mymutex
	Stop         bool
	TaskdoneCount   int
	TotalTask       int
	usedworker  int
	usedworkerlock Mutex.Mymutex
	Done chan  int

}


func NewTredpool (workerCount int , queuesize int) *MyTreadPool {


	 nthreadpool := &MyTreadPool {
	Tqueue : thread_q.Newqueue(queuesize) ,
    Workersnum  : workerCount ,
	Taskdone :   *Mutex.NMutex() ,
	Stop  :      false ,
	TaskdoneCount :   0 ,
	usedworker: 0,
	usedworkerlock: *Mutex.NMutex(),
	Done:  make(chan int) ,
  }

  return nthreadpool

}


func (tp *MyTreadPool) ReadFromFiles (f1 string , f2 string ) ( []thread_q.Task , error) {
	
    var addtask []thread_q.Task
	artime , err1 := read.ReadFromFile(f1)
	if err1 != nil { return nil , err1 }
	btime , err2 := read.ReadFromFile(f2)
    if err2 != nil { return nil ,err2 }
	for i:= 0 ; i< len(artime)  ; i++ {
	var temp thread_q.Task
	temp.TaskNum = i+1
	temp.Arrtime = artime[i]
	temp.Bursttime = btime[i]
	addtask = append(addtask , temp)
	}
     tp.TotalTask = len(addtask)
    return addtask , nil
}


func (tp *MyTreadPool) AddToQueue (file1 string , file2 string ) error{

	treadtask , err := tp.ReadFromFiles(file1 , file2 )
    if err != nil {return err}
	//fmt.Println(treadtask)
    count := 0
	ntime := time.Now()
	for _ , tsk := range treadtask {
		tp.Taskdone.Lock()
		tp.TaskdoneCount++
		tp.Taskdone.Unlock()
		dealy := time.Duration(tsk.Arrtime)*time.Second
    	addtime := ntime.Add(dealy)
		time.Sleep(time.Until(addtime))
		//fmt.Printf(" Task %v arrived \n" , tsk.TaskNum)
		tp.Tqueue.Enqueue(treadtask[count])
		count ++
		}

		tp.Done <- 1
    return nil

}



func (tp *MyTreadPool) Work() {

	flag := false

	for {
          	
	Tsk , err := tp.Tqueue.Dequeue()
	if err == nil {
		
		fmt.Printf(" Task %v started at %v \n" , Tsk.TaskNum , time.Now().Format(time.RFC1123))
		time.Sleep(time.Duration(Tsk.Bursttime)*time.Second)
		fmt.Printf(" Task %v finished at %v \n" , Tsk.TaskNum , time.Now().Format(time.RFC1123))
		if !flag{
			tp.usedworkerlock.Lock()
			tp.usedworker++
			tp.usedworkerlock.Unlock()
			flag=true
		}
       // tp.workermap[num]=1
		tp.Taskdone.Lock()
		tp.TaskdoneCount --
		tp.Taskdone.Unlock()
	
	} else {
     
		if tp.Stop {
             break
		}

	//	time.Sleep(100*time.Millisecond)
	}

	


	}
}

func (tp *MyTreadPool) Wait() {

	// time.Sleep(2 * time.Second) 
	
	<- tp.Done

	for {
		tp.Taskdone.Lock()
		if(tp.TaskdoneCount == 0){
			tp.Taskdone.Unlock()
			tp.Stop = true
			time.Sleep(2*time.Second)
			break
		}
		tp.Taskdone.Unlock()
		time.Sleep(2 * time.Second)
 

	}
}

func (tp *MyTreadPool) PoolManager(F1 string , F2 string ) {

	startTime := time.Now()

	go tp.AddToQueue(F1 , F2 )

	for i:= 0 ; i < tp.Workersnum ; i++ {
		go tp.Work()
	}

	tp.Wait()
	fmt.Println(" Done! " )
	endTime := time.Since(startTime)
	fmt.Printf(" Program execution time: %.2f \n TotalTasks: %v \n NumberofworkerThreads: %v \n NumberofUsedworkerThreads: %v \n Queuesize: %v \n" , 
	endTime.Seconds() , tp.TotalTask , tp.Workersnum , tp.usedworker , tp.Tqueue.Qsize )
  


}

