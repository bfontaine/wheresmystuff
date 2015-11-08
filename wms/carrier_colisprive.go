package wms

import (
	"net/url"
	"regexp"

	"github.com/bfontaine/gostruct"
)

func init() {
	const urlPrefix = "https://www.colisprive.com/moncolis/pages/detailColis.aspx?numColis="

	// this is from my (limited) experience and some posts on the Web. The last
	// 5 digits are the postal code.
	var pattern = regexp.MustCompile(`^5\d{2}000\d{11}$`)

	RegisterCarrier(GenericCarrier{
		Name:      "Colis Privé",
		ShortName: "colisprive",
		Match: func(p PackageID) bool {
			return pattern.MatchString(p.String())
		},
		GetInfo: func(p PackageID) (*PackageInfo, error) {
			var info struct {
				Info string `gostruct:".BandeauInfoColis .divStatut .tdText"`
			}

			err := gostruct.Fetch(&info, urlPrefix+url.QueryEscape(p.String()))
			if err != nil {
				return nil, err
			}

			return &PackageInfo{
				PackageID: p,
				Info:      info.Info,
			}, nil
		},
	})
}
