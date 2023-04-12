package collector

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
	"strings"
)

// indeedSource is a type that represents an Indeed job listing source.
type indeedSource struct{}

// NewIndeedSource creates a new Indeed job listing source.
func NewIndeedSource() Source {
	return &indeedSource{}
}

// Name returns the name of the job listing source.
func (s *indeedSource) Name() string {
	return "indeed"
}

// Collect collects job listings from Indeed based on a search term and location.
func (s *indeedSource) Collect(ctx context.Context, term string, location string, jobChan chan<- Job) error {

	// Création d'un nouveau collecteur Colly
	collector := colly.NewCollector(
		colly.AllowURLRevisit(), // Permet la revisite des URL pour le cas où le site web a été mis à jour
		colly.Async(true),       // Permet de collecter plusieurs pages en parallèle
	)

	// Limite le nombre de connexions simultanées à 2 pour éviter de surcharger le site web
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
	})

	// En cas d'erreur, affiche l'URL de la requête, le code de réponse et l'erreur
	collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request URL: %s\nFailed with response code: %d\nError: %v\n", r.Request.URL, r.StatusCode, err)
	})

	// Pour chaque élément HTML avec l'ID "resultsCol"
	collector.OnHTML("#resultsCol", func(e *colly.HTMLElement) {
		// Pour chaque élément HTML avec la classe "result" dans "resultsCol"
		e.ForEach(".result", func(i int, e *colly.HTMLElement) {
			// Récupère le titre, l'URL, l'entreprise et la localisation de l'offre d'emploi
			title := e.ChildText(".title>a")
			url := e.Request.AbsoluteURL(e.ChildAttr(".title>a", "href"))
			company := e.ChildText(".company")
			location := e.ChildText(".location")

			// Envoie une nouvelle Job struct contenant les données récupérées sur le channel
			jobChan <- Job{
				Title:    title,
				URL:      url,
				Company:  company,
				Location: location,
			}
		})
	})

	// Crée un nouvel URL avec les paramètres de recherche
	q := url.Values{}
	q.Set("q", term)
	q.Set("l", location)
	q.Set("fromage", "30")
	q.Set("limit", "50")
	q.Set("sort", "date")

	var jobCount int
	var lastPage int

	// Récupère le nombre total d'offres d'emploi et calcule le nombre de pages à parcourir
	collector.OnHTML("#searchCountPages", func(e *colly.HTMLElement) {
		countString := strings.TrimSpace(e.Text)
		parts := strings.Split(countString, " ")
		if len(parts) < 3 {
			fmt.Println("Unexpected job count string format: ", countString)
			return
		}

		jobCount, _ = strconv.Atoi(strings.ReplaceAll(parts[2], ",", ""))
		perPage := 50
		lastPage = jobCount / perPage
		if jobCount%perPage > 0 {
			lastPage += 1
		}
	})

	baseURL := "https://www.indeed.com/jobs?"
	for i := 0; i < lastPage; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			q.Set("start", strconv.Itoa(i*50))
			err := collector.Visit(baseURL + q.Encode())
			if err != nil {
				return errors.Wrap(err, "error visiting Indeed URL")
			}
		}
	}

	collector.Wait()

	return nil

}
