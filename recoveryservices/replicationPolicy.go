// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Protection profile details.
// Azure REST API version: 2023-04-01. Prior API version in Azure Native 1.x: 2018-07-10.
//
// Other available API versions: 2023-06-01, 2023-08-01, 2024-01-01.
type ReplicationPolicy struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The custom data.
	Properties PolicyPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationPolicy registers a new resource with the given unique name, arguments, and options.
func NewReplicationPolicy(ctx *pulumi.Context,
	name string, args *ReplicationPolicyArgs, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230201:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230401:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230601:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230801:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20240101:ReplicationPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReplicationPolicy
	err := ctx.RegisterResource("azure-native:recoveryservices:ReplicationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationPolicy gets an existing ReplicationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationPolicyState, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	var resource ReplicationPolicy
	err := ctx.ReadResource("azure-native:recoveryservices:ReplicationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationPolicy resources.
type replicationPolicyState struct {
}

type ReplicationPolicyState struct {
}

func (ReplicationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyState)(nil)).Elem()
}

type replicationPolicyArgs struct {
	// Replication policy name.
	PolicyName *string `pulumi:"policyName"`
	// Policy creation properties.
	Properties *CreatePolicyInputProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationPolicy resource.
type ReplicationPolicyArgs struct {
	// Replication policy name.
	PolicyName pulumi.StringPtrInput
	// Policy creation properties.
	Properties CreatePolicyInputPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyArgs)(nil)).Elem()
}

type ReplicationPolicyInput interface {
	pulumi.Input

	ToReplicationPolicyOutput() ReplicationPolicyOutput
	ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput
}

func (*ReplicationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationPolicy)(nil)).Elem()
}

func (i *ReplicationPolicy) ToReplicationPolicyOutput() ReplicationPolicyOutput {
	return i.ToReplicationPolicyOutputWithContext(context.Background())
}

func (i *ReplicationPolicy) ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationPolicyOutput)
}

type ReplicationPolicyOutput struct{ *pulumi.OutputState }

func (ReplicationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationPolicy)(nil)).Elem()
}

func (o ReplicationPolicyOutput) ToReplicationPolicyOutput() ReplicationPolicyOutput {
	return o
}

func (o ReplicationPolicyOutput) ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput {
	return o
}

// Resource Location
func (o ReplicationPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The custom data.
func (o ReplicationPolicyOutput) Properties() PolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationPolicy) PolicyPropertiesResponseOutput { return v.Properties }).(PolicyPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationPolicyOutput{})
}
