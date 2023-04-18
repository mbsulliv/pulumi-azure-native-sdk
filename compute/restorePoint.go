// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Restore Point details.
// API Version: 2021-03-01.
type RestorePoint struct {
	pulumi.CustomResourceState

	// Gets the consistency mode for the restore point. Please refer to https://aka.ms/RestorePoints for more details.
	ConsistencyMode pulumi.StringOutput `pulumi:"consistencyMode"`
	// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
	ExcludeDisks ApiEntityReferenceResponseArrayOutput `pulumi:"excludeDisks"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state of the restore point.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets the details of the VM captured at the time of the restore point creation.
	SourceMetadata RestorePointSourceMetadataResponseOutput `pulumi:"sourceMetadata"`
	// Gets the creation time of the restore point.
	TimeCreated pulumi.StringPtrOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRestorePoint registers a new resource with the given unique name, arguments, and options.
func NewRestorePoint(ctx *pulumi.Context,
	name string, args *RestorePointArgs, opts ...pulumi.ResourceOption) (*RestorePoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RestorePointCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RestorePointCollectionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20210301:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:RestorePoint"),
		},
	})
	opts = append(opts, aliases)
	var resource RestorePoint
	err := ctx.RegisterResource("azure-native:compute:RestorePoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestorePoint gets an existing RestorePoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestorePoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestorePointState, opts ...pulumi.ResourceOption) (*RestorePoint, error) {
	var resource RestorePoint
	err := ctx.ReadResource("azure-native:compute:RestorePoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestorePoint resources.
type restorePointState struct {
}

type RestorePointState struct {
}

func (RestorePointState) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointState)(nil)).Elem()
}

type restorePointArgs struct {
	// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
	ExcludeDisks []ApiEntityReference `pulumi:"excludeDisks"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName string `pulumi:"restorePointCollectionName"`
	// The name of the restore point.
	RestorePointName *string `pulumi:"restorePointName"`
	// Gets the creation time of the restore point.
	TimeCreated *string `pulumi:"timeCreated"`
}

// The set of arguments for constructing a RestorePoint resource.
type RestorePointArgs struct {
	// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
	ExcludeDisks ApiEntityReferenceArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the restore point collection.
	RestorePointCollectionName pulumi.StringInput
	// The name of the restore point.
	RestorePointName pulumi.StringPtrInput
	// Gets the creation time of the restore point.
	TimeCreated pulumi.StringPtrInput
}

func (RestorePointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointArgs)(nil)).Elem()
}

type RestorePointInput interface {
	pulumi.Input

	ToRestorePointOutput() RestorePointOutput
	ToRestorePointOutputWithContext(ctx context.Context) RestorePointOutput
}

func (*RestorePoint) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePoint)(nil)).Elem()
}

func (i *RestorePoint) ToRestorePointOutput() RestorePointOutput {
	return i.ToRestorePointOutputWithContext(context.Background())
}

func (i *RestorePoint) ToRestorePointOutputWithContext(ctx context.Context) RestorePointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePointOutput)
}

type RestorePointOutput struct{ *pulumi.OutputState }

func (RestorePointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePoint)(nil)).Elem()
}

func (o RestorePointOutput) ToRestorePointOutput() RestorePointOutput {
	return o
}

func (o RestorePointOutput) ToRestorePointOutputWithContext(ctx context.Context) RestorePointOutput {
	return o
}

// Gets the consistency mode for the restore point. Please refer to https://aka.ms/RestorePoints for more details.
func (o RestorePointOutput) ConsistencyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.ConsistencyMode }).(pulumi.StringOutput)
}

// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
func (o RestorePointOutput) ExcludeDisks() ApiEntityReferenceResponseArrayOutput {
	return o.ApplyT(func(v *RestorePoint) ApiEntityReferenceResponseArrayOutput { return v.ExcludeDisks }).(ApiEntityReferenceResponseArrayOutput)
}

// Resource name
func (o RestorePointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the restore point.
func (o RestorePointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the details of the VM captured at the time of the restore point creation.
func (o RestorePointOutput) SourceMetadata() RestorePointSourceMetadataResponseOutput {
	return o.ApplyT(func(v *RestorePoint) RestorePointSourceMetadataResponseOutput { return v.SourceMetadata }).(RestorePointSourceMetadataResponseOutput)
}

// Gets the creation time of the restore point.
func (o RestorePointOutput) TimeCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringPtrOutput { return v.TimeCreated }).(pulumi.StringPtrOutput)
}

// Resource type
func (o RestorePointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RestorePointOutput{})
}
