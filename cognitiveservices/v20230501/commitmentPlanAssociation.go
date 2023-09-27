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

// The commitment plan association.
type CommitmentPlanAssociation struct {
	pulumi.CustomResourceState

	// The Azure resource id of the account.
	AccountId pulumi.StringPtrOutput `pulumi:"accountId"`
	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCommitmentPlanAssociation registers a new resource with the given unique name, arguments, and options.
func NewCommitmentPlanAssociation(ctx *pulumi.Context,
	name string, args *CommitmentPlanAssociationArgs, opts ...pulumi.ResourceOption) (*CommitmentPlanAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommitmentPlanName == nil {
		return nil, errors.New("invalid value for required argument 'CommitmentPlanName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:CommitmentPlanAssociation"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221201:CommitmentPlanAssociation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CommitmentPlanAssociation
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20230501:CommitmentPlanAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommitmentPlanAssociation gets an existing CommitmentPlanAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommitmentPlanAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommitmentPlanAssociationState, opts ...pulumi.ResourceOption) (*CommitmentPlanAssociation, error) {
	var resource CommitmentPlanAssociation
	err := ctx.ReadResource("azure-native:cognitiveservices/v20230501:CommitmentPlanAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommitmentPlanAssociation resources.
type commitmentPlanAssociationState struct {
}

type CommitmentPlanAssociationState struct {
}

func (CommitmentPlanAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanAssociationState)(nil)).Elem()
}

type commitmentPlanAssociationArgs struct {
	// The Azure resource id of the account.
	AccountId *string `pulumi:"accountId"`
	// The name of the commitment plan association with the Cognitive Services Account
	CommitmentPlanAssociationName *string `pulumi:"commitmentPlanAssociationName"`
	// The name of the commitmentPlan associated with the Cognitive Services Account
	CommitmentPlanName string `pulumi:"commitmentPlanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a CommitmentPlanAssociation resource.
type CommitmentPlanAssociationArgs struct {
	// The Azure resource id of the account.
	AccountId pulumi.StringPtrInput
	// The name of the commitment plan association with the Cognitive Services Account
	CommitmentPlanAssociationName pulumi.StringPtrInput
	// The name of the commitmentPlan associated with the Cognitive Services Account
	CommitmentPlanName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (CommitmentPlanAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanAssociationArgs)(nil)).Elem()
}

type CommitmentPlanAssociationInput interface {
	pulumi.Input

	ToCommitmentPlanAssociationOutput() CommitmentPlanAssociationOutput
	ToCommitmentPlanAssociationOutputWithContext(ctx context.Context) CommitmentPlanAssociationOutput
}

func (*CommitmentPlanAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanAssociation)(nil)).Elem()
}

func (i *CommitmentPlanAssociation) ToCommitmentPlanAssociationOutput() CommitmentPlanAssociationOutput {
	return i.ToCommitmentPlanAssociationOutputWithContext(context.Background())
}

func (i *CommitmentPlanAssociation) ToCommitmentPlanAssociationOutputWithContext(ctx context.Context) CommitmentPlanAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanAssociationOutput)
}

func (i *CommitmentPlanAssociation) ToOutput(ctx context.Context) pulumix.Output[*CommitmentPlanAssociation] {
	return pulumix.Output[*CommitmentPlanAssociation]{
		OutputState: i.ToCommitmentPlanAssociationOutputWithContext(ctx).OutputState,
	}
}

type CommitmentPlanAssociationOutput struct{ *pulumi.OutputState }

func (CommitmentPlanAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanAssociation)(nil)).Elem()
}

func (o CommitmentPlanAssociationOutput) ToCommitmentPlanAssociationOutput() CommitmentPlanAssociationOutput {
	return o
}

func (o CommitmentPlanAssociationOutput) ToCommitmentPlanAssociationOutputWithContext(ctx context.Context) CommitmentPlanAssociationOutput {
	return o
}

func (o CommitmentPlanAssociationOutput) ToOutput(ctx context.Context) pulumix.Output[*CommitmentPlanAssociation] {
	return pulumix.Output[*CommitmentPlanAssociation]{
		OutputState: o.OutputState,
	}
}

// The Azure resource id of the account.
func (o CommitmentPlanAssociationOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringPtrOutput { return v.AccountId }).(pulumi.StringPtrOutput)
}

// Resource Etag.
func (o CommitmentPlanAssociationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource
func (o CommitmentPlanAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CommitmentPlanAssociationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CommitmentPlanAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommitmentPlanAssociationOutput{})
}
