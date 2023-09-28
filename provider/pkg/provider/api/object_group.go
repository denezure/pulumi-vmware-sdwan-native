package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	models "github.com/nick-barrett/pulumi-veco/provider/pkg/provider/models/swagger"
	"github.com/nick-barrett/pulumi-veco/provider/pkg/provider/utils"
)

type Prefix string

func NewPrefixFromNetworkMask(network string, mask string) (Prefix, error) {
	s, e := utils.BuildCidrString(network, mask)
	if e == nil {
		return Prefix(s), nil
	}
	return Prefix(s), e
}

type Domain string

func DeleteObjectGroup(c *Client, id int64) error {
	_, err := c.DoPortal("enterprise/deleteObjectGroup", getAddrGroupReq{ID: id})

	return err
}

type insertResponse struct {
	ID int64 `json:"id"`
}

func InsertAddressGroup(c *Client, name string, prefixes []Prefix, domains []Domain) (int64, error) {
	req := addrGrp{Name: name, Type: "address_group", Data: buildAddressGroupData(prefixes, domains)}

	raw_res, err := c.DoPortal("enterprise/insertObjectGroup", &req)
	if err != nil {
		return 0, err
	}

	var res insertResponse
	err = json.Unmarshal(raw_res, &res)
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func UpdateAddressGroup(c *Client, id int64, name string, prefixes []Prefix, domains []Domain) error {
	_, err := c.DoPortal("enterprise/updateObjectGroup", addrGrp{
		ID:   id,
		Data: buildAddressGroupData(prefixes, domains),
	})

	return err
}

func GetAddressGroup(c *Client, id int64) (name string, prefixes []Prefix, domains []Domain, e error) {
	rawGroup, err := c.DoPortal("enterprise/getObjectGroups", getAddrGroupReq{ID: id})
	if err != nil {
		return "", nil, nil, err
	}

	var addrGroups addrGrpResult
	err = json.Unmarshal(rawGroup, &addrGroups)
	if err != nil {
		return "", nil, nil, err
	}

	if len(addrGroups) != 1 {
		return "", nil, nil, fmt.Errorf("API provided incorrect number of address groups")
	}

	addrGroup := addrGroups[0]

	domains = make([]Domain, 0, len(addrGroup.Data))
	prefixes = make([]Prefix, 0, len(addrGroup.Data))
	for _, e := range addrGroup.Data {
		if len(e.Domain) > 0 {
			domains = append(domains, Domain(e.Domain))
		} else if len(e.IP) > 0 && len(e.Mask) > 0 {
			prefix, err := NewPrefixFromNetworkMask(e.IP, e.Mask)
			if err != nil {
				return "", nil, nil, fmt.Errorf("API provided invalid IPs in address group: %v / %v", e.IP, e.Mask)
			}
			prefixes = append(prefixes, prefix)
		} else {
			return "", nil, nil, fmt.Errorf("API provided address group entry which is neither domain nor IP")
		}
	}

	return addrGroup.Name, prefixes, domains, nil
}

type getAddrGroupReq struct {
	ID int64 `json:"id"`
}

type addrGrpResult []addrGrp

type addrGrp struct {
	ID   int64             `json:"id,omitempty"`
	Name string            `json:"name,omitempty"`
	Type string            `json:"type,omitempty"`
	Data []addrGrpDataElem `json:"data"`
}

type addrGrpDataElem struct {
	Domain   string `json:"domain,omitempty"`
	IP       string `json:"ip,omitempty"`
	Mask     string `json:"mask,omitempty"`
	RuleType string `json:"rule_type,omitempty"`
}

func buildAddressGroupData(prefixes []Prefix, domains []Domain) []addrGrpDataElem {
	addresses := make([]addrGrpDataElem, 0, len(prefixes)+len(domains))
	for _, domain := range domains {
		addresses = append(addresses, addrGrpDataElem{Domain: string(domain)})
	}
	for _, prefix := range prefixes {
		_, ipNet, _ := net.ParseCIDR(string(prefix))
		ipString, maskString := utils.CidrStringToNetworkMask(ipNet)

		var ruleType string
		if leadingOnes, _ := ipNet.Mask.Size(); leadingOnes == 32 || leadingOnes == 128 {
			ruleType = "exact"
		} else {
			ruleType = "prefix"
		}

		addresses = append(addresses, addrGrpDataElem{IP: ipString, Mask: maskString, RuleType: ruleType})
	}

	return addresses
}

func InsertServiceGroup(c *Client, name string, services []NetworkService) (int64, error) {
	servicesImpl := make([]networkServiceImpl, 0, len(services))
	for _, s := range services {
		servicesImpl = append(servicesImpl, networkServiceImpl{s})
	}
	req := svcGrp{
		Name: name,
		Type: "port_group",
		Data: servicesImpl,
	}

	raw_res, err := c.DoPortal("enterprise/insertObjectGroup", &req)
	if err != nil {
		return 0, err
	}

	var res models.InsertionConfirmation
	err = json.Unmarshal(raw_res, &res)
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func UpdateServiceGroup(c *Client, id int64, services []NetworkService) error {
	servicesImpl := make([]networkServiceImpl, 0, len(services))
	for _, s := range services {
		servicesImpl = append(servicesImpl, networkServiceImpl{s})
	}
	req := svcGrp{
		ID:   id,
		Data: servicesImpl,
	}
	_, err := c.DoPortal("enterprise/updateObjectGroup", req)

	return err
}

type svcGrp struct {
	ID   int64                `json:"id,omitempty"`
	Name string               `json:"name,omitempty"`
	Type string               `json:"type,omitempty"`
	Data []networkServiceImpl `json:"data"`
}

/*
{
	"proto": 1,
	"type": 1,
	"code_low": 2,
	"code_high": 2
},
{
	"proto": 58,
	"type": 3,
	"code_low": 4,
	"code_high": 4
},
{
	"proto": 6,
	"port_low": 1000,
	"port_high": 1000
},
{
	"proto": 17,
	"port_low": 1000,
	"port_high": 2000
}
*/

type networkServiceImpl struct {
	NetworkService
}

func (s networkServiceImpl) MarshalJSON() ([]byte, error) {
	var e bytes.Buffer

	e.WriteString(`{"proto":`)
	e.WriteString(strconv.Itoa(s.Protocol()))

	if portLow, err := s.PortLow(); err == nil {
		portHigh, err := s.PortHigh()
		if err != nil {
			return nil, ErrInvalidNetworkService
		}

		e.WriteString(`,"port_low":`)
		e.WriteString(strconv.Itoa(portLow))
		e.WriteString(`,"port_high":`)
		e.WriteString(strconv.Itoa(portHigh))
	} else if codeLow, err := s.CodeLow(); err == nil {
		codeHigh, err := s.CodeHigh()
		if err != nil {
			return nil, ErrInvalidNetworkService
		}
		icmpType, err := s.Type()
		if err != nil {
			return nil, ErrInvalidNetworkService
		}

		e.WriteString(`,"type":`)
		e.WriteString(strconv.Itoa(icmpType))
		e.WriteString(`,"code_low":`)
		e.WriteString(strconv.Itoa(codeLow))
		e.WriteString(`,"code_high":`)
		e.WriteString(strconv.Itoa(codeHigh))
	}

	e.WriteByte('}')

	return e.Bytes(), nil
}

var ErrNotImplemented = fmt.Errorf("method not implemented")
var ErrInvalidNetworkService = fmt.Errorf("network service bad")

type TransportProtocol int

const (
	ProtoIcmp  TransportProtocol = 1
	ProtoTcp   TransportProtocol = 6
	ProtoUdp   TransportProtocol = 17
	ProtoIcmp6 TransportProtocol = 58
)

type NetworkService interface {
	Protocol() int
	PortLow() (int, error)
	PortHigh() (int, error)
	Type() (int, error)
	CodeLow() (int, error)
	CodeHigh() (int, error)
}

var _ = (NetworkService)((*baseNetworkService)(nil))

type baseNetworkService struct {
	protocol int
}

func (s baseNetworkService) Protocol() int {
	return s.protocol
}
func (s baseNetworkService) PortLow() (int, error) {
	return 0, ErrNotImplemented
}
func (s baseNetworkService) PortHigh() (int, error) {
	return 0, ErrNotImplemented
}
func (s baseNetworkService) Type() (int, error) {
	return 0, ErrNotImplemented
}
func (s baseNetworkService) CodeLow() (int, error) {
	return 0, ErrNotImplemented
}
func (s baseNetworkService) CodeHigh() (int, error) {
	return 0, ErrNotImplemented
}

var _ = (NetworkService)((*portNetworkService)(nil))

type portNetworkService struct {
	baseNetworkService
	portLow  int
	portHigh int
}

func (s portNetworkService) PortLow() (int, error) {
	return s.portLow, nil
}
func (s portNetworkService) PortHigh() (int, error) {
	return s.portHigh, nil
}

func TcpSingle(port int) NetworkService {
	return TcpRange(port, port)
}
func TcpRange(portLow int, portHigh int) NetworkService {
	return portNetworkService{baseNetworkService: baseNetworkService{protocol: 6}, portLow: portLow, portHigh: portHigh}
}
func UdpSingle(port int) NetworkService {
	return UdpRange(port, port)
}
func UdpRange(portLow int, portHigh int) NetworkService {
	return portNetworkService{baseNetworkService: baseNetworkService{protocol: 17}, portLow: portLow, portHigh: portHigh}
}

var _ = (NetworkService)((*icmpNetworkService)(nil))

type icmpNetworkService struct {
	baseNetworkService
	icmpType     int
	icmpCodeLow  int
	icmpCodeHigh int
}

func (s icmpNetworkService) Type() (int, error) {
	return s.icmpType, nil
}
func (s icmpNetworkService) CodeLow() (int, error) {
	return s.icmpCodeLow, nil
}
func (s icmpNetworkService) CodeHigh() (int, error) {
	return s.icmpCodeHigh, nil
}

func IcmpSingle(icmpType int, icmpCode int) NetworkService {
	return IcmpRange(icmpType, icmpCode, icmpCode)
}
func IcmpRange(icmpType int, icmpCodeLow int, icmpCodeHigh int) NetworkService {
	return icmpNetworkService{
		baseNetworkService: baseNetworkService{protocol: 1},
		icmpType:           icmpType,
		icmpCodeLow:        icmpCodeLow,
		icmpCodeHigh:       icmpCodeHigh,
	}
}
func Icmp6Single(icmpType int, icmpCode int) NetworkService {
	return Icmp6Range(icmpType, icmpCode, icmpCode)
}
func Icmp6Range(icmpType int, icmpCodeLow int, icmpCodeHigh int) NetworkService {
	return icmpNetworkService{
		baseNetworkService: baseNetworkService{protocol: 58},
		icmpType:           icmpType,
		icmpCodeLow:        icmpCodeLow,
		icmpCodeHigh:       icmpCodeHigh,
	}
}
