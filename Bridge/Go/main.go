package main

// This is our client that uses the remote.

func main() {

	carRadio := &Radio{
		maxChannel:     100,
		minChannel:     0,
		hasPower:       false,
		currentChannel: 0,
		currentVolume:  0,
	}

	tv := &TV{
		maxChannel:     50,
		minChannel:     0,
		hasPower:       false,
		currentChannel: 0,
		currentVolume:  0,
	}

	// Our remote for a radio. It's a cool car radio with a remote!
	radioRemote := Remote{device: carRadio}
	radioRemote.togglePower()
	radioRemote.volumeUp()
	radioRemote.channelDown()
	radioRemote.volumeDown()
	radioRemote.channelUp()

	// Our remote for a tv
	tvRemote := Remote{device: tv}
	tvRemote.togglePower()
	tvRemote.volumeUp()
	tvRemote.channelDown()
	tvRemote.volumeDown()
	tvRemote.channelUp()
}
