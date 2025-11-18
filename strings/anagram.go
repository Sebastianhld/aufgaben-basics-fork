package strings

import (
	"sort"
	"strings"
)

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	// die Funktion strings.ToLower schreibt den String in Unicode um

	if len([]rune(s1)) != len([]rune(s2)) {
		//Wandelt den String s1 in ein Slice aus einzelnen Zeichen um, wenn die nicht gleich sind, dann falsch

		return false
	}

	r1 := []rune(s1)
	r2 := []rune(s2)
	//macht aus einem String eine Liste einzelner, echter Zeichen.

	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	//wir sortieren in alphabetischer Reihenfolge

	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	// TODO
	return false
}
