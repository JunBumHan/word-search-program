// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// type FILE_LINE_TEXT struct {
// 	fileLineNumber int
// 	fileSentence   string
// }

// type FILE struct {
// 	fileName string
// 	fileText []FILE_LINE_TEXT
// }

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Printf("Usage : %s [word] [directory]\n", os.Args[0])
// 		return
// 	}

// 	// 단어 추출
// 	word := os.Args[1]
// 	path := os.Args[2:]
// 	res := []FILE{}

// 	for _, dir := range path {
// 		res = append(res, findWord(word, dir)...)
// 	}

// 	for _, t := range res {
// 		fmt.Println(t.fileName)
// 		for _, b := range t.fileText {
// 			fmt.Println("\t", b.fileLineNumber, " : ", "\t", b.fileSentence)
// 		}
// 	}
// }

// func findWord(word, dir string) []FILE {

// 	dirRes := []FILE{}
// 	list, err := filepath.Glob(dir)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, fileName := range list {
// 		dirRes = append(dirRes, FINDinFILEtoWORLD(word, fileName))
// 	}

// 	return dirRes

// }

// func FINDinFILEtoWORLD(word, fileName string) FILE {
// 	fileRes := FILE{fileName: fileName, fileText: []FILE_LINE_TEXT{}}
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	lineNo := 1
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.Contains(line, word) {
// 			fileRes.fileText = append(fileRes.fileText, FILE_LINE_TEXT{lineNo, line})
// 		}
// 		lineNo++
// 	}

// 	return fileRes

// }
