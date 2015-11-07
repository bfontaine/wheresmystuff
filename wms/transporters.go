package wms

type Transporter interface {
	GetName() string
	MatchPackage(p PackageID) bool
	GetPackageInfo(p PackageID) (*PackageInfo, error)
}

type GenericTransporter struct {
	name           string
	matchPackage   func(p PackageID) bool
	getPackageInfo func(p PackageID) (*PackageInfo, error)
}

func (t GenericTransporter) GetName() string { return t.name }
func (t GenericTransporter) MatchPackage(p PackageID) bool {
	if t.matchPackage == nil {
		return false
	}
	return t.matchPackage(p)
}
func (t GenericTransporter) GetPackageInfo(p PackageID) (*PackageInfo, error) {
	return t.getPackageInfo(p)
}

var _ Transporter = GenericTransporter{}

var transporters = make(map[string]Transporter)

func RegisterTransporter(t Transporter) {
	transporters[t.GetName()] = t
}

func GetTransporter(name string) (Transporter, bool) {
	t, ok := transporters[name]
	if !ok {
		return nil, false
	}
	return t, true
}

func GetTransporters() (ts []Transporter) {
	for _, t := range transporters {
		ts = append(ts, t)
	}

	return
}

func GetTransporterForPackage(p PackageID) (Transporter, error) {
	for _, t := range transporters {
		if t.MatchPackage(p) {
			return t, nil
		}
	}

	return nil, ErrUnknownTransporter
}
