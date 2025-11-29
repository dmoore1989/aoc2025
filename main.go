package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// UPDATE THE YEAR HERE TO PULL THE CORRECT DATA
const YEAR string = "2025"

const (
	AOC_LINK_FOR_YEAR  string      = "https://adventofcode.com/%s/day/%s"
	ERR_MISSING_COOKIE string      = "Puzzle inputs differ by user.  Please log in to get your puzzle input."
	ERR_MISSING_DAY    string      = "Please don't repeatedly request this endpoint before it unlocks! The calendar countdown is synchronized with the server time; the link will be enabled on the calendar the instant this puzzle becomes available."
	FILE_PERMISSION    os.FileMode = 0777
)

func main() {
	var day string
	args := os.Args
	if len(args) > 1 {
		day = args[1]
	} else {
		day = strconv.Itoa(time.Now().Day())
	}
	body, err := getAOCData(day)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = writeInputFile(body, day)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = writeCodeFile(day)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = exec.Command("open", fmt.Sprintf(AOC_LINK_FOR_YEAR, YEAR, day)).Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func getAOCData(day string) ([]byte, error) {
	_, err := os.ReadFile(fmt.Sprintf("lib/day%s/main.go", day))
	if err == nil {
		return nil, fmt.Errorf("day already exists")
	}

	aocCookie := os.Getenv("AOC_COOKIE")
	email := os.Getenv("EMAIL")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(AOC_LINK_FOR_YEAR, YEAR, strings.Join([]string{day, "/input"}, "")), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", email)
	req.AddCookie(&http.Cookie{Name: "session", Value: aocCookie})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func writeInputFile(body []byte, day string) error {
	switch strings.TrimSpace(string(body)) {
	case ERR_MISSING_COOKIE:
		return fmt.Errorf("invalid cookie")
	case ERR_MISSING_DAY:
		return fmt.Errorf("not time for new puzzle")
	}

	err := os.WriteFile(fmt.Sprintf("lib/day%s.txt", day), body, FILE_PERMISSION)
	if err != nil {
		return err
	}
	return nil
}

func writeCodeFile(day string) error {
	if !hasCmdDir() {
		err := os.Mkdir("cmd", FILE_PERMISSION)
		if err != nil {
			return err
		}
	}

	err := os.Mkdir(fmt.Sprintf("cmd/day%s", day), FILE_PERMISSION)
	if err != nil {
		return err
	}

	templateData, err := os.ReadFile("lib/aoc_template.txt")
	if err != nil {
		return err
	}

	templateString := fmt.Sprintf(string(templateData), day)
	err = os.WriteFile(fmt.Sprintf("cmd/day%s/main.go", day), []byte(templateString), FILE_PERMISSION)
	if err != nil {
		return err
	}

	return nil
}

func hasCmdDir() bool {
	info, err := os.Stat("cmd")
	if err != nil {
		return false
	}

	return info.IsDir()
}
