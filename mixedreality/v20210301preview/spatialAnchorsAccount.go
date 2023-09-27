// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// SpatialAnchorsAccount Response.
type SpatialAnchorsAccount struct {
	pulumi.CustomResourceState

	// Correspond domain name of certain Spatial Anchors Account
	AccountDomain pulumi.StringOutput `pulumi:"accountDomain"`
	// unique id of certain account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The identity associated with this account
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The kind of account, if supported
	Kind SkuResponsePtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The plan associated with this account
	Plan IdentityResponsePtrOutput `pulumi:"plan"`
	// The sku associated with this account
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The name of the storage account associated with this accountId
	StorageAccountName pulumi.StringPtrOutput `pulumi:"storageAccountName"`
	// System metadata for this account
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSpatialAnchorsAccount registers a new resource with the given unique name, arguments, and options.
func NewSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, args *SpatialAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mixedreality:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20190228preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20191202preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20200501:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210101:SpatialAnchorsAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SpatialAnchorsAccount
	err := ctx.RegisterResource("azure-native:mixedreality/v20210301preview:SpatialAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpatialAnchorsAccount gets an existing SpatialAnchorsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpatialAnchorsAccountState, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	var resource SpatialAnchorsAccount
	err := ctx.ReadResource("azure-native:mixedreality/v20210301preview:SpatialAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpatialAnchorsAccount resources.
type spatialAnchorsAccountState struct {
}

type SpatialAnchorsAccountState struct {
}

func (SpatialAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountState)(nil)).Elem()
}

type spatialAnchorsAccountArgs struct {
	// Name of an Mixed Reality Account.
	AccountName *string `pulumi:"accountName"`
	// The identity associated with this account
	Identity *Identity `pulumi:"identity"`
	// The kind of account, if supported
	Kind *Sku `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The plan associated with this account
	Plan *Identity `pulumi:"plan"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku associated with this account
	Sku *Sku `pulumi:"sku"`
	// The name of the storage account associated with this accountId
	StorageAccountName *string `pulumi:"storageAccountName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SpatialAnchorsAccount resource.
type SpatialAnchorsAccountArgs struct {
	// Name of an Mixed Reality Account.
	AccountName pulumi.StringPtrInput
	// The identity associated with this account
	Identity IdentityPtrInput
	// The kind of account, if supported
	Kind SkuPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The plan associated with this account
	Plan IdentityPtrInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// The sku associated with this account
	Sku SkuPtrInput
	// The name of the storage account associated with this accountId
	StorageAccountName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SpatialAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountArgs)(nil)).Elem()
}

type SpatialAnchorsAccountInput interface {
	pulumi.Input

	ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput
	ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput
}

func (*SpatialAnchorsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil)).Elem()
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return i.ToSpatialAnchorsAccountOutputWithContext(context.Background())
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountOutput)
}

func (i *SpatialAnchorsAccount) ToOutput(ctx context.Context) pulumix.Output[*SpatialAnchorsAccount] {
	return pulumix.Output[*SpatialAnchorsAccount]{
		OutputState: i.ToSpatialAnchorsAccountOutputWithContext(ctx).OutputState,
	}
}

type SpatialAnchorsAccountOutput struct{ *pulumi.OutputState }

func (SpatialAnchorsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil)).Elem()
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*SpatialAnchorsAccount] {
	return pulumix.Output[*SpatialAnchorsAccount]{
		OutputState: o.OutputState,
	}
}

// Correspond domain name of certain Spatial Anchors Account
func (o SpatialAnchorsAccountOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.AccountDomain }).(pulumi.StringOutput)
}

// unique id of certain account.
func (o SpatialAnchorsAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The identity associated with this account
func (o SpatialAnchorsAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The kind of account, if supported
func (o SpatialAnchorsAccountOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) SkuResponsePtrOutput { return v.Kind }).(SkuResponsePtrOutput)
}

// The geo-location where the resource lives
func (o SpatialAnchorsAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SpatialAnchorsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The plan associated with this account
func (o SpatialAnchorsAccountOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) IdentityResponsePtrOutput { return v.Plan }).(IdentityResponsePtrOutput)
}

// The sku associated with this account
func (o SpatialAnchorsAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The name of the storage account associated with this accountId
func (o SpatialAnchorsAccountOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringPtrOutput { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

// System metadata for this account
func (o SpatialAnchorsAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SpatialAnchorsAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SpatialAnchorsAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SpatialAnchorsAccountOutput{})
}
