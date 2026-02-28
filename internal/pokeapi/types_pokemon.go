package pokeapi

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	ID             int     `json:"id"`
	Name           *string `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (p Pokemon) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Name: %s\n", *p.Name)
	fmt.Fprintf(&b, "Height: %d\n", p.Height)
	fmt.Fprintf(&b, "Weight: %d\n", p.Weight)

	b.WriteString("Stats:\n")
	for _, stat := range p.Stats {
		fmt.Fprintf(&b, "  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	b.WriteString("Types:\n")
	for _, t := range p.Types {
		fmt.Fprintf(&b, "  - %s\n", t.Type.Name)
	}

	return b.String()
}
