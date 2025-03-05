package main

import (
	//"fmt"
	//"fmt"
	"os_prj/TreadPool"
	//"os_prj/thread_q"
	//"sync"
	//"testing"
	//"os_prj/Mutex"
	//"time"
)

func main() {

	// ids, err := read.ReadFromFile("id.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// burst_times, err := read.ReadFromFile("burst_time.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// arriving_times, err := read.ReadFromFile("arriving_time.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// var tq thread_q.TaskQueue
	// tq.Set_id(ids)
	// tq.Set_burst(burst_times)
	// tq.Set_arriving(arriving_times)
	// // fmt.Println(ids, " ")
	// // fmt.Println(burst_times, " ")
	// // fmt.Println(arriving_times," ")

	// // Initialize Thread Pool
	
	// // we can get the number - at present i put 5 
	// pool := Thread_pool.NewThreadPool(5, &tq)
	// pool.Start()

	// // Assign tasks
	// go pool.AssignTasks()

	// // Run simulation for 5 seconds
	// time.Sleep(5 * time.Second)
	// pool.Stop()

// 	// fmt.Println("Simulation Complete!")
//    var arr []int
//    arr = append(arr, 1)
//    arr = append(arr, 1)
//    arr = append(arr, 1)
//    arr = append(arr, 5)
// 	stime := time.Now()
// 	for _ , t := range arr {
// 		dtime := time.Duration(t)*time.Second
// 		addtime := stime.Add(dtime)
// 		fmt.Println(addtime.Format(time.RFC1123))
// 		time.Sleep(time.Until(addtime))
// 		fmt.Println(t)
// 	}
      a := TreadPool.NewTredpool(5)
	  //err := a.AddToQueue("id.txt" , "arriving_time.txt" , "burst_time.txt")
	 // if err != nil {fmt.Println(err.Error())}
	 a.PoolManager("id.txt" , "arriving_time.txt" , "burst_time.txt")
	// a.Print()
}


