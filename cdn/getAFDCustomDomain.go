// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource group and profile.
// Azure REST API version: 2023-05-01.
//
// Other available API versions: 2023-07-01-preview.
func LookupAFDCustomDomain(ctx *pulumi.Context, args *LookupAFDCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupAFDCustomDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAFDCustomDomainResult
	err := ctx.Invoke("azure-native:cdn:getAFDCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDCustomDomainArgs struct {
	// Name of the domain under the profile which is unique globally.
	CustomDomainName string `pulumi:"customDomainName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
type LookupAFDCustomDomainResult struct {
	// Resource reference to the Azure DNS zone
	AzureDnsZone     *ResourceReferenceResponse `pulumi:"azureDnsZone"`
	DeploymentStatus string                     `pulumi:"deploymentStatus"`
	// Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step. DCV stands for DomainControlValidation.
	DomainValidationState string `pulumi:"domainValidationState"`
	// Key-Value pair representing migration properties for domains.
	ExtendedProperties map[string]string `pulumi:"extendedProperties"`
	// The host name of the domain. Must be a domain name.
	HostName string `pulumi:"hostName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource reference to the Azure resource where custom domain ownership was prevalidated
	PreValidatedCustomDomainResourceId *ResourceReferenceResponse `pulumi:"preValidatedCustomDomainResourceId"`
	// The name of the profile which holds the domain.
	ProfileName string `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The configuration specifying how to enable HTTPS for the domain - using AzureFrontDoor managed certificate or user's own certificate. If not specified, enabling ssl uses AzureFrontDoor managed certificate by default.
	TlsSettings *AFDDomainHttpsParametersResponse `pulumi:"tlsSettings"`
	// Resource type.
	Type string `pulumi:"type"`
	// Values the customer needs to validate domain ownership
	ValidationProperties DomainValidationPropertiesResponse `pulumi:"validationProperties"`
}

func LookupAFDCustomDomainOutput(ctx *pulumi.Context, args LookupAFDCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupAFDCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAFDCustomDomainResult, error) {
			args := v.(LookupAFDCustomDomainArgs)
			r, err := LookupAFDCustomDomain(ctx, &args, opts...)
			var s LookupAFDCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAFDCustomDomainResultOutput)
}

type LookupAFDCustomDomainOutputArgs struct {
	// Name of the domain under the profile which is unique globally.
	CustomDomainName pulumi.StringInput `pulumi:"customDomainName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDCustomDomainArgs)(nil)).Elem()
}

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
type LookupAFDCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupAFDCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDCustomDomainResult)(nil)).Elem()
}

func (o LookupAFDCustomDomainResultOutput) ToLookupAFDCustomDomainResultOutput() LookupAFDCustomDomainResultOutput {
	return o
}

func (o LookupAFDCustomDomainResultOutput) ToLookupAFDCustomDomainResultOutputWithContext(ctx context.Context) LookupAFDCustomDomainResultOutput {
	return o
}

// Resource reference to the Azure DNS zone
func (o LookupAFDCustomDomainResultOutput) AzureDnsZone() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) *ResourceReferenceResponse { return v.AzureDnsZone }).(ResourceReferenceResponsePtrOutput)
}

func (o LookupAFDCustomDomainResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step. DCV stands for DomainControlValidation.
func (o LookupAFDCustomDomainResultOutput) DomainValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.DomainValidationState }).(pulumi.StringOutput)
}

// Key-Value pair representing migration properties for domains.
func (o LookupAFDCustomDomainResultOutput) ExtendedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) map[string]string { return v.ExtendedProperties }).(pulumi.StringMapOutput)
}

// The host name of the domain. Must be a domain name.
func (o LookupAFDCustomDomainResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupAFDCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupAFDCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource reference to the Azure resource where custom domain ownership was prevalidated
func (o LookupAFDCustomDomainResultOutput) PreValidatedCustomDomainResourceId() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) *ResourceReferenceResponse {
		return v.PreValidatedCustomDomainResourceId
	}).(ResourceReferenceResponsePtrOutput)
}

// The name of the profile which holds the domain.
func (o LookupAFDCustomDomainResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o LookupAFDCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o LookupAFDCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The configuration specifying how to enable HTTPS for the domain - using AzureFrontDoor managed certificate or user's own certificate. If not specified, enabling ssl uses AzureFrontDoor managed certificate by default.
func (o LookupAFDCustomDomainResultOutput) TlsSettings() AFDDomainHttpsParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) *AFDDomainHttpsParametersResponse { return v.TlsSettings }).(AFDDomainHttpsParametersResponsePtrOutput)
}

// Resource type.
func (o LookupAFDCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

// Values the customer needs to validate domain ownership
func (o LookupAFDCustomDomainResultOutput) ValidationProperties() DomainValidationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) DomainValidationPropertiesResponse { return v.ValidationProperties }).(DomainValidationPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDCustomDomainResultOutput{})
}
