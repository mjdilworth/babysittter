package main

import "testing"

//Test cases
//ForOneHourBeforeBedtimeReturns12USD
//ForTwoHoursBeforeBedTimeReturns24USD()
//ForThreeHoursTillBedTimeReturns36USD()
//ForFoursHoursPastBedTimeReturns44USD()
//For530PMTo630PMReturns12USD()
//For510PMTo530PMReturns12USD()
//For510PMTo750PMReturns36USD()
//For740PMTo830PMReturns20USD()
//For10PMTo3AMReturns64USD()
//For5PMTo4AMReturns132USD()
//WhenStartingBefore5PMTill6PMReturns12USD()

func TestcalculateCost(t *testing.T) {

	//table driven test
	var tests = []struct {
		a int
		b int
		c int
		d int
	}{
		{"i", 1},

	}

	for _, test := range tests {
		want := test.d
		if got := calculateCost(test.a, test.b, test.c); got != want {
			t.Errorf("calculateCost() = %d, want %d", got, want)
		}

	}
}

/*
- starts no earlier than 5:00PM
- leaves no later than 4:00AM
- gets paid $12/hour from start-time to bedtime
- gets paid $8/hour from bedtime to midnight
- gets paid $16/hour from midnight to end of job
*/