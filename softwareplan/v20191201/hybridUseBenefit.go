// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response on GET of a hybrid use benefit
type HybridUseBenefit struct {
	pulumi.CustomResourceState

	// Created date
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Indicates the revision of the hybrid use benefit
	Etag pulumi.IntOutput `pulumi:"etag"`
	// Last updated date
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Hybrid use benefit SKU
	Sku SkuResponseOutput `pulumi:"sku"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHybridUseBenefit registers a new resource with the given unique name, arguments, and options.
func NewHybridUseBenefit(ctx *pulumi.Context,
	name string, args *HybridUseBenefitArgs, opts ...pulumi.ResourceOption) (*HybridUseBenefit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:softwareplan:HybridUseBenefit"),
		},
		{
			Type: pulumi.String("azure-native:softwareplan/v20190601preview:HybridUseBenefit"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HybridUseBenefit
	err := ctx.RegisterResource("azure-native:softwareplan/v20191201:HybridUseBenefit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridUseBenefit gets an existing HybridUseBenefit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridUseBenefit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridUseBenefitState, opts ...pulumi.ResourceOption) (*HybridUseBenefit, error) {
	var resource HybridUseBenefit
	err := ctx.ReadResource("azure-native:softwareplan/v20191201:HybridUseBenefit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridUseBenefit resources.
type hybridUseBenefitState struct {
}

type HybridUseBenefitState struct {
}

func (HybridUseBenefitState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridUseBenefitState)(nil)).Elem()
}

type hybridUseBenefitArgs struct {
	// This is a unique identifier for a plan. Should be a guid.
	PlanId *string `pulumi:"planId"`
	// The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
	Scope string `pulumi:"scope"`
	// Hybrid use benefit SKU
	Sku Sku `pulumi:"sku"`
}

// The set of arguments for constructing a HybridUseBenefit resource.
type HybridUseBenefitArgs struct {
	// This is a unique identifier for a plan. Should be a guid.
	PlanId pulumi.StringPtrInput
	// The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
	Scope pulumi.StringInput
	// Hybrid use benefit SKU
	Sku SkuInput
}

func (HybridUseBenefitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridUseBenefitArgs)(nil)).Elem()
}

type HybridUseBenefitInput interface {
	pulumi.Input

	ToHybridUseBenefitOutput() HybridUseBenefitOutput
	ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput
}

func (*HybridUseBenefit) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridUseBenefit)(nil)).Elem()
}

func (i *HybridUseBenefit) ToHybridUseBenefitOutput() HybridUseBenefitOutput {
	return i.ToHybridUseBenefitOutputWithContext(context.Background())
}

func (i *HybridUseBenefit) ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridUseBenefitOutput)
}

func (i *HybridUseBenefit) ToOutput(ctx context.Context) pulumix.Output[*HybridUseBenefit] {
	return pulumix.Output[*HybridUseBenefit]{
		OutputState: i.ToHybridUseBenefitOutputWithContext(ctx).OutputState,
	}
}

type HybridUseBenefitOutput struct{ *pulumi.OutputState }

func (HybridUseBenefitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridUseBenefit)(nil)).Elem()
}

func (o HybridUseBenefitOutput) ToHybridUseBenefitOutput() HybridUseBenefitOutput {
	return o
}

func (o HybridUseBenefitOutput) ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput {
	return o
}

func (o HybridUseBenefitOutput) ToOutput(ctx context.Context) pulumix.Output[*HybridUseBenefit] {
	return pulumix.Output[*HybridUseBenefit]{
		OutputState: o.OutputState,
	}
}

// Created date
func (o HybridUseBenefitOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// Indicates the revision of the hybrid use benefit
func (o HybridUseBenefitOutput) Etag() pulumi.IntOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.IntOutput { return v.Etag }).(pulumi.IntOutput)
}

// Last updated date
func (o HybridUseBenefitOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

// The name of the resource
func (o HybridUseBenefitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state
func (o HybridUseBenefitOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Hybrid use benefit SKU
func (o HybridUseBenefitOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *HybridUseBenefit) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HybridUseBenefitOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridUseBenefitOutput{})
}
