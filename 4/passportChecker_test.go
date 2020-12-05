package main

import (
	"fmt"
	"testing"

	"github.com/james-wallis/adventofcode/utils"
)

const sampleFile = "./sample.txt"

func TestConvertPassportToStruct(t *testing.T) {
	t.Run("converts a valid string into a Passport", func(t *testing.T) {
		line := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"

		want := Passport{
			eyeColor:       "gry",
			passportID:     "860033327",
			expirationYear: "2020",
			hairColor:      "#fffffd",
			birthYear:      "1937",
			issueYear:      "2017",
			countryID:      "147",
			height:         "183cm",
		}
		got := ConvertLineToPassport(line)

		if got != want {
			t.Errorf("got %+v want %+v given, %s and", got, want, line)
		}
	})

	t.Run("converts a valid string into a Passport  (missing countryID)", func(t *testing.T) {
		line := "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm"

		want := Passport{
			eyeColor:       "brn",
			passportID:     "760753108",
			expirationYear: "2024",
			hairColor:      "#ae17e1",
			birthYear:      "1931",
			issueYear:      "2013",
			height:         "179cm",
			countryID:      "",
		}
		got := ConvertLineToPassport(line)

		if got != want {
			t.Errorf("got %+v want %+v given, %s and", got, want, line)
		}
	})

	t.Run("converts an invalid string into a Passport (missing passportID, birthYear and Height)", func(t *testing.T) {
		line := "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn cid:147"

		want := Passport{
			eyeColor:       "brn",
			passportID:     "",
			expirationYear: "2024",
			hairColor:      "#ae17e1",
			birthYear:      "",
			issueYear:      "2013",
			height:         "",
			countryID:      "147",
		}
		got := ConvertLineToPassport(line)

		if got != want {
			t.Errorf("got %+v want %+v given, %s and", got, want, line)
		}
	})
}

func TestIsPassportValidPart1(t *testing.T) {
	t.Run("returns true as all fields have values", func(t *testing.T) {
		want := true
		got := IsPassportValidPart1(Passport{
			eyeColor:       "gry",
			passportID:     "860033327",
			expirationYear: "2020",
			hairColor:      "#fffffd",
			birthYear:      "1937",
			issueYear:      "2017",
			countryID:      "147",
			height:         "183cm",
		})

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("returns true as all fields have values apart from countryID (not required)", func(t *testing.T) {
		want := true
		got := IsPassportValidPart1(Passport{
			eyeColor:       "gry",
			passportID:     "860033327",
			expirationYear: "2020",
			hairColor:      "#fffffd",
			birthYear:      "1937",
			issueYear:      "2017",
			countryID:      "",
			height:         "183cm",
		})

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("returns false as passportID is blank (required)", func(t *testing.T) {
		want := false
		got := IsPassportValidPart1(Passport{
			eyeColor:       "gry",
			passportID:     "",
			expirationYear: "2020",
			hairColor:      "#fffffd",
			birthYear:      "1937",
			issueYear:      "2017",
			countryID:      "147",
			height:         "183cm",
		})

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestIsPassportValidPart2(t *testing.T) {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.
	var cases = []struct {
		title    string
		passport Passport
		want     bool
	}{
		{
			"returns true as as the passport is valid",
			Passport{
				eyeColor:       "gry",
				passportID:     "060033327",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			true,
		},
		{
			"returns false as passportID is blank (required field)",
			Passport{
				eyeColor:       "gry",
				passportID:     "",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as birthYear is less than 1920",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1919",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as expirationYear is greater than 2030",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2050",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as the height is greater than 193cm",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "200cm",
			},
			false,
		},
		{
			"returns false as the height is greater than 76inches",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "77in",
			},
			false,
		},
		{
			"returns false as the height is doesn't end with cm or in",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "77",
			},
			false,
		},
		{
			"returns false as as the hair colour is invalid (doesn't start with #)",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "1234567",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as as the hair colour is invalid (length isn't 7)",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#fffffffff",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as as the hair colour is invalid (characters are not letters and numbers after the #)",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#!?!?!?",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as as the hair colour is invalid (has letters that are outside of the a-f boundary)",
			Passport{
				eyeColor:       "gry",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#GOHABC",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as as the eye colour is invalid",
			Passport{
				eyeColor:       "org",
				passportID:     "860033327",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as as the passportID is invalid (too many numbers)",
			Passport{
				eyeColor:       "gry",
				passportID:     "8600333273",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
		{
			"returns false as as the passportID is invalid (contains letters)",
			Passport{
				eyeColor:       "gry",
				passportID:     "123456SOM",
				expirationYear: "2020",
				hairColor:      "#fffffd",
				birthYear:      "1937",
				issueYear:      "2017",
				countryID:      "147",
				height:         "183cm",
			},
			false,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s", test.title), func(t *testing.T) {
			got := IsPassportValidPart2(test.passport)
			if got != test.want {
				t.Errorf("got %t want %t", got, test.want)
			}
		})
	}

}

func TestCalculateValidPassports(t *testing.T) {
	t.Run("converts an array of strings separated by a blank string into Passports (test first part)", func(t *testing.T) {
		lines := []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", // valid - has all fields
			"",
			"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", // valid - missing only countryID
			"",
			"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931", // invalid - missing height
		}

		want := 2
		got, _ := CalculateValidPassports(lines)

		if got != want {
			t.Errorf("got %d want %d given, %s", got, want, lines)
		}
	})

	t.Run("tests the validPart2.txt file", func(t *testing.T) {
		lines, _ := utils.ReadLines("./validPart2.txt")

		want := 4
		_, got := CalculateValidPassports(lines)

		if got != want {
			t.Errorf("got %d want %d given, %s", got, want, lines)
		}
	})

	t.Run("tests the invalidPart2.txt file", func(t *testing.T) {
		lines, _ := utils.ReadLines("./invalidPart2.txt")

		want := 0
		_, got := CalculateValidPassports(lines)

		if got != want {
			t.Errorf("got %d want %d given, %s", got, want, lines)
		}
	})
}
