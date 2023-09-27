package resources

import (
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-xyz/provider/pkg/provider/api"
	"github.com/pulumi/pulumi-xyz/provider/pkg/provider/config"
)

var _ = (infer.CustomResource[AddressGroupInputs, AddressGroupOutputs])((*AddressGroup)(nil))
var _ = (infer.CustomDelete[AddressGroupOutputs])((*AddressGroup)(nil))
var _ = (infer.CustomDiff[AddressGroupInputs, AddressGroupOutputs])((*AddressGroup)(nil))
var _ = (infer.CustomUpdate[AddressGroupInputs, AddressGroupOutputs])((*AddressGroup)(nil))
var _ = (infer.CustomRead[AddressGroupInputs, AddressGroupOutputs])((*AddressGroup)(nil))

type AddressGroup struct{}

type AddressGroupInputs struct {
	Prefixes []api.Prefix `pulumi:"prefixes"`
	Domains  []api.Domain `pulumi:"domains"`
}

type AddressGroupOutputs struct {
	AddressGroupInputs
	AddressGroupId int64 `pulumi:"addressGroupId"`
}

func (ag *AddressGroup) Create(ctx p.Context, name string, input AddressGroupInputs, preview bool) (string, AddressGroupOutputs, error) {
	state := AddressGroupOutputs{AddressGroupInputs: input}

	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[config.XyzConfig](ctx)

	objectId, err := api.InsertAddressGroup(config.Client(), name, input.Prefixes, input.Domains)
	if err != nil {
		return name, state, err
	}

	state.AddressGroupId = objectId

	return name, state, nil
}

func (ag *AddressGroup) Delete(ctx p.Context, id string, props AddressGroupOutputs) error {
	config := infer.GetConfig[config.XyzConfig](ctx)

	return api.DeleteObjectGroup(config.Client(), props.AddressGroupId)
}

func (ag *AddressGroup) Update(ctx p.Context, id string, olds AddressGroupOutputs, news AddressGroupInputs, preview bool) (AddressGroupOutputs, error) {
	state := AddressGroupOutputs{AddressGroupInputs: news, AddressGroupId: olds.AddressGroupId}

	if preview {
		return state, nil
	}

	config := infer.GetConfig[config.XyzConfig](ctx)

	if err := api.UpdateAddressGroup(config.Client(), olds.AddressGroupId, id, news.Prefixes, news.Domains); err != nil {
		return olds, err
	}

	return state, nil
}

func (ag *AddressGroup) Diff(ctx p.Context, id string, olds AddressGroupOutputs, news AddressGroupInputs) (p.DiffResponse, error) {
	diffs := make(map[string]p.PropertyDiff)

	// the output has ugly nulls because there's no real set type in the schema
	// map values are indices into old/new in case we can fix this somehow
	// maybe PropertyDiff.InputDiff?

	// TODO: refactor to use diffArrays helper

	olds_prefix_set := make(map[api.Prefix]int)
	for i, prefix := range olds.Prefixes {
		olds_prefix_set[prefix] = i
	}
	news_prefix_set := make(map[api.Prefix]int)
	for i, prefix := range news.Prefixes {
		if _, in_old := olds_prefix_set[prefix]; !in_old {
			diffs["prefixes[\""+string(prefix)+"\"]"] = p.PropertyDiff{Kind: p.Add}
		}

		news_prefix_set[prefix] = i
	}

	for prefix := range olds_prefix_set {
		if _, in_new := news_prefix_set[prefix]; !in_new {
			diffs["prefixes[\""+string(prefix)+"\"]"] = p.PropertyDiff{Kind: p.Delete}
		}
	}

	olds_dom_set := make(map[api.Domain]int)
	for i, domain := range olds.Domains {
		olds_dom_set[domain] = i
	}
	news_dom_set := make(map[api.Domain]int)
	for i, domain := range news.Domains {
		if _, in_old := olds_dom_set[domain]; !in_old {
			diffs["domains[\""+string(domain)+"\"]"] = p.PropertyDiff{Kind: p.Add}
		}

		news_dom_set[domain] = i
	}

	for domain := range olds_dom_set {
		if _, in_new := news_dom_set[domain]; !in_new {
			diffs["domains[\""+string(domain)+"\"]"] = p.PropertyDiff{Kind: p.Delete}
		}
	}

	resp := p.DiffResponse{DeleteBeforeReplace: false, HasChanges: len(diffs) > 0, DetailedDiff: diffs}
	return resp, nil
}

