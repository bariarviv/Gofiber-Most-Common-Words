package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"sync"
	"testing"
)

// Test_getMostCommonWordsHandler tests the getMostCommonWordsHandler() function
func Test_getMostCommonWordsHandler(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	tests := []struct {
		name    string
		ctx     *fiber.Ctx
		input   string
		wantErr bool
	}{
		{"Gets a string and processes it successfully", ctx, "a,a,a,r,r,t,y,u,i,d,d,d", false},
		{"Gets a string and processes it successfully", ctx, "z,z,z,z,z,z,z,z,v,b,r", false},
		{"Gets a string and processes it successfully", ctx, "w,r,u,o,o,o,o,o,o,o,o,o,o,o,o,o,o", false},
		{"Gets an empty string successfully", ctx, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctx.Request().SetBody([]byte(tt.input))
			if err := getStrHandler(tt.ctx); err == nil {
				if err = getMostCommonWordsHandler(tt.ctx); (err != nil) != tt.wantErr {
					t.Errorf("getMostCommonWordsHandler() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

// Test_insertToGlobalMap tests the insertToGlobalMap() function
func Test_insertToGlobalMap(t *testing.T) {
	// Initializes the AllWords global map
	initGlobalMapTest()
	tests := []struct {
		name string
		word string
		val  int
	}{
		{"Inserts an empty word - not updated", "", 10},
		{"Inserts a new word successfully", "t", 10},
		{"Inserts a new word successfully", "k", 4},
		{"Updates an existing word successfully", "a", 3},
		{"Updates an existing word successfully", "z", 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			max, min := TopWords.GetMinMax()
			t.Run(tt.name, func(t *testing.T) {
				insertToGlobalMap(tt.word, tt.val, max, min)
			})
			printSyncMap(AllWords)
		})
	}
}

// Test_insertToTopN tests the insertToTopN() function
func Test_insertToTopN(t *testing.T) {
	// topTest = [TopNum]KV{{"a", 5}, {"z", 1}, {"d", 3}, {"y", 3}, {"f", 1}}
	TopWords.TopKV = topTest
	TopWords.Sort()
	tests := []struct {
		name string
		word string
		val  int
	}{
		{"Inserts an empty word - not updated", "", 5},
		{"Does not update", TopWords.GetKey(0), TopWords.GetVal(0)},
		{"Updates an existing word successfully", TopWords.GetKey(2), TopWords.GetVal(2) + 1},
		{"Inserts a new word successfully", "t", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertToTopN(tt.word, tt.val)
			TopWords.Print()
		})
	}
}

// Test_getStrHandler tests the getStrHandler() function
func Test_getStrHandler(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	tests := []struct {
		name    string
		ctx     *fiber.Ctx
		input   string
		wantErr bool
	}{
		{"Gets a string and processes it successfully", ctx, "a,a,a,r,r,t,y,u,i,d,d,d", false},
		{"Gets a string and processes it successfully", ctx, "z,z,z,z,z,z,z,z,v,b,r", false},
		{"Gets a string and processes it successfully", ctx, "w,r,u,o,o,o,o,o,o,o,o,o,o,o,o,o,o", false},
		{"Gets an empty string successfully", ctx, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctx.Request().SetBody([]byte(tt.input))
			if err := getStrHandler(tt.ctx); (err != nil) != tt.wantErr {
				t.Errorf("getStrHandler() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Println(TopWords.ToString())
			}
		})
	}
}

// initGlobalMapTest initializes the AllWords global map for testing
func initGlobalMapTest() {
	AllWords.Store("a", 1)
	AllWords.Store("v", 2)
	AllWords.Store("z", 3)
	AllWords.Store("b", 1)
	AllWords.Store("u", 5)
}

// printSyncMap prints the sync.map and length for testing
func printSyncMap(m sync.Map) {
	i := 0
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("\t[%d] key: %v, value: %v\n", i, key, value)
		i++
		return true
	})
	fmt.Printf("\tLength: %d\n", i)
}
