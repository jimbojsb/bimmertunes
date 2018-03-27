package bimmertunes

import (
	"github.com/jimbojsb/ibus"
	"go.uber.org/zap"
	"time"
)

type CdChangerEmulator struct {
	hasBeenPinged bool
	logger        *zap.SugaredLogger
}

func NewCdChangerEmulator(logger *zap.SugaredLogger) *CdChangerEmulator {
	c := CdChangerEmulator{
		hasBeenPinged: false,
		logger:        logger,
	}
	return &c
}

func (c *CdChangerEmulator) Run() {
	ibus.Events.Subscribe(ibus.EVENT_CDPLAYER_PING, func() {
		c.logger.Debugf("cd changer received ping from radio")
		c.hasBeenPinged = true
		ibus.CdPlayer.Pong()
	})

	ibus.Events.Subscribe(ibus.EVENT_CDPLAYER_STATUS, func() {
		c.logger.Debugf("cd changer received status request")
		ibus.CdPlayer.RespondToStatusRequest(false, 1, 1)
	})

	ibus.Events.Subscribe(ibus.EVENT_CDPLAYER_CONTROL_PLAY, func() {
		c.logger.Debugf("cd changer start playback")
		ibus.CdPlayer.RespondToStatusRequest(true, 1, 1)
	})

	ibus.Events.Subscribe(ibus.EVENT_CDPLAYER_CONTROL_STOP, func() {
		c.logger.Debugf("cd changer stop playback")
		ibus.CdPlayer.RespondToStatusRequest(false, 1, 1)
	})

	ibus.Events.Subscribe(ibus.EVENT_CDPLAYER_CONTROL_NEXT_TRACK, func() {
		c.logger.Debugf("cd changer next track")
		ibus.CdPlayer.RespondToStatusRequest(true, 1, 1)
	})

	ibus.Events.Subscribe(ibus.EVENT_CDPLAYER_CONTROL_PREVIOUS_TRACK, func() {
		c.logger.Debugf("cd changer previous track")
		ibus.CdPlayer.RespondToStatusRequest(true, 1, 1)
	})

	go func() {
		time.Sleep(1 * time.Second)
		for {
			if !c.hasBeenPinged {
				c.logger.Debugf("cd changer announcement ever 30 seconds until pinged")
				ibus.CdPlayer.Announce()
				time.Sleep(30 * time.Second)
			} else {
				return
			}
		}
	}()
}
