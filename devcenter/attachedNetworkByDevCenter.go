// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents an attached NetworkConnection.
// Azure REST API version: 2023-04-01. Prior API version in Azure Native 1.x: 2022-09-01-preview
type AttachedNetworkByDevCenter struct {
	pulumi.CustomResourceState

	// AAD Join type of the network. This is populated based on the referenced Network Connection.
	DomainJoinType pulumi.StringOutput `pulumi:"domainJoinType"`
	// Health check status values
	HealthCheckStatus pulumi.StringOutput `pulumi:"healthCheckStatus"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the NetworkConnection you want to attach.
	NetworkConnectionId pulumi.StringOutput `pulumi:"networkConnectionId"`
	// The geo-location where the NetworkConnection resource specified in 'networkConnectionResourceId' property lives.
	NetworkConnectionLocation pulumi.StringOutput `pulumi:"networkConnectionLocation"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAttachedNetworkByDevCenter registers a new resource with the given unique name, arguments, and options.
func NewAttachedNetworkByDevCenter(ctx *pulumi.Context,
	name string, args *AttachedNetworkByDevCenterArgs, opts ...pulumi.ResourceOption) (*AttachedNetworkByDevCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.NetworkConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConnectionId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:AttachedNetworkByDevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:AttachedNetworkByDevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:AttachedNetworkByDevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:AttachedNetworkByDevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:AttachedNetworkByDevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230401:AttachedNetworkByDevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230801preview:AttachedNetworkByDevCenter"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AttachedNetworkByDevCenter
	err := ctx.RegisterResource("azure-native:devcenter:AttachedNetworkByDevCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedNetworkByDevCenter gets an existing AttachedNetworkByDevCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedNetworkByDevCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedNetworkByDevCenterState, opts ...pulumi.ResourceOption) (*AttachedNetworkByDevCenter, error) {
	var resource AttachedNetworkByDevCenter
	err := ctx.ReadResource("azure-native:devcenter:AttachedNetworkByDevCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedNetworkByDevCenter resources.
type attachedNetworkByDevCenterState struct {
}

type AttachedNetworkByDevCenterState struct {
}

func (AttachedNetworkByDevCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedNetworkByDevCenterState)(nil)).Elem()
}

type attachedNetworkByDevCenterArgs struct {
	// The name of the attached NetworkConnection.
	AttachedNetworkConnectionName *string `pulumi:"attachedNetworkConnectionName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The resource ID of the NetworkConnection you want to attach.
	NetworkConnectionId string `pulumi:"networkConnectionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AttachedNetworkByDevCenter resource.
type AttachedNetworkByDevCenterArgs struct {
	// The name of the attached NetworkConnection.
	AttachedNetworkConnectionName pulumi.StringPtrInput
	// The name of the devcenter.
	DevCenterName pulumi.StringInput
	// The resource ID of the NetworkConnection you want to attach.
	NetworkConnectionId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (AttachedNetworkByDevCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedNetworkByDevCenterArgs)(nil)).Elem()
}

type AttachedNetworkByDevCenterInput interface {
	pulumi.Input

	ToAttachedNetworkByDevCenterOutput() AttachedNetworkByDevCenterOutput
	ToAttachedNetworkByDevCenterOutputWithContext(ctx context.Context) AttachedNetworkByDevCenterOutput
}

func (*AttachedNetworkByDevCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedNetworkByDevCenter)(nil)).Elem()
}

func (i *AttachedNetworkByDevCenter) ToAttachedNetworkByDevCenterOutput() AttachedNetworkByDevCenterOutput {
	return i.ToAttachedNetworkByDevCenterOutputWithContext(context.Background())
}

func (i *AttachedNetworkByDevCenter) ToAttachedNetworkByDevCenterOutputWithContext(ctx context.Context) AttachedNetworkByDevCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedNetworkByDevCenterOutput)
}

func (i *AttachedNetworkByDevCenter) ToOutput(ctx context.Context) pulumix.Output[*AttachedNetworkByDevCenter] {
	return pulumix.Output[*AttachedNetworkByDevCenter]{
		OutputState: i.ToAttachedNetworkByDevCenterOutputWithContext(ctx).OutputState,
	}
}

type AttachedNetworkByDevCenterOutput struct{ *pulumi.OutputState }

func (AttachedNetworkByDevCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedNetworkByDevCenter)(nil)).Elem()
}

func (o AttachedNetworkByDevCenterOutput) ToAttachedNetworkByDevCenterOutput() AttachedNetworkByDevCenterOutput {
	return o
}

func (o AttachedNetworkByDevCenterOutput) ToAttachedNetworkByDevCenterOutputWithContext(ctx context.Context) AttachedNetworkByDevCenterOutput {
	return o
}

func (o AttachedNetworkByDevCenterOutput) ToOutput(ctx context.Context) pulumix.Output[*AttachedNetworkByDevCenter] {
	return pulumix.Output[*AttachedNetworkByDevCenter]{
		OutputState: o.OutputState,
	}
}

// AAD Join type of the network. This is populated based on the referenced Network Connection.
func (o AttachedNetworkByDevCenterOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.DomainJoinType }).(pulumi.StringOutput)
}

// Health check status values
func (o AttachedNetworkByDevCenterOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

// The name of the resource
func (o AttachedNetworkByDevCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the NetworkConnection you want to attach.
func (o AttachedNetworkByDevCenterOutput) NetworkConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.NetworkConnectionId }).(pulumi.StringOutput)
}

// The geo-location where the NetworkConnection resource specified in 'networkConnectionResourceId' property lives.
func (o AttachedNetworkByDevCenterOutput) NetworkConnectionLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.NetworkConnectionLocation }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o AttachedNetworkByDevCenterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AttachedNetworkByDevCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AttachedNetworkByDevCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedNetworkByDevCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachedNetworkByDevCenterOutput{})
}
