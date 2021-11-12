package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// TODO 11/12 버그 있음.
// 참고
// https://stackoverflow.com/questions/43610646/want-to-write-from-os-stdin-to-os-stdout-using-channels
// https://stackoverflow.com/questions/48353768/capture-stdout-from-command-exec-in-real-time
// https://gist.github.com/mxschmitt/6c07b5b97853f05455c3fdaf48b1a8b6

func main() {
	s := "./client_test/date_tester.sh"
	ScriptRunnerA(s)
	ch := Reply()

	for m := range ch {
		fmt.Scanln(">", m)
	}
}

// stdio 에 바로 넣음.
func ScriptRunnerA(s string) {
	cmd := exec.Command(s)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func Reply() <-chan string {
	r := make(chan string, 1)

	// TODO 11/12 error prone. 언제 끝날지 생각하자.
	go func() {
		// 왜 고루틴에 넣는지 잘 생각할 것
		defer close(r)
		scan := bufio.NewScanner(os.Stdout)

		for scan.Scan() {
			s := scan.Text()
			r <- s
		}
	}()
	return r
}
