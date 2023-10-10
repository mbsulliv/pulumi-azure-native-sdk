// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type CodeVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	CodeVersionProperties CodeVersionResponseOutput `pulumi:"codeVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCodeVersion registers a new resource with the given unique name, arguments, and options.
func NewCodeVersion(ctx *pulumi.Context,
	name string, args *CodeVersionArgs, opts ...pulumi.ResourceOption) (*CodeVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CodeVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'CodeVersionProperties'")
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
	args.CodeVersionProperties = args.CodeVersionProperties.ToCodeVersionTypeOutput().ApplyT(func(v CodeVersionType) CodeVersionType { return *v.Defaults() }).(CodeVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:CodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:CodeVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CodeVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230801preview:CodeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCodeVersion gets an existing CodeVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCodeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodeVersionState, opts ...pulumi.ResourceOption) (*CodeVersion, error) {
	var resource CodeVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230801preview:CodeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CodeVersion resources.
type codeVersionState struct {
}

type CodeVersionState struct {
}

func (CodeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*codeVersionState)(nil)).Elem()
}

type codeVersionArgs struct {
	// [Required] Additional attributes of the entity.
	CodeVersionProperties CodeVersionType `pulumi:"codeVersionProperties"`
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a CodeVersion resource.
type CodeVersionArgs struct {
	// [Required] Additional attributes of the entity.
	CodeVersionProperties CodeVersionTypeInput
	// Container name. This is case-sensitive.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier. This is case-sensitive.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (CodeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codeVersionArgs)(nil)).Elem()
}

type CodeVersionInput interface {
	pulumi.Input

	ToCodeVersionOutput() CodeVersionOutput
	ToCodeVersionOutputWithContext(ctx context.Context) CodeVersionOutput
}

func (*CodeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeVersion)(nil)).Elem()
}

func (i *CodeVersion) ToCodeVersionOutput() CodeVersionOutput {
	return i.ToCodeVersionOutputWithContext(context.Background())
}

func (i *CodeVersion) ToCodeVersionOutputWithContext(ctx context.Context) CodeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeVersionOutput)
}

func (i *CodeVersion) ToOutput(ctx context.Context) pulumix.Output[*CodeVersion] {
	return pulumix.Output[*CodeVersion]{
		OutputState: i.ToCodeVersionOutputWithContext(ctx).OutputState,
	}
}

type CodeVersionOutput struct{ *pulumi.OutputState }

func (CodeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeVersion)(nil)).Elem()
}

func (o CodeVersionOutput) ToCodeVersionOutput() CodeVersionOutput {
	return o
}

func (o CodeVersionOutput) ToCodeVersionOutputWithContext(ctx context.Context) CodeVersionOutput {
	return o
}

func (o CodeVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*CodeVersion] {
	return pulumix.Output[*CodeVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o CodeVersionOutput) CodeVersionProperties() CodeVersionResponseOutput {
	return o.ApplyT(func(v *CodeVersion) CodeVersionResponseOutput { return v.CodeVersionProperties }).(CodeVersionResponseOutput)
}

// The name of the resource
func (o CodeVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CodeVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CodeVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CodeVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CodeVersionOutput{})
}
