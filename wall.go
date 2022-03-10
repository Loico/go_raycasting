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

// readMap reads JSON file that stores the map
func readMap() {
	// Add walls on around the sreen
	var mapBoundary wall = wall{coordinates{-1, -1}, coordinates{-1, mapHeight}}
	walls = append(walls, mapBoundary)
	mapBoundary = wall{coordinates{-1, mapHeight}, coordinates{mapWidth, mapHeight}}
	walls = append(walls, mapBoundary)
	mapBoundary = wall{coordinates{mapWidth, mapHeight}, coordinates{mapWidth, -1}}
	walls = append(walls, mapBoundary)
	mapBoundary = wall{coordinates{mapWidth, -1}, coordinates{-1, -1}}
	walls = append(walls, mapBoundary)

	jsonFile, err := os.Open("map.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	type JSONA struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}

	type JSONB struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}

	type JSONWall struct {
		A JSONA `json:"a"`
		B JSONB `json:"b"`
	}

	type JSONWalls struct {
		Walls []JSONWall `json:"walls"`
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var m JSONWalls
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
