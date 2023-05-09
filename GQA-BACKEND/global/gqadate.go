package global

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type GqaDate time.Time

func (d *GqaDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*d = GqaDate(vt)
	default:
		return errors.New("time scan error")
	}
	return nil
}

func (d GqaDate) Value() (driver.Value, error) {
	tTime := time.Time(d)
	return tTime.Format(time.DateOnly), nil
}

func (d GqaDate) String() string {
	return time.Time(d).Format(time.DateOnly)
}

func (d GqaDate) GormDataType() string {
	return "date"
}

func (d GqaDate) GobEncode() ([]byte, error) {
	return time.Time(d).GobEncode()
}

func (d *GqaDate) GobDecode(b []byte) error {
	return (*time.Time)(d).GobDecode(b)
}

func (d GqaDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(d).Format(time.DateOnly))
	return []byte(formatted), nil
}

func (d *GqaDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(time.DateOnly, timeStr)
	*d = GqaDate(t1)
	return err
}
