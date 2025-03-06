package main

import (
	"os_prj/TreadPool"

)

func main() {
       
       a:=TreadPool.NewTredpool(5 , 20)
	   a.PoolManager( "arriving_time1.txt" , "burst_time1.txt")


	//   b:=TreadPool.NewTredpool(5 , 10)
	//   b.PoolManager( "arriving_time2.txt" , "burst_time2.txt")
	
	//   c:=TreadPool.NewTredpool(8 , 5)
	//   c.PoolManager( "arriving_time3.txt" , "burst_time3.txt")

}
