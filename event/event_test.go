//Package for testing events in the riverworld
//Possible events: put(item), getin(), getout(), crossriver(), takeout()
//start(), reset() osv.
package event


import "testing"

func TestPutIn(t *testing.T)  {
	wanted := "[kylling rev korn mann ---V \\_korn_/___________/ Ø---]"
	got := Put("korn")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}

func TestGetIn(t *testing.T)  {
	wanted := "[kylling rev ---V \\_mann+korn_/___________/ Ø---]"
	got := GetIn("korn")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}
