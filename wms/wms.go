package wms

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

// ErrUnknownCarrier is returned when no carrier can match a given package ID
var ErrUnknownCarrier = errors.New("Unknown transporter for this package")

// PackageID is a package ID
type PackageID string

func (p PackageID) String() string {
	// we define a String() method so if we decide to change the type later we
	// won't have to change all the `string(p)` calls in the code because
	// `p.String()` will still work.
	return string(p)
}

// PackageInfo contains info describing the state of a package (WIP)
type PackageInfo struct {
	PackageID  PackageID
	LastUpdate time.Time
	Info       string
}

func (pi PackageInfo) String() string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Package %s: %s\n", pi.PackageID, pi.Info)

	if !pi.LastUpdate.IsZero() {
		fmt.Fprintf(&buf, "LastUpdate: %v\n", pi.LastUpdate)
	}

	return buf.String()
}
