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

	// 指定されたフラグのみ参照
	// flag.Visit(func(f *flag.Flag) {
	// 	fmt.Printf("%v param => type: %T, value: %v\n", f.Name, f.Value, f.Value)
	// })

	// 全てのフラグを参照
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%v param => type: %T, value: %v\n", f.Name, f.Value, f.Value)
	})

}
