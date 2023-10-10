// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type ComponentVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionResponseOutput `pulumi:"componentVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewComponentVersion registers a new resource with the given unique name, arguments, and options.
func NewComponentVersion(ctx *pulumi.Context,
	name string, args *ComponentVersionArgs, opts ...pulumi.ResourceOption) (*ComponentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ComponentVersionProperties'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ComponentVersionProperties = args.ComponentVersionProperties.ToComponentVersionTypeOutput().ApplyT(func(v ComponentVersionType) ComponentVersionType { return *v.Defaults() }).(ComponentVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:ComponentVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ComponentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401:ComponentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponentVersion gets an existing ComponentVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentVersionState, opts ...pulumi.ResourceOption) (*ComponentVersion, error) {
	var resource ComponentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401:ComponentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComponentVersion resources.
type componentVersionState struct {
}

type ComponentVersionState struct {
}

func (ComponentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentVersionState)(nil)).Elem()
}

type componentVersionArgs struct {
	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionType `pulumi:"componentVersionProperties"`
	// Container name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ComponentVersion resource.
type ComponentVersionArgs struct {
	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionTypeInput
	// Container name.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (ComponentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentVersionArgs)(nil)).Elem()
}

type ComponentVersionInput interface {
	pulumi.Input

	ToComponentVersionOutput() ComponentVersionOutput
	ToComponentVersionOutputWithContext(ctx context.Context) ComponentVersionOutput
}

func (*ComponentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentVersion)(nil)).Elem()
}

func (i *ComponentVersion) ToComponentVersionOutput() ComponentVersionOutput {
	return i.ToComponentVersionOutputWithContext(context.Background())
}

func (i *ComponentVersion) ToComponentVersionOutputWithContext(ctx context.Context) ComponentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentVersionOutput)
}

func (i *ComponentVersion) ToOutput(ctx context.Context) pulumix.Output[*ComponentVersion] {
	return pulumix.Output[*ComponentVersion]{
		OutputState: i.ToComponentVersionOutputWithContext(ctx).OutputState,
	}
}

type ComponentVersionOutput struct{ *pulumi.OutputState }

func (ComponentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentVersion)(nil)).Elem()
}

func (o ComponentVersionOutput) ToComponentVersionOutput() ComponentVersionOutput {
	return o
}

func (o ComponentVersionOutput) ToComponentVersionOutputWithContext(ctx context.Context) ComponentVersionOutput {
	return o
}

func (o ComponentVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*ComponentVersion] {
	return pulumix.Output[*ComponentVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o ComponentVersionOutput) ComponentVersionProperties() ComponentVersionResponseOutput {
	return o.ApplyT(func(v *ComponentVersion) ComponentVersionResponseOutput { return v.ComponentVersionProperties }).(ComponentVersionResponseOutput)
}

// The name of the resource
func (o ComponentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ComponentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ComponentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ComponentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ComponentVersionOutput{})
}
