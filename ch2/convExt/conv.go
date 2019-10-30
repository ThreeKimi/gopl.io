// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package convExt

func CToM(c CentiMeter) Meter { return Meter(c/100.0) }

func MToC(m Meter) CentiMeter { return CentiMeter(m * 100.0) }

//!-
