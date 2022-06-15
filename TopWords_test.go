package main

import (
	"reflect"
	"sync"
	"testing"
)

var topTest = [TopNum]KV{{"a", 14}, {"z", 8}, {"d", 3}, {"y", 2}, {"f", 1}}

// TestNewTopN tests the NewTopN() function
func TestNewTopN(t *testing.T) {
	tests := []struct {
		name string
		val  int
		want *TopN
	}{
		{"Initializes the TopKV slice values to 0", 0, NewTopN(0)},
		{"Initializes the TopKV slice values to 1", 1, NewTopN(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTopN(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTopN() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTopN_GetKV tests the GetKV() function
func TestTopN_GetKV(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
		want  [TopNum]KV
	}{
		{"Gets empty KV slice", sync.RWMutex{}, [TopNum]KV{}, [TopNum]KV{}},
		{"Gets KV slice successfully", sync.RWMutex{}, topTest, topTest},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			if got := t.GetKV(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetKV() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTopN_GetKey tests the GetKey() function
func TestTopN_GetKey(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
		i     int
		want  string
	}{
		{"Gets an empty key from New struct", sync.RWMutex{}, NewTopN(0).GetKV(), 0, ""},
		{"Gets key successfully", sync.RWMutex{}, topTest, 2, "d"},
		{"Out of range index - gets an empty key", sync.RWMutex{}, topTest, TopNum + 2, ""},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			if got := t.GetKey(tt.i); got != tt.want {
				t1.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTopN_GetVal tests the GetVal() function
func TestTopN_GetVal(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
		i     int
		want  int
	}{
		{"Gets an empty value from New struct", sync.RWMutex{}, NewTopN(0).GetKV(), 0, 0},
		{"Gets value successfully", sync.RWMutex{}, topTest, 2, 3},
		{"Out of range index - gets -1", sync.RWMutex{}, topTest, TopNum + 2, -1},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			if got := t.GetVal(tt.i); got != tt.want {
				t1.Errorf("GetVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTopN_Len tests the Len() function
func TestTopN_Len(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
		want  int
	}{
		{"Gets len for an empty KV slice", sync.RWMutex{}, [TopNum]KV{}, TopNum},
		{"Gets len successfully", sync.RWMutex{}, topTest, TopNum},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			if got := t.Len(); got != tt.want {
				t1.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTopN_Print tests the Print() function
func TestTopN_Print(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
	}{
		{"Prints an empty KV slice", sync.RWMutex{}, [TopNum]KV{}},
		{"Prints successfully", sync.RWMutex{}, topTest},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			t.Print()
		})
	}
}

// TestTopN_Search tests the Search() function
func TestTopN_Search(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
		key   string
		want  int
	}{
		{"Len of empty KV slice", sync.RWMutex{}, [TopNum]KV{}, "a", TopNum},
		{"Success get KV slice", sync.RWMutex{}, topTest, "a", 0},
		{"Success get KV slice", sync.RWMutex{}, topTest, "d", 2},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			if got := t.Search(tt.key); got != tt.want {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTopN_Sort tests the Sort() function
func TestTopN_Sort(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
	}{
		{"Sorts an empty KV slice", sync.RWMutex{}, [TopNum]KV{}},
		{
			name:  "Gets KV slice successfully",
			Lock:  sync.RWMutex{},
			TopKV: [TopNum]KV{{"a", 3}, {"z", 8}, {"d", 2}, {"y", 1}, {"f", 19}},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			t.Sort()
			t.Print()
		})
	}
}

// TestTopN_UpdateKV tests the UpdateKV() function
func TestTopN_UpdateKV(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
		key   string
		val   int
		i     int
	}{
		{"Updates an empty TopKV slice", sync.RWMutex{}, [TopNum]KV{}, "a", 3, 0},
		{"Out of range index - not updated", sync.RWMutex{}, [TopNum]KV{}, "a", 3, TopNum + 2},
		{"Updates the TopKV slice successfully", sync.RWMutex{}, topTest, "r", 5, 0},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			t.UpdateKV(tt.key, tt.val, tt.i)
			t.Print()
		})
	}
}

// TestTopN_UpdateVal tests the UpdateVal() function
func TestTopN_UpdateVal(t1 *testing.T) {
	tests := []struct {
		name  string
		Lock  sync.RWMutex
		TopKV [TopNum]KV
		val   int
		i     int
	}{
		{"Updates the value in an empty TopKV slice", sync.RWMutex{}, [TopNum]KV{}, 3, 0},
		{"Updates the value successfully", sync.RWMutex{}, topTest, 5, 0},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TopN{
				Lock:  tt.Lock,
				TopKV: tt.TopKV,
			}
			t.UpdateVal(tt.val, tt.i)
		})
	}
}
