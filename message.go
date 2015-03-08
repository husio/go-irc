package irc

import (
	"strings"
)

// IRC message format:
//
// :<prefix> <command> <params> :<trailing>
type Message struct {
	Raw      string
	Prefix   string
	Command  string
	Params   []string
	Trailing string
}

// ParseLine return Message created from parsing given line.
func ParseLine(raw string) (*Message, error) {
	raw = strings.TrimSpace(raw)
	m := &Message{Raw: raw}
	if raw[0] == ':' {
		chunks := strings.SplitN(raw, " ", 2)
		m.Prefix = chunks[0][1:]
		raw = chunks[1]
	}
	chunks := strings.SplitN(raw, " ", 2)
	m.Command = chunks[0]
	raw = chunks[1]

	if raw[0] != ':' {
		chunks := strings.SplitN(raw, " :", 2)
		m.Params = strings.Split(chunks[0], " ")
		if len(chunks) == 2 {
			raw = chunks[1]
		} else {
			raw = ""
		}
	}

	if len(raw) > 0 {
		if raw[0] == ':' {
			raw = raw[1:]
		}
		m.Trailing = raw
	}
	return m, nil
}

func (m *Message) String() string {
	return m.Raw
}

// Nick return author nickname or empty string if cannot be determined.
func (m *Message) Nick() string {
	if m.Prefix == "" {
		return ""
	}
	return strings.SplitN(m.Prefix, "!", 2)[0]
}
