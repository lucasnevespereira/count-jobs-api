package collector

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

// Source is a type that represents a job listing source.
type Source interface {
	Name() string
	Collect(ctx context.Context, term string, location string, jobChan chan<- Job) error
}

// Job is a type that represents a job listing.
type Job struct {
	Title    string
	Location string
	Company  string
	URL      string
}

// Collector is a type that represents a job listing collector.
type Collector struct {
	sources []Source
}

// NewCollector creates a new job listing collector.
func NewCollector(sources ...Source) *Collector {
	return &Collector{
		sources: sources,
	}
}

// Start starts the job listing collector and returns a channel of jobs.
func (c *Collector) Start(ctx context.Context, term string, location string) (<-chan Job, error) {
	jobChan := make(chan Job)

	var wg sync.WaitGroup
	wg.Add(len(c.sources))
	for _, source := range c.sources {
		go func(s Source) {
			err := s.Collect(ctx, term, location, jobChan)
			if err != nil {
				// handle error
			}
			wg.Done()
		}(source)
	}

	wg.Wait()
	close(jobChan)

	return jobChan, nil
}

// SourceNotFoundError is an error that indicates that a job listing source was not found.
type SourceNotFoundError struct {
	name string
}

func (e *SourceNotFoundError) Error() string {
	return errors.Errorf("source %q not found", e.name).Error()
}

// GetSource returns a job listing source by name.
func GetSource(name string) (Source, error) {
	switch name {
	case "indeed":
		return NewIndeedSource(), nil
	case "linkedin":
		return NewLinkedInSource(), nil
	default:
		return nil, &SourceNotFoundError{name: name}
	}
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/B08C390",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
	// TODO: ajouter d'autres
}
