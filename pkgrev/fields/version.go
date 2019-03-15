package fields

import (
	"database/sql/driver"
	"errors"

	"github.com/hashicorp/go-version"
)

type Version struct {
	*version.Version
}

func (v *Version) Value() (driver.Value, error) {
	if v.Version == nil {
		return "", nil
	}
	return v.Version.String(), nil
}

func (v *Version) IsZero() bool {
	return v.Version == nil
}

func (v *Version) Scan(src interface{}) (err error) {
	switch t := src.(type) {
	case string:
		if t != "" {
			v.Version, err = version.NewVersion(t)
		}
		return
	case []byte:
		return v.Scan(string(t))
	default:
		return errors.New("Version.Scan: invalid type")
	}
}
