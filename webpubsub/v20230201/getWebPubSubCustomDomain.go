// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a custom domain.
func LookupWebPubSubCustomDomain(ctx *pulumi.Context, args *LookupWebPubSubCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubCustomDomainResult, error) {
	var rv LookupWebPubSubCustomDomainResult
	err := ctx.Invoke("azure-native:webpubsub/v20230201:getWebPubSubCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubCustomDomainArgs struct {
	// Custom domain name.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A custom domain
type LookupWebPubSubCustomDomainResult struct {
	// Reference to a resource.
	CustomCertificate ResourceReferenceResponse `pulumi:"customCertificate"`
	// The custom domain name.
	DomainName string `pulumi:"domainName"`
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type string `pulumi:"type"`
}

func LookupWebPubSubCustomDomainOutput(ctx *pulumi.Context, args LookupWebPubSubCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubCustomDomainResult, error) {
			args := v.(LookupWebPubSubCustomDomainArgs)
			r, err := LookupWebPubSubCustomDomain(ctx, &args, opts...)
			var s LookupWebPubSubCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubCustomDomainResultOutput)
}

type LookupWebPubSubCustomDomainOutputArgs struct {
	// Custom domain name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomDomainArgs)(nil)).Elem()
}

// A custom domain
type LookupWebPubSubCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomDomainResult)(nil)).Elem()
}

func (o LookupWebPubSubCustomDomainResultOutput) ToLookupWebPubSubCustomDomainResultOutput() LookupWebPubSubCustomDomainResultOutput {
	return o
}

func (o LookupWebPubSubCustomDomainResultOutput) ToLookupWebPubSubCustomDomainResultOutputWithContext(ctx context.Context) LookupWebPubSubCustomDomainResultOutput {
	return o
}

// Reference to a resource.
func (o LookupWebPubSubCustomDomainResultOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) ResourceReferenceResponse { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

// The custom domain name.
func (o LookupWebPubSubCustomDomainResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource.
func (o LookupWebPubSubCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupWebPubSubCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupWebPubSubCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupWebPubSubCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o LookupWebPubSubCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubCustomDomainResultOutput{})
}