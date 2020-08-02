package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
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
		panic(e)
	}
}

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

	// Vars to keep track of the number of questions asked and number answered correctly
	count, corr := 0, 0

	fmt.Printf("Press enter when you would like to start...")
	fmt.Scanf("\n")

	t1 := time.NewTimer(time.Duration(*timer) * time.Second)

	// Read the csv file in a loop
problem:
	for {
		ansCh := make(chan string)

		q, err := r.Read()

		if err == io.EOF {
			break problem
		}

		check(err)

		count++

		// Ask question and get answer from input
		fmt.Printf("\nQuestion %d: %s\nEnter your answer: ", count, q[0])

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansCh <- answer
		}()

		select {
		case <-t1.C:
			break problem
		case ans := <-ansCh:
			// Check if correct
			if ans == q[1] {
				corr++
			}
		}
	}

	score := (float64(corr) / float64(count)) * 100
	fmt.Printf("\n\nNumber of correct answers: %d\nNumber of total questions: %d\nScore: %.2f\n", corr, count, score)

	f.Close()
}
