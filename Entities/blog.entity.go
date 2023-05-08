package blogEntity

import "time"

type Blog struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	WebsiteUrl  string    `json:"websiteUrl"`
	CreatedAt   time.Time `json:"createdAt" format:"2006-01-02T15:04:05Z07:00"`
}

type BlogInput struct {
	Name        string `json:"name" validate:"required,max=15,min=1"`
	Description string `json:"description" validate:"required,min=1,max=500"`
	WebsiteUrl  string `json:"websiteUrl" validate:"required,url"`
}

func (b *Blog) Update(name, description, websiteUrl string) {
	b.Name = name
	b.Description = description
	b.WebsiteUrl = websiteUrl
}

func New(name, description, websiteUrl string) Blog {
	return Blog{
		"1",
		name,
		description,
		websiteUrl,
		time.Now(),
	}
}

func NewFromDto(dto BlogInput) Blog {
	name, description, websiteUrl := dto.Name, dto.Description, dto.WebsiteUrl
	return New(name, description, websiteUrl)
}
