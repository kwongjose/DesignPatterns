package main

import "fmt"

/**
** The TV stuct matches our device interface
** This is a different implimnation of setChannel than Radio
**/
type TV struct {
	maxChannel     int
	minChannel     int
	hasPower       bool
	currentChannel int
	currentVolume  int
}

func (tv *TV) isOn() bool {
	return tv.hasPower
}

func (tv *TV) powerOff() {
	tv.hasPower = false
	fmt.Println("TV is off")
}

func (tv *TV) powerOn() {
	tv.hasPower = true
	fmt.Println("TV is on")
}

func (tv *TV) getChannel() int {
	fmt.Printf("TV is tuned to %d \n", tv.currentChannel)
	return tv.currentChannel
}

// setChannel sets the channel for the TV.
// If the channel is higher than the max or lower than the min it sets the channel to either the min or max channel
func (tv *TV) setChannel(channel int) {
	if channel > tv.maxChannel {
		tv.currentChannel = tv.minChannel
	} else if channel < tv.minChannel {
		tv.currentChannel = tv.maxChannel
	} else {
		tv.currentChannel = channel
	}
	fmt.Printf("TV channel set to %d \n", tv.currentChannel)
}

func (tv *TV) getVolume() int {
	fmt.Printf("TV volume set at %d\n", tv.currentVolume)
	return tv.currentVolume
}

func (tv *TV) setVolume(volume int) {
	tv.currentVolume = volume
	fmt.Printf("TV volume set to %d\n", tv.currentVolume)
}
