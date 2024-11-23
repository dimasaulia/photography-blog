package converter

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

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
		html.WithHardWraps(),
		html.WithXHTML(),
		html.WithUnsafe(),
	),
)

func MarkdownConverter(filename string) (string, error) {
	filePath := filepath.Join("./markdown_files", filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", errors.New("file tidak ditemukan")
	}

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := md.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return "", errors.New("error ketika konversi markdown")
	}

	return buf.String(), nil
}
