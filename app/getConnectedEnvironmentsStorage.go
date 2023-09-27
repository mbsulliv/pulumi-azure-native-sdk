// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get storage for a connectedEnvironment.
// Azure REST API version: 2022-10-01.
func LookupConnectedEnvironmentsStorage(ctx *pulumi.Context, args *LookupConnectedEnvironmentsStorageArgs, opts ...pulumi.InvokeOption) (*LookupConnectedEnvironmentsStorageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectedEnvironmentsStorageResult
	err := ctx.Invoke("azure-native:app:getConnectedEnvironmentsStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectedEnvironmentsStorageArgs struct {
	// Name of the Environment.
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the storage.
	StorageName string `pulumi:"storageName"`
}

// Storage resource for connectedEnvironment.
type LookupConnectedEnvironmentsStorageResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Storage properties
	Properties ConnectedEnvironmentStorageResponseProperties `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupConnectedEnvironmentsStorageOutput(ctx *pulumi.Context, args LookupConnectedEnvironmentsStorageOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedEnvironmentsStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedEnvironmentsStorageResult, error) {
			args := v.(LookupConnectedEnvironmentsStorageArgs)
			r, err := LookupConnectedEnvironmentsStorage(ctx, &args, opts...)
			var s LookupConnectedEnvironmentsStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedEnvironmentsStorageResultOutput)
}

type LookupConnectedEnvironmentsStorageOutputArgs struct {
	// Name of the Environment.
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the storage.
	StorageName pulumi.StringInput `pulumi:"storageName"`
}

func (LookupConnectedEnvironmentsStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsStorageArgs)(nil)).Elem()
}

// Storage resource for connectedEnvironment.
type LookupConnectedEnvironmentsStorageResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedEnvironmentsStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsStorageResult)(nil)).Elem()
}

func (o LookupConnectedEnvironmentsStorageResultOutput) ToLookupConnectedEnvironmentsStorageResultOutput() LookupConnectedEnvironmentsStorageResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsStorageResultOutput) ToLookupConnectedEnvironmentsStorageResultOutputWithContext(ctx context.Context) LookupConnectedEnvironmentsStorageResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsStorageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectedEnvironmentsStorageResult] {
	return pulumix.Output[LookupConnectedEnvironmentsStorageResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConnectedEnvironmentsStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupConnectedEnvironmentsStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Storage properties
func (o LookupConnectedEnvironmentsStorageResultOutput) Properties() ConnectedEnvironmentStorageResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) ConnectedEnvironmentStorageResponseProperties {
		return v.Properties
	}).(ConnectedEnvironmentStorageResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConnectedEnvironmentsStorageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConnectedEnvironmentsStorageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedEnvironmentsStorageResultOutput{})
}
