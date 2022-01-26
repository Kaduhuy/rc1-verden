package main


var tilstand string

func ViewState()string{
	return Tilstander()
}

func Tilstander()string{
	tilstand = "[kylling rev korn mann ---V \\____/___________/ Ø---]"
	return tilstand
}