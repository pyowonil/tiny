package cmd

import (
	"os"
	"testing"
)

func TestNewCommand(t *testing.T) {
	_ = NewCommand("ls")
	_ = NewCommand("ls", "-al")
}

func TestCommand_GetPath(t *testing.T) {
	command := NewCommand("ls")
	_ = command.GetPath()
}

func TestCommand_GetName(t *testing.T) {
	command := NewCommand("ls")
	if command.GetName() != "ls" {
		t.Fail()
	}
}

func TestCommand_GetArgs(t *testing.T) {
	command := NewCommand("ls")
	if len(command.GetArgs()) != 0 {
		t.Fail()
	}

	command = NewCommand("ls", "-al")
	if len(command.GetArgs()) != 1 {
		t.Fail()
	}
}

func TestCommand_GetEnvs(t *testing.T) {
	command := NewCommand("ls")
	if len(command.GetEnvs()) != len(os.Environ()) {
		t.Fail()
	}

	command.SetEnvs("KEY=VALUE")
	if len(command.GetEnvs()) != 1 {
		t.Fail()
	}
}

func TestCommand_GetWorkDir(t *testing.T) {
	command := NewCommand("ls")
	if command.GetWorkDir() != "." {
		t.Fail()
	}
}

func TestCommand_GetStdout(t *testing.T) {
	command := NewCommand("ls")
	if command.GetStdout() != nil {
		t.Fail()
	}
	if err := command.Run(); err != nil {
		t.Error(err)
	}
	if command.GetStdout() == nil {
		t.Fail()
	}
}

func TestCommand_GetStderr(t *testing.T) {
	command := NewCommand("ls", "-0INVALIDFLAG")
	if command.GetStderr() != nil {
		t.Fail()
	}
	if err := command.Run(); err == nil {
		t.Fail()
	}
	if command.GetStderr() == nil {
		t.Fail()
	}
}

func TestCommand_SetArgs(t *testing.T) {
	command := NewCommand("ls")
	command.SetArgs("-al")
}

func TestCommand_SetEnvs(t *testing.T) {
	command := NewCommand("ls")
	command.SetEnvs("KEY=VALUE")
}

func TestCommand_SetWorkDir(t *testing.T) {
	command := NewCommand("ls")
	command.SetWorkDir("/")
}

func TestCommand_Run(t *testing.T) {
	command := NewCommand("ls")
	if err := command.Run(); err != nil {
		t.Error(err)
	}
	// do not support reusable.
	if err := command.Run(); err == nil {
		t.Fail()
	}
	command = NewCommand("INVALID-COMMAND")
	if err := command.Run(); err == nil {
		t.Fail()
	}
}

func TestCommand_Clone(t *testing.T) {
	command := NewCommand("ls")
	if err := command.Clone().Run(); err != nil {
		t.Error(err)
	}
	if err := command.Clone().Run(); err != nil {
		t.Error(err)
	}
	command.SetWorkDir("/")
	command.SetEnvs("KEY=VALUE")
	_ = command.Run()
	if err := command.Clone().Run(); err != nil {
		t.Error(err)
	}
}
