// Package lengthconv performs centimeter and inch conversions.
package lengthconv

import "fmt"

type Centimeter float64
type Inch float64

const (
	OneInch Centimeter = 2.54
)

func (cm Centimeter) String() string { return fmt.Sprintf("%gcm", cm) }
func (inch Inch) String() string     { return fmt.Sprintf("%g″", inch) }

// !-
