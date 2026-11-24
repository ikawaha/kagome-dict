package dict

import (
	"slices"
	"testing"
)

func TestIndexTable_DupCache(t *testing.T) {
	keys := []string{"a", "a", "a", "ab", "b", "b"}
	idx, err := BuildIndexTable(keys)
	if err != nil {
		t.Fatal(err)
	}
	if idx.dup == nil {
		t.Fatal("expected the dense dup table to be built")
	}
	check := func(t *testing.T, idx IndexTable) {
		t.Helper()
		for _, tt := range []struct {
			in   string
			want []int
		}{
			{"a", []int{0, 1, 2}},
			{"ab", []int{3}},
			{"b", []int{4, 5}},
			{"c", nil},
		} {
			if got := idx.Search(tt.in); !slices.Equal(got, tt.want) {
				t.Errorf("Search(%q) = %v, want %v", tt.in, got, tt.want)
			}
		}
		var got [][2]int
		idx.CommonPrefixSearchCallback("abc", func(id, l int) {
			got = append(got, [2]int{id, l})
		})
		if want := [][2]int{{0, 1}, {1, 1}, {2, 1}, {3, 2}}; !slices.Equal(got, want) {
			t.Errorf("CommonPrefixSearchCallback = %v, want %v", got, want)
		}
	}
	t.Run("dense", func(t *testing.T) { check(t, idx) })

	t.Run("map fallback", func(t *testing.T) {
		idx := idx
		idx.Dup = map[int32]int32{0: 2, 4: 1, 1000: 300} // 300 does not fit in uint8
		idx.RebuildDupCache()
		if idx.dup != nil {
			t.Fatal("expected no dense dup table")
		}
		check(t, idx)
	})

	t.Run("rebuild after change", func(t *testing.T) {
		idx := idx
		idx.Dup = map[int32]int32{0: 1}
		idx.RebuildDupCache()
		if got, want := idx.Search("a"), []int{0, 1}; !slices.Equal(got, want) {
			t.Errorf("Search(a) = %v, want %v", got, want)
		}
		if got, want := idx.Search("b"), []int{4}; !slices.Equal(got, want) {
			t.Errorf("Search(b) = %v, want %v", got, want)
		}
	})
}
