// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package convExt

import "fmt"

type Meter float64
type CentiMeter float64

func (c CentiMeter) String() string    { return fmt.Sprintf("%g cm", c) }
func (m Meter) String() string { return fmt.Sprintf("%g meter", m) }

//!-
