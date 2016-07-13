package main

import "github.com/nsf/termbox-go"

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
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				modifyDirection(true)
			case termbox.KeyArrowDown:
				modifySpeed(false)
			case termbox.KeyArrowRight:
				modifyDirection(false)
			case termbox.KeyArrowUp:
				modifySpeed(true)
			case termbox.KeyEsc:
				fmt.Println("esc?")
        os.Exit(0)
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
  config.heading += dir
  config.heading %= 360
  // need to call ollie bot
}

func modifySpeed(isForward bool){
  speed := -5
  if isFoward {
    speed = 5
  }
  config.speed += speed
  //need to call ollie bot here
}
