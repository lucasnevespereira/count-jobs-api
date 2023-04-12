package collector

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"net/url"
)

// LinkedInSource is a type that represents a LinkedIn job listing source.
type LinkedInSource struct{}

// Name returns the name of the LinkedIn job listing source.
func (s *LinkedInSource) Name() string {
	return "linkedin"
}

func (s *LinkedInSource) Collect(ctx context.Context, term string, location string, jobChan chan<- Job) error {
	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.linkedin.com"),
	)

	// Set user agent
	c.UserAgent = userAgents[rand.Intn(len(userAgents))]

	// Visit search page
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL.String())
	})

	c.OnHTML("ul.jobs-search__results-list", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			job := Job{
				Title:    el.ChildText("h3 a"),
				Location: el.ChildText("span.job-search-card__location"),
				Company:  el.ChildText("a.job-search-card__subtitle"),
				URL:      el.ChildAttr("h3 a", "href"),
			}
			jobChan <- job
		})
	})

	url := fmt.Sprintf("https://www.linkedin.com/jobs/search/?keywords=%s&location=%s", url.QueryEscape(term), url.QueryEscape(location))
	c.Visit(url)

	return nil
}

// NewLinkedInSource creates a new LinkedIn job listing source.
func NewLinkedInSource() Source {
	return &LinkedInSource{}
}
