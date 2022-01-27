package main

import (
	"fmt"
	"github.com/Kaduhuy/rc1-verden/event"
	"github.com/digitnow/rivercrossing/state"

	//"rc1-verden/event"
)

var items = [] string{

"kylling",
"rev",
"korn",
}


func main() {
	fmt.Println(state.ViewState())
	fmt.Println(event.FirstPut(items[0]))
	fmt.Println(event.GetIn(items[0]))
	fmt.Println(event.CrossRiver(items[0]))
	fmt.Println(event.TakeOut(items[0]))
	fmt.Println(event.GetOut(items[0]))
}
