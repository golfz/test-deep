package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	fileName    = "testfile.dat"
	dataSize    = 100 * 1024 // 100KB
	measureTime = 1 * time.Minute
)

func generateRandomData(size int) ([]byte, error) {
	data := make([]byte, size)
	_, err := rand.Read(data)
	return data, err
}

func writeFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}

func readFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}

func main() {
	var writeDuration, readDuration time.Duration
	var loopCount int

	startTime := time.Now()

	fmt.Printf("Start testing, @time: %s\n", startTime)
	fmt.Printf("Result will be printed every %s, please wait...\n", measureTime)

	for {
		// Generate random data
		data, err := generateRandomData(dataSize)
		if err != nil {
			log.Fatalf("Failed to generate random data: %v", err)
		}

		// Write data to file
		startWrite := time.Now()
		err = writeFile(fileName, data)
		writeDuration += time.Since(startWrite)
		if err != nil {
			log.Fatalf("Failed to write data to file: %v", err)
		}

		// Read data from file
		startRead := time.Now()
		readData, err := readFile(fileName)
		readDuration += time.Since(startRead)
		if err != nil {
			log.Fatalf("Failed to read data from file: %v", err)
		}

		// Compare data
		if !bytes.Equal(data, readData) {
			log.Fatalf("Data mismatch!")
		}

		loopCount++

		// Calculate average count of loops in 1 minute
		elapsed := time.Since(startTime)
		if elapsed >= measureTime {
			averageLoops := float64(loopCount) / elapsed.Minutes()
			fmt.Println("--------------------------------------------------")
			fmt.Printf("Time: %s\n", time.Now())
			fmt.Printf("Average loops per minute: %.2f\n", averageLoops)
			fmt.Printf("Total loops: %d, Total time: %s\n", loopCount, elapsed)
			fmt.Printf("Average write duration: %s\n", writeDuration/time.Duration(loopCount))
			fmt.Printf("Average read duration: %s\n", readDuration/time.Duration(loopCount))
			// Reset counters
			startTime = time.Now()
			loopCount = 0
			writeDuration = 0
			readDuration = 0
		}
	}
}
