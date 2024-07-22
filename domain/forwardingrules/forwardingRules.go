package forwardingrules

import "strings"

type Rule struct {
	Type        string   `json:"type"`
	Subject     string   `json:"subject"`
	Source      string   `json:"source"`
	Target      string   `json:"target"`
	Identifiers []string `json:"identifiers"`
}

func (r Rule) GetAllSubjects() []string {
	methods := make([]string, len(r.Identifiers))
	for i := range methods {
		methods[i] = strings.ReplaceAll(r.Subject, "{{id}}", r.Identifiers[i])
	}
	return methods
}
