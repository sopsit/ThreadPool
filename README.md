# ***Thread Pool***


## Table of Contents
1. [Introduction](#introduction)
2. [Project Structure](#ProjectStructure)
3. [Explanation](#Explanation)
4. [Usage](#Usage)

## Introduction

In this project we define the safe queue, override the mutex, create the Thread pool and managing the task by using FIFO algorithm.

## Project Structure
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
└── main.go                # Main of the project
```

## Explanation
first of all, in the main function, a defined as a Thread pool, the files are read and store in structure. after reading from files, if the tasks arriving time has passed, it being Enqueued. in Work() function, the threads tasks are dequeued. it means that task is executing. in output, id of task, starting and finishing time, total time of proceccing and the number of Workers will be showed.

- tips:
  - The size of queue is important and is considerd
  - The mutex that has been wrote in mutex file has used
  - Handling situations where the entry time has not yet arrived  

## Usage

1. **Clone the Repository:** `https://github.com/sopsit/ThreadPool.git`
2. **Cd into the Tpool Directory**
3. **run this command to See the result** `go run main.go`
