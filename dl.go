package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	// "syscall"
	"time"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	path := fmt.Sprintf("/app/data/%s", time.Now().Format("200601/02"))
	pathExists, _ := PathExists(path)
	if !pathExists {
		// fmt.Printf("%s not exist.\n", path)
		mkdirErr := os.MkdirAll(path, 0777)
		if mkdirErr != nil {
			fmt.Println("mkdir Err :" + mkdirErr.Error())
		}
	}

	cmdStr := fmt.Sprintf("cd %s && youtube-dl -c --no-warnings -x --audio-format mp3 https://www.youtube.com/watch?v=ud5m4oaBrUY | grep .mp3", path)
	fmt.Println(cmdStr)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	// Handle Stdout
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}
	s := strings.Split(string(bytes), ": ")
	_, mp3 := s[0], s[1]
	fmt.Println(mp3)

	// // Handle Stderr
	// stderr, err := cmd.StderrPipe()
	// if err != nil {
	// 	fmt.Println("StderrPipe: ", err.Error())
	// 	return
	// }

	// bytesErr, err := ioutil.ReadAll(stderr)
	// if err != nil {
	// 	fmt.Println("ReadAll stderr: ", err.Error())
	// 	return
	// }

	// if len(bytesErr) != 0 {
	// 	fmt.Printf("stderr is not nil: %s", bytesErr)
	// 	return
	// }

	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println("Wait: ", err.Error())
	// 	return
	// }
}
