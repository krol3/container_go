package pkg

import (
	"log"
	"os"
	"os/exec"
)

func ExecuteRun(input string, arg ...string) {
	log.Printf("Running %v \n", input)

	cmd := exec.Command(input)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
