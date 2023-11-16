// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230620preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Arc PrivateLinkScope definition.
type PrivateLinkScope struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties that define a Azure Arc PrivateLinkScope resource.
	Properties HybridComputePrivateLinkScopePropertiesResponseOutput `pulumi:"properties"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkScope registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkScope(ctx *pulumi.Context,
	name string, args *PrivateLinkScopeArgs, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220310:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220510preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220811preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221110:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221227:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221227preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20230315preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20231003preview:PrivateLinkScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkScope
	err := ctx.RegisterResource("azure-native:hybridcompute/v20230620preview:PrivateLinkScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkScope gets an existing PrivateLinkScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkScopeState, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	var resource PrivateLinkScope
	err := ctx.ReadResource("azure-native:hybridcompute/v20230620preview:PrivateLinkScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkScope resources.
type privateLinkScopeState struct {
}

type PrivateLinkScopeState struct {
}

func (PrivateLinkScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeState)(nil)).Elem()
}

type privateLinkScopeArgs struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Properties that define a Azure Arc PrivateLinkScope resource.
	Properties *HybridComputePrivateLinkScopeProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName *string `pulumi:"scopeName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateLinkScope resource.
type PrivateLinkScopeArgs struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Properties that define a Azure Arc PrivateLinkScope resource.
	Properties HybridComputePrivateLinkScopePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (PrivateLinkScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeArgs)(nil)).Elem()
}

type PrivateLinkScopeInput interface {
	pulumi.Input

	ToPrivateLinkScopeOutput() PrivateLinkScopeOutput
	ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput
}

func (*PrivateLinkScope) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return i.ToPrivateLinkScopeOutputWithContext(context.Background())
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopeOutput)
}

type PrivateLinkScopeOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return o
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return o
}

// Resource location
func (o PrivateLinkScopeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o PrivateLinkScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties that define a Azure Arc PrivateLinkScope resource.
func (o PrivateLinkScopeOutput) Properties() HybridComputePrivateLinkScopePropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkScope) HybridComputePrivateLinkScopePropertiesResponseOutput { return v.Properties }).(HybridComputePrivateLinkScopePropertiesResponseOutput)
}

// The system meta data relating to this resource.
func (o PrivateLinkScopeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkScope) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o PrivateLinkScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o PrivateLinkScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopeOutput{})
}
