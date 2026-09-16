* Thread vs. process
  * A thread runs inside a process
  * You can run multiple threads inside the same process
  * Different threads share the same memory
* ADVANTAGES of using threads
  1. I/O concurrency
     * Can run another task while we are waiting for one to complete, especially input/output tasks since we know that these take time
  2. Multicore performance
  3. Convenience
* DISADVANTAGES of using threads
  1. Races
     * Two different threads reading shared memory but they have different values than what it should actually be
     * "Two threads use the same memory at the same time and at least one writes"
     * Difficult to reproduce
     * We don't want this AT ALL
  2. Deadlock
     * Two threads each require information from the other so the program simply stops executing
  3. Livelock
  4. Bad performance
     * there is overhead with regards to using threads
* Human analogy between deadlock and livelock
  * Deadlock: both people are trying to go through the doorway at the same time
  * Livelock: the two people are trying to get through, except they both try at the same time, then let the other person pass at the same time
* Is there an alternative to using multiple threads?
  * YES: "Write code that explicitly interleaves activities in a single thread - usually called 'event-driven'"
* Go is very convenient when using threads
  * It prevents many things that may happen compared to when using C
* C is often more difficult compared to using Go
* `-race` is a flag that can be used in order to determine where race conditions may be occurring
* **Channels** are very important
  * They are more powerful than condition variables
* If we are writing into 


Action Items:
* Complete the tutorials here:
  * https://go.dev/tour/concurrency/2
  * https://go.dev/tour/concurrency/10
* Understand the exercise at the end about consumer/producer
* Install the correct version of Go