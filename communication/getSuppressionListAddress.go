// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package communication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a SuppressionListAddress.
// Azure REST API version: 2023-06-01-preview.
func LookupSuppressionListAddress(ctx *pulumi.Context, args *LookupSuppressionListAddressArgs, opts ...pulumi.InvokeOption) (*LookupSuppressionListAddressResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSuppressionListAddressResult
	err := ctx.Invoke("azure-native:communication:getSuppressionListAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSuppressionListAddressArgs struct {
	// The id of the address in a suppression list.
	AddressId string `pulumi:"addressId"`
	// The name of the Domains resource.
	DomainName string `pulumi:"domainName"`
	// The name of the EmailService resource.
	EmailServiceName string `pulumi:"emailServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the suppression list.
	SuppressionListName string `pulumi:"suppressionListName"`
}

// A object that represents a SuppressionList record.
type LookupSuppressionListAddressResult struct {
	// The location where the SuppressionListAddress data is stored at rest. This value is inherited from the parent Domains resource.
	DataLocation string `pulumi:"dataLocation"`
	// Email address of the recipient.
	Email string `pulumi:"email"`
	// The first name of the email recipient.
	FirstName *string `pulumi:"firstName"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The date the address was last updated in a suppression list.
	LastModified string `pulumi:"lastModified"`
	// The last name of the email recipient.
	LastName *string `pulumi:"lastName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// An optional property to provide contextual notes or a description for an address.
	Notes *string `pulumi:"notes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSuppressionListAddressOutput(ctx *pulumi.Context, args LookupSuppressionListAddressOutputArgs, opts ...pulumi.InvokeOption) LookupSuppressionListAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSuppressionListAddressResult, error) {
			args := v.(LookupSuppressionListAddressArgs)
			r, err := LookupSuppressionListAddress(ctx, &args, opts...)
			var s LookupSuppressionListAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSuppressionListAddressResultOutput)
}

type LookupSuppressionListAddressOutputArgs struct {
	// The id of the address in a suppression list.
	AddressId pulumi.StringInput `pulumi:"addressId"`
	// The name of the Domains resource.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the EmailService resource.
	EmailServiceName pulumi.StringInput `pulumi:"emailServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the suppression list.
	SuppressionListName pulumi.StringInput `pulumi:"suppressionListName"`
}

func (LookupSuppressionListAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionListAddressArgs)(nil)).Elem()
}

// A object that represents a SuppressionList record.
type LookupSuppressionListAddressResultOutput struct{ *pulumi.OutputState }

func (LookupSuppressionListAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionListAddressResult)(nil)).Elem()
}

func (o LookupSuppressionListAddressResultOutput) ToLookupSuppressionListAddressResultOutput() LookupSuppressionListAddressResultOutput {
	return o
}

func (o LookupSuppressionListAddressResultOutput) ToLookupSuppressionListAddressResultOutputWithContext(ctx context.Context) LookupSuppressionListAddressResultOutput {
	return o
}

// The location where the SuppressionListAddress data is stored at rest. This value is inherited from the parent Domains resource.
func (o LookupSuppressionListAddressResultOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) string { return v.DataLocation }).(pulumi.StringOutput)
}

// Email address of the recipient.
func (o LookupSuppressionListAddressResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) string { return v.Email }).(pulumi.StringOutput)
}

// The first name of the email recipient.
func (o LookupSuppressionListAddressResultOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSuppressionListAddressResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) string { return v.Id }).(pulumi.StringOutput)
}

// The date the address was last updated in a suppression list.
func (o LookupSuppressionListAddressResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The last name of the email recipient.
func (o LookupSuppressionListAddressResultOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupSuppressionListAddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) string { return v.Name }).(pulumi.StringOutput)
}

// An optional property to provide contextual notes or a description for an address.
func (o LookupSuppressionListAddressResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSuppressionListAddressResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSuppressionListAddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionListAddressResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSuppressionListAddressResultOutput{})
}
