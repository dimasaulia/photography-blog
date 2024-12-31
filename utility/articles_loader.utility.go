package utility

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	s "blog/provider/storage"

	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

const dateTimeFormat = "2006-01-02T15:04:05"

func LoadAllData(path string) error {

	// Buka directory nya
	dir, err := os.Open(path)
	if err != nil {
		return err
	}

	// Ambil Seluruh File di Directory
	fileInDir, err := dir.ReadDir(-1)
	if err != nil {
		return err
	}

	fmt.Println("============ DAFTAR SELURUH FILE ============")
	for _, file := range fileInDir {
		if file.Type().IsRegular() {
			fmt.Println(fmt.Sprintf("File Name %v", file.Name()))
			filePath := filepath.Join("./markdown_files", file.Name())

			content, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			context := parser.NewContext()
			if err := md.Convert(content, &buf, parser.WithContext(context)); err != nil {
				return err
			}

			metaData := meta.Get(
				context,
			)

			createdDateParsedTime, err := time.Parse(dateTimeFormat, fmt.Sprintf("%v", metaData["created_date"]))
			if err != nil {
				return err
			}

			updateDateParsedTime, err := time.Parse(dateTimeFormat, fmt.Sprintf("%v", metaData["update_date"]))
			if err != nil {
				return err
			}

			var articleMeta s.ArticlesData = s.ArticlesData{
				Title:        fmt.Sprintf("%v", metaData["title"]),
				TitleSlug:    fmt.Sprintf("%v", metaData["title_slug"]),
				Description:  fmt.Sprintf("%v", metaData["description"]),
				CreatedDate:  createdDateParsedTime,
				UpdateDate:   updateDateParsedTime,
				BannerImage:  fmt.Sprintf("%v", metaData["banner_image"]),
				Author:       fmt.Sprintf("%v", metaData["author"]),
				Category:     fmt.Sprintf("%v", metaData["category"]),
				Location:     strings.Split(fmt.Sprintf("%v", metaData["location"]), ","),
				IsPinned:     metaData["is_pinned"].(bool),
				IsPublish:    metaData["is_publish"].(bool),
				CategorySlug: fmt.Sprintf("%v", metaData["category_slug"]),
				FilePath:     fmt.Sprintf("./markdown_files/%v", file.Name()),
			}

			if articleMeta.IsPublish {
				*s.ArticlesStorage = append(*s.ArticlesStorage, articleMeta)

				var markdown s.ArticlesMarkdown = s.ArticlesMarkdown{
					Html:         buf.String(),
					ArticlesData: &articleMeta,
				}
				(*s.ArticlesContent)[articleMeta.TitleSlug] = markdown
				fmt.Println(fmt.Sprintf("Slug %v", articleMeta.TitleSlug))
			}
		}
	}

	return nil
}
