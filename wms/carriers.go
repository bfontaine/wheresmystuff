package wms

type Carrier interface {
	GetName() string
	MatchPackage(p PackageID) bool
	GetPackageInfo(p PackageID) (*PackageInfo, error)
}

type GenericCarrier struct {
	name           string
	matchPackage   func(p PackageID) bool
	getPackageInfo func(p PackageID) (*PackageInfo, error)
}

func (c GenericCarrier) GetName() string { return c.name }
func (c GenericCarrier) MatchPackage(p PackageID) bool {
	if c.matchPackage == nil {
		return false
	}
	return c.matchPackage(p)
}
func (c GenericCarrier) GetPackageInfo(p PackageID) (*PackageInfo, error) {
	return c.getPackageInfo(p)
}

var _ Carrier = GenericCarrier{}

var carriers = make(map[string]Carrier)

func RegisterCarrier(c Carrier) {
	carriers[c.GetName()] = c
}

func GetCarrier(name string) (Carrier, bool) {
	c, ok := carriers[name]
	if !ok {
		return nil, false
	}
	return c, true
}

func GetCarriers() (cs []Carrier) {
	for _, c := range carriers {
		cs = append(cs, c)
	}

	return
}

func GetCarrierForPackage(p PackageID) (Carrier, error) {
	for _, c := range carriers {
		if c.MatchPackage(p) {
			return c, nil
		}
	}

	return nil, ErrUnknownCarrier
}
