// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX VM Group
type WorkloadNetworkVMGroup struct {
	pulumi.CustomResourceState

	// Display name of the VM group.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Virtual machine members of this group.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// VM Group status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkVMGroup registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkVMGroup(ctx *pulumi.Context,
	name string, args *WorkloadNetworkVMGroupArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkVMGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkVMGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkVMGroup
	err := ctx.RegisterResource("azure-native:avs/v20230301:WorkloadNetworkVMGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkVMGroup gets an existing WorkloadNetworkVMGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkVMGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkVMGroupState, opts ...pulumi.ResourceOption) (*WorkloadNetworkVMGroup, error) {
	var resource WorkloadNetworkVMGroup
	err := ctx.ReadResource("azure-native:avs/v20230301:WorkloadNetworkVMGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkVMGroup resources.
type workloadNetworkVMGroupState struct {
}

type WorkloadNetworkVMGroupState struct {
}

func (WorkloadNetworkVMGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkVMGroupState)(nil)).Elem()
}

type workloadNetworkVMGroupArgs struct {
	// Display name of the VM group.
	DisplayName *string `pulumi:"displayName"`
	// Virtual machine members of this group.
	Members []string `pulumi:"members"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// NSX VM Group identifier. Generally the same as the VM Group's display name
	VmGroupId *string `pulumi:"vmGroupId"`
}

// The set of arguments for constructing a WorkloadNetworkVMGroup resource.
type WorkloadNetworkVMGroupArgs struct {
	// Display name of the VM group.
	DisplayName pulumi.StringPtrInput
	// Virtual machine members of this group.
	Members pulumi.StringArrayInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
	// NSX VM Group identifier. Generally the same as the VM Group's display name
	VmGroupId pulumi.StringPtrInput
}

func (WorkloadNetworkVMGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkVMGroupArgs)(nil)).Elem()
}

type WorkloadNetworkVMGroupInput interface {
	pulumi.Input

	ToWorkloadNetworkVMGroupOutput() WorkloadNetworkVMGroupOutput
	ToWorkloadNetworkVMGroupOutputWithContext(ctx context.Context) WorkloadNetworkVMGroupOutput
}

func (*WorkloadNetworkVMGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkVMGroup)(nil)).Elem()
}

func (i *WorkloadNetworkVMGroup) ToWorkloadNetworkVMGroupOutput() WorkloadNetworkVMGroupOutput {
	return i.ToWorkloadNetworkVMGroupOutputWithContext(context.Background())
}

func (i *WorkloadNetworkVMGroup) ToWorkloadNetworkVMGroupOutputWithContext(ctx context.Context) WorkloadNetworkVMGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkVMGroupOutput)
}

type WorkloadNetworkVMGroupOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkVMGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkVMGroup)(nil)).Elem()
}

func (o WorkloadNetworkVMGroupOutput) ToWorkloadNetworkVMGroupOutput() WorkloadNetworkVMGroupOutput {
	return o
}

func (o WorkloadNetworkVMGroupOutput) ToWorkloadNetworkVMGroupOutputWithContext(ctx context.Context) WorkloadNetworkVMGroupOutput {
	return o
}

// Display name of the VM group.
func (o WorkloadNetworkVMGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Virtual machine members of this group.
func (o WorkloadNetworkVMGroupOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o WorkloadNetworkVMGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state
func (o WorkloadNetworkVMGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o WorkloadNetworkVMGroupOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

// VM Group status.
func (o WorkloadNetworkVMGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource type.
func (o WorkloadNetworkVMGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkVMGroupOutput{})
}
