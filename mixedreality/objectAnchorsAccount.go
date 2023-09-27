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

// ObjectAnchorsAccount Response.
// Azure REST API version: 2021-03-01-preview. Prior API version in Azure Native 1.x: 2021-03-01-preview
type ObjectAnchorsAccount struct {
	pulumi.CustomResourceState

	// Correspond domain name of certain Spatial Anchors Account
	AccountDomain pulumi.StringOutput `pulumi:"accountDomain"`
	// unique id of certain account.
	AccountId pulumi.StringOutput                           `pulumi:"accountId"`
	Identity  ObjectAnchorsAccountResponseIdentityPtrOutput `pulumi:"identity"`
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
	// The system metadata related to an object anchors account.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewObjectAnchorsAccount registers a new resource with the given unique name, arguments, and options.
func NewObjectAnchorsAccount(ctx *pulumi.Context,
	name string, args *ObjectAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*ObjectAnchorsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:ObjectAnchorsAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ObjectAnchorsAccount
	err := ctx.RegisterResource("azure-native:mixedreality:ObjectAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectAnchorsAccount gets an existing ObjectAnchorsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectAnchorsAccountState, opts ...pulumi.ResourceOption) (*ObjectAnchorsAccount, error) {
	var resource ObjectAnchorsAccount
	err := ctx.ReadResource("azure-native:mixedreality:ObjectAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectAnchorsAccount resources.
type objectAnchorsAccountState struct {
}

type ObjectAnchorsAccountState struct {
}

func (ObjectAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAnchorsAccountState)(nil)).Elem()
}

type objectAnchorsAccountArgs struct {
	// Name of an Mixed Reality Account.
	AccountName *string                       `pulumi:"accountName"`
	Identity    *ObjectAnchorsAccountIdentity `pulumi:"identity"`
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

// The set of arguments for constructing a ObjectAnchorsAccount resource.
type ObjectAnchorsAccountArgs struct {
	// Name of an Mixed Reality Account.
	AccountName pulumi.StringPtrInput
	Identity    ObjectAnchorsAccountIdentityPtrInput
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

func (ObjectAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAnchorsAccountArgs)(nil)).Elem()
}

type ObjectAnchorsAccountInput interface {
	pulumi.Input

	ToObjectAnchorsAccountOutput() ObjectAnchorsAccountOutput
	ToObjectAnchorsAccountOutputWithContext(ctx context.Context) ObjectAnchorsAccountOutput
}

func (*ObjectAnchorsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAnchorsAccount)(nil)).Elem()
}

func (i *ObjectAnchorsAccount) ToObjectAnchorsAccountOutput() ObjectAnchorsAccountOutput {
	return i.ToObjectAnchorsAccountOutputWithContext(context.Background())
}

func (i *ObjectAnchorsAccount) ToObjectAnchorsAccountOutputWithContext(ctx context.Context) ObjectAnchorsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAnchorsAccountOutput)
}

func (i *ObjectAnchorsAccount) ToOutput(ctx context.Context) pulumix.Output[*ObjectAnchorsAccount] {
	return pulumix.Output[*ObjectAnchorsAccount]{
		OutputState: i.ToObjectAnchorsAccountOutputWithContext(ctx).OutputState,
	}
}

type ObjectAnchorsAccountOutput struct{ *pulumi.OutputState }

func (ObjectAnchorsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAnchorsAccount)(nil)).Elem()
}

func (o ObjectAnchorsAccountOutput) ToObjectAnchorsAccountOutput() ObjectAnchorsAccountOutput {
	return o
}

func (o ObjectAnchorsAccountOutput) ToObjectAnchorsAccountOutputWithContext(ctx context.Context) ObjectAnchorsAccountOutput {
	return o
}

func (o ObjectAnchorsAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*ObjectAnchorsAccount] {
	return pulumix.Output[*ObjectAnchorsAccount]{
		OutputState: o.OutputState,
	}
}

// Correspond domain name of certain Spatial Anchors Account
func (o ObjectAnchorsAccountOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) pulumi.StringOutput { return v.AccountDomain }).(pulumi.StringOutput)
}

// unique id of certain account.
func (o ObjectAnchorsAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o ObjectAnchorsAccountOutput) Identity() ObjectAnchorsAccountResponseIdentityPtrOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) ObjectAnchorsAccountResponseIdentityPtrOutput { return v.Identity }).(ObjectAnchorsAccountResponseIdentityPtrOutput)
}

// The kind of account, if supported
func (o ObjectAnchorsAccountOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) SkuResponsePtrOutput { return v.Kind }).(SkuResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ObjectAnchorsAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ObjectAnchorsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The plan associated with this account
func (o ObjectAnchorsAccountOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) IdentityResponsePtrOutput { return v.Plan }).(IdentityResponsePtrOutput)
}

// The sku associated with this account
func (o ObjectAnchorsAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The name of the storage account associated with this accountId
func (o ObjectAnchorsAccountOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) pulumi.StringPtrOutput { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

// The system metadata related to an object anchors account.
func (o ObjectAnchorsAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ObjectAnchorsAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ObjectAnchorsAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAnchorsAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ObjectAnchorsAccountOutput{})
}
