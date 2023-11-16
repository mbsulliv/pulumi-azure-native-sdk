// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Blueprint definition.
type Blueprint struct {
	pulumi.CustomResourceState

	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Layout view of the blueprint definition for UI reference.
	Layout pulumi.AnyOutput `pulumi:"layout"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters required by this blueprint definition.
	Parameters ParameterDefinitionResponseMapOutput `pulumi:"parameters"`
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups ResourceGroupDefinitionResponseMapOutput `pulumi:"resourceGroups"`
	// Status of the blueprint. This field is readonly.
	Status BlueprintStatusResponseOutput `pulumi:"status"`
	// The scope where this blueprint definition can be assigned.
	TargetScope pulumi.StringOutput `pulumi:"targetScope"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Published versions of this blueprint definition.
	Versions pulumi.AnyOutput `pulumi:"versions"`
}

// NewBlueprint registers a new resource with the given unique name, arguments, and options.
func NewBlueprint(ctx *pulumi.Context,
	name string, args *BlueprintArgs, opts ...pulumi.ResourceOption) (*Blueprint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	if args.TargetScope == nil {
		return nil, errors.New("invalid value for required argument 'TargetScope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blueprint:Blueprint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Blueprint
	err := ctx.RegisterResource("azure-native:blueprint/v20181101preview:Blueprint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlueprint gets an existing Blueprint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlueprint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlueprintState, opts ...pulumi.ResourceOption) (*Blueprint, error) {
	var resource Blueprint
	err := ctx.ReadResource("azure-native:blueprint/v20181101preview:Blueprint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Blueprint resources.
type blueprintState struct {
}

type BlueprintState struct {
}

func (BlueprintState) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintState)(nil)).Elem()
}

type blueprintArgs struct {
	// Name of the blueprint definition.
	BlueprintName *string `pulumi:"blueprintName"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Parameters required by this blueprint definition.
	Parameters map[string]ParameterDefinition `pulumi:"parameters"`
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups map[string]ResourceGroupDefinition `pulumi:"resourceGroups"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope string `pulumi:"resourceScope"`
	// The scope where this blueprint definition can be assigned.
	TargetScope string `pulumi:"targetScope"`
	// Published versions of this blueprint definition.
	Versions interface{} `pulumi:"versions"`
}

// The set of arguments for constructing a Blueprint resource.
type BlueprintArgs struct {
	// Name of the blueprint definition.
	BlueprintName pulumi.StringPtrInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Parameters required by this blueprint definition.
	Parameters ParameterDefinitionMapInput
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups ResourceGroupDefinitionMapInput
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope pulumi.StringInput
	// The scope where this blueprint definition can be assigned.
	TargetScope pulumi.StringInput
	// Published versions of this blueprint definition.
	Versions pulumi.Input
}

func (BlueprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintArgs)(nil)).Elem()
}

type BlueprintInput interface {
	pulumi.Input

	ToBlueprintOutput() BlueprintOutput
	ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput
}

func (*Blueprint) ElementType() reflect.Type {
	return reflect.TypeOf((**Blueprint)(nil)).Elem()
}

func (i *Blueprint) ToBlueprintOutput() BlueprintOutput {
	return i.ToBlueprintOutputWithContext(context.Background())
}

func (i *Blueprint) ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintOutput)
}

type BlueprintOutput struct{ *pulumi.OutputState }

func (BlueprintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Blueprint)(nil)).Elem()
}

func (o BlueprintOutput) ToBlueprintOutput() BlueprintOutput {
	return o
}

func (o BlueprintOutput) ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput {
	return o
}

// Multi-line explain this resource.
func (o BlueprintOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o BlueprintOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Layout view of the blueprint definition for UI reference.
func (o BlueprintOutput) Layout() pulumi.AnyOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.AnyOutput { return v.Layout }).(pulumi.AnyOutput)
}

// Name of this resource.
func (o BlueprintOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Parameters required by this blueprint definition.
func (o BlueprintOutput) Parameters() ParameterDefinitionResponseMapOutput {
	return o.ApplyT(func(v *Blueprint) ParameterDefinitionResponseMapOutput { return v.Parameters }).(ParameterDefinitionResponseMapOutput)
}

// Resource group placeholders defined by this blueprint definition.
func (o BlueprintOutput) ResourceGroups() ResourceGroupDefinitionResponseMapOutput {
	return o.ApplyT(func(v *Blueprint) ResourceGroupDefinitionResponseMapOutput { return v.ResourceGroups }).(ResourceGroupDefinitionResponseMapOutput)
}

// Status of the blueprint. This field is readonly.
func (o BlueprintOutput) Status() BlueprintStatusResponseOutput {
	return o.ApplyT(func(v *Blueprint) BlueprintStatusResponseOutput { return v.Status }).(BlueprintStatusResponseOutput)
}

// The scope where this blueprint definition can be assigned.
func (o BlueprintOutput) TargetScope() pulumi.StringOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringOutput { return v.TargetScope }).(pulumi.StringOutput)
}

// Type of this resource.
func (o BlueprintOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Published versions of this blueprint definition.
func (o BlueprintOutput) Versions() pulumi.AnyOutput {
	return o.ApplyT(func(v *Blueprint) pulumi.AnyOutput { return v.Versions }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(BlueprintOutput{})
}
