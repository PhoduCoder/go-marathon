#Why Concurrency

Speeds up the code execution by organizing code into coroutines
which can run independently 
and use the fast CPU speed, which a single coroutine can't because of multiplexing 
in CPU 

Note that the CPU speed doesn't change irrespective of concurrency. We just utilize it more effectively 
The CPU is already fast.
By splitting tasks into multiple independent goroutines, you let the CPU keep working on other tasks while some are waiting (I/O, sleep, etc.).

This leverages the CPU’s speed more efficiently without changing the CPU itself.