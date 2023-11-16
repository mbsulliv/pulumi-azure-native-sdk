// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Tag description in scope of API
func LookupApiTagDescription(ctx *pulumi.Context, args *LookupApiTagDescriptionArgs, opts ...pulumi.InvokeOption) (*LookupApiTagDescriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiTagDescriptionResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getApiTagDescription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiTagDescriptionArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag description identifier. Used when creating tagDescription for API/Tag association. Based on API and Tag names.
	TagDescriptionId string `pulumi:"tagDescriptionId"`
}

// Contract details.
type LookupApiTagDescriptionResult struct {
	// Description of the Tag.
	Description *string `pulumi:"description"`
	// Tag name.
	DisplayName *string `pulumi:"displayName"`
	// Description of the external resources describing the tag.
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	// Absolute URL of external resources describing the tag.
	ExternalDocsUrl *string `pulumi:"externalDocsUrl"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Identifier of the tag in the form of /tags/{tagId}
	TagId *string `pulumi:"tagId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupApiTagDescriptionOutput(ctx *pulumi.Context, args LookupApiTagDescriptionOutputArgs, opts ...pulumi.InvokeOption) LookupApiTagDescriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiTagDescriptionResult, error) {
			args := v.(LookupApiTagDescriptionArgs)
			r, err := LookupApiTagDescription(ctx, &args, opts...)
			var s LookupApiTagDescriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiTagDescriptionResultOutput)
}

type LookupApiTagDescriptionOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag description identifier. Used when creating tagDescription for API/Tag association. Based on API and Tag names.
	TagDescriptionId pulumi.StringInput `pulumi:"tagDescriptionId"`
}

func (LookupApiTagDescriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiTagDescriptionArgs)(nil)).Elem()
}

// Contract details.
type LookupApiTagDescriptionResultOutput struct{ *pulumi.OutputState }

func (LookupApiTagDescriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiTagDescriptionResult)(nil)).Elem()
}

func (o LookupApiTagDescriptionResultOutput) ToLookupApiTagDescriptionResultOutput() LookupApiTagDescriptionResultOutput {
	return o
}

func (o LookupApiTagDescriptionResultOutput) ToLookupApiTagDescriptionResultOutputWithContext(ctx context.Context) LookupApiTagDescriptionResultOutput {
	return o
}

// Description of the Tag.
func (o LookupApiTagDescriptionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Tag name.
func (o LookupApiTagDescriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Description of the external resources describing the tag.
func (o LookupApiTagDescriptionResultOutput) ExternalDocsDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.ExternalDocsDescription }).(pulumi.StringPtrOutput)
}

// Absolute URL of external resources describing the tag.
func (o LookupApiTagDescriptionResultOutput) ExternalDocsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.ExternalDocsUrl }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiTagDescriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiTagDescriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Identifier of the tag in the form of /tags/{tagId}
func (o LookupApiTagDescriptionResultOutput) TagId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.TagId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiTagDescriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiTagDescriptionResultOutput{})
}
