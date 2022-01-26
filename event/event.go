package event

func Put(item string)string  {
	return "[rev korn hs ---V \\_"+item+"___/___________/ Ø---]"

}

func GetIn(item string)string  {
	return "[rev korn ---V \\_hs+"+item+"_/___________/ Ø---]"

}

/* Tester ut kode
func GetIn(item string, items []string) string{
	returnString := "["
	for i:=0; i < len(items); i++{
		if item != items[i]{
			returnString += items[i] + " "
		}
	}
	returnString += " ---V \\_mann+"+item+"_/___________/ Ø---]"
	return returnString
}*/

func CrossRiver(item string)string  {
	return "[rev korn ---V \\___________\\_hs+"+item+"___/ Ø---]"

}

func TakeOut(item string)string  {
	return "[rev korn ---V \\___________\\_hs___/ Ø--- "+item+"]"

}

func GetOut(item string)string{
	return "[rev korn ---V \\___________\\____/ Ø--- hs "+item+"]"
}