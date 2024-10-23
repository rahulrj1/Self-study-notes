Multithreading is a technique that allows for concurrent (simultaneous) execution of two or more parts of a program for maximum utilisation of a CPU.

Program: A program is an executable file like chrome.exe  Process: A process is an executing instance of a program. Each process is able to run concurrent subtasks called threads.

Thread: Thread is the smallest executable unit of a process. Threads can give the illusion of multitasking even though at any given point in time the CPU is executing only one thread  Concurrency v/s Parallelism: Concurrency is the ability of your program to deal (not doing) with many things at once and is achieved through multithreading. On the other hand, parallelism is about doing many things at once.

Context switching: Context switching is the technique where CPU time is shared across all running processes and is key for multitasking.  Thread Pool: A thread pool consists of homogenous worker threads that are assigned to execute tasks. Thread pools allow you to decouple task submission and execution. Once a worker thread finishes a task, it is returned to the pool. A thread pool can be fine-tuned for the size of the threads it holds which will help us to control the throughput of the system.

Lock: Locks are a synchronisation technique used to limit access to a resource in an environment where there are many threads of execution.  Mutex: Mutex implies mutual exclusion. A mutex allows only a single thread to access a resource.  Thread Safety: Thread safety is a concept that means different threads can access the same resources without exposing erroneous behaviour or producing unpredictable results like a race condition or a deadlock

Deadlock: Deadlocks happen when two or more threads aren’t able to make any progress because the resource required by the first thread is held by the second and the resource required by the second thread is held by the first. Necessary conditions are MHNC: Mutual exclusion, Hold & Wait, No pre-emption, Circular wait.

Critical Section: Critical section is any piece of code that has the possibility of being executed concurrently by more than one thread of the application and exposes any shared data or resources

Race condition: A race condition occurs in concurrent programming when two or more threads access shared data at the same time without thread synchronisation, and the final outcome depends on the timing of their execution. The operations being performed are non-atomic.
 Example of a rack condition (handled using mutex lock [It was giving wrong result without mutex lock])

<img width="1198" alt="Screenshot 2024-09-05 at 7 16 12 PM" src="https://github.com/user-attachments/assets/7d950ebb-2444-4d4f-9ebe-c57f64f25521">
￼
If mutex lock is used, final counter value would be 20000000 always.

Livelock: Livelock is a condition in concurrent programming where two or more threads or processes continuously change their states in response to each other without making any real progress. Unlike a deadlock, where processes are stuck waiting for each other, in livelock, the processes are active but they keep adjusting their actions to avoid a conflict, leading to an endless cycle of activity without achieving their goal.  Task: A task is a higher-level abstraction compared to threads. Tasks represent a unit of work that can be executed concurrently, but unlike threads, tasks do not require the programmer to directly manage
