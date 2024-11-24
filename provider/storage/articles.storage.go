package articles_storage

import (
	"sync"
	"time"
)

// MarkdownData represents the structure of the Markdown file data.
type ArticlesData struct {
	Title        string    `json:"title"`
	TitleSlug    string    `json:"title_slug"`
	Description  string    `json:"description"`
	CreatedDate  time.Time `json:"created_date"`
	UpdateDate   time.Time `json:"update_date"`
	BannerImage  string    `json:"banner_image"`
	Author       string    `json:"author"`
	Category     string    `json:"category"`
	Location     []string  `json:"location"`
	IsPinned     bool      `json:"is_pinned"`
	CategorySlug string    `json:"category_slug"`
	FilePath     string    `json:"file_path"`
}

type ArticlesMarkdown struct {
	ArticlesData *ArticlesData `json:"articles_data"`
	Html         string
}

var (
	ArticlesStorage *[]ArticlesData
	ArticlesContent *map[string]ArticlesMarkdown
	storageOnce     sync.Once
	contentOnce     sync.Once
)

func InitArticlesStorage() {
	storageOnce.Do(func() {
		ArticlesStorage = &[]ArticlesData{}
	})
}

func InitArticlesContent() {
	contentOnce.Do(func() {
		temp := make(map[string]ArticlesMarkdown)
		ArticlesContent = &temp
	})
}
