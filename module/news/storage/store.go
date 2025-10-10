package newsstorage

import (
	"context"
	"my-app/common"
	newsmodel "my-app/module/news/model"
	"net/url"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
)

type store struct{}

type Category struct {
	Name string
	Path string
}

type Source struct {
	Name string
	URL  string
}

var newsSource = []Source{
	{
		Name: "VnExpress",
		URL:  "https://vnexpress.net/rss/",
	},
	{
		Name: "Tuổi Trẻ",
		URL:  "https://tuoitre.vn/rss/",
	},
	{
		Name: "Thanh Niên",
		URL:  "https://thanhnien.vn/rss/",
	},
	{
		Name: "Dân Trí",
		URL:  "https://dantri.com.vn/rss/",
	},
	{
		Name: "VietnamNet",
		URL:  "https://vietnamnet.vn/rss/",
	},
}

var newsCategory = []Category{
	{
		Name: "Thời sự",
		Path: "thoi-su.rss",
	},
	{
		Name: "Thế giới",
		Path: "the-gioi.rss",
	},
	{
		Name: "Kinh doanh",
		Path: "kinh-doanh.rss",
	},
	{
		Name: "Giải trí",
		Path: "giai-tri.rss",
	},
	{
		Name: "Thể thao",
		Path: "the-thao.rss",
	},
}

func NewNewsStore() *store {
	return &store{}
}

func (s store) FetchLatestNews(ctx context.Context, data *newsmodel.GetArticle) ([]newsmodel.Article, error) {
	parser := gofeed.NewParser()

	results := make(chan []newsmodel.Article, len(newsSource)*len(newsCategory))
	var wg sync.WaitGroup

	for _, src := range newsSource {
		for _, cat := range newsCategory {
			wg.Add(1)
			go func(sourceName, sourceURL, category, catPath string) {
				defer wg.Done()

				base, _ := url.Parse("https://tuoitre.vn/rss/")
				rel, _ := url.Parse("kinh-doanh.rss")
				fullURL := base.ResolveReference(rel).String()
				feed, err := parser.ParseURLWithContext(fullURL, ctx)

				if err != nil {
					panic(common.NewStorageErrorResponse(err))
				}
				articles := []newsmodel.Article{}

				for i, item := range feed.Items {
					if i >= data.Limit {
						break
					}
					articles = append(articles, newsmodel.Article{
						Title:       item.Title,
						Link:        item.Link,
						PublishedAt: item.Published,
						Category:    category,
						Source:      sourceName,
					})
				}

				results <- articles
			}(src.Name, src.URL, cat.Name, cat.Path)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var final []newsmodel.Article

	timeoutCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			return final, nil
		case articles, ok := <-results:
			if !ok {
				return final, nil
			}
			final = append(final, articles...)
		}
	}
}
