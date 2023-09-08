// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the access keys of the CommunicationService resource.
func ListCommunicationServiceKeys(ctx *pulumi.Context, args *ListCommunicationServiceKeysArgs, opts ...pulumi.InvokeOption) (*ListCommunicationServiceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListCommunicationServiceKeysResult
	err := ctx.Invoke("azure-native:communication/v20230401preview:listCommunicationServiceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCommunicationServiceKeysArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName string `pulumi:"communicationServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A class representing the access keys of a CommunicationService.
type ListCommunicationServiceKeysResult struct {
	// CommunicationService connection string constructed via the primaryKey
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The primary access key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// CommunicationService connection string constructed via the secondaryKey
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The secondary access key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListCommunicationServiceKeysOutput(ctx *pulumi.Context, args ListCommunicationServiceKeysOutputArgs, opts ...pulumi.InvokeOption) ListCommunicationServiceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCommunicationServiceKeysResult, error) {
			args := v.(ListCommunicationServiceKeysArgs)
			r, err := ListCommunicationServiceKeys(ctx, &args, opts...)
			var s ListCommunicationServiceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCommunicationServiceKeysResultOutput)
}

type ListCommunicationServiceKeysOutputArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName pulumi.StringInput `pulumi:"communicationServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListCommunicationServiceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCommunicationServiceKeysArgs)(nil)).Elem()
}

// A class representing the access keys of a CommunicationService.
type ListCommunicationServiceKeysResultOutput struct{ *pulumi.OutputState }

func (ListCommunicationServiceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCommunicationServiceKeysResult)(nil)).Elem()
}

func (o ListCommunicationServiceKeysResultOutput) ToListCommunicationServiceKeysResultOutput() ListCommunicationServiceKeysResultOutput {
	return o
}

func (o ListCommunicationServiceKeysResultOutput) ToListCommunicationServiceKeysResultOutputWithContext(ctx context.Context) ListCommunicationServiceKeysResultOutput {
	return o
}

// CommunicationService connection string constructed via the primaryKey
func (o ListCommunicationServiceKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

// The primary access key.
func (o ListCommunicationServiceKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// CommunicationService connection string constructed via the secondaryKey
func (o ListCommunicationServiceKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

// The secondary access key.
func (o ListCommunicationServiceKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCommunicationServiceKeysResultOutput{})
}