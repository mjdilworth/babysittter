package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting")

	fmt.Println(calculateCost(1,1,1))
}

func calculateCost(startTime int, endTime int, bedTime int) int {

	iRet := 0

	return iRet;
}
/*

The babysitter
- starts no earlier than 5:00PM
- leaves no later than 4:00AM
- gets paid $12/hour from start-time to bedtime
- gets paid $8/hour from bedtime to midnight
- gets paid $16/hour from midnight to end of job
/*