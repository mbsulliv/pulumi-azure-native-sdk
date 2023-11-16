// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Concrete tracked resource types can be created by aliasing this type using a specific property type.
type TrafficControllerInterface struct {
	pulumi.CustomResourceState

	// Associations References List
	Associations ResourceIDResponseArrayOutput `pulumi:"associations"`
	// Configuration Endpoints.
	ConfigurationEndpoints pulumi.StringArrayOutput `pulumi:"configurationEndpoints"`
	// Frontends References List
	Frontends ResourceIDResponseArrayOutput `pulumi:"frontends"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrafficControllerInterface registers a new resource with the given unique name, arguments, and options.
func NewTrafficControllerInterface(ctx *pulumi.Context,
	name string, args *TrafficControllerInterfaceArgs, opts ...pulumi.ResourceOption) (*TrafficControllerInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicenetworking:TrafficControllerInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20230501preview:TrafficControllerInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20231101:TrafficControllerInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TrafficControllerInterface
	err := ctx.RegisterResource("azure-native:servicenetworking/v20221001preview:TrafficControllerInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficControllerInterface gets an existing TrafficControllerInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficControllerInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficControllerInterfaceState, opts ...pulumi.ResourceOption) (*TrafficControllerInterface, error) {
	var resource TrafficControllerInterface
	err := ctx.ReadResource("azure-native:servicenetworking/v20221001preview:TrafficControllerInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficControllerInterface resources.
type trafficControllerInterfaceState struct {
}

type TrafficControllerInterfaceState struct {
}

func (TrafficControllerInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficControllerInterfaceState)(nil)).Elem()
}

type trafficControllerInterfaceArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// traffic controller name for path
	TrafficControllerName *string `pulumi:"trafficControllerName"`
}

// The set of arguments for constructing a TrafficControllerInterface resource.
type TrafficControllerInterfaceArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// traffic controller name for path
	TrafficControllerName pulumi.StringPtrInput
}

func (TrafficControllerInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficControllerInterfaceArgs)(nil)).Elem()
}

type TrafficControllerInterfaceInput interface {
	pulumi.Input

	ToTrafficControllerInterfaceOutput() TrafficControllerInterfaceOutput
	ToTrafficControllerInterfaceOutputWithContext(ctx context.Context) TrafficControllerInterfaceOutput
}

func (*TrafficControllerInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficControllerInterface)(nil)).Elem()
}

func (i *TrafficControllerInterface) ToTrafficControllerInterfaceOutput() TrafficControllerInterfaceOutput {
	return i.ToTrafficControllerInterfaceOutputWithContext(context.Background())
}

func (i *TrafficControllerInterface) ToTrafficControllerInterfaceOutputWithContext(ctx context.Context) TrafficControllerInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficControllerInterfaceOutput)
}

type TrafficControllerInterfaceOutput struct{ *pulumi.OutputState }

func (TrafficControllerInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficControllerInterface)(nil)).Elem()
}

func (o TrafficControllerInterfaceOutput) ToTrafficControllerInterfaceOutput() TrafficControllerInterfaceOutput {
	return o
}

func (o TrafficControllerInterfaceOutput) ToTrafficControllerInterfaceOutputWithContext(ctx context.Context) TrafficControllerInterfaceOutput {
	return o
}

// Associations References List
func (o TrafficControllerInterfaceOutput) Associations() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) ResourceIDResponseArrayOutput { return v.Associations }).(ResourceIDResponseArrayOutput)
}

// Configuration Endpoints.
func (o TrafficControllerInterfaceOutput) ConfigurationEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringArrayOutput { return v.ConfigurationEndpoints }).(pulumi.StringArrayOutput)
}

// Frontends References List
func (o TrafficControllerInterfaceOutput) Frontends() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) ResourceIDResponseArrayOutput { return v.Frontends }).(ResourceIDResponseArrayOutput)
}

// The geo-location where the resource lives
func (o TrafficControllerInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o TrafficControllerInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o TrafficControllerInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o TrafficControllerInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o TrafficControllerInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TrafficControllerInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrafficControllerInterfaceOutput{})
}
