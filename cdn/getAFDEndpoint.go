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

// Gets an existing AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
// Azure REST API version: 2023-05-01.
func LookupAFDEndpoint(ctx *pulumi.Context, args *LookupAFDEndpointArgs, opts ...pulumi.InvokeOption) (*LookupAFDEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAFDEndpointResult
	err := ctx.Invoke("azure-native:cdn:getAFDEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDEndpointArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Front Door endpoint is the entity within a Azure Front Door profile containing configuration information such as origin, protocol, content caching and delivery behavior. The AzureFrontDoor endpoint uses the URL format <endpointname>.azureedge.net.
type LookupAFDEndpointResult struct {
	// Indicates the endpoint name reuse scope. The default value is TenantReuse.
	AutoGeneratedDomainNameLabelScope *string `pulumi:"autoGeneratedDomainNameLabelScope"`
	DeploymentStatus                  string  `pulumi:"deploymentStatus"`
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *string `pulumi:"enabledState"`
	// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
	HostName string `pulumi:"hostName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The name of the profile which holds the endpoint.
	ProfileName string `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAFDEndpointOutput(ctx *pulumi.Context, args LookupAFDEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupAFDEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAFDEndpointResult, error) {
			args := v.(LookupAFDEndpointArgs)
			r, err := LookupAFDEndpoint(ctx, &args, opts...)
			var s LookupAFDEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAFDEndpointResultOutput)
}

type LookupAFDEndpointOutputArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDEndpointArgs)(nil)).Elem()
}

// Azure Front Door endpoint is the entity within a Azure Front Door profile containing configuration information such as origin, protocol, content caching and delivery behavior. The AzureFrontDoor endpoint uses the URL format <endpointname>.azureedge.net.
type LookupAFDEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupAFDEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDEndpointResult)(nil)).Elem()
}

func (o LookupAFDEndpointResultOutput) ToLookupAFDEndpointResultOutput() LookupAFDEndpointResultOutput {
	return o
}

func (o LookupAFDEndpointResultOutput) ToLookupAFDEndpointResultOutputWithContext(ctx context.Context) LookupAFDEndpointResultOutput {
	return o
}

func (o LookupAFDEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAFDEndpointResult] {
	return pulumix.Output[LookupAFDEndpointResult]{
		OutputState: o.OutputState,
	}
}

// Indicates the endpoint name reuse scope. The default value is TenantReuse.
func (o LookupAFDEndpointResultOutput) AutoGeneratedDomainNameLabelScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) *string { return v.AutoGeneratedDomainNameLabelScope }).(pulumi.StringPtrOutput)
}

func (o LookupAFDEndpointResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
func (o LookupAFDEndpointResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
func (o LookupAFDEndpointResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupAFDEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupAFDEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupAFDEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the profile which holds the endpoint.
func (o LookupAFDEndpointResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o LookupAFDEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o LookupAFDEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAFDEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupAFDEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDEndpointResultOutput{})
}
