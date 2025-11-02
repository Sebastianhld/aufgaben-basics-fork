package forms

import "math"

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	// TODO
	return math.Sqrt(a*a+b*b) + a + b
	// c über Pythagroas berechnet, dann Umfang = c + a + b
}
