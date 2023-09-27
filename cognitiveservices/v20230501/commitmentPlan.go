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

// Cognitive Services account commitment plan.
type CommitmentPlan struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The Kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of Cognitive Services account commitment plan.
	Properties CommitmentPlanPropertiesResponseOutput `pulumi:"properties"`
	// The resource model definition representing SKU
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCommitmentPlan registers a new resource with the given unique name, arguments, and options.
func NewCommitmentPlan(ctx *pulumi.Context,
	name string, args *CommitmentPlanArgs, opts ...pulumi.ResourceOption) (*CommitmentPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:CommitmentPlan"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20211001:CommitmentPlan"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20220301:CommitmentPlan"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221001:CommitmentPlan"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221201:CommitmentPlan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CommitmentPlan
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20230501:CommitmentPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommitmentPlan gets an existing CommitmentPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommitmentPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommitmentPlanState, opts ...pulumi.ResourceOption) (*CommitmentPlan, error) {
	var resource CommitmentPlan
	err := ctx.ReadResource("azure-native:cognitiveservices/v20230501:CommitmentPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommitmentPlan resources.
type commitmentPlanState struct {
}

type CommitmentPlanState struct {
}

func (CommitmentPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanState)(nil)).Elem()
}

type commitmentPlanArgs struct {
	// The name of Cognitive Services account.
	AccountName string `pulumi:"accountName"`
	// The name of the commitmentPlan associated with the Cognitive Services Account
	CommitmentPlanName *string `pulumi:"commitmentPlanName"`
	// The Kind of the resource.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties of Cognitive Services account commitment plan.
	Properties *CommitmentPlanProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource model definition representing SKU
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CommitmentPlan resource.
type CommitmentPlanArgs struct {
	// The name of Cognitive Services account.
	AccountName pulumi.StringInput
	// The name of the commitmentPlan associated with the Cognitive Services Account
	CommitmentPlanName pulumi.StringPtrInput
	// The Kind of the resource.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties of Cognitive Services account commitment plan.
	Properties CommitmentPlanPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource model definition representing SKU
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CommitmentPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanArgs)(nil)).Elem()
}

type CommitmentPlanInput interface {
	pulumi.Input

	ToCommitmentPlanOutput() CommitmentPlanOutput
	ToCommitmentPlanOutputWithContext(ctx context.Context) CommitmentPlanOutput
}

func (*CommitmentPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlan)(nil)).Elem()
}

func (i *CommitmentPlan) ToCommitmentPlanOutput() CommitmentPlanOutput {
	return i.ToCommitmentPlanOutputWithContext(context.Background())
}

func (i *CommitmentPlan) ToCommitmentPlanOutputWithContext(ctx context.Context) CommitmentPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanOutput)
}

func (i *CommitmentPlan) ToOutput(ctx context.Context) pulumix.Output[*CommitmentPlan] {
	return pulumix.Output[*CommitmentPlan]{
		OutputState: i.ToCommitmentPlanOutputWithContext(ctx).OutputState,
	}
}

type CommitmentPlanOutput struct{ *pulumi.OutputState }

func (CommitmentPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlan)(nil)).Elem()
}

func (o CommitmentPlanOutput) ToCommitmentPlanOutput() CommitmentPlanOutput {
	return o
}

func (o CommitmentPlanOutput) ToCommitmentPlanOutputWithContext(ctx context.Context) CommitmentPlanOutput {
	return o
}

func (o CommitmentPlanOutput) ToOutput(ctx context.Context) pulumix.Output[*CommitmentPlan] {
	return pulumix.Output[*CommitmentPlan]{
		OutputState: o.OutputState,
	}
}

// Resource Etag.
func (o CommitmentPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The Kind of the resource.
func (o CommitmentPlanOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o CommitmentPlanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o CommitmentPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Cognitive Services account commitment plan.
func (o CommitmentPlanOutput) Properties() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *CommitmentPlan) CommitmentPlanPropertiesResponseOutput { return v.Properties }).(CommitmentPlanPropertiesResponseOutput)
}

// The resource model definition representing SKU
func (o CommitmentPlanOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *CommitmentPlan) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CommitmentPlanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CommitmentPlan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CommitmentPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CommitmentPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommitmentPlanOutput{})
}
