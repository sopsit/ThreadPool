
package main

import (
	"fmt"
	"os_prj/read"
	"os_prj/thread_q"
	//"sync"
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
}
