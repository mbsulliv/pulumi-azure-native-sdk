// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the properties of a lab resource.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2018-10-15, 2023-06-07.
func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:labservices:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupLabArgs struct {
	// The name of the lab that uniquely identifies it within containing lab plan. Used in resource URIs.
	LabName string `pulumi:"labName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The lab resource.
type LookupLabResult struct {
	// The resource auto shutdown configuration for the lab. This controls whether actions are taken on resources that are sitting idle.
	AutoShutdownProfile AutoShutdownProfileResponse `pulumi:"autoShutdownProfile"`
	// The connection profile for the lab. This controls settings such as web access to lab resources or whether RDP or SSH ports are open.
	ConnectionProfile ConnectionProfileResponse `pulumi:"connectionProfile"`
	// The description of the lab.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The ID of the lab plan. Used during resource creation to provide defaults and acts as a permission container when creating a lab via labs.azure.com. Setting a labPlanId on an existing lab provides organization..
	LabPlanId *string `pulumi:"labPlanId"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The network profile for the lab, typically applied via a lab plan. This profile cannot be modified once a lab has been created.
	NetworkProfile *LabNetworkProfileResponse `pulumi:"networkProfile"`
	// Current provisioning state of the lab.
	ProvisioningState string `pulumi:"provisioningState"`
	// The lab user list management profile.
	RosterProfile *RosterProfileResponse `pulumi:"rosterProfile"`
	// The lab security profile.
	SecurityProfile SecurityProfileResponse `pulumi:"securityProfile"`
	// The lab state.
	State string `pulumi:"state"`
	// Metadata pertaining to creation and last modification of the lab.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The title of the lab.
	Title *string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The profile used for creating lab virtual machines.
	VirtualMachineProfile VirtualMachineProfileResponse `pulumi:"virtualMachineProfile"`
}

// Defaults sets the appropriate defaults for LookupLabResult
func (val *LookupLabResult) Defaults() *LookupLabResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.AutoShutdownProfile = *tmp.AutoShutdownProfile.Defaults()

	tmp.ConnectionProfile = *tmp.ConnectionProfile.Defaults()

	tmp.VirtualMachineProfile = *tmp.VirtualMachineProfile.Defaults()

	return &tmp
}

func LookupLabOutput(ctx *pulumi.Context, args LookupLabOutputArgs, opts ...pulumi.InvokeOption) LookupLabResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResult, error) {
			args := v.(LookupLabArgs)
			r, err := LookupLab(ctx, &args, opts...)
			var s LookupLabResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResultOutput)
}

type LookupLabOutputArgs struct {
	// The name of the lab that uniquely identifies it within containing lab plan. Used in resource URIs.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLabOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabArgs)(nil)).Elem()
}

// The lab resource.
type LookupLabResultOutput struct{ *pulumi.OutputState }

func (LookupLabResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResult)(nil)).Elem()
}

func (o LookupLabResultOutput) ToLookupLabResultOutput() LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ToLookupLabResultOutputWithContext(ctx context.Context) LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLabResult] {
	return pulumix.Output[LookupLabResult]{
		OutputState: o.OutputState,
	}
}

// The resource auto shutdown configuration for the lab. This controls whether actions are taken on resources that are sitting idle.
func (o LookupLabResultOutput) AutoShutdownProfile() AutoShutdownProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) AutoShutdownProfileResponse { return v.AutoShutdownProfile }).(AutoShutdownProfileResponseOutput)
}

// The connection profile for the lab. This controls settings such as web access to lab resources or whether RDP or SSH ports are open.
func (o LookupLabResultOutput) ConnectionProfile() ConnectionProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) ConnectionProfileResponse { return v.ConnectionProfile }).(ConnectionProfileResponseOutput)
}

// The description of the lab.
func (o LookupLabResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLabResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the lab plan. Used during resource creation to provide defaults and acts as a permission container when creating a lab via labs.azure.com. Setting a labPlanId on an existing lab provides organization..
func (o LookupLabResultOutput) LabPlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.LabPlanId }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupLabResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLabResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Name }).(pulumi.StringOutput)
}

// The network profile for the lab, typically applied via a lab plan. This profile cannot be modified once a lab has been created.
func (o LookupLabResultOutput) NetworkProfile() LabNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupLabResult) *LabNetworkProfileResponse { return v.NetworkProfile }).(LabNetworkProfileResponsePtrOutput)
}

// Current provisioning state of the lab.
func (o LookupLabResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The lab user list management profile.
func (o LookupLabResultOutput) RosterProfile() RosterProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupLabResult) *RosterProfileResponse { return v.RosterProfile }).(RosterProfileResponsePtrOutput)
}

// The lab security profile.
func (o LookupLabResultOutput) SecurityProfile() SecurityProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponseOutput)
}

// The lab state.
func (o LookupLabResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.State }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the lab.
func (o LookupLabResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLabResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupLabResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The title of the lab.
func (o LookupLabResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLabResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Type }).(pulumi.StringOutput)
}

// The profile used for creating lab virtual machines.
func (o LookupLabResultOutput) VirtualMachineProfile() VirtualMachineProfileResponseOutput {
	return o.ApplyT(func(v LookupLabResult) VirtualMachineProfileResponse { return v.VirtualMachineProfile }).(VirtualMachineProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResultOutput{})
}
