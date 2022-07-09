package main

import (
	"encoding/xml"

	svg "github.com/metaleap/go-xsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
	"github.com/metaleap/go-xsd/xsd-makepkg/tests"
	udevgo "github.com/wwsheng009/go-util/udev/go"
)

type SvgDoc struct {
	XMLName xml.Name `xml:"svg"`
	svg.TsvgType
}

func main() {
	var (
		dirBasePath  = udevgo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-svg")
		makeEmptyDoc = func() interface{} { return &SvgDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
