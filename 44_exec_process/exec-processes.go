package main

import (
    "os"
    "os/exec"
    "syscall"
)

// to completely replace the current Go process with another (perhaps non-Go) one
func main() {

	//  Go requires an absolute path to the binary we want to execute, so we’ll use exec.LookPath to find it
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

	// Exec requires arguments in slice form (as opposed to one big string).
	// Note that the first argument should be the program name
    args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use
    env := os.Environ()

	// If this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process.
	// If there is an error we’ll get a return value.
	// our program will be replaced by ls
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}