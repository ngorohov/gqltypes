package gqltypes

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

var (
	ErrWrongWeekday   = errors.New("invalid Weekday given")
	ErrWrongEnumValue = errors.New("enums must be strings")
)

type Weekday time.Weekday

const (
	WeekdaySunday    = Weekday(time.Sunday)
	WeekdayMonday    = Weekday(time.Monday)
	WeekdayTuesday   = Weekday(time.Tuesday)
	WeekdayWednesday = Weekday(time.Wednesday)
	WeekdayThursday  = Weekday(time.Thursday)
	WeekdayFriday    = Weekday(time.Friday)
	WeekdaySaturday  = Weekday(time.Saturday)

	Sunday    = "SUNDAY"
	Monday    = "MONDAY"
	Tuesday   = "TUESDAY"
	Wednesday = "WEDNESDAY"
	Thursday  = "THURSDAY"
	Friday    = "FRIDAY"
	Saturday  = "SATURDAY"
)

func weekDayFromString(str string) (Weekday, error) {
	switch str {
	case Sunday:
		return WeekdaySunday, nil
	case Monday:
		return WeekdayMonday, nil
	case Tuesday:
		return WeekdayTuesday, nil
	case Wednesday:
		return WeekdayWednesday, nil
	case Thursday:
		return WeekdayThursday, nil
	case Friday:
		return WeekdayFriday, nil
	case Saturday:
		return WeekdaySaturday, nil
	}

	return Weekday(-1), ErrWrongWeekday
}

func (e Weekday) String() string {
	switch e {
	case WeekdaySunday:
		return Sunday
	case WeekdayMonday:
		return Monday
	case WeekdayTuesday:
		return Tuesday
	case WeekdayWednesday:
		return Wednesday
	case WeekdayThursday:
		return Thursday
	case WeekdayFriday:
		return Friday
	case WeekdaySaturday:
		return Saturday
	default:
		panic("invalid Weekday value")
	}
}

// UnmarshalGQL subj
func (e *Weekday) UnmarshalGQL(v interface{}) (err error) {
	str, ok := v.(string)
	if !ok {
		return ErrWrongEnumValue
	}

	*e, err = weekDayFromString(str)

	return err
}

// MarshalGQL subj
func (e Weekday) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(e.String()))
}
