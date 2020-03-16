// Package lengthconv converts cm to inches and vice versa
package lengthconv

// CmToInch converts a centimeter length to inch.
func CmToInch(cm Centimeter) Inch { return Inch(cm / OneInch) }

// InchToCm converts an inch length to centimeter.
func InchToCm(inch Inch) Centimeter { return Centimeter(inch) * OneInch }
