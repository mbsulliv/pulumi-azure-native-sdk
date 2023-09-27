// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the blob inventory policy associated with the specified storage account.
func LookupBlobInventoryPolicy(ctx *pulumi.Context, args *LookupBlobInventoryPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBlobInventoryPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlobInventoryPolicyResult
	err := ctx.Invoke("azure-native:storage/v20220901:getBlobInventoryPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobInventoryPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the storage account blob inventory policy. It should always be 'default'
	BlobInventoryPolicyName string `pulumi:"blobInventoryPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The storage account blob inventory policy.
type LookupBlobInventoryPolicyResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Returns the last modified date and time of the blob inventory policy.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The storage account blob inventory policy object. It is composed of policy rules.
	Policy BlobInventoryPolicySchemaResponse `pulumi:"policy"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBlobInventoryPolicyOutput(ctx *pulumi.Context, args LookupBlobInventoryPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBlobInventoryPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobInventoryPolicyResult, error) {
			args := v.(LookupBlobInventoryPolicyArgs)
			r, err := LookupBlobInventoryPolicy(ctx, &args, opts...)
			var s LookupBlobInventoryPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobInventoryPolicyResultOutput)
}

type LookupBlobInventoryPolicyOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the storage account blob inventory policy. It should always be 'default'
	BlobInventoryPolicyName pulumi.StringInput `pulumi:"blobInventoryPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobInventoryPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobInventoryPolicyArgs)(nil)).Elem()
}

// The storage account blob inventory policy.
type LookupBlobInventoryPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBlobInventoryPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobInventoryPolicyResult)(nil)).Elem()
}

func (o LookupBlobInventoryPolicyResultOutput) ToLookupBlobInventoryPolicyResultOutput() LookupBlobInventoryPolicyResultOutput {
	return o
}

func (o LookupBlobInventoryPolicyResultOutput) ToLookupBlobInventoryPolicyResultOutputWithContext(ctx context.Context) LookupBlobInventoryPolicyResultOutput {
	return o
}

func (o LookupBlobInventoryPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBlobInventoryPolicyResult] {
	return pulumix.Output[LookupBlobInventoryPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBlobInventoryPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Returns the last modified date and time of the blob inventory policy.
func (o LookupBlobInventoryPolicyResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBlobInventoryPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The storage account blob inventory policy object. It is composed of policy rules.
func (o LookupBlobInventoryPolicyResultOutput) Policy() BlobInventoryPolicySchemaResponseOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) BlobInventoryPolicySchemaResponse { return v.Policy }).(BlobInventoryPolicySchemaResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupBlobInventoryPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBlobInventoryPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobInventoryPolicyResultOutput{})
}
