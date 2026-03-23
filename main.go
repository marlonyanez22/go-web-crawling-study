package main

import (
	"fmt"
	"strings"
	"time"
)

// Product representa la estructura de un dato extraído de eCommerce
type Product struct {
	Name  string
	Price float64
	URL   string
}

func main() {
	// Simulamos una lista de "Raw Data" extraída por un crawler
	rawScrapedData := []string{
		"Smartphone Samsung | $899.99 | https://store.com/s24",
		"Laptop Pro 16 | $1250.00 | https://store.com/laptop",
		"Wireless Headphones | $150.50 | https://store.com/audio",
	}

	fmt.Println("--- TrackStreet Crawler Monitor System ---")
	fmt.Printf("Started at: %v\n\n", time.Now().Format(time.RFC822))

	var processedProducts []Product

	// Procesamos cada línea simulando la limpieza de selectores (Troubleshooting)
	for _, entry := range rawScrapedData {
		parts := strings.Split(entry, " | ")
		
		if len(parts) == 3 {
			// Simulación de conversión de datos (Data Parsing)
			var price float64
			fmt.Sscanf(parts[1], "$%f", &price)

			p := Product{
				Name:  parts[0],
				Price: price,
				URL:   parts[2],
			}
			processedProducts = append(processedProducts, p)
		}
	}

	// Mostramos los resultados finales como un profesional
	for _, p := range processedProducts {
		fmt.Printf("[SUCCESS] Scraped: %-20s | Price: %8.2f | Source: %s\n", p.Name, p.Price, p.URL)
	}

	fmt.Printf("\nTotal items processed: %d\n", len(processedProducts))
	fmt.Println("Status: All crawlers healthy.")
}
