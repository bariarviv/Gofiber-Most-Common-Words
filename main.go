package main

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"sync"
)

var (
	AllWords sync.Map
	TopWords *TopN
)

func init() {
	TopWords = NewTopN(0)
}

func main() {
	// Setup the server and the routers according to the HTTP requests
	router := fiber.New()
	router.Get("/", getMostCommonWordsHandler)
	router.Post("/", getStrHandler)
	// Listen on port 3000
	router.Listen(":3000")
}

// getMostCommonWordsHandler returns the most common words
func getMostCommonWordsHandler(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusOK).SendString(TopWords.ToString())
	return nil
}

// getStrHandler gets a string and processes it
func getStrHandler(ctx *fiber.Ctx) error {
	// Step 1: Gets the parameter from the raw body
	body := ctx.Body()
	words := strings.Split(string(body), ",")
	// Step 2: Creates a histogram for each word
	hist := map[string]int{}
	for _, word := range words {
		if _, ok := hist[word]; ok {
			hist[word] += 1
		} else {
			hist[word] = 1
		}
	}
	// Step 3: Inserts/Updates all words to the AllWords global map
	insertToGlobalMap(hist)
	// Step 4: Inserts/Updates the most common words in the TopWords struct
	insertToTopN(hist)
	ctx.Status(fiber.StatusOK).SendString("All words uploaded!\n")
	return nil
}

// insertToGlobalMap inserts/updates all words to the AllWords global map
func insertToGlobalMap(hist map[string]int) {
	// Step 1: For each word
	for word, _ := range hist {
		// Step 2: Checks if the word exists
		if val, ok := AllWords.Load(word); ok {
			// Updates the actual value of the word
			count := val.(int) + hist[word]
			AllWords.Store(word, count)
			hist[word] = count
		} else {
			// Adds a new key to the global map
			AllWords.Store(word, hist[word])
		}
	}
}

// insertToTopN inserts/updates the most common words in the TopWords struct
func insertToTopN(hist map[string]int) {
	size := TopWords.Len() - 1
	// Step 1: For each word
	for word, _ := range hist {
		val := hist[word]
		// Step 2: Checks if the word exists in the slice
		if idx := TopWords.Search(word); idx <= size {
			// Updates the value and sorts the slice
			TopWords.UpdateVal(val, idx)
			TopWords.Sort()
			continue
		}
		// Step 3: Checks if the value is less than or equal to the smallest value in the slice
		if val <= TopWords.GetVal(size) {
			// Skips to the next word
			continue
		}
		// Step 4: Updates the word in the last entry in the slice and sorts
		TopWords.UpdateKV(word, val, size)
		TopWords.Sort()
	}
}
