// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A multi-stage process to perform update operations across members of a Fleet.
type UpdateRun struct {
	pulumi.CustomResourceState

	// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.
	ManagedClusterUpdate ManagedClusterUpdateResponseOutput `pulumi:"managedClusterUpdate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the UpdateRun resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The status of the UpdateRun.
	Status UpdateRunStatusResponseOutput `pulumi:"status"`
	// The strategy defines the order in which the clusters will be updated.
	// If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members.
	// The strategy of the UpdateRun can be modified until the run is started.
	Strategy UpdateRunStrategyResponsePtrOutput `pulumi:"strategy"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUpdateRun registers a new resource with the given unique name, arguments, and options.
func NewUpdateRun(ctx *pulumi.Context,
	name string, args *UpdateRunArgs, opts ...pulumi.ResourceOption) (*UpdateRun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FleetName == nil {
		return nil, errors.New("invalid value for required argument 'FleetName'")
	}
	if args.ManagedClusterUpdate == nil {
		return nil, errors.New("invalid value for required argument 'ManagedClusterUpdate'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230315preview:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230815preview:UpdateRun"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource UpdateRun
	err := ctx.RegisterResource("azure-native:containerservice/v20230615preview:UpdateRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpdateRun gets an existing UpdateRun resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpdateRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdateRunState, opts ...pulumi.ResourceOption) (*UpdateRun, error) {
	var resource UpdateRun
	err := ctx.ReadResource("azure-native:containerservice/v20230615preview:UpdateRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpdateRun resources.
type updateRunState struct {
}

type UpdateRunState struct {
}

func (UpdateRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*updateRunState)(nil)).Elem()
}

type updateRunArgs struct {
	// The name of the Fleet resource.
	FleetName string `pulumi:"fleetName"`
	// The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.
	ManagedClusterUpdate ManagedClusterUpdate `pulumi:"managedClusterUpdate"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The strategy defines the order in which the clusters will be updated.
	// If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members.
	// The strategy of the UpdateRun can be modified until the run is started.
	Strategy *UpdateRunStrategy `pulumi:"strategy"`
	// The name of the UpdateRun resource.
	UpdateRunName *string `pulumi:"updateRunName"`
}

// The set of arguments for constructing a UpdateRun resource.
type UpdateRunArgs struct {
	// The name of the Fleet resource.
	FleetName pulumi.StringInput
	// The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.
	ManagedClusterUpdate ManagedClusterUpdateInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The strategy defines the order in which the clusters will be updated.
	// If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members.
	// The strategy of the UpdateRun can be modified until the run is started.
	Strategy UpdateRunStrategyPtrInput
	// The name of the UpdateRun resource.
	UpdateRunName pulumi.StringPtrInput
}

func (UpdateRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updateRunArgs)(nil)).Elem()
}

type UpdateRunInput interface {
	pulumi.Input

	ToUpdateRunOutput() UpdateRunOutput
	ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput
}

func (*UpdateRun) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateRun)(nil)).Elem()
}

func (i *UpdateRun) ToUpdateRunOutput() UpdateRunOutput {
	return i.ToUpdateRunOutputWithContext(context.Background())
}

func (i *UpdateRun) ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateRunOutput)
}

func (i *UpdateRun) ToOutput(ctx context.Context) pulumix.Output[*UpdateRun] {
	return pulumix.Output[*UpdateRun]{
		OutputState: i.ToUpdateRunOutputWithContext(ctx).OutputState,
	}
}

type UpdateRunOutput struct{ *pulumi.OutputState }

func (UpdateRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateRun)(nil)).Elem()
}

func (o UpdateRunOutput) ToUpdateRunOutput() UpdateRunOutput {
	return o
}

func (o UpdateRunOutput) ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput {
	return o
}

func (o UpdateRunOutput) ToOutput(ctx context.Context) pulumix.Output[*UpdateRun] {
	return pulumix.Output[*UpdateRun]{
		OutputState: o.OutputState,
	}
}

// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o UpdateRunOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.
func (o UpdateRunOutput) ManagedClusterUpdate() ManagedClusterUpdateResponseOutput {
	return o.ApplyT(func(v *UpdateRun) ManagedClusterUpdateResponseOutput { return v.ManagedClusterUpdate }).(ManagedClusterUpdateResponseOutput)
}

// The name of the resource
func (o UpdateRunOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the UpdateRun resource.
func (o UpdateRunOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The status of the UpdateRun.
func (o UpdateRunOutput) Status() UpdateRunStatusResponseOutput {
	return o.ApplyT(func(v *UpdateRun) UpdateRunStatusResponseOutput { return v.Status }).(UpdateRunStatusResponseOutput)
}

// The strategy defines the order in which the clusters will be updated.
// If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members.
// The strategy of the UpdateRun can be modified until the run is started.
func (o UpdateRunOutput) Strategy() UpdateRunStrategyResponsePtrOutput {
	return o.ApplyT(func(v *UpdateRun) UpdateRunStrategyResponsePtrOutput { return v.Strategy }).(UpdateRunStrategyResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o UpdateRunOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UpdateRun) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o UpdateRunOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UpdateRunOutput{})
}
