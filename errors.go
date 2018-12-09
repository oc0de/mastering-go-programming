package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errCrewNotFound = errors.New("Crew number not found")
var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

var ErrCrewNotFound = errors.New("Crew member not found")

type findError struct {
	Name, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	if v, ok := scMapping[name]; !ok {
		return -1, ErrCrewNotFound
	} else {
		return v, nil
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	if clearance, err := findSC("James", "Server 1"); err != nil {
		fmt.Println("Error occured while searching for clearance level: ", err)
	} else {
		fmt.Println("Clearance level found: ", clearance)
	}
}
