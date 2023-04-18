// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
type CodeContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	CodeContainerProperties CodeContainerResponseOutput `pulumi:"codeContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCodeContainer registers a new resource with the given unique name, arguments, and options.
func NewCodeContainer(ctx *pulumi.Context,
	name string, args *CodeContainerArgs, opts ...pulumi.ResourceOption) (*CodeContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CodeContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'CodeContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.CodeContainerProperties = args.CodeContainerProperties.ToCodeContainerTypeOutput().ApplyT(func(v CodeContainerType) CodeContainerType { return *v.Defaults() }).(CodeContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:CodeContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource CodeContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220601preview:CodeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCodeContainer gets an existing CodeContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCodeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodeContainerState, opts ...pulumi.ResourceOption) (*CodeContainer, error) {
	var resource CodeContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220601preview:CodeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CodeContainer resources.
type codeContainerState struct {
}

type CodeContainerState struct {
}

func (CodeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*codeContainerState)(nil)).Elem()
}

type codeContainerArgs struct {
	// [Required] Additional attributes of the entity.
	CodeContainerProperties CodeContainerType `pulumi:"codeContainerProperties"`
	// Container name. This is case-sensitive.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a CodeContainer resource.
type CodeContainerArgs struct {
	// [Required] Additional attributes of the entity.
	CodeContainerProperties CodeContainerTypeInput
	// Container name. This is case-sensitive.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (CodeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codeContainerArgs)(nil)).Elem()
}

type CodeContainerInput interface {
	pulumi.Input

	ToCodeContainerOutput() CodeContainerOutput
	ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput
}

func (*CodeContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeContainer)(nil)).Elem()
}

func (i *CodeContainer) ToCodeContainerOutput() CodeContainerOutput {
	return i.ToCodeContainerOutputWithContext(context.Background())
}

func (i *CodeContainer) ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeContainerOutput)
}

type CodeContainerOutput struct{ *pulumi.OutputState }

func (CodeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeContainer)(nil)).Elem()
}

func (o CodeContainerOutput) ToCodeContainerOutput() CodeContainerOutput {
	return o
}

func (o CodeContainerOutput) ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o CodeContainerOutput) CodeContainerProperties() CodeContainerResponseOutput {
	return o.ApplyT(func(v *CodeContainer) CodeContainerResponseOutput { return v.CodeContainerProperties }).(CodeContainerResponseOutput)
}

// The name of the resource
func (o CodeContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CodeContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CodeContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CodeContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CodeContainerOutput{})
}
