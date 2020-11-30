package sh

import (
	"testing"
)

func TestNewCommand(t *testing.T) {
	_ = NewCommand()
	_ = NewCommand("ls -al")
}

func TestCommand_GetCmdlines(t *testing.T) {
	command := NewCommand()
	_ = command.GetCmdlines()
}

func TestCommand_GetEnvs(t *testing.T) {
	command := NewCommand()
	_ = command.GetEnvs()
}

func TestCommand_GetWorkDir(t *testing.T) {
	command := NewCommand()
	_ = command.GetWorkDir()
}

func TestCommand_GetStdout(t *testing.T) {
	command := NewCommand("ls", "ls -a", "ls -al")
	if err := command.Run(); err != nil {
		t.Error(err)
	}
	if out := command.GetStdout(); out == nil {
		t.Fail()
	}
	if out := command.GetStdout(); out == nil {
		t.Fail()
	}

	command = NewCommand("ls > /dev/null")
	if err := command.Run(); err != nil {
		t.Error(err)
	}
	if out := command.GetStdout(); out != nil {
		t.Fail()
	}
}

func TestCommand_GetStderr(t *testing.T) {
	command := NewCommand("ls", "ls -0INVALIDFLAG", "ls -al")
	if err := command.Run(); err == nil {
		t.Fail()
	}
	if err := command.GetStderr(); err == nil {
		t.Fail()
	}
	if err := command.GetStderr(); err == nil {
		t.Fail()
	}

	command = NewCommand("ls > /dev/null")
	if err := command.Run(); err != nil {
		t.Error(err)
	}
	if err := command.GetStderr(); err != nil {
		t.Fail()
	}
}

func TestCommand_SetCmdlines(t *testing.T) {
	command := NewCommand()
	command.SetCmdlines("ls", "ls -al")
}

func TestCommand_SetEnvs(t *testing.T) {
	command := NewCommand()
	command.SetEnvs("KEY=VALUE")
}

func TestCommand_SetWorkDir(t *testing.T) {
	command := NewCommand()
	command.SetWorkDir("/")
}

func TestCommand_Run(t *testing.T) {
	command := NewCommand()
	if err := command.Run(); err != nil {
		t.Error(err)
	}
	command.SetCmdlines("ls", "ls -al", "ls -al / |grep tmp")
	if err := command.Run(); err != nil {
		t.Error(err)
	}

	command = NewCommand("INVALID-COMMAND")
	if err := command.Run(); err == nil {
		t.Fail()
	}
}
