package command

import (
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

type ProcessManager struct {
	sync.RWMutex
	cmd *exec.Cmd
}

func NewProcessManager(cmd *exec.Cmd) *ProcessManager {
	return &ProcessManager{cmd: cmd}
}

func (p *ProcessManager) Signal(sig os.Signal) {
	p.RLock()
	defer p.RUnlock()
	if p.Available() {
		p.cmd.Process.Signal(sig)
	}
}

func (p *ProcessManager) Graceful(sig os.Signal) {
	p.RLock()
	defer p.RUnlock()

	p.Signal(sig)

	time.Sleep(1 * time.Second)

	p.Signal(syscall.SIGKILL)
	p.cmd = nil
}

func (p *ProcessManager) Available() bool {
	p.RLock()
	defer p.RUnlock()

	return p.cmd != nil
}
