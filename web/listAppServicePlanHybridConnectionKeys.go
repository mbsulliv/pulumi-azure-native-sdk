// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Get the send key name and value of a Hybrid Connection.
// Azure REST API version: 2022-09-01.
func ListAppServicePlanHybridConnectionKeys(ctx *pulumi.Context, args *ListAppServicePlanHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListAppServicePlanHybridConnectionKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAppServicePlanHybridConnectionKeysResult
	err := ctx.Invoke("azure-native:web:listAppServicePlanHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAppServicePlanHybridConnectionKeysArgs struct {
	// Name of the App Service plan.
	Name string `pulumi:"name"`
	// The name of the Service Bus namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the Service Bus relay.
	RelayName string `pulumi:"relayName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Hybrid Connection key contract. This has the send key name and value for a Hybrid Connection.
type ListAppServicePlanHybridConnectionKeysResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// The name of the send key.
	SendKeyName string `pulumi:"sendKeyName"`
	// The value of the send key.
	SendKeyValue string `pulumi:"sendKeyValue"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListAppServicePlanHybridConnectionKeysOutput(ctx *pulumi.Context, args ListAppServicePlanHybridConnectionKeysOutputArgs, opts ...pulumi.InvokeOption) ListAppServicePlanHybridConnectionKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAppServicePlanHybridConnectionKeysResult, error) {
			args := v.(ListAppServicePlanHybridConnectionKeysArgs)
			r, err := ListAppServicePlanHybridConnectionKeys(ctx, &args, opts...)
			var s ListAppServicePlanHybridConnectionKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAppServicePlanHybridConnectionKeysResultOutput)
}

type ListAppServicePlanHybridConnectionKeysOutputArgs struct {
	// Name of the App Service plan.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Service Bus namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the Service Bus relay.
	RelayName pulumi.StringInput `pulumi:"relayName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAppServicePlanHybridConnectionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppServicePlanHybridConnectionKeysArgs)(nil)).Elem()
}

// Hybrid Connection key contract. This has the send key name and value for a Hybrid Connection.
type ListAppServicePlanHybridConnectionKeysResultOutput struct{ *pulumi.OutputState }

func (ListAppServicePlanHybridConnectionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppServicePlanHybridConnectionKeysResult)(nil)).Elem()
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) ToListAppServicePlanHybridConnectionKeysResultOutput() ListAppServicePlanHybridConnectionKeysResultOutput {
	return o
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) ToListAppServicePlanHybridConnectionKeysResultOutputWithContext(ctx context.Context) ListAppServicePlanHybridConnectionKeysResultOutput {
	return o
}

func (o ListAppServicePlanHybridConnectionKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListAppServicePlanHybridConnectionKeysResult] {
	return pulumix.Output[ListAppServicePlanHybridConnectionKeysResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListAppServicePlanHybridConnectionKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListAppServicePlanHybridConnectionKeysResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListAppServicePlanHybridConnectionKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the send key.
func (o ListAppServicePlanHybridConnectionKeysResultOutput) SendKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.SendKeyName }).(pulumi.StringOutput)
}

// The value of the send key.
func (o ListAppServicePlanHybridConnectionKeysResultOutput) SendKeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.SendKeyValue }).(pulumi.StringOutput)
}

// Resource type.
func (o ListAppServicePlanHybridConnectionKeysResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListAppServicePlanHybridConnectionKeysResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAppServicePlanHybridConnectionKeysResultOutput{})
}
