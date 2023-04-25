package main

import "testing"

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
var avenger2 = Avenger{
	Name:              "Captain America",
	Abilities:         []string{"Abilities: Highly Intelligent Suit of Armor"},
	MissionsAssigned:  2,
	MissionsCompleted: 0,
}
var mission = Mission{
	Name:    "TestMission",
	Details: "TestDetails",
	Status:  "Assigned",
}
var mission1 = Mission{
	Name:    "TestMission1",
	Details: "TestDetails1",
	Status:  "Completed",
}

func TestGetNamesForAvengers(t *testing.T) {
	//Arrange
	insertAvengers(avenger)
	insertAvengers(avenger1)
	avengersAssigned := &[]Avenger{
		avenger1,
	}

	//Act
	avengerNames := getNamesFromAvengers(avengersAssigned)

	//Assert
	if avengerNames[0] != avenger1.Name {
		t.Errorf("Expected %s got %s", avengerNames[0], avenger1.Name)
	}

	//cleanup
	Avengers = []Avenger{}
}

func TestGetMissionNameForAvenger(t *testing.T) {
	//Arrange
	AvengersMissions[mission] = &[]Avenger{avenger}

	//Act
	missionNames := getMissionNameForAvenger("Tony Stark")

	//Assert
	if missionNames[0] != mission.Name {
		t.Errorf("Expected %s got %s", missionNames[0], mission.Name)
	}

	//Cleanup
	AvengersMissions = make(map[Mission]*[]Avenger)
}

func TestUpdateMissionStatusInDictionary(t *testing.T) {
	//Arrange
	AvengersMissions[mission] = &[]Avenger{avenger}
	missionCompleted := Mission{
		Name:    mission.Name,
		Details: mission.Details,
		Status:  "Completed",
	}

	//Act
	updateMissionStatusInDictionary(missionCompleted)

	//Assert
	if _, ok := AvengersMissions[missionCompleted]; !ok {
		t.Errorf("Mission Status is not updated")
	}

	//Cleanup
	AvengersMissions = make(map[Mission]*[]Avenger)
}

func TestIsValidAvenger(t *testing.T) {
	//Arrange
	insertAvengers(avenger1)

	//Act
	isValid, avengers := isValidAvenger([]string{avenger1.Name})

	//Assert
	if !isValid {
		t.Error("Expected the avenger to be valid")
	}
	if len(avengers) == 0 {
		t.Error("Avengers should have been obtained")
	}
	if avengers[0].Name != avenger1.Name {
		t.Error("Avengers doesnt match")
	}

	//cleanup
	Avengers = []Avenger{}
}

func TestUpdateMissionStatusAvenger(t *testing.T) {

	t.Run("MissionCompleted", func(t *testing.T) {
		//Arrange
		avenger := Avenger{
			Name:              "Thor",
			Abilities:         []string{"test"},
			MissionsAssigned:  2,
			MissionsCompleted: 0,
		}
		insertAvengers(avenger)
		previousCountAssigned := avenger.MissionsAssigned
		previousCountCompleted := avenger.MissionsCompleted

		//Act
		updateMissionStatusAvenger([]Avenger{avenger}, true)

		//Assert
		if Avengers[0].MissionsAssigned != previousCountAssigned-1 {
			t.Errorf("expected count %d got %d", previousCountAssigned-1, avenger.MissionsAssigned)
		}
		if Avengers[0].MissionsCompleted != previousCountCompleted+1 {
			t.Errorf("expected count %d got %d", previousCountCompleted+1, avenger.MissionsCompleted)
		}

		//cleanup
		Avengers = []Avenger{}
	})

	t.Run("MissionAssigned", func(t *testing.T) {
		//Arrange
		avenger := Avenger{
			Name:              "Thor",
			Abilities:         []string{"test"},
			MissionsAssigned:  1,
			MissionsCompleted: 0,
		}
		insertAvengers(avenger)
		previousCountAssigned := avenger.MissionsAssigned
		previousCountCompleted := avenger.MissionsCompleted

		//Act
		updateMissionStatusAvenger([]Avenger{avenger}, false)

		//Assert
		if Avengers[0].MissionsAssigned != previousCountAssigned+1 {
			t.Errorf("expected count %d got %d", previousCountAssigned+1, avenger.MissionsAssigned)
		}
		if Avengers[0].MissionsCompleted != previousCountCompleted {
			t.Errorf("expected count %d got %d", previousCountCompleted, avenger.MissionsCompleted)
		}

		//cleanup
		Avengers = []Avenger{}
	})

}

func TestIsAvengerAvailable(t *testing.T) {
	t.Run("AvengerIsAssigned", func(t *testing.T) {
		//Arrange
		insertAvengers(avenger2)

		//Act
		isAvengerAvailable, name := isAvengerAvailable([]string{avenger2.Name})

		//Assert
		if isAvengerAvailable {
			t.Error("Expected the Avenger to be assigned")
		}
		if name != avenger2.Name {
			t.Errorf("Expected the name to be %s", avenger2.Name)
		}
	})

	t.Run("AvengerIsAvailable", func(t *testing.T) {
		//Arrange
		insertAvengers(avenger1)

		//Act
		isAvengerAvailable, name := isAvengerAvailable([]string{avenger1.Name})

		//Assert
		if !isAvengerAvailable {
			t.Error("Expected the Avenger to be available")
		}
		if name != "" {
			t.Errorf("Expected the name to be empty")
		}
	})

	//cleanup
	Avengers = []Avenger{}
}
