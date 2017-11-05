package command

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var pm *ProcessManager

func ExecuteCommand(args []string) {
	if err := executeChild(args); err != nil {
		log.Fatalf("error executing command %q: %v", args, err)
	}

	waitForSignal()
}

func waitForSignal() {
	defer func() { log.Printf("all processes terminated") }()

	interrupts := make(chan os.Signal)
	signal.Notify(interrupts, syscall.SIGINT, syscall.SIGTERM)

	otherSigs := make(chan os.Signal)
	signal.Notify(otherSigs, syscall.SIGHUP)

	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-tick:
			if !pm.Available() {
				return
			}
		case sig := <-otherSigs:
			pm.Signal(sig)
		case sig := <-interrupts:
			pm.Graceful(sig)
		}
	}
}

func executeChild(args []string) error {
	var cmd *exec.Cmd

	cmd = &exec.Cmd{Stdout: os.Stdout, Stderr: os.Stderr}
	path, err := exec.LookPath(os.ExpandEnv(args[0]))

	if err != nil {
		return err
	}

	if len(args) > 1 {
		cmd.Args = append([]string{path}, args[1:]...)
	}
	cmd.Path = path

	if err := cmd.Start(); err != nil {
		return err
	}

	pm = NewProcessManager(cmd)

	go func() {
		err := cmd.Wait()
		pid := cmd.Process.Pid

		switch err {
		default:
			_, ok := err.(*os.SyscallError)
			if !ok {
				log.Printf("pid %d finished: %v, error: %v", pid, cmd.Args, err)
				break
			}
			fallthrough
		case nil:
			log.Printf("pid %d finished: %v", pid, cmd.Args)

		}
		pm.Graceful(syscall.SIGINT)
	}()

	return nil
}
