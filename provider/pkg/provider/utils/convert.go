package utils

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func CidrStringToNetworkMask(ipNet *net.IPNet) (network string, mask string) {
	network = ipNet.IP.String()

	var b strings.Builder

	_, maskLen := ipNet.Mask.Size()
	if maskLen == 128 {
		for i := 0; i < 7; i++ {
			s := ipNet.Mask[2*i : (2*i)+2]
			b.WriteString(hex.EncodeToString(s))
			b.WriteRune(':')
		}
		s := ipNet.Mask[14:16]
		b.WriteString(hex.EncodeToString(s))
	} else if maskLen == 32 {
		for i := 0; i < 3; i++ {
			v := int(ipNet.Mask[i])
			b.WriteString(strconv.Itoa(v))
			b.WriteRune('.')
		}
		b.WriteString(strconv.Itoa(int(ipNet.Mask[3])))
	} else {
		panic("CIDR mask len not 32 or 128")
	}

	return network, b.String()
}

func NetMaskToPrefixLength(mask string) (int, error) {
	maskIp := net.ParseIP(mask)
	if maskIp == nil {
		return 0, fmt.Errorf("invalid netmask %s", mask)
	}

	var prefixLength int
	if maskIp4 := maskIp.To4(); maskIp4 != nil {
		prefixLength, _ = net.IPMask(maskIp4).Size()
	} else if maskIp6 := maskIp.To16(); maskIp6 != nil {
		prefixLength, _ = net.IPMask(maskIp6).Size()
	} else {
		return 0, fmt.Errorf("invalid netmask %s", mask)
	}

	return prefixLength, nil
}

func BuildCidrString(ip string, mask string) (string, error) {
	prefixLen, err := NetMaskToPrefixLength(mask)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.WriteString(ip)
	b.WriteRune('/')
	b.WriteString(strconv.Itoa(prefixLen))

	return b.String(), nil
}
