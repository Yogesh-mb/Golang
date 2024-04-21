package hello

import (
	"fmt"
	"strings"
)

func Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func Hey(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}
