package utils

import (
	"net"
	"testing"
)

func TestCidrConvertV4(t *testing.T) {
	_, cidrInput, _ := net.ParseCIDR("192.168.20.0/23")
	netResult, maskResult := CidrStringToNetworkMask(cidrInput)
	netExpect := "192.168.20.0"
	maskExpect := "255.255.254.0"

	if netResult != netExpect {
		t.Fatalf("%s != %s", netResult, netExpect)
	}
	if maskResult != maskExpect {
		t.Fatalf("%s != %s", maskResult, maskExpect)
	}
}

func TestCidrConvertV6(t *testing.T) {
	_, cidrInput, _ := net.ParseCIDR("fd00::dead:beef/60")
	netResult, maskResult := CidrStringToNetworkMask(cidrInput)
	netExpect := "fd00::"
	maskExpect := "ffff:ffff:ffff:fff0:0000:0000:0000:0000"

	if netResult != netExpect {
		t.Fatalf("%s != %s", netResult, netExpect)
	}
	if maskResult != maskExpect {
		t.Fatalf("%s != %s", maskResult, maskExpect)
	}
}

func TestNetMaskToPrefixLength(t *testing.T) {
	inputs := []string{
		"8000::", "c000::", "e000::", "f000::", "ff00::",
	}
	outputs := []int{
		1, 2, 3, 4, 8,
	}

	for i := 0; i < len(inputs); i++ {
		if v, e := NetMaskToPrefixLength(inputs[i]); e != nil || v != outputs[i] {
			if e == nil {
				t.Fatalf("%x != %x", v, outputs[i])
			} else {
				t.Fatalf("err %v", e)
			}
		}
	}
}
