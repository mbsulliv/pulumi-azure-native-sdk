// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managedidentity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Describes an Azure resource that is attached to an identity.
type AzureResourceResponse struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The name of this resource.
	Name string `pulumi:"name"`
	// The name of the resource group this resource belongs to.
	ResourceGroup string `pulumi:"resourceGroup"`
	// The name of the subscription this resource belongs to.
	SubscriptionDisplayName string `pulumi:"subscriptionDisplayName"`
	// The ID of the subscription this resource belongs to.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The type of this resource.
	Type string `pulumi:"type"`
}

// Describes an Azure resource that is attached to an identity.
type AzureResourceResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceResponse)(nil)).Elem()
}

func (o AzureResourceResponseOutput) ToAzureResourceResponseOutput() AzureResourceResponseOutput {
	return o
}

func (o AzureResourceResponseOutput) ToAzureResourceResponseOutputWithContext(ctx context.Context) AzureResourceResponseOutput {
	return o
}

// The ID of this resource.
func (o AzureResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The name of this resource.
func (o AzureResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the resource group this resource belongs to.
func (o AzureResourceResponseOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// The name of the subscription this resource belongs to.
func (o AzureResourceResponseOutput) SubscriptionDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.SubscriptionDisplayName }).(pulumi.StringOutput)
}

// The ID of the subscription this resource belongs to.
func (o AzureResourceResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// The type of this resource.
func (o AzureResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AzureResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureResourceResponse)(nil)).Elem()
}

func (o AzureResourceResponseArrayOutput) ToAzureResourceResponseArrayOutput() AzureResourceResponseArrayOutput {
	return o
}

func (o AzureResourceResponseArrayOutput) ToAzureResourceResponseArrayOutputWithContext(ctx context.Context) AzureResourceResponseArrayOutput {
	return o
}

func (o AzureResourceResponseArrayOutput) Index(i pulumi.IntInput) AzureResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureResourceResponse {
		return vs[0].([]AzureResourceResponse)[vs[1].(int)]
	}).(AzureResourceResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureResourceResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
