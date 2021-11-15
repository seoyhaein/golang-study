package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// 참고
// https://stackoverflow.com/questions/43610646/want-to-write-from-os-stdin-to-os-stdout-using-channels
// https://stackoverflow.com/questions/48353768/capture-stdout-from-command-exec-in-real-time
// https://gist.github.com/mxschmitt/6c07b5b97853f05455c3fdaf48b1a8b6
// https://groups.google.com/g/golang-nuts/c/MN_W1_oAFrs
// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
// https://medium.com/rungo/executing-shell-commands-script-files-and-executables-in-go-894814f1c0f7

func main() {
	// s := "./client_test/date_tester.sh"
	s := "./date_tester.sh"
	cmd, r := ScriptRunner(s)

	go func(cmd *exec.Cmd) {
		if cmd != nil {
			if err := cmd.Start(); err != nil {
				log.Printf("Error starting Cmd: %v", err)
				return
			}
			if err := cmd.Wait(); err != nil {
				log.Printf("Error waiting for Cmd: %v", err)
				return
			}
		}
	}(cmd)

	ch := Reply(r)

	for m := range ch {
		fmt.Println(">", m)
	}
}

func ScriptRunner(s string) (*exec.Cmd, io.Reader) {
	cmd := exec.Command(s)

	// StdoutPipe 쓰면 Run 및 기타 Run 을 포함한 method 를 쓰면 에러난다.
	r, err := cmd.StdoutPipe()
	if err != nil {
		log.Panicf("Error stdout pipe for Cmd: %v", err)
	}

	return cmd, r
}

func Reply(i io.Reader) <-chan string {
	r := make(chan string, 1)
	go func() {
		// 왜 고루틴에 넣는지 잘 생각할 것
		// 스크립트 실행을 기다리지 않고 실시간으로 결과를 출력하기 위해서?
		defer close(r)
		scan := bufio.NewScanner(i)

		for scan.Scan() {
			s := scan.Text()
			r <- s
		}
	}()
	return r
}
