package util

import (
	"encoding/binary"
	"errors"
	"net"
	"strings"
)

func Ip2long(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

func Long2Ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func IsLocal(addr net.Addr) bool {
	sl := strings.Split(addr.String(), ":")
	ip, err := Ip2long(sl[0])
	if err != nil {
		return false
	}
	return ip == 2130706433 || ip>>24 == 10 || ip>>20 == 2753 || ip>>16 == 49320
}
