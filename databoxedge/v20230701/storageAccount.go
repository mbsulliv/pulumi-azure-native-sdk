// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Storage Account on the  Data Box Edge/Gateway device.
type StorageAccount struct {
	pulumi.CustomResourceState

	// BlobEndpoint of Storage Account
	BlobEndpoint pulumi.StringOutput `pulumi:"blobEndpoint"`
	// The Container Count. Present only for Storage Accounts with DataPolicy set to Cloud.
	ContainerCount pulumi.IntOutput `pulumi:"containerCount"`
	// Data policy of the storage Account.
	DataPolicy pulumi.StringOutput `pulumi:"dataPolicy"`
	// Description for the storage Account.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Storage Account Credential Id
	StorageAccountCredentialId pulumi.StringPtrOutput `pulumi:"storageAccountCredentialId"`
	// Current status of the storage account
	StorageAccountStatus pulumi.StringPtrOutput `pulumi:"storageAccountStatus"`
	// Metadata pertaining to creation and last modification of StorageAccount
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageAccount registers a new resource with the given unique name, arguments, and options.
func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataPolicy == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicy'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StorageAccount
	err := ctx.RegisterResource("azure-native:databoxedge/v20230701:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAccount gets an existing StorageAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:databoxedge/v20230701:StorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAccount resources.
type storageAccountState struct {
}

type StorageAccountState struct {
}

func (StorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountState)(nil)).Elem()
}

type storageAccountArgs struct {
	// Data policy of the storage Account.
	DataPolicy string `pulumi:"dataPolicy"`
	// Description for the storage Account.
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Storage Account Credential Id
	StorageAccountCredentialId *string `pulumi:"storageAccountCredentialId"`
	// The StorageAccount name.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// Current status of the storage account
	StorageAccountStatus *string `pulumi:"storageAccountStatus"`
}

// The set of arguments for constructing a StorageAccount resource.
type StorageAccountArgs struct {
	// Data policy of the storage Account.
	DataPolicy pulumi.StringInput
	// Description for the storage Account.
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Storage Account Credential Id
	StorageAccountCredentialId pulumi.StringPtrInput
	// The StorageAccount name.
	StorageAccountName pulumi.StringPtrInput
	// Current status of the storage account
	StorageAccountStatus pulumi.StringPtrInput
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountArgs)(nil)).Elem()
}

type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput
}

func (*StorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *StorageAccount) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i *StorageAccount) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

// BlobEndpoint of Storage Account
func (o StorageAccountOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.BlobEndpoint }).(pulumi.StringOutput)
}

// The Container Count. Present only for Storage Accounts with DataPolicy set to Cloud.
func (o StorageAccountOutput) ContainerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.IntOutput { return v.ContainerCount }).(pulumi.IntOutput)
}

// Data policy of the storage Account.
func (o StorageAccountOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.DataPolicy }).(pulumi.StringOutput)
}

// Description for the storage Account.
func (o StorageAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The object name.
func (o StorageAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Storage Account Credential Id
func (o StorageAccountOutput) StorageAccountCredentialId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.StorageAccountCredentialId }).(pulumi.StringPtrOutput)
}

// Current status of the storage account
func (o StorageAccountOutput) StorageAccountStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.StorageAccountStatus }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of StorageAccount
func (o StorageAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
}
