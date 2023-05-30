package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func readCsvFile(filePath string) [][]string {
	currentPath, _ := os.Getwd()
	f, err := os.Open(currentPath + filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func main() {
	records := readCsvFile("/problems.csv")
	var correctAns *int
	correct := 0
	correctAns = &correct
	var enter string
	timeFlag := 30
	fmt.Printf("You have %d seconds to answer the questions, Please enter any key to start", timeFlag)
	fmt.Scanf("%s", &enter)
	newTimer := time.NewTimer(time.Duration(timeFlag) * time.Second)
	go func() {
		<-newTimer.C
		fmt.Println("\n Sorry out of time")
		fmt.Printf("You have got %d out of %d correct\n", *correctAns, len(records))
		os.Exit(0)
	}()
	for _, j := range records {
		var ans string
		fmt.Printf("%s ", j[0])
		fmt.Scanf("%s", &ans)
		if ans == j[1] {
			correct++
		}
	}
	fmt.Printf("You have got %d out of %d correct\n", *correctAns, len(records))
	return
}
