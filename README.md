# ***Thread Pool***


## Table of Contents
1. [Introduction](#introduction)
2. [Project Structure](#ProjectStructure)
3. [Packages](#Packages)
4. [Explanation](#Explanation)
5. [Usage](#Usage)

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
## Packages
- main -> Begening of program, Create Thread pool
- TreadPool -> defenition of TreadPool, Creating function, Reading from files (by calling from the other package), Work function, Wait Function and managing function for handling the tasks finishing.
- Mutex -> override the Lock and Unlock method
- read -> read the single file, which has a single string in each line
- thread_q -> defenition of safe queue using the mutex


## Explanation

first of all, in the main function, a defined as a Thread pool, the files are read and store in structure. after reading from files, if the tasks arriving time has passed, it being Enqueued. in Work() function, the threads tasks are dequeued. it means that task is executing. in output, id of task, starting and finishing time, total time of proceccing and the number of Workers will be showed.

- tips:
  - The arriving file is sorted
  - The mutex that has been wrote in mutex file has used
  - Handling situations where the entry time has not yet arrived  

## Usage

1. **Clone the Repository:** `https://github.com/sopsit/ThreadPool.git`
2. **Cd into the Tpool Directory** `ThreadPool/Tpoll`
3. **run this command to See the result** ‍‍‍‍
   
   ```
   go mod tidy
   go run main.go
   ```
