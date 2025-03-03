package models

type News struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	CategoryID       uint   `json:"category_id"`
	IsCrawl          bool   `json:"is_crawl"`
	AuhtorCrawl      string `json:"author_crawl"`
	SourceCrawl      string `json:"source_crawl"`
	Title            string `json:"title"`
	Slug             string `json:"slug"`
	Image            string `json:"image"`
	ImageDescription string `json:"image_description"`
	Excerpt          string `json:"excerpt"`
	PublishedAt      string `json:"published_at"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`

	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}