func (ag *AddressGroup) Read(ctx p.Context, id string, inputs AddressGroupInputs, state AddressGroupOutputs) (
	string, AddressGroupInputs, AddressGroupOutputs, error) {
	config := infer.GetConfig[config.XyzConfig](ctx)

	_, prefixes, domains, err := api.GetAddressGroup(config.Client(), state.AddressGroupId)
	if err != nil {
		return id, inputs, state, err
	}

	trueInputs := AddressGroupInputs{Prefixes: prefixes, Domains: domains}
	trueOutputs := AddressGroupOutputs{AddressGroupInputs: trueInputs, AddressGroupId: state.AddressGroupId}

	return id, trueInputs, trueOutputs, nil
}

type ServiceGroupStringify interface {
	StringifyServiceGroup() string
}

type ServiceGroupTcp struct {
	PortLow  int `pulumi:"portStart"`
	PortHigh int `pulumi:"portEnd"`
}

func (sg ServiceGroupTcp) StringifyServiceGroup() string {
	return fmt.Sprint(sg.PortLow, "-", sg.PortHigh)
}

type ServiceGroupUdp struct {
	PortLow  int `pulumi:"portStart"`
	PortHigh int `pulumi:"portEnd"`
}

func (sg ServiceGroupUdp) StringifyServiceGroup() string {
	return fmt.Sprint(sg.PortLow, "-", sg.PortHigh)
}

type ServiceGroupIcmp struct {
	Type     int `pulumi:"icmpType"`
	CodeLow  int `pulumi:"codeLow"`
	CodeHigh int `pulumi:"codeHigh"`
}

func (sg ServiceGroupIcmp) StringifyServiceGroup() string {
	return fmt.Sprint(sg.Type, ", ", sg.CodeLow, "-", sg.CodeHigh)
}

type ServiceGroup struct{}
type ServiceGroupInputs struct {
	Tcp   []ServiceGroupTcp  `pulumi:"tcp,optional"`
	Udp   []ServiceGroupUdp  `pulumi:"udp,optional"`
	Icmp  []ServiceGroupIcmp `pulumi:"icmp,optional"`
	Icmp6 []ServiceGroupIcmp `pulumi:"icmp6,optional"`
}
type ServiceGroupOutputs struct {
	ServiceGroupInputs
	ServiceGroupId int64 `pulumi:"serviceGroupId"`
}

var _ = (infer.CustomResource[ServiceGroupInputs, ServiceGroupOutputs])((*ServiceGroup)(nil))
var _ = (infer.CustomDelete[ServiceGroupOutputs])((*ServiceGroup)(nil))
var _ = (infer.CustomUpdate[ServiceGroupInputs, ServiceGroupOutputs])((*ServiceGroup)(nil))

/*
var _ = (infer.CustomDiff[ServiceGroupInputs, ServiceGroupOutputs])((*ServiceGroup)(nil))
var _ = (infer.CustomRead[ServiceGroupInputs, ServiceGroupOutputs])((*ServiceGroup)(nil))
*/

// TODO: Should this be done by implementing NetworkService on the Pulumi types?

func (sg *ServiceGroup) Create(ctx p.Context, name string, input ServiceGroupInputs, preview bool) (string, ServiceGroupOutputs, error) {
	state := ServiceGroupOutputs{ServiceGroupInputs: input}

	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[config.XyzConfig](ctx)

	inputsInterfaces := make([]api.NetworkService, 0, len(input.Tcp)+len(input.Udp)+len(input.Icmp)+len(input.Icmp6))
	for _, i := range input.Tcp {
		inputsInterfaces = append(inputsInterfaces, api.TcpRange(i.PortLow, i.PortHigh))
	}
	for _, i := range input.Udp {
		inputsInterfaces = append(inputsInterfaces, api.UdpRange(i.PortLow, i.PortHigh))
	}
	for _, i := range input.Icmp {
		inputsInterfaces = append(inputsInterfaces, api.IcmpRange(i.Type, i.CodeLow, i.CodeHigh))
	}
	for _, i := range input.Icmp6 {
		inputsInterfaces = append(inputsInterfaces, api.Icmp6Range(i.Type, i.CodeLow, i.CodeHigh))
	}

	objectId, err := api.InsertServiceGroup(config.Client(), name, inputsInterfaces)
	if err != nil {
		return name, state, err
	}

	state.ServiceGroupId = objectId

	return name, state, nil
}

