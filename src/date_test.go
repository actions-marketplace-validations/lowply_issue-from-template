package main

import (
	"testing"
	"time"
)

type testCase struct {
	now    string
	should string
}

func TestNewDate(t *testing.T) {
	testCases := []testCase{
		// Monday when Jan 1st is Monday
		{now: "2018-01-01T00:00:00Z", should: "2018 Week 01"},
		// Monday when Jan 1st is Tuesday
		{now: "2018-12-31T00:00:00Z", should: "2019 Week 01"},
		// Monday when Jan 1st is Wednesday
		{now: "2019-12-30T00:00:00Z", should: "2020 Week 01"},
		// Monday when Jan 1st is Thursday
		{now: "2025-12-29T00:00:00Z", should: "2026 Week 01"},
		// Monday when Jan 1st is Friday
		{now: "2020-12-28T00:00:00Z", should: "2020 Week 53"},
		// Monday when Jan 1st is Saturday
		{now: "2021-12-27T00:00:00Z", should: "2021 Week 52"},
		// Monday when Jan 1st is Saturday and it's a leap year
		{now: "2032-12-27T00:00:00Z", should: "2032 Week 53"},
		// Monday when Jan 1st is Sunday
		{now: "2022-12-26T00:00:00Z", should: "2022 Week 52"},
	}

	for _, v := range testCases {
		t.Logf("Testing %v...", v.now)
		now, err := time.Parse(time.RFC3339, v.now)
		if err != nil {
			t.Fatal(err)
		}
		d := NewDate(now)
		current := d.WeekNumberYear + " Week " + d.WeekNumber
		if current != v.should {
			t.Errorf("Actual: %v, Should: %v\n", current, v.should)
		}
	}
}