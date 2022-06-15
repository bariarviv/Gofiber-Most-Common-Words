package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"reflect"
	"sync"
	"testing"
)

// Test_insertToGlobalMap tests the insertToGlobalMap() function
func Test_insertToGlobalMap(t *testing.T) {
	// Initializes the AllWords global map
	initGlobalMapTest()
	tests := []struct {
		name string
		hist map[string]int
		want map[string]int
	}{
		{"Inserts an empty map", map[string]int{}, map[string]int{}},
		{
			name: "Updates an existing word successfully",
			hist: map[string]int{"a": 2, "d": 9},
			want: map[string]int{"a": 3, "d": 9},
		},
		{
			name: "Updates an existing word successfully",
			hist: map[string]int{"g": 2, "z": 13, "w": 1},
			want: map[string]int{"g": 2, "z": 16, "w": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if insertToGlobalMap(tt.hist); !reflect.DeepEqual(tt.hist, tt.want) {
				t.Errorf("insertToGlobalMap() = %v, want %v", tt.hist, tt.want)
			}
			printSyncMap(AllWords)
		})
	}
}

// Test_insertToTopN tests the insertToTopN() function
func Test_insertToTopN(t *testing.T) {
	// topTest = [TopNum]KV{{"a", 14}, {"z", 8}, {"d", 3}, {"y", 2}, {"f", 1}}
	TopWords.TopKV = topTest
	tests := []struct {
		name string
		hist map[string]int
	}{
		{"Inserts an empty map", map[string]int{}},
		{"Does not update", map[string]int{"c": 1, "e": 1, "m": 1}},
		{"Updates an existing word successfully", map[string]int{"b": 1, "d": 5}},
		{"Inserts a new word successfully", map[string]int{"r": 1, "q": 15, "x": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertToTopN(tt.hist)
			TopWords.Print()
		})
	}
}

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
				t.Log(TopWords.ToString())
			}
		})
	}
}

// initGlobalMapTest initializes the AllWords global map for testing
func initGlobalMapTest() {
	AllWords.Store("a", 1)
	AllWords.Store("v", 16)
	AllWords.Store("z", 3)
	AllWords.Store("b", 7)
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
