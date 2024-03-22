// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A datastore resource
type Datastore struct {
	pulumi.CustomResourceState

	// An iSCSI volume
	DiskPoolVolume DiskPoolVolumeResponsePtrOutput `pulumi:"diskPoolVolume"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// An Azure NetApp Files volume
	NetAppVolume NetAppVolumeResponsePtrOutput `pulumi:"netAppVolume"`
	// The state of the datastore provisioning
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The operational status of the datastore
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatastore registers a new resource with the given unique name, arguments, and options.
func NewDatastore(ctx *pulumi.Context,
	name string, args *DatastoreArgs, opts ...pulumi.ResourceOption) (*Datastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DiskPoolVolume != nil {
		args.DiskPoolVolume = args.DiskPoolVolume.ToDiskPoolVolumePtrOutput().ApplyT(func(v *DiskPoolVolume) *DiskPoolVolume { return v.Defaults() }).(DiskPoolVolumePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230901:Datastore"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Datastore
	err := ctx.RegisterResource("azure-native:avs/v20230301:Datastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatastore gets an existing Datastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreState, opts ...pulumi.ResourceOption) (*Datastore, error) {
	var resource Datastore
	err := ctx.ReadResource("azure-native:avs/v20230301:Datastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datastore resources.
type datastoreState struct {
}

type DatastoreState struct {
}

func (DatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreState)(nil)).Elem()
}

type datastoreArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// Name of the datastore in the private cloud cluster
	DatastoreName *string `pulumi:"datastoreName"`
	// An iSCSI volume
	DiskPoolVolume *DiskPoolVolume `pulumi:"diskPoolVolume"`
	// An Azure NetApp Files volume
	NetAppVolume *NetAppVolume `pulumi:"netAppVolume"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Datastore resource.
type DatastoreArgs struct {
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringInput
	// Name of the datastore in the private cloud cluster
	DatastoreName pulumi.StringPtrInput
	// An iSCSI volume
	DiskPoolVolume DiskPoolVolumePtrInput
	// An Azure NetApp Files volume
	NetAppVolume NetAppVolumePtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (DatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreArgs)(nil)).Elem()
}

type DatastoreInput interface {
	pulumi.Input

	ToDatastoreOutput() DatastoreOutput
	ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput
}

func (*Datastore) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (i *Datastore) ToDatastoreOutput() DatastoreOutput {
	return i.ToDatastoreOutputWithContext(context.Background())
}

func (i *Datastore) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreOutput)
}

type DatastoreOutput struct{ *pulumi.OutputState }

func (DatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (o DatastoreOutput) ToDatastoreOutput() DatastoreOutput {
	return o
}

func (o DatastoreOutput) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return o
}

// An iSCSI volume
func (o DatastoreOutput) DiskPoolVolume() DiskPoolVolumeResponsePtrOutput {
	return o.ApplyT(func(v *Datastore) DiskPoolVolumeResponsePtrOutput { return v.DiskPoolVolume }).(DiskPoolVolumeResponsePtrOutput)
}

// Resource name.
func (o DatastoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An Azure NetApp Files volume
func (o DatastoreOutput) NetAppVolume() NetAppVolumeResponsePtrOutput {
	return o.ApplyT(func(v *Datastore) NetAppVolumeResponsePtrOutput { return v.NetAppVolume }).(NetAppVolumeResponsePtrOutput)
}

// The state of the datastore provisioning
func (o DatastoreOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The operational status of the datastore
func (o DatastoreOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource type.
func (o DatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatastoreOutput{})
}