func (sg *ServiceGroup) Delete(ctx p.Context, id string, props ServiceGroupOutputs) error {
	config := infer.GetConfig[config.XyzConfig](ctx)

	return api.DeleteObjectGroup(config.Client(), props.ServiceGroupId)
}

func (sg *ServiceGroup) Update(ctx p.Context, id string, olds ServiceGroupOutputs, news ServiceGroupInputs, preview bool) (ServiceGroupOutputs, error) {
	state := ServiceGroupOutputs{ServiceGroupInputs: news, ServiceGroupId: olds.ServiceGroupId}

	if preview {
		return state, nil
	}

	config := infer.GetConfig[config.XyzConfig](ctx)

	inputsInterfaces := make([]api.NetworkService, 0, len(news.Tcp)+len(news.Udp)+len(news.Icmp)+len(news.Icmp6))
	for _, i := range news.Tcp {
		inputsInterfaces = append(inputsInterfaces, api.TcpRange(i.PortLow, i.PortHigh))
	}
	for _, i := range news.Udp {
		inputsInterfaces = append(inputsInterfaces, api.UdpRange(i.PortLow, i.PortHigh))
	}
	for _, i := range news.Icmp {
		inputsInterfaces = append(inputsInterfaces, api.IcmpRange(i.Type, i.CodeLow, i.CodeHigh))
	}
	for _, i := range news.Icmp6 {
		inputsInterfaces = append(inputsInterfaces, api.Icmp6Range(i.Type, i.CodeLow, i.CodeHigh))
	}

	if err := api.UpdateServiceGroup(config.Client(), olds.ServiceGroupId, inputsInterfaces); err != nil {
		return olds, err
	}

	return state, nil
}

func (sg *ServiceGroup) Diff(ctx p.Context, id string, olds ServiceGroupOutputs, news ServiceGroupInputs) (p.DiffResponse, error) {
	tcpInserted, tcpDeleted := diffArrays(olds.Tcp, news.Tcp)
	udpInserted, udpDeleted := diffArrays(olds.Udp, news.Udp)
	icmpInserted, icmpDeleted := diffArrays(olds.Icmp, news.Icmp)
	icmp6Inserted, icmp6Deleted := diffArrays(olds.Icmp6, news.Icmp6)

	diffCount := len(tcpInserted) + len(tcpDeleted) + len(udpInserted) + len(udpDeleted) + len(icmpInserted) + len(icmpDeleted) + len(icmp6Inserted) + len(icmp6Deleted)

	if diffCount == 0 {
		return p.DiffResponse{DeleteBeforeReplace: false, HasChanges: false}, nil
	}

	diffs := make(map[string]p.PropertyDiff, diffCount)

	diffs = buildDetailedDiffs(diffs, "tcp", tcpInserted, tcpDeleted)
	diffs = buildDetailedDiffs(diffs, "udp", udpInserted, udpDeleted)
	diffs = buildDetailedDiffs(diffs, "icmp", icmpInserted, icmpDeleted)
	diffs = buildDetailedDiffs(diffs, "icmp6", icmp6Inserted, icmp6Deleted)

	resp := p.DiffResponse{DeleteBeforeReplace: false, HasChanges: true, DetailedDiff: diffs}
	return resp, nil
}

func diffArrays[I comparable](olds []I, news []I) (inserted []I, deleted []I) {
	inserted = make([]I, 0, 4)
	deleted = make([]I, 0, 4)

	oldsSet := setFromArray(olds)
	newsSet := setFromArray(news)

	for _, newItem := range news {
		if _, ok := oldsSet[newItem]; !ok {
			inserted = append(inserted, newItem)
		}
	}

	for _, oldItem := range olds {
		if _, ok := newsSet[oldItem]; !ok {
			deleted = append(deleted, oldItem)
		}
	}

	return inserted, deleted
}

func setFromArray[I comparable](a []I) map[I]struct{} {
	s := make(map[I]struct{}, len(a))

	for _, i := range a {
		s[i] = struct{}{}
	}

	return s
}

func buildDetailedDiffs[I ServiceGroupStringify](diffs map[string]p.PropertyDiff, propName string, insertions []I, deletions []I) map[string]p.PropertyDiff {
	for _, i := range insertions {
		diffs[propName+`["`+i.StringifyServiceGroup()+`"]`] = p.PropertyDiff{Kind: p.Add}
	}
	for _, d := range deletions {
		diffs[propName+`["`+d.StringifyServiceGroup()+`"]`] = p.PropertyDiff{Kind: p.Delete}
	}

	return diffs
}
