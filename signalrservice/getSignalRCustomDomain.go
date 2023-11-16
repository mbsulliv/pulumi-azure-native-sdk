// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a custom domain.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview.
func LookupSignalRCustomDomain(ctx *pulumi.Context, args *LookupSignalRCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupSignalRCustomDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalRCustomDomainResult
	err := ctx.Invoke("azure-native:signalrservice:getSignalRCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRCustomDomainArgs struct {
	// Custom domain name.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A custom domain
type LookupSignalRCustomDomainResult struct {
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

func LookupSignalRCustomDomainOutput(ctx *pulumi.Context, args LookupSignalRCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRCustomDomainResult, error) {
			args := v.(LookupSignalRCustomDomainArgs)
			r, err := LookupSignalRCustomDomain(ctx, &args, opts...)
			var s LookupSignalRCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRCustomDomainResultOutput)
}

type LookupSignalRCustomDomainOutputArgs struct {
	// Custom domain name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalRCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomDomainArgs)(nil)).Elem()
}

// A custom domain
type LookupSignalRCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomDomainResult)(nil)).Elem()
}

func (o LookupSignalRCustomDomainResultOutput) ToLookupSignalRCustomDomainResultOutput() LookupSignalRCustomDomainResultOutput {
	return o
}

func (o LookupSignalRCustomDomainResultOutput) ToLookupSignalRCustomDomainResultOutputWithContext(ctx context.Context) LookupSignalRCustomDomainResultOutput {
	return o
}

// Reference to a resource.
func (o LookupSignalRCustomDomainResultOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) ResourceReferenceResponse { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

// The custom domain name.
func (o LookupSignalRCustomDomainResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource.
func (o LookupSignalRCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupSignalRCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupSignalRCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupSignalRCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o LookupSignalRCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRCustomDomainResultOutput{})
}
