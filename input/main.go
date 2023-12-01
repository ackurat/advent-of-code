package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/adamliliemark/advent-of-code/utils"
)

func fetchInput(year, day int) []byte {
	sessionString := utils.ReadFileToString("./input/session")
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: sessionString,
	}
	req.AddCookie(&sessionCookie)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return body
}

func writeToFile(filename string, contents []byte) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}
	err = os.WriteFile(filename, contents, os.FileMode(0644))
	if err != nil {
		log.Fatalf("writing file: %s", err)
	}
}

func ParseFlags() (day, year int) {
	today := time.Now()
	flag.IntVar(&day, "day", today.Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "year", today.Year(), "AOC year")
	flag.Parse()

	if day > 25 || day < 1 {
		log.Fatalf("day out of range: %d", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	return day, year
}

func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}

func main() {
	day, year := ParseFlags()
	body := fetchInput(year, day)
	filename := filepath.Join(Dirname(), "../", fmt.Sprintf("%d/day%02d/input.txt", year, day))
	writeToFile(filename, body)

	fmt.Println("Wrote to file: ", filename)

	fmt.Println("Done!")
}
