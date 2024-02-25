package main

import (
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

// Recipe struct to store recipe information
type Recipe struct {
	URL            string
	Title          string
	Name           string
	Ingredients    []string
	IngredientsUnit []string
}

func main() {
	// Check if URL argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL")
		return
	}

	url := os.Args[1]
	collector := colly.NewCollector()

	var recipe Recipe

	// Extract recipe data from the HTML
	collector.OnHTML(".container", func(e *colly.HTMLElement) {
		// Find and assign recipe title
		recipe.Name = e.ChildText("h1")
		fmt.Println("Scraping recipe for:", url)
		// Find and assign recipe title
		recipe.Name = e.ChildText("h1")
		// Print recipe details
		fmt.Println("Recipe Name:", recipe.Name)

		// Extract recipe ingredients and their units
		e.ForEach(".recipe-details-ingredients .ingredient-name", func(i int, ingredient *colly.HTMLElement) {
			recipe.Ingredients = append(recipe.Ingredients, ingredient.Text)
		})
		e.ForEach(".recipe-details-ingredients .ingredient-unit", func(i int, unit *colly.HTMLElement) {
			recipe.IngredientsUnit = append(recipe.IngredientsUnit, unit.Text)
		})

		fmt.Println("Ingredients:")
		for i, ingredient := range recipe.Ingredients {
			fmt.Printf(" %s: %s --", ingredient, recipe.IngredientsUnit[i])
		}
	})

	// Handle errors
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error:", e)
	})

	// Start scraping
	collector.Visit(url)
}