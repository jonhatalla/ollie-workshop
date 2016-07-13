package main

import (
	"os"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ble"
)

var ollie *ble.SpheroOllieDriver

func main() {
	gbot := gobot.NewGobot()

	bleAdaptor := ble.NewBLEClientAdaptor("ble", os.Args[1])
	ollie = ble.NewSpheroOllieDriver(bleAdaptor, "ollie")

	robot := gobot.NewRobot("ollieBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ollie},
		listenToKeyboard,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
