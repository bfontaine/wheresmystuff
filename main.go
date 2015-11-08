package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bfontaine/wheresmystuff/wms"
)

func main() {
	var flagName string
	var flagList bool

	var info *wms.PackageInfo
	var err error

	flag.StringVar(&flagName, "c", "", "Force to use a carrier")
	flag.BoolVar(&flagList, "list", false, "List the available carriers")
	flag.Parse()

	if flagList {
		fmt.Println("Available carriers:")
		for _, c := range wms.GetCarriers() {
			fmt.Printf("- %s\n", c.GetShortName())
		}
		return
	}

	pkg := wms.PackageID(flag.Arg(0))
	if pkg == "" {
		fmt.Fprintln(os.Stderr, "Please give me a package ID")
		os.Exit(1)
	}

	if flagName != "" {
		c, ok := wms.GetCarrier(flagName)
		if !ok {
			fmt.Fprintf(os.Stderr, "The carrier '%s' doesn't exist.\n", flagName)
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
