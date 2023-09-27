// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220521preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines the HybridIdentityMetadata.
type HybridIdentityMetadata struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity IdentityResponseOutput `pulumi:"identity"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the Public Key.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// Gets or sets the Vm Id.
	ResourceUid pulumi.StringPtrOutput `pulumi:"resourceUid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHybridIdentityMetadata registers a new resource with the given unique name, arguments, and options.
func NewHybridIdentityMetadata(ctx *pulumi.Context,
	name string, args *HybridIdentityMetadataArgs, opts ...pulumi.ResourceOption) (*HybridIdentityMetadata, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:HybridIdentityMetadata"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20230401preview:HybridIdentityMetadata"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HybridIdentityMetadata
	err := ctx.RegisterResource("azure-native:scvmm/v20220521preview:HybridIdentityMetadata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridIdentityMetadata gets an existing HybridIdentityMetadata resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridIdentityMetadata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridIdentityMetadataState, opts ...pulumi.ResourceOption) (*HybridIdentityMetadata, error) {
	var resource HybridIdentityMetadata
	err := ctx.ReadResource("azure-native:scvmm/v20220521preview:HybridIdentityMetadata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridIdentityMetadata resources.
type hybridIdentityMetadataState struct {
}

type HybridIdentityMetadataState struct {
}

func (HybridIdentityMetadataState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadataState)(nil)).Elem()
}

type hybridIdentityMetadataArgs struct {
	// Name of the hybridIdentityMetadata.
	MetadataName *string `pulumi:"metadataName"`
	// Gets or sets the Public Key.
	PublicKey *string `pulumi:"publicKey"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the Vm Id.
	ResourceUid *string `pulumi:"resourceUid"`
	// Name of the vm.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// The set of arguments for constructing a HybridIdentityMetadata resource.
type HybridIdentityMetadataArgs struct {
	// Name of the hybridIdentityMetadata.
	MetadataName pulumi.StringPtrInput
	// Gets or sets the Public Key.
	PublicKey pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the Vm Id.
	ResourceUid pulumi.StringPtrInput
	// Name of the vm.
	VirtualMachineName pulumi.StringInput
}

func (HybridIdentityMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadataArgs)(nil)).Elem()
}

type HybridIdentityMetadataInput interface {
	pulumi.Input

	ToHybridIdentityMetadataOutput() HybridIdentityMetadataOutput
	ToHybridIdentityMetadataOutputWithContext(ctx context.Context) HybridIdentityMetadataOutput
}

func (*HybridIdentityMetadata) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadata)(nil)).Elem()
}

func (i *HybridIdentityMetadata) ToHybridIdentityMetadataOutput() HybridIdentityMetadataOutput {
	return i.ToHybridIdentityMetadataOutputWithContext(context.Background())
}

func (i *HybridIdentityMetadata) ToHybridIdentityMetadataOutputWithContext(ctx context.Context) HybridIdentityMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridIdentityMetadataOutput)
}

func (i *HybridIdentityMetadata) ToOutput(ctx context.Context) pulumix.Output[*HybridIdentityMetadata] {
	return pulumix.Output[*HybridIdentityMetadata]{
		OutputState: i.ToHybridIdentityMetadataOutputWithContext(ctx).OutputState,
	}
}

type HybridIdentityMetadataOutput struct{ *pulumi.OutputState }

func (HybridIdentityMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadata)(nil)).Elem()
}

func (o HybridIdentityMetadataOutput) ToHybridIdentityMetadataOutput() HybridIdentityMetadataOutput {
	return o
}

func (o HybridIdentityMetadataOutput) ToHybridIdentityMetadataOutputWithContext(ctx context.Context) HybridIdentityMetadataOutput {
	return o
}

func (o HybridIdentityMetadataOutput) ToOutput(ctx context.Context) pulumix.Output[*HybridIdentityMetadata] {
	return pulumix.Output[*HybridIdentityMetadata]{
		OutputState: o.OutputState,
	}
}

// The identity of the resource.
func (o HybridIdentityMetadataOutput) Identity() IdentityResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadata) IdentityResponseOutput { return v.Identity }).(IdentityResponseOutput)
}

// The name of the resource
func (o HybridIdentityMetadataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadata) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o HybridIdentityMetadataOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadata) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the Public Key.
func (o HybridIdentityMetadataOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadata) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Gets or sets the Vm Id.
func (o HybridIdentityMetadataOutput) ResourceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadata) pulumi.StringPtrOutput { return v.ResourceUid }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HybridIdentityMetadataOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadata) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HybridIdentityMetadataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadata) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridIdentityMetadataOutput{})
}
