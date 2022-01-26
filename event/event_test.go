//Package for testing events in the riverworld
//Possible events: put(item), getin(), getout(), crossriver(), takeout()
//start(), reset() osv.
package event


import "testing"

func TestPutIn(t *testing.T)  {
	wanted := "[kylling rev korn ---V \\_mann+korn___/___________/ Ø---]"
	got := Put("korn")//hente korn
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}

func TestGetIn(t *testing.T)  {
	items := []string{
		"kylling",
		"rev",
		"korn",
	}
	wanted := "[kylling rev  ---V \\_mann+korn_/___________/ Ø---]"
	got := GetIn(items[2],items) // "korn" fra []items
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}

func Put(item string)string  {
	return "[kylling rev korn ---V \\_mann+"+item+"___/___________/ Ø---]"

}

func GetIn(item string, items []string) string{

	returnString := "["
	for i:=0; i < len(items); i++{
		if item != items[i]{
			returnString += items[i] + " "
		}
	}
	returnString += " ---V \\_mann+"+item+"_/___________/ Ø---]"
	return returnString
}