/*
=============================================================================
                           📡 GO CHANNELS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Channels are Go's way for goroutines to communicate and synchronize.
Think of them as "pipes" that safely pass data between concurrent functions.

🔑 KEY FEATURES:
• Thread-safe communication between goroutines
• Blocking operations (synchronization built-in)
• Directional channels (send-only, receive-only)
• Buffered vs unbuffered channels

💡 REAL-WORLD ANALOGY:
Channel = Postal Service
- Sender: Puts letter in mailbox (send operation)
- Receiver: Waits for letter to arrive (receive operation)
- Mailbox: Channel that holds the message
- Both sender and receiver must be ready for delivery!

🎯 WHY USE CHANNELS?
• Safe data sharing between goroutines
• Built-in synchronization (no need for locks)
• "Don't communicate by sharing memory; share memory by communicating"

=============================================================================
*/

// channels help provide communicatio between goroutines

package main

import (
	"fmt"
	"time"
)

// 📨 CHANNEL RECEIVER FUNCTION: Waits for data from channel
func processNum(numChan chan int){
  fmt.Println("Processing channel, received number : ", <-numChan)  // ⬅️ RECEIVE: Gets value from channel
  // <-numChan blocks until someone sends a value
}

// 📤 CHANNEL SENDER FUNCTION: Sends data to channel
func sendNumbers(numChan chan int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("📤 Sending: %d\n", i)
		numChan <- i  // ➡️ SEND: Puts value into channel
		time.Sleep(500 * time.Millisecond)
	}
	close(numChan)  // 🔒 CLOSE: Signal no more values coming
}

func main(){  // main function is a goroutine by default

	fmt.Println("🚀 CHANNELS DEMO")
	fmt.Println("================")

	// 🏗️ CREATE CHANNEL: make(chan type)
	numChan := make(chan int)  // Unbuffered channel (synchronous)

	// 🎯 DEMO 1: Basic Send/Receive
	go processNum(numChan)  // 2nd goroutine - waits for data
     
	//sending value to channel
	numChan <- 10  // ➡️ SEND: Main goroutine sends 10

	time.Sleep(1 * time.Second)  // wait for 1 second)

	fmt.Println("\n🎯 DEMO 2: Multiple Values")
	fmt.Println("==========================")

	// 🔄 DEMO 2: Sending multiple values
	numChan2 := make(chan int)
	go sendNumbers(numChan2)

	// 📥 RECEIVE in loop until channel is closed
	for num := range numChan2 {  // range automatically handles channel closing
		fmt.Printf("📥 Received: %d\n", num)
	}

	fmt.Println("\n🎯 DEMO 3: Buffered Channel")
	fmt.Println("===========================")

	// 🗂️ BUFFERED CHANNEL: Can hold multiple values without blocking
	bufferedChan := make(chan string, 2)  // Buffer size = 2
	
	bufferedChan <- "Hello"    // ✅ Doesn't block (buffer has space)
	bufferedChan <- "World"    // ✅ Doesn't block (buffer still has space)
	// bufferedChan <- "Third" // ❌ Would block (buffer full)

	fmt.Println("📥", <-bufferedChan)  // "Hello"
	fmt.Println("📥", <-bufferedChan)  // "World"

//  💡 COMMENTED EXAMPLE: Blocking behavior demonstration
//  messageChan := make(chan string)
//  messageChan <- "ping"  // ❌ DEADLOCK! No receiver ready
//    messageRecevied := <-messageChan
//    fmt.Println(messageRecevied)

	fmt.Println("\n✨ All channel demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📡 CHANNEL OPERATIONS:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Operation     │     Syntax      │           Behavior                  │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ Create          │ make(chan int)  │ Creates unbuffered channel          │
│ Send            │ ch <- value     │ Blocks until receiver ready         │
│ Receive         │ value := <-ch   │ Blocks until sender sends           │
│ Close           │ close(ch)       │ Signals no more values coming       │
│ Range           │ for v := range  │ Receives until channel closed       │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🔄 CHANNEL TYPES:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Unbuffered (synchronous)                                             │
│ ch := make(chan int)        // Send blocks until receive               │
│                                                                         │
│ // Buffered (asynchronous)                                              │
│ ch := make(chan int, 5)     // Can hold 5 values before blocking       │
│                                                                         │
│ // Directional channels                                                 │
│ var sendOnly chan<- int     // Can only send                           │
│ var recvOnly <-chan int     // Can only receive                        │
└─────────────────────────────────────────────────────────────────────────┘

⚡ BLOCKING BEHAVIOR:
• Unbuffered: Send blocks until receive, receive blocks until send
• Buffered: Send blocks only when buffer full, receive blocks when empty
• Closed channel: Receive returns zero value + false

🎯 COMMON PATTERNS:
• Worker pools: Multiple goroutines processing from same channel
• Fan-out: One sender, multiple receivers
• Fan-in: Multiple senders, one receiver
• Pipeline: Chain of processing stages

🚨 GOTCHAS:
❌ Sending to closed channel = panic
❌ Closing already closed channel = panic
❌ Receiving from nil channel = blocks forever
❌ Sending to nil channel = blocks forever

🔧 BEST PRACTICES:
• Close channels from sender side, not receiver
• Use buffered channels for async communication
• Use select for non-blocking operations
• Don't close channels unless necessary (GC will handle)

💡 CHANNEL AXIOMS:
• A send to a nil channel blocks forever
• A receive from a nil channel blocks forever
• A send to a closed channel panics
• A receive from a closed channel returns zero value immediately

=============================================================================
*/