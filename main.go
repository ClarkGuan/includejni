package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "lack of runnable program")
		os.Exit(1)
	}

	javahome, err := findJavaHome(runtime.GOOS)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	headerName, err := osHeaderName(runtime.GOOS)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cflags := os.Getenv("CGO_CFLAGS")
	cxxflags := os.Getenv("CGO_CXXFLAGS")

	cflags = fmt.Sprintf("CGO_CFLAGS=-I%s/include -I%s/include/%s %s", javahome, javahome, headerName, cflags)
	cxxflags = fmt.Sprintf("CGO_CXXFLAGS=-I%s/include -I%s/include/%s %s", javahome, javahome, headerName, cxxflags)

	command := exec.Command(os.Args[1], os.Args[2:]...)
	command.Env = append([]string{}, os.Environ()...)
	command.Env = append(command.Env, cflags, cxxflags)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin

	//log.Println(command.Env)

	if err = command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func findJavaHome(osName string) (string, error) {
	if env := os.Getenv("JAVA_HOME"); len(env) > 0 {
		return env, nil
	}

	switch osName {
	case "darwin":
		command := exec.Command("/usr/libexec/java_home")
		content, err := command.CombinedOutput()
		return string(bytes.TrimSpace(content)), err

	default:
		return "", errors.New("no $JAVA_HOME")
	}
}

func osHeaderName(osName string) (string, error) {
	switch osName {
	case "windows":
		return "win32", nil

	case "darwin", "linux":
		return osName, nil

	default:
		return "", fmt.Errorf("unsupport os %q", osName)
	}
}
