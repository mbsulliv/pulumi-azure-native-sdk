// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SIM group resource.
type SimGroup struct {
	pulumi.CustomResourceState

	// A key to encrypt the SIM data that belongs to this SIM group.
	EncryptionKey KeyVaultKeyResponsePtrOutput `pulumi:"encryptionKey"`
	// The identity used to retrieve the encryption key from Azure key vault.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Mobile network that this SIM group belongs to. The mobile network must be in the same location as the SIM group.
	MobileNetwork MobileNetworkResourceIdResponsePtrOutput `pulumi:"mobileNetwork"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the SIM group resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSimGroup registers a new resource with the given unique name, arguments, and options.
func NewSimGroup(ctx *pulumi.Context,
	name string, args *SimGroupArgs, opts ...pulumi.ResourceOption) (*SimGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:SimGroup"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:SimGroup"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:SimGroup"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230901:SimGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SimGroup
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20230601:SimGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSimGroup gets an existing SimGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSimGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimGroupState, opts ...pulumi.ResourceOption) (*SimGroup, error) {
	var resource SimGroup
	err := ctx.ReadResource("azure-native:mobilenetwork/v20230601:SimGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SimGroup resources.
type simGroupState struct {
}

type SimGroupState struct {
}

func (SimGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*simGroupState)(nil)).Elem()
}

type simGroupArgs struct {
	// A key to encrypt the SIM data that belongs to this SIM group.
	EncryptionKey *KeyVaultKey `pulumi:"encryptionKey"`
	// The identity used to retrieve the encryption key from Azure key vault.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Mobile network that this SIM group belongs to. The mobile network must be in the same location as the SIM group.
	MobileNetwork *MobileNetworkResourceId `pulumi:"mobileNetwork"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SIM Group.
	SimGroupName *string `pulumi:"simGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SimGroup resource.
type SimGroupArgs struct {
	// A key to encrypt the SIM data that belongs to this SIM group.
	EncryptionKey KeyVaultKeyPtrInput
	// The identity used to retrieve the encryption key from Azure key vault.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Mobile network that this SIM group belongs to. The mobile network must be in the same location as the SIM group.
	MobileNetwork MobileNetworkResourceIdPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the SIM Group.
	SimGroupName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SimGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*simGroupArgs)(nil)).Elem()
}

type SimGroupInput interface {
	pulumi.Input

	ToSimGroupOutput() SimGroupOutput
	ToSimGroupOutputWithContext(ctx context.Context) SimGroupOutput
}

func (*SimGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SimGroup)(nil)).Elem()
}

func (i *SimGroup) ToSimGroupOutput() SimGroupOutput {
	return i.ToSimGroupOutputWithContext(context.Background())
}

func (i *SimGroup) ToSimGroupOutputWithContext(ctx context.Context) SimGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimGroupOutput)
}

type SimGroupOutput struct{ *pulumi.OutputState }

func (SimGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimGroup)(nil)).Elem()
}

func (o SimGroupOutput) ToSimGroupOutput() SimGroupOutput {
	return o
}

func (o SimGroupOutput) ToSimGroupOutputWithContext(ctx context.Context) SimGroupOutput {
	return o
}

// A key to encrypt the SIM data that belongs to this SIM group.
func (o SimGroupOutput) EncryptionKey() KeyVaultKeyResponsePtrOutput {
	return o.ApplyT(func(v *SimGroup) KeyVaultKeyResponsePtrOutput { return v.EncryptionKey }).(KeyVaultKeyResponsePtrOutput)
}

// The identity used to retrieve the encryption key from Azure key vault.
func (o SimGroupOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SimGroup) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o SimGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Mobile network that this SIM group belongs to. The mobile network must be in the same location as the SIM group.
func (o SimGroupOutput) MobileNetwork() MobileNetworkResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *SimGroup) MobileNetworkResourceIdResponsePtrOutput { return v.MobileNetwork }).(MobileNetworkResourceIdResponsePtrOutput)
}

// The name of the resource
func (o SimGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the SIM group resource.
func (o SimGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SimGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SimGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SimGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SimGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SimGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SimGroupOutput{})
}
