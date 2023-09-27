// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Snapshot of a Volume
type Snapshot struct {
	pulumi.CustomResourceState

	// The creation date of the snapshot
	Created pulumi.StringOutput `pulumi:"created"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// UUID v4 used to identify the Snapshot
	SnapshotId pulumi.StringOutput `pulumi:"snapshotId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101preview:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:netapp/v20230501:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:netapp/v20230501:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
}

type SnapshotState struct {
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the snapshot
	SnapshotName *string `pulumi:"snapshotName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the capacity pool
	PoolName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the snapshot
	SnapshotName pulumi.StringPtrInput
	// The name of the volume
	VolumeName pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

func (i *Snapshot) ToOutput(ctx context.Context) pulumix.Output[*Snapshot] {
	return pulumix.Output[*Snapshot]{
		OutputState: i.ToSnapshotOutputWithContext(ctx).OutputState,
	}
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToOutput(ctx context.Context) pulumix.Output[*Snapshot] {
	return pulumix.Output[*Snapshot]{
		OutputState: o.OutputState,
	}
}

// The creation date of the snapshot
func (o SnapshotOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Resource location
func (o SnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o SnapshotOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// UUID v4 used to identify the Snapshot
func (o SnapshotOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SnapshotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Snapshot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
