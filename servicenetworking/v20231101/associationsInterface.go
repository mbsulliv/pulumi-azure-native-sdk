// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Association Subresource of Traffic Controller
type AssociationsInterface struct {
	pulumi.CustomResourceState

	// Association Type
	AssociationType pulumi.StringOutput `pulumi:"associationType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning State of Traffic Controller Association Resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Association Subnet
	Subnet AssociationSubnetResponsePtrOutput `pulumi:"subnet"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssociationsInterface registers a new resource with the given unique name, arguments, and options.
func NewAssociationsInterface(ctx *pulumi.Context,
	name string, args *AssociationsInterfaceArgs, opts ...pulumi.ResourceOption) (*AssociationsInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssociationType == nil {
		return nil, errors.New("invalid value for required argument 'AssociationType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrafficControllerName == nil {
		return nil, errors.New("invalid value for required argument 'TrafficControllerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicenetworking:AssociationsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20221001preview:AssociationsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20230501preview:AssociationsInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AssociationsInterface
	err := ctx.RegisterResource("azure-native:servicenetworking/v20231101:AssociationsInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssociationsInterface gets an existing AssociationsInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssociationsInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssociationsInterfaceState, opts ...pulumi.ResourceOption) (*AssociationsInterface, error) {
	var resource AssociationsInterface
	err := ctx.ReadResource("azure-native:servicenetworking/v20231101:AssociationsInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssociationsInterface resources.
type associationsInterfaceState struct {
}

type AssociationsInterfaceState struct {
}

func (AssociationsInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*associationsInterfaceState)(nil)).Elem()
}

type associationsInterfaceArgs struct {
	// Name of Association
	AssociationName *string `pulumi:"associationName"`
	// Association Type
	AssociationType string `pulumi:"associationType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Association Subnet
	Subnet *AssociationSubnet `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// traffic controller name for path
	TrafficControllerName string `pulumi:"trafficControllerName"`
}

// The set of arguments for constructing a AssociationsInterface resource.
type AssociationsInterfaceArgs struct {
	// Name of Association
	AssociationName pulumi.StringPtrInput
	// Association Type
	AssociationType pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Association Subnet
	Subnet AssociationSubnetPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// traffic controller name for path
	TrafficControllerName pulumi.StringInput
}

func (AssociationsInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*associationsInterfaceArgs)(nil)).Elem()
}

type AssociationsInterfaceInput interface {
	pulumi.Input

	ToAssociationsInterfaceOutput() AssociationsInterfaceOutput
	ToAssociationsInterfaceOutputWithContext(ctx context.Context) AssociationsInterfaceOutput
}

func (*AssociationsInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationsInterface)(nil)).Elem()
}

func (i *AssociationsInterface) ToAssociationsInterfaceOutput() AssociationsInterfaceOutput {
	return i.ToAssociationsInterfaceOutputWithContext(context.Background())
}

func (i *AssociationsInterface) ToAssociationsInterfaceOutputWithContext(ctx context.Context) AssociationsInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationsInterfaceOutput)
}

func (i *AssociationsInterface) ToOutput(ctx context.Context) pulumix.Output[*AssociationsInterface] {
	return pulumix.Output[*AssociationsInterface]{
		OutputState: i.ToAssociationsInterfaceOutputWithContext(ctx).OutputState,
	}
}

type AssociationsInterfaceOutput struct{ *pulumi.OutputState }

func (AssociationsInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationsInterface)(nil)).Elem()
}

func (o AssociationsInterfaceOutput) ToAssociationsInterfaceOutput() AssociationsInterfaceOutput {
	return o
}

func (o AssociationsInterfaceOutput) ToAssociationsInterfaceOutputWithContext(ctx context.Context) AssociationsInterfaceOutput {
	return o
}

func (o AssociationsInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*AssociationsInterface] {
	return pulumix.Output[*AssociationsInterface]{
		OutputState: o.OutputState,
	}
}

// Association Type
func (o AssociationsInterfaceOutput) AssociationType() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.AssociationType }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o AssociationsInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AssociationsInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning State of Traffic Controller Association Resource
func (o AssociationsInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Association Subnet
func (o AssociationsInterfaceOutput) Subnet() AssociationSubnetResponsePtrOutput {
	return o.ApplyT(func(v *AssociationsInterface) AssociationSubnetResponsePtrOutput { return v.Subnet }).(AssociationSubnetResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AssociationsInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AssociationsInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AssociationsInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AssociationsInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssociationsInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssociationsInterfaceOutput{})
}
