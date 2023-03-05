/*
	1 와일드 카드란 패턴 매칭에 사용되는 문자 또는 기호이다.
	  그러면 패턴 매칭이란?
	  패턴 매칭은 주어진 패턴과 일차하는 문자열을 찾는 과정이다.

	2 실행 인수란 linux에서 cd 명령어를 사용할 때, [cd 디렉토리 이름] 과 같은 형식을 사용하는데, 여기서 디렉토리 이름을 실행 인수라고 한다.

	3 Go에서 어떻게 파일을 열고, 읽고, 파일 목록을 가져올까?
	  파일을 열려면 os 패키지의 open() 함수를 이용해서 파일을 열어 파일 핸들을 가져와야 합니다. (파일 핸들은 컴퓨터에서 파일을 식별하는데 사용되는 일종의 식별자이다. 파일 핸들을 사용하면 열고, 읽고, 쓰고, 닫고가 가능하다)

*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//실행 인수 갯수 검사
	if len(os.Args) < 3 {
		fmt.Printf("usage : %s [WORD] [FILE_PATH]\n", os.Args[0])
		return
	}

	word := os.Args[1]
	// 와일드 카드 입력 가능
	files := os.Args[2:]
	fmt.Println("찾으려는 단어 : ", word, "\n파일 경로 : ", files)
	// 함수 실행
	PrintAllFiles(files)

}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path) // 모든 파일 이름의 목록을 반환하는 함수입니다.
}

func PrintAllFiles(files []string) {
	// fmt.Println(">>> files : ", files)
	// /bin /boot /dev /etc /home /lib /lib64 /lost+found /media /mnt /opt /proc /root /run /sbin /srv /sys /tmp /usr /var
	for i, path /* filelist = /bin, /boot, /dev */ := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Printf("파일 경로가 잘못 되었습니다. [%e] [%s]", err, path)
			return
		}

		fmt.Printf("[%d] 찾으시려는 문자열", i)
		for j, name := range filelist {
			fmt.Printf("[j : %d] %s\n", j, name)
		}

	}
}
