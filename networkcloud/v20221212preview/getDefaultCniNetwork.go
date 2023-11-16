// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221212preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of the provided default CNI network.
func LookupDefaultCniNetwork(ctx *pulumi.Context, args *LookupDefaultCniNetworkArgs, opts ...pulumi.InvokeOption) (*LookupDefaultCniNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultCniNetworkResult
	err := ctx.Invoke("azure-native:networkcloud/v20221212preview:getDefaultCniNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDefaultCniNetworkArgs struct {
	// The name of the default CNI network.
	DefaultCniNetworkName string `pulumi:"defaultCniNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupDefaultCniNetworkResult struct {
	// The resource ID of the Network Cloud cluster this default CNI network is associated with.
	ClusterId string `pulumi:"clusterId"`
	// The autonomous system number that the fabric expects to peer with, derived from the associated L3 isolation domain.
	CniAsNumber float64 `pulumi:"cniAsNumber"`
	// The Calico BGP configuration.
	CniBgpConfiguration *CniBgpConfigurationResponse `pulumi:"cniBgpConfiguration"`
	// The more detailed status of the default CNI network.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The L3 isolation fabric BGP peering connectivity information necessary for BGP peering the Hybrid AKS Cluster with the switch fabric.
	FabricBgpPeers []BgpPeerResponse `pulumi:"fabricBgpPeers"`
	// The list of Hybrid AKS cluster resource ID(s) that are associated with this default CNI network.
	HybridAksClustersAssociatedIds []string `pulumi:"hybridAksClustersAssociatedIds"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the interface that will be present in the virtual machine to represent this network.
	InterfaceName string `pulumi:"interfaceName"`
	// The type of the IP address allocation.
	IpAllocationType *string `pulumi:"ipAllocationType"`
	// The IPV4 prefix (CIDR) assigned to this default CNI network. It is required when the IP allocation type
	// is IPV4 or DualStack.
	Ipv4ConnectedPrefix *string `pulumi:"ipv4ConnectedPrefix"`
	// The IPV6 prefix (CIDR) assigned to this default CNI network. It is required when the IP allocation type
	// is IPV6 or DualStack.
	Ipv6ConnectedPrefix *string `pulumi:"ipv6ConnectedPrefix"`
	// The resource ID of the Network Fabric l3IsolationDomain.
	L3IsolationDomainId string `pulumi:"l3IsolationDomainId"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the default CNI network.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The VLAN from the l3IsolationDomain that is used for this network.
	Vlan float64 `pulumi:"vlan"`
}

// Defaults sets the appropriate defaults for LookupDefaultCniNetworkResult
func (val *LookupDefaultCniNetworkResult) Defaults() *LookupDefaultCniNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.IpAllocationType == nil {
		ipAllocationType_ := "DualStack"
		tmp.IpAllocationType = &ipAllocationType_
	}
	return &tmp
}

func LookupDefaultCniNetworkOutput(ctx *pulumi.Context, args LookupDefaultCniNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultCniNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultCniNetworkResult, error) {
			args := v.(LookupDefaultCniNetworkArgs)
			r, err := LookupDefaultCniNetwork(ctx, &args, opts...)
			var s LookupDefaultCniNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultCniNetworkResultOutput)
}

type LookupDefaultCniNetworkOutputArgs struct {
	// The name of the default CNI network.
	DefaultCniNetworkName pulumi.StringInput `pulumi:"defaultCniNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDefaultCniNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultCniNetworkArgs)(nil)).Elem()
}

type LookupDefaultCniNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultCniNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultCniNetworkResult)(nil)).Elem()
}

func (o LookupDefaultCniNetworkResultOutput) ToLookupDefaultCniNetworkResultOutput() LookupDefaultCniNetworkResultOutput {
	return o
}

func (o LookupDefaultCniNetworkResultOutput) ToLookupDefaultCniNetworkResultOutputWithContext(ctx context.Context) LookupDefaultCniNetworkResultOutput {
	return o
}

// The resource ID of the Network Cloud cluster this default CNI network is associated with.
func (o LookupDefaultCniNetworkResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The autonomous system number that the fabric expects to peer with, derived from the associated L3 isolation domain.
func (o LookupDefaultCniNetworkResultOutput) CniAsNumber() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) float64 { return v.CniAsNumber }).(pulumi.Float64Output)
}

// The Calico BGP configuration.
func (o LookupDefaultCniNetworkResultOutput) CniBgpConfiguration() CniBgpConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) *CniBgpConfigurationResponse { return v.CniBgpConfiguration }).(CniBgpConfigurationResponsePtrOutput)
}

// The more detailed status of the default CNI network.
func (o LookupDefaultCniNetworkResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupDefaultCniNetworkResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupDefaultCniNetworkResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The L3 isolation fabric BGP peering connectivity information necessary for BGP peering the Hybrid AKS Cluster with the switch fabric.
func (o LookupDefaultCniNetworkResultOutput) FabricBgpPeers() BgpPeerResponseArrayOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) []BgpPeerResponse { return v.FabricBgpPeers }).(BgpPeerResponseArrayOutput)
}

// The list of Hybrid AKS cluster resource ID(s) that are associated with this default CNI network.
func (o LookupDefaultCniNetworkResultOutput) HybridAksClustersAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) []string { return v.HybridAksClustersAssociatedIds }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDefaultCniNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the interface that will be present in the virtual machine to represent this network.
func (o LookupDefaultCniNetworkResultOutput) InterfaceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.InterfaceName }).(pulumi.StringOutput)
}

// The type of the IP address allocation.
func (o LookupDefaultCniNetworkResultOutput) IpAllocationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) *string { return v.IpAllocationType }).(pulumi.StringPtrOutput)
}

// The IPV4 prefix (CIDR) assigned to this default CNI network. It is required when the IP allocation type
// is IPV4 or DualStack.
func (o LookupDefaultCniNetworkResultOutput) Ipv4ConnectedPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) *string { return v.Ipv4ConnectedPrefix }).(pulumi.StringPtrOutput)
}

// The IPV6 prefix (CIDR) assigned to this default CNI network. It is required when the IP allocation type
// is IPV6 or DualStack.
func (o LookupDefaultCniNetworkResultOutput) Ipv6ConnectedPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) *string { return v.Ipv6ConnectedPrefix }).(pulumi.StringPtrOutput)
}

// The resource ID of the Network Fabric l3IsolationDomain.
func (o LookupDefaultCniNetworkResultOutput) L3IsolationDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.L3IsolationDomainId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDefaultCniNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDefaultCniNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the default CNI network.
func (o LookupDefaultCniNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDefaultCniNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDefaultCniNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDefaultCniNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VLAN from the l3IsolationDomain that is used for this network.
func (o LookupDefaultCniNetworkResultOutput) Vlan() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDefaultCniNetworkResult) float64 { return v.Vlan }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultCniNetworkResultOutput{})
}
