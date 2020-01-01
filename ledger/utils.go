package ledger

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func ExecLedgerCommand(filePath string, cmdList ...string) (bytes.Buffer, bytes.Buffer) {
	var outb, errb bytes.Buffer

	// NOTE: @mesut restrict executed command with ledger filePath
	args := append([]string{"-f", filePath}, cmdList...)

	cmd := exec.Command("ledger", args...)
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	cmd.Run()
	return outb, errb

}

func ExecSedCommandOnFile(filePath, command string) error {
	var stdout, stderr bytes.Buffer
	sedCmd := fmt.Sprintf("/usr/bin/sed -i %s %s", command, filePath)
	cmd := exec.Command("/usr/bin/bash", "-c", sedCmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"cmd":    cmd.String(),
			"stderr": stderr.String(),
		}).Error("ExecSedCommandOnFile Error")
		return err
	}

	return nil
}
