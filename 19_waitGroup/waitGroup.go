/*
=============================================================================
                           🚀 GO WAITGROUP TUTORIAL
=============================================================================

📚 CORE CONCEPT:
WaitGroup is Go's way to wait for multiple goroutines to finish their work.
Think of it as a "counter" that tracks how many tasks are still running.

🔑 KEY METHODS:
• Add(n)  → Increases counter by n (usually 1)
• Done()  → Decreases counter by 1 (task finished)
• Wait()  → Blocks until counter reaches 0 (all tasks done)

💡 REAL-WORLD ANALOGY:
Imagine you're a manager waiting for your team to finish their tasks:
- Add(1): "I'm giving you one more task"
- Done(): "I finished my task"
- Wait(): "I'll wait here until everyone is done"

⚠️  IMPORTANT RULES:
1. Always call Add() BEFORE starting the goroutine
2. Always call Done() when goroutine finishes (use defer!)
3. Call Wait() from main thread to wait for all goroutines

=============================================================================  */


package main

import (
	"fmt"
	"sync"
)

// WaitGroup helps coordinate multiple goroutines:
// - Acts like a counter to track running goroutines
// - Add(1): Increases counter when starting a goroutine
// - Done(): Decreases counter when goroutine finishes  
// - Wait(): Blocks until all goroutines complete (counter = 0)

// task() runs a single unit of work:
// - Takes an ID to identify the task
// - Uses defer to ensure cleanup happens
// - Prints when task is running
func task(id int, w *sync.WaitGroup) {
	defer w.Done()  
	fmt.Println("Running task", id)
}

func main() {
	// Create a WaitGroup to coordinate all our goroutines
	var wg sync.WaitGroup

   for i := 0; i <=10 ; i++ {
	// Tell WaitGroup we're adding a new task
	wg.Add(1)
	// Start task in background thread (goroutine)
	go task(i, &wg)
   }

   // Pause here until all tasks have finished
   wg.Wait()
}
		


/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🔍 WHAT HAPPENS WITHOUT WAITGROUP?
Without WaitGroup, main() would exit immediately, killing all goroutines
before they finish their work!

📊 EXECUTION FLOW:
1. Main creates WaitGroup (counter = 0)
2. For each worker: Add(1) → counter++
3. Launch goroutine → worker starts running
4. Main calls Wait() → blocks until counter = 0
5. Each worker calls Done() → counter--
6. When counter = 0 → Wait() unblocks → main continues

🎯 COMMON MISTAKES:
❌ Calling Add() inside goroutine (race condition)
❌ Forgetting to call Done() (deadlock)
❌ Calling Add() after Wait() (undefined behavior)
❌ Not using defer for Done() (might skip if panic)

✅ BEST PRACTICES:
• Always Add() before launching goroutine
• Always use defer wg.Done()
• Pass WaitGroup by pointer (&wg)
• One WaitGroup per logical group of tasks

🚀 WHEN TO USE WAITGROUP:
• Parallel processing of independent tasks
• Fan-out/fan-in patterns
• Coordinating multiple API calls
• Batch processing operations

=============================================================================
*/
		
