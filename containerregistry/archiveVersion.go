// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An object that represents an export pipeline for a container registry.
// Azure REST API version: 2023-06-01-preview.
//
// Other available API versions: 2023-08-01-preview.
type ArchiveVersion struct {
	pulumi.CustomResourceState

	// The detailed error message for the archive version in the case of failure.
	ArchiveVersionErrorMessage pulumi.StringPtrOutput `pulumi:"archiveVersionErrorMessage"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the archive at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewArchiveVersion registers a new resource with the given unique name, arguments, and options.
func NewArchiveVersion(ctx *pulumi.Context,
	name string, args *ArchiveVersionArgs, opts ...pulumi.ResourceOption) (*ArchiveVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArchiveName == nil {
		return nil, errors.New("invalid value for required argument 'ArchiveName'")
	}
	if args.PackageType == nil {
		return nil, errors.New("invalid value for required argument 'PackageType'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:ArchiveVersion"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230801preview:ArchiveVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ArchiveVersion
	err := ctx.RegisterResource("azure-native:containerregistry:ArchiveVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArchiveVersion gets an existing ArchiveVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArchiveVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArchiveVersionState, opts ...pulumi.ResourceOption) (*ArchiveVersion, error) {
	var resource ArchiveVersion
	err := ctx.ReadResource("azure-native:containerregistry:ArchiveVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArchiveVersion resources.
type archiveVersionState struct {
}

type ArchiveVersionState struct {
}

func (ArchiveVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*archiveVersionState)(nil)).Elem()
}

type archiveVersionArgs struct {
	// The name of the archive resource.
	ArchiveName string `pulumi:"archiveName"`
	// The name of the archive version resource.
	ArchiveVersionName *string `pulumi:"archiveVersionName"`
	// The type of the package resource.
	PackageType string `pulumi:"packageType"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ArchiveVersion resource.
type ArchiveVersionArgs struct {
	// The name of the archive resource.
	ArchiveName pulumi.StringInput
	// The name of the archive version resource.
	ArchiveVersionName pulumi.StringPtrInput
	// The type of the package resource.
	PackageType pulumi.StringInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ArchiveVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*archiveVersionArgs)(nil)).Elem()
}

type ArchiveVersionInput interface {
	pulumi.Input

	ToArchiveVersionOutput() ArchiveVersionOutput
	ToArchiveVersionOutputWithContext(ctx context.Context) ArchiveVersionOutput
}

func (*ArchiveVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ArchiveVersion)(nil)).Elem()
}

func (i *ArchiveVersion) ToArchiveVersionOutput() ArchiveVersionOutput {
	return i.ToArchiveVersionOutputWithContext(context.Background())
}

func (i *ArchiveVersion) ToArchiveVersionOutputWithContext(ctx context.Context) ArchiveVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArchiveVersionOutput)
}

func (i *ArchiveVersion) ToOutput(ctx context.Context) pulumix.Output[*ArchiveVersion] {
	return pulumix.Output[*ArchiveVersion]{
		OutputState: i.ToArchiveVersionOutputWithContext(ctx).OutputState,
	}
}

type ArchiveVersionOutput struct{ *pulumi.OutputState }

func (ArchiveVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArchiveVersion)(nil)).Elem()
}

func (o ArchiveVersionOutput) ToArchiveVersionOutput() ArchiveVersionOutput {
	return o
}

func (o ArchiveVersionOutput) ToArchiveVersionOutputWithContext(ctx context.Context) ArchiveVersionOutput {
	return o
}

func (o ArchiveVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*ArchiveVersion] {
	return pulumix.Output[*ArchiveVersion]{
		OutputState: o.OutputState,
	}
}

// The detailed error message for the archive version in the case of failure.
func (o ArchiveVersionOutput) ArchiveVersionErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArchiveVersion) pulumi.StringPtrOutput { return v.ArchiveVersionErrorMessage }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ArchiveVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArchiveVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the archive at the time the operation was called.
func (o ArchiveVersionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ArchiveVersion) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ArchiveVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ArchiveVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ArchiveVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArchiveVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArchiveVersionOutput{})
}
