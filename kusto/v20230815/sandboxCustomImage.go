// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230815

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Class representing a Kusto sandbox custom image.
type SandboxCustomImage struct {
	pulumi.CustomResourceState

	// The language name, for example Python.
	Language pulumi.StringOutput `pulumi:"language"`
	// The version of the language.
	LanguageVersion pulumi.StringOutput `pulumi:"languageVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The requirements file content.
	RequirementsFileContent pulumi.StringPtrOutput `pulumi:"requirementsFileContent"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSandboxCustomImage registers a new resource with the given unique name, arguments, and options.
func NewSandboxCustomImage(ctx *pulumi.Context,
	name string, args *SandboxCustomImageArgs, opts ...pulumi.ResourceOption) (*SandboxCustomImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Language == nil {
		return nil, errors.New("invalid value for required argument 'Language'")
	}
	if args.LanguageVersion == nil {
		return nil, errors.New("invalid value for required argument 'LanguageVersion'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:SandboxCustomImage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SandboxCustomImage
	err := ctx.RegisterResource("azure-native:kusto/v20230815:SandboxCustomImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSandboxCustomImage gets an existing SandboxCustomImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSandboxCustomImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SandboxCustomImageState, opts ...pulumi.ResourceOption) (*SandboxCustomImage, error) {
	var resource SandboxCustomImage
	err := ctx.ReadResource("azure-native:kusto/v20230815:SandboxCustomImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SandboxCustomImage resources.
type sandboxCustomImageState struct {
}

type SandboxCustomImageState struct {
}

func (SandboxCustomImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*sandboxCustomImageState)(nil)).Elem()
}

type sandboxCustomImageArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The language name, for example Python.
	Language string `pulumi:"language"`
	// The version of the language.
	LanguageVersion string `pulumi:"languageVersion"`
	// The requirements file content.
	RequirementsFileContent *string `pulumi:"requirementsFileContent"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the sandbox custom image.
	SandboxCustomImageName *string `pulumi:"sandboxCustomImageName"`
}

// The set of arguments for constructing a SandboxCustomImage resource.
type SandboxCustomImageArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The language name, for example Python.
	Language pulumi.StringInput
	// The version of the language.
	LanguageVersion pulumi.StringInput
	// The requirements file content.
	RequirementsFileContent pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the sandbox custom image.
	SandboxCustomImageName pulumi.StringPtrInput
}

func (SandboxCustomImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sandboxCustomImageArgs)(nil)).Elem()
}

type SandboxCustomImageInput interface {
	pulumi.Input

	ToSandboxCustomImageOutput() SandboxCustomImageOutput
	ToSandboxCustomImageOutputWithContext(ctx context.Context) SandboxCustomImageOutput
}

func (*SandboxCustomImage) ElementType() reflect.Type {
	return reflect.TypeOf((**SandboxCustomImage)(nil)).Elem()
}

func (i *SandboxCustomImage) ToSandboxCustomImageOutput() SandboxCustomImageOutput {
	return i.ToSandboxCustomImageOutputWithContext(context.Background())
}

func (i *SandboxCustomImage) ToSandboxCustomImageOutputWithContext(ctx context.Context) SandboxCustomImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandboxCustomImageOutput)
}

func (i *SandboxCustomImage) ToOutput(ctx context.Context) pulumix.Output[*SandboxCustomImage] {
	return pulumix.Output[*SandboxCustomImage]{
		OutputState: i.ToSandboxCustomImageOutputWithContext(ctx).OutputState,
	}
}

type SandboxCustomImageOutput struct{ *pulumi.OutputState }

func (SandboxCustomImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SandboxCustomImage)(nil)).Elem()
}

func (o SandboxCustomImageOutput) ToSandboxCustomImageOutput() SandboxCustomImageOutput {
	return o
}

func (o SandboxCustomImageOutput) ToSandboxCustomImageOutputWithContext(ctx context.Context) SandboxCustomImageOutput {
	return o
}

func (o SandboxCustomImageOutput) ToOutput(ctx context.Context) pulumix.Output[*SandboxCustomImage] {
	return pulumix.Output[*SandboxCustomImage]{
		OutputState: o.OutputState,
	}
}

// The language name, for example Python.
func (o SandboxCustomImageOutput) Language() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxCustomImage) pulumi.StringOutput { return v.Language }).(pulumi.StringOutput)
}

// The version of the language.
func (o SandboxCustomImageOutput) LanguageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxCustomImage) pulumi.StringOutput { return v.LanguageVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o SandboxCustomImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxCustomImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o SandboxCustomImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxCustomImage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The requirements file content.
func (o SandboxCustomImageOutput) RequirementsFileContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SandboxCustomImage) pulumi.StringPtrOutput { return v.RequirementsFileContent }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SandboxCustomImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxCustomImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SandboxCustomImageOutput{})
}
