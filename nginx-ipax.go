package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type LogRecord struct {
	Time        int64
	Date        string
	IP          string
	AccessCount int
}

type ByTime []*LogRecord

func main() {

	inputFileNames := getFileNames()
	// Map to hold records
	logMap := make(map[string]*LogRecord)
	lineCount := 0

	// Cycle through provided files
	for _, inputFileName := range inputFileNames {
		// Open up log file
		logFile, err := os.Open(inputFileName)
		checkErr(err)

		scanner := bufio.NewScanner(logFile)

		// Loop through file
		for scanner.Scan() {
			line := scanner.Text()
			ip := getLogIP(line)
			date := getLogDate(line)
			mapKey := date + ":" + ip

			// If record exists increment count
			// otherwise create record
			if record, ok := logMap[mapKey]; ok {
				record.AccessCount++
			} else {
				time, err := time.Parse("02/Jan/2006", date)
				checkErr(err)
				logMap[mapKey] = &LogRecord{
					Date:        time.Format("2006-01-02"),
					Time:        time.Unix(),
					IP:          ip,
					AccessCount: 1,
				}
			}

			lineCount++
		}

		// Ensure file reader has no errors
		if err = scanner.Err(); err != nil {
			checkErr(err)
		}

		logFile.Close()
	}

	printResults(logMap)
}

func getLogIP(line string) string {
	return line[:strings.Index(line, " -")]
}

func getLogDate(line string) string {
	start := strings.Index(line, "[") + 1
	end := strings.Index(line, ":")
	return line[start:end]
}

func printResults(logMap map[string]*LogRecord) {

	// Move records into slice
	records := make([]*LogRecord, len(logMap))
	index := 0
	for _, record := range logMap {
		records[index] = record
		index++
	}
	// Sort records
	sort.Sort(ByTime(records))
	fmt.Println("Date, IP, Access Count")
	for _, record := range records {

		line := fmt.Sprintf("%s, %s, %d", record.Date, record.IP, record.AccessCount)
		fmt.Println(line)

	}
}

func getFileNames() []string {
	args := os.Args
	var fileNames []string

	if len(args) > 2 && args[1] == "run" {
		fileNames = args[2:]
	} else if len(args) > 1 {
		fileNames = args[1:]
	} else {
		errorOut("No Input File Specified")
	}
	return fileNames
}

func errorOut(message string) {
	fmt.Println("Error: " + message)
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Time sort functions
func (s ByTime) Len() int {
	return len(s)
}

func (s ByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByTime) Less(i, j int) bool {
	if s[i].Time == s[j].Time {
		return s[i].IP < s[j].IP
	}
	return s[i].Time < s[j].Time
}
