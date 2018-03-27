package main

import (
	"github.com/jimbojsb/bimmertunes"
	"github.com/jimbojsb/ibus"
	"go.uber.org/zap"
	"sync"
)

func main() {

	logger, _ := zap.NewDevelopment()
	sugaredLogger := logger.Sugar()

	ibus.Events.Subscribe(ibus.EVENT_PACKET_SENT, func(p *ibus.Packet) {
		sugaredLogger.Debugf("TX: %v", p.String())
	})

	ibus.Events.Subscribe(ibus.EVENT_PACKET_RECEIVED, func(p *ibus.Packet) {
		sugaredLogger.Debugf("RX: %v", p.String())
	})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go ibus.Run("/dev/tty.SLAB_USBtoUART")

	//sugaredLogger.Debugf("waking up ibus, simulating IKE")
	//ibus.Ike.WakeUp()

	cdChangerEmulator := bimmertunes.NewCdChangerEmulator(sugaredLogger)
	cdChangerEmulator.Run()

	wg.Wait()
}
