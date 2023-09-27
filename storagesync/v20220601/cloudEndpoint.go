// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Cloud Endpoint object.
type CloudEndpoint struct {
	pulumi.CustomResourceState

	// Azure file share name
	AzureFileShareName pulumi.StringPtrOutput `pulumi:"azureFileShareName"`
	// Backup Enabled
	BackupEnabled pulumi.StringOutput `pulumi:"backupEnabled"`
	// Cloud endpoint change enumeration status
	ChangeEnumerationStatus CloudEndpointChangeEnumerationStatusResponseOutput `pulumi:"changeEnumerationStatus"`
	// Friendly Name
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// Resource Last Operation Name
	LastOperationName pulumi.StringPtrOutput `pulumi:"lastOperationName"`
	// CloudEndpoint lastWorkflowId
	LastWorkflowId pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Partnership Id
	PartnershipId pulumi.StringPtrOutput `pulumi:"partnershipId"`
	// CloudEndpoint Provisioning State
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Storage Account Resource Id
	StorageAccountResourceId pulumi.StringPtrOutput `pulumi:"storageAccountResourceId"`
	// Storage Account Tenant Id
	StorageAccountTenantId pulumi.StringPtrOutput `pulumi:"storageAccountTenantId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCloudEndpoint registers a new resource with the given unique name, arguments, and options.
func NewCloudEndpoint(ctx *pulumi.Context,
	name string, args *CloudEndpointArgs, opts ...pulumi.ResourceOption) (*CloudEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageSyncServiceName == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncServiceName'")
	}
	if args.SyncGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SyncGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20170605preview:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200901:CloudEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CloudEndpoint
	err := ctx.RegisterResource("azure-native:storagesync/v20220601:CloudEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudEndpoint gets an existing CloudEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudEndpointState, opts ...pulumi.ResourceOption) (*CloudEndpoint, error) {
	var resource CloudEndpoint
	err := ctx.ReadResource("azure-native:storagesync/v20220601:CloudEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudEndpoint resources.
type cloudEndpointState struct {
}

type CloudEndpointState struct {
}

func (CloudEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEndpointState)(nil)).Elem()
}

type cloudEndpointArgs struct {
	// Azure file share name
	AzureFileShareName *string `pulumi:"azureFileShareName"`
	// Name of Cloud Endpoint object.
	CloudEndpointName *string `pulumi:"cloudEndpointName"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Storage Account Resource Id
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
	// Storage Account Tenant Id
	StorageAccountTenantId *string `pulumi:"storageAccountTenantId"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
}

// The set of arguments for constructing a CloudEndpoint resource.
type CloudEndpointArgs struct {
	// Azure file share name
	AzureFileShareName pulumi.StringPtrInput
	// Name of Cloud Endpoint object.
	CloudEndpointName pulumi.StringPtrInput
	// Friendly Name
	FriendlyName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Storage Account Resource Id
	StorageAccountResourceId pulumi.StringPtrInput
	// Storage Account Tenant Id
	StorageAccountTenantId pulumi.StringPtrInput
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput
	// Name of Sync Group resource.
	SyncGroupName pulumi.StringInput
}

func (CloudEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEndpointArgs)(nil)).Elem()
}

type CloudEndpointInput interface {
	pulumi.Input

	ToCloudEndpointOutput() CloudEndpointOutput
	ToCloudEndpointOutputWithContext(ctx context.Context) CloudEndpointOutput
}

func (*CloudEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudEndpoint)(nil)).Elem()
}

func (i *CloudEndpoint) ToCloudEndpointOutput() CloudEndpointOutput {
	return i.ToCloudEndpointOutputWithContext(context.Background())
}

func (i *CloudEndpoint) ToCloudEndpointOutputWithContext(ctx context.Context) CloudEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudEndpointOutput)
}

func (i *CloudEndpoint) ToOutput(ctx context.Context) pulumix.Output[*CloudEndpoint] {
	return pulumix.Output[*CloudEndpoint]{
		OutputState: i.ToCloudEndpointOutputWithContext(ctx).OutputState,
	}
}

type CloudEndpointOutput struct{ *pulumi.OutputState }

func (CloudEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudEndpoint)(nil)).Elem()
}

func (o CloudEndpointOutput) ToCloudEndpointOutput() CloudEndpointOutput {
	return o
}

func (o CloudEndpointOutput) ToCloudEndpointOutputWithContext(ctx context.Context) CloudEndpointOutput {
	return o
}

func (o CloudEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*CloudEndpoint] {
	return pulumix.Output[*CloudEndpoint]{
		OutputState: o.OutputState,
	}
}

// Azure file share name
func (o CloudEndpointOutput) AzureFileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.AzureFileShareName }).(pulumi.StringPtrOutput)
}

// Backup Enabled
func (o CloudEndpointOutput) BackupEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringOutput { return v.BackupEnabled }).(pulumi.StringOutput)
}

// Cloud endpoint change enumeration status
func (o CloudEndpointOutput) ChangeEnumerationStatus() CloudEndpointChangeEnumerationStatusResponseOutput {
	return o.ApplyT(func(v *CloudEndpoint) CloudEndpointChangeEnumerationStatusResponseOutput {
		return v.ChangeEnumerationStatus
	}).(CloudEndpointChangeEnumerationStatusResponseOutput)
}

// Friendly Name
func (o CloudEndpointOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// Resource Last Operation Name
func (o CloudEndpointOutput) LastOperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.LastOperationName }).(pulumi.StringPtrOutput)
}

// CloudEndpoint lastWorkflowId
func (o CloudEndpointOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o CloudEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Partnership Id
func (o CloudEndpointOutput) PartnershipId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.PartnershipId }).(pulumi.StringPtrOutput)
}

// CloudEndpoint Provisioning State
func (o CloudEndpointOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Storage Account Resource Id
func (o CloudEndpointOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

// Storage Account Tenant Id
func (o CloudEndpointOutput) StorageAccountTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.StorageAccountTenantId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CloudEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CloudEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CloudEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudEndpointOutput{})
}
