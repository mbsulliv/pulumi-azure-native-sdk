// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified SIM policy.
func LookupSimPolicy(ctx *pulumi.Context, args *LookupSimPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSimPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSimPolicyResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20230901:getSimPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSimPolicyArgs struct {
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SIM policy.
	SimPolicyName string `pulumi:"simPolicyName"`
}

// SIM policy resource.
type LookupSimPolicyResult struct {
	// The default slice to use if the UE does not explicitly specify it. This slice must exist in the `sliceConfigurations` map. The slice must be in the same location as the SIM policy.
	DefaultSlice SliceResourceIdResponse `pulumi:"defaultSlice"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the SIM policy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// UE periodic registration update timer (5G) or UE periodic tracking area update timer (4G), in seconds.
	RegistrationTimer *int `pulumi:"registrationTimer"`
	// RAT/Frequency Selection Priority Index, defined in 3GPP TS 36.413. This is an optional setting and by default is unspecified.
	RfspIndex *int `pulumi:"rfspIndex"`
	// A dictionary of sites to the provisioning state of this SIM policy on that site.
	SiteProvisioningState map[string]string `pulumi:"siteProvisioningState"`
	// The allowed slices and the settings to use for them. The list must not contain duplicate items and must contain at least one item.
	SliceConfigurations []SliceConfigurationResponse `pulumi:"sliceConfigurations"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Aggregate maximum bit rate across all non-GBR QoS flows of all PDU sessions of a given UE. See 3GPP TS23.501 section 5.7.2.6 for a full description of the UE-AMBR.
	UeAmbr AmbrResponse `pulumi:"ueAmbr"`
}

// Defaults sets the appropriate defaults for LookupSimPolicyResult
func (val *LookupSimPolicyResult) Defaults() *LookupSimPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.RegistrationTimer == nil {
		registrationTimer_ := 3240
		tmp.RegistrationTimer = &registrationTimer_
	}
	return &tmp
}

func LookupSimPolicyOutput(ctx *pulumi.Context, args LookupSimPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSimPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimPolicyResult, error) {
			args := v.(LookupSimPolicyArgs)
			r, err := LookupSimPolicy(ctx, &args, opts...)
			var s LookupSimPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimPolicyResultOutput)
}

type LookupSimPolicyOutputArgs struct {
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the SIM policy.
	SimPolicyName pulumi.StringInput `pulumi:"simPolicyName"`
}

func (LookupSimPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimPolicyArgs)(nil)).Elem()
}

// SIM policy resource.
type LookupSimPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSimPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimPolicyResult)(nil)).Elem()
}

func (o LookupSimPolicyResultOutput) ToLookupSimPolicyResultOutput() LookupSimPolicyResultOutput {
	return o
}

func (o LookupSimPolicyResultOutput) ToLookupSimPolicyResultOutputWithContext(ctx context.Context) LookupSimPolicyResultOutput {
	return o
}

// The default slice to use if the UE does not explicitly specify it. This slice must exist in the `sliceConfigurations` map. The slice must be in the same location as the SIM policy.
func (o LookupSimPolicyResultOutput) DefaultSlice() SliceResourceIdResponseOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) SliceResourceIdResponse { return v.DefaultSlice }).(SliceResourceIdResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSimPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSimPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSimPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the SIM policy resource.
func (o LookupSimPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// UE periodic registration update timer (5G) or UE periodic tracking area update timer (4G), in seconds.
func (o LookupSimPolicyResultOutput) RegistrationTimer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *int { return v.RegistrationTimer }).(pulumi.IntPtrOutput)
}

// RAT/Frequency Selection Priority Index, defined in 3GPP TS 36.413. This is an optional setting and by default is unspecified.
func (o LookupSimPolicyResultOutput) RfspIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) *int { return v.RfspIndex }).(pulumi.IntPtrOutput)
}

// A dictionary of sites to the provisioning state of this SIM policy on that site.
func (o LookupSimPolicyResultOutput) SiteProvisioningState() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) map[string]string { return v.SiteProvisioningState }).(pulumi.StringMapOutput)
}

// The allowed slices and the settings to use for them. The list must not contain duplicate items and must contain at least one item.
func (o LookupSimPolicyResultOutput) SliceConfigurations() SliceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) []SliceConfigurationResponse { return v.SliceConfigurations }).(SliceConfigurationResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSimPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSimPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSimPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Aggregate maximum bit rate across all non-GBR QoS flows of all PDU sessions of a given UE. See 3GPP TS23.501 section 5.7.2.6 for a full description of the UE-AMBR.
func (o LookupSimPolicyResultOutput) UeAmbr() AmbrResponseOutput {
	return o.ApplyT(func(v LookupSimPolicyResult) AmbrResponse { return v.UeAmbr }).(AmbrResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimPolicyResultOutput{})
}
