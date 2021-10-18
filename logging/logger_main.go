package main

import (
	"log"
	"os"
)

/*
	golang 에 있는 기본 log 패키지를 사용하여 대부분의 로깅을 처리할 수 있다.
	http://golang.site/go/article/114-Logging 를 참고하면 기본적인 설명이 잘 설명되어 있다.

	먼저 개발자 자신이 표준로그를 그냥 사용할지 아니면 custom 로그를 제작?할지 결정해야한다.
	custom log 를 제작할 경우는 New 함수를 써서 제작하면 된다.
*/

var Info *log.Logger

// main 보다 먼저 실행되는 함수, 일반적으로 app 실행전 초기화 코드를 많이 적용한다.
func init() {

	/*
		log.New(os.Stdout, "Info: ", log.LstdFlags)
		에서
		첫번째 파라미터는

		type Writer interface {
			Write(p []byte) (n int, err error)
		}
		타입으로 받을 녀석이다. 향후 이러한 방식의 함수들을 많이 만들 수 있는데 이 부분은 향후 인터페이스에서 다루도록 한다.

		나는 여기서 os.Stdout 으로 적용했다. 여기서 왜 os.Stdout 이 가능한지는 file.go 를 찾아보면 알 수 있다.

		질문) file.go 에서 무엇 때문에 os.Stdout 이 가능한 것일까? 향후 인터페이스를 다루면 좀더 이해할 수 있는데, 혹시 이 내용을 아는 사람이 있으면 설명 부탁드린다.

		두번째 파라미터는 prefix 이다. 즉 로그로 나타낼 메세지의 앞에 나타나는 메세지 이다.
		세번재 파라미터는 log flag 이다.

		자세한 정보는 아래에서 확인 할 수 있다. ( log.go 파일에서 가져옴. )
		참고로 아래 설명에서  1 << iota 의 예는 숫자가 하나씩 증가하는 것을 의미한다.

		일반적으로

		const (
		number_one =  iota
		number_tow
		number_three
		)

		이렇게 처리할 경우 number_one, number_two, ... 등은 정수의 값이 차례로 들어간다.

		아래의 예는 쉬프트 연산을 해서 1 비트가 왼쪽으로 하나씩 밀려 나가는 형태가 된다.
		00000001
		00000010
		00000100

		// These flags define which text to prefix to each log entry generated by the Logger.
		// Bits are or'ed together to control what's printed.
		// With the exception of the Lmsgprefix flag, there is no
		// control over the order they appear (the order listed here)
		// or the format they present (as described in the comments).
		// The prefix is followed by a colon only when Llongfile or Lshortfile
		// is specified.
		// For example, flags Ldate | Ltime (or LstdFlags) produce,
		//	2009/01/23 01:23:23 message
		// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
		//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
		const (
			Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
			Ltime                         // the time in the local time zone: 01:23:23
			Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
			Llongfile                     // full file name and line number: /a/b/c/d.go:23
			Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
			LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
			Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
			LstdFlags     = Ldate | Ltime // initial values for the standard logger
		)

	*/

	Info = log.New(os.Stdout, "Info: ", log.LstdFlags)
}

func main() {

	// 표준 로그를 사용한 경우
	log.Println("test")

	// 사용자 정의 로그를 만드는 경우
	Info.Println("info test")
}
