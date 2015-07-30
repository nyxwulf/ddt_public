package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"unicode/utf8"
)

var delimiter = flag.String("delimiter", "\t", "String delimiter")
var numeric = flag.Bool("n", false, "Sort keys as ascending numeric sequence")

func main() {
	flag.Parse()
	stdin := bufio.NewReader(os.Stdin)
	reader := csv.NewReader(stdin)
	sep, _ := utf8.DecodeRuneInString(*delimiter)
	reader.Comma = sep

	counter := make(map[string]int)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			genericError(err)
		}

		countRecord(counter, record[:])
	}

	if *numeric {
		displayNumericCount(counter)
	} else {
		for key, value := range counter {
			fmt.Printf("%s%s%v\n", key, *delimiter, value)
		}

	}

}

func countRecord(counter map[string]int, record []string) {
	if len(record) == 0 {
		return
	}

	counter[record[0]] += 1

	return
}

func displayNumericCount(counter map[string]int) error {
	keys := []int{}

	for k, _ := range counter {
		key, err := strconv.Atoi(k)

		if err != nil {
			err := fmt.Errorf("Cannot convert %v to integer", k)
			return err
		}

		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, k := range keys {
		key := strconv.Itoa(k)
		fmt.Printf("%s%s%v\n", key, *delimiter, counter[key])
	}

	return nil

}

func genericError(err error) {
	os.Stdout.Sync()
	fmt.Fprintf(os.Stderr, "ERROR: %s", err)
	os.Exit(2)
}
