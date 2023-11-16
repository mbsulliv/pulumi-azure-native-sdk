// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Network Virtual Appliance.
func LookupNetworkVirtualAppliance(ctx *pulumi.Context, args *LookupNetworkVirtualApplianceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkVirtualApplianceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkVirtualApplianceResult
	err := ctx.Invoke("azure-native:network/v20230501:getNetworkVirtualAppliance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkVirtualApplianceArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// NetworkVirtualAppliance Resource.
type LookupNetworkVirtualApplianceResult struct {
	// Details required for Additional Network Interface.
	AdditionalNics []VirtualApplianceAdditionalNicPropertiesResponse `pulumi:"additionalNics"`
	// Address Prefix.
	AddressPrefix string `pulumi:"addressPrefix"`
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs []string `pulumi:"bootStrapConfigurationBlobs"`
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration *string `pulumi:"cloudInitConfiguration"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs []string `pulumi:"cloudInitConfigurationBlobs"`
	// The delegation for the Virtual Appliance
	Delegation *DelegationPropertiesResponse `pulumi:"delegation"`
	// The deployment type. PartnerManaged for the SaaS NVA
	DeploymentType string `pulumi:"deploymentType"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The service principal that has read access to cloud-init and config blob.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// List of references to InboundSecurityRules.
	InboundSecurityRules []SubResourceResponse `pulumi:"inboundSecurityRules"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Network Virtual Appliance SKU.
	NvaSku *VirtualApplianceSkuPropertiesResponse `pulumi:"nvaSku"`
	// The delegation for the Virtual Appliance
	PartnerManagedResource *PartnerManagedResourcePropertiesResponse `pulumi:"partnerManagedResource"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Public key for SSH login.
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// VirtualAppliance ASN. Microsoft private, public and IANA reserved ASN are not supported.
	VirtualApplianceAsn *float64 `pulumi:"virtualApplianceAsn"`
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics []VirtualApplianceNicPropertiesResponse `pulumi:"virtualApplianceNics"`
	// List of references to VirtualApplianceSite.
	VirtualApplianceSites []SubResourceResponse `pulumi:"virtualApplianceSites"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
}

func LookupNetworkVirtualApplianceOutput(ctx *pulumi.Context, args LookupNetworkVirtualApplianceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkVirtualApplianceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkVirtualApplianceResult, error) {
			args := v.(LookupNetworkVirtualApplianceArgs)
			r, err := LookupNetworkVirtualAppliance(ctx, &args, opts...)
			var s LookupNetworkVirtualApplianceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkVirtualApplianceResultOutput)
}

type LookupNetworkVirtualApplianceOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringInput `pulumi:"networkVirtualApplianceName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkVirtualApplianceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkVirtualApplianceArgs)(nil)).Elem()
}

// NetworkVirtualAppliance Resource.
type LookupNetworkVirtualApplianceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkVirtualApplianceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkVirtualApplianceResult)(nil)).Elem()
}

func (o LookupNetworkVirtualApplianceResultOutput) ToLookupNetworkVirtualApplianceResultOutput() LookupNetworkVirtualApplianceResultOutput {
	return o
}

func (o LookupNetworkVirtualApplianceResultOutput) ToLookupNetworkVirtualApplianceResultOutputWithContext(ctx context.Context) LookupNetworkVirtualApplianceResultOutput {
	return o
}

// Details required for Additional Network Interface.
func (o LookupNetworkVirtualApplianceResultOutput) AdditionalNics() VirtualApplianceAdditionalNicPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []VirtualApplianceAdditionalNicPropertiesResponse {
		return v.AdditionalNics
	}).(VirtualApplianceAdditionalNicPropertiesResponseArrayOutput)
}

// Address Prefix.
func (o LookupNetworkVirtualApplianceResultOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

// BootStrapConfigurationBlobs storage URLs.
func (o LookupNetworkVirtualApplianceResultOutput) BootStrapConfigurationBlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []string { return v.BootStrapConfigurationBlobs }).(pulumi.StringArrayOutput)
}

// CloudInitConfiguration string in plain text.
func (o LookupNetworkVirtualApplianceResultOutput) CloudInitConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.CloudInitConfiguration }).(pulumi.StringPtrOutput)
}

// CloudInitConfigurationBlob storage URLs.
func (o LookupNetworkVirtualApplianceResultOutput) CloudInitConfigurationBlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []string { return v.CloudInitConfigurationBlobs }).(pulumi.StringArrayOutput)
}

// The delegation for the Virtual Appliance
func (o LookupNetworkVirtualApplianceResultOutput) Delegation() DelegationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *DelegationPropertiesResponse { return v.Delegation }).(DelegationPropertiesResponsePtrOutput)
}

// The deployment type. PartnerManaged for the SaaS NVA
func (o LookupNetworkVirtualApplianceResultOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.DeploymentType }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkVirtualApplianceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNetworkVirtualApplianceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The service principal that has read access to cloud-init and config blob.
func (o LookupNetworkVirtualApplianceResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// List of references to InboundSecurityRules.
func (o LookupNetworkVirtualApplianceResultOutput) InboundSecurityRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []SubResourceResponse { return v.InboundSecurityRules }).(SubResourceResponseArrayOutput)
}

// Resource location.
func (o LookupNetworkVirtualApplianceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNetworkVirtualApplianceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network Virtual Appliance SKU.
func (o LookupNetworkVirtualApplianceResultOutput) NvaSku() VirtualApplianceSkuPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *VirtualApplianceSkuPropertiesResponse { return v.NvaSku }).(VirtualApplianceSkuPropertiesResponsePtrOutput)
}

// The delegation for the Virtual Appliance
func (o LookupNetworkVirtualApplianceResultOutput) PartnerManagedResource() PartnerManagedResourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *PartnerManagedResourcePropertiesResponse {
		return v.PartnerManagedResource
	}).(PartnerManagedResourcePropertiesResponsePtrOutput)
}

// The provisioning state of the resource.
func (o LookupNetworkVirtualApplianceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Public key for SSH login.
func (o LookupNetworkVirtualApplianceResultOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.SshPublicKey }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupNetworkVirtualApplianceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNetworkVirtualApplianceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.Type }).(pulumi.StringOutput)
}

// VirtualAppliance ASN. Microsoft private, public and IANA reserved ASN are not supported.
func (o LookupNetworkVirtualApplianceResultOutput) VirtualApplianceAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *float64 { return v.VirtualApplianceAsn }).(pulumi.Float64PtrOutput)
}

// List of Virtual Appliance Network Interfaces.
func (o LookupNetworkVirtualApplianceResultOutput) VirtualApplianceNics() VirtualApplianceNicPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []VirtualApplianceNicPropertiesResponse {
		return v.VirtualApplianceNics
	}).(VirtualApplianceNicPropertiesResponseArrayOutput)
}

// List of references to VirtualApplianceSite.
func (o LookupNetworkVirtualApplianceResultOutput) VirtualApplianceSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []SubResourceResponse { return v.VirtualApplianceSites }).(SubResourceResponseArrayOutput)
}

// The Virtual Hub where Network Virtual Appliance is being deployed.
func (o LookupNetworkVirtualApplianceResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkVirtualApplianceResultOutput{})
}
