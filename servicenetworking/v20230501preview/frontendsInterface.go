// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Frontend Subresource of Traffic Controller.
type FrontendsInterface struct {
	pulumi.CustomResourceState

	// The Fully Qualified Domain Name of the DNS record associated to a Traffic Controller frontend.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning State of Traffic Controller Frontend Resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFrontendsInterface registers a new resource with the given unique name, arguments, and options.
func NewFrontendsInterface(ctx *pulumi.Context,
	name string, args *FrontendsInterfaceArgs, opts ...pulumi.ResourceOption) (*FrontendsInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrafficControllerName == nil {
		return nil, errors.New("invalid value for required argument 'TrafficControllerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicenetworking:FrontendsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20221001preview:FrontendsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20231101:FrontendsInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FrontendsInterface
	err := ctx.RegisterResource("azure-native:servicenetworking/v20230501preview:FrontendsInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFrontendsInterface gets an existing FrontendsInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFrontendsInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FrontendsInterfaceState, opts ...pulumi.ResourceOption) (*FrontendsInterface, error) {
	var resource FrontendsInterface
	err := ctx.ReadResource("azure-native:servicenetworking/v20230501preview:FrontendsInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FrontendsInterface resources.
type frontendsInterfaceState struct {
}

type FrontendsInterfaceState struct {
}

func (FrontendsInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendsInterfaceState)(nil)).Elem()
}

type frontendsInterfaceArgs struct {
	// Frontends
	FrontendName *string `pulumi:"frontendName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// traffic controller name for path
	TrafficControllerName string `pulumi:"trafficControllerName"`
}

// The set of arguments for constructing a FrontendsInterface resource.
type FrontendsInterfaceArgs struct {
	// Frontends
	FrontendName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// traffic controller name for path
	TrafficControllerName pulumi.StringInput
}

func (FrontendsInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendsInterfaceArgs)(nil)).Elem()
}

type FrontendsInterfaceInput interface {
	pulumi.Input

	ToFrontendsInterfaceOutput() FrontendsInterfaceOutput
	ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput
}

func (*FrontendsInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendsInterface)(nil)).Elem()
}

func (i *FrontendsInterface) ToFrontendsInterfaceOutput() FrontendsInterfaceOutput {
	return i.ToFrontendsInterfaceOutputWithContext(context.Background())
}

func (i *FrontendsInterface) ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendsInterfaceOutput)
}

func (i *FrontendsInterface) ToOutput(ctx context.Context) pulumix.Output[*FrontendsInterface] {
	return pulumix.Output[*FrontendsInterface]{
		OutputState: i.ToFrontendsInterfaceOutputWithContext(ctx).OutputState,
	}
}

type FrontendsInterfaceOutput struct{ *pulumi.OutputState }

func (FrontendsInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendsInterface)(nil)).Elem()
}

func (o FrontendsInterfaceOutput) ToFrontendsInterfaceOutput() FrontendsInterfaceOutput {
	return o
}

func (o FrontendsInterfaceOutput) ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput {
	return o
}

func (o FrontendsInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*FrontendsInterface] {
	return pulumix.Output[*FrontendsInterface]{
		OutputState: o.OutputState,
	}
}

// The Fully Qualified Domain Name of the DNS record associated to a Traffic Controller frontend.
func (o FrontendsInterfaceOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o FrontendsInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o FrontendsInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning State of Traffic Controller Frontend Resource
func (o FrontendsInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FrontendsInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FrontendsInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o FrontendsInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FrontendsInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FrontendsInterfaceOutput{})
}
