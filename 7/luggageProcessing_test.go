package main

import (
	"fmt"
	"testing"
)

var input = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
	"dark teal bags contain 1 dark olive bag, 1 drab lavender bag, 2 mirrored purple bags, 1 pale teal bag.",
}

var part2Input = []string{
	"shiny gold bags contain 2 dark red bags.",
	"dark red bags contain 2 dark orange bags.",
	"dark orange bags contain 2 dark yellow bags.",
	"dark yellow bags contain 2 dark green bags.",
	"dark green bags contain 2 dark blue bags.",
	"dark blue bags contain 2 dark violet bags.",
	"dark violet bags contain no other bags.",
}

var part2InputSimple = []string{
	"shiny gold bags contain 2 dark red bags.",
	"dark red bags contain no other bags.",
}

func TestFindDirectHoldersOfBag(t *testing.T) {
	var cases = []struct {
		bag  string
		want int
	}{
		{

			"shiny gold",
			2,
		},
		{

			"bright white",
			2,
		},
		{

			"faded blue",
			3,
		},
		{

			"dotted black",
			2,
		},
		{
			"muted yellow",
			2,
		},
		{
			"dark olive",
			2,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.bag), func(t *testing.T) {
			got := FindDirectHoldersOfBag(test.bag, input)
			if len(got) != test.want {
				t.Errorf("got %d want %d", len(got), test.want)
			}
		})
	}
}

func TestGetNumEventualBagHolders(t *testing.T) {
	var cases = []struct {
		bag  string
		want int
	}{
		{

			"shiny gold",
			4,
		},
		{

			"faded blue",
			8,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.bag), func(t *testing.T) {
			got := GetNumEventualBagHolders(test.bag, input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}

func TestCountBagsInside(t *testing.T) {
	var cases = []struct {
		bag   string
		want  int64
		input []string
	}{
		{

			"shiny gold",
			126,
			part2Input,
		},
		{

			"shiny gold",
			2,
			part2InputSimple,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.bag), func(t *testing.T) {
			got := CountBagsInside(test.bag, test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
