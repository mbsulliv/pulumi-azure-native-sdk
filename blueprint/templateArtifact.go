// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Blueprint artifact that deploys a Resource Manager template.
// Azure REST API version: 2018-11-01-preview. Prior API version in Azure Native 1.x: 2018-11-01-preview
type TemplateArtifact struct {
	pulumi.CustomResourceState

	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayOutput `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Specifies the kind of blueprint artifact.
	// Expected value is 'template'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource Manager template blueprint artifact parameter values.
	Parameters ParameterValueResponseMapOutput `pulumi:"parameters"`
	// If applicable, the name of the resource group placeholder to which the Resource Manager template blueprint artifact will be deployed.
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// The Resource Manager template blueprint artifact body.
	Template pulumi.AnyOutput `pulumi:"template"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTemplateArtifact registers a new resource with the given unique name, arguments, and options.
func NewTemplateArtifact(ctx *pulumi.Context,
	name string, args *TemplateArtifactArgs, opts ...pulumi.ResourceOption) (*TemplateArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	args.Kind = pulumi.String("template")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blueprint/v20181101preview:TemplateArtifact"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TemplateArtifact
	err := ctx.RegisterResource("azure-native:blueprint:TemplateArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateArtifact gets an existing TemplateArtifact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateArtifactState, opts ...pulumi.ResourceOption) (*TemplateArtifact, error) {
	var resource TemplateArtifact
	err := ctx.ReadResource("azure-native:blueprint:TemplateArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateArtifact resources.
type templateArtifactState struct {
}

type TemplateArtifactState struct {
}

func (TemplateArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArtifactState)(nil)).Elem()
}

type templateArtifactArgs struct {
	// Name of the blueprint artifact.
	ArtifactName *string `pulumi:"artifactName"`
	// Name of the blueprint definition.
	BlueprintName string `pulumi:"blueprintName"`
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []string `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the kind of blueprint artifact.
	// Expected value is 'template'.
	Kind string `pulumi:"kind"`
	// Resource Manager template blueprint artifact parameter values.
	Parameters map[string]ParameterValue `pulumi:"parameters"`
	// If applicable, the name of the resource group placeholder to which the Resource Manager template blueprint artifact will be deployed.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope string `pulumi:"resourceScope"`
	// The Resource Manager template blueprint artifact body.
	Template interface{} `pulumi:"template"`
}

// The set of arguments for constructing a TemplateArtifact resource.
type TemplateArtifactArgs struct {
	// Name of the blueprint artifact.
	ArtifactName pulumi.StringPtrInput
	// Name of the blueprint definition.
	BlueprintName pulumi.StringInput
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Specifies the kind of blueprint artifact.
	// Expected value is 'template'.
	Kind pulumi.StringInput
	// Resource Manager template blueprint artifact parameter values.
	Parameters ParameterValueMapInput
	// If applicable, the name of the resource group placeholder to which the Resource Manager template blueprint artifact will be deployed.
	ResourceGroup pulumi.StringPtrInput
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope pulumi.StringInput
	// The Resource Manager template blueprint artifact body.
	Template pulumi.Input
}

func (TemplateArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArtifactArgs)(nil)).Elem()
}

type TemplateArtifactInput interface {
	pulumi.Input

	ToTemplateArtifactOutput() TemplateArtifactOutput
	ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput
}

func (*TemplateArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateArtifact)(nil)).Elem()
}

func (i *TemplateArtifact) ToTemplateArtifactOutput() TemplateArtifactOutput {
	return i.ToTemplateArtifactOutputWithContext(context.Background())
}

func (i *TemplateArtifact) ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArtifactOutput)
}

func (i *TemplateArtifact) ToOutput(ctx context.Context) pulumix.Output[*TemplateArtifact] {
	return pulumix.Output[*TemplateArtifact]{
		OutputState: i.ToTemplateArtifactOutputWithContext(ctx).OutputState,
	}
}

type TemplateArtifactOutput struct{ *pulumi.OutputState }

func (TemplateArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateArtifact)(nil)).Elem()
}

func (o TemplateArtifactOutput) ToTemplateArtifactOutput() TemplateArtifactOutput {
	return o
}

func (o TemplateArtifactOutput) ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput {
	return o
}

func (o TemplateArtifactOutput) ToOutput(ctx context.Context) pulumix.Output[*TemplateArtifact] {
	return pulumix.Output[*TemplateArtifact]{
		OutputState: o.OutputState,
	}
}

// Artifacts which need to be deployed before the specified artifact.
func (o TemplateArtifactOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringArrayOutput { return v.DependsOn }).(pulumi.StringArrayOutput)
}

// Multi-line explain this resource.
func (o TemplateArtifactOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o TemplateArtifactOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Specifies the kind of blueprint artifact.
// Expected value is 'template'.
func (o TemplateArtifactOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of this resource.
func (o TemplateArtifactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource Manager template blueprint artifact parameter values.
func (o TemplateArtifactOutput) Parameters() ParameterValueResponseMapOutput {
	return o.ApplyT(func(v *TemplateArtifact) ParameterValueResponseMapOutput { return v.Parameters }).(ParameterValueResponseMapOutput)
}

// If applicable, the name of the resource group placeholder to which the Resource Manager template blueprint artifact will be deployed.
func (o TemplateArtifactOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// The Resource Manager template blueprint artifact body.
func (o TemplateArtifactOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.AnyOutput { return v.Template }).(pulumi.AnyOutput)
}

// Type of this resource.
func (o TemplateArtifactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateArtifactOutput{})
}
