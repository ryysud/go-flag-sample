package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	b := flag.Bool("bool", false, "specify bool type")
	d := flag.Duration("duration", time.Duration(1), "specify time.Duration type")
	f64 := flag.Float64("float64", 0.00, "specify float64 type")
	i := flag.Int("int", 0, "specify int type")
	s := flag.String("string", "default", "specify string type")

	flag.Parse()

	fmt.Printf("%d flags are specified!\n\n", flag.NFlag())

	fmt.Printf("bool param => type: %T, value: %v\n", *b, *b)
	fmt.Printf("duration param => type: %T, value: %v\n", *d, *d)
	fmt.Printf("float64 param => type: %T, value: %v\n", *f64, *f64)
	fmt.Printf("int param => type: %T, value: %v\n", *i, *i)
	fmt.Printf("string param => type: %T, value: %v\n", *s, *s)

}
