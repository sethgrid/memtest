package main

import "testing"

func TestGetOptions(t *testing.T) {
	opts := getOptions(3)

	if len(opts) != 3 {
		t.Errorf("got %d options, want 3", len(opts))
	}

	if !allUnique(opts) {
		t.Errorf("not all options are unique")
	}
}

func TestShuffle(t *testing.T) {
	opts := getOptions(6)
	shuffled := shuffle(opts)

	allTheSame := true

	for i, _ := range opts {
		if opts[i] != shuffled[i] {
			allTheSame = false
		}
	}

	if allTheSame {
		t.Errorf("shuffling does not change order")
	}
}

func TestAddOption(t *testing.T) {
	opts := getOptions(1)
	newOpts, _ := addOption(opts)

	if len(newOpts) != len(opts)+1 {
		t.Fatalf("got %d new options, want %d", len(newOpts), len(opts)+1)
	}
}

func allUnique(rs []rune) bool {
	set := make(map[rune]bool)
	for _, r := range rs {
		set[r] = true
	}
	return len(set) == len(rs)
}
