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
	return tTime.Format("2006-01-02"), nil
}

func (d GqaDate) String() string {
	return time.Time(d).Format("2006-01-02")
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
	formatted := fmt.Sprintf("\"%v\"", time.Time(d).Format("2006-01-02"))
	return []byte(formatted), nil
}

func (d *GqaDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02", timeStr)
	*d = GqaDate(t1)
	return err
}
