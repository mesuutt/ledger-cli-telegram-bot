package ledger

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func ExecLedgerCommand(filePath string, cmdList ...string) (bytes.Buffer, error) {
	var stdout, stderr bytes.Buffer

	// NOTE: @mesut restrict executed command with ledger filePath
	args := append([]string{"ledger", "-f", filePath}, cmdList...)

	cmd := exec.Command("bash", "-c", strings.Join(args, " "))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"cmd":    cmd.String(),
			"stderr": stderr.String(),
		}).Error("ExecLedgerCommand Error")
		return stdout, err
	}
	return stdout, nil

}

func ExecSedCommandOnFile(filePath, command string) error {
	var stdout, stderr bytes.Buffer
	sedCmd := fmt.Sprintf("sed -i %s %s", command, filePath)
	// replace interpreted NL with raw NL which couse to error when adding new alias
	sedCmd = strings.Replace(sedCmd, "\n", `\n`, -1)
	cmd := exec.Command("bash", "-c", sedCmd)
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


func InsertToBeginningOfFile(filePath, text string) error {
	sedCmd := fmt.Sprintf(`'1s/^/%s/'`, text)
	ExecSedCommandOnFile(filePath, sedCmd)
	return nil
}
