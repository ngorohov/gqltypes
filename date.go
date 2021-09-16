package gqltypes

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

const dateLayout = "2006-01-02"

var ErrWrongDate = errors.New("date must be in YYYY-MM-DD format")

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
	return []byte(strconv.Quote(fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day))), nil
}

func (d *Date) UnmarshalJSON(v []byte) (err error) {
	str, err := strconv.Unquote(string(v))
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", str)
	if err != nil {
		return ErrWrongDate
	}

	*d = FromTime(parsedDate)

	return nil
}

// UnmarshalGQL subj
func (d *Date) UnmarshalGQL(v interface{}) (err error) {
	str, ok := v.(string)
	if !ok {
		return ErrWrongDate
	}

	var parsedDate time.Time

	if parsedDate, err = time.Parse(dateLayout, str); err != nil {
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
