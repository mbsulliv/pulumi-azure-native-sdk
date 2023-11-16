// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines a multi-stage process to perform update operations across members of a Fleet.
// Azure REST API version: 2023-08-15-preview.
//
// Other available API versions: 2023-10-15.
type FleetUpdateStrategy struct {
	pulumi.CustomResourceState

	// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the UpdateStrategy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines the update sequence of the clusters.
	Strategy UpdateRunStrategyResponseOutput `pulumi:"strategy"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFleetUpdateStrategy registers a new resource with the given unique name, arguments, and options.
func NewFleetUpdateStrategy(ctx *pulumi.Context,
	name string, args *FleetUpdateStrategyArgs, opts ...pulumi.ResourceOption) (*FleetUpdateStrategy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FleetName == nil {
		return nil, errors.New("invalid value for required argument 'FleetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Strategy == nil {
		return nil, errors.New("invalid value for required argument 'Strategy'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice/v20230815preview:FleetUpdateStrategy"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20231015:FleetUpdateStrategy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FleetUpdateStrategy
	err := ctx.RegisterResource("azure-native:containerservice:FleetUpdateStrategy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleetUpdateStrategy gets an existing FleetUpdateStrategy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleetUpdateStrategy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetUpdateStrategyState, opts ...pulumi.ResourceOption) (*FleetUpdateStrategy, error) {
	var resource FleetUpdateStrategy
	err := ctx.ReadResource("azure-native:containerservice:FleetUpdateStrategy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FleetUpdateStrategy resources.
type fleetUpdateStrategyState struct {
}

type FleetUpdateStrategyState struct {
}

func (FleetUpdateStrategyState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetUpdateStrategyState)(nil)).Elem()
}

type fleetUpdateStrategyArgs struct {
	// The name of the Fleet resource.
	FleetName string `pulumi:"fleetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Defines the update sequence of the clusters.
	Strategy UpdateRunStrategy `pulumi:"strategy"`
	// The name of the UpdateStrategy resource.
	UpdateStrategyName *string `pulumi:"updateStrategyName"`
}

// The set of arguments for constructing a FleetUpdateStrategy resource.
type FleetUpdateStrategyArgs struct {
	// The name of the Fleet resource.
	FleetName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Defines the update sequence of the clusters.
	Strategy UpdateRunStrategyInput
	// The name of the UpdateStrategy resource.
	UpdateStrategyName pulumi.StringPtrInput
}

func (FleetUpdateStrategyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetUpdateStrategyArgs)(nil)).Elem()
}

type FleetUpdateStrategyInput interface {
	pulumi.Input

	ToFleetUpdateStrategyOutput() FleetUpdateStrategyOutput
	ToFleetUpdateStrategyOutputWithContext(ctx context.Context) FleetUpdateStrategyOutput
}

func (*FleetUpdateStrategy) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetUpdateStrategy)(nil)).Elem()
}

func (i *FleetUpdateStrategy) ToFleetUpdateStrategyOutput() FleetUpdateStrategyOutput {
	return i.ToFleetUpdateStrategyOutputWithContext(context.Background())
}

func (i *FleetUpdateStrategy) ToFleetUpdateStrategyOutputWithContext(ctx context.Context) FleetUpdateStrategyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetUpdateStrategyOutput)
}

type FleetUpdateStrategyOutput struct{ *pulumi.OutputState }

func (FleetUpdateStrategyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetUpdateStrategy)(nil)).Elem()
}

func (o FleetUpdateStrategyOutput) ToFleetUpdateStrategyOutput() FleetUpdateStrategyOutput {
	return o
}

func (o FleetUpdateStrategyOutput) ToFleetUpdateStrategyOutputWithContext(ctx context.Context) FleetUpdateStrategyOutput {
	return o
}

// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o FleetUpdateStrategyOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetUpdateStrategy) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// The name of the resource
func (o FleetUpdateStrategyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetUpdateStrategy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the UpdateStrategy resource.
func (o FleetUpdateStrategyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetUpdateStrategy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines the update sequence of the clusters.
func (o FleetUpdateStrategyOutput) Strategy() UpdateRunStrategyResponseOutput {
	return o.ApplyT(func(v *FleetUpdateStrategy) UpdateRunStrategyResponseOutput { return v.Strategy }).(UpdateRunStrategyResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FleetUpdateStrategyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FleetUpdateStrategy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FleetUpdateStrategyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetUpdateStrategy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetUpdateStrategyOutput{})
}
