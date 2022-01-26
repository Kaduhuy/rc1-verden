//Package for testing events in the riverworld
//Possible events: put(item), getin(), getout(), crossriver(), takeout()
//start(), reset() osv.
package event

import (
	"testing"
)

func TestPutIn(t *testing.T)  {
	wanted := "[rev korn hs ---V \\_kylling___/___________/ Ø---]"
	got := Put("kylling")//hente korn
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestGetIn(t *testing.T)  {
	wanted := "[rev korn ---V \\_hs+kylling_/___________/ Ø---]"
	got := GetIn("kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}

func TestCrossRiver(t *testing.T)  {
	wanted := "[rev korn ---V \\___________\\_hs+kylling___/ Ø---]"
	got := CrossRiver("kylling")//levere kylling
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}

func TestTakeOut(t *testing.T){
	wanted := "[rev korn ---V \\___________\\_hs___/ Ø--- kylling]"
	got := TakeOut("kylling") //levere kylling
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestGetOut(t *testing.T){
	wanted := "[rev korn ---V \\___________\\____/ Ø--- hs kylling]"
	got := GetOut("kylling")//levere kylling
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}


