// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220402preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A managed cluster snapshot resource.
type ManagedClusterSnapshot struct {
	pulumi.CustomResourceState

	// CreationData to be used to specify the source resource ID to create this snapshot.
	CreationData CreationDataResponsePtrOutput `pulumi:"creationData"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// What the properties will be showed when getting managed cluster snapshot. Those properties are read-only.
	ManagedClusterPropertiesReadOnly ManagedClusterPropertiesForSnapshotResponseOutput `pulumi:"managedClusterPropertiesReadOnly"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of a snapshot. The default is NodePool.
	SnapshotType pulumi.StringPtrOutput `pulumi:"snapshotType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedClusterSnapshot registers a new resource with the given unique name, arguments, and options.
func NewManagedClusterSnapshot(ctx *pulumi.Context,
	name string, args *ManagedClusterSnapshotArgs, opts ...pulumi.ResourceOption) (*ManagedClusterSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221102preview:ManagedClusterSnapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230102preview:ManagedClusterSnapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedClusterSnapshot
	err := ctx.RegisterResource("azure-native:containerservice/v20220402preview:ManagedClusterSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedClusterSnapshot gets an existing ManagedClusterSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedClusterSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterSnapshotState, opts ...pulumi.ResourceOption) (*ManagedClusterSnapshot, error) {
	var resource ManagedClusterSnapshot
	err := ctx.ReadResource("azure-native:containerservice/v20220402preview:ManagedClusterSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedClusterSnapshot resources.
type managedClusterSnapshotState struct {
}

type ManagedClusterSnapshotState struct {
}

func (ManagedClusterSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterSnapshotState)(nil)).Elem()
}

type managedClusterSnapshotArgs struct {
	// CreationData to be used to specify the source resource ID to create this snapshot.
	CreationData *CreationData `pulumi:"creationData"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName *string `pulumi:"resourceName"`
	// The type of a snapshot. The default is NodePool.
	SnapshotType *string `pulumi:"snapshotType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedClusterSnapshot resource.
type ManagedClusterSnapshotArgs struct {
	// CreationData to be used to specify the source resource ID to create this snapshot.
	CreationData CreationDataPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringPtrInput
	// The type of a snapshot. The default is NodePool.
	SnapshotType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ManagedClusterSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterSnapshotArgs)(nil)).Elem()
}

type ManagedClusterSnapshotInput interface {
	pulumi.Input

	ToManagedClusterSnapshotOutput() ManagedClusterSnapshotOutput
	ToManagedClusterSnapshotOutputWithContext(ctx context.Context) ManagedClusterSnapshotOutput
}

func (*ManagedClusterSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSnapshot)(nil)).Elem()
}

func (i *ManagedClusterSnapshot) ToManagedClusterSnapshotOutput() ManagedClusterSnapshotOutput {
	return i.ToManagedClusterSnapshotOutputWithContext(context.Background())
}

func (i *ManagedClusterSnapshot) ToManagedClusterSnapshotOutputWithContext(ctx context.Context) ManagedClusterSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSnapshotOutput)
}

type ManagedClusterSnapshotOutput struct{ *pulumi.OutputState }

func (ManagedClusterSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSnapshot)(nil)).Elem()
}

func (o ManagedClusterSnapshotOutput) ToManagedClusterSnapshotOutput() ManagedClusterSnapshotOutput {
	return o
}

func (o ManagedClusterSnapshotOutput) ToManagedClusterSnapshotOutputWithContext(ctx context.Context) ManagedClusterSnapshotOutput {
	return o
}

// CreationData to be used to specify the source resource ID to create this snapshot.
func (o ManagedClusterSnapshotOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) CreationDataResponsePtrOutput { return v.CreationData }).(CreationDataResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ManagedClusterSnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// What the properties will be showed when getting managed cluster snapshot. Those properties are read-only.
func (o ManagedClusterSnapshotOutput) ManagedClusterPropertiesReadOnly() ManagedClusterPropertiesForSnapshotResponseOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) ManagedClusterPropertiesForSnapshotResponseOutput {
		return v.ManagedClusterPropertiesReadOnly
	}).(ManagedClusterPropertiesForSnapshotResponseOutput)
}

// The name of the resource
func (o ManagedClusterSnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of a snapshot. The default is NodePool.
func (o ManagedClusterSnapshotOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringPtrOutput { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ManagedClusterSnapshotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ManagedClusterSnapshotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagedClusterSnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterSnapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterSnapshotOutput{})
}
