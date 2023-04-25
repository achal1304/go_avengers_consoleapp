package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	listmenu()

	for {
		choice, err := strconv.Atoi(takeInputText("Enter your choice : "))
		if err != nil || choice > 6 || choice < 1 {
			fmt.Println("Please choose a valid option")
			continue
		}
		switch choice {
		case 1:
			getAllMissionDetails()
		case 2:
			var mission Mission
			input := takeInput("Enter a comma-separated list of strings: ")
			avengers := strings.Split(input, ",")
			isValidAvenger, avengersAssigned := isValidAvenger(avengers)
			hasBeenAssigned, avengerAssigned := isAvengerAssigned(avengers)
			if !hasBeenAssigned {
				fmt.Printf("Sorry, %s has already been working on two missions.", avengerAssigned)
				continue
			}
			if !isValidAvenger {
				fmt.Print("Not a valid avengers list")
				continue
			}

			mission.Name = takeInputText("Enter Mission: ")
			mission.Details = takeInputText("Enter Mission Details: ")
			mission.Status = "Assigned"

			assignMissions(avengersAssigned, mission)
		case 3:
			checkMissionDetails(takeInputText("Enter Mission Name : "))
		case 4:
			getAvengerByName(takeInputText("Enter Avenger Name : "))
		case 5:
			updateMissionStatus(takeInputText("Enter Mission Name : "))
		case 6:
			getAvengerStatusAndMissions()
		default:
			fmt.Println("This is default message")
		}
	}
}
