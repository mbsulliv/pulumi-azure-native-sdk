// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List test keys for a Service.
func ListServiceTestKeys(ctx *pulumi.Context, args *ListServiceTestKeysArgs, opts ...pulumi.InvokeOption) (*ListServiceTestKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListServiceTestKeysResult
	err := ctx.Invoke("azure-native:appplatform/v20230501preview:listServiceTestKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServiceTestKeysArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Test keys payload
type ListServiceTestKeysResult struct {
	// Indicates whether the test endpoint feature enabled or not
	Enabled *bool `pulumi:"enabled"`
	// Primary key
	PrimaryKey *string `pulumi:"primaryKey"`
	// Primary test endpoint
	PrimaryTestEndpoint *string `pulumi:"primaryTestEndpoint"`
	// Secondary key
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Secondary test endpoint
	SecondaryTestEndpoint *string `pulumi:"secondaryTestEndpoint"`
}

func ListServiceTestKeysOutput(ctx *pulumi.Context, args ListServiceTestKeysOutputArgs, opts ...pulumi.InvokeOption) ListServiceTestKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServiceTestKeysResult, error) {
			args := v.(ListServiceTestKeysArgs)
			r, err := ListServiceTestKeys(ctx, &args, opts...)
			var s ListServiceTestKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServiceTestKeysResultOutput)
}

type ListServiceTestKeysOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListServiceTestKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceTestKeysArgs)(nil)).Elem()
}

// Test keys payload
type ListServiceTestKeysResultOutput struct{ *pulumi.OutputState }

func (ListServiceTestKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceTestKeysResult)(nil)).Elem()
}

func (o ListServiceTestKeysResultOutput) ToListServiceTestKeysResultOutput() ListServiceTestKeysResultOutput {
	return o
}

func (o ListServiceTestKeysResultOutput) ToListServiceTestKeysResultOutputWithContext(ctx context.Context) ListServiceTestKeysResultOutput {
	return o
}

// Indicates whether the test endpoint feature enabled or not
func (o ListServiceTestKeysResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Primary key
func (o ListServiceTestKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// Primary test endpoint
func (o ListServiceTestKeysResultOutput) PrimaryTestEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.PrimaryTestEndpoint }).(pulumi.StringPtrOutput)
}

// Secondary key
func (o ListServiceTestKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

// Secondary test endpoint
func (o ListServiceTestKeysResultOutput) SecondaryTestEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceTestKeysResult) *string { return v.SecondaryTestEndpoint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServiceTestKeysResultOutput{})
}
