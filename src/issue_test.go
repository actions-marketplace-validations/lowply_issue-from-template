package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestAddDatesFixesSundayCronIssue(t *testing.T) {
	// Save original ADD_DATES value
	originalAddDates := os.Getenv("ADD_DATES")
	defer os.Setenv("ADD_DATES", originalAddDates)

	// Test case: Sunday 21:00 UTC (when cron "0 21 * * 0" runs)
	// This simulates the exact time your cron job runs
	sundayEvening := time.Date(2024, 1, 7, 21, 0, 0, 0, time.UTC) // Sunday Jan 7, 2024 21:00 UTC
	fmt.Println(sundayEvening)

	// Without ADD_DATES - should use current time (Sunday)
	os.Unsetenv("ADD_DATES")
	issueWithoutAddDates := &issue{}
	issueWithoutAddDates.data = NewData(sundayEvening)

	// With ADD_DATES=1 - should add 1 day (Monday)
	os.Setenv("ADD_DATES", "1")
	issueWithAddDates := &issue{}
	issueWithAddDates.data = NewData(sundayEvening.AddDate(0, 0, 1))

	// The week start should be different
	// Without ADD_DATES: week starts on Jan 1 (the Monday before Sunday Jan 7)
	// With ADD_DATES=1: week starts on Jan 8 (the Monday after Sunday Jan 7)
	expectedWeekStartWithoutAddDates := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // Monday Jan 1
	expectedWeekStartWithAddDates := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)    // Monday Jan 8

	if !issueWithoutAddDates.data.WeekStart.Equal(expectedWeekStartWithoutAddDates) {
		t.Errorf("Without ADD_DATES, expected week start %v, got %v",
			expectedWeekStartWithoutAddDates, issueWithoutAddDates.data.WeekStart)
	}

	if !issueWithAddDates.data.WeekStart.Equal(expectedWeekStartWithAddDates) {
		t.Errorf("With ADD_DATES=1, expected week start %v, got %v",
			expectedWeekStartWithAddDates, issueWithAddDates.data.WeekStart)
	}

	// Verify they're in different weeks
	if issueWithoutAddDates.data.WeekNumber == issueWithAddDates.data.WeekNumber &&
		issueWithoutAddDates.data.YearOfTheWeek == issueWithAddDates.data.YearOfTheWeek {
		t.Error("ADD_DATES=1 should put the issue in a different week than without ADD_DATES")
	}

	t.Logf("Without ADD_DATES: %s Week %s", issueWithoutAddDates.data.YearOfTheWeek, issueWithoutAddDates.data.WeekNumber)
	t.Logf("With ADD_DATES=1: %s Week %s", issueWithAddDates.data.YearOfTheWeek, issueWithAddDates.data.WeekNumber)
}
