// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the access keys of the resource.
func ListWebPubSubKeys(ctx *pulumi.Context, args *ListWebPubSubKeysArgs, opts ...pulumi.InvokeOption) (*ListWebPubSubKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebPubSubKeysResult
	err := ctx.Invoke("azure-native:webpubsub/v20210401preview:listWebPubSubKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebPubSubKeysArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A class represents the access keys of the resource.
type ListWebPubSubKeysResult struct {
	// Connection string constructed via the primaryKey
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The primary access key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Connection string constructed via the secondaryKey
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The secondary access key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListWebPubSubKeysOutput(ctx *pulumi.Context, args ListWebPubSubKeysOutputArgs, opts ...pulumi.InvokeOption) ListWebPubSubKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebPubSubKeysResult, error) {
			args := v.(ListWebPubSubKeysArgs)
			r, err := ListWebPubSubKeys(ctx, &args, opts...)
			var s ListWebPubSubKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebPubSubKeysResultOutput)
}

type ListWebPubSubKeysOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListWebPubSubKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebPubSubKeysArgs)(nil)).Elem()
}

// A class represents the access keys of the resource.
type ListWebPubSubKeysResultOutput struct{ *pulumi.OutputState }

func (ListWebPubSubKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebPubSubKeysResult)(nil)).Elem()
}

func (o ListWebPubSubKeysResultOutput) ToListWebPubSubKeysResultOutput() ListWebPubSubKeysResultOutput {
	return o
}

func (o ListWebPubSubKeysResultOutput) ToListWebPubSubKeysResultOutputWithContext(ctx context.Context) ListWebPubSubKeysResultOutput {
	return o
}

// Connection string constructed via the primaryKey
func (o ListWebPubSubKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

// The primary access key.
func (o ListWebPubSubKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// Connection string constructed via the secondaryKey
func (o ListWebPubSubKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

// The secondary access key.
func (o ListWebPubSubKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebPubSubKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebPubSubKeysResultOutput{})
}
