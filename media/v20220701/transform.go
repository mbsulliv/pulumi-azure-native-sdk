// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.
type Transform struct {
	pulumi.CustomResourceState

	// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created pulumi.StringOutput `pulumi:"created"`
	// An optional verbose description of the Transform.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs TransformOutputResponseArrayOutput `pulumi:"outputs"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTransform registers a new resource with the given unique name, arguments, and options.
func NewTransform(ctx *pulumi.Context,
	name string, args *TransformArgs, opts ...pulumi.ResourceOption) (*Transform, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Outputs == nil {
		return nil, errors.New("invalid value for required argument 'Outputs'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220501preview:Transform"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Transform
	err := ctx.RegisterResource("azure-native:media/v20220701:Transform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransform gets an existing Transform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransformState, opts ...pulumi.ResourceOption) (*Transform, error) {
	var resource Transform
	err := ctx.ReadResource("azure-native:media/v20220701:Transform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Transform resources.
type transformState struct {
}

type TransformState struct {
}

func (TransformState) ElementType() reflect.Type {
	return reflect.TypeOf((*transformState)(nil)).Elem()
}

type transformArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// An optional verbose description of the Transform.
	Description *string `pulumi:"description"`
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs []TransformOutputType `pulumi:"outputs"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Transform name.
	TransformName *string `pulumi:"transformName"`
}

// The set of arguments for constructing a Transform resource.
type TransformArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// An optional verbose description of the Transform.
	Description pulumi.StringPtrInput
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs TransformOutputTypeArrayInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The Transform name.
	TransformName pulumi.StringPtrInput
}

func (TransformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transformArgs)(nil)).Elem()
}

type TransformInput interface {
	pulumi.Input

	ToTransformOutput() TransformOutput
	ToTransformOutputWithContext(ctx context.Context) TransformOutput
}

func (*Transform) ElementType() reflect.Type {
	return reflect.TypeOf((**Transform)(nil)).Elem()
}

func (i *Transform) ToTransformOutput() TransformOutput {
	return i.ToTransformOutputWithContext(context.Background())
}

func (i *Transform) ToTransformOutputWithContext(ctx context.Context) TransformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutput)
}

func (i *Transform) ToOutput(ctx context.Context) pulumix.Output[*Transform] {
	return pulumix.Output[*Transform]{
		OutputState: i.ToTransformOutputWithContext(ctx).OutputState,
	}
}

type TransformOutput struct{ *pulumi.OutputState }

func (TransformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transform)(nil)).Elem()
}

func (o TransformOutput) ToTransformOutput() TransformOutput {
	return o
}

func (o TransformOutput) ToTransformOutputWithContext(ctx context.Context) TransformOutput {
	return o
}

func (o TransformOutput) ToOutput(ctx context.Context) pulumix.Output[*Transform] {
	return pulumix.Output[*Transform]{
		OutputState: o.OutputState,
	}
}

// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
func (o TransformOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Transform) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// An optional verbose description of the Transform.
func (o TransformOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transform) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
func (o TransformOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Transform) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The name of the resource
func (o TransformOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Transform) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of one or more TransformOutputs that the Transform should generate.
func (o TransformOutput) Outputs() TransformOutputResponseArrayOutput {
	return o.ApplyT(func(v *Transform) TransformOutputResponseArrayOutput { return v.Outputs }).(TransformOutputResponseArrayOutput)
}

// The system metadata relating to this resource.
func (o TransformOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Transform) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TransformOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Transform) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TransformOutput{})
}
