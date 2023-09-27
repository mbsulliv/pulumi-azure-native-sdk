// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type SkusNestedResourceTypeThird struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSkusNestedResourceTypeThird registers a new resource with the given unique name, arguments, and options.
func NewSkusNestedResourceTypeThird(ctx *pulumi.Context,
	name string, args *SkusNestedResourceTypeThirdArgs, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeThird, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NestedResourceTypeFirst == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeFirst'")
	}
	if args.NestedResourceTypeSecond == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeSecond'")
	}
	if args.NestedResourceTypeThird == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeThird'")
	}
	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:SkusNestedResourceTypeThird"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SkusNestedResourceTypeThird
	err := ctx.RegisterResource("azure-native:providerhub/v20210901preview:SkusNestedResourceTypeThird", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSkusNestedResourceTypeThird gets an existing SkusNestedResourceTypeThird resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSkusNestedResourceTypeThird(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusNestedResourceTypeThirdState, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeThird, error) {
	var resource SkusNestedResourceTypeThird
	err := ctx.ReadResource("azure-native:providerhub/v20210901preview:SkusNestedResourceTypeThird", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SkusNestedResourceTypeThird resources.
type skusNestedResourceTypeThirdState struct {
}

type SkusNestedResourceTypeThirdState struct {
}

func (SkusNestedResourceTypeThirdState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeThirdState)(nil)).Elem()
}

type skusNestedResourceTypeThirdArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst string `pulumi:"nestedResourceTypeFirst"`
	// The second child resource type.
	NestedResourceTypeSecond string `pulumi:"nestedResourceTypeSecond"`
	// The third child resource type.
	NestedResourceTypeThird string                 `pulumi:"nestedResourceTypeThird"`
	Properties              *SkuResourceProperties `pulumi:"properties"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType string `pulumi:"resourceType"`
	// The SKU.
	Sku *string `pulumi:"sku"`
}

// The set of arguments for constructing a SkusNestedResourceTypeThird resource.
type SkusNestedResourceTypeThirdArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst pulumi.StringInput
	// The second child resource type.
	NestedResourceTypeSecond pulumi.StringInput
	// The third child resource type.
	NestedResourceTypeThird pulumi.StringInput
	Properties              SkuResourcePropertiesPtrInput
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput
	// The resource type.
	ResourceType pulumi.StringInput
	// The SKU.
	Sku pulumi.StringPtrInput
}

func (SkusNestedResourceTypeThirdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeThirdArgs)(nil)).Elem()
}

type SkusNestedResourceTypeThirdInput interface {
	pulumi.Input

	ToSkusNestedResourceTypeThirdOutput() SkusNestedResourceTypeThirdOutput
	ToSkusNestedResourceTypeThirdOutputWithContext(ctx context.Context) SkusNestedResourceTypeThirdOutput
}

func (*SkusNestedResourceTypeThird) ElementType() reflect.Type {
	return reflect.TypeOf((**SkusNestedResourceTypeThird)(nil)).Elem()
}

func (i *SkusNestedResourceTypeThird) ToSkusNestedResourceTypeThirdOutput() SkusNestedResourceTypeThirdOutput {
	return i.ToSkusNestedResourceTypeThirdOutputWithContext(context.Background())
}

func (i *SkusNestedResourceTypeThird) ToSkusNestedResourceTypeThirdOutputWithContext(ctx context.Context) SkusNestedResourceTypeThirdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusNestedResourceTypeThirdOutput)
}

func (i *SkusNestedResourceTypeThird) ToOutput(ctx context.Context) pulumix.Output[*SkusNestedResourceTypeThird] {
	return pulumix.Output[*SkusNestedResourceTypeThird]{
		OutputState: i.ToSkusNestedResourceTypeThirdOutputWithContext(ctx).OutputState,
	}
}

type SkusNestedResourceTypeThirdOutput struct{ *pulumi.OutputState }

func (SkusNestedResourceTypeThirdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkusNestedResourceTypeThird)(nil)).Elem()
}

func (o SkusNestedResourceTypeThirdOutput) ToSkusNestedResourceTypeThirdOutput() SkusNestedResourceTypeThirdOutput {
	return o
}

func (o SkusNestedResourceTypeThirdOutput) ToSkusNestedResourceTypeThirdOutputWithContext(ctx context.Context) SkusNestedResourceTypeThirdOutput {
	return o
}

func (o SkusNestedResourceTypeThirdOutput) ToOutput(ctx context.Context) pulumix.Output[*SkusNestedResourceTypeThird] {
	return pulumix.Output[*SkusNestedResourceTypeThird]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o SkusNestedResourceTypeThirdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeThird) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SkusNestedResourceTypeThirdOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeThird) SkuResourceResponsePropertiesOutput { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o SkusNestedResourceTypeThirdOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeThird) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SkusNestedResourceTypeThirdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeThird) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SkusNestedResourceTypeThirdOutput{})
}
