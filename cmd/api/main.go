package main

import "fmt"

type AvengersMissions map[string][]string

type Avenger struct {
	Name              string
	Abilities         []string
	MissionsAssigned  int
	MissionsCompleted int
}

var Avengers []Avenger

type Mission struct {
	Details string
	Status  string
}

var Missions []Mission

func main() {
	var avenger = Avenger{
		Name:              "Tony Stark",
		Abilities:         []string{"Abilities: Highly Intelligent Suit of Armor"},
		MissionsAssigned:  0,
		MissionsCompleted: 0,
	}
	var avenger1 = Avenger{
		Name:              "Thor",
		Abilities:         []string{"Abilities: Highly Intelligent Suit of Armor"},
		MissionsAssigned:  0,
		MissionsCompleted: 0,
	}
	insertAvengers(avenger)
	insertAvengers(avenger1)

	fmt.Println("Avengers Assemble")
	menu()
}

func listmenu() {
	fmt.Println("1. Check the missions")
	fmt.Println("2. Assign mission to Avengers")
	fmt.Println("3. Check mission’s details")
	fmt.Println("4. Check Avenger’s details")
	fmt.Println("5. Update Mission’s status")
	fmt.Println("6. List Avengers")
}

func insertAvengers(avenger Avenger) {
	Avengers = append(Avengers, avenger)
}

func menu() {
	var choice int

	listmenu()

	for {
		fmt.Scan(&choice)

		if choice > 6 || choice < 1 {
			fmt.Println("Plese choose correct option")
			break
		}
		switch choice {
		case 1:
			getAvengers()
		default:
			fmt.Println("THis is default message")
		}
	}
}

func getAvengers() {
	fmt.Println("Get avengers called")
	for _, avenger := range Avengers {
		fmt.Println(avenger.Name)
	}
}
