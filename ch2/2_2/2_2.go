package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"go/ch2/tempconv"
)

func convert(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

// It's nonsense to implement a convert function between meter and inch, which
// just looks like tempcomv package. Just implementation of read from stdin there.
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		convert(t)
	}
	// If no args given, read from stdin
	if len(os.Args) <= 1 {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			t, err := strconv.ParseFloat(in.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			convert(t)
		}
		if err := in.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
