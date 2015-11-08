package wms

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/bfontaine/gostruct"
)

func init() {
	// unfortunately it doesn't seem we can shorten this URL
	const targetURL = "http://www.chronopost.fr/transport-express/livraison-colis/accueil/suivi?appid=9680_718&appparams=http%3A%2F%2Fwww.chronopost.fr%3A54711%2Fwebclipping%2Fservlet%2Fwebclip%3Fjahia_url_web_clipping%3Dhttp%3A%2F%2Flocalhost%3A54702%2Fexpedier%2FinputLTNumbers.do"

	// http://www.chronopost.fr/transport-express/webdav/site/chronov4/users/chronopost/public/pdf/track.pdf
	var pattern = regexp.MustCompile(`^(?:[a-z]{2}\d{9}[a-z]{2}|\d{15}|\d{14}[a-z])$`)

	RegisterCarrier(GenericCarrier{
		Name:      "Chronopost",
		ShortName: "chronopost",
		Match: func(p PackageID) bool {
			return pattern.MatchString(p.String())
		},
		GetInfo: func(p PackageID) (*PackageInfo, error) {
			resp, err := http.PostForm(targetURL, url.Values{"chronoNumbers": {p.String()}})
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

			var text string
			if len(parts) > 2 {
				text = parts[1]
			} else {
				text = result.Info
			}

			info := PackageInfo{
				PackageID: p,
				Info:      text,
			}

			return &info, nil
		},
	})
}
