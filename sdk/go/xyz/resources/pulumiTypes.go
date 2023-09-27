// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"internal"
)

var _ = internal.GetEnvOrDefault

type ServiceGroupIcmp struct {
	CodeHigh int `pulumi:"codeHigh"`
	CodeLow  int `pulumi:"codeLow"`
	IcmpType int `pulumi:"icmpType"`
}

// ServiceGroupIcmpInput is an input type that accepts ServiceGroupIcmpArgs and ServiceGroupIcmpOutput values.
// You can construct a concrete instance of `ServiceGroupIcmpInput` via:
//
//	ServiceGroupIcmpArgs{...}
type ServiceGroupIcmpInput interface {
	pulumi.Input

	ToServiceGroupIcmpOutput() ServiceGroupIcmpOutput
	ToServiceGroupIcmpOutputWithContext(context.Context) ServiceGroupIcmpOutput
}

type ServiceGroupIcmpArgs struct {
	CodeHigh pulumi.IntInput `pulumi:"codeHigh"`
	CodeLow  pulumi.IntInput `pulumi:"codeLow"`
	IcmpType pulumi.IntInput `pulumi:"icmpType"`
}

func (ServiceGroupIcmpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGroupIcmp)(nil)).Elem()
}

func (i ServiceGroupIcmpArgs) ToServiceGroupIcmpOutput() ServiceGroupIcmpOutput {
	return i.ToServiceGroupIcmpOutputWithContext(context.Background())
}

func (i ServiceGroupIcmpArgs) ToServiceGroupIcmpOutputWithContext(ctx context.Context) ServiceGroupIcmpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupIcmpOutput)
}

func (i ServiceGroupIcmpArgs) ToOutput(ctx context.Context) pulumix.Output[ServiceGroupIcmp] {
	return pulumix.Output[ServiceGroupIcmp]{
		OutputState: i.ToServiceGroupIcmpOutputWithContext(ctx).OutputState,
	}
}

// ServiceGroupIcmpArrayInput is an input type that accepts ServiceGroupIcmpArray and ServiceGroupIcmpArrayOutput values.
// You can construct a concrete instance of `ServiceGroupIcmpArrayInput` via:
//
//	ServiceGroupIcmpArray{ ServiceGroupIcmpArgs{...} }
type ServiceGroupIcmpArrayInput interface {
	pulumi.Input

	ToServiceGroupIcmpArrayOutput() ServiceGroupIcmpArrayOutput
	ToServiceGroupIcmpArrayOutputWithContext(context.Context) ServiceGroupIcmpArrayOutput
}

type ServiceGroupIcmpArray []ServiceGroupIcmpInput

func (ServiceGroupIcmpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceGroupIcmp)(nil)).Elem()
}

func (i ServiceGroupIcmpArray) ToServiceGroupIcmpArrayOutput() ServiceGroupIcmpArrayOutput {
	return i.ToServiceGroupIcmpArrayOutputWithContext(context.Background())
}

func (i ServiceGroupIcmpArray) ToServiceGroupIcmpArrayOutputWithContext(ctx context.Context) ServiceGroupIcmpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupIcmpArrayOutput)
}

func (i ServiceGroupIcmpArray) ToOutput(ctx context.Context) pulumix.Output[[]ServiceGroupIcmp] {
	return pulumix.Output[[]ServiceGroupIcmp]{
		OutputState: i.ToServiceGroupIcmpArrayOutputWithContext(ctx).OutputState,
	}
}

type ServiceGroupIcmpOutput struct{ *pulumi.OutputState }

func (ServiceGroupIcmpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGroupIcmp)(nil)).Elem()
}

func (o ServiceGroupIcmpOutput) ToServiceGroupIcmpOutput() ServiceGroupIcmpOutput {
	return o
}

func (o ServiceGroupIcmpOutput) ToServiceGroupIcmpOutputWithContext(ctx context.Context) ServiceGroupIcmpOutput {
	return o
}

func (o ServiceGroupIcmpOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceGroupIcmp] {
	return pulumix.Output[ServiceGroupIcmp]{
		OutputState: o.OutputState,
	}
}

