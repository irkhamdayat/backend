package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

func NewNullDecimalFromFloatPtr(val *float64) decimal.NullDecimal {
	if val == nil {
		return decimal.NullDecimal{}
	}

	return decimal.NullDecimal{Valid: val != nil, Decimal: decimal.NewFromFloat(*val)}
}

func NewNullUUIDFromStringPtr(val *string) uuid.NullUUID {
	if val == nil {
		return uuid.NullUUID{}
	}

	UUID, err := uuid.Parse(*val)
	if err != nil {
		logrus.WithField("val", val).Error(err)
		return uuid.NullUUID{}
	}

	return uuid.NullUUID{Valid: val != nil, UUID: UUID}
}

// StringLabelToValue to convert string ex: "super admin" to system name "SUPER_ADMIN"
// specific format to value format
func StringLabelToValue(vesselTypeSpecific string) string {
	// Remove any leading/trailing spaces
	vesselTypeSpecific = strings.TrimSpace(vesselTypeSpecific)
	// Replace non-alphanumeric characters (except spaces) with a single underscore
	vesselTypeSpecific = regexp.MustCompile(`[^a-zA-Z0-9\s]`).ReplaceAllString(vesselTypeSpecific, "_")
	// Convert to uppercase
	vesselTypeSpecific = strings.ToUpper(vesselTypeSpecific)
	// Replace spaces with underscores
	vesselTypeSpecific = strings.ReplaceAll(vesselTypeSpecific, " ", "_")
	// Replace multiple underscores with a single underscore
	vesselTypeSpecific = regexp.MustCompile(`_+`).ReplaceAllString(vesselTypeSpecific, "_")
	// Remove special characters at the end of the sequence
	vesselTypeSpecific = regexp.MustCompile(`[_-]+$`).ReplaceAllString(vesselTypeSpecific, "")
	return vesselTypeSpecific
}

// EstimateReadDuration calculates the estimated reading duration in seconds for a given text.
// The goal is to provide a time estimate based on the text's word count and a predefined words-per-minute rate.
func EstimateReadDuration(text string, avgWordPerMinute float64) int64 {
	words := strings.Split(text, " ")

	nWord := len(words)

	est := float64(nWord) / avgWordPerMinute

	s := fmt.Sprintf("%v", est)
	split := strings.Split(s, ".")
	decimalStr := "0." + split[1]
	estSecond, _ := strconv.ParseFloat(decimalStr, 64)
	estMinute, _ := strconv.ParseInt(split[0], 10, 0)
	finalEst := estMinute*60 + int64(estSecond*0.6*100.0)

	if finalEst <= 0 {
		return 1
	}

	return finalEst
}

func ParseReadDeadline(readDeadline string, startTime time.Time) (null.Time, error) {
	parts := strings.Fields(readDeadline)

	if len(parts) != 2 {
		return null.Time{}, fmt.Errorf("invalid input format")
	}

	value, err := strconv.Atoi(parts[0])
	if err != nil {
		return null.Time{}, err
	}

	unit := parts[1]
	switch unit {
	case "DAYS":
		return null.TimeFrom(startTime.AddDate(0, 0, value)), nil
	case "WEEKS":
		return null.TimeFrom(startTime.AddDate(0, 0, 7*value)), nil
	case "MONTHS":
		return null.TimeFrom(startTime.AddDate(0, value, 0)), nil
	case "YEARS":
		return null.TimeFrom(startTime.AddDate(value, 0, 0)), nil
	default:
		return null.Time{}, fmt.Errorf("unsupported unit: %s", unit)
	}
}

func GetLastDayOfPreviousMonth(t time.Time) time.Time {
	firstDayOfCurrentMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	lastDayOfPreviousMonth := firstDayOfCurrentMonth.Add(-time.Second)
	return lastDayOfPreviousMonth
}

func GetMonthDiff(start, end time.Time) int {
	months := (end.Year()-start.Year())*12 + int(end.Month()) - int(start.Month())
	return months
}

func IsValidDay(year, month, day int) bool {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date.Year() == year && int(date.Month()) == month && date.Day() == day
}

func GetDate(inputTime time.Time) time.Time {
	return time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 0, 0, 0, 0, inputTime.Location())
}
