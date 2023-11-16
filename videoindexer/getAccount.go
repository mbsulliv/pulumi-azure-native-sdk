// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videoindexer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of an Azure Video Indexer account.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2024-01-01.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:videoindexer:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	// The name of the Azure Video Indexer account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Video Indexer account.
type LookupAccountResult struct {
	// The account's data-plane ID. This can be set only when connecting an existing classic account
	AccountId *string `pulumi:"accountId"`
	// The account's name
	AccountName string `pulumi:"accountName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The media services details
	MediaServices *MediaServicesForPutRequestResponse `pulumi:"mediaServices"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the status of the account at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The account's tenant id
	TenantId string `pulumi:"tenantId"`
	// An integer representing the total seconds that have been indexed on the account
	TotalSecondsIndexed int `pulumi:"totalSecondsIndexed"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAccountResult
func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AccountId == nil {
		accountId_ := "00000000-0000-0000-0000-000000000000"
		tmp.AccountId = &accountId_
	}
	return &tmp
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	// The name of the Azure Video Indexer account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

// An Azure Video Indexer account.
type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

// The account's data-plane ID. This can be set only when connecting an existing classic account
func (o LookupAccountResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// The account's name
func (o LookupAccountResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.AccountName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o LookupAccountResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// The media services details
func (o LookupAccountResultOutput) MediaServices() MediaServicesForPutRequestResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *MediaServicesForPutRequestResponse { return v.MediaServices }).(MediaServicesForPutRequestResponsePtrOutput)
}

// The name of the resource
func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the status of the account at the time the operation was called.
func (o LookupAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o LookupAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The account's tenant id
func (o LookupAccountResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// An integer representing the total seconds that have been indexed on the account
func (o LookupAccountResultOutput) TotalSecondsIndexed() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.TotalSecondsIndexed }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
