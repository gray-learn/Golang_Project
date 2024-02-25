package main  // Declares the package name for the Go program.

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Recipe struct to store recipe information
type Recipe struct {
	URL         string
	Name        string
	Specifications Specifications
	Ingredients []map[string]string
}

// Specifications struct to store recipe specifications
type Specifications struct {
	Difficulty   string
	PrepTime     string
	CookingTime  string
	ServingSize  string
	PriceTier    string
}

func main() {  // Start of the main function.
	// Check if URL argument is provided
	if len(os.Args) < 2 {  // Check if command-line arguments are provided.
		fmt.Println("Please provide a URL")  // Print a message if URL is not provided.
		return
	}

	url := os.Args[1]  // Retrieve URL from command-line arguments.
	collector := colly.NewCollector()  // Create a new Colly collector.

	var recipes []Recipe  // Declare a slice to store recipes.

	// Print the visited URL before making a request
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)  // Print the visited URL.
	})

	// Extract recipe data from the HTML
	collector.OnHTML("main", func(main *colly.HTMLElement) {
		recipe := Recipe{  // Initialize a new recipe struct.
			URL: url,
		}

		// Find and assign recipe title
		recipe.Name = main.ChildText(".gz-title-recipe")  // Extract and assign recipe name.
		fmt.Println("Scraping recipe for:", recipe.Name)  // Print the name of the scraped recipe.

		// Extract recipe specifications
		main.ForEach(".gz-name-featured-data", func(i int, specListElement *colly.HTMLElement) {
			specText := specListElement.Text  // Extract text content of specification element.
			if strings.Contains(specText, "Difficoltà: ") {  // Check if specification contains difficulty information.
				recipe.Specifications.Difficulty = specListElement.ChildText("strong")  // Extract and assign difficulty.
			}
			if strings.Contains(specText, "Preparazione: ") {  // Check if specification contains preparation time information.
				recipe.Specifications.PrepTime = specListElement.ChildText("strong")  // Extract and assign preparation time.
			}
			if strings.Contains(specText, "Cottura: ") {  // Check if specification contains cooking time information.
				recipe.Specifications.CookingTime = specListElement.ChildText("strong")  // Extract and assign cooking time.
			}
			if strings.Contains(specText, "Dosi per: ") {  // Check if specification contains serving size information.
				recipe.Specifications.ServingSize = specListElement.ChildText("strong")  // Extract and assign serving size.
			}
			if strings.Contains(specText, "Costo: ") {  // Check if specification contains price tier information.
				recipe.Specifications.PriceTier = specListElement.ChildText("strong")  // Extract and assign price tier.
			}
		})

		// Extract recipe ingredients
		ingredients := make(map[string]string)  // Initialize a map to store ingredients.
		main.ForEach(".gz-ingredient", func(i int, ingredient *colly.HTMLElement) {
			ingredients[ingredient.ChildText("a")] = ingredient.ChildText("span")  // Extract and assign ingredients.
		})
		recipe.Ingredients = append(recipe.Ingredients, ingredients)  // Append ingredients to recipe.

		recipes = append(recipes, recipe)  // Append recipe to the recipes slice.
	})

	// Print response URL after receiving a response
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)  // Print response URL.
	})

	// Handle errors
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error:", e)  // Print error message.
	})

	// Start scraping
	collector.Visit(url)  // Initiate the scraping process.

	// Print scraped recipes
	printRecipes(recipes)  // Print the scraped recipes.
}

// Function to print recipes
func printRecipes(recipes []Recipe) {
	fmt.Println("Scraped Recipes:")  // Print header for scraped recipes.
	for _, recipe := range recipes {  // Iterate over each recipe.
		fmt.Printf("Recipe Name: %s\n", recipe.Name)  // Print recipe name.
		fmt.Printf("URL: %s\n", recipe.URL)  // Print recipe URL.
		fmt.Printf("Difficulty: %s\n", recipe.Specifications.Difficulty)  // Print recipe difficulty.
		fmt.Printf("Preparation Time: %s\n", recipe.Specifications.PrepTime)  // Print recipe preparation time.
		fmt.Printf("Cooking Time: %s\n", recipe.Specifications.CookingTime)  // Print recipe cooking time.
		fmt.Printf("Serving Size: %s\n", recipe.Specifications.ServingSize)  // Print recipe serving size.
		fmt.Printf("Price Tier: %s\n", recipe.Specifications.PriceTier)  // Print recipe price tier.

	
		fmt.Println()  // Print an empty line to separate recipes.
	}
}
