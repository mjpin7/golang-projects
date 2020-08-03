package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

/**
* Helper Function
*
* Helps streamline checking for errors
**/
func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

/* MAIN FUNCTION */
func main() {
	// Get the filename from the flag '-f' or else just default to "problems.csv"
	filename := flag.String("f", "problems.csv", "CSV file to read from")
	timer := flag.Int("t", 30, "Amount of time given for quiz (in seconds)")
	flag.Parse()

	// Open file for reading
	f, err := os.Open(*filename)
	check(err)

	// Create the new csv reader
	r := csv.NewReader(f)

	// Var to keep track of the number of questions answered correctly
	corr := 0

	fmt.Printf("Press enter when you would like to start...")
	fmt.Scanf("\n")

	// Create timer with the inputed amount of seconds
	t1 := time.NewTimer(time.Duration(*timer) * time.Second)

	q, err := r.ReadAll()
	check(err)

	// Read the problems
problem:
	for i, p := range q {
		// Create a new channel for the input
		ansCh := make(chan string)

		// Ask question and get answer from input
		fmt.Printf("\nQuestion %d: %s\nEnter your answer: ", i+1, p[0])

		// This will create a nw goroutine, asynchronously call
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)

			// Put the input into the ansCh input channel
			ansCh <- answer
		}()

		select {
		// If the timer has run out, then break out and end
		case <-t1.C:
			fmt.Printf("\n\nTimes up!!")
			t1.Stop()
			break problem
		// If we get a input answer, check if correct
		case ans := <-ansCh:
			// Check if correct
			if ans == p[1] {
				corr++
			}
		}
	}

	// Calculate score
	score := (float64(corr) / float64(len(q))) * 100
	fmt.Printf("\n\nNumber of correct answers: %d\nNumber of total questions: %d\nScore: %.2f\n", corr, len(q), score)

	f.Close()
}
