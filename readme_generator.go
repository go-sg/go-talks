package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type talk struct {
	Topic   string `json:"topic"`
	Speaker string `json:"speaker"`
	Slides  string `json:"slides"`
	Code    string `json:"code"`
	Video   string `json:"video"`
	Meetup  meetup `json:"meetup"`
}

type meetup struct {
	Date string `json:"date"`
	Link string `json:"link"`
}

const headerString string = "# Go Talks\n" +
	"Click on the ✓ to see the slides/code/video.\n\n" +
	"| Topic | Speaker | Slides | Code | Video | Meetup |\n" +
	"| --- | --- |:---:|:---:|:---:|:---:|\n"

const mark string = "✓"

func main() {
	goTalks := getDataFromJSON("talks.json")

	bodyStr := generateBodyStr(goTalks)

	ioutil.WriteFile("README.md", []byte(headerString+bodyStr), 0644)
}

func getDataFromJSON(jsonFileName string) []talk {
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var goTalks []talk
	err = json.Unmarshal(byteVal, &goTalks)
	if err != nil {
		panic(err)
	}

	return goTalks
}

func generateBodyStr(talks []talk) string {
	tableStr := ""
	for _, t := range talks {
		tableStr += fmt.Sprintf("| %v ", t.Topic) + fmt.Sprintf("| %v ", t.Speaker)
		tableStr += getCellsText(mark, t.Slides, t.Code, t.Video)
		tableStr += fmt.Sprintf("| [%v](%v) |\n", t.Meetup.Date, t.Meetup.Link)
	}

	return tableStr
}

func getCellsText(mark string, fields ...string) string {
	cellsTxt := ""
	for _, f := range fields {
		if f != "" {
			cellsTxt += fmt.Sprintf("| [%v](%v) ", mark, f)
		} else {
			cellsTxt += "| "
		}
	}
	return cellsTxt
}
