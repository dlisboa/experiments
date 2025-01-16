package main

type Draft struct{}
type Published struct{}
type States interface {
	Draft | Published
}

type Article[S States] struct {
	content string
}

func (a Article[S]) Content() string {
	return a.content
}

type DraftArticle Article[Draft]

// Edit sets the content of the article
func (a *DraftArticle) Edit(text string) {
	a.content = text
}

// Transitions from Draft to Published
func (a *DraftArticle) Publish() *PublishedArticle {
	return &PublishedArticle{content: a.content}
}

type PublishedArticle Article[Published]

func (p *PublishedArticle) Link() string {
	return "https://example.com/article"
}
func (p *PublishedArticle) Unpublish() *DraftArticle {
	return &DraftArticle{content: p.content}
}
