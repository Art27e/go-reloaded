package functions

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CheckErrors(text string, err error) {
	if err != nil {
		fmt.Println(errors.New(text))
	}
}

func Corrector(text []string) []string {
	for i := 0; i < len(text); i++ {
		tags := regexp.MustCompile(`\(low\,\d+\)|\(up\,\d+\)|\(cap\,\d+\)`)
		check := tags.MatchString(text[i])
		if check {
			switch {
			case strings.HasPrefix(text[i], "(low"):
				for range text[i] {
					testing := regexp.MustCompile(`\d`)
					numTest := testing.FindAllString(string(text[i]), -1)
					numStr := strings.Join(numTest, "")
					num, _ := strconv.Atoi(numStr)
					if i-num >= 0 {
						for x := 1; x <= num; x++ {
							text[i-x] = strings.ToLower(text[i-x])
						}
					}
					if i-num < 0 {
						for x := 1; x <= i; x++ {
							text[i-x] = strings.ToLower(text[i-x])
						}
					}
				}
			case strings.HasPrefix(text[i], "(up"):
				for range text[i] {
					testing := regexp.MustCompile(`\d`)
					numTest := testing.FindAllString(string(text[i]), -1)
					numStr := strings.Join(numTest, "")
					num, _ := strconv.Atoi(numStr)
					if i-num >= 0 {
						for x := 1; x <= num; x++ {
							text[i-x] = strings.ToUpper(text[i-x])
						}
					}
					if i-num < 0 {
						for x := 1; x <= i; x++ {
							text[i-x] = strings.ToUpper(text[i-x])
						}
					}
				}
			case strings.HasPrefix(text[i], "(cap"):
				for range text[i] {
					testing := regexp.MustCompile(`\d`)
					numTest := testing.FindAllString(string(text[i]), -1)
					numStr := strings.Join(numTest, "")
					num, _ := strconv.Atoi(numStr)
					if i-num >= 0 {
						for x := 1; x <= num; x++ {
							text[i-x] = strings.Title(text[i-x])
						}
					}
					if i-num < 0 {
						for x := 1; x <= i; x++ {
							text[i-x] = strings.Title(text[i-x])
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(text); i++ {
		tags2 := regexp.MustCompile(`\(low\)|\(up\)|\(cap\)|\(hex\)|\(bin\)`)
		check2 := tags2.MatchString(text[i])
		if check2 {
			re := regexp.MustCompile(`\(low\)`)
			checkForLow := re.MatchString(text[i])
			if checkForLow && len(text) > 1 {
				text[i-1] = strings.ToLower(text[i-1])
			}
			re2 := regexp.MustCompile(`\(cap\)`)
			checkForCap := re2.MatchString(text[i])
			if checkForCap && len(text) > 1 {
				text[i-1] = strings.Title(text[i-1])
			}
			re3 := regexp.MustCompile(`\(up\)`)
			checkForUp := re3.MatchString(text[i])
			if checkForUp && len(text) > 1 {
				text[i-1] = strings.ToUpper(text[i-1])
			}
			re4 := regexp.MustCompile(`\(hex\)`)
			checkForHex := re4.MatchString(text[i])
			if checkForHex && len(text) > 1 {
				decimalNum, _ := strconv.ParseInt(text[i-1], 16, 64)
				decimalStr := strconv.FormatInt(decimalNum, 10)
				text[i-1] = decimalStr
			}
			re5 := regexp.MustCompile(`\(bin\)`)
			checkForBin := re5.MatchString(text[i])
			if checkForBin && len(text) > 1 {
				binNum, _ := strconv.ParseInt(text[i-1], 2, 64)
				decimalStr2 := strconv.FormatInt(binNum, 10)
				text[i-1] = decimalStr2
			}
		}
	}
	return text
}

func RegexCorrections(s string) string {
	regex := regexp.MustCompile(`\(bin\)|\(hex\)|\(cap\)|\(up\)|\(low\)|\(up,\d+\)|\(low,\d+\)|\(cap,\d+\)`)
	s = regex.ReplaceAllString(s, "")
	regex2 := regexp.MustCompile(`\s\s`)
	s = regex2.ReplaceAllString(s, " ")
	regex3 := regexp.MustCompile(`(\s*)([\.,!\?:;]+)(\s*)`) // punctuations
	s = regex3.ReplaceAllString(s, `$2 `)
	regex4 := regexp.MustCompile(`(')(\s+)(.+?)(\s+)(')`) // quotes
	s = regex4.ReplaceAllString(s, `$1$3$5`)
	regex5 := regexp.MustCompile(`(^|\s)([aA])(\s[aeiouhAEIOUUH])`) // a and an
	s = regex5.ReplaceAllString(s, `$1${2}n$3`)
	return s
}
