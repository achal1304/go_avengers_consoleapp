package main

import (
	"bufio"
	"fmt"
	"os"
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

func getNamesFromAvengers(avengers *[]Avenger) []string {
	var avengersAssigned []string
	for _, avenger := range *avengers {
		avengersAssigned = append(avengersAssigned, avenger.Name)
	}
	return avengersAssigned
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
