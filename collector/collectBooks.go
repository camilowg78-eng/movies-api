package collector

import (
	"fmt"
	"strconv"
	"strings"

	"movies-api/models"

	"github.com/gocolly/colly"
)

func CollectorB() ([]models.Book, error) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	var books []models.Book

	c.OnHTML("article.product_pod", func(e *colly.HTMLElement) {

		// Título completo desde el atributo title del <a> dentro de h3
		title := e.ChildAttr("h3 a", "title")

		// Precio
		price := strings.TrimSpace(e.ChildText("p.price_color"))

		convertedPrice, err := parsePrice(price)
		if err != nil {
			fmt.Printf("Error parsing price for book '%s': %v\n", title, err)
			return
		}

		book := models.Book{
			Title: title,
			Price: convertedPrice,
		}
		books = append(books, book)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://books.toscrape.com/")
	return books, nil
}

func parsePrice(priceStr string) (float64, error) {
	// Remove the currency symbol and any whitespace
	priceStr = strings.TrimSpace(priceStr)
	priceStr = strings.TrimPrefix(priceStr, "£")

	// Convert the string to a float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse price '%s': %v", priceStr, err)
	}

	return price, nil
}
