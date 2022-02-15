package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type wall struct {
	a coordinates
	b coordinates
}

var walls []wall

func readMap() {
	// Add walls on around the sreen
	var screenBoundary wall = wall{coordinates{-1, -1}, coordinates{-1, screenHeight}}
	walls = append(walls, screenBoundary)
	screenBoundary = wall{coordinates{-1, screenHeight}, coordinates{screenWidth, screenHeight}}
	walls = append(walls, screenBoundary)
	screenBoundary = wall{coordinates{screenWidth, screenHeight}, coordinates{screenWidth, -1}}
	walls = append(walls, screenBoundary)
	screenBoundary = wall{coordinates{screenWidth, -1}, coordinates{-1, -1}}
	walls = append(walls, screenBoundary)

	jsonFile, err := os.Open("map.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	type JsonA struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}

	type JsonB struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}

	type JsonWall struct {
		A JsonA `json:"a"`
		B JsonB `json:"b"`
	}

	type JsonWalls struct {
		Walls []JsonWall `json:"walls"`
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var m JsonWalls
	if err := json.Unmarshal(byteValue, &m); err != nil {
		log.Fatal(err)
	}

	for _, jsonW := range m.Walls {
		var w wall
		w.a.x = jsonW.A.X
		w.a.y = jsonW.A.Y
		w.b.x = jsonW.B.X
		w.b.y = jsonW.B.Y
		walls = append(walls, w)
	}
}
