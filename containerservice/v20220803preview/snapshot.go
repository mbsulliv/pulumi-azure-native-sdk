// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220803preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A node pool snapshot resource.
type Snapshot struct {
	pulumi.CustomResourceState

	// CreationData to be used to specify the source agent pool resource ID to create this snapshot.
	CreationData CreationDataResponsePtrOutput `pulumi:"creationData"`
	// Whether to use a FIPS-enabled OS.
	EnableFIPS pulumi.BoolOutput `pulumi:"enableFIPS"`
	// The version of Kubernetes.
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The version of node image.
	NodeImageVersion pulumi.StringOutput `pulumi:"nodeImageVersion"`
	// Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.
	OsSku pulumi.StringOutput `pulumi:"osSku"`
	// The operating system type. The default is Linux.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The type of a snapshot. The default is NodePool.
	SnapshotType pulumi.StringPtrOutput `pulumi:"snapshotType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The size of the VM.
	VmSize pulumi.StringOutput `pulumi:"vmSize"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220901:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221102preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230102preview:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:containerservice/v20220803preview:Snapshot", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerservice/v20220803preview:Snapshot", name, id, state, &resource, opts...)
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
	// CreationData to be used to specify the source agent pool resource ID to create this snapshot.
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

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// CreationData to be used to specify the source agent pool resource ID to create this snapshot.
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

// CreationData to be used to specify the source agent pool resource ID to create this snapshot.
func (o SnapshotOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) CreationDataResponsePtrOutput { return v.CreationData }).(CreationDataResponsePtrOutput)
}

// Whether to use a FIPS-enabled OS.
func (o SnapshotOutput) EnableFIPS() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.EnableFIPS }).(pulumi.BoolOutput)
}

// The version of Kubernetes.
func (o SnapshotOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o SnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The version of node image.
func (o SnapshotOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.NodeImageVersion }).(pulumi.StringOutput)
}

// Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.
func (o SnapshotOutput) OsSku() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OsSku }).(pulumi.StringOutput)
}

// The operating system type. The default is Linux.
func (o SnapshotOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// The type of a snapshot. The default is NodePool.
func (o SnapshotOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SnapshotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Snapshot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SnapshotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The size of the VM.
func (o SnapshotOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.VmSize }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
