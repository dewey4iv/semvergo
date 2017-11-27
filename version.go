package semvergo

import (
	"fmt"
	"strconv"
	"strings"
)

// Version is a type
type Version struct {
	Major  uint
	Minor  uint
	Patch  uint
	Branch string
	Meta   string
}

// NewFromString returns a new Version
func NewFromString(s string) (*Version, error) {
	return fromString(s)
}

// String implements Stringer
func (v Version) String() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)

	if strings.TrimSpace(v.Branch) != "" {
		s += fmt.Sprintf("-%s", v.Branch)
	}

	return s
}

func fromString(s string) (*Version, error) {
	var v Version

	a := strings.Split(s, "-")
	if len(a) == 2 {
		v.Branch = a[1]
	}

	parts := strings.Split(a[0], ".")
	if len(parts) == 3 {
		major, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid major: %v", err)
		}

		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid minor: %v", err)
		}

		patch, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid patch: %v", err)
		}

		v.Major = uint(major)
		v.Minor = uint(minor)
		v.Patch = uint(patch)
	}

	return &v, nil
}
