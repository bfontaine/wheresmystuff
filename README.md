# Where’s My Stuff?

**wheresmystuff** is a personal command-line utility to check where’s the stuff
I ordered at Amazon the day before.

[![GoDoc](https://godoc.org/github.com/bfontaine/wheresmystuff?status.svg)](https://godoc.org/github.com/bfontaine/wheresmystuff)
[![Build Status](https://travis-ci.org/bfontaine/wheresmystuff.svg?branch=master)](https://travis-ci.org/bfontaine/wheresmystuff)

Right now the most important thing is to know where my stuff is so don’t expect
fancy code.

## Usage

    $ ./whereismystuff <package ID>

## Supported Carriers

* Chronopost France
* Colis Privé (France)

## Install

    go get github.com/bfontaine/wheresmystuff

## API

You can check where’s my^Wyour stuff with `wms.GetPackageInfo`:

```go
info, err := wms.GetPackageInfo(wms.PackageID("XYZ123456789"))
```

### Extensions

Adding a new carrier is easy:

1. Create a type implementing the `wms.Carrier` interface
2. Call `wms.Register()` with an instance of your type
3. ???
4. Profit!

Since most carriers follow the same pattern you can use `wms.GenericCarrier`
for an even easier experience:

```go
wms.RegisterCarrier(wms.GenericCarrier{
    Name: "The Name",
    ShortName: "thename",
    Match: func(p PackageID) bool {
        return len(p.String()) == 12
    },

    GetInfo: func(p PackageID) (*PackageInfo, error) {
        // request the carrier website
        // ...
        return &PackageInfo{Info: "something something"}, nil
    },
})
```
