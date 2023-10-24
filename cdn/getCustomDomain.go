// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an existing custom domain within an endpoint.
// Azure REST API version: 2023-05-01.
//
// Other available API versions: 2016-10-02, 2023-07-01-preview.
func LookupCustomDomain(ctx *pulumi.Context, args *LookupCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupCustomDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomDomainResult
	err := ctx.Invoke("azure-native:cdn:getCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDomainArgs struct {
	// Name of the custom domain within an endpoint.
	CustomDomainName string `pulumi:"customDomainName"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
type LookupCustomDomainResult struct {
	// Certificate parameters for securing custom HTTPS
	CustomHttpsParameters interface{} `pulumi:"customHttpsParameters"`
	// Provisioning status of the custom domain.
	CustomHttpsProvisioningState string `pulumi:"customHttpsProvisioningState"`
	// Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
	CustomHttpsProvisioningSubstate string `pulumi:"customHttpsProvisioningSubstate"`
	// The host name of the custom domain. Must be a domain name.
	HostName string `pulumi:"hostName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Provisioning status of Custom Https of the custom domain.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource status of the custom domain.
	ResourceState string `pulumi:"resourceState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
	// Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
	ValidationData *string `pulumi:"validationData"`
}

func LookupCustomDomainOutput(ctx *pulumi.Context, args LookupCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomDomainResult, error) {
			args := v.(LookupCustomDomainArgs)
			r, err := LookupCustomDomain(ctx, &args, opts...)
			var s LookupCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomDomainResultOutput)
}

type LookupCustomDomainOutputArgs struct {
	// Name of the custom domain within an endpoint.
	CustomDomainName pulumi.StringInput `pulumi:"customDomainName"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainArgs)(nil)).Elem()
}

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
type LookupCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainResult)(nil)).Elem()
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutput() LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutputWithContext(ctx context.Context) LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCustomDomainResult] {
	return pulumix.Output[LookupCustomDomainResult]{
		OutputState: o.OutputState,
	}
}

// Certificate parameters for securing custom HTTPS
func (o LookupCustomDomainResultOutput) CustomHttpsParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) interface{} { return v.CustomHttpsParameters }).(pulumi.AnyOutput)
}

// Provisioning status of the custom domain.
func (o LookupCustomDomainResultOutput) CustomHttpsProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.CustomHttpsProvisioningState }).(pulumi.StringOutput)
}

// Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
func (o LookupCustomDomainResultOutput) CustomHttpsProvisioningSubstate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.CustomHttpsProvisioningSubstate }).(pulumi.StringOutput)
}

// The host name of the custom domain. Must be a domain name.
func (o LookupCustomDomainResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning status of Custom Https of the custom domain.
func (o LookupCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the custom domain.
func (o LookupCustomDomainResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// Read only system data
func (o LookupCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

// Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
func (o LookupCustomDomainResultOutput) ValidationData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) *string { return v.ValidationData }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomDomainResultOutput{})
}
