package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Dreiecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	// Nichts zeichnen bei unsinnigen Werten
	if length <= 0 {
		return
	}

	// b = "Höhe" von unten (0 = unterste Zeile)
	// Wir laufen von oben nach unten, deshalb: length-1 runter bis 0
	for b := length - 1; b >= 0; b-- {
		// Anzahl der "Spalten" in dieser Zeile
		// Unten (b=0) sollen es `length` Zeichen sein, oben nur 1
		nCells := length - b

		for c := 0; c < nCells; c++ {
			// Randbedingungen:
			// 1. Untere Seite: b == 0
			// 2. Linke Seite: c == 0
			// 3. Hypotenuse (schräge Seite): c == nCells-1
			if b == 0 || c == 0 || c == nCells-1 {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}
		}
		fmt.Println()
	}
}
