package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const (
	maxInt32     = 2147483647
	ulimitedTime = maxInt32
)

func exitErr(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func catchOSSignals() {
	// Catch OS signals
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		fmt.Printf("Signal from OS: %s", s)
	}()
}

func showMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func genMem(total, block int) {
	// allocate memory 1MB at a time
	rand.Seed(time.Now().UTC().UnixNano())
	loop := int(math.Ceil(float64(total) / float64(block)))
	var res = make([][]byte, loop)
	for i := 0; i < loop; i++ {
		blockMB := block * 1024 * 1024
		if i == (loop - 1) {
			blockMB = (total - block*i) * 1024 * 1024
		}
		res[i] = make([]byte, blockMB)
		// populate array so it takes up memory
		rand.Read(res[i])
		showMem()
		// Wait for 100ms
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("                       Final result                              ")
	fmt.Println("-----------------------------------------------------------------")
	// Final
	showMem()
}

func genCPU(cores, timeSec int) {
	runtime.GOMAXPROCS(cores)

	done := make(chan struct{})
	for i := 0; i <= cores; i++ {
		runtime.LockOSThread()
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	// how long
	time.Sleep(time.Duration(timeSec) * time.Second)
	close(done)
}

func main() {
	catchOSSignals()

	memCmd := flag.NewFlagSet("mem", flag.ExitOnError)
	memTotal := memCmd.Int("total", 1024, "total memory in MB be generated")
	memBlock := memCmd.Int("block", 1, "size of a single block in MB will be allocated each time")

	cpuCmd := flag.NewFlagSet("cpu", flag.ExitOnError)
	cpuCores := cpuCmd.Int("cores", runtime.NumCPU(), "number of CPU cores be used")
	cpuTime := cpuCmd.Int("time", ulimitedTime, "time in seconds to run the load generator")

	if len(os.Args) < 2 {
		exitErr("expected 'mem' or 'cpu' subcommands")
	}

	switch os.Args[1] {
	case "mem":
		memCmd.Parse(os.Args[2:])
		// Validate flags
		if *memTotal < 1 {
			exitErr("total is invalid, must be positive int")
		}
		if *memBlock < 1 || *memBlock > *memTotal {
			exitErr("block is invalid, must be between 1 - `total*")
		}
		genMem(*memTotal, *memBlock)
	case "cpu":
		cpuCmd.Parse(os.Args[2:])
		// Validate flags
		if *cpuCores < 1 || *cpuCores > runtime.NumCPU() {
			exitErr("cores is invalid, must be between 1 - `max CPU cores`")
		}
		if *cpuTime < 1 {
			exitErr("time is invalid, must be positive int")
		}

		genCPU(*cpuCores, *cpuTime)
	default:
		exitErr("expected 'mem' or 'cpu' subcommands")
	}
}
