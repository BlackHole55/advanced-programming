package Agents

import (
	"fmt"
)

type Agent interface {
	GetID() string
	GetName() string
}

type HumanAgent struct {
	ID   string
	Name string
}

type BotAgent struct {
	ID      string
	Name    string
	Version string
}

func (h HumanAgent) GetID() string {
	return fmt.Sprintf("%s", h.ID)
}

func (h HumanAgent) GetName() string {
	return fmt.Sprintf("%s", h.Name)
}

func (b BotAgent) GetID() string {
	return fmt.Sprintf("%s", b.ID)
}
func (b BotAgent) GetName() string {
	return fmt.Sprintf("%s", b.Name)
}
func (b BotAgent) GetVersion() string {
	return fmt.Sprintf("%s", b.Version)
}

func FormatAgent(a Agent) string {
	formatedAgent := ""

	switch v := a.(type) {
	case HumanAgent:
		formatedAgent = fmt.Sprintf("%s | %s", v.GetID(), v.GetName())
	case BotAgent:
		formatedAgent = fmt.Sprintf("%s | %s | %s", v.GetID(), v.GetName(), v.Version)
	}

	return formatedAgent
}
