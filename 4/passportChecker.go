package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/james-wallis/adventofcode/utils"
)

const input = "./input.txt"

// Passport contains information about a person
type Passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

// ConvertLineToPassport take a line and converts it to the Passport struct
func ConvertLineToPassport(line string) Passport {
	// The order here has to match the order what the values are added to the Passport at the end of the statement
	passportKeys := [8]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	passportValues := [8]string{}

	for i, key := range passportKeys {
		re := regexp.MustCompile(key + ":(.*?)(\\s|\\n|\\z)")
		match := re.FindStringSubmatch(line)
		if len(match) >= 2 {
			passportValues[i] = match[1]
		}
	}
	passport := Passport{
		birthYear:      passportValues[0],
		issueYear:      passportValues[1],
		expirationYear: passportValues[2],
		height:         passportValues[3],
		hairColor:      passportValues[4],
		eyeColor:       passportValues[5],
		passportID:     passportValues[6],
		countryID:      passportValues[7],
	}
	return passport
}

// IsPassportValidPart1 returns true if a passport is valid (all fields have values apart from countryID)
func IsPassportValidPart1(passport Passport) bool {
	return !(passport.birthYear == "" ||
		passport.issueYear == "" ||
		passport.expirationYear == "" ||
		passport.height == "" ||
		passport.hairColor == "" ||
		passport.eyeColor == "" ||
		passport.passportID == "")
}

// IsPassportValidPart2 returns true if a passport matches the criteria for part 2:
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
func IsPassportValidPart2(passport Passport) bool {
	// Use part 1 as an initial check
	if !IsPassportValidPart1(passport) {
		return false
	}

	// Handle years first
	// Ensure they are the correct length
	if len(passport.birthYear) != 4 || len(passport.issueYear) != 4 || len(passport.expirationYear) != 4 {
		return false
	}

	// Ensure they are in the correct range
	birthYear, birthYearErr := strconv.Atoi(passport.birthYear)
	issueYear, issueYearErr := strconv.Atoi(passport.issueYear)
	expirationYear, expirationYearErr := strconv.Atoi(passport.expirationYear)
	if (birthYearErr != nil || issueYearErr != nil || expirationYearErr != nil) ||
		(birthYear < 1920 || birthYear > 2002) ||
		(issueYear < 2010 || issueYear > 2020) ||
		(expirationYear < 2020 || expirationYear > 2030) {
		return false
	}

	// Handle height
	heightAsInt, heightAsIntErr := strconv.Atoi(passport.height[:len(passport.height)-2])
	endsWithCm := strings.HasSuffix(passport.height, "cm")
	endsWithIn := strings.HasSuffix(passport.height, "in")
	// Validate CM, then IN
	if heightAsIntErr == nil && (endsWithCm || endsWithIn) {
		if (endsWithCm && (heightAsInt < 150 || heightAsInt > 193)) ||
			(endsWithIn && (heightAsInt < 59 || heightAsInt > 76)) {
			return false
		}
	} else {
		return false
	}

	// Handle hair colour (a # followed by exactly six characters 0-9 or a-f)
	hairColor := passport.hairColor
	hairColorRe := regexp.MustCompile("^#[a-fA-F0-9]*$")
	hairColorMatch := hairColorRe.FindString(hairColor)
	if hairColorMatch == "" || !strings.HasPrefix(hairColor, "#") || len(hairColor) != 7 {
		return false
	}

	// Handle eye colour (amb blu brn gry grn hzl oth)
	eyeColor := passport.eyeColor
	if eyeColor != "amb" && eyeColor != "blu" && eyeColor != "brn" &&
		eyeColor != "gry" && eyeColor != "grn" && eyeColor != "hzl" && eyeColor != "oth" {
		return false
	}

	// Handle Passport ID
	passportID := passport.passportID
	passportIDRe := regexp.MustCompile("^[0-9]*$")
	passportIDMatch := passportIDRe.FindString(passportID)
	if passportIDMatch == "" || len(passportID) != 9 {
		return false
	}

	return true
}

// CalculateValidPassports takes an array of strings representing passports (lines of a file) and calculates how many are valid:
//	* The next passport is denoted than an extra newline
// 	* Converts a string passport into the Passport struct using the key given in the challenge:
// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)
func CalculateValidPassports(lines []string) (part1Valid int, part2Valid int) {
	part1Valid, part2Valid = 0, 0
	var sanitizedLines []string
	sanitizedIndex := 0
	for i, line := range lines {
		if line == "" {
			// If empty line increment sanitized index (next line will be a new passport)
			sanitizedIndex++
		} else if (len(sanitizedLines)) == sanitizedIndex {
			// the current index needs to be created
			sanitizedLines = append(sanitizedLines, line)
		} else {
			// the current index already exists, append to it
			sanitizedLines[sanitizedIndex] += " " + line
		}

		if i+1 == len(lines) || lines[i+1] == "" {
			// If the next line is blank or it is the end of the slice,
			// convert the current sanitizedLine to a passport and append it to the passports slice
			passport := ConvertLineToPassport(sanitizedLines[sanitizedIndex])

			if IsPassportValidPart1(passport) {
				part1Valid++
			}

			if IsPassportValidPart2(passport) {
				part2Valid++
			}
		}
	}
	return
}

func main() {
	lines, _ := utils.ReadLines(input)

	IsPassportValidPart2(Passport{
		eyeColor:       "gry",
		passportID:     "860033327",
		expirationYear: "2020",
		hairColor:      "#fffffd",
		birthYear:      "1937",
		issueYear:      "2017",
		countryID:      "147",
		height:         "183cm",
	})

	part1Valid, part2Valid := CalculateValidPassports(lines)
	fmt.Println("*** Part 1 ***")
	fmt.Printf("\nNumber of valid passports = %d\n\n", part1Valid)

	fmt.Println("*** Part 2 ***")
	fmt.Printf("\nNumber of valid passports = %d\n\n", part2Valid)
}
