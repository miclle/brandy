package compressor

import (
	"io"
	"regexp"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

// MediaType may take the form of 'text/plain', 'text/*', '*/*' or 'text/plain; charset=UTF-8; version=2.0'.
type MediaType string

const (
	// MediaTypeTextCSS is text/css media
	MediaTypeTextCSS MediaType = "text/css"
	// MediaTypeTextHTML  is text/html media
	MediaTypeTextHTML MediaType = "text/html"
	// MediaTypeTextJavaScript is text/javascript media
	MediaTypeTextJavaScript MediaType = "text/javascript"
	// MediaTypeimageSvgXML is image/svg+xml media
	MediaTypeimageSvgXML MediaType = "image/svg+xml"
)

// Minify is compress file func
func Minify(mt MediaType, w io.Writer, r io.Reader) error {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	// Or use the following for better minification of JS but lower speed:
	// m.AddCmd("text/javascript", exec.Command("java", "-jar", "build/compiler.jar"))

	return m.Minify(string(mt), w, r)
}
