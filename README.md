# tcpdump

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/cfea20057d8740899fa1ef66a90dfe7e)](https://app.codacy.com/gh/vorteil/tcpdump?utm_source=github.com&utm_medium=referral&utm_content=vorteil/tcpdump&utm_campaign=Badge_Grade_Dashboard) [![Build Status](https://travis-ci.org/vorteil/tcpdump.svg?branch=master)](https://travis-ci.org/vorteil/tcpdump) [![codecov](https://codecov.io/gh/vorteil/tcpdump/branch/master/graph/badge.svg)](https://codecov.io/gh/vorteil/tcpdump)

tcpdump is a simple tcp packet logging tool for [vorteil.io](http://www.vorteil.io).

To enable tcpdump on a [vorteil.io](http://www.vorteil.io) micro vm it needs to be enabled in the network configuration:

```toml
[[network]]
  tcpdump = true
```

This captures traffic on the network device and prints the results to stdout:

```sh
PACKET: 150 bytes, wire length 150 cap length 150 @ 2020-08-28 12:52:12.164536 +1000 AEST
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..136..] SrcMAC=2c:4d:54:56:38:a0 DstMAC=bc:30:d9:a7:58:65 EthernetType=IPv6 Length=0}
- Layer 2 (40 bytes) = IPv6	{Contents=[..40..] Payload=[..96..] Version=6 TrafficClass=0 FlowLabel=139359 Length=96 NextHeader=TCP HopLimit=64 SrcIP=2001:8003:749f:9f01:d92e:44e2:2691:4d7b DstIP=2404:6800:4006:805::200a HopByHop=nil}
- Layer 3 (32 bytes) = TCP	{Contents=[..32..] Payload=[..64..] SrcPort=39140 DstPort=443(https) Seq=2245258945 Ack=2682595307 DataOffset=8 FIN=false SYN=false RST=false PSH=true ACK=true URG=false ECE=false CWR=false NS=false Window=506 Checksum=14915 Urgent=0 Options=[TCPOption(NOP:), TCPOption(NOP:), TCPOption(Timestamps:484163095/1738587300 0x1cdbbe1767a0bca4)] Padding=[]}
- Layer 4 (64 bytes) = Payload	64 byte(s)

```
