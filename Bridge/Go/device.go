package main

// Device is our interface for similar classes
type Device interface {
	isOn() bool
	powerOn()
	powerOff()
	getVolume() int
	setVolume(percent int)
	getChannel() int
	setChannel(channel int)
}

const (
	MAX_VOLUME = 100
	MIN_VOLUME = 0
)

// Remote is our struct that works with our Device interface.
// It does not care about the impelmentation only  that it matches.
type Remote struct {
	device Device
}

func (remote *Remote) togglePower() {
	if remote.device.isOn() {
		remote.device.powerOff()
	} else {
		remote.device.powerOn()
	}
}

// volumeUp increases the volume by 1 until it reaches 100
func (remote *Remote) volumeUp() {
	currentVol := remote.device.getVolume()
	if currentVol == MAX_VOLUME {
		return
	} else {
		remote.device.setVolume(currentVol + 1)
	}
}

// volumeDown decreases the volume by 1 until it reaches 0
func (remote *Remote) volumeDown() {
	currentVol := remote.device.getVolume()
	if currentVol == MIN_VOLUME {
		return
	} else {
		remote.device.setVolume(currentVol - 1)
	}
}

func (remote *Remote) channelDown() {
	currentChan := remote.device.getChannel()
	remote.device.setChannel(currentChan - 1)
}

func (remote *Remote) channelUp() {
	currentChan := remote.device.getChannel()
	remote.device.setChannel(currentChan + 1)
}
