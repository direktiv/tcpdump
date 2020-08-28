package tcpdump

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDevice(t *testing.T) {

	name := "lo"
	filter := "tcp"

	m := NewPacketCaptureManager(512, DefaultPromiscuousMode, DefaultTimeout)

	assert.Equal(t, m.snapshotLen, int32(512), "snapshot length not equal")
	assert.Equal(t, m.promiscuousMode, false, "promisuous bool not equal")
	assert.Equal(t, m.timeout, DefaultTimeout, "timeout not equal")

	m.SetDevice(name)
	assert.Equal(t, m.targetDevice, name, "names not equal")

	m.SetFilter(filter)
	assert.Equal(t, m.bpfFilter, filter, "filter not equal")

	m.StartCapturing()

}
