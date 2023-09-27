// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a InternalNetworks.
func LookupInternalNetwork(ctx *pulumi.Context, args *LookupInternalNetworkArgs, opts ...pulumi.InvokeOption) (*LookupInternalNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInternalNetworkResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230201preview:getInternalNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupInternalNetworkArgs struct {
	// Name of the InternalNetwork
	InternalNetworkName string `pulumi:"internalNetworkName"`
	// Name of the L3IsolationDomain
	L3IsolationDomainName string `pulumi:"l3IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the InternalNetwork item.
type LookupInternalNetworkResult struct {
	// Administrative state of the InternalNetwork. Example: Enabled | Disabled.
	AdministrativeState string `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// List of resources the BFD for BGP is disabled on. Can be either entire NetworkFabric or NetworkRack.
	BfdDisabledOnResources []string `pulumi:"bfdDisabledOnResources"`
	// List of resources the BFD of StaticRoutes is disabled on. Can be either entire NetworkFabric or NetworkRack.
	BfdForStaticRoutesDisabledOnResources []string `pulumi:"bfdForStaticRoutesDisabledOnResources"`
	// BGP configuration properties
	BgpConfiguration *BgpConfigurationResponse `pulumi:"bgpConfiguration"`
	// List of resources the BGP is disabled on. Can be either entire NetworkFabric or NetworkRack.
	BgpDisabledOnResources []string `pulumi:"bgpDisabledOnResources"`
	// List with object connected IPv4 Subnets.
	ConnectedIPv4Subnets []ConnectedSubnetResponse `pulumi:"connectedIPv4Subnets"`
	// List with object connected IPv6 Subnets.
	ConnectedIPv6Subnets []ConnectedSubnetResponse `pulumi:"connectedIPv6Subnets"`
	// List of resources the InternalNetwork is disabled on. Can be either entire NetworkFabric or NetworkRack.
	DisabledOnResources []string `pulumi:"disabledOnResources"`
	// ARM resource ID of importRoutePolicy.
	ExportRoutePolicyId *string `pulumi:"exportRoutePolicyId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// ARM resource ID of importRoutePolicy.
	ImportRoutePolicyId *string `pulumi:"importRoutePolicyId"`
	// Maximum transmission unit. Default value is 1500.
	Mtu *int `pulumi:"mtu"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Static Route Configuration properties.
	StaticRouteConfiguration *StaticRouteConfigurationResponse `pulumi:"staticRouteConfiguration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Vlan identifier. Example: 1001.
	VlanId int `pulumi:"vlanId"`
}

// Defaults sets the appropriate defaults for LookupInternalNetworkResult
func (val *LookupInternalNetworkResult) Defaults() *LookupInternalNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BgpConfiguration = tmp.BgpConfiguration.Defaults()

	if tmp.Mtu == nil {
		mtu_ := 1500
		tmp.Mtu = &mtu_
	}
	return &tmp
}

func LookupInternalNetworkOutput(ctx *pulumi.Context, args LookupInternalNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupInternalNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInternalNetworkResult, error) {
			args := v.(LookupInternalNetworkArgs)
			r, err := LookupInternalNetwork(ctx, &args, opts...)
			var s LookupInternalNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInternalNetworkResultOutput)
}

type LookupInternalNetworkOutputArgs struct {
	// Name of the InternalNetwork
	InternalNetworkName pulumi.StringInput `pulumi:"internalNetworkName"`
	// Name of the L3IsolationDomain
	L3IsolationDomainName pulumi.StringInput `pulumi:"l3IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInternalNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternalNetworkArgs)(nil)).Elem()
}

// Defines the InternalNetwork item.
type LookupInternalNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupInternalNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternalNetworkResult)(nil)).Elem()
}

func (o LookupInternalNetworkResultOutput) ToLookupInternalNetworkResultOutput() LookupInternalNetworkResultOutput {
	return o
}

func (o LookupInternalNetworkResultOutput) ToLookupInternalNetworkResultOutputWithContext(ctx context.Context) LookupInternalNetworkResultOutput {
	return o
}

func (o LookupInternalNetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInternalNetworkResult] {
	return pulumix.Output[LookupInternalNetworkResult]{
		OutputState: o.OutputState,
	}
}

// Administrative state of the InternalNetwork. Example: Enabled | Disabled.
func (o LookupInternalNetworkResultOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) string { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupInternalNetworkResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// List of resources the BFD for BGP is disabled on. Can be either entire NetworkFabric or NetworkRack.
func (o LookupInternalNetworkResultOutput) BfdDisabledOnResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) []string { return v.BfdDisabledOnResources }).(pulumi.StringArrayOutput)
}

// List of resources the BFD of StaticRoutes is disabled on. Can be either entire NetworkFabric or NetworkRack.
func (o LookupInternalNetworkResultOutput) BfdForStaticRoutesDisabledOnResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) []string { return v.BfdForStaticRoutesDisabledOnResources }).(pulumi.StringArrayOutput)
}

// BGP configuration properties
func (o LookupInternalNetworkResultOutput) BgpConfiguration() BgpConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) *BgpConfigurationResponse { return v.BgpConfiguration }).(BgpConfigurationResponsePtrOutput)
}

// List of resources the BGP is disabled on. Can be either entire NetworkFabric or NetworkRack.
func (o LookupInternalNetworkResultOutput) BgpDisabledOnResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) []string { return v.BgpDisabledOnResources }).(pulumi.StringArrayOutput)
}

// List with object connected IPv4 Subnets.
func (o LookupInternalNetworkResultOutput) ConnectedIPv4Subnets() ConnectedSubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) []ConnectedSubnetResponse { return v.ConnectedIPv4Subnets }).(ConnectedSubnetResponseArrayOutput)
}

// List with object connected IPv6 Subnets.
func (o LookupInternalNetworkResultOutput) ConnectedIPv6Subnets() ConnectedSubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) []ConnectedSubnetResponse { return v.ConnectedIPv6Subnets }).(ConnectedSubnetResponseArrayOutput)
}

// List of resources the InternalNetwork is disabled on. Can be either entire NetworkFabric or NetworkRack.
func (o LookupInternalNetworkResultOutput) DisabledOnResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) []string { return v.DisabledOnResources }).(pulumi.StringArrayOutput)
}

// ARM resource ID of importRoutePolicy.
func (o LookupInternalNetworkResultOutput) ExportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) *string { return v.ExportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupInternalNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// ARM resource ID of importRoutePolicy.
func (o LookupInternalNetworkResultOutput) ImportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) *string { return v.ImportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// Maximum transmission unit. Default value is 1500.
func (o LookupInternalNetworkResultOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) *int { return v.Mtu }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupInternalNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o LookupInternalNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Static Route Configuration properties.
func (o LookupInternalNetworkResultOutput) StaticRouteConfiguration() StaticRouteConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) *StaticRouteConfigurationResponse {
		return v.StaticRouteConfiguration
	}).(StaticRouteConfigurationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupInternalNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupInternalNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// Vlan identifier. Example: 1001.
func (o LookupInternalNetworkResultOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInternalNetworkResult) int { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternalNetworkResultOutput{})
}
