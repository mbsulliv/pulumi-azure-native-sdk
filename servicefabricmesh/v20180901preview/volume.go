// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This type describes a volume resource.
type Volume struct {
	pulumi.CustomResourceState

	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters VolumeProviderParametersAzureFileResponsePtrOutput `pulumi:"azureFileParameters"`
	// User readable description of the volume.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provider of the volume.
	Provider pulumi.StringOutput `pulumi:"provider"`
	// State of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Status of the volume.
	Status pulumi.StringOutput `pulumi:"status"`
	// Gives additional information about the current status of the volume.
	StatusDetails pulumi.StringOutput `pulumi:"statusDetails"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Provider == nil {
		return nil, errors.New("invalid value for required argument 'Provider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh:Volume"),
		},
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180701preview:Volume"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("azure-native:servicefabricmesh/v20180901preview:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:servicefabricmesh/v20180901preview:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
}

type VolumeState struct {
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters *VolumeProviderParametersAzureFile `pulumi:"azureFileParameters"`
	// User readable description of the volume.
	Description *string `pulumi:"description"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Provider of the volume.
	Provider string `pulumi:"provider"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The identity of the volume.
	VolumeResourceName *string `pulumi:"volumeResourceName"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters VolumeProviderParametersAzureFilePtrInput
	// User readable description of the volume.
	Description pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Provider of the volume.
	Provider pulumi.StringInput
	// Azure resource group name
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The identity of the volume.
	VolumeResourceName pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

// This type describes a volume provided by an Azure Files file share.
func (o VolumeOutput) AzureFileParameters() VolumeProviderParametersAzureFileResponsePtrOutput {
	return o.ApplyT(func(v *Volume) VolumeProviderParametersAzureFileResponsePtrOutput { return v.AzureFileParameters }).(VolumeProviderParametersAzureFileResponsePtrOutput)
}

// User readable description of the volume.
func (o VolumeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o VolumeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provider of the volume.
func (o VolumeOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Provider }).(pulumi.StringOutput)
}

// State of the resource.
func (o VolumeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Status of the volume.
func (o VolumeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Gives additional information about the current status of the volume.
func (o VolumeOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.StatusDetails }).(pulumi.StringOutput)
}

// Resource tags.
func (o VolumeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
