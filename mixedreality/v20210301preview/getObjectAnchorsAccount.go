// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieve an Object Anchors Account.
func LookupObjectAnchorsAccount(ctx *pulumi.Context, args *LookupObjectAnchorsAccountArgs, opts ...pulumi.InvokeOption) (*LookupObjectAnchorsAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupObjectAnchorsAccountResult
	err := ctx.Invoke("azure-native:mixedreality/v20210301preview:getObjectAnchorsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObjectAnchorsAccountArgs struct {
	// Name of an Mixed Reality Account.
	AccountName string `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ObjectAnchorsAccount Response.
type LookupObjectAnchorsAccountResult struct {
	// Correspond domain name of certain Spatial Anchors Account
	AccountDomain string `pulumi:"accountDomain"`
	// unique id of certain account.
	AccountId string `pulumi:"accountId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id       string                                `pulumi:"id"`
	Identity *ObjectAnchorsAccountResponseIdentity `pulumi:"identity"`
	// The kind of account, if supported
	Kind *SkuResponse `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The plan associated with this account
	Plan *IdentityResponse `pulumi:"plan"`
	// The sku associated with this account
	Sku *SkuResponse `pulumi:"sku"`
	// The name of the storage account associated with this accountId
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The system metadata related to an object anchors account.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupObjectAnchorsAccountOutput(ctx *pulumi.Context, args LookupObjectAnchorsAccountOutputArgs, opts ...pulumi.InvokeOption) LookupObjectAnchorsAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectAnchorsAccountResult, error) {
			args := v.(LookupObjectAnchorsAccountArgs)
			r, err := LookupObjectAnchorsAccount(ctx, &args, opts...)
			var s LookupObjectAnchorsAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObjectAnchorsAccountResultOutput)
}

type LookupObjectAnchorsAccountOutputArgs struct {
	// Name of an Mixed Reality Account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupObjectAnchorsAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectAnchorsAccountArgs)(nil)).Elem()
}

// ObjectAnchorsAccount Response.
type LookupObjectAnchorsAccountResultOutput struct{ *pulumi.OutputState }

func (LookupObjectAnchorsAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectAnchorsAccountResult)(nil)).Elem()
}

func (o LookupObjectAnchorsAccountResultOutput) ToLookupObjectAnchorsAccountResultOutput() LookupObjectAnchorsAccountResultOutput {
	return o
}

func (o LookupObjectAnchorsAccountResultOutput) ToLookupObjectAnchorsAccountResultOutputWithContext(ctx context.Context) LookupObjectAnchorsAccountResultOutput {
	return o
}

func (o LookupObjectAnchorsAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupObjectAnchorsAccountResult] {
	return pulumix.Output[LookupObjectAnchorsAccountResult]{
		OutputState: o.OutputState,
	}
}

// Correspond domain name of certain Spatial Anchors Account
func (o LookupObjectAnchorsAccountResultOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.AccountDomain }).(pulumi.StringOutput)
}

// unique id of certain account.
func (o LookupObjectAnchorsAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupObjectAnchorsAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Identity() ObjectAnchorsAccountResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *ObjectAnchorsAccountResponseIdentity { return v.Identity }).(ObjectAnchorsAccountResponseIdentityPtrOutput)
}

// The kind of account, if supported
func (o LookupObjectAnchorsAccountResultOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *SkuResponse { return v.Kind }).(SkuResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupObjectAnchorsAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupObjectAnchorsAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The plan associated with this account
func (o LookupObjectAnchorsAccountResultOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *IdentityResponse { return v.Plan }).(IdentityResponsePtrOutput)
}

// The sku associated with this account
func (o LookupObjectAnchorsAccountResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The name of the storage account associated with this accountId
func (o LookupObjectAnchorsAccountResultOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

// The system metadata related to an object anchors account.
func (o LookupObjectAnchorsAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupObjectAnchorsAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupObjectAnchorsAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectAnchorsAccountResultOutput{})
}
