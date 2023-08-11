// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230402

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// disk encryption set resource.
type DiskEncryptionSet struct {
	pulumi.CustomResourceState

	// The key vault key which is currently used by this disk encryption set.
	ActiveKey KeyForDiskEncryptionSetResponsePtrOutput `pulumi:"activeKey"`
	// The error that was encountered during auto-key rotation. If an error is present, then auto-key rotation will not be attempted until the error on this disk encryption set is fixed.
	AutoKeyRotationError ApiErrorResponseOutput `pulumi:"autoKeyRotationError"`
	// The type of key used to encrypt the data of the disk.
	EncryptionType pulumi.StringPtrOutput `pulumi:"encryptionType"`
	// Multi-tenant application client id to access key vault in a different tenant. Setting the value to 'None' will clear the property.
	FederatedClientId pulumi.StringPtrOutput `pulumi:"federatedClientId"`
	// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
	Identity EncryptionSetIdentityResponsePtrOutput `pulumi:"identity"`
	// The time when the active key of this disk encryption set was updated.
	LastKeyRotationTimestamp pulumi.StringOutput `pulumi:"lastKeyRotationTimestamp"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// A readonly collection of key vault keys previously used by this disk encryption set while a key rotation is in progress. It will be empty if there is no ongoing key rotation.
	PreviousKeys KeyForDiskEncryptionSetResponseArrayOutput `pulumi:"previousKeys"`
	// The disk encryption set provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Set this flag to true to enable auto-updating of this disk encryption set to the latest key version.
	RotationToLatestKeyVersionEnabled pulumi.BoolPtrOutput `pulumi:"rotationToLatestKeyVersionEnabled"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiskEncryptionSet registers a new resource with the given unique name, arguments, and options.
func NewDiskEncryptionSet(ctx *pulumi.Context,
	name string, args *DiskEncryptionSetArgs, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211201:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220702:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230102:DiskEncryptionSet"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskEncryptionSet
	err := ctx.RegisterResource("azure-native:compute/v20230402:DiskEncryptionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskEncryptionSet gets an existing DiskEncryptionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskEncryptionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskEncryptionSetState, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	var resource DiskEncryptionSet
	err := ctx.ReadResource("azure-native:compute/v20230402:DiskEncryptionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskEncryptionSet resources.
type diskEncryptionSetState struct {
}

type DiskEncryptionSetState struct {
}

func (DiskEncryptionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetState)(nil)).Elem()
}

type diskEncryptionSetArgs struct {
	// The key vault key which is currently used by this disk encryption set.
	ActiveKey *KeyForDiskEncryptionSet `pulumi:"activeKey"`
	// The name of the disk encryption set that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskEncryptionSetName *string `pulumi:"diskEncryptionSetName"`
	// The type of key used to encrypt the data of the disk.
	EncryptionType *string `pulumi:"encryptionType"`
	// Multi-tenant application client id to access key vault in a different tenant. Setting the value to 'None' will clear the property.
	FederatedClientId *string `pulumi:"federatedClientId"`
	// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
	Identity *EncryptionSetIdentity `pulumi:"identity"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Set this flag to true to enable auto-updating of this disk encryption set to the latest key version.
	RotationToLatestKeyVersionEnabled *bool `pulumi:"rotationToLatestKeyVersionEnabled"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DiskEncryptionSet resource.
type DiskEncryptionSetArgs struct {
	// The key vault key which is currently used by this disk encryption set.
	ActiveKey KeyForDiskEncryptionSetPtrInput
	// The name of the disk encryption set that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskEncryptionSetName pulumi.StringPtrInput
	// The type of key used to encrypt the data of the disk.
	EncryptionType pulumi.StringPtrInput
	// Multi-tenant application client id to access key vault in a different tenant. Setting the value to 'None' will clear the property.
	FederatedClientId pulumi.StringPtrInput
	// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
	Identity EncryptionSetIdentityPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Set this flag to true to enable auto-updating of this disk encryption set to the latest key version.
	RotationToLatestKeyVersionEnabled pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DiskEncryptionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetArgs)(nil)).Elem()
}

type DiskEncryptionSetInput interface {
	pulumi.Input

	ToDiskEncryptionSetOutput() DiskEncryptionSetOutput
	ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput
}

func (*DiskEncryptionSet) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSet)(nil)).Elem()
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return i.ToDiskEncryptionSetOutputWithContext(context.Background())
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetOutput)
}

type DiskEncryptionSetOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSet)(nil)).Elem()
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return o
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return o
}

// The key vault key which is currently used by this disk encryption set.
func (o DiskEncryptionSetOutput) ActiveKey() KeyForDiskEncryptionSetResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) KeyForDiskEncryptionSetResponsePtrOutput { return v.ActiveKey }).(KeyForDiskEncryptionSetResponsePtrOutput)
}

// The error that was encountered during auto-key rotation. If an error is present, then auto-key rotation will not be attempted until the error on this disk encryption set is fixed.
func (o DiskEncryptionSetOutput) AutoKeyRotationError() ApiErrorResponseOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) ApiErrorResponseOutput { return v.AutoKeyRotationError }).(ApiErrorResponseOutput)
}

// The type of key used to encrypt the data of the disk.
func (o DiskEncryptionSetOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringPtrOutput { return v.EncryptionType }).(pulumi.StringPtrOutput)
}

// Multi-tenant application client id to access key vault in a different tenant. Setting the value to 'None' will clear the property.
func (o DiskEncryptionSetOutput) FederatedClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringPtrOutput { return v.FederatedClientId }).(pulumi.StringPtrOutput)
}

// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
func (o DiskEncryptionSetOutput) Identity() EncryptionSetIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) EncryptionSetIdentityResponsePtrOutput { return v.Identity }).(EncryptionSetIdentityResponsePtrOutput)
}

// The time when the active key of this disk encryption set was updated.
func (o DiskEncryptionSetOutput) LastKeyRotationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.LastKeyRotationTimestamp }).(pulumi.StringOutput)
}

// Resource location
func (o DiskEncryptionSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o DiskEncryptionSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A readonly collection of key vault keys previously used by this disk encryption set while a key rotation is in progress. It will be empty if there is no ongoing key rotation.
func (o DiskEncryptionSetOutput) PreviousKeys() KeyForDiskEncryptionSetResponseArrayOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) KeyForDiskEncryptionSetResponseArrayOutput { return v.PreviousKeys }).(KeyForDiskEncryptionSetResponseArrayOutput)
}

// The disk encryption set provisioning state.
func (o DiskEncryptionSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Set this flag to true to enable auto-updating of this disk encryption set to the latest key version.
func (o DiskEncryptionSetOutput) RotationToLatestKeyVersionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.BoolPtrOutput { return v.RotationToLatestKeyVersionEnabled }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o DiskEncryptionSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o DiskEncryptionSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskEncryptionSetOutput{})
}
