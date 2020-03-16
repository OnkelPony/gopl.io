// Cf converts its numeric argument to inch and centimeter.
package main

import (
	"fmt"
	lengthconv "github.com/gopl.io/ch2/lengthconv2.2"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		inch := lengthconv.Inch(l)
		cm := lengthconv.Centimeter(l)
		fmt.Printf("%s = %s, %s = %s\n", inch, lengthconv.InchToCm(inch), cm, lengthconv.CmToInch(cm))

	}
}
