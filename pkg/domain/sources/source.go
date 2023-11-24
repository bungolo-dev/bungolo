package sources

import "github.com/bungolow-dev/bungolow/pkg/domain/commands"

type Source struct {
}

type SourceCommands interface {
	commands.PowerControl
	commands.Navigate
}
