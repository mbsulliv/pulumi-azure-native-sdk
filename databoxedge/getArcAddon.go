// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a specific addon by name.
// Azure REST API version: 2022-03-01.
func LookupArcAddon(ctx *pulumi.Context, args *LookupArcAddonArgs, opts ...pulumi.InvokeOption) (*LookupArcAddonResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupArcAddonResult
	err := ctx.Invoke("azure-native:databoxedge:getArcAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArcAddonArgs struct {
	// The addon name.
	AddonName string `pulumi:"addonName"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The role name.
	RoleName string `pulumi:"roleName"`
}

// Arc Addon.
type LookupArcAddonResult struct {
	// Host OS supported by the Arc addon.
	HostPlatform string `pulumi:"hostPlatform"`
	// Platform where the runtime is hosted.
	HostPlatformType string `pulumi:"hostPlatformType"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// Addon type.
	// Expected value is 'ArcForKubernetes'.
	Kind string `pulumi:"kind"`
	// The object name.
	Name string `pulumi:"name"`
	// Addon Provisioning State
	ProvisioningState string `pulumi:"provisioningState"`
	// Arc resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Arc resource location
	ResourceLocation string `pulumi:"resourceLocation"`
	// Arc resource Name
	ResourceName string `pulumi:"resourceName"`
	// Arc resource subscription Id
	SubscriptionId string `pulumi:"subscriptionId"`
	// Metadata pertaining to creation and last modification of Addon
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// Arc resource version
	Version string `pulumi:"version"`
}

func LookupArcAddonOutput(ctx *pulumi.Context, args LookupArcAddonOutputArgs, opts ...pulumi.InvokeOption) LookupArcAddonResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArcAddonResult, error) {
			args := v.(LookupArcAddonArgs)
			r, err := LookupArcAddon(ctx, &args, opts...)
			var s LookupArcAddonResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArcAddonResultOutput)
}

type LookupArcAddonOutputArgs struct {
	// The addon name.
	AddonName pulumi.StringInput `pulumi:"addonName"`
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The role name.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (LookupArcAddonOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcAddonArgs)(nil)).Elem()
}

// Arc Addon.
type LookupArcAddonResultOutput struct{ *pulumi.OutputState }

func (LookupArcAddonResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcAddonResult)(nil)).Elem()
}

func (o LookupArcAddonResultOutput) ToLookupArcAddonResultOutput() LookupArcAddonResultOutput {
	return o
}

func (o LookupArcAddonResultOutput) ToLookupArcAddonResultOutputWithContext(ctx context.Context) LookupArcAddonResultOutput {
	return o
}

func (o LookupArcAddonResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupArcAddonResult] {
	return pulumix.Output[LookupArcAddonResult]{
		OutputState: o.OutputState,
	}
}

// Host OS supported by the Arc addon.
func (o LookupArcAddonResultOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.HostPlatform }).(pulumi.StringOutput)
}

// Platform where the runtime is hosted.
func (o LookupArcAddonResultOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.HostPlatformType }).(pulumi.StringOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupArcAddonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Id }).(pulumi.StringOutput)
}

// Addon type.
// Expected value is 'ArcForKubernetes'.
func (o LookupArcAddonResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o LookupArcAddonResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Name }).(pulumi.StringOutput)
}

// Addon Provisioning State
func (o LookupArcAddonResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Arc resource group name
func (o LookupArcAddonResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// Arc resource location
func (o LookupArcAddonResultOutput) ResourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ResourceLocation }).(pulumi.StringOutput)
}

// Arc resource Name
func (o LookupArcAddonResultOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ResourceName }).(pulumi.StringOutput)
}

// Arc resource subscription Id
func (o LookupArcAddonResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Addon
func (o LookupArcAddonResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupArcAddonResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupArcAddonResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Type }).(pulumi.StringOutput)
}

// Arc resource version
func (o LookupArcAddonResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArcAddonResultOutput{})
}
