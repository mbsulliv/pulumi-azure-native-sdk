// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a dedicated host group.
func LookupDedicatedHostGroup(ctx *pulumi.Context, args *LookupDedicatedHostGroupArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDedicatedHostGroupResult
	err := ctx.Invoke("azure-native:compute/v20230701:getDedicatedHostGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostGroupArgs struct {
	// The expand expression to apply on the operation. 'InstanceView' will retrieve the list of instance views of the dedicated hosts under the dedicated host group. 'UserData' is not supported for dedicated host group.
	Expand *string `pulumi:"expand"`
	// The name of the dedicated host group.
	HostGroupName string `pulumi:"hostGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the dedicated host group that the dedicated hosts should be assigned to. Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.
type LookupDedicatedHostGroupResult struct {
	// Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
	AdditionalCapabilities *DedicatedHostGroupPropertiesResponseAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// A list of references to all dedicated hosts in the dedicated host group.
	Hosts []SubResourceReadOnlyResponse `pulumi:"hosts"`
	// Resource Id
	Id string `pulumi:"id"`
	// The dedicated host group instance view, which has the list of instance view of the dedicated hosts under the dedicated host group.
	InstanceView DedicatedHostGroupInstanceViewResponse `pulumi:"instanceView"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount int `pulumi:"platformFaultDomainCount"`
	// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
	SupportAutomaticPlacement *bool `pulumi:"supportAutomaticPlacement"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones []string `pulumi:"zones"`
}

func LookupDedicatedHostGroupOutput(ctx *pulumi.Context, args LookupDedicatedHostGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedHostGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedHostGroupResult, error) {
			args := v.(LookupDedicatedHostGroupArgs)
			r, err := LookupDedicatedHostGroup(ctx, &args, opts...)
			var s LookupDedicatedHostGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedHostGroupResultOutput)
}

type LookupDedicatedHostGroupOutputArgs struct {
	// The expand expression to apply on the operation. 'InstanceView' will retrieve the list of instance views of the dedicated hosts under the dedicated host group. 'UserData' is not supported for dedicated host group.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the dedicated host group.
	HostGroupName pulumi.StringInput `pulumi:"hostGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDedicatedHostGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostGroupArgs)(nil)).Elem()
}

// Specifies information about the dedicated host group that the dedicated hosts should be assigned to. Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.
type LookupDedicatedHostGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedHostGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostGroupResult)(nil)).Elem()
}

func (o LookupDedicatedHostGroupResultOutput) ToLookupDedicatedHostGroupResultOutput() LookupDedicatedHostGroupResultOutput {
	return o
}

func (o LookupDedicatedHostGroupResultOutput) ToLookupDedicatedHostGroupResultOutputWithContext(ctx context.Context) LookupDedicatedHostGroupResultOutput {
	return o
}

// Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
func (o LookupDedicatedHostGroupResultOutput) AdditionalCapabilities() DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) *DedicatedHostGroupPropertiesResponseAdditionalCapabilities {
		return v.AdditionalCapabilities
	}).(DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput)
}

// A list of references to all dedicated hosts in the dedicated host group.
func (o LookupDedicatedHostGroupResultOutput) Hosts() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) []SubResourceReadOnlyResponse { return v.Hosts }).(SubResourceReadOnlyResponseArrayOutput)
}

// Resource Id
func (o LookupDedicatedHostGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The dedicated host group instance view, which has the list of instance view of the dedicated hosts under the dedicated host group.
func (o LookupDedicatedHostGroupResultOutput) InstanceView() DedicatedHostGroupInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) DedicatedHostGroupInstanceViewResponse { return v.InstanceView }).(DedicatedHostGroupInstanceViewResponseOutput)
}

// Resource location
func (o LookupDedicatedHostGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupDedicatedHostGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of fault domains that the host group can span.
func (o LookupDedicatedHostGroupResultOutput) PlatformFaultDomainCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) int { return v.PlatformFaultDomainCount }).(pulumi.IntOutput)
}

// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
func (o LookupDedicatedHostGroupResultOutput) SupportAutomaticPlacement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) *bool { return v.SupportAutomaticPlacement }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o LookupDedicatedHostGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupDedicatedHostGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
func (o LookupDedicatedHostGroupResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHostGroupResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedHostGroupResultOutput{})
}
