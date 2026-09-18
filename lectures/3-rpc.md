RPC:
* What are the challenges of using RPC?
  * Timeouts
    * Server crashes or client crashes
* Unavoidable differences between RPC and local function calls?
  * TODO
* RPC
  * Arguments are put into a packet and sent through the network
  * the client-side and server-side stubs are simply trying to hide that we are calling a function remotely
* What are the challenges of implementing RPC framework?
  * The client and server need a shared version of the interface
  * Different languages store the different types in different sizes -> so it's language dependent
    * this makes the implementation of thr RPC framework challenging
  * Endianness issues on different machines

-------
TODO
-------


* Two different semantics:  
  * At-least-once
    * (1) The client sends, the server processes, but the response is lost. (2) The client then sends the response again, then the server processes again, and the response is sent to the client and the client now receives it.
  * At-most-once
    * We will send the request AT MOST ONCE (or maybe even zero times)
    * Client assigns unique id for each RPC
    * If RPC raises an exception -> zero or once (don't know which one)
* Is it possible to have exactly once?
  * TODO
* Should we use at-most-once or at-least-once?
  * Easier to use at-least-once
  *  