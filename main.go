package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Schedule struct {
	Date             time.Time
	SpiritualThought string
	Message          string
	Game             string
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: fhe <year> <month> <people>")
		return
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid year: ", os.Args[1])
		return
	}
	month, err := strconv.Atoi(os.Args[2])
	if err != nil || month < 1 || month > 12 {
		fmt.Println("invalid month: ", os.Args[2])
		return
	}
	people := strings.Split(os.Args[3], ",")
	shuffle(people)

	sundays := sundays(year, month)

	schedules := []Schedule{}
	for i, sunday := range sundays {
		schedule := Schedule{
			Date:             sunday,
			SpiritualThought: people[i%len(people)],
			Message:          people[(i+1)%len(people)],
			Game:             people[(i+2)%len(people)],
		}
		schedules = append(schedules, schedule)
	}

	prettyPrint(schedules)
}

func sundays(year int, month int) []time.Time {
	sundays := []time.Time{}
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	for date.Month() == time.Month(month) {
		if date.Weekday() == time.Sunday {
			sundays = append(sundays, date)
		}
		date = date.AddDate(0, 0, 1)
	}
	return sundays
}

func prettyPrint(schedules []Schedule) {
	for _, schedule := range schedules {
		fmt.Printf("Date: %v\n", schedule.Date.Format("2006-01-02"))
		fmt.Printf("Spiritual Thought: %s\n", schedule.SpiritualThought)
		fmt.Printf("Message: %s\n", schedule.Message)
		fmt.Printf("Game: %s\n", schedule.Game)
		fmt.Println()
	}
}

func shuffle(people []string) {
	rand := rand.New(rand.NewSource(time.Now().Unix()))

	rand.Shuffle(len(people), func(i, j int) {
		people[i], people[j] = people[j], people[i]
	})
}
