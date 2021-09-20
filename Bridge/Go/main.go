package main

func main() {

	carRadio := &Radio{
		maxChannel:     100,
		minChannel:     0,
		hasPower:       false,
		currentChannel: 0,
		currentVolume:  0,
	}

	radioRemote := Remote{device: carRadio}
	radioRemote.togglePower()
	radioRemote.volumeUp()
	radioRemote.channelDown()
}
