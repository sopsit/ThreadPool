# ***Thread Pool***

In this project we define the safe queue, override the mutex, create the Thread pool and managing the task by using FIFO algorithm.

# Project Structure
```
ThreadPool/ TPool
├── Mutex/                # Contains implitation of Mutex
    ├── Mymutex.go
├── ThreadPool/           # Contains implitation of Thread pool and doing the tasks
    ├──TP.go
├── read/                 # Reading from files
    ├── read.go
├── thread_q/             # Contains implitation of safe queue using overrided mutex
    ├── thread_q.go
├── arriving_time.txt              
├── burst_time.txt                        
├── id.txt               
└── main.go                # Main of the project
```
# Usage
run this command to See the result
```
go run main.go
```
