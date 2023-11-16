// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211130

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Azure dedicated HSM.
func LookupDedicatedHsm(ctx *pulumi.Context, args *LookupDedicatedHsmArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHsmResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDedicatedHsmResult
	err := ctx.Invoke("azure-native:hardwaresecuritymodules/v20211130:getDedicatedHsm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHsmArgs struct {
	// The name of the dedicated HSM.
	Name string `pulumi:"name"`
	// The name of the Resource Group to which the dedicated hsm belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Resource information with extended details.
type LookupDedicatedHsmResult struct {
	// The Azure Resource Manager resource ID for the dedicated HSM.
	Id string `pulumi:"id"`
	// The supported Azure location where the dedicated HSM should be created.
	Location string `pulumi:"location"`
	// Specifies the management network interfaces of the dedicated hsm.
	ManagementNetworkProfile *NetworkProfileResponse `pulumi:"managementNetworkProfile"`
	// The name of the dedicated HSM.
	Name string `pulumi:"name"`
	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// SKU details
	Sku SkuResponse `pulumi:"sku"`
	// This field will be used when RP does not support Availability zones.
	StampId *string `pulumi:"stampId"`
	// Resource Status Message.
	StatusMessage string `pulumi:"statusMessage"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The resource type of the dedicated HSM.
	Type string `pulumi:"type"`
	// The Dedicated Hsm zones.
	Zones []string `pulumi:"zones"`
}

func LookupDedicatedHsmOutput(ctx *pulumi.Context, args LookupDedicatedHsmOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedHsmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedHsmResult, error) {
			args := v.(LookupDedicatedHsmArgs)
			r, err := LookupDedicatedHsm(ctx, &args, opts...)
			var s LookupDedicatedHsmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedHsmResultOutput)
}

type LookupDedicatedHsmOutputArgs struct {
	// The name of the dedicated HSM.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group to which the dedicated hsm belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDedicatedHsmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHsmArgs)(nil)).Elem()
}

// Resource information with extended details.
type LookupDedicatedHsmResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedHsmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHsmResult)(nil)).Elem()
}

func (o LookupDedicatedHsmResultOutput) ToLookupDedicatedHsmResultOutput() LookupDedicatedHsmResultOutput {
	return o
}

func (o LookupDedicatedHsmResultOutput) ToLookupDedicatedHsmResultOutputWithContext(ctx context.Context) LookupDedicatedHsmResultOutput {
	return o
}

// The Azure Resource Manager resource ID for the dedicated HSM.
func (o LookupDedicatedHsmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Id }).(pulumi.StringOutput)
}

// The supported Azure location where the dedicated HSM should be created.
func (o LookupDedicatedHsmResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Location }).(pulumi.StringOutput)
}

// Specifies the management network interfaces of the dedicated hsm.
func (o LookupDedicatedHsmResultOutput) ManagementNetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) *NetworkProfileResponse { return v.ManagementNetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// The name of the dedicated HSM.
func (o LookupDedicatedHsmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the network interfaces of the dedicated hsm.
func (o LookupDedicatedHsmResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// Provisioning state.
func (o LookupDedicatedHsmResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SKU details
func (o LookupDedicatedHsmResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// This field will be used when RP does not support Availability zones.
func (o LookupDedicatedHsmResultOutput) StampId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) *string { return v.StampId }).(pulumi.StringPtrOutput)
}

// Resource Status Message.
func (o LookupDedicatedHsmResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o LookupDedicatedHsmResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupDedicatedHsmResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type of the dedicated HSM.
func (o LookupDedicatedHsmResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) string { return v.Type }).(pulumi.StringOutput)
}

// The Dedicated Hsm zones.
func (o LookupDedicatedHsmResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHsmResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedHsmResultOutput{})
}
