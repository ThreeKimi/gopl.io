// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

package main

import (
  "bufio"
	"fmt"
	"os"
	"strconv"
	"gopl.io/ch2/convExt"
)

func main() {
  params := os.Args[1:]
  if len(params) == 0 {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
      t, err := strconv.ParseFloat(input.Text(),64)
      if err != nil {
        fmt.Fprintf(os.Stderr, "cf: %v\n", err)
        os.Exit(1)
      }
      PrintInfo(t)
    }
  } else {
    for _, arg := range params {
      t, err := strconv.ParseFloat(arg, 64)
      if err != nil {
        fmt.Fprintf(os.Stderr, "cf: %v\n", err)
        os.Exit(1)
      }
      PrintInfo(t)
    }
  }
}

func PrintInfo(t float64){
  m := convExt.Meter(t)
  c := convExt.CentiMeter(t)
  fmt.Printf("%s = %s, %s = %s\n",
    m, convExt.MToC(m), c, convExt.CToM(c))
}

//!-
