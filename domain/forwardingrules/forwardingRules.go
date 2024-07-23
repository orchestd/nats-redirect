package forwardingrules

type Rule struct {
	Type      string   `json:"type"`
	Source    string   `json:"source"`
	Target    string   `json:"target"`
	Subjects  []string `json:"subjects"`
	TimeoutMS int      `json:"timeoutMS"`
}
