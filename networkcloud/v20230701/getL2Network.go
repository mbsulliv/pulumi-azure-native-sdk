// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of the provided layer 2 (L2) network.
func LookupL2Network(ctx *pulumi.Context, args *LookupL2NetworkArgs, opts ...pulumi.InvokeOption) (*LookupL2NetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupL2NetworkResult
	err := ctx.Invoke("azure-native:networkcloud/v20230701:getL2Network", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupL2NetworkArgs struct {
	// The name of the L2 network.
	L2NetworkName string `pulumi:"l2NetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupL2NetworkResult struct {
	// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
	AssociatedResourceIds []string `pulumi:"associatedResourceIds"`
	// The resource ID of the Network Cloud cluster this L2 network is associated with.
	ClusterId string `pulumi:"clusterId"`
	// The more detailed status of the L2 network.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource ID(s) that are associated with this L2 network.
	HybridAksClustersAssociatedIds []string `pulumi:"hybridAksClustersAssociatedIds"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType *string `pulumi:"hybridAksPluginType"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The default interface name for this L2 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName *string `pulumi:"interfaceName"`
	// The resource ID of the Network Fabric l2IsolationDomain.
	L2IsolationDomainId string `pulumi:"l2IsolationDomainId"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the L2 network.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource ID(s), excluding any Hybrid AKS virtual machines, that are currently using this L2 network.
	VirtualMachinesAssociatedIds []string `pulumi:"virtualMachinesAssociatedIds"`
}

// Defaults sets the appropriate defaults for LookupL2NetworkResult
func (val *LookupL2NetworkResult) Defaults() *LookupL2NetworkResult {
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

func LookupL2NetworkOutput(ctx *pulumi.Context, args LookupL2NetworkOutputArgs, opts ...pulumi.InvokeOption) LookupL2NetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupL2NetworkResult, error) {
			args := v.(LookupL2NetworkArgs)
			r, err := LookupL2Network(ctx, &args, opts...)
			var s LookupL2NetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupL2NetworkResultOutput)
}

type LookupL2NetworkOutputArgs struct {
	// The name of the L2 network.
	L2NetworkName pulumi.StringInput `pulumi:"l2NetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupL2NetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2NetworkArgs)(nil)).Elem()
}

type LookupL2NetworkResultOutput struct{ *pulumi.OutputState }

func (LookupL2NetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2NetworkResult)(nil)).Elem()
}

func (o LookupL2NetworkResultOutput) ToLookupL2NetworkResultOutput() LookupL2NetworkResultOutput {
	return o
}

func (o LookupL2NetworkResultOutput) ToLookupL2NetworkResultOutputWithContext(ctx context.Context) LookupL2NetworkResultOutput {
	return o
}

func (o LookupL2NetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupL2NetworkResult] {
	return pulumix.Output[LookupL2NetworkResult]{
		OutputState: o.OutputState,
	}
}

// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
func (o LookupL2NetworkResultOutput) AssociatedResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) []string { return v.AssociatedResourceIds }).(pulumi.StringArrayOutput)
}

// The resource ID of the Network Cloud cluster this L2 network is associated with.
func (o LookupL2NetworkResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The more detailed status of the L2 network.
func (o LookupL2NetworkResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupL2NetworkResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupL2NetworkResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource ID(s) that are associated with this L2 network.
func (o LookupL2NetworkResultOutput) HybridAksClustersAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) []string { return v.HybridAksClustersAssociatedIds }).(pulumi.StringArrayOutput)
}

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
func (o LookupL2NetworkResultOutput) HybridAksPluginType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) *string { return v.HybridAksPluginType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupL2NetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The default interface name for this L2 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
func (o LookupL2NetworkResultOutput) InterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) *string { return v.InterfaceName }).(pulumi.StringPtrOutput)
}

// The resource ID of the Network Fabric l2IsolationDomain.
func (o LookupL2NetworkResultOutput) L2IsolationDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.L2IsolationDomainId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupL2NetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupL2NetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the L2 network.
func (o LookupL2NetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupL2NetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupL2NetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupL2NetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource ID(s), excluding any Hybrid AKS virtual machines, that are currently using this L2 network.
func (o LookupL2NetworkResultOutput) VirtualMachinesAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupL2NetworkResult) []string { return v.VirtualMachinesAssociatedIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupL2NetworkResultOutput{})
}
