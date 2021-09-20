package main

// appliance is our interface for similar classes
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

func (remote *Remote) volumeUp() {
	currentVol := remote.device.getVolume()
	if currentVol == MAX_VOLUME {
		return
	} else {
		remote.device.setVolume(currentVol + 1)
	}
}

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
