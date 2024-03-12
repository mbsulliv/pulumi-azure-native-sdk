// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2021-02-01.
//
// Other available API versions: 2023-01-01, 2023-04-01.
type ObjectReplicationPolicy struct {
	pulumi.CustomResourceState

	// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
	DestinationAccount pulumi.StringOutput `pulumi:"destinationAccount"`
	// Indicates when the policy is enabled on the source account.
	EnabledTime pulumi.StringOutput `pulumi:"enabledTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A unique id for object replication policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The storage account object replication rules.
	Rules ObjectReplicationPolicyRuleResponseArrayOutput `pulumi:"rules"`
	// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
	SourceAccount pulumi.StringOutput `pulumi:"sourceAccount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewObjectReplicationPolicy registers a new resource with the given unique name, arguments, and options.
func NewObjectReplicationPolicy(ctx *pulumi.Context,
	name string, args *ObjectReplicationPolicyArgs, opts ...pulumi.ResourceOption) (*ObjectReplicationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DestinationAccount == nil {
		return nil, errors.New("invalid value for required argument 'DestinationAccount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceAccount == nil {
		return nil, errors.New("invalid value for required argument 'SourceAccount'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage/v20190601:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230101:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230401:ObjectReplicationPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ObjectReplicationPolicy
	err := ctx.RegisterResource("azure-native:storage:ObjectReplicationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectReplicationPolicy gets an existing ObjectReplicationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectReplicationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectReplicationPolicyState, opts ...pulumi.ResourceOption) (*ObjectReplicationPolicy, error) {
	var resource ObjectReplicationPolicy
	err := ctx.ReadResource("azure-native:storage:ObjectReplicationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectReplicationPolicy resources.
type objectReplicationPolicyState struct {
}

type ObjectReplicationPolicyState struct {
}

func (ObjectReplicationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectReplicationPolicyState)(nil)).Elem()
}

type objectReplicationPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
	DestinationAccount string `pulumi:"destinationAccount"`
	// For the destination account, provide the value 'default'. Configure the policy on the destination account first. For the source account, provide the value of the policy ID that is returned when you download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
	ObjectReplicationPolicyId *string `pulumi:"objectReplicationPolicyId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage account object replication rules.
	Rules []ObjectReplicationPolicyRule `pulumi:"rules"`
	// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
	SourceAccount string `pulumi:"sourceAccount"`
}

// The set of arguments for constructing a ObjectReplicationPolicy resource.
type ObjectReplicationPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
	DestinationAccount pulumi.StringInput
	// For the destination account, provide the value 'default'. Configure the policy on the destination account first. For the source account, provide the value of the policy ID that is returned when you download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
	ObjectReplicationPolicyId pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The storage account object replication rules.
	Rules ObjectReplicationPolicyRuleArrayInput
	// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
	SourceAccount pulumi.StringInput
}

func (ObjectReplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectReplicationPolicyArgs)(nil)).Elem()
}

type ObjectReplicationPolicyInput interface {
	pulumi.Input

	ToObjectReplicationPolicyOutput() ObjectReplicationPolicyOutput
	ToObjectReplicationPolicyOutputWithContext(ctx context.Context) ObjectReplicationPolicyOutput
}

func (*ObjectReplicationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicy)(nil)).Elem()
}

func (i *ObjectReplicationPolicy) ToObjectReplicationPolicyOutput() ObjectReplicationPolicyOutput {
	return i.ToObjectReplicationPolicyOutputWithContext(context.Background())
}

func (i *ObjectReplicationPolicy) ToObjectReplicationPolicyOutputWithContext(ctx context.Context) ObjectReplicationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyOutput)
}

type ObjectReplicationPolicyOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicy)(nil)).Elem()
}

func (o ObjectReplicationPolicyOutput) ToObjectReplicationPolicyOutput() ObjectReplicationPolicyOutput {
	return o
}

func (o ObjectReplicationPolicyOutput) ToObjectReplicationPolicyOutputWithContext(ctx context.Context) ObjectReplicationPolicyOutput {
	return o
}

// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
func (o ObjectReplicationPolicyOutput) DestinationAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.DestinationAccount }).(pulumi.StringOutput)
}

// Indicates when the policy is enabled on the source account.
func (o ObjectReplicationPolicyOutput) EnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.EnabledTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o ObjectReplicationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A unique id for object replication policy.
func (o ObjectReplicationPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The storage account object replication rules.
func (o ObjectReplicationPolicyOutput) Rules() ObjectReplicationPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) ObjectReplicationPolicyRuleResponseArrayOutput { return v.Rules }).(ObjectReplicationPolicyRuleResponseArrayOutput)
}

// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
func (o ObjectReplicationPolicyOutput) SourceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.SourceAccount }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ObjectReplicationPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ObjectReplicationPolicyOutput{})
}
