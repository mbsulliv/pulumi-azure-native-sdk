// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230815preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A member of the Fleet. It contains a reference to an existing Kubernetes cluster on Azure.
type FleetMember struct {
	pulumi.CustomResourceState

	// The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{clusterName}'.
	ClusterResourceId pulumi.StringOutput `pulumi:"clusterResourceId"`
	// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// The group this member belongs to for multi-cluster update management.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFleetMember registers a new resource with the given unique name, arguments, and options.
func NewFleetMember(ctx *pulumi.Context,
	name string, args *FleetMemberArgs, opts ...pulumi.ResourceOption) (*FleetMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceId'")
	}
	if args.FleetName == nil {
		return nil, errors.New("invalid value for required argument 'FleetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:FleetMember"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:FleetMember"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:FleetMember"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:FleetMember"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230315preview:FleetMember"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230615preview:FleetMember"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20231015:FleetMember"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FleetMember
	err := ctx.RegisterResource("azure-native:containerservice/v20230815preview:FleetMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleetMember gets an existing FleetMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleetMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetMemberState, opts ...pulumi.ResourceOption) (*FleetMember, error) {
	var resource FleetMember
	err := ctx.ReadResource("azure-native:containerservice/v20230815preview:FleetMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FleetMember resources.
type fleetMemberState struct {
}

type FleetMemberState struct {
}

func (FleetMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetMemberState)(nil)).Elem()
}

type fleetMemberArgs struct {
	// The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{clusterName}'.
	ClusterResourceId string `pulumi:"clusterResourceId"`
	// The name of the Fleet member resource.
	FleetMemberName *string `pulumi:"fleetMemberName"`
	// The name of the Fleet resource.
	FleetName string `pulumi:"fleetName"`
	// The group this member belongs to for multi-cluster update management.
	Group *string `pulumi:"group"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a FleetMember resource.
type FleetMemberArgs struct {
	// The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{clusterName}'.
	ClusterResourceId pulumi.StringInput
	// The name of the Fleet member resource.
	FleetMemberName pulumi.StringPtrInput
	// The name of the Fleet resource.
	FleetName pulumi.StringInput
	// The group this member belongs to for multi-cluster update management.
	Group pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (FleetMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetMemberArgs)(nil)).Elem()
}

type FleetMemberInput interface {
	pulumi.Input

	ToFleetMemberOutput() FleetMemberOutput
	ToFleetMemberOutputWithContext(ctx context.Context) FleetMemberOutput
}

func (*FleetMember) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetMember)(nil)).Elem()
}

func (i *FleetMember) ToFleetMemberOutput() FleetMemberOutput {
	return i.ToFleetMemberOutputWithContext(context.Background())
}

func (i *FleetMember) ToFleetMemberOutputWithContext(ctx context.Context) FleetMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetMemberOutput)
}

func (i *FleetMember) ToOutput(ctx context.Context) pulumix.Output[*FleetMember] {
	return pulumix.Output[*FleetMember]{
		OutputState: i.ToFleetMemberOutputWithContext(ctx).OutputState,
	}
}

type FleetMemberOutput struct{ *pulumi.OutputState }

func (FleetMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetMember)(nil)).Elem()
}

func (o FleetMemberOutput) ToFleetMemberOutput() FleetMemberOutput {
	return o
}

func (o FleetMemberOutput) ToFleetMemberOutputWithContext(ctx context.Context) FleetMemberOutput {
	return o
}

func (o FleetMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*FleetMember] {
	return pulumix.Output[*FleetMember]{
		OutputState: o.OutputState,
	}
}

// The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{clusterName}'.
func (o FleetMemberOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.ClusterResourceId }).(pulumi.StringOutput)
}

// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o FleetMemberOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// The group this member belongs to for multi-cluster update management.
func (o FleetMemberOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o FleetMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o FleetMemberOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FleetMemberOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FleetMember) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FleetMemberOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetMemberOutput{})
}
