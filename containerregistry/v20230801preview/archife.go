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

// An object that represents a archive for a container registry.
type Archife struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The package source of the archive.
	PackageSource ArchivePackageSourcePropertiesResponsePtrOutput `pulumi:"packageSource"`
	// The provisioning state of the archive at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The published version of the archive.
	PublishedVersion         pulumi.StringPtrOutput `pulumi:"publishedVersion"`
	RepositoryEndpoint       pulumi.StringOutput    `pulumi:"repositoryEndpoint"`
	RepositoryEndpointPrefix pulumi.StringPtrOutput `pulumi:"repositoryEndpointPrefix"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewArchife registers a new resource with the given unique name, arguments, and options.
func NewArchife(ctx *pulumi.Context,
	name string, args *ArchifeArgs, opts ...pulumi.ResourceOption) (*Archife, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:containerregistry:Archife"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:Archife"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Archife
	err := ctx.RegisterResource("azure-native:containerregistry/v20230801preview:Archife", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArchife gets an existing Archife resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArchife(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArchifeState, opts ...pulumi.ResourceOption) (*Archife, error) {
	var resource Archife
	err := ctx.ReadResource("azure-native:containerregistry/v20230801preview:Archife", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Archife resources.
type archifeState struct {
}

type ArchifeState struct {
}

func (ArchifeState) ElementType() reflect.Type {
	return reflect.TypeOf((*archifeState)(nil)).Elem()
}

type archifeArgs struct {
	// The name of the archive resource.
	ArchiveName *string `pulumi:"archiveName"`
	// The package source of the archive.
	PackageSource *ArchivePackageSourceProperties `pulumi:"packageSource"`
	// The type of the package resource.
	PackageType string `pulumi:"packageType"`
	// The published version of the archive.
	PublishedVersion *string `pulumi:"publishedVersion"`
	// The name of the container registry.
	RegistryName             string  `pulumi:"registryName"`
	RepositoryEndpointPrefix *string `pulumi:"repositoryEndpointPrefix"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Archife resource.
type ArchifeArgs struct {
	// The name of the archive resource.
	ArchiveName pulumi.StringPtrInput
	// The package source of the archive.
	PackageSource ArchivePackageSourcePropertiesPtrInput
	// The type of the package resource.
	PackageType pulumi.StringInput
	// The published version of the archive.
	PublishedVersion pulumi.StringPtrInput
	// The name of the container registry.
	RegistryName             pulumi.StringInput
	RepositoryEndpointPrefix pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ArchifeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*archifeArgs)(nil)).Elem()
}

type ArchifeInput interface {
	pulumi.Input

	ToArchifeOutput() ArchifeOutput
	ToArchifeOutputWithContext(ctx context.Context) ArchifeOutput
}

func (*Archife) ElementType() reflect.Type {
	return reflect.TypeOf((**Archife)(nil)).Elem()
}

func (i *Archife) ToArchifeOutput() ArchifeOutput {
	return i.ToArchifeOutputWithContext(context.Background())
}

func (i *Archife) ToArchifeOutputWithContext(ctx context.Context) ArchifeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArchifeOutput)
}

func (i *Archife) ToOutput(ctx context.Context) pulumix.Output[*Archife] {
	return pulumix.Output[*Archife]{
		OutputState: i.ToArchifeOutputWithContext(ctx).OutputState,
	}
}

type ArchifeOutput struct{ *pulumi.OutputState }

func (ArchifeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Archife)(nil)).Elem()
}

func (o ArchifeOutput) ToArchifeOutput() ArchifeOutput {
	return o
}

func (o ArchifeOutput) ToArchifeOutputWithContext(ctx context.Context) ArchifeOutput {
	return o
}

func (o ArchifeOutput) ToOutput(ctx context.Context) pulumix.Output[*Archife] {
	return pulumix.Output[*Archife]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o ArchifeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Archife) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The package source of the archive.
func (o ArchifeOutput) PackageSource() ArchivePackageSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Archife) ArchivePackageSourcePropertiesResponsePtrOutput { return v.PackageSource }).(ArchivePackageSourcePropertiesResponsePtrOutput)
}

// The provisioning state of the archive at the time the operation was called.
func (o ArchifeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Archife) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The published version of the archive.
func (o ArchifeOutput) PublishedVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Archife) pulumi.StringPtrOutput { return v.PublishedVersion }).(pulumi.StringPtrOutput)
}

func (o ArchifeOutput) RepositoryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Archife) pulumi.StringOutput { return v.RepositoryEndpoint }).(pulumi.StringOutput)
}

func (o ArchifeOutput) RepositoryEndpointPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Archife) pulumi.StringPtrOutput { return v.RepositoryEndpointPrefix }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ArchifeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Archife) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ArchifeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Archife) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArchifeOutput{})
}
