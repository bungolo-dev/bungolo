package commands

type Navigate interface {
	ArrowCommands
	EnterCommands
}

type ArrowCommands interface {
	PressUp()
	ReleaseUp()
	PressDown()
	ReleaseDown()
	PressLeft()
	ReleaseLeft()
	PressRight()
	ReleaseRight()
}

type EnterCommands interface {
	Select()
}
