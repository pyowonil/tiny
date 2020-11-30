package cmd

import (
	"bytes"
	"os"
	"os/exec"
)

// Command is a tiny command.
type Command struct {
	cmd *exec.Cmd
}

// NewCommand creates a new Command.
func NewCommand(name string, args ...string) *Command {
	logger.Debugf("NewCommand(%v)", name)
	command := &Command{
		cmd: exec.Command(name, args...),
	}
	logger.Debugf("create a new command (%p *Command)", command)
	command.cmd.Stdout = &bytes.Buffer{}
	command.cmd.Stderr = &bytes.Buffer{}
	return command
}

// GetPath gets command's path.
func (command *Command) GetPath() string {
	logger.Debugf("(%p *Command).GetPath()", command)
	return command.cmd.Path
}

// GetName gets command's name.
func (command *Command) GetName() string {
	logger.Debugf("(%p *Command).GetName()", command)
	return command.cmd.Args[0]
}

// GetArgs gets command's args.
func (command *Command) GetArgs() []string {
	logger.Debugf("(%p *Command).GetArgs()", command)
	return command.cmd.Args[1:]
}

// GetEnvs gets command's envs.
func (command *Command) GetEnvs() []string {
	logger.Debugf("(%p *Command).GetEnvs()", command)
	env := command.cmd.Env
	if env == nil {
		env = os.Environ()
	}
	return env
}

// GetWorkDir gets command's working directory.
func (command *Command) GetWorkDir() string {
	logger.Debugf("(%p *Command).GetWorkDir()", command)
	wd := command.cmd.Dir
	if wd == "" {
		wd = "."
	}
	return wd
}

// GetStdout gets command's stdout.
func (command *Command) GetStdout() []byte {
	logger.Debugf("(%p *Command).GetStdout()", command)
	stdout := command.cmd.Stdout.(*bytes.Buffer)
	out := stdout.Bytes()
	if out != nil {
		logger.Infof("command's stdout: \n%s", out)
	} else {
		logger.Infof("command's stdout is empty")
	}
	return out
}

// GetStderr gets command's stderr.
func (command *Command) GetStderr() []byte {
	logger.Debugf("(%p *Command).GetStderr()", command)
	stderr := command.cmd.Stderr.(*bytes.Buffer)
	err := stderr.Bytes()
	if err != nil {
		logger.Warnf("command's stderr: \n%s", err)
	} else {
		logger.Infof("command's stderr is empty")
	}
	return err
}

// SetArgs sets command's args.
func (command *Command) SetArgs(args ...string) {
	logger.Debugf("(%p *Command).SetArgs(%v)", command, args)
	command.cmd.Args = append([]string{command.GetName()}, args...)
}

// SetEnvs sets command's envs.
func (command *Command) SetEnvs(envs ...string) {
	logger.Debugf("(%p *Command).SetEnvs(%v)", command, envs)
	command.cmd.Env = envs
}

// SetWorkDir sets command's working directory.
func (command *Command) SetWorkDir(path string) {
	logger.Debugf("(%p *Command).SetWorkDir(%s)", command, path)
	command.cmd.Dir = path
}

// Run runs the command.
func (command *Command) Run() error {
	logger.Debugf("(%p *Command).Run()", command)
	logger.Infof("(*Command).Run() - name: %v, args: %v", command.GetName(), command.GetArgs())
	err := command.cmd.Run()
	if err != nil {
		logger.Errorf("failed to run command: %v", err)
	}
	return err
}

// Clone clones the command.
func (command *Command) Clone() *Command {
	logger.Debugf("(%p *Command).Clone()", command)
	clone := NewCommand(command.GetName(), command.GetArgs()...)
	logger.Debugf("%p is %p's clone", clone, command)
	if envs := command.cmd.Env; envs != nil {
		clone.SetEnvs(envs...)
	}
	if wd := command.cmd.Dir; wd != "" {
		clone.SetWorkDir(wd)
	}
	return clone
}
