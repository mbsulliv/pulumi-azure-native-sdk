// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get storage for a managedEnvironment.
// Azure REST API version: 2022-10-01.
//
// Other available API versions: 2022-01-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview.
func LookupManagedEnvironmentsStorage(ctx *pulumi.Context, args *LookupManagedEnvironmentsStorageArgs, opts ...pulumi.InvokeOption) (*LookupManagedEnvironmentsStorageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedEnvironmentsStorageResult
	err := ctx.Invoke("azure-native:app:getManagedEnvironmentsStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedEnvironmentsStorageArgs struct {
	// Name of the Environment.
	EnvironmentName string `pulumi:"environmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the storage.
	StorageName string `pulumi:"storageName"`
}

// Storage resource for managedEnvironment.
type LookupManagedEnvironmentsStorageResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Storage properties
	Properties ManagedEnvironmentStorageResponseProperties `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupManagedEnvironmentsStorageOutput(ctx *pulumi.Context, args LookupManagedEnvironmentsStorageOutputArgs, opts ...pulumi.InvokeOption) LookupManagedEnvironmentsStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedEnvironmentsStorageResult, error) {
			args := v.(LookupManagedEnvironmentsStorageArgs)
			r, err := LookupManagedEnvironmentsStorage(ctx, &args, opts...)
			var s LookupManagedEnvironmentsStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedEnvironmentsStorageResultOutput)
}

type LookupManagedEnvironmentsStorageOutputArgs struct {
	// Name of the Environment.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the storage.
	StorageName pulumi.StringInput `pulumi:"storageName"`
}

func (LookupManagedEnvironmentsStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentsStorageArgs)(nil)).Elem()
}

// Storage resource for managedEnvironment.
type LookupManagedEnvironmentsStorageResultOutput struct{ *pulumi.OutputState }

func (LookupManagedEnvironmentsStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentsStorageResult)(nil)).Elem()
}

func (o LookupManagedEnvironmentsStorageResultOutput) ToLookupManagedEnvironmentsStorageResultOutput() LookupManagedEnvironmentsStorageResultOutput {
	return o
}

func (o LookupManagedEnvironmentsStorageResultOutput) ToLookupManagedEnvironmentsStorageResultOutputWithContext(ctx context.Context) LookupManagedEnvironmentsStorageResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedEnvironmentsStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedEnvironmentsStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Storage properties
func (o LookupManagedEnvironmentsStorageResultOutput) Properties() ManagedEnvironmentStorageResponsePropertiesOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) ManagedEnvironmentStorageResponseProperties {
		return v.Properties
	}).(ManagedEnvironmentStorageResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupManagedEnvironmentsStorageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedEnvironmentsStorageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedEnvironmentsStorageResultOutput{})
}
