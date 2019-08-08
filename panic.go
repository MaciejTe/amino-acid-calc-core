package main

import (
	"fmt"
	"os"
)

// This output is shown if a panic happens.
const panicOutput = `
!!!!!!!!!!!!!!!!!!!!!!!!!!! APPLICATION CRASH !!!!!!!!!!!!!!!!!!!!!!!!!!!!

Go project template crashed! This is always indicative of a bug within application.
A crash log saving and other necessary actions can be run inside panicHandler() function.
If you are experiencing this, please report crash logs and other helpful information by submitting a bug in project issues[1]. 
Thank you!
[1]: https://github.com/MaciejTe/go-project-template/issues

!!!!!!!!!!!!!!!!!!!!!!!!!!! APPLICATION CRASH !!!!!!!!!!!!!!!!!!!!!!!!!!!!
`

func panicHandler(output string) {
	// output contains the full output (including stack traces) of the
	// panic. Put it in a file or something.
	f, err := os.Create("crash.log")
	if err != nil {
		panicErr := fmt.Errorf("Failed to create crash.log file! Details: %s ", err)
		panic(panicErr)
	}
	saveDataToFile(f, output)
	fmt.Printf("%s:\n\n%s\n", panicOutput, output)
	os.Exit(1)
}


// save string to given file
func saveDataToFile(f *os.File, output string) {
	defer f.Close()
	_, err := f.WriteString(output)
	if err != nil {
		panic(err)
	}
}
