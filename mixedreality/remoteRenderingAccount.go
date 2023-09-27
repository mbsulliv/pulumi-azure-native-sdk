// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mixedreality

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// RemoteRenderingAccount Response.
// Azure REST API version: 2021-01-01. Prior API version in Azure Native 1.x: 2021-01-01
type RemoteRenderingAccount struct {
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

// NewRemoteRenderingAccount registers a new resource with the given unique name, arguments, and options.
func NewRemoteRenderingAccount(ctx *pulumi.Context,
	name string, args *RemoteRenderingAccountArgs, opts ...pulumi.ResourceOption) (*RemoteRenderingAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mixedreality/v20191202preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20200406preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210101:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:RemoteRenderingAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RemoteRenderingAccount
	err := ctx.RegisterResource("azure-native:mixedreality:RemoteRenderingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteRenderingAccount gets an existing RemoteRenderingAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteRenderingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteRenderingAccountState, opts ...pulumi.ResourceOption) (*RemoteRenderingAccount, error) {
	var resource RemoteRenderingAccount
	err := ctx.ReadResource("azure-native:mixedreality:RemoteRenderingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteRenderingAccount resources.
type remoteRenderingAccountState struct {
}

type RemoteRenderingAccountState struct {
}

func (RemoteRenderingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRenderingAccountState)(nil)).Elem()
}

type remoteRenderingAccountArgs struct {
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

// The set of arguments for constructing a RemoteRenderingAccount resource.
type RemoteRenderingAccountArgs struct {
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

func (RemoteRenderingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRenderingAccountArgs)(nil)).Elem()
}

type RemoteRenderingAccountInput interface {
	pulumi.Input

	ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput
	ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput
}

func (*RemoteRenderingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteRenderingAccount)(nil)).Elem()
}

func (i *RemoteRenderingAccount) ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput {
	return i.ToRemoteRenderingAccountOutputWithContext(context.Background())
}

func (i *RemoteRenderingAccount) ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteRenderingAccountOutput)
}

func (i *RemoteRenderingAccount) ToOutput(ctx context.Context) pulumix.Output[*RemoteRenderingAccount] {
	return pulumix.Output[*RemoteRenderingAccount]{
		OutputState: i.ToRemoteRenderingAccountOutputWithContext(ctx).OutputState,
	}
}

type RemoteRenderingAccountOutput struct{ *pulumi.OutputState }

func (RemoteRenderingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteRenderingAccount)(nil)).Elem()
}

func (o RemoteRenderingAccountOutput) ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput {
	return o
}

func (o RemoteRenderingAccountOutput) ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput {
	return o
}

func (o RemoteRenderingAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*RemoteRenderingAccount] {
	return pulumix.Output[*RemoteRenderingAccount]{
		OutputState: o.OutputState,
	}
}

// Correspond domain name of certain Spatial Anchors Account
func (o RemoteRenderingAccountOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.AccountDomain }).(pulumi.StringOutput)
}

// unique id of certain account.
func (o RemoteRenderingAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The identity associated with this account
func (o RemoteRenderingAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The kind of account, if supported
func (o RemoteRenderingAccountOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) SkuResponsePtrOutput { return v.Kind }).(SkuResponsePtrOutput)
}

// The geo-location where the resource lives
func (o RemoteRenderingAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o RemoteRenderingAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The plan associated with this account
func (o RemoteRenderingAccountOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) IdentityResponsePtrOutput { return v.Plan }).(IdentityResponsePtrOutput)
}

// The sku associated with this account
func (o RemoteRenderingAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The name of the storage account associated with this accountId
func (o RemoteRenderingAccountOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringPtrOutput { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

// System metadata for this account
func (o RemoteRenderingAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o RemoteRenderingAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RemoteRenderingAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RemoteRenderingAccountOutput{})
}
