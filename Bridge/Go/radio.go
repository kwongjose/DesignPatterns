package main

import "fmt"

/**
** The radio stuct matches our device interface
** This is a different implimnation of setChannel than TV
**/
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

// setChannel sets the channel for the radio.
// If the desired channell is higher than the max or lower than the min channel it does not change the channel
func (radio *Radio) setChannel(channel int) {
	if channel > radio.maxChannel || channel < radio.minChannel {
		return
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
