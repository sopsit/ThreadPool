
package main

import (
	"fmt"
	"os_prj/read"
	"os_prj/thread_q"
	//"sync"
	"os_prj/Thread_pool.go"
	"time"
)

func main() {

	ids, err := read.ReadFromFile("id.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	burst_times, err := read.ReadFromFile("burst_time.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	arriving_times, err := read.ReadFromFile("arriving_time.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var tq thread_q.TaskQueue
	tq.Set_id(ids)
	tq.Set_burst(burst_times)
	tq.Set_arriving(arriving_times)
	// fmt.Println(ids, " ")
	// fmt.Println(burst_times, " ")
	// fmt.Println(arriving_times," ")

	// Initialize Thread Pool
	
	// we can get the number - at present i put 5 
	pool := Thread_pool.NewThreadPool(5, &tq)
	pool.Start()

	// Assign tasks
	go pool.AssignTasks()

	// Run simulation for 5 seconds
	time.Sleep(5 * time.Second)
	pool.Stop()

	fmt.Println("Simulation Complete!")
}
