// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sentinel onboarding state
type SentinelOnboardingState struct {
	pulumi.CustomResourceState

	// Flag that indicates the status of the CMK setting
	CustomerManagedKey pulumi.BoolPtrOutput `pulumi:"customerManagedKey"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSentinelOnboardingState registers a new resource with the given unique name, arguments, and options.
func NewSentinelOnboardingState(ctx *pulumi.Context,
	name string, args *SentinelOnboardingStateArgs, opts ...pulumi.ResourceOption) (*SentinelOnboardingState, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:SentinelOnboardingState"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SentinelOnboardingState
	err := ctx.RegisterResource("azure-native:securityinsights/v20230701preview:SentinelOnboardingState", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSentinelOnboardingState gets an existing SentinelOnboardingState resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSentinelOnboardingState(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SentinelOnboardingStateState, opts ...pulumi.ResourceOption) (*SentinelOnboardingState, error) {
	var resource SentinelOnboardingState
	err := ctx.ReadResource("azure-native:securityinsights/v20230701preview:SentinelOnboardingState", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SentinelOnboardingState resources.
type sentinelOnboardingStateState struct {
}

type SentinelOnboardingStateState struct {
}

func (SentinelOnboardingStateState) ElementType() reflect.Type {
	return reflect.TypeOf((*sentinelOnboardingStateState)(nil)).Elem()
}

type sentinelOnboardingStateArgs struct {
	// Flag that indicates the status of the CMK setting
	CustomerManagedKey *bool `pulumi:"customerManagedKey"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Sentinel onboarding state name. Supports - default
	SentinelOnboardingStateName *string `pulumi:"sentinelOnboardingStateName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SentinelOnboardingState resource.
type SentinelOnboardingStateArgs struct {
	// Flag that indicates the status of the CMK setting
	CustomerManagedKey pulumi.BoolPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The Sentinel onboarding state name. Supports - default
	SentinelOnboardingStateName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (SentinelOnboardingStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sentinelOnboardingStateArgs)(nil)).Elem()
}

type SentinelOnboardingStateInput interface {
	pulumi.Input

	ToSentinelOnboardingStateOutput() SentinelOnboardingStateOutput
	ToSentinelOnboardingStateOutputWithContext(ctx context.Context) SentinelOnboardingStateOutput
}

func (*SentinelOnboardingState) ElementType() reflect.Type {
	return reflect.TypeOf((**SentinelOnboardingState)(nil)).Elem()
}

func (i *SentinelOnboardingState) ToSentinelOnboardingStateOutput() SentinelOnboardingStateOutput {
	return i.ToSentinelOnboardingStateOutputWithContext(context.Background())
}

func (i *SentinelOnboardingState) ToSentinelOnboardingStateOutputWithContext(ctx context.Context) SentinelOnboardingStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentinelOnboardingStateOutput)
}

type SentinelOnboardingStateOutput struct{ *pulumi.OutputState }

func (SentinelOnboardingStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SentinelOnboardingState)(nil)).Elem()
}

func (o SentinelOnboardingStateOutput) ToSentinelOnboardingStateOutput() SentinelOnboardingStateOutput {
	return o
}

func (o SentinelOnboardingStateOutput) ToSentinelOnboardingStateOutputWithContext(ctx context.Context) SentinelOnboardingStateOutput {
	return o
}

// Flag that indicates the status of the CMK setting
func (o SentinelOnboardingStateOutput) CustomerManagedKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SentinelOnboardingState) pulumi.BoolPtrOutput { return v.CustomerManagedKey }).(pulumi.BoolPtrOutput)
}

// Etag of the azure resource
func (o SentinelOnboardingStateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SentinelOnboardingState) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SentinelOnboardingStateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SentinelOnboardingState) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SentinelOnboardingStateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SentinelOnboardingState) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SentinelOnboardingStateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SentinelOnboardingState) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SentinelOnboardingStateOutput{})
}
