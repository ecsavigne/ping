package internal

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func Ping(host string, timeout time.Duration) (duration time.Duration, size, ttl int, ip string, err_ error) {
	dstIpAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return 0, 0, 0, "", err
	}
	socketConn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return 0, 0, 0, "", err
	}
	defer socketConn.Close()

	data := [56]byte{}
	copy(data[:], "ping")
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: data[:],
		},
	}

	icmpPacket, err := msg.Marshal(nil)
	if err != nil {
		return 0, 0, 0, "", err
	}

	dstUDPAddr := &net.UDPAddr{IP: dstIpAddr.IP, Zone: dstIpAddr.Zone}
	start := time.Now()
	_, err = socketConn.WriteTo(icmpPacket, dstUDPAddr)
	if err != nil {
		return 0, 0, 0, "", err
	}

	ipConn := socketConn.IPv4PacketConn()
	if ipConn == nil {
		return 0, 0, 0, "", fmt.Errorf("error getting ipv4 connection")
	}

	if err := ipConn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
		fmt.Printf("timeout: %v. error: (%v)\n", timeout, err)
		return 0, 0, 0, "", err
	}

	if err := ipConn.SetControlMessage(ipv4.FlagTTL, true); err != nil {
		return 0, 0, 0, "", err
	}

	reply := make([]byte, 1500)
	n, control, netAddr, err := ipConn.ReadFrom(reply)
	if err != nil {
		return 0, 0, 0, "", err
	}
	size = n
	ttl = control.TTL
	ip = netAddr.String()

	response, err := icmp.ParseMessage(1, reply[:n]) // 1 = ICMPv4
	if err != nil {
		return 0, 0, 0, "", err
	}

	if response.Type != ipv4.ICMPTypeEchoReply {
		return 0, 0, 0, "", fmt.Errorf("response ICMP not waited: %v", response.Type)
	}

	duration = time.Since(start)
	return duration, size, ttl, ip, nil
}
