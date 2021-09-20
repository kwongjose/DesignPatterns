package main

type TV struct {
	maxChannel     float64
	minChannel     float64
	hasPower       bool
	currentChannel float64
	currentVolume  int
}

func (tv *TV) isOn() bool {
	return tv.hasPower
}

func (tv *TV) powerOff() {
	tv.hasPower = false
}

func (tv *TV) powerOn() {
	tv.hasPower = true
}

func (tv *TV) getChannel() float64 {
	return tv.currentChannel
}

func (tv *TV) setChannel(channel float64) {
	if channel > tv.maxChannel {
		tv.currentChannel = tv.minChannel
	} else if channel < tv.minChannel {
		tv.currentChannel = tv.maxChannel
	} else {
		tv.currentChannel = channel
	}
}

func (tv *TV) getVolume() int {
	return tv.currentVolume
}

func (tv *TV) setVolume(volume int) {
	tv.currentVolume = volume
}
