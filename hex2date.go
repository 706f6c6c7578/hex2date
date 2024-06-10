package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Flag for decimal input
	decimal := flag.Bool("d", false, "Set this flag to process decimal values")
	flag.Parse()

	// Usage instruction
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-d] [hex|dec value]\n", os.Args[0])
		flag.PrintDefaults()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		var seconds int64
		var err error

		// Process the input value based on the set flag
		if *decimal {
			seconds, err = strconv.ParseInt(input, 10, 64)
		} else {
			seconds, err = strconv.ParseInt(input, 16, 64)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing the value: %v\n", err)
			continue
		}

		// Convert Unix timestamp to a date in UTC
		date := time.Unix(seconds, 0).UTC()
		// Output the date in UTC
		fmt.Println(date.Format("Monday, 02. January 2006 15:04:05 +0000 (UTC)"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
	}
}
