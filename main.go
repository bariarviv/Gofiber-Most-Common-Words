package main

import (
	"github.com/gofiber/fiber/v2"
	"math"
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
	max, min := TopWords.GetMinMax()
	// Gets the parameter from the raw body
	body := ctx.Body()
	words := strings.Split(string(body), ",")
	// Creates a histogram for each word
	hist := map[string]int{}
	for _, word := range words {
		if word == "" {
			continue
		}
		if _, ok := hist[word]; ok {
			hist[word] += 1
			if hist[word] > max {
				max = hist[word]
			}
		} else {
			hist[word] = 1
		}
	}
	// Inserts/Updates all words to the AllWords global map and inserts/Updates the most common words in the TopWords
	for word, val := range hist {
		go insertToGlobalMap(word, val, max, min)
	}
	ctx.Status(fiber.StatusOK).SendString("All words uploaded!\n")
	return nil
}

// insertToGlobalMap inserts/updates all words to the AllWords global map
func insertToGlobalMap(word string, val, max, min int) {
	if word == "" {
		return
	}
	if v, ok := AllWords.Load(word); ok {
		val += v.(int)
	}
	// Adds/Updates the value of the word
	AllWords.Store(word, val)

	histFreq := 1
	if val >= max {
		histFreq = TopNum
	} else if val > min {
		histFreq = int(math.Ceil((float64(val-min)/float64(max-min))*4 + 1))
	}
	insertToTopN(word, histFreq)
}

// insertToTopN inserts/updates the most common words in the TopWords struct
func insertToTopN(word string, val int) {
	if word == "" {
		return
	}
	size := TopWords.Len() - 1
	// Checks if the word exists in the slice
	if idx, v := TopWords.Search(word); idx <= size {
		if val != v {
			// Updates the value and sorts the slice
			TopWords.UpdateVal(val, idx)
			TopWords.Sort()
		}
		return
	}
	// Checks if the value is less than or equal to the smallest value in the slice
	if val <= TopWords.GetVal(size) {
		// Skips to the next word
		return
	}
	// Updates the word in the last entry in the slice and sorts
	TopWords.UpdateKV(word, val, size)
	TopWords.Sort()
}
