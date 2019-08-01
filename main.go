package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Commnad not found...")
	}

}

func run() {
	fmt.Printf("Run /proc/self/exe %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// PID namespace:
	// UTS namespace: isolate hostname and domainName
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("child %v as PID %d \n", os.Args[2:], os.Getegid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Chroot("/home/vagrant/ubuntu_xenial_1604"))
	must(os.Chdir("/"))
	must(cmd.Run())

}

func must(err error) {
	if err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		panic(err)
	}
}
