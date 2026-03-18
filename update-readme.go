package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/etree"
)

// https://stackoverflow.com/questions/36530251/time-since-with-months-and-years
func diff(a, b time.Time) (year, month, day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func updateAgeData(newAgeData, file string) {
	doc := etree.NewDocument()

	if err := doc.ReadFromFile(file); err != nil {
		log.Fatal(err)
	}

	if el := doc.FindElement("//tspan[@id='age_data']"); el != nil {
		el.SetText(newAgeData)
	} else {
		log.Println("Element not found")
	}

	if err := doc.WriteToFile(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	birthday, _ := time.Parse(time.RFC3339, "1995-11-02T15:40:00+02:00")
	year, month, day := diff(birthday, time.Now())

	updateAgeData(fmt.Sprintf("%d years, %d months, %d days", year, month, day), "light_mode.svg")
	updateAgeData(fmt.Sprintf("%d years, %d months, %d days", year, month, day), "dark_mode.svg")

	log.Println("Finished uptime update of light_mode.svg and dark_mode.svg")
}
