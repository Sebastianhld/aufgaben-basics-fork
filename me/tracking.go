package strava

import (
	"fmt"
)

func PickKategorie() {
	//Wähle eine Kategorie aus
	var gym = []string{"push", "pull", "legs"}
	fmt.Println(gym)

	var bike = []string{"enduro", "dirtjump", "bikepark", "zone 2"}
	fmt.Println(bike)

	var dance = []string{"alone", "together"}
	fmt.Println(dance)

}
func Main() {
	fmt.Println("Hello, what's your name?")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello ", name, "welche Sportart willst du machen")
	//Hallo _ , welche Sportart willst du machen?
	//Wähle eine Sportart aus
	sport := []string{gym, bike, dance}
	fmt.Println(sport)
}
