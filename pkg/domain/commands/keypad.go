package commands

type Keypad interface {
	Keypad0()
	Keypad1()
	Keypad2()
	Keypad3()
	Keypad4()
	Keypad5()
	Keypad6()
	Keypad7()
	Keypad8()
	Keypad9()
}

type KeypadEnter interface {
	KeypadEnter()
}

type PhoneKeypad interface {
	Star()
	Pound()
}
