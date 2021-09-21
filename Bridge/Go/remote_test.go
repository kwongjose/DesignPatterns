package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	radio = &Radio{
		maxChannel:     100,
		minChannel:     0,
		hasPower:       false,
		currentChannel: 0,
		currentVolume:  0,
	}

	tv = &TV{
		maxChannel:     50,
		minChannel:     0,
		hasPower:       false,
		currentChannel: 0,
		currentVolume:  0,
	}

	tvRemote    = Remote{device: tv}
	radioRemote = Remote{device: radio}
)

func TestChannelDown(t *testing.T) {
	tvRemote.channelDown()
	assert.Equal(t, tv.getChannel(), tv.maxChannel)

	radioRemote.channelDown()
	assert.Equal(t, radio.getChannel(), radio.minChannel)

}
