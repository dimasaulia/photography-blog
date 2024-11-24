package utility

import (
	"sync"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithExtensions(extension.NewFootnote()),
	goldmark.WithExtensions(meta.Meta),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		// html.WithHardWraps(),
		html.WithXHTML(),
		html.WithUnsafe(),
	),
)

var MD *goldmark.Markdown
var once sync.Once

func InitGoldmark() {
	once.Do(func() {
		MD = &md
	})
}
