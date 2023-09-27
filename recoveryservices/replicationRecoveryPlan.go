// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Recovery plan details.
// Azure REST API version: 2023-04-01. Prior API version in Azure Native 1.x: 2018-07-10
type ReplicationRecoveryPlan struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The custom details.
	Properties RecoveryPlanPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationRecoveryPlan registers a new resource with the given unique name, arguments, and options.
func NewReplicationRecoveryPlan(ctx *pulumi.Context,
	name string, args *ReplicationRecoveryPlanArgs, opts ...pulumi.ResourceOption) (*ReplicationRecoveryPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230201:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230401:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230601:ReplicationRecoveryPlan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReplicationRecoveryPlan
	err := ctx.RegisterResource("azure-native:recoveryservices:ReplicationRecoveryPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationRecoveryPlan gets an existing ReplicationRecoveryPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationRecoveryPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationRecoveryPlanState, opts ...pulumi.ResourceOption) (*ReplicationRecoveryPlan, error) {
	var resource ReplicationRecoveryPlan
	err := ctx.ReadResource("azure-native:recoveryservices:ReplicationRecoveryPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationRecoveryPlan resources.
type replicationRecoveryPlanState struct {
}

type ReplicationRecoveryPlanState struct {
}

func (ReplicationRecoveryPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryPlanState)(nil)).Elem()
}

type replicationRecoveryPlanArgs struct {
	// Recovery plan creation properties.
	Properties CreateRecoveryPlanInputProperties `pulumi:"properties"`
	// Recovery plan name.
	RecoveryPlanName *string `pulumi:"recoveryPlanName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationRecoveryPlan resource.
type ReplicationRecoveryPlanArgs struct {
	// Recovery plan creation properties.
	Properties CreateRecoveryPlanInputPropertiesInput
	// Recovery plan name.
	RecoveryPlanName pulumi.StringPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationRecoveryPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryPlanArgs)(nil)).Elem()
}

type ReplicationRecoveryPlanInput interface {
	pulumi.Input

	ToReplicationRecoveryPlanOutput() ReplicationRecoveryPlanOutput
	ToReplicationRecoveryPlanOutputWithContext(ctx context.Context) ReplicationRecoveryPlanOutput
}

func (*ReplicationRecoveryPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationRecoveryPlan)(nil)).Elem()
}

func (i *ReplicationRecoveryPlan) ToReplicationRecoveryPlanOutput() ReplicationRecoveryPlanOutput {
	return i.ToReplicationRecoveryPlanOutputWithContext(context.Background())
}

func (i *ReplicationRecoveryPlan) ToReplicationRecoveryPlanOutputWithContext(ctx context.Context) ReplicationRecoveryPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationRecoveryPlanOutput)
}

func (i *ReplicationRecoveryPlan) ToOutput(ctx context.Context) pulumix.Output[*ReplicationRecoveryPlan] {
	return pulumix.Output[*ReplicationRecoveryPlan]{
		OutputState: i.ToReplicationRecoveryPlanOutputWithContext(ctx).OutputState,
	}
}

type ReplicationRecoveryPlanOutput struct{ *pulumi.OutputState }

func (ReplicationRecoveryPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationRecoveryPlan)(nil)).Elem()
}

func (o ReplicationRecoveryPlanOutput) ToReplicationRecoveryPlanOutput() ReplicationRecoveryPlanOutput {
	return o
}

func (o ReplicationRecoveryPlanOutput) ToReplicationRecoveryPlanOutputWithContext(ctx context.Context) ReplicationRecoveryPlanOutput {
	return o
}

func (o ReplicationRecoveryPlanOutput) ToOutput(ctx context.Context) pulumix.Output[*ReplicationRecoveryPlan] {
	return pulumix.Output[*ReplicationRecoveryPlan]{
		OutputState: o.OutputState,
	}
}

// Resource Location
func (o ReplicationRecoveryPlanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationRecoveryPlan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationRecoveryPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationRecoveryPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The custom details.
func (o ReplicationRecoveryPlanOutput) Properties() RecoveryPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationRecoveryPlan) RecoveryPlanPropertiesResponseOutput { return v.Properties }).(RecoveryPlanPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationRecoveryPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationRecoveryPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationRecoveryPlanOutput{})
}
