// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure storage blob container data set.
type BlobContainerDataSet struct {
	pulumi.CustomResourceState

	// BLOB Container name.
	ContainerName pulumi.StringOutput `pulumi:"containerName"`
	// Unique id for identifying a data set resource
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Kind of data set.
	// Expected value is 'Container'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource group of storage account
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Storage account name of the source data set
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// Subscription id of storage account
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBlobContainerDataSet registers a new resource with the given unique name, arguments, and options.
func NewBlobContainerDataSet(ctx *pulumi.Context,
	name string, args *BlobContainerDataSetArgs, opts ...pulumi.ResourceOption) (*BlobContainerDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("Container")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:BlobContainerDataSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BlobContainerDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20210801:BlobContainerDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobContainerDataSet gets an existing BlobContainerDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobContainerDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerDataSetState, opts ...pulumi.ResourceOption) (*BlobContainerDataSet, error) {
	var resource BlobContainerDataSet
	err := ctx.ReadResource("azure-native:datashare/v20210801:BlobContainerDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobContainerDataSet resources.
type blobContainerDataSetState struct {
}

type BlobContainerDataSetState struct {
}

func (BlobContainerDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetState)(nil)).Elem()
}

type blobContainerDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// BLOB Container name.
	ContainerName string `pulumi:"containerName"`
	// The name of the dataSet.
	DataSetName *string `pulumi:"dataSetName"`
	// Kind of data set.
	// Expected value is 'Container'.
	Kind string `pulumi:"kind"`
	// Resource group of storage account
	ResourceGroup string `pulumi:"resourceGroup"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share to add the data set to.
	ShareName string `pulumi:"shareName"`
	// Storage account name of the source data set
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a BlobContainerDataSet resource.
type BlobContainerDataSetArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// BLOB Container name.
	ContainerName pulumi.StringInput
	// The name of the dataSet.
	DataSetName pulumi.StringPtrInput
	// Kind of data set.
	// Expected value is 'Container'.
	Kind pulumi.StringInput
	// Resource group of storage account
	ResourceGroup pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share to add the data set to.
	ShareName pulumi.StringInput
	// Storage account name of the source data set
	StorageAccountName pulumi.StringInput
	// Subscription id of storage account
	SubscriptionId pulumi.StringInput
}

func (BlobContainerDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetArgs)(nil)).Elem()
}

type BlobContainerDataSetInput interface {
	pulumi.Input

	ToBlobContainerDataSetOutput() BlobContainerDataSetOutput
	ToBlobContainerDataSetOutputWithContext(ctx context.Context) BlobContainerDataSetOutput
}

func (*BlobContainerDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerDataSet)(nil)).Elem()
}

func (i *BlobContainerDataSet) ToBlobContainerDataSetOutput() BlobContainerDataSetOutput {
	return i.ToBlobContainerDataSetOutputWithContext(context.Background())
}

func (i *BlobContainerDataSet) ToBlobContainerDataSetOutputWithContext(ctx context.Context) BlobContainerDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobContainerDataSetOutput)
}

type BlobContainerDataSetOutput struct{ *pulumi.OutputState }

func (BlobContainerDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerDataSet)(nil)).Elem()
}

func (o BlobContainerDataSetOutput) ToBlobContainerDataSetOutput() BlobContainerDataSetOutput {
	return o
}

func (o BlobContainerDataSetOutput) ToBlobContainerDataSetOutputWithContext(ctx context.Context) BlobContainerDataSetOutput {
	return o
}

// BLOB Container name.
func (o BlobContainerDataSetOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.ContainerName }).(pulumi.StringOutput)
}

// Unique id for identifying a data set resource
func (o BlobContainerDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'Container'.
func (o BlobContainerDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o BlobContainerDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource group of storage account
func (o BlobContainerDataSetOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set
func (o BlobContainerDataSetOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account
func (o BlobContainerDataSetOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o BlobContainerDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o BlobContainerDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobContainerDataSetOutput{})
}
