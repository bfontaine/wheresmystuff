package wms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageInfoString(t *testing.T) {
	p := PackageInfo{Info: "yo", PackageID: "xyz"}

	assert.Equal(t, "Package xyz: yo\n", p.String())
}
