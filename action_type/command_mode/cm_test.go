package command_mode

import "testing"

func Test(t *testing.T) {
	receiver := &Receiver{}
	command := &ConcreteCommand{receiver: receiver}
	invoker := &Invoker{command: command}
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
