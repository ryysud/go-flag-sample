package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	var (
		b   bool
		d   time.Duration
		f64 float64
		i   int
		s   string
	)

	flag.BoolVar(&b, "bool", false, "specify bool type")
	flag.DurationVar(&d, "duration", time.Duration(1), "specify time.Duration type")
	flag.Float64Var(&f64, "float64", 0.00, "specify float64 type")
	flag.IntVar(&i, "int", 0, "specify int type")
	flag.StringVar(&s, "string", "default", "specify string type")

	flag.Parse()

	fmt.Printf("%d flags are specified!\n\n", flag.NFlag())

	fmt.Printf("bool param => type: %T, value: %v\n", b, b)
	fmt.Printf("duration param => type: %T, value: %v\n", d, d)
	fmt.Printf("float64 param => type: %T, value: %v\n", f64, f64)
	fmt.Printf("int param => type: %T, value: %v\n", i, i)
	fmt.Printf("string param => type: %T, value: %v\n", s, s)

}