func (o ServiceGroupIcmpOutput) CodeHigh() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceGroupIcmp) int { return v.CodeHigh }).(pulumi.IntOutput)
}

func (o ServiceGroupIcmpOutput) CodeLow() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceGroupIcmp) int { return v.CodeLow }).(pulumi.IntOutput)
}

func (o ServiceGroupIcmpOutput) IcmpType() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceGroupIcmp) int { return v.IcmpType }).(pulumi.IntOutput)
}

type ServiceGroupIcmpArrayOutput struct{ *pulumi.OutputState }

func (ServiceGroupIcmpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceGroupIcmp)(nil)).Elem()
}

func (o ServiceGroupIcmpArrayOutput) ToServiceGroupIcmpArrayOutput() ServiceGroupIcmpArrayOutput {
	return o
}

func (o ServiceGroupIcmpArrayOutput) ToServiceGroupIcmpArrayOutputWithContext(ctx context.Context) ServiceGroupIcmpArrayOutput {
	return o
}

func (o ServiceGroupIcmpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ServiceGroupIcmp] {
	return pulumix.Output[[]ServiceGroupIcmp]{
		OutputState: o.OutputState,
	}
}

func (o ServiceGroupIcmpArrayOutput) Index(i pulumi.IntInput) ServiceGroupIcmpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceGroupIcmp {
		return vs[0].([]ServiceGroupIcmp)[vs[1].(int)]
	}).(ServiceGroupIcmpOutput)
}

type ServiceGroupTcp struct {
	PortEnd   int `pulumi:"portEnd"`
	PortStart int `pulumi:"portStart"`
}

// ServiceGroupTcpInput is an input type that accepts ServiceGroupTcpArgs and ServiceGroupTcpOutput values.
// You can construct a concrete instance of `ServiceGroupTcpInput` via:
//
//	ServiceGroupTcpArgs{...}
type ServiceGroupTcpInput interface {
	pulumi.Input

	ToServiceGroupTcpOutput() ServiceGroupTcpOutput
	ToServiceGroupTcpOutputWithContext(context.Context) ServiceGroupTcpOutput
}

type ServiceGroupTcpArgs struct {
	PortEnd   pulumi.IntInput `pulumi:"portEnd"`
	PortStart pulumi.IntInput `pulumi:"portStart"`
}

func (ServiceGroupTcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGroupTcp)(nil)).Elem()
}

func (i ServiceGroupTcpArgs) ToServiceGroupTcpOutput() ServiceGroupTcpOutput {
	return i.ToServiceGroupTcpOutputWithContext(context.Background())
}

func (i ServiceGroupTcpArgs) ToServiceGroupTcpOutputWithContext(ctx context.Context) ServiceGroupTcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupTcpOutput)
}

func (i ServiceGroupTcpArgs) ToOutput(ctx context.Context) pulumix.Output[ServiceGroupTcp] {
	return pulumix.Output[ServiceGroupTcp]{
		OutputState: i.ToServiceGroupTcpOutputWithContext(ctx).OutputState,
	}
}

// ServiceGroupTcpArrayInput is an input type that accepts ServiceGroupTcpArray and ServiceGroupTcpArrayOutput values.
// You can construct a concrete instance of `ServiceGroupTcpArrayInput` via:
//
//	ServiceGroupTcpArray{ ServiceGroupTcpArgs{...} }
type ServiceGroupTcpArrayInput interface {
	pulumi.Input

	ToServiceGroupTcpArrayOutput() ServiceGroupTcpArrayOutput
	ToServiceGroupTcpArrayOutputWithContext(context.Context) ServiceGroupTcpArrayOutput
}

type ServiceGroupTcpArray []ServiceGroupTcpInput

func (ServiceGroupTcpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceGroupTcp)(nil)).Elem()
}

func (i ServiceGroupTcpArray) ToServiceGroupTcpArrayOutput() ServiceGroupTcpArrayOutput {
	return i.ToServiceGroupTcpArrayOutputWithContext(context.Background())
}

func (i ServiceGroupTcpArray) ToServiceGroupTcpArrayOutputWithContext(ctx context.Context) ServiceGroupTcpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupTcpArrayOutput)
}

func (i ServiceGroupTcpArray) ToOutput(ctx context.Context) pulumix.Output[[]ServiceGroupTcp] {
	return pulumix.Output[[]ServiceGroupTcp]{
		OutputState: i.ToServiceGroupTcpArrayOutputWithContext(ctx).OutputState,
	}
}

type ServiceGroupTcpOutput struct{ *pulumi.OutputState }

func (ServiceGroupTcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGroupTcp)(nil)).Elem()
}

func (o ServiceGroupTcpOutput) ToServiceGroupTcpOutput() ServiceGroupTcpOutput {
	return o
}

func (o ServiceGroupTcpOutput) ToServiceGroupTcpOutputWithContext(ctx context.Context) ServiceGroupTcpOutput {
	return o
}

func (o ServiceGroupTcpOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceGroupTcp] {
	return pulumix.Output[ServiceGroupTcp]{
		OutputState: o.OutputState,
	}
}

func (o ServiceGroupTcpOutput) PortEnd() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceGroupTcp) int { return v.PortEnd }).(pulumi.IntOutput)
}

func (o ServiceGroupTcpOutput) PortStart() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceGroupTcp) int { return v.PortStart }).(pulumi.IntOutput)
}

type ServiceGroupTcpArrayOutput struct{ *pulumi.OutputState }

func (ServiceGroupTcpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceGroupTcp)(nil)).Elem()
}

func (o ServiceGroupTcpArrayOutput) ToServiceGroupTcpArrayOutput() ServiceGroupTcpArrayOutput {
	return o
}

func (o ServiceGroupTcpArrayOutput) ToServiceGroupTcpArrayOutputWithContext(ctx context.Context) ServiceGroupTcpArrayOutput {
	return o
}

func (o ServiceGroupTcpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ServiceGroupTcp] {
	return pulumix.Output[[]ServiceGroupTcp]{
		OutputState: o.OutputState,
	}
}

func (o ServiceGroupTcpArrayOutput) Index(i pulumi.IntInput) ServiceGroupTcpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceGroupTcp {
		return vs[0].([]ServiceGroupTcp)[vs[1].(int)]
	}).(ServiceGroupTcpOutput)
}

type ServiceGroupUdp struct {
	PortEnd   int `pulumi:"portEnd"`
	PortStart int `pulumi:"portStart"`
}

// ServiceGroupUdpInput is an input type that accepts ServiceGroupUdpArgs and ServiceGroupUdpOutput values.
// You can construct a concrete instance of `ServiceGroupUdpInput` via:
//
//	ServiceGroupUdpArgs{...}
type ServiceGroupUdpInput interface {
	pulumi.Input

	ToServiceGroupUdpOutput() ServiceGroupUdpOutput
	ToServiceGroupUdpOutputWithContext(context.Context) ServiceGroupUdpOutput
}

type ServiceGroupUdpArgs struct {
	PortEnd   pulumi.IntInput `pulumi:"portEnd"`
	PortStart pulumi.IntInput `pulumi:"portStart"`
}

func (ServiceGroupUdpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGroupUdp)(nil)).Elem()
}

func (i ServiceGroupUdpArgs) ToServiceGroupUdpOutput() ServiceGroupUdpOutput {
	return i.ToServiceGroupUdpOutputWithContext(context.Background())
}

func (i ServiceGroupUdpArgs) ToServiceGroupUdpOutputWithContext(ctx context.Context) ServiceGroupUdpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupUdpOutput)
}

func (i ServiceGroupUdpArgs) ToOutput(ctx context.Context) pulumix.Output[ServiceGroupUdp] {
	return pulumix.Output[ServiceGroupUdp]{
		OutputState: i.ToServiceGroupUdpOutputWithContext(ctx).OutputState,
	}
}

