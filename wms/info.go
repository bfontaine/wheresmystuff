package wms

import (
	"bytes"
	"fmt"
	"time"
)

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
