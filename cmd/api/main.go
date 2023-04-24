package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var AvengersMissions = make(map[Mission]*[]Avenger)

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
			fmt.Print("Enter Mission Name: ")
			scanner.Scan()
			checkMissionDetails(scanner.Text())
		case 4:
			fmt.Print("Enter Avenger Name: ")
			scanner.Scan()
			getAvengerByName(scanner.Text())
		case 5:
			fmt.Print("Enter Mission Name: ")
			scanner.Scan()
			updateMissionStatus(scanner.Text())
		case 6:
			getAvengerStatusAndMissions()
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

func isValidAvenger(avengers []string) (bool, []Avenger) {
	fmt.Printf("Avenger: %v\n", avengers)
	var countAvenger = 0
	var avengerDetails []Avenger
	for _, avenger := range avengers {
		for _, a := range Avengers {
			if strings.TrimSpace(avenger) == strings.TrimSpace(a.Name) {
				fmt.Printf("Match found: %s\n", avenger)
				avengerDetails = append(avengerDetails, a)
				countAvenger += 1
			}
		}
	}
	fmt.Printf("Count: %d, CountAv: %d\n", countAvenger, len(avengers))
	return countAvenger == len(avengers), avengerDetails
}

func isAvengerAssigned(avengers []string) (bool, string) {
	for _, avengerMission := range Avengers {
		for _, name := range avengers {
			if strings.TrimSpace(name) == strings.TrimSpace(avengerMission.Name) {
				if avengerMission.MissionsAssigned >= 2 {
					return false, name
				}
			}
		}
	}
	return true, ""
}

func assignMissions(avengers []Avenger, mission Mission) {
	Missions = append(Missions, mission)
	AvengersMissions[mission] = &avengers
	updateMissionStatusAvenger(avengers, false)
}

func updateMissionStatusAvenger(avengers []Avenger, isCompleted bool) {
	for i := 0; i < len(Avengers); i++ {
		for _, avenger := range avengers {
			if strings.TrimSpace(avenger.Name) == strings.TrimSpace(Avengers[i].Name) {
				if !isCompleted {
					Avengers[i].MissionsAssigned += 1
				} else {
					Avengers[i].MissionsAssigned -= 1
					Avengers[i].MissionsCompleted += 1
				}
			}
		}
	}
}

func insertAvengers(avenger Avenger) {
	Avengers = append(Avengers, avenger)
}

func getAvengerByName(avengerName string) {
	for _, avenger := range Avengers {
		if avenger.Name == avengerName {
			fmt.Println("Name : ", avenger.Name)
			fmt.Println("Abilities : ", strings.Join(avenger.Abilities, ", "))
			fmt.Println("Missions Assigned : ", avenger.MissionsAssigned)
			fmt.Println("Missions Completed : ", avenger.MissionsCompleted)
			return
		}
	}
	fmt.Println("Please Enter Avenger that exists")
}

func getAllMissionDetails() {
	if len(AvengersMissions) == 0 {
		fmt.Println("No Mission has been assigned to an Avenger.")
		return
	}

	fmt.Println("Mission Name               Status              Avenger")
	for key, value := range AvengersMissions {
		fmt.Printf("%s               ", key.Name)
		fmt.Printf("%s               ", key.Status)
		fmt.Println(strings.Join(getNamesFromAvengers(value), ", "))
	}

}

func checkMissionDetails(missionName string) {
	for key, value := range AvengersMissions {
		if key.Name == missionName {
			fmt.Println("Mission Details: ", key.Details)
			fmt.Println("Mission Status : ", key.Status)
			fmt.Println("Avengers assigned", strings.Join(getNamesFromAvengers(value), ", "))
			return
		}
	}
	fmt.Println("No Mission with such name")
}

func getNamesFromAvengers(avengers *[]Avenger) []string {
	var avengersAssigned []string
	for _, avenger := range *avengers {
		avengersAssigned = append(avengersAssigned, avenger.Name)
	}
	return avengersAssigned
}

func updateMissionStatus(missionName string) {
	for i := 0; i < len(Missions); i++ {
		if Missions[i].Name == missionName {
			fmt.Print("Enter new status: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newStatus := scanner.Text()
			updateMissionStatusAvenger(*AvengersMissions[Missions[i]], true)
			Missions[i].Status = newStatus
			updateMissionStatusInDictionary(Missions[i])
			return
		}
	}
	fmt.Println("Mission Name doesnt match")
}

func updateMissionStatusInDictionary(mission Mission) {
	var avengerDetails *[]Avenger
	for key, value := range AvengersMissions {
		if key.Name == mission.Name {
			avengerDetails = value
			delete(AvengersMissions, key)
			break
		}
	}
	AvengersMissions[mission] = avengerDetails
}

func getAvengerStatusAndMissions() {
	fmt.Println("Name          Status            Assigned Mission")
	for _, avenger := range Avengers {
		fmt.Print(avenger.Name)
		if avenger.MissionsAssigned > 0 {
			fmt.Print("         On Mission")
		} else {
			fmt.Print("        Available")
		}
		fmt.Println("            ", strings.Join(getMissionNameForAvenger(avenger.Name), ", "))
	}
}

func getMissionNameForAvenger(avengerName string) []string {
	var missions []string
	for key, value := range AvengersMissions {
		for _, avenger := range *value {
			if avenger.Name == avengerName && key.Status == "Assigned" {
				missions = append(missions, key.Name)
			}
		}
	}
	return missions
}
