package weather_time

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var _ fmt.Stringer = &WeatherTime{}
var _ json.Unmarshaler = &WeatherTime{}

type WeatherTime struct {
	t time.Time
	d time.Duration
}

func (w *WeatherTime) String() string {
	location, err := time.LoadLocation("MST")
	if err != nil {
		fmt.Println(err)
	}

	return w.t.
		In(location).
		Format(time.Kitchen)
}

func (w *WeatherTime) UnmarshalJSON(data []byte) error {
	rawString := string(data)
	rawString = strings.TrimPrefix(rawString, `"`)
	rawString = strings.TrimSuffix(rawString, `"`)
	timeAndDuration := strings.Split(rawString, "/")

	if len(timeAndDuration) != 2 {
		return errors.New(fmt.Sprintf("expected time to be 2 parts, but was %d: [%s]", len(timeAndDuration), rawString))
	}

	t, err := time.Parse(time.RFC3339, timeAndDuration[0])
	if err != nil {
		return err
	}

	d, err := w.parseDuration(timeAndDuration[1])
	if err != nil {
		return err
	}

	w.t = t
	w.d = d

	return nil
}

func (w *WeatherTime) Day() string {
	location, err := time.LoadLocation("MST")
	if err != nil {
		fmt.Println(err)
	}

	return w.t.
		In(location).
		Format(time.DateOnly)
}

func (w *WeatherTime) Midnight() time.Time {
	location, err := time.LoadLocation("MST")
	if err != nil {
		fmt.Println(err)
	}

	return w.t.
		In(location).
		Truncate(time.Hour)
}

func (w *WeatherTime) Hour() int {
	location, err := time.LoadLocation("MST")
	if err != nil {
		fmt.Println(err)
	}

	return w.t.
		In(location).
		Hour()
}

func (w *WeatherTime) Duration() time.Duration {
	return w.d
}

func (w *WeatherTime) IsAfterNoon() bool {
	return w.t.Hour() >= 12
}

// parseDuration uses ISO8601 Duration Standard: https://www.digi.com/resources/documentation/digidocs/90001488-13/reference/r_iso_8601_duration_format.htm
func (w *WeatherTime) parseDuration(durString string) (time.Duration, error) {
	if !strings.HasPrefix(durString, "P") {
		return 0, fmt.Errorf("expected duration string to start with [P], but was [%s]", durString)
	}
	durString = durString[1:]

	designator := durString[0]
	switch designator {
	case 'T':
		return parseTime(durString[1:])
	default:
		// calendarDesignator
		calendarString := strings.SplitN(durString, "T", 2)

		duration := time.Duration(0)
		num := ""
		for _, c := range calendarString[0] {
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num = fmt.Sprintf("%s%c", num, c)
			case 'Y': // year
				// How many days are in a year? lol
				return 0, fmt.Errorf("year based duration are unsupported")
			case 'M': // month
				// How many days are in a month?  lol
				return 0, fmt.Errorf("month based duration are unsupported")
			case 'W': // week
				numWeeks, err := strconv.Atoi(num)
				if err != nil {
					return 0, fmt.Errorf("invalid weeks parsing duration: [%s]", num)
				}
				duration += time.Hour * 24 * 7 * time.Duration(numWeeks)
				num = ""
			case 'D': // day
				numDays, err := strconv.Atoi(num)
				if err != nil {
					return 0, fmt.Errorf("invalid days parsing duration: [%s]", num)
				}
				duration += time.Hour * 24 * time.Duration(numDays)
				num = ""
			}
		}
		if len(calendarString) == 2 {
			timeDuration, err := parseTime(calendarString[1])
			if err != nil {
				return 0, err
			}
			duration += timeDuration
		}

		return duration, nil
	}
}

func parseTime(timeString string) (time.Duration, error) {
	timeString = strings.ReplaceAll(timeString, "H", "h")
	timeString = strings.ReplaceAll(timeString, "M", "m")
	timeString = strings.ReplaceAll(timeString, "S", "s")

	return time.ParseDuration(timeString)
}
