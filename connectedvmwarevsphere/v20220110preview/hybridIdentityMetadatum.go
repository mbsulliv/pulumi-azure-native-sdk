// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the HybridIdentityMetadata.
type HybridIdentityMetadatum struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity IdentityResponseOutput `pulumi:"identity"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the Public Key.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets the Vm Id.
	VmId pulumi.StringPtrOutput `pulumi:"vmId"`
}

// NewHybridIdentityMetadatum registers a new resource with the given unique name, arguments, and options.
func NewHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, args *HybridIdentityMetadatumArgs, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
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
			Type: pulumi.String("azure-native:connectedvmwarevsphere:HybridIdentityMetadatum"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:HybridIdentityMetadatum"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:HybridIdentityMetadatum"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridIdentityMetadatum
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20220110preview:HybridIdentityMetadatum", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridIdentityMetadatum gets an existing HybridIdentityMetadatum resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridIdentityMetadatumState, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
	var resource HybridIdentityMetadatum
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20220110preview:HybridIdentityMetadatum", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridIdentityMetadatum resources.
type hybridIdentityMetadatumState struct {
}

type HybridIdentityMetadatumState struct {
}

func (HybridIdentityMetadatumState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadatumState)(nil)).Elem()
}

type hybridIdentityMetadatumArgs struct {
	// Name of the hybridIdentityMetadata.
	MetadataName *string `pulumi:"metadataName"`
	// Gets or sets the Public Key.
	PublicKey *string `pulumi:"publicKey"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the vm.
	VirtualMachineName string `pulumi:"virtualMachineName"`
	// Gets or sets the Vm Id.
	VmId *string `pulumi:"vmId"`
}

// The set of arguments for constructing a HybridIdentityMetadatum resource.
type HybridIdentityMetadatumArgs struct {
	// Name of the hybridIdentityMetadata.
	MetadataName pulumi.StringPtrInput
	// Gets or sets the Public Key.
	PublicKey pulumi.StringPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Name of the vm.
	VirtualMachineName pulumi.StringInput
	// Gets or sets the Vm Id.
	VmId pulumi.StringPtrInput
}

func (HybridIdentityMetadatumArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadatumArgs)(nil)).Elem()
}

type HybridIdentityMetadatumInput interface {
	pulumi.Input

	ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput
	ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput
}

func (*HybridIdentityMetadatum) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadatum)(nil)).Elem()
}

func (i *HybridIdentityMetadatum) ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput {
	return i.ToHybridIdentityMetadatumOutputWithContext(context.Background())
}

func (i *HybridIdentityMetadatum) ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridIdentityMetadatumOutput)
}

type HybridIdentityMetadatumOutput struct{ *pulumi.OutputState }

func (HybridIdentityMetadatumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadatum)(nil)).Elem()
}

func (o HybridIdentityMetadatumOutput) ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput {
	return o
}

func (o HybridIdentityMetadatumOutput) ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput {
	return o
}

// The identity of the resource.
func (o HybridIdentityMetadatumOutput) Identity() IdentityResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) IdentityResponseOutput { return v.Identity }).(IdentityResponseOutput)
}

// The name of the resource
func (o HybridIdentityMetadatumOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o HybridIdentityMetadatumOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the Public Key.
func (o HybridIdentityMetadatumOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// The system data.
func (o HybridIdentityMetadatumOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HybridIdentityMetadatumOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the Vm Id.
func (o HybridIdentityMetadatumOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringPtrOutput { return v.VmId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridIdentityMetadatumOutput{})
}
