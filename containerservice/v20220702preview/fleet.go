// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220702preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Fleet resource which contains multiple Kubernetes clusters as its members.
type Fleet struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The FleetHubProfile configures the Fleet's hub.
	HubProfile FleetHubProfileResponsePtrOutput `pulumi:"hubProfile"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the last accepted operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOption) (*Fleet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230315preview:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230615preview:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230815preview:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20231015:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20240202preview:Fleet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Fleet
	err := ctx.RegisterResource("azure-native:containerservice/v20220702preview:Fleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleet gets an existing Fleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetState, opts ...pulumi.ResourceOption) (*Fleet, error) {
	var resource Fleet
	err := ctx.ReadResource("azure-native:containerservice/v20220702preview:Fleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fleet resources.
type fleetState struct {
}

type FleetState struct {
}

func (FleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetState)(nil)).Elem()
}

type fleetArgs struct {
	// The name of the Fleet resource.
	FleetName *string `pulumi:"fleetName"`
	// The FleetHubProfile configures the Fleet's hub.
	HubProfile *FleetHubProfile `pulumi:"hubProfile"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// The name of the Fleet resource.
	FleetName pulumi.StringPtrInput
	// The FleetHubProfile configures the Fleet's hub.
	HubProfile FleetHubProfilePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (FleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetArgs)(nil)).Elem()
}

type FleetInput interface {
	pulumi.Input

	ToFleetOutput() FleetOutput
	ToFleetOutputWithContext(ctx context.Context) FleetOutput
}

func (*Fleet) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil)).Elem()
}

func (i *Fleet) ToFleetOutput() FleetOutput {
	return i.ToFleetOutputWithContext(context.Background())
}

func (i *Fleet) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetOutput)
}

type FleetOutput struct{ *pulumi.OutputState }

func (FleetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil)).Elem()
}

func (o FleetOutput) ToFleetOutput() FleetOutput {
	return o
}

func (o FleetOutput) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return o
}

// Resource Etag.
func (o FleetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The FleetHubProfile configures the Fleet's hub.
func (o FleetOutput) HubProfile() FleetHubProfileResponsePtrOutput {
	return o.ApplyT(func(v *Fleet) FleetHubProfileResponsePtrOutput { return v.HubProfile }).(FleetHubProfileResponsePtrOutput)
}

// The geo-location where the resource lives
func (o FleetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o FleetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the last accepted operation.
func (o FleetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FleetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Fleet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o FleetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FleetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetOutput{})
}
