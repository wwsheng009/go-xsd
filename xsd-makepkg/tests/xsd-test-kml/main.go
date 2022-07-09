package main

import (
	"encoding/xml"

	kml "github.com/metaleap/go-xsd-pkg/schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go"
	"github.com/metaleap/go-xsd/xsd-makepkg/tests"
	udevgo "github.com/wwsheng009/go-util/udev/go"
)

type KmlDoc struct {
	XMLName xml.Name `xml:"http://www.opengis.net/kml/2.2 kml"`
	kml.TKmlType
}

func main() {
	var (
		dirBasePath  = udevgo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-kml")
		makeEmptyDoc = func() interface{} { return &KmlDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
