// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognitiveservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cognitive Services account commitment plan.
// Azure REST API version: 2023-05-01.
//
// Other available API versions: 2023-10-01-preview.
type SharedCommitmentPlan struct {
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

// NewSharedCommitmentPlan registers a new resource with the given unique name, arguments, and options.
func NewSharedCommitmentPlan(ctx *pulumi.Context,
	name string, args *SharedCommitmentPlanArgs, opts ...pulumi.ResourceOption) (*SharedCommitmentPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221201:SharedCommitmentPlan"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20230501:SharedCommitmentPlan"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20231001preview:SharedCommitmentPlan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SharedCommitmentPlan
	err := ctx.RegisterResource("azure-native:cognitiveservices:SharedCommitmentPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedCommitmentPlan gets an existing SharedCommitmentPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedCommitmentPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedCommitmentPlanState, opts ...pulumi.ResourceOption) (*SharedCommitmentPlan, error) {
	var resource SharedCommitmentPlan
	err := ctx.ReadResource("azure-native:cognitiveservices:SharedCommitmentPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedCommitmentPlan resources.
type sharedCommitmentPlanState struct {
}

type SharedCommitmentPlanState struct {
}

func (SharedCommitmentPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedCommitmentPlanState)(nil)).Elem()
}

type sharedCommitmentPlanArgs struct {
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

// The set of arguments for constructing a SharedCommitmentPlan resource.
type SharedCommitmentPlanArgs struct {
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

func (SharedCommitmentPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedCommitmentPlanArgs)(nil)).Elem()
}

type SharedCommitmentPlanInput interface {
	pulumi.Input

	ToSharedCommitmentPlanOutput() SharedCommitmentPlanOutput
	ToSharedCommitmentPlanOutputWithContext(ctx context.Context) SharedCommitmentPlanOutput
}

func (*SharedCommitmentPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedCommitmentPlan)(nil)).Elem()
}

func (i *SharedCommitmentPlan) ToSharedCommitmentPlanOutput() SharedCommitmentPlanOutput {
	return i.ToSharedCommitmentPlanOutputWithContext(context.Background())
}

func (i *SharedCommitmentPlan) ToSharedCommitmentPlanOutputWithContext(ctx context.Context) SharedCommitmentPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedCommitmentPlanOutput)
}

type SharedCommitmentPlanOutput struct{ *pulumi.OutputState }

func (SharedCommitmentPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedCommitmentPlan)(nil)).Elem()
}

func (o SharedCommitmentPlanOutput) ToSharedCommitmentPlanOutput() SharedCommitmentPlanOutput {
	return o
}

func (o SharedCommitmentPlanOutput) ToSharedCommitmentPlanOutputWithContext(ctx context.Context) SharedCommitmentPlanOutput {
	return o
}

// Resource Etag.
func (o SharedCommitmentPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The Kind of the resource.
func (o SharedCommitmentPlanOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o SharedCommitmentPlanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SharedCommitmentPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Cognitive Services account commitment plan.
func (o SharedCommitmentPlanOutput) Properties() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) CommitmentPlanPropertiesResponseOutput { return v.Properties }).(CommitmentPlanPropertiesResponseOutput)
}

// The resource model definition representing SKU
func (o SharedCommitmentPlanOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o SharedCommitmentPlanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SharedCommitmentPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SharedCommitmentPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SharedCommitmentPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SharedCommitmentPlanOutput{})
}
