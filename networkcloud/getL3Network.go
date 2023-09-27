// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of the provided layer 3 (L3) network.
// Azure REST API version: 2023-05-01-preview.
func LookupL3Network(ctx *pulumi.Context, args *LookupL3NetworkArgs, opts ...pulumi.InvokeOption) (*LookupL3NetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupL3NetworkResult
	err := ctx.Invoke("azure-native:networkcloud:getL3Network", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupL3NetworkArgs struct {
	// The name of the L3 network.
	L3NetworkName string `pulumi:"l3NetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupL3NetworkResult struct {
	// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
	AssociatedResourceIds []string `pulumi:"associatedResourceIds"`
	// The resource ID of the Network Cloud cluster this L3 network is associated with.
	ClusterId string `pulumi:"clusterId"`
	// The more detailed status of the L3 network.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this L3 network.
	HybridAksClustersAssociatedIds []string `pulumi:"hybridAksClustersAssociatedIds"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The indicator of whether or not to disable IPAM allocation on the network attachment definition injected into the Hybrid AKS Cluster.
	HybridAksIpamEnabled *string `pulumi:"hybridAksIpamEnabled"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType *string `pulumi:"hybridAksPluginType"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The default interface name for this L3 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName *string `pulumi:"interfaceName"`
	// The type of the IP address allocation, defaulted to "DualStack".
	IpAllocationType *string `pulumi:"ipAllocationType"`
	// The IPV4 prefix (CIDR) assigned to this L3 network. Required when the IP allocation type
	// is IPV4 or DualStack.
	Ipv4ConnectedPrefix *string `pulumi:"ipv4ConnectedPrefix"`
	// The IPV6 prefix (CIDR) assigned to this L3 network. Required when the IP allocation type
	// is IPV6 or DualStack.
	Ipv6ConnectedPrefix *string `pulumi:"ipv6ConnectedPrefix"`
	// The resource ID of the Network Fabric l3IsolationDomain.
	L3IsolationDomainId string `pulumi:"l3IsolationDomainId"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the L3 network.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this L3 network.
	VirtualMachinesAssociatedIds []string `pulumi:"virtualMachinesAssociatedIds"`
	// The VLAN from the l3IsolationDomain that is used for this network.
	Vlan float64 `pulumi:"vlan"`
}

// Defaults sets the appropriate defaults for LookupL3NetworkResult
func (val *LookupL3NetworkResult) Defaults() *LookupL3NetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.HybridAksIpamEnabled == nil {
		hybridAksIpamEnabled_ := "True"
		tmp.HybridAksIpamEnabled = &hybridAksIpamEnabled_
	}
	if tmp.HybridAksPluginType == nil {
		hybridAksPluginType_ := "SRIOV"
		tmp.HybridAksPluginType = &hybridAksPluginType_
	}
	if tmp.IpAllocationType == nil {
		ipAllocationType_ := "DualStack"
		tmp.IpAllocationType = &ipAllocationType_
	}
	return &tmp
}

func LookupL3NetworkOutput(ctx *pulumi.Context, args LookupL3NetworkOutputArgs, opts ...pulumi.InvokeOption) LookupL3NetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupL3NetworkResult, error) {
			args := v.(LookupL3NetworkArgs)
			r, err := LookupL3Network(ctx, &args, opts...)
			var s LookupL3NetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupL3NetworkResultOutput)
}

type LookupL3NetworkOutputArgs struct {
	// The name of the L3 network.
	L3NetworkName pulumi.StringInput `pulumi:"l3NetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupL3NetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL3NetworkArgs)(nil)).Elem()
}

type LookupL3NetworkResultOutput struct{ *pulumi.OutputState }

func (LookupL3NetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL3NetworkResult)(nil)).Elem()
}

func (o LookupL3NetworkResultOutput) ToLookupL3NetworkResultOutput() LookupL3NetworkResultOutput {
	return o
}

func (o LookupL3NetworkResultOutput) ToLookupL3NetworkResultOutputWithContext(ctx context.Context) LookupL3NetworkResultOutput {
	return o
}

func (o LookupL3NetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupL3NetworkResult] {
	return pulumix.Output[LookupL3NetworkResult]{
		OutputState: o.OutputState,
	}
}

// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
func (o LookupL3NetworkResultOutput) AssociatedResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) []string { return v.AssociatedResourceIds }).(pulumi.StringArrayOutput)
}

// The resource ID of the Network Cloud cluster this L3 network is associated with.
func (o LookupL3NetworkResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The more detailed status of the L3 network.
func (o LookupL3NetworkResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupL3NetworkResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupL3NetworkResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this L3 network.
func (o LookupL3NetworkResultOutput) HybridAksClustersAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) []string { return v.HybridAksClustersAssociatedIds }).(pulumi.StringArrayOutput)
}

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The indicator of whether or not to disable IPAM allocation on the network attachment definition injected into the Hybrid AKS Cluster.
func (o LookupL3NetworkResultOutput) HybridAksIpamEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) *string { return v.HybridAksIpamEnabled }).(pulumi.StringPtrOutput)
}

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
func (o LookupL3NetworkResultOutput) HybridAksPluginType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) *string { return v.HybridAksPluginType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupL3NetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The default interface name for this L3 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
func (o LookupL3NetworkResultOutput) InterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) *string { return v.InterfaceName }).(pulumi.StringPtrOutput)
}

// The type of the IP address allocation, defaulted to "DualStack".
func (o LookupL3NetworkResultOutput) IpAllocationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) *string { return v.IpAllocationType }).(pulumi.StringPtrOutput)
}

// The IPV4 prefix (CIDR) assigned to this L3 network. Required when the IP allocation type
// is IPV4 or DualStack.
func (o LookupL3NetworkResultOutput) Ipv4ConnectedPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) *string { return v.Ipv4ConnectedPrefix }).(pulumi.StringPtrOutput)
}

// The IPV6 prefix (CIDR) assigned to this L3 network. Required when the IP allocation type
// is IPV6 or DualStack.
func (o LookupL3NetworkResultOutput) Ipv6ConnectedPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) *string { return v.Ipv6ConnectedPrefix }).(pulumi.StringPtrOutput)
}

// The resource ID of the Network Fabric l3IsolationDomain.
func (o LookupL3NetworkResultOutput) L3IsolationDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.L3IsolationDomainId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupL3NetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupL3NetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the L3 network.
func (o LookupL3NetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupL3NetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupL3NetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupL3NetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this L3 network.
func (o LookupL3NetworkResultOutput) VirtualMachinesAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupL3NetworkResult) []string { return v.VirtualMachinesAssociatedIds }).(pulumi.StringArrayOutput)
}

// The VLAN from the l3IsolationDomain that is used for this network.
func (o LookupL3NetworkResultOutput) Vlan() pulumi.Float64Output {
	return o.ApplyT(func(v LookupL3NetworkResult) float64 { return v.Vlan }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupL3NetworkResultOutput{})
}
