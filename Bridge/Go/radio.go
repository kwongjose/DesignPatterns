package main

import "fmt"

type Radio struct {
	maxChannel     int
	minChannel     int
	hasPower       bool
	currentChannel int
	currentVolume  int
}

func (radio *Radio) isOn() bool {
	return radio.hasPower
}

func (radio *Radio) powerOff() {
	radio.hasPower = false
	fmt.Println("Radio is off")
}

func (radio *Radio) powerOn() {
	radio.hasPower = true
	fmt.Println("Radio is on")
}

func (radio *Radio) getChannel() int {
	fmt.Printf("Radio is tuned to %d \n", radio.currentChannel)
	return radio.currentChannel
}

func (radio *Radio) setChannel(channel int) {
	if channel > radio.maxChannel {
		radio.currentChannel = radio.minChannel
	} else if channel < radio.minChannel {
		radio.currentChannel = radio.maxChannel
	} else {
		radio.currentChannel = channel
	}
	fmt.Printf("radio channel set to %d \n", radio.currentChannel)
}

func (radio *Radio) getVolume() int {
	fmt.Printf("Radio volume set at %d\n", radio.currentVolume)
	return radio.currentVolume
}

func (radio *Radio) setVolume(volume int) {
	radio.currentVolume = volume
	fmt.Printf("Radio volume set to %d\n", radio.currentVolume)
}
