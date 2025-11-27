Keep in mind that mutex is just an formal arrangement between coroutines 
to avoid race conditions when trying to access the same variables or data

So say when coroutine1 uses a mutex and then 
call lock, then everything between the lock 
and unlock is said to be changed only by coroutine1 
at that time. 

However some coroutine can still access the variable
even without using mutex and thus lead to race condition and 
non-deterministic results. 

As a result, mutex and lock unlock is just a way of saying 
let's coordinate by using lock and unlock when dealing with shared data
It is not binding and coroutines can break it and cause race condition 

=======

A mutex is a coordination tool, not an automatic safety mechanism.

Inside a locked section, only the locking goroutine can modify the protected data.

Another goroutine can bypass the mutex and touch the variable, but doing so breaks correctness.

The mutex protects only the code that actually uses it.


The mutex protects the entire system only if all goroutines cooperate.
This is like traffic lights:
A red light doesn’t physically block your car.
It works only if everyone follows it.
If someone runs the red light, the system breaks.
Mutexes are exactly this: cooperative concurrency control.

=====
so say coroutine1 acquiures a lock and then modify a global vvariable age, another coroutine doesn't acqire the mutex, but just tries updating age , it will still be able to. If another coroutine also agrees to be a good citixen by using lock unlock,, then that coroutine won't be able to get that lock during which the previous service is unlocked,, buut canmake the changes afterwards

=====
So with cond, the waiting and the signaling coroutine MUST use the same mutex 
So when the coroutine with cond.wait is called, mutex is unlocked
now this enables the mutex to be acquired by the other coroutine 
and then when cond.signal is done , then?

=====
Step-by-step:

Waiter acquires mutex → sits in the chair.

Waiter calls Wait() → releases mutex (stands up) and sleeps.

Signaler acquires mutex → sits, signals the waiter.

Signaler releases mutex → leaves the chair.

Waiter wakes up, reacquires mutex (sits again) → continues safely.

Waiter eventually releases mutex → leaves chair.
