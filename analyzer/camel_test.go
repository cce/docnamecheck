package analyzer

import "testing"

func TestHasCamelChunkReplacement(t *testing.T) {
	tests := []struct {
		doc, sym string
		max      int
		want     bool
	}{
		{"processCIDRs", "validateCIDRs", 2, true},
		{"processCIDRs", "validateCIDRs", 1, true},
		{"processCIDRs", "validateFooCIDRs", 1, false},
		{"oneWord", "oneWord", 1, false},
	}
	for _, tt := range tests {
		if got := hasCamelChunkReplacement(tt.doc, tt.sym, tt.max); got != tt.want {
			t.Errorf("hasCamelChunkReplacement(%q,%q,%d)=%v, want %v", tt.doc, tt.sym, tt.max, got, tt.want)
		}
	}
}

func TestHasCamelChunkInsertionOrRemoval(t *testing.T) {
	tests := []struct {
		doc, sym string
		max      int
		want     bool
	}{
		{"handleVolume", "handleEphemeralVolume", 2, true},
		{"syncHandler", "sync", 2, true},
		{"syncHandler", "sync", 0, false},
		{"fooBar", "fooBar", 2, false},
	}
	for _, tt := range tests {
		if got := hasCamelChunkInsertionOrRemoval(tt.doc, tt.sym, tt.max); got != tt.want {
			t.Errorf("hasCamelChunkInsertionOrRemoval(%q,%q,%d)=%v, want %v", tt.doc, tt.sym, tt.max, got, tt.want)
		}
	}
}

func TestIsCamelSwapVariant(t *testing.T) {
	tests := []struct {
		doc, sym string
		want     bool
	}{
		{"getPodsReady", "getReadyPods", true},
		{"getPodIPs", "getPodIPs", false},
	}
	for _, tt := range tests {
		if got := isCamelSwapVariant(tt.doc, tt.sym); got != tt.want {
			t.Errorf("isCamelSwapVariant(%q,%q)=%v, want %v", tt.doc, tt.sym, got, tt.want)
		}
	}
}
