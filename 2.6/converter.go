// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"2.6/lengthconv"
	"2.6/tempconv"
	"2.6/weightconv"
)

var t = flag.Bool("t", false, "prepares temperature conversions")
var l = flag.Bool("l", false, "prepares length conversions")
var w = flag.Bool("w", false, "prepares weight conversions")

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguements.")
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			processVal(input.Text())
		}
	} else {
		for _, arg := range os.Args[2:] {
			processVal(arg)
		}
	}
}

func getVal(val string) float64 {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return t
}

func processVal(val string) {
	v := getVal(val)

	if *t {
		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	} else if *l {
		f := lengthconv.Foot(v)
		m := lengthconv.Meter(v)
		fmt.Printf("%s = %s, %s = %s\n",
			f, lengthconv.FToM(f), m, lengthconv.MToF(m))
	} else if *w {
		p := weightconv.Pound(v)
		k := weightconv.Kilogram(v)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.PToK(p), k, weightconv.KToP(k))
	} else {
		fmt.Println("No conversion specified.")
	}

}
