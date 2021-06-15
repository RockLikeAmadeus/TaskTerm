package todotxt

import (
	"time"
)

type Task struct {
	complete       bool
	priority       rune
	creationDate   time.Time
	completionDate time.Time
	description    string
	project        string
	context        string
	addMetadata    map[string]string
}

func Test() Task {
	test := Task{description: "this is a test task"}
	return test
}
