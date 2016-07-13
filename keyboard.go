package main

import ("github.com/nsf/termbox-go"
// "os"
"fmt")

type keyboardEventType int

type Config struct {
  heading   uint16
  speed     uint8
}

var config Config

type keyboardEvent struct {
	eventType keyboardEventType
	key       termbox.Key
}

func listenToKeyboard() {
  if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				modifyHeading(true)
			case termbox.KeyArrowDown:
				modifySpeed(false)
			case termbox.KeyArrowRight:
				modifyHeading(false)
			case termbox.KeyArrowUp:
				modifySpeed(true)
			case termbox.KeyEsc:
				fmt.Println("esc?")
        return
    }
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func modifyHeading(isLeft bool){
  dir := 5
  if isLeft {
      dir = -5
  }
  config.heading += uint16(dir)
  config.heading %= 360
  // need to call ollie bot
  fmt.Println(config)
}

func modifySpeed(isForward bool){
  speed := -5
  if isForward {
    speed = 5
  }
  config.speed += uint8(speed)
  //need to call ollie bot here
}
