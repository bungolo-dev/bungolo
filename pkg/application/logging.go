package bungolow

import (
	golog "github.com/ewilliams0305/golog/logger"
	writer "github.com/ewilliams0305/golog/writers/fmtsink"
)

var Logger golog.Logger = golog.LoggingConfiguration().
	Configure(golog.Debug, "").
	WriteTo(&writer.FmtPrinter{}).
	CreateLogger()
