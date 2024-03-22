// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240205preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a trusted Signing Account.
func LookupCodeSigningAccount(ctx *pulumi.Context, args *LookupCodeSigningAccountArgs, opts ...pulumi.InvokeOption) (*LookupCodeSigningAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCodeSigningAccountResult
	err := ctx.Invoke("azure-native:codesigning/v20240205preview:getCodeSigningAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodeSigningAccountArgs struct {
	// Trusted Signing account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Trusted signing account resource.
type LookupCodeSigningAccountResult struct {
	// The URI of the trusted signing account which is used during signing files.
	AccountUri string `pulumi:"accountUri"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Status of the current operation on trusted signing account.
	ProvisioningState string `pulumi:"provisioningState"`
	// SKU of the trusted signing account.
	Sku *AccountSkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCodeSigningAccountOutput(ctx *pulumi.Context, args LookupCodeSigningAccountOutputArgs, opts ...pulumi.InvokeOption) LookupCodeSigningAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodeSigningAccountResult, error) {
			args := v.(LookupCodeSigningAccountArgs)
			r, err := LookupCodeSigningAccount(ctx, &args, opts...)
			var s LookupCodeSigningAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodeSigningAccountResultOutput)
}

type LookupCodeSigningAccountOutputArgs struct {
	// Trusted Signing account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCodeSigningAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeSigningAccountArgs)(nil)).Elem()
}

// Trusted signing account resource.
type LookupCodeSigningAccountResultOutput struct{ *pulumi.OutputState }

func (LookupCodeSigningAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeSigningAccountResult)(nil)).Elem()
}

func (o LookupCodeSigningAccountResultOutput) ToLookupCodeSigningAccountResultOutput() LookupCodeSigningAccountResultOutput {
	return o
}

func (o LookupCodeSigningAccountResultOutput) ToLookupCodeSigningAccountResultOutputWithContext(ctx context.Context) LookupCodeSigningAccountResultOutput {
	return o
}

// The URI of the trusted signing account which is used during signing files.
func (o LookupCodeSigningAccountResultOutput) AccountUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) string { return v.AccountUri }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupCodeSigningAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupCodeSigningAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCodeSigningAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Status of the current operation on trusted signing account.
func (o LookupCodeSigningAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SKU of the trusted signing account.
func (o LookupCodeSigningAccountResultOutput) Sku() AccountSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) *AccountSkuResponse { return v.Sku }).(AccountSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCodeSigningAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupCodeSigningAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCodeSigningAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeSigningAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodeSigningAccountResultOutput{})
}
