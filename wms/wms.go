package wms

import "errors"

var (
	// ErrUnknownCarrier is returned when no carrier can match a given package
	// ID
	ErrUnknownCarrier = errors.New("Unknown carrier for this package")

	// ErrUnknownPackage is returned when a carrier doesn't know a package
	ErrUnknownPackage = errors.New("Unknown package for this carrier")
)

// PackageID is a package ID
type PackageID string

func (p PackageID) String() string {
	// we define a String() method so if we decide to change the type later we
	// won't have to change all the `string(p)` calls in the code because
	// `p.String()` will still work.
	return string(p)
}

// GetPackageInfo tries to find the carrier for a given package and return the
// package info for the given ID
func GetPackageInfo(p PackageID) (*PackageInfo, error) {
	for _, c := range carriers {
		if !c.MatchPackage(p) {
			continue
		}

		info, err := c.GetPackageInfo(p)
		if err == ErrUnknownPackage {
			continue
		}
		if err != nil {
			return nil, err
		}
		return info, nil
	}

	return nil, ErrUnknownCarrier
}
