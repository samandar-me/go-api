package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	DecodeJson()
}

type club struct {
	ShortName string
	League    string
	Place     int
	Ball      int
	FullName  string
	Players   []string
}

func EncodeJson() {
	clubs := []club{
		{"Tot", "Apl", 3, 26, "Tottenham HotSpur", []string{"Kane", "Son"}},
		{"Rma", "Laliga", 2, 30, "Real Madrid", []string{"Benzema", "Vini"}},
		{"FCB", "Bundesliga", 2, 30, "Bayern Munich", []string{"Sane", "Mane"}},
		{"MUN", "Apl", 5, 23, "Man Utd", []string{"Ronaldo", "Fernandesh"}},
		{"PSG", "Liga 1", 3, 30, "Paris Saint-Germain", []string{"Neymar", "Messi"}},
	}

	finalJson, err := json.MarshalIndent(clubs, "", "\t")

	printError(err)
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
			{
                "ShortName": "Tot",
                "League": "Apl",
                "Place": 3,
                "Ball": 26,
                "FullName": "Tottenham HotSpur",
                "Players": ["Kane","Son"]
        	}`,
	)
	var myClub club
	isValid := json.Valid(jsonData)
	if isValid {
		fmt.Println("Valid")
		json.Unmarshal(jsonData, &myClub)
		fmt.Printf("%#v\n", myClub)
	} else {
		fmt.Println("Json isn't valid")
	}
}

func printError(err error) {
	if err != nil {
		panic(err)
	}
}
