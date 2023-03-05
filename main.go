// // /*
// // 	1 와일드 카드란 패턴 매칭에 사용되는 문자 또는 기호이다.
// // 	  그러면 패턴 매칭이란?
// // 	  패턴 매칭은 주어진 패턴과 일차하는 문자열을 찾는 과정이다.

// // 	2 실행 인수란 linux에서 cd 명령어를 사용할 때, [cd 디렉토리 이름] 과 같은 형식을 사용하는데, 여기서 디렉토리 이름을 실행 인수라고 한다.

// // 	3 Go에서 어떻게 파일을 열고, 읽고, 파일 목록을 가져올까?
// // 	  파일을 열려면 os 패키지의 open() 함수를 이용해서 파일을 열어 파일 핸들을 가져와야 합니다. (파일 핸들은 컴퓨터에서 파일을 식별하는데 사용되는 일종의 식별자이다. 파일 핸들을 사용하면 열고, 읽고, 쓰고, 닫고가 가능하다)

// // */

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct {
	lineNo int
	line   string
}

// 파일 내 라인 정보
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	// 실행 인수 검사
	if len(os.Args) < 3 { // 만약 실행 인수가 3보다 작다면
		fmt.Printf("Usage : %s [WORD] [FILE_PATH]", os.Args[0]) // 올바른 사용법 출력
		return                                                  // return
	}

	// 실행 인수 추출
	word := os.Args[1]        // word는 찾을 단어 저장
	files := os.Args[2:]      // 파일 경로 저장
	findInfos := []FindInfo{} // FindInfo 슬라이스 생성

	for _, path := range files /* /etc, /boot, /dev, /lib  */ {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...) // findInfos 슬라이스에 FindWordInAllFiles 반환 구조체 요소값 추가
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------")
		fmt.Println()
	}
}

func GetFileList(path string) ([]string /*문자열 슬라이스*/, error) {
	return filepath.Glob(path)
}

// 모든 파일에서 단어를 찾는 함수 입니다.
func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{} //FindInfo 구조체 슬라이스 생성

	filelist, err := GetFileList(path) // 파일 리스트 가져오기
	if err != nil {                    // 파일 리스트가 오료나면
		fmt.Println("파일 경로가 잘못되었습니다. err :", err, "path:", path)
		return findInfos
	}

	for _, filename := range filelist { // 파일 이름 하나 하나를 findInfos 슬라이스에 추가
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}

	return findInfos

}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file) // scanner 생성 (file)
	for scanner.Scan() {              // Scan이 가능하냐?
		line := scanner.Text()            // o -> 한줄 읽어와
		if strings.Contains(line, word) { // 찾은 단어가 있냐?
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line}) // 있으면 findInfo에 단어 추가

		}
		lineNo++
	}
	return findInfo
}
