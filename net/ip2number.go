package net

import (
	"net"

	"github.com/pkg/errors"
)

func IP2Number(ip net.IP) (uint, error) {
	b := ip.To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}
	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

func IPCountBetweenTwoIP(ip1, ip2 net.IP) (uint, error) {
	num1, err := IP2Number(ip2)
	num2, err := IP2Number(ip1)
	if err != nil {
		return 0, errors.Wrap(err, "Get count between two ip")
	}
	return num2 - num1, nil
}

func IPString2Number(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}
	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

func IPCountBetweenTwoIPString(ip1, ip2 string) (uint, error) {
	num1, err := IPString2Number(ip2)
	num2, err := IPString2Number(ip1)
	if err != nil {
		return 0, errors.Wrap(err, "Get count between two ip")
	}
	return num2 - num1, nil
}
