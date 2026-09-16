* What is a distributed system?
  * "A distributed system in which the failure of a computer you didn't even know existed can render your own computer unusable" -> main issue of a distributed system
  * A collection of nodes connected between each other by a network -> they all collaborate with each other to provide a certain service
  * In a distributed system, there is no shared memory or clock (each of the different nodes have their own)
* Examples of distributed systems
  * Peer-to-peer systems
  * Any program where you are using APIs
    * If the API doesn't work, then our application won't work (since we are relying on the API)
  * DNS is considered a distributed system
* History of distributed system:
  * 70s: Xerox PARC
  * 80s: DEC SRC
* RPC (remote procedure call): key idea in distributed systems
* We need systems to be able to work even when servers go down since this is inevitable
* Why do we still used distributed systems even though there are risks
  * Fault tolerance: If we have everything in one place, then if something happens, then it doesn't bring down the entire system
  * Performance: faster when we run things on multiple different machines
  * Scaling: you often need to have multiple nodes
  
    ---
  * Service architecture: split responsibility into components that talk over the network, not the local calls
  * Scalability: add machines instead of hitting the ceiling of one bigger machine
  * Availability: Survive the failure of any single machine by replicating
  * Geographic Locality: Keep data 
* Biggest reason for delay: propagation delay
  * When servers are very far from one another, there is nothing we can do about it to make it faster
  * We can however add more machines around the world so that the information is closer to the intended place where it will be used
* Recurring problems with distributed systems:
  1. Programming model - communication
     * No shared memory
     * Network delay with RPC
     * Failure ambiguity (what actually happened)
  2. Programming mode - synchronization
     * No shared clock or memory (no atomic ordering)
     * No global order (no shared clock)
  3. Replication and consistency
     * If a write lands on replica A but not yet on B -> a read from B is stale
     * How do we order different events (e.g. do we apply interest before or after a certain sum of money was added)
  4. Event ordering and consistency
     * Strong consistency = coordination on every write
  5. Failure and availability
     * Node failure or network failure
     * CAP theorem: strong consistency vs. strong availability
       * There is a tradeoff between choosing one of these. You cannot have both
    ---
  * If we have a distributed system, there are guaranteed servers that will crash
  * A big issue is debugging -> we don't know where the issue is coming from
    * We may not be able to reproduce the issue given teh async nature of the problem
  * If one of th hosts is compromised, how do we ensure that the other nodes are not also compromised -> large security problems arise when developing distributed systems