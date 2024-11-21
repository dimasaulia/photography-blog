package criticalcss_utility

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var BaseCss *string

func LoadCss(filename string) (string, error) {
	filePath := filepath.Join("./public/css/", filename+".css")
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", errors.New("file tidak ditemukan")
	}
	minify := strings.ReplaceAll(string(content), "\n", "")
	minify = strings.ReplaceAll(minify, "\t", "")
	minify = strings.ReplaceAll(minify, "\r", "")
	minify = strings.ReplaceAll(minify, " ", "")
	minify = strings.ReplaceAll(minify, "  ", "")
	minify = strings.TrimSpace(minify)
	return minify, nil
}
