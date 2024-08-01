package mediator

import "testing"

func Test(t *testing.T) {
	mediator := &ConcreteMediator{
		colleagueMap: make(map[string]Colleague),
	}

	colleagueA := &ConcreteColleague_A{
		mediator: mediator,
		name:     "ColleagueA",
	}

	colleagueB := &ConcreteColleague_B{
		mediator: mediator,
		name:     "ColleagueB",
	}

	mediator.Register("ColleagueA", colleagueA)
	mediator.Register("ColleagueB", colleagueB)

	colleagueA.SendMessage("Hello, ColleagueB")
	colleagueB.SendMessage("Hello, ColleagueA")
}
