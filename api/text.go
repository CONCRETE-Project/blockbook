package api

import (
	"strings"

	"github.com/gobuffalo/packr"
)

// Text contains static overridable texts used in explorer
var Text struct {
	BlockbookAbout, TOSLink string
}

func init() {
	box := packr.NewBox("../build/text")
	if about, err := box.MustString("about"); err == nil {
		Text.BlockbookAbout = strings.TrimSpace(about)
	} else {
		panic(err)
	}
}
