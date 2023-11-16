// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of the provided trunked network.
func LookupTrunkedNetwork(ctx *pulumi.Context, args *LookupTrunkedNetworkArgs, opts ...pulumi.InvokeOption) (*LookupTrunkedNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTrunkedNetworkResult
	err := ctx.Invoke("azure-native:networkcloud/v20230701:getTrunkedNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTrunkedNetworkArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the trunked network.
	TrunkedNetworkName string `pulumi:"trunkedNetworkName"`
}

type LookupTrunkedNetworkResult struct {
	// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
	AssociatedResourceIds []string `pulumi:"associatedResourceIds"`
	// The resource ID of the Network Cloud cluster this trunked network is associated with.
	ClusterId string `pulumi:"clusterId"`
	// The more detailed status of the trunked network.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this trunked network.
	HybridAksClustersAssociatedIds []string `pulumi:"hybridAksClustersAssociatedIds"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType *string `pulumi:"hybridAksPluginType"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The default interface name for this trunked network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName *string `pulumi:"interfaceName"`
	// The list of resource IDs representing the Network Fabric isolation domains. It can be any combination of l2IsolationDomain and l3IsolationDomain resources.
	IsolationDomainIds []string `pulumi:"isolationDomainIds"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the trunked network.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this trunked network.
	VirtualMachinesAssociatedIds []string `pulumi:"virtualMachinesAssociatedIds"`
	// The list of vlans that are selected from the isolation domains for trunking.
	Vlans []float64 `pulumi:"vlans"`
}

// Defaults sets the appropriate defaults for LookupTrunkedNetworkResult
func (val *LookupTrunkedNetworkResult) Defaults() *LookupTrunkedNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.HybridAksPluginType == nil {
		hybridAksPluginType_ := "SRIOV"
		tmp.HybridAksPluginType = &hybridAksPluginType_
	}
	return &tmp
}

func LookupTrunkedNetworkOutput(ctx *pulumi.Context, args LookupTrunkedNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupTrunkedNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrunkedNetworkResult, error) {
			args := v.(LookupTrunkedNetworkArgs)
			r, err := LookupTrunkedNetwork(ctx, &args, opts...)
			var s LookupTrunkedNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrunkedNetworkResultOutput)
}

type LookupTrunkedNetworkOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the trunked network.
	TrunkedNetworkName pulumi.StringInput `pulumi:"trunkedNetworkName"`
}

func (LookupTrunkedNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrunkedNetworkArgs)(nil)).Elem()
}

type LookupTrunkedNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupTrunkedNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrunkedNetworkResult)(nil)).Elem()
}

func (o LookupTrunkedNetworkResultOutput) ToLookupTrunkedNetworkResultOutput() LookupTrunkedNetworkResultOutput {
	return o
}

func (o LookupTrunkedNetworkResultOutput) ToLookupTrunkedNetworkResultOutputWithContext(ctx context.Context) LookupTrunkedNetworkResultOutput {
	return o
}

// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
func (o LookupTrunkedNetworkResultOutput) AssociatedResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) []string { return v.AssociatedResourceIds }).(pulumi.StringArrayOutput)
}

// The resource ID of the Network Cloud cluster this trunked network is associated with.
func (o LookupTrunkedNetworkResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The more detailed status of the trunked network.
func (o LookupTrunkedNetworkResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupTrunkedNetworkResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupTrunkedNetworkResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this trunked network.
func (o LookupTrunkedNetworkResultOutput) HybridAksClustersAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) []string { return v.HybridAksClustersAssociatedIds }).(pulumi.StringArrayOutput)
}

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
func (o LookupTrunkedNetworkResultOutput) HybridAksPluginType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) *string { return v.HybridAksPluginType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupTrunkedNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The default interface name for this trunked network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
func (o LookupTrunkedNetworkResultOutput) InterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) *string { return v.InterfaceName }).(pulumi.StringPtrOutput)
}

// The list of resource IDs representing the Network Fabric isolation domains. It can be any combination of l2IsolationDomain and l3IsolationDomain resources.
func (o LookupTrunkedNetworkResultOutput) IsolationDomainIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) []string { return v.IsolationDomainIds }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o LookupTrunkedNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTrunkedNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the trunked network.
func (o LookupTrunkedNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupTrunkedNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupTrunkedNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTrunkedNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this trunked network.
func (o LookupTrunkedNetworkResultOutput) VirtualMachinesAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) []string { return v.VirtualMachinesAssociatedIds }).(pulumi.StringArrayOutput)
}

// The list of vlans that are selected from the isolation domains for trunking.
func (o LookupTrunkedNetworkResultOutput) Vlans() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v LookupTrunkedNetworkResult) []float64 { return v.Vlans }).(pulumi.Float64ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrunkedNetworkResultOutput{})
}
