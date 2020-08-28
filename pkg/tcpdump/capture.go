/**
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2020 vorteil.io Pty Ltd
 */

package tcpdump

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	CaptureLogInfo func(string, ...interface{}) = func(s string, i ...interface{}) {
		os.Stdout.WriteString(fmt.Sprintf(s, i...))
	}
)

type packetCaptureManager struct {
	targetDevice    string
	capturing       bool
	snapshotLen     int32
	bpfFilter       string
	promiscuousMode bool
	timeout         time.Duration
}

func NewPacketCapture(snapshotLen int32, promiscuousMode bool, timeout time.Duration) *packetCaptureManager {
	return &packetCaptureManager{
		capturing:       false,
		snapshotLen:     snapshotLen,
		bpfFilter:       "",
		targetDevice:    "",
		timeout:         timeout,
		promiscuousMode: promiscuousMode,
	}
}

func (pCM *packetCaptureManager) SetDevice(deviceName string) error {
	if err := pCM.isCapturing(); err != nil {
		return err
	}

	pCM.targetDevice = deviceName

	return nil
}

func (pCM *packetCaptureManager) SetFilter(bpfFilter string) error {
	if err := pCM.isCapturing(); err != nil {
		return err
	}

	pCM.bpfFilter = bpfFilter

	return nil
}

func (pCM *packetCaptureManager) StartCapturing() error {
	if pCM.targetDevice == "" {
		return fmt.Errorf("no target devices set, please add a device before capturing")
	}

	if err := pCM.isCapturing(); err != nil {
		return err
	}

	// Open device
	CaptureLogInfo("Opening Device \"%s\" for Packet Capturing\n", pCM.targetDevice)
	tcpdumpHandler, err := pcap.OpenLive(pCM.targetDevice, pCM.snapshotLen, pCM.promiscuousMode, pCM.timeout)
	if err != nil {
		return fmt.Errorf("could not open device \"%s\", error=%v", pCM.targetDevice, err)
	}

	// Set Filter
	CaptureLogInfo("Setting Device \"%s\" Packet Capturing BPF Filter to \"%s\"\n", pCM.targetDevice, pCM.bpfFilter)
	err = tcpdumpHandler.SetBPFFilter(pCM.bpfFilter)
	if err != nil {
		return fmt.Errorf("could not open bpf filter \"%s\" on device \"%s\", error=%v", pCM.bpfFilter, pCM.targetDevice, err)
	}

	pCM.capturing = true

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(tcpdumpHandler, tcpdumpHandler.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		CaptureLogInfo("%v\n", packet)
	}

	return nil
}

func (pCM *packetCaptureManager) isCapturing() error {
	if pCM.capturing {
		return fmt.Errorf("cannot edit/start capture manager, already capturing packets")
	}

	return nil
}
