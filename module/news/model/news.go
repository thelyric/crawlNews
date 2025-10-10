package newsmodel

type Article struct {
	Avatar      string `json:"avatar"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	PublishedAt string `json:"published_at"`
	Category    string `json:"category"`
	Source      string `json:"source"`
}

type GetArticle struct {
	Limit int `json:"limit"`
}
