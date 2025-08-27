/*
=============================================================================
                        👷 GO WORKER POOLS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Worker pools are a concurrency pattern where a fixed number of goroutines
(workers) process jobs from a shared queue. This pattern provides controlled
concurrency and efficient resource utilization.

🔑 KEY FEATURES:
• Controlled concurrency with fixed worker count
• Job queue for distributing work
• Result collection and error handling
• Graceful shutdown and cleanup

💡 REAL-WORLD ANALOGY:
Worker Pool = Restaurant Kitchen
- Workers = Chefs working in parallel
- Job queue = Order tickets waiting to be prepared
- Results = Completed dishes ready to serve
- Pool size = Number of chefs on duty

🎯 WHY USE WORKER POOLS?
• Limit resource usage (memory, connections)
• Process large amounts of work efficiently
• Control system load and prevent overload
• Implement rate limiting and backpressure

=============================================================================
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 📋 JOB DEFINITIONS
type Job struct {
	ID   int
	Data string
}

type Result struct {
	Job    Job
	Output string
	Error  error
}

// 🏭 BASIC WORKER POOL
func basicWorkerPool() {
	fmt.Println("🏭 Basic Worker Pool")
	fmt.Println("===================")

	const numWorkers = 3
	const numJobs = 10

	// Create channels
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job.ID)
				
				// Simulate work
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				
				// Send result
				results <- Result{
					Job:    job,
					Output: fmt.Sprintf("Processed by worker %d", id),
					Error:  nil,
				}
			}
			fmt.Printf("Worker %d finished\n", id)
		}(w)
	}

	// Send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{ID: j, Data: fmt.Sprintf("job-%d", j)}
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	for result := range results {
		fmt.Printf("Result: Job %d -> %s\n", result.Job.ID, result.Output)
	}
}

// 🎯 ADVANCED WORKER POOL WITH STRUCT
type WorkerPool struct {
	workers    int
	jobQueue   chan Job
	resultChan chan Result
	quit       chan bool
	wg         sync.WaitGroup
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
	return &WorkerPool{
		workers:    workers,
		jobQueue:   make(chan Job, queueSize),
		resultChan: make(chan Result, queueSize),
		quit:       make(chan bool),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i + 1)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	
	for {
		select {
		case job := <-wp.jobQueue:
			fmt.Printf("🔧 Worker %d processing job %d\n", id, job.ID)
			
			// Simulate processing time
			processingTime := time.Duration(rand.Intn(500)+100) * time.Millisecond
			time.Sleep(processingTime)
			
			// Simulate occasional errors
			var err error
			if rand.Float32() < 0.1 { // 10% error rate
				err = fmt.Errorf("processing failed for job %d", job.ID)
			}
			
			result := Result{
				Job:    job,
				Output: fmt.Sprintf("Completed by worker %d in %v", id, processingTime),
				Error:  err,
			}
			
			wp.resultChan <- result
			
		case <-wp.quit:
			fmt.Printf("🛑 Worker %d stopping\n", id)
			return
		}
	}
}

func (wp *WorkerPool) Submit(job Job) {
	wp.jobQueue <- job
}

func (wp *WorkerPool) Results() <-chan Result {
	return wp.resultChan
}

func (wp *WorkerPool) Stop() {
	close(wp.quit)
	wp.wg.Wait()
	close(wp.jobQueue)
	close(wp.resultChan)
}

// 📊 BATCH PROCESSOR
type BatchProcessor struct {
	workerPool *WorkerPool
	batchSize  int
	timeout    time.Duration
}

func NewBatchProcessor(workers, queueSize, batchSize int, timeout time.Duration) *BatchProcessor {
	return &BatchProcessor{
		workerPool: NewWorkerPool(workers, queueSize),
		batchSize:  batchSize,
		timeout:    timeout,
	}
}

func (bp *BatchProcessor) ProcessBatch(jobs []Job) []Result {
	bp.workerPool.Start()
	defer bp.workerPool.Stop()

	// Submit all jobs
	for _, job := range jobs {
		bp.workerPool.Submit(job)
	}

	// Collect results with timeout
	var results []Result
	timeout := time.After(bp.timeout)
	
	for i := 0; i < len(jobs); i++ {
		select {
		case result := <-bp.workerPool.Results():
			results = append(results, result)
		case <-timeout:
			fmt.Printf("⏰ Timeout reached, collected %d/%d results\n", len(results), len(jobs))
			return results
		}
	}
	
	return results
}

// 🌐 HTTP REQUEST WORKER POOL EXAMPLE
type URLJob struct {
	ID  int
	URL string
}

type URLResult struct {
	Job        URLJob
	StatusCode int
	Error      error
	Duration   time.Duration
}

func httpWorkerPool() {
	fmt.Println("\n🌐 HTTP Request Worker Pool")
	fmt.Println("===========================")

	const numWorkers = 5
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
	}

	jobs := make(chan URLJob, len(urls))
	results := make(chan URLResult, len(urls))

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				
				// Simulate HTTP request
				fmt.Printf("🌐 Worker %d requesting %s\n", id, job.URL)
				time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
				
				// Simulate response
				statusCode := 200
				if rand.Float32() < 0.2 {
					statusCode = 500
				}
				
				results <- URLResult{
					Job:        job,
					StatusCode: statusCode,
					Duration:   time.Since(start),
					Error:      nil,
				}
			}
		}(w)
	}

	// Submit jobs
	for i, url := range urls {
		jobs <- URLJob{ID: i + 1, URL: url}
	}
	close(jobs)

	// Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		status := "✅"
		if result.StatusCode >= 400 {
			status = "❌"
		}
		fmt.Printf("%s Job %d: %s -> %d (%v)\n", 
			status, result.Job.ID, result.Job.URL, result.StatusCode, result.Duration)
	}
}

func main() {
	fmt.Println("👷 WORKER POOLS TUTORIAL")
	fmt.Println("========================")

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// 🎯 DEMO 1: Basic Worker Pool
	basicWorkerPool()

	// 🎯 DEMO 2: Advanced Worker Pool
	fmt.Println("\n🎯 Advanced Worker Pool")
	fmt.Println("=======================")

	pool := NewWorkerPool(4, 20)
	pool.Start()

	// Submit jobs
	numJobs := 15
	go func() {
		for i := 1; i <= numJobs; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("task-%d", i),
			}
			pool.Submit(job)
		}
	}()

	// Collect results
	var successCount, errorCount int
	for i := 0; i < numJobs; i++ {
		result := <-pool.Results()
		if result.Error != nil {
			fmt.Printf("❌ Job %d failed: %v\n", result.Job.ID, result.Error)
			errorCount++
		} else {
			fmt.Printf("✅ Job %d: %s\n", result.Job.ID, result.Output)
			successCount++
		}
	}

	pool.Stop()
	fmt.Printf("📊 Summary: %d successful, %d failed\n", successCount, errorCount)

	// 🎯 DEMO 3: Batch Processor
	fmt.Println("\n🎯 Batch Processor")
	fmt.Println("==================")

	batchProcessor := NewBatchProcessor(3, 10, 5, 5*time.Second)
	
	// Create batch of jobs
	var batchJobs []Job
	for i := 1; i <= 8; i++ {
		batchJobs = append(batchJobs, Job{
			ID:   i,
			Data: fmt.Sprintf("batch-job-%d", i),
		})
	}

	fmt.Printf("Processing batch of %d jobs...\n", len(batchJobs))
	results := batchProcessor.ProcessBatch(batchJobs)
	
	fmt.Printf("Batch completed: %d results\n", len(results))
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("❌ Batch job %d: %v\n", result.Job.ID, result.Error)
		} else {
			fmt.Printf("✅ Batch job %d: %s\n", result.Job.ID, result.Output)
		}
	}

	// 🎯 DEMO 4: HTTP Worker Pool (simulated)
	httpWorkerPool()

	// 🎯 DEMO 5: Performance Comparison
	fmt.Println("\n🎯 Performance Comparison")
	fmt.Println("=========================")

	// Sequential processing
	start := time.Now()
	fmt.Println("Sequential processing...")
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
	sequentialTime := time.Since(start)
	fmt.Printf("Sequential time: %v\n", sequentialTime)

	// Parallel processing with worker pool
	start = time.Now()
	fmt.Println("Parallel processing with worker pool...")
	
	parallelPool := NewWorkerPool(5, 10)
	parallelPool.Start()

	// Submit jobs
	for i := 1; i <= 10; i++ {
		parallelPool.Submit(Job{ID: i, Data: fmt.Sprintf("perf-test-%d", i)})
	}

	// Collect results
	for i := 0; i < 10; i++ {
		<-parallelPool.Results()
	}

	parallelPool.Stop()
	parallelTime := time.Since(start)
	fmt.Printf("Parallel time: %v\n", parallelTime)
	fmt.Printf("Speedup: %.2fx\n", float64(sequentialTime)/float64(parallelTime))

	fmt.Println("\n✨ All worker pool demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

👷 WORKER POOL PATTERN:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic structure                                                      │
│ jobs := make(chan Job, queueSize)                                       │
│ results := make(chan Result, queueSize)                                 │
│                                                                         │
│ // Start workers                                                        │
│ for w := 1; w <= numWorkers; w++ {                                      │
│     go func(id int) {                                                   │
│         for job := range jobs {                                         │
│             // Process job                                              │
│             result := processJob(job)                                   │
│             results <- result                                           │
│         }                                                               │
│     }(w)                                                                │
│ }                                                                       │
│                                                                         │
│ // Submit jobs                                                          │
│ for _, job := range jobList {                                           │
│     jobs <- job                                                         │
│ }                                                                       │
│ close(jobs)                                                             │
│                                                                         │
│ // Collect results                                                      │
│ for i := 0; i < len(jobList); i++ {                                     │
│     result := <-results                                                 │
│     // Handle result                                                    │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

🎯 WORKER POOL COMPONENTS:
┌─────────────────────────────────────────────────────────────────────────┐
│ type WorkerPool struct {                                                │
│     workers    int              // Number of worker goroutines          │
│     jobQueue   chan Job         // Channel for incoming jobs            │
│     resultChan chan Result      // Channel for results                  │
│     quit       chan bool        // Channel for shutdown signal          │
│     wg         sync.WaitGroup   // Wait for workers to finish           │
│ }                                                                       │
│                                                                         │
│ // Key methods                                                          │
│ func (wp *WorkerPool) Start()                // Start workers           │
│ func (wp *WorkerPool) Submit(job Job)        // Submit job              │
│ func (wp *WorkerPool) Results() <-chan Result // Get results channel    │
│ func (wp *WorkerPool) Stop()                 // Graceful shutdown       │
└─────────────────────────────────────────────────────────────────────────┘

🔧 IMPLEMENTATION PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Worker function with select                                          │
│ func (wp *WorkerPool) worker(id int) {                                  │
│     for {                                                               │
│         select {                                                        │
│         case job := <-wp.jobQueue:                                      │
│             // Process job                                              │
│             result := processJob(job)                                   │
│             wp.resultChan <- result                                     │
│         case <-wp.quit:                                                 │
│             return                                                      │
│         }                                                               │
│     }                                                                   │
│ }                                                                       │
│                                                                         │
│ // Graceful shutdown                                                    │
│ func (wp *WorkerPool) Stop() {                                          │
│     close(wp.quit)        // Signal workers to stop                     │
│     wp.wg.Wait()          // Wait for workers to finish                 │
│     close(wp.jobQueue)    // Close job queue                            │
│     close(wp.resultChan)  // Close result channel                       │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

📊 SIZING CONSIDERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Worker count guidelines                                              │
│ • CPU-bound tasks: runtime.NumCPU()                                     │
│ • I/O-bound tasks: Higher count (10-100+ workers)                       │
│ • Memory-limited: Consider memory per worker                            │
│ • External API: Respect rate limits                                     │
│                                                                         │
│ // Queue size guidelines                                                │
│ • Small queue: Better backpressure, less memory                         │
│ • Large queue: Better throughput, more memory                           │
│ • Buffered channels: len(jobs) = 2-10x number of workers                │
│                                                                         │
│ // Monitoring metrics                                                   │
│ • Queue length: len(jobQueue)                                           │
│ • Active workers: Track in worker function                              │
│ • Processing time: Measure per job                                      │
│ • Error rate: Track failed jobs                                         │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE OPTIMIZATION:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Batch processing                                                     │
│ type Batch []Job                                                        │
│ func processBatch(batch Batch) []Result {                               │
│     // Process multiple jobs together                                   │
│ }                                                                       │
│                                                                         │
│ // Worker pools with different priorities                               │
│ highPriorityJobs := make(chan Job, 100)                                 │
│ lowPriorityJobs := make(chan Job, 1000)                                 │
│                                                                         │
│ // Worker with priority handling                                        │
│ select {                                                                │
│ case job := <-highPriorityJobs:                                         │
│     // Process high priority job                                        │
│ case job := <-lowPriorityJobs:                                          │
│     // Process low priority job                                         │
│ case <-quit:                                                            │
│     return                                                              │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Size worker pool based on workload characteristics
• Use buffered channels for better performance
• Implement graceful shutdown with context or quit channel
• Monitor queue length and processing times
• Handle errors gracefully and provide retry mechanisms
• Use timeouts for long-running jobs
• Consider using sync.Pool for object reuse

🚨 COMMON MISTAKES:
❌ Too many workers for CPU-bound tasks
❌ Not implementing graceful shutdown
❌ Unbounded job queues causing memory issues
❌ Not handling worker panics
❌ Blocking operations in worker functions
❌ Not monitoring pool performance

🎯 REAL-WORLD USE CASES:
• Image/video processing pipelines
• HTTP request processing
• Database batch operations
• File processing and ETL
• Email sending systems
• Log processing and analysis
• API rate-limited operations

🔍 MONITORING AND DEBUGGING:
• Track active workers and queue sizes
• Measure job processing times
• Monitor error rates and types
• Use pprof for goroutine analysis
• Implement health checks
• Log worker lifecycle events

=============================================================================
*/