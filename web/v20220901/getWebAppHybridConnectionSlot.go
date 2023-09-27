// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Retrieves a specific Service Bus Hybrid Connection used by this Web App.
func LookupWebAppHybridConnectionSlot(ctx *pulumi.Context, args *LookupWebAppHybridConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppHybridConnectionSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppHybridConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20220901:getWebAppHybridConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppHybridConnectionSlotArgs struct {
	// The name of the web app.
	Name string `pulumi:"name"`
	// The namespace for this hybrid connection.
	NamespaceName string `pulumi:"namespaceName"`
	// The relay name for this hybrid connection.
	RelayName string `pulumi:"relayName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the slot for the web app.
	Slot string `pulumi:"slot"`
}

// Hybrid Connection contract. This is used to configure a Hybrid Connection.
type LookupWebAppHybridConnectionSlotResult struct {
	// The hostname of the endpoint.
	Hostname *string `pulumi:"hostname"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// The port of the endpoint.
	Port *int `pulumi:"port"`
	// The ARM URI to the Service Bus relay.
	RelayArmUri *string `pulumi:"relayArmUri"`
	// The name of the Service Bus relay.
	RelayName *string `pulumi:"relayName"`
	// The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
	SendKeyName *string `pulumi:"sendKeyName"`
	// The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
	// normally, use the POST /listKeys API instead.
	SendKeyValue *string `pulumi:"sendKeyValue"`
	// The name of the Service Bus namespace.
	ServiceBusNamespace *string `pulumi:"serviceBusNamespace"`
	// The suffix for the service bus endpoint. By default this is .servicebus.windows.net
	ServiceBusSuffix *string `pulumi:"serviceBusSuffix"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppHybridConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppHybridConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppHybridConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppHybridConnectionSlotResult, error) {
			args := v.(LookupWebAppHybridConnectionSlotArgs)
			r, err := LookupWebAppHybridConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppHybridConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppHybridConnectionSlotResultOutput)
}

type LookupWebAppHybridConnectionSlotOutputArgs struct {
	// The name of the web app.
	Name pulumi.StringInput `pulumi:"name"`
	// The namespace for this hybrid connection.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The relay name for this hybrid connection.
	RelayName pulumi.StringInput `pulumi:"relayName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the slot for the web app.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppHybridConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHybridConnectionSlotArgs)(nil)).Elem()
}

// Hybrid Connection contract. This is used to configure a Hybrid Connection.
type LookupWebAppHybridConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppHybridConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHybridConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppHybridConnectionSlotResultOutput) ToLookupWebAppHybridConnectionSlotResultOutput() LookupWebAppHybridConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppHybridConnectionSlotResultOutput) ToLookupWebAppHybridConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppHybridConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppHybridConnectionSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppHybridConnectionSlotResult] {
	return pulumix.Output[LookupWebAppHybridConnectionSlotResult]{
		OutputState: o.OutputState,
	}
}

// The hostname of the endpoint.
func (o LookupWebAppHybridConnectionSlotResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppHybridConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppHybridConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppHybridConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// The port of the endpoint.
func (o LookupWebAppHybridConnectionSlotResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The ARM URI to the Service Bus relay.
func (o LookupWebAppHybridConnectionSlotResultOutput) RelayArmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.RelayArmUri }).(pulumi.StringPtrOutput)
}

// The name of the Service Bus relay.
func (o LookupWebAppHybridConnectionSlotResultOutput) RelayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.RelayName }).(pulumi.StringPtrOutput)
}

// The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
func (o LookupWebAppHybridConnectionSlotResultOutput) SendKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.SendKeyName }).(pulumi.StringPtrOutput)
}

// The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
// normally, use the POST /listKeys API instead.
func (o LookupWebAppHybridConnectionSlotResultOutput) SendKeyValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.SendKeyValue }).(pulumi.StringPtrOutput)
}

// The name of the Service Bus namespace.
func (o LookupWebAppHybridConnectionSlotResultOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

// The suffix for the service bus endpoint. By default this is .servicebus.windows.net
func (o LookupWebAppHybridConnectionSlotResultOutput) ServiceBusSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) *string { return v.ServiceBusSuffix }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWebAppHybridConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHybridConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppHybridConnectionSlotResultOutput{})
}
