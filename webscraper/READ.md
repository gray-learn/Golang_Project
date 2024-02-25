# Recipe Scraper


## Description

This project is a Go-based web scraper that extracts detailed recipe information from specified URLs. It leverages the Colly web scraping library to navigate web pages and gather data such as the recipe's name, specifications (difficulty, preparation time, cooking time, serving size, and price tier), and ingredients.

## Features

- Extract recipe name and URL.
- Gather detailed recipe specifications.
- List ingredients with quantities.

## Getting Started

### Dependencies

- Go (version 1.15 or newer recommended).
- Colly web scraping library.

### Environment Setup and Installation

1. **Install Go**: Ensure Go is installed on your system. For Windows users, you can install Go using Chocolatey:

    ```bash
    choco install golang
    ```

    For users on other operating systems, please visit the [official Go website](https://golang.org/dl/) for installation instructions.

2. **Verify Go Installation**: Ensure Go is correctly installed and accessible via your PATH by running:

    ```bash
    go version
    ```

    You should see the installed version of Go printed to the console.

3. **Clone the Repository**: Clone the Recipe Scraper project to your local machine using:

    ```bash
    git clone https://github.com/gray-learn/Golang_Project.git
    cd webscraper    
    ```

4. **Initialize the Go Module**: Set up a new Go module for your project (if not already done):

    ```bash
    go mod init webscraper
    ```

    Or, if you're setting up the project under a specific domain:

    ```bash
    go mod init oxylabs.io/web-scraping-with-go
    ```

5. **Install Dependencies**: Install the required Go dependencies, including Colly:

    ```bash
    go mod tidy
    ```

### Usage

To run the scraper, use the following command:

```bash
go run ingredient_lunch.go https://icook.tw/recipes/403213
go run ingredient_dessert.go https://ricette.giallozafferano.it/Schiacciata-fiorentina.html