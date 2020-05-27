//go:generate go run doc/gen_markdown.go
//go:generate go run doc/gen_manpage.go
//go:generate go run misc/gen_bash_completion.go
//go:generate go run misc/gen_powershell_completion.go
//go:generate go run misc/gen_zsh_completion.go

package main

import (
	"github.com/daedaleanai/git-ticket/commands"

	// minimal go version is 1.11
	_ "github.com/theckman/goconstraint/go1.11/gte"
)

func main() {
	commands.Execute()
}
