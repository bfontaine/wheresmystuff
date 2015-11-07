package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bfontaine/wheresmystuff/wms"
)

func main() {
	var name string
	var c wms.Carrier

	flag.StringVar(&name, "t", "", "Carrier")
	flag.Parse()

	pkg := wms.PackageID(flag.Arg(0))
	if pkg == "" {
		fmt.Fprintln(os.Stderr, "Please give me a package ID")
		os.Exit(1)
	}

	if name != "" {
		c, _ = wms.GetCarrier(name)
		if c == nil {
			fmt.Fprintf(os.Stderr, "The transporter '%s' doesn't exist.\n", name)
			os.Exit(2)
		}
	} else {
		c, _ = wms.GetCarrierForPackage(pkg)
		if c == nil {
			fmt.Fprintln(os.Stderr, "I couldn't find a transporter for this package.")
			fmt.Fprintln(os.Stderr, "You can give its name with the -t option.")
			os.Exit(3)
		}
	}

	info, err := c.GetPackageInfo(pkg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while retrieving package info: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("%s: %s\n", pkg, info.Info)
}