// ServiceGroupUdpArrayInput is an input type that accepts ServiceGroupUdpArray and ServiceGroupUdpArrayOutput values.
// You can construct a concrete instance of `ServiceGroupUdpArrayInput` via:
//
//	ServiceGroupUdpArray{ ServiceGroupUdpArgs{...} }
type ServiceGroupUdpArrayInput interface {
	pulumi.Input

	ToServiceGroupUdpArrayOutput() ServiceGroupUdpArrayOutput
	ToServiceGroupUdpArrayOutputWithContext(context.Context) ServiceGroupUdpArrayOutput
}

type ServiceGroupUdpArray []ServiceGroupUdpInput

func (ServiceGroupUdpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceGroupUdp)(nil)).Elem()
}

func (i ServiceGroupUdpArray) ToServiceGroupUdpArrayOutput() ServiceGroupUdpArrayOutput {
	return i.ToServiceGroupUdpArrayOutputWithContext(context.Background())
}

func (i ServiceGroupUdpArray) ToServiceGroupUdpArrayOutputWithContext(ctx context.Context) ServiceGroupUdpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupUdpArrayOutput)
}

func (i ServiceGroupUdpArray) ToOutput(ctx context.Context) pulumix.Output[[]ServiceGroupUdp] {
	return pulumix.Output[[]ServiceGroupUdp]{
		OutputState: i.ToServiceGroupUdpArrayOutputWithContext(ctx).OutputState,
	}
}

type ServiceGroupUdpOutput struct{ *pulumi.OutputState }

func (ServiceGroupUdpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGroupUdp)(nil)).Elem()
}

func (o ServiceGroupUdpOutput) ToServiceGroupUdpOutput() ServiceGroupUdpOutput {
	return o
}

func (o ServiceGroupUdpOutput) ToServiceGroupUdpOutputWithContext(ctx context.Context) ServiceGroupUdpOutput {
	return o
}

func (o ServiceGroupUdpOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceGroupUdp] {
	return pulumix.Output[ServiceGroupUdp]{
		OutputState: o.OutputState,
	}
}

func (o ServiceGroupUdpOutput) PortEnd() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceGroupUdp) int { return v.PortEnd }).(pulumi.IntOutput)
}

func (o ServiceGroupUdpOutput) PortStart() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceGroupUdp) int { return v.PortStart }).(pulumi.IntOutput)
}

type ServiceGroupUdpArrayOutput struct{ *pulumi.OutputState }

func (ServiceGroupUdpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceGroupUdp)(nil)).Elem()
}

func (o ServiceGroupUdpArrayOutput) ToServiceGroupUdpArrayOutput() ServiceGroupUdpArrayOutput {
	return o
}

func (o ServiceGroupUdpArrayOutput) ToServiceGroupUdpArrayOutputWithContext(ctx context.Context) ServiceGroupUdpArrayOutput {
	return o
}

func (o ServiceGroupUdpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ServiceGroupUdp] {
	return pulumix.Output[[]ServiceGroupUdp]{
		OutputState: o.OutputState,
	}
}

func (o ServiceGroupUdpArrayOutput) Index(i pulumi.IntInput) ServiceGroupUdpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceGroupUdp {
		return vs[0].([]ServiceGroupUdp)[vs[1].(int)]
	}).(ServiceGroupUdpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupIcmpInput)(nil)).Elem(), ServiceGroupIcmpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupIcmpArrayInput)(nil)).Elem(), ServiceGroupIcmpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupTcpInput)(nil)).Elem(), ServiceGroupTcpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupTcpArrayInput)(nil)).Elem(), ServiceGroupTcpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupUdpInput)(nil)).Elem(), ServiceGroupUdpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupUdpArrayInput)(nil)).Elem(), ServiceGroupUdpArray{})
	pulumi.RegisterOutputType(ServiceGroupIcmpOutput{})
	pulumi.RegisterOutputType(ServiceGroupIcmpArrayOutput{})
	pulumi.RegisterOutputType(ServiceGroupTcpOutput{})
	pulumi.RegisterOutputType(ServiceGroupTcpArrayOutput{})
	pulumi.RegisterOutputType(ServiceGroupUdpOutput{})
	pulumi.RegisterOutputType(ServiceGroupUdpArrayOutput{})
}
