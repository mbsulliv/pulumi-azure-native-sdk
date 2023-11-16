// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure ML commitment plan resource.
type CommitmentPlan struct {
	pulumi.CustomResourceState

	// An entity tag used to enforce optimistic concurrency.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The commitment plan properties.
	Properties CommitmentPlanPropertiesResponseOutput `pulumi:"properties"`
	// The commitment plan SKU.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// User-defined tags for the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCommitmentPlan registers a new resource with the given unique name, arguments, and options.
func NewCommitmentPlan(ctx *pulumi.Context,
	name string, args *CommitmentPlanArgs, opts ...pulumi.ResourceOption) (*CommitmentPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearning:CommitmentPlan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CommitmentPlan
	err := ctx.RegisterResource("azure-native:machinelearning/v20160501preview:CommitmentPlan", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:machinelearning/v20160501preview:CommitmentPlan", name, id, state, &resource, opts...)
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
	// The Azure ML commitment plan name.
	CommitmentPlanName *string `pulumi:"commitmentPlanName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The commitment plan SKU.
	Sku *ResourceSku `pulumi:"sku"`
	// User-defined tags for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CommitmentPlan resource.
type CommitmentPlanArgs struct {
	// The Azure ML commitment plan name.
	CommitmentPlanName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The commitment plan SKU.
	Sku ResourceSkuPtrInput
	// User-defined tags for the resource.
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

// An entity tag used to enforce optimistic concurrency.
func (o CommitmentPlanOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o CommitmentPlanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o CommitmentPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The commitment plan properties.
func (o CommitmentPlanOutput) Properties() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *CommitmentPlan) CommitmentPlanPropertiesResponseOutput { return v.Properties }).(CommitmentPlanPropertiesResponseOutput)
}

// The commitment plan SKU.
func (o CommitmentPlanOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *CommitmentPlan) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// User-defined tags for the resource.
func (o CommitmentPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o CommitmentPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommitmentPlanOutput{})
}
