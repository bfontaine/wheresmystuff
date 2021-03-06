package wms

import "errors"

// A Carrier is some entity which knows where a package is. It has a
// human-readable name and two methods: one to check if it might be responsible
// for a given package ID; another to actually get the info for this package.
type Carrier interface {
	GetName() string
	GetShortName() string
	MatchPackage(p PackageID) bool
	GetPackageInfo(p PackageID) (*PackageInfo, error)
}

// GenericCarrier is a convenient way to declare a new carrier.
type GenericCarrier struct {
	Name      string
	ShortName string
	Match     func(p PackageID) bool
	GetInfo   func(p PackageID) (*PackageInfo, error)
}

// GetName implements the Carrier interface.
func (c GenericCarrier) GetName() string { return c.Name }

// GetShortName implements the Carrier interface.
func (c GenericCarrier) GetShortName() string { return c.ShortName }

// MatchPackage implements the Carrier interface.
func (c GenericCarrier) MatchPackage(p PackageID) bool {
	if c.Match == nil {
		return false
	}
	return c.Match(p)
}

// GetPackageInfo implements the Carrier interface.
func (c GenericCarrier) GetPackageInfo(p PackageID) (*PackageInfo, error) {
	if c.GetInfo == nil {
		return nil, errors.New("undefined func attribute: GetInfo")
	}
	return c.GetInfo(p)
}

var _ Carrier = GenericCarrier{}

var carriers = make(map[string]Carrier)

// RegisterCarrier registers a carrier in the global register
func RegisterCarrier(c Carrier) {
	carriers[c.GetShortName()] = c
}

// GetCarrier gets a carrier by name from the global register
func GetCarrier(name string) (Carrier, bool) {
	c, ok := carriers[name]
	if !ok {
		return nil, false
	}
	return c, true
}

// GetCarriers returns a slice of all the carriers in the global register
func GetCarriers() (cs []Carrier) {
	for _, c := range carriers {
		cs = append(cs, c)
	}

	return
}
