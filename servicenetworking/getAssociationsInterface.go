// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicenetworking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Association
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2022-10-01-preview.
func LookupAssociationsInterface(ctx *pulumi.Context, args *LookupAssociationsInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupAssociationsInterfaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAssociationsInterfaceResult
	err := ctx.Invoke("azure-native:servicenetworking:getAssociationsInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssociationsInterfaceArgs struct {
	// Name of Association
	AssociationName string `pulumi:"associationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// traffic controller name for path
	TrafficControllerName string `pulumi:"trafficControllerName"`
}

// Association Subresource of Traffic Controller
type LookupAssociationsInterfaceResult struct {
	// Association Type
	AssociationType string `pulumi:"associationType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning State of Traffic Controller Association Resource
	ProvisioningState string `pulumi:"provisioningState"`
	// Association Subnet
	Subnet *AssociationSubnetResponse `pulumi:"subnet"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAssociationsInterfaceOutput(ctx *pulumi.Context, args LookupAssociationsInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupAssociationsInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssociationsInterfaceResult, error) {
			args := v.(LookupAssociationsInterfaceArgs)
			r, err := LookupAssociationsInterface(ctx, &args, opts...)
			var s LookupAssociationsInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssociationsInterfaceResultOutput)
}

type LookupAssociationsInterfaceOutputArgs struct {
	// Name of Association
	AssociationName pulumi.StringInput `pulumi:"associationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// traffic controller name for path
	TrafficControllerName pulumi.StringInput `pulumi:"trafficControllerName"`
}

func (LookupAssociationsInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssociationsInterfaceArgs)(nil)).Elem()
}

// Association Subresource of Traffic Controller
type LookupAssociationsInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupAssociationsInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssociationsInterfaceResult)(nil)).Elem()
}

func (o LookupAssociationsInterfaceResultOutput) ToLookupAssociationsInterfaceResultOutput() LookupAssociationsInterfaceResultOutput {
	return o
}

func (o LookupAssociationsInterfaceResultOutput) ToLookupAssociationsInterfaceResultOutputWithContext(ctx context.Context) LookupAssociationsInterfaceResultOutput {
	return o
}

func (o LookupAssociationsInterfaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAssociationsInterfaceResult] {
	return pulumix.Output[LookupAssociationsInterfaceResult]{
		OutputState: o.OutputState,
	}
}

// Association Type
func (o LookupAssociationsInterfaceResultOutput) AssociationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.AssociationType }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAssociationsInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupAssociationsInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAssociationsInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning State of Traffic Controller Association Resource
func (o LookupAssociationsInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Association Subnet
func (o LookupAssociationsInterfaceResultOutput) Subnet() AssociationSubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) *AssociationSubnetResponse { return v.Subnet }).(AssociationSubnetResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAssociationsInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAssociationsInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAssociationsInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationsInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssociationsInterfaceResultOutput{})
}
