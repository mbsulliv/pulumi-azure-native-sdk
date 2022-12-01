// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource information, as returned by the resource provider.
type Vault struct {
	pulumi.CustomResourceState

	// Optional ETag.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Identity for the resource.
	Identity IdentityDataResponsePtrOutput `pulumi:"identity"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the vault.
	Properties VaultPropertiesResponseOutput `pulumi:"properties"`
	// Identifies the unique system identifier for each Azure resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160601:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20200202:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220131preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220930preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:Vault"),
		},
	})
	opts = append(opts, aliases)
	var resource Vault
	err := ctx.RegisterResource("azure-native:recoveryservices/v20220201:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("azure-native:recoveryservices/v20220201:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
}

type VaultState struct {
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// Identity for the resource.
	Identity *IdentityData `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Properties of the vault.
	Properties *VaultProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Identifies the unique system identifier for each Azure resource.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName *string `pulumi:"vaultName"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// Identity for the resource.
	Identity IdentityDataPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Properties of the vault.
	Properties VaultPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// Identifies the unique system identifier for each Azure resource.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the recovery services vault.
	VaultName pulumi.StringPtrInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

func (*Vault) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil)).Elem()
}

func (i *Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i *Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

type VaultOutput struct{ *pulumi.OutputState }

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil)).Elem()
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

// Optional ETag.
func (o VaultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Identity for the resource.
func (o VaultOutput) Identity() IdentityDataResponsePtrOutput {
	return o.ApplyT(func(v *Vault) IdentityDataResponsePtrOutput { return v.Identity }).(IdentityDataResponsePtrOutput)
}

// Resource location.
func (o VaultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name associated with the resource.
func (o VaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the vault.
func (o VaultOutput) Properties() VaultPropertiesResponseOutput {
	return o.ApplyT(func(v *Vault) VaultPropertiesResponseOutput { return v.Properties }).(VaultPropertiesResponseOutput)
}

// Identifies the unique system identifier for each Azure resource.
func (o VaultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Vault) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o VaultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Vault) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o VaultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o VaultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VaultOutput{})
}
