package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", `youtube-dl -c --no-warnings -x --audio-format mp3 https://www.youtube.com/watch?v=ud5m4oaBrUY | grep .mp3`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}

	// Handle Stdout
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}

	fmt.Printf("stdout: %s", bytes)
}
