package main

import (
	"bufio"
	"errors"
	"fmt"
	"go-reloaded/functions"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(errors.New("please enter all arguments"))
		return
	}
	openFile := os.Args[1]
	file, err := os.OpenFile(openFile, os.O_RDONLY, os.FileMode(0600))
	functions.CheckErrors("choose correct file", err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	s := []string{}
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	functions.CheckErrors("scan error, try again", scanner.Err())

	str := strings.Join(s, "") // string from slice of strings

	mods := regexp.MustCompile(`\,\s`)
	str = mods.ReplaceAllString(str, ",") // removing spaces before comma

	resS := strings.Split(str, " ") // string to slice of strings
	modifiedRes := functions.Corrector(resS)
	strFinal := strings.Join(modifiedRes, " ") // our final string before last corrections
	res := functions.RegexCorrections(strFinal)
	// Writing to file
	writeFile := os.Args[2]
	writer, _ := os.Create(writeFile)
	defer writer.Close()
	writer.WriteString(res)
}
