
package main

import (
	"fmt"
	"os_prj/read"
	"os_prj/thread_q"
	//"sync"
)

func main() {

	ids, err := read.readFromFile("id.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	burst_times, err := read.readFromFile("burst_time.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	arriving_times, err := read.readFromFile("arriving_time.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var tq thread_q.TaskQueue
	tq.Set_id(ids)
	tq.Set_burst(burst_times)
	tq.Set_arriving(arriving_times)
}
