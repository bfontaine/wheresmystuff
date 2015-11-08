package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bfontaine/wheresmystuff/wms"
)

func main() {
	var name string
	var info *wms.PackageInfo
	var err error

	flag.StringVar(&name, "t", "", "Carrier")
	flag.Parse()

	pkg := wms.PackageID(flag.Arg(0))
	if pkg == "" {
		fmt.Fprintln(os.Stderr, "Please give me a package ID")
		os.Exit(1)
	}

	if name != "" {
		c, ok := wms.GetCarrier(name)
		if !ok {
			fmt.Fprintf(os.Stderr, "The transporter '%s' doesn't exist.\n", name)
			os.Exit(2)
		}

		info, err = c.GetPackageInfo(pkg)
	} else {
		info, err = wms.GetPackageInfo(pkg)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while retrieving package info: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("%s: %s\n", pkg, info.Info)
}
