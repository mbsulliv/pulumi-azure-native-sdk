// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type DataVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	DataVersionBaseDetails pulumi.AnyOutput `pulumi:"dataVersionBaseDetails"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataVersion registers a new resource with the given unique name, arguments, and options.
func NewDataVersion(ctx *pulumi.Context,
	name string, args *DataVersionArgs, opts ...pulumi.ResourceOption) (*DataVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataVersionBaseDetails == nil {
		return nil, errors.New("invalid value for required argument 'DataVersionBaseDetails'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:DataVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:DataVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataVersion gets an existing DataVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataVersionState, opts ...pulumi.ResourceOption) (*DataVersion, error) {
	var resource DataVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:DataVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataVersion resources.
type dataVersionState struct {
}

type DataVersionState struct {
}

func (DataVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataVersionState)(nil)).Elem()
}

type dataVersionArgs struct {
	// [Required] Additional attributes of the entity.
	DataVersionBaseDetails interface{} `pulumi:"dataVersionBaseDetails"`
	// Container name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataVersion resource.
type DataVersionArgs struct {
	// [Required] Additional attributes of the entity.
	DataVersionBaseDetails pulumi.Input
	// Container name.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (DataVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataVersionArgs)(nil)).Elem()
}

type DataVersionInput interface {
	pulumi.Input

	ToDataVersionOutput() DataVersionOutput
	ToDataVersionOutputWithContext(ctx context.Context) DataVersionOutput
}

func (*DataVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**DataVersion)(nil)).Elem()
}

func (i *DataVersion) ToDataVersionOutput() DataVersionOutput {
	return i.ToDataVersionOutputWithContext(context.Background())
}

func (i *DataVersion) ToDataVersionOutputWithContext(ctx context.Context) DataVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataVersionOutput)
}

func (i *DataVersion) ToOutput(ctx context.Context) pulumix.Output[*DataVersion] {
	return pulumix.Output[*DataVersion]{
		OutputState: i.ToDataVersionOutputWithContext(ctx).OutputState,
	}
}

type DataVersionOutput struct{ *pulumi.OutputState }

func (DataVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataVersion)(nil)).Elem()
}

func (o DataVersionOutput) ToDataVersionOutput() DataVersionOutput {
	return o
}

func (o DataVersionOutput) ToDataVersionOutputWithContext(ctx context.Context) DataVersionOutput {
	return o
}

func (o DataVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*DataVersion] {
	return pulumix.Output[*DataVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o DataVersionOutput) DataVersionBaseDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataVersion) pulumi.AnyOutput { return v.DataVersionBaseDetails }).(pulumi.AnyOutput)
}

// The name of the resource
func (o DataVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DataVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataVersionOutput{})
}
