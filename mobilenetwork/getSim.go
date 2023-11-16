// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mobilenetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified SIM.
// Azure REST API version: 2023-06-01.
//
// Other available API versions: 2022-03-01-preview, 2022-04-01-preview, 2022-11-01, 2023-09-01.
func LookupSim(ctx *pulumi.Context, args *LookupSimArgs, opts ...pulumi.InvokeOption) (*LookupSimResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSimResult
	err := ctx.Invoke("azure-native:mobilenetwork:getSim", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SIM Group.
	SimGroupName string `pulumi:"simGroupName"`
	// The name of the SIM.
	SimName string `pulumi:"simName"`
}

// SIM resource.
type LookupSimResult struct {
	// An optional free-form text field that can be used to record the device type this SIM is associated with, for example 'Video camera'. The Azure portal allows SIMs to be grouped and filtered based on this value.
	DeviceType *string `pulumi:"deviceType"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The integrated circuit card ID (ICCID) for the SIM.
	IntegratedCircuitCardIdentifier *string `pulumi:"integratedCircuitCardIdentifier"`
	// The international mobile subscriber identity (IMSI) for the SIM.
	InternationalMobileSubscriberIdentity string `pulumi:"internationalMobileSubscriberIdentity"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the SIM resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SIM policy used by this SIM. The SIM policy must be in the same location as the SIM.
	SimPolicy *SimPolicyResourceIdResponse `pulumi:"simPolicy"`
	// The state of the SIM resource.
	SimState string `pulumi:"simState"`
	// A dictionary of sites to the provisioning state of this SIM on that site.
	SiteProvisioningState map[string]string `pulumi:"siteProvisioningState"`
	// A list of static IP addresses assigned to this SIM. Each address is assigned at a defined network scope, made up of {attached data network, slice}.
	StaticIpConfiguration []SimStaticIpPropertiesResponse `pulumi:"staticIpConfiguration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The public key fingerprint of the SIM vendor who provided this SIM, if any.
	VendorKeyFingerprint string `pulumi:"vendorKeyFingerprint"`
	// The name of the SIM vendor who provided this SIM, if any.
	VendorName string `pulumi:"vendorName"`
}

func LookupSimOutput(ctx *pulumi.Context, args LookupSimOutputArgs, opts ...pulumi.InvokeOption) LookupSimResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimResult, error) {
			args := v.(LookupSimArgs)
			r, err := LookupSim(ctx, &args, opts...)
			var s LookupSimResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimResultOutput)
}

type LookupSimOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the SIM Group.
	SimGroupName pulumi.StringInput `pulumi:"simGroupName"`
	// The name of the SIM.
	SimName pulumi.StringInput `pulumi:"simName"`
}

func (LookupSimOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimArgs)(nil)).Elem()
}

// SIM resource.
type LookupSimResultOutput struct{ *pulumi.OutputState }

func (LookupSimResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimResult)(nil)).Elem()
}

func (o LookupSimResultOutput) ToLookupSimResultOutput() LookupSimResultOutput {
	return o
}

func (o LookupSimResultOutput) ToLookupSimResultOutputWithContext(ctx context.Context) LookupSimResultOutput {
	return o
}

// An optional free-form text field that can be used to record the device type this SIM is associated with, for example 'Video camera'. The Azure portal allows SIMs to be grouped and filtered based on this value.
func (o LookupSimResultOutput) DeviceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.DeviceType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSimResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.Id }).(pulumi.StringOutput)
}

// The integrated circuit card ID (ICCID) for the SIM.
func (o LookupSimResultOutput) IntegratedCircuitCardIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.IntegratedCircuitCardIdentifier }).(pulumi.StringPtrOutput)
}

// The international mobile subscriber identity (IMSI) for the SIM.
func (o LookupSimResultOutput) InternationalMobileSubscriberIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.InternationalMobileSubscriberIdentity }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSimResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the SIM resource.
func (o LookupSimResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SIM policy used by this SIM. The SIM policy must be in the same location as the SIM.
func (o LookupSimResultOutput) SimPolicy() SimPolicyResourceIdResponsePtrOutput {
	return o.ApplyT(func(v LookupSimResult) *SimPolicyResourceIdResponse { return v.SimPolicy }).(SimPolicyResourceIdResponsePtrOutput)
}

// The state of the SIM resource.
func (o LookupSimResultOutput) SimState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.SimState }).(pulumi.StringOutput)
}

// A dictionary of sites to the provisioning state of this SIM on that site.
func (o LookupSimResultOutput) SiteProvisioningState() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSimResult) map[string]string { return v.SiteProvisioningState }).(pulumi.StringMapOutput)
}

// A list of static IP addresses assigned to this SIM. Each address is assigned at a defined network scope, made up of {attached data network, slice}.
func (o LookupSimResultOutput) StaticIpConfiguration() SimStaticIpPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupSimResult) []SimStaticIpPropertiesResponse { return v.StaticIpConfiguration }).(SimStaticIpPropertiesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSimResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSimResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSimResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.Type }).(pulumi.StringOutput)
}

// The public key fingerprint of the SIM vendor who provided this SIM, if any.
func (o LookupSimResultOutput) VendorKeyFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.VendorKeyFingerprint }).(pulumi.StringOutput)
}

// The name of the SIM vendor who provided this SIM, if any.
func (o LookupSimResultOutput) VendorName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.VendorName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimResultOutput{})
}
