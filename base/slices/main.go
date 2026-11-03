package main

import (
	"fmt"
)

type Agent struct {
	ID      int
	Enabled bool
}

func main() {
	agents := make([]Agent, 0, 10)
	for i := 0; i < 3; i++ {
		agents = append(agents, Agent{ID: i})
	}

	fmt.Println(cap(agents), len(agents))
	agents = addThirdPartyAgents(agents)
	addFourthPartyAgents(&agents)
	fmt.Println(cap(agents), len(agents))

	fmt.Println(agents)
}

func addThirdPartyAgents(agents []Agent) []Agent {
	thirdParty := []Agent{
		{ID: 3},
		{ID: 4},
	}

	agents = append(agents, thirdParty...)

	return agents
}

func addFourthPartyAgents(agents *[]Agent) {
	fourthParty := []Agent{
		{ID: 5},
		{ID: 6},
	}

	*agents = append(*agents, fourthParty...)
}
