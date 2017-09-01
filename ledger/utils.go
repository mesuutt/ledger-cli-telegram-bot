package ledger

import (
	"bytes"
	"os/exec"
)

type Utils struct{}

func (u *Utils) ExecLedgerCommand(filePath string, cmdList ...string) (stdout bytes.Buffer, stderr bytes.Buffer) {
	var outb, errb bytes.Buffer

	// DONTFIXME: @mesut restrict executed command with ledger filePath
	args := append([]string{"-f", filePath}, cmdList...)

	cmd := exec.Command("ledger", args...)
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	cmd.Run()
	stdout = outb
	stderr = errb
	return

}
