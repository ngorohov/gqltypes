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
	ErrWrongDate      = errors.New("date must be in YYYY-MM-DD format")
)

type Weekday time.Weekday

const (
	// WeekdaySunday subj
	WeekdaySunday = Weekday(time.Sunday)
	// WeekdayMonday subj
	WeekdayMonday = Weekday(time.Monday)
	// WeekdayTuesday subj
	WeekdayTuesday = Weekday(time.Tuesday)
	// WeekdayWednesday subj
	WeekdayWednesday = Weekday(time.Wednesday)
	// WeekdayThursday subj
	WeekdayThursday = Weekday(time.Thursday)
	// WeekdayFriday subj
	WeekdayFriday = Weekday(time.Friday)
	// WeekdaySaturday subj
	WeekdaySaturday = Weekday(time.Saturday)

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

//isValid subj
//func (e Weekday) isValid() bool {
//	switch e {
//	case WeekdayMonday, WeekdayTuesday, WeekdayWednesday, WeekdayThursday, WeekdayFriday, WeekdaySaturday, WeekdaySunday:
//		return true
//	}
//	return false
//}

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

// Date graphql date type
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// MarshalGQL subj
func (d Date) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)))
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)), nil
}

// UnmarshalGQL subj
func (d *Date) UnmarshalGQL(v interface{}) (err error) {
	str, ok := v.(string)
	if !ok {
		return ErrWrongDate
	}

	var parsedDate time.Time

	if parsedDate, err = time.Parse("2006-01-02", str); err != nil {
		return ErrWrongDate
	}

	*d = FromTime(parsedDate)

	return nil
}

// FromTime subj
func FromTime(t time.Time) Date {
	return Date{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}
}
