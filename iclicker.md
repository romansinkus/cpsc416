# CPSC 416 — iClicker Questions

## Question #1 — Lock Granularity

![iClicker Question 1: Lock Granularity](iclicker-assets/q1-lock-granularity.png)

**Question:** See the two versions of `safe_increment` on the right. Which of these is true?

**Version 1 — lock inside the loop**

```go
for i := 0; i < e7; i++ {
    safe_counter.mu.Lock()
    safe_counter.value++
    safe_counter.mu.Unlock()
}
```

**Version 2 — lock outside the loop**

```go
safe_counter.mu.Lock()
for i := 0; i < e7; i++ {
    safe_counter.value++
}
safe_counter.mu.Unlock()
```

**Options**

- **A.** Both versions are safe
- **B.** Version 1 is unsafe
- **C.** Version 2 is unsafe
- **D.** Version 1 is more efficient
- **E.** Version 2 is more efficient

**Answer: A and E**

- **A — both versions are safe.** In each, every mutation of `safe_counter.value` happens while the lock is held, so no two goroutines ever touch it concurrently. The final count is exactly `2e7` either way.
- **E — Version 2 is more efficient.** It acquires and releases the lock once, instead of 10 million times. Each `Lock`/`Unlock` pair costs an atomic operation, and under contention the losing goroutine parks and reschedules — so V1 pays that overhead on every iteration.

**Notes:**

- Safety and efficiency are separate axes here. B and C are wrong because both versions are correctly synchronized; D is wrong because V1 has strictly more lock overhead.
- The catch with V2: holding the lock for the entire loop fully serializes the goroutines — B cannot start until A finishes. It is faster in this microbenchmark precisely *because* it gives up all concurrency. In real code, a coarse lock held across a long critical section is usually the wrong call, especially if it spans I/O or an RPC.
- The general rule: hold a lock for as short a section as correctness allows, but don't reacquire it in a tight loop when one acquisition covers the whole operation.

---

## Question #2 — Scheduling Order

![iClicker Question 2: Scheduling Order](iclicker-assets/q2-scheduling-order.png)

**Question:** Given the creation order in `main()` (A, then B), which of these is true for **Version 2**?

```go
main()

go safe_increment("A")  // Thread A
go safe_increment("B")  // Thread B
time.Sleep(20 * time.Second)
```

**Version 2 — lock outside the loop**

```go
safe_counter.mu.Lock()
for i := 0; i < e7; i++ {
    safe_counter.value++
}
safe_counter.mu.Unlock()
```

**Options**

- **A.** Thread A will start (print "begin") first
- **B.** Thread B will start (print "begin") first
- **C.** Thread A will finish (print "end") first
- **D.** Thread B will finish (print "end") first
- **E.** None of the above

**Answer: E — none of the above**

When you launch a goroutine, you don't know exactly when it will start, and you don't know which one will run first. `go f()` only marks the goroutine as runnable; the Go scheduler decides when it actually gets a thread. Creation order in `main()` is not a promise about execution order.

**Notes:**

- **Why not A or B:** `fmt.Println(arg, ": begin")` happens *before* `mu.Lock()`, so the lock does nothing to order the two "begin" prints. They race, and either can win.
- **Why not C or D:** whichever goroutine acquires the lock first runs all `1e7` iterations while holding it and finishes first — but which one wins the lock is exactly the nondeterministic part. Go's `sync.Mutex` gives no FIFO or fairness guarantee on acquisition.
- **Seen in practice:** running the unsynchronized version of this program printed `B : begin` before `A : begin`, even though A was created first.
- **Takeaway:** if you need ordering between goroutines, you must create it explicitly — a channel, a `sync.WaitGroup`, or a mutex around the part you care about. Never infer ordering from the order of `go` statements.
