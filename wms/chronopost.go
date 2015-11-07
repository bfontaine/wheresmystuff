package wms

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/bfontaine/gostruct"
)

func init() {

	const targetURL = "http://www.chronopost.fr/transport-express/livraison-colis/accueil/suivi?appid=9680_718&appparams=http%3A%2F%2Fwww.chronopost.fr%3A54711%2Fwebclipping%2Fservlet%2Fwebclip%3Fjahia_url_web_clipping%3Dhttp%3A%2F%2Flocalhost%3A54702%2Fexpedier%2FinputLTNumbers.do"

	RegisterCarrier(GenericCarrier{
		name: "Chronopost",
		matchPackage: func(p PackageID) bool {
			return true
		},
		getPackageInfo: func(p PackageID) (*PackageInfo, error) {
			resp, err := http.PostForm(targetURL, url.Values{"chronoNumbers": {string(p)}})
			if err != nil {
				return nil, err
			}

			var result struct {
				Info string `gostruct:"div.numeroColi2"`
			}

			if err := gostruct.PopulateFromResponse(&result, resp); err != nil {
				return nil, err
			}

			parts := strings.Split(result.Info, "\"")

			info := PackageInfo{
				PackageID: p,
				Info:      parts[1],
			}

			return &info, nil
		},
	})
}
