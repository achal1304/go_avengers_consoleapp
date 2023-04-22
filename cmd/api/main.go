package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AvengersMissions map[string][]string

type Avenger struct {
	Name              string
	Abilities         []string
	MissionsAssigned  int
	MissionsCompleted int
}

var Avengers []Avenger

type Mission struct {
	Name    string
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

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	listmenu()

	for {
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil || choice > 6 || choice < 1 {
			fmt.Println("Please choose a valid option")
			continue
		}
		switch choice {
		case 1:
			getAvengers()
		case 2:
			var mission Mission
			input := takeInput("Enter a comma-separated list of strings: ")
			avengers := strings.Split(input, ",")
			isValidAvenger := isValidAvenger(avengers)
			if !isValidAvenger {
				fmt.Print("Not a valid avengers list")
			}

			mission.Name = takeInputText("Enter Mission: ")
			mission.Details = takeInputText("Enter Mission Details: ")
			mission.Status = "Assigned"

		default:
			fmt.Println("This is default message")
		}
	}
}

func takeInput(inputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputText)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	text = strings.TrimSuffix(text, "\n")
	return text
}

func takeInputText(inputText string) string {
	fmt.Print(inputText)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	return choice
}

func isValidAvenger(avengers []string) bool {
	fmt.Printf("Avenger: %v\n", avengers)
	var countAvenger = 0
	for _, avenger := range avengers {
		for _, a := range Avengers {
			if strings.TrimSpace(avenger) == strings.TrimSpace(a.Name) {
				fmt.Printf("Match found: %s\n", avenger)
				countAvenger += 1
			}
		}
	}
	fmt.Printf("Count: %d, CountAv: %d\n", countAvenger, len(avengers))
	return countAvenger == len(avengers)
}

func assignMissions(avenger Avenger, mission Mission) {

}

func insertAvengers(avenger Avenger) {
	Avengers = append(Avengers, avenger)
}

func getAvengers() {
	fmt.Println("Get avengers called")
	for _, avenger := range Avengers {
		fmt.Println(avenger.Name)
	}
}

func checkMissionDetails() {

}
