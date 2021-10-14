package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	s1 string //test server https address
	s2 string
)

func main() {

	// flagset 만들어주기
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// 처음에는 StringVal 로 하나 받는 것만 해준다. 추후에는 struct 로 만들어 준다.
	// 디폴트 값은 https://daum.net 이다.
	fs.StringVar(&s1, "u", "https://daum.net", "https address")

	// 사용자 입력값을 파싱한다. 즉, 입력된 파라미터를 s1 에 집어 넣는다.
	fs.Parse(os.Args[1:])

	// 여기에 문제가 있다. 찾을 수 있을까?
	if len(os.Args) < 2 {
		// 에러 보고 및 exit
		// 에러 보고 코드 들어가야 함.
		os.Exit(1)
	}

	// 여기부터 사용자의 입력 파라미터를 사용하는 코드가 들어간다. 실제적인 코딩 부분
	fmt.Println("사용자 입력 파라미터", s1)
}

// https://pkg.go.dev/flag#NewFlagSet
// https://golang.org/src/flag/flag.go
