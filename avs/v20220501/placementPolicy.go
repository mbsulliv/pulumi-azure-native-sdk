// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A vSphere Distributed Resource Scheduler (DRS) placement policy
type PlacementPolicy struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// placement policy properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPlacementPolicy registers a new resource with the given unique name, arguments, and options.
func NewPlacementPolicy(ctx *pulumi.Context,
	name string, args *PlacementPolicyArgs, opts ...pulumi.ResourceOption) (*PlacementPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:PlacementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:PlacementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:PlacementPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PlacementPolicy
	err := ctx.RegisterResource("azure-native:avs/v20220501:PlacementPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlacementPolicy gets an existing PlacementPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlacementPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlacementPolicyState, opts ...pulumi.ResourceOption) (*PlacementPolicy, error) {
	var resource PlacementPolicy
	err := ctx.ReadResource("azure-native:avs/v20220501:PlacementPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlacementPolicy resources.
type placementPolicyState struct {
}

type PlacementPolicyState struct {
}

func (PlacementPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementPolicyState)(nil)).Elem()
}

type placementPolicyArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
	PlacementPolicyName *string `pulumi:"placementPolicyName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// placement policy properties
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PlacementPolicy resource.
type PlacementPolicyArgs struct {
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringInput
	// Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
	PlacementPolicyName pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// placement policy properties
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (PlacementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementPolicyArgs)(nil)).Elem()
}

type PlacementPolicyInput interface {
	pulumi.Input

	ToPlacementPolicyOutput() PlacementPolicyOutput
	ToPlacementPolicyOutputWithContext(ctx context.Context) PlacementPolicyOutput
}

func (*PlacementPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementPolicy)(nil)).Elem()
}

func (i *PlacementPolicy) ToPlacementPolicyOutput() PlacementPolicyOutput {
	return i.ToPlacementPolicyOutputWithContext(context.Background())
}

func (i *PlacementPolicy) ToPlacementPolicyOutputWithContext(ctx context.Context) PlacementPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementPolicyOutput)
}

type PlacementPolicyOutput struct{ *pulumi.OutputState }

func (PlacementPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementPolicy)(nil)).Elem()
}

func (o PlacementPolicyOutput) ToPlacementPolicyOutput() PlacementPolicyOutput {
	return o
}

func (o PlacementPolicyOutput) ToPlacementPolicyOutputWithContext(ctx context.Context) PlacementPolicyOutput {
	return o
}

// Resource name.
func (o PlacementPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PlacementPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// placement policy properties
func (o PlacementPolicyOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *PlacementPolicy) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource type.
func (o PlacementPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PlacementPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PlacementPolicyOutput{})
}
