package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bfontaine/wheresmystuff/wms"
)

func main() {
	var name string
	var t wms.Transporter

	flag.StringVar(&name, "t", "", "Transporter")
	flag.Parse()

	pkg := wms.PackageID(flag.Arg(0))
	if pkg == "" {
		fmt.Fprintln(os.Stderr, "Please give me a package ID")
		os.Exit(1)
	}

	if name != "" {
		t, _ = wms.GetTransporter(name)
		if t == nil {
			fmt.Fprintf(os.Stderr, "The transporter '%s' doesn't exist.\n", name)
			os.Exit(2)
		}
	} else {
		t, _ = wms.GetTransporterForPackage(pkg)
		if t == nil {
			fmt.Fprintln(os.Stderr, "I couldn't find a transporter for this package.")
			fmt.Fprintln(os.Stderr, "You can give its name with the -t option.")
			os.Exit(3)
		}
	}

	info, err := t.GetPackageInfo(pkg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while retrieving package info: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("%s: %s\n", pkg, info.Info)
}
