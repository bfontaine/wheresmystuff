package wms

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

var ErrUnknownCarrier = errors.New("Unknown transporter for this package")

type PackageID string

// WIP
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
