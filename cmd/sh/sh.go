package sh

import (
	"bytes"
	"fmt"

	"github.com/pyowonil/tiny/cmd"
)

// Command is a shell command.
type Command struct {
	sh       *cmd.Command
	cmdlines []string
	stdout   *bytes.Buffer
	stderr   *bytes.Buffer
}

// NewCommand creates a new Command.
func NewCommand(cmds ...string) *Command {
	logger.Debugf("NewCommand(%v)", cmds)
	command := &Command{
		sh:       cmd.NewCommand("sh", "-c"),
		cmdlines: cmds,
		stdout:   &bytes.Buffer{},
		stderr:   &bytes.Buffer{},
	}
	return command
}

// GetCmdlines gets sh's cmdlines.
func (command *Command) GetCmdlines() []string {
	logger.Debugf("(%p *Command).GetCmdlines()", command)
	return command.cmdlines
}

// GetEnvs gets sh's envs.
func (command *Command) GetEnvs() []string {
	logger.Debugf("(%p *Command).GetEnvs()", command)
	return command.sh.GetEnvs()
}

// GetWorkDir gets sh's working directory.
func (command *Command) GetWorkDir() string {
	logger.Debugf("(%p *Command).GetWorkDir()", command)
	return command.sh.GetWorkDir()
}

// GetStdout gets command's stdout.
func (command *Command) GetStdout() []byte {
	logger.Debugf("(%p *Command).GetStdout()", command)
	out := command.stdout.Bytes()
	if out != nil {
		logger.Infof("sh's stdout: \n%s", out)
	} else {
		logger.Infof("sh's stdout is empty")
	}
	return out
}

// GetStderr gets command's stderr.
func (command *Command) GetStderr() []byte {
	logger.Debugf("(%p *Command).GetStderr()", command)
	err := command.stderr.Bytes()
	if err != nil {
		logger.Warnf("sh's stderr: \n%s", err)
	} else {
		logger.Infof("sh's stderr is empty")
	}
	return err
}

// SetCmdlines sets sh's cmdlines.
func (command *Command) SetCmdlines(cmds ...string) {
	logger.Debugf("(%p *Command).SetCmdlines(%v)", command, cmds)
	command.cmdlines = cmds
}

// SetEnvs sets sh's envs.
func (command *Command) SetEnvs(envs ...string) {
	logger.Debugf("(%p *Command).SetEnvs(%v)", command, envs)
	command.sh.SetEnvs(envs...)
}

// SetWorkDir sets sh's working directory.
func (command *Command) SetWorkDir(path string) {
	logger.Debugf("(%p *Command).SetWorkDir(%v)", command, path)
	command.sh.SetWorkDir(path)
}

// Run runs the cmdlines on the sh.
func (command *Command) Run() error {
	logger.Debugf("(%p *Command).Run()", command)
	logger.Infof("reset %p's stdout, stderr buffers", command)
	command.stdout.Reset()
	command.stderr.Reset()
	for i, cmd := range command.cmdlines {
		logger.Infof("[%d/%d] %s", i+1, len(command.cmdlines), cmd)
		command.sh.SetArgs("-c", cmd)
		if err := command.run(); err != nil {
			err := fmt.Errorf("failed to run %s command on sh: %v", cmd, err)
			logger.Errorf("%s", err)
			return err
		}
	}
	return nil
}

// run runs the cmd on the sh.
func (command *Command) run() error {
	logger.Debugf("(%p *Command).run()", command)
	cmd := command.sh.Clone()
	err := cmd.Run()
	if err := command.writeStdout(cmd.GetStdout()); err != nil {
		err = fmt.Errorf("failed to set stdout: %v", err)
		logger.Errorf("%s", err)
		return err
	}
	if err := command.writeStderr(cmd.GetStderr()); err != nil {
		err = fmt.Errorf("failed to set stderr: %v", err)
		logger.Errorf("%s", err)
		return err
	}
	return err
}

// writeStdout writes sh's stdout.
func (command *Command) writeStdout(b []byte) error {
	logger.Debugf("(%p *Command).writeStdout(%v)", command, b)
	if _, err := command.stdout.Write(b); err != nil {
		err = fmt.Errorf("write to stdout failed: %v", err)
		logger.Errorf("%s", err)
		return err
	}
	return nil
}

// writeStderr writes sh's stderr.
func (command *Command) writeStderr(b []byte) error {
	logger.Debugf("(%p *Command).writeStderr(%v)", command, b)
	if _, err := command.stderr.Write(b); err != nil {
		err = fmt.Errorf("write to stderr failed: %v", err)
		logger.Errorf("%s", err)
		return err
	}
	return nil
}
