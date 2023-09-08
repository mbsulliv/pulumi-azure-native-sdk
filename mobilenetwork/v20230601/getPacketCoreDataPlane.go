// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified packet core data plane.
func LookupPacketCoreDataPlane(ctx *pulumi.Context, args *LookupPacketCoreDataPlaneArgs, opts ...pulumi.InvokeOption) (*LookupPacketCoreDataPlaneResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPacketCoreDataPlaneResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20230601:getPacketCoreDataPlane", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPacketCoreDataPlaneArgs struct {
	// The name of the packet core control plane.
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	// The name of the packet core data plane.
	PacketCoreDataPlaneName string `pulumi:"packetCoreDataPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Packet core data plane resource. Must be created in the same location as its parent packet core control plane.
type LookupPacketCoreDataPlaneResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the packet core data plane resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The user plane interface on the access network. For 5G networks, this is the N3 interface. For 4G networks, this is the S1-U interface.
	UserPlaneAccessInterface InterfacePropertiesResponse `pulumi:"userPlaneAccessInterface"`
}

func LookupPacketCoreDataPlaneOutput(ctx *pulumi.Context, args LookupPacketCoreDataPlaneOutputArgs, opts ...pulumi.InvokeOption) LookupPacketCoreDataPlaneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPacketCoreDataPlaneResult, error) {
			args := v.(LookupPacketCoreDataPlaneArgs)
			r, err := LookupPacketCoreDataPlane(ctx, &args, opts...)
			var s LookupPacketCoreDataPlaneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPacketCoreDataPlaneResultOutput)
}

type LookupPacketCoreDataPlaneOutputArgs struct {
	// The name of the packet core control plane.
	PacketCoreControlPlaneName pulumi.StringInput `pulumi:"packetCoreControlPlaneName"`
	// The name of the packet core data plane.
	PacketCoreDataPlaneName pulumi.StringInput `pulumi:"packetCoreDataPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPacketCoreDataPlaneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreDataPlaneArgs)(nil)).Elem()
}

// Packet core data plane resource. Must be created in the same location as its parent packet core control plane.
type LookupPacketCoreDataPlaneResultOutput struct{ *pulumi.OutputState }

func (LookupPacketCoreDataPlaneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreDataPlaneResult)(nil)).Elem()
}

func (o LookupPacketCoreDataPlaneResultOutput) ToLookupPacketCoreDataPlaneResultOutput() LookupPacketCoreDataPlaneResultOutput {
	return o
}

func (o LookupPacketCoreDataPlaneResultOutput) ToLookupPacketCoreDataPlaneResultOutputWithContext(ctx context.Context) LookupPacketCoreDataPlaneResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupPacketCoreDataPlaneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPacketCoreDataPlaneResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPacketCoreDataPlaneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the packet core data plane resource.
func (o LookupPacketCoreDataPlaneResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPacketCoreDataPlaneResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPacketCoreDataPlaneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPacketCoreDataPlaneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Type }).(pulumi.StringOutput)
}

// The user plane interface on the access network. For 5G networks, this is the N3 interface. For 4G networks, this is the S1-U interface.
func (o LookupPacketCoreDataPlaneResultOutput) UserPlaneAccessInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) InterfacePropertiesResponse { return v.UserPlaneAccessInterface }).(InterfacePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPacketCoreDataPlaneResultOutput{})
}