// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the access keys of the resource.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2018-10-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview.
func ListSignalRKeys(ctx *pulumi.Context, args *ListSignalRKeysArgs, opts ...pulumi.InvokeOption) (*ListSignalRKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSignalRKeysResult
	err := ctx.Invoke("azure-native:signalrservice:listSignalRKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSignalRKeysArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A class represents the access keys of the resource.
type ListSignalRKeysResult struct {
	// Connection string constructed via the primaryKey
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The primary access key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Connection string constructed via the secondaryKey
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The secondary access key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListSignalRKeysOutput(ctx *pulumi.Context, args ListSignalRKeysOutputArgs, opts ...pulumi.InvokeOption) ListSignalRKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSignalRKeysResult, error) {
			args := v.(ListSignalRKeysArgs)
			r, err := ListSignalRKeys(ctx, &args, opts...)
			var s ListSignalRKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSignalRKeysResultOutput)
}

type ListSignalRKeysOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListSignalRKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSignalRKeysArgs)(nil)).Elem()
}

// A class represents the access keys of the resource.
type ListSignalRKeysResultOutput struct{ *pulumi.OutputState }

func (ListSignalRKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSignalRKeysResult)(nil)).Elem()
}

func (o ListSignalRKeysResultOutput) ToListSignalRKeysResultOutput() ListSignalRKeysResultOutput {
	return o
}

func (o ListSignalRKeysResultOutput) ToListSignalRKeysResultOutputWithContext(ctx context.Context) ListSignalRKeysResultOutput {
	return o
}

func (o ListSignalRKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListSignalRKeysResult] {
	return pulumix.Output[ListSignalRKeysResult]{
		OutputState: o.OutputState,
	}
}

// Connection string constructed via the primaryKey
func (o ListSignalRKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

// The primary access key.
func (o ListSignalRKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// Connection string constructed via the secondaryKey
func (o ListSignalRKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

// The secondary access key.
func (o ListSignalRKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSignalRKeysResultOutput{})
}
