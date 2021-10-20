package main

import (
	"flag"
	"fmt"
	"os"

	conf "github.com/seoyhaein/golang-study/config"
)

// 10/15
// 예전에 세팅해놓은 것이라 틀릴 수도 있음.
// 위에 import 를 하기 위해서는 또한 github 에 올릴때 충돌?(디렉토리가 이상해지는) 이 없도록 하기 위해서는 몇가지 디렉토리 계층을 가지도록 하는게 좋다.
// go 의 소스를 놓는 위치인 src 에 github.com/자신의github/프로젝트네임 으로 해주면 별 문제가 없다.
// 나는 ~/go/src/github.com/seoyhaein/golang-study 에서 작업을 하고 github 에 올린다.
// 참고 하나는 GOROOT 설정까지 해누는 것이고 다른 하나는 GOBIN 만 설정해주는 것. 요즘은 GOROOT 설정 않해주는 지 모르겠음.
// https://medium.com/chequer/goroot%EC%99%80-gopath-77f44cbaa1d8
// https://golang.org/doc/install

/*var (
	s1 string //test server https address
	s2 string
)*/

// 10/18
var Version = "0.0.0"

func main() {

	// flagset 만들어주기
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// 처음에는 StringVar 로 하나 받는 것만 해준다. 추후에는 struct 로 만들어 준다.
	// 디폴트 값은 https://daum.net 이다.

	/*
			Command line flag syntax

			The following forms are permitted:

				-flag
				-flag=x
				-flag x  // non-boolean flags only
			One or two minus signs may be used; they are equivalent.
			The last form is not permitted for boolean flags because the
			meaning of the command
				cmd -x *
			where * is a Unix shell wildcard, will change if there is a file
			called 0, false, etc. You must use the -flag=false form to turn
			off a boolean flag.

			Flag parsing stops just before the first non-flag argument
			("-" is a non-flag argument) or after the terminator "--".

			Integer flags accept 1234, 0664, 0x1234 and may be negative.
			Boolean flags may be:
				1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
			Duration flags accept any input valid for time.ParseDuration.

			The default set of command-line flags is controlled by
			top-level functions.  The FlagSet type allows one to define
			independent sets of flags, such as to implement subcommands
			in a command-line interface. The methods of FlagSet are
			analogous to the top-level functions for the command-line
			flag set.

			핵심적인 것은 코드상에서 - 나 -- 를 쓰면 안된다.
			아래 예에서 fs.StringVar(&s1, "-u" 또는 "--u",value,usage) 안된 다는 것이다.
		    하지만 실행단에서 즉 파라미터에서는 "-u" 또는 "--u" 가능하고 또한 동일 하다.
	*/
	// 10/14
	//fs.StringVar(&s1, "u", "https://daum.net", "https address")
	// 10/15 config struct 를 가지고 온고 디폴트 값을 저장한다.
	// 10/16 아래 두 함수를 개선할 필요가 있다.
	c := conf.DefaultConfig()
	// 이 함수에서 사용자의 입력 파라미터가 저장된다.
	conf, err := c.RegisterConfig(fs)

	if err != nil {
		os.Exit(1)
	}

	// 사용자 입력값을 파싱한다. 즉, 입력된 파라미터를 s1 에 집어 넣는다.
	fs.Parse(os.Args[1:]) // command line 의 slice 의 첫번째 파라미터 부터 끝까지

	// 여기에 문제가 있다. 찾을 수 있을까?
	if len(os.Args) < 2 {
		// 에러 보고 및 exit
		// 에러 보고 코드 들어가야 함.

		// 사용설명 출력
		fs.Usage()
		// os.Exit 의 경우 0 ~ 125 까지의 exit status 를 내보낼 수 있는데
		// 일반적으로 0 의 경우는 정상적인 종료를 의미하고 그 이외는 에러에 따른 종료로 나타낼 수 있다.
		// golang 같은 경우는 main 함수가 리턴값이 없기 때문에 이러한 방식으로 Exit status 를 확인 할 수 있다.
		// 또한 os.Exit 이 있는 경우는 defer 가 적용되지 않는다.
		os.Exit(1)
	}

	// 여기부터 사용자의 입력 파라미터를 사용하는 코드가 들어간다. 실제적인 코딩 부분
	// 10/14
	//fmt.Println("사용자 입력 파라미터", s1)
	// 10/15
	if !c.Silent {
		fmt.Println("사용자 입력 파라미터", c.S1)
		fmt.Println("config 파일로부터 읽은 데이터", conf.Filename)
		// 10/18 Version 수정
		// 이것을 Config 에 넣는 개선 방안은?
		fmt.Println("현재 Version", Version)
	}
}

// https://pkg.go.dev/flag#NewFlagSet
// https://golang.org/src/flag/flag.go
