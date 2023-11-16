// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datareplication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the protected item.
// Azure REST API version: 2021-02-16-preview.
func LookupProtectedItem(ctx *pulumi.Context, args *LookupProtectedItemArgs, opts ...pulumi.InvokeOption) (*LookupProtectedItemResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProtectedItemResult
	err := ctx.Invoke("azure-native:datareplication:getProtectedItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectedItemArgs struct {
	// The protected item name.
	ProtectedItemName string `pulumi:"protectedItemName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName string `pulumi:"vaultName"`
}

// Protected item model.
type LookupProtectedItemResult struct {
	// Gets or sets the Id of the resource.
	Id string `pulumi:"id"`
	// Gets or sets the name of the resource.
	Name string `pulumi:"name"`
	// Protected item model properties.
	Properties ProtectedItemModelPropertiesResponse `pulumi:"properties"`
	SystemData ProtectedItemModelResponseSystemData `pulumi:"systemData"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
}

func LookupProtectedItemOutput(ctx *pulumi.Context, args LookupProtectedItemOutputArgs, opts ...pulumi.InvokeOption) LookupProtectedItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectedItemResult, error) {
			args := v.(LookupProtectedItemArgs)
			r, err := LookupProtectedItem(ctx, &args, opts...)
			var s LookupProtectedItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProtectedItemResultOutput)
}

type LookupProtectedItemOutputArgs struct {
	// The protected item name.
	ProtectedItemName pulumi.StringInput `pulumi:"protectedItemName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupProtectedItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectedItemArgs)(nil)).Elem()
}

// Protected item model.
type LookupProtectedItemResultOutput struct{ *pulumi.OutputState }

func (LookupProtectedItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectedItemResult)(nil)).Elem()
}

func (o LookupProtectedItemResultOutput) ToLookupProtectedItemResultOutput() LookupProtectedItemResultOutput {
	return o
}

func (o LookupProtectedItemResultOutput) ToLookupProtectedItemResultOutputWithContext(ctx context.Context) LookupProtectedItemResultOutput {
	return o
}

// Gets or sets the Id of the resource.
func (o LookupProtectedItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the name of the resource.
func (o LookupProtectedItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) string { return v.Name }).(pulumi.StringOutput)
}

// Protected item model properties.
func (o LookupProtectedItemResultOutput) Properties() ProtectedItemModelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) ProtectedItemModelPropertiesResponse { return v.Properties }).(ProtectedItemModelPropertiesResponseOutput)
}

func (o LookupProtectedItemResultOutput) SystemData() ProtectedItemModelResponseSystemDataOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) ProtectedItemModelResponseSystemData { return v.SystemData }).(ProtectedItemModelResponseSystemDataOutput)
}

// Gets or sets the type of the resource.
func (o LookupProtectedItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectedItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectedItemResultOutput{})
}
