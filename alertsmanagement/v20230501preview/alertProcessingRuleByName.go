// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Alert processing rule object containing target scopes, conditions and scheduling logic.
type AlertProcessingRuleByName struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Alert processing rule properties.
	Properties AlertProcessingRulePropertiesResponseOutput `pulumi:"properties"`
	// Alert processing rule system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertProcessingRuleByName registers a new resource with the given unique name, arguments, and options.
func NewAlertProcessingRuleByName(ctx *pulumi.Context,
	name string, args *AlertProcessingRuleByNameArgs, opts ...pulumi.ResourceOption) (*AlertProcessingRuleByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToAlertProcessingRulePropertiesPtrOutput().ApplyT(func(v *AlertProcessingRuleProperties) *AlertProcessingRuleProperties { return v.Defaults() }).(AlertProcessingRulePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:alertsmanagement:AlertProcessingRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20181102privatepreview:AlertProcessingRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190505preview:AlertProcessingRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20210808:AlertProcessingRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20210808preview:AlertProcessingRuleByName"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AlertProcessingRuleByName
	err := ctx.RegisterResource("azure-native:alertsmanagement/v20230501preview:AlertProcessingRuleByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertProcessingRuleByName gets an existing AlertProcessingRuleByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertProcessingRuleByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertProcessingRuleByNameState, opts ...pulumi.ResourceOption) (*AlertProcessingRuleByName, error) {
	var resource AlertProcessingRuleByName
	err := ctx.ReadResource("azure-native:alertsmanagement/v20230501preview:AlertProcessingRuleByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertProcessingRuleByName resources.
type alertProcessingRuleByNameState struct {
}

type AlertProcessingRuleByNameState struct {
}

func (AlertProcessingRuleByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertProcessingRuleByNameState)(nil)).Elem()
}

type alertProcessingRuleByNameArgs struct {
	// The name of the alert processing rule that needs to be created/updated.
	AlertProcessingRuleName *string `pulumi:"alertProcessingRuleName"`
	// Resource location
	Location *string `pulumi:"location"`
	// Alert processing rule properties.
	Properties *AlertProcessingRuleProperties `pulumi:"properties"`
	// Resource group name where the resource is created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AlertProcessingRuleByName resource.
type AlertProcessingRuleByNameArgs struct {
	// The name of the alert processing rule that needs to be created/updated.
	AlertProcessingRuleName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Alert processing rule properties.
	Properties AlertProcessingRulePropertiesPtrInput
	// Resource group name where the resource is created.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (AlertProcessingRuleByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertProcessingRuleByNameArgs)(nil)).Elem()
}

type AlertProcessingRuleByNameInput interface {
	pulumi.Input

	ToAlertProcessingRuleByNameOutput() AlertProcessingRuleByNameOutput
	ToAlertProcessingRuleByNameOutputWithContext(ctx context.Context) AlertProcessingRuleByNameOutput
}

func (*AlertProcessingRuleByName) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertProcessingRuleByName)(nil)).Elem()
}

func (i *AlertProcessingRuleByName) ToAlertProcessingRuleByNameOutput() AlertProcessingRuleByNameOutput {
	return i.ToAlertProcessingRuleByNameOutputWithContext(context.Background())
}

func (i *AlertProcessingRuleByName) ToAlertProcessingRuleByNameOutputWithContext(ctx context.Context) AlertProcessingRuleByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertProcessingRuleByNameOutput)
}

type AlertProcessingRuleByNameOutput struct{ *pulumi.OutputState }

func (AlertProcessingRuleByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertProcessingRuleByName)(nil)).Elem()
}

func (o AlertProcessingRuleByNameOutput) ToAlertProcessingRuleByNameOutput() AlertProcessingRuleByNameOutput {
	return o
}

func (o AlertProcessingRuleByNameOutput) ToAlertProcessingRuleByNameOutputWithContext(ctx context.Context) AlertProcessingRuleByNameOutput {
	return o
}

// Resource location
func (o AlertProcessingRuleByNameOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertProcessingRuleByName) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o AlertProcessingRuleByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertProcessingRuleByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Alert processing rule properties.
func (o AlertProcessingRuleByNameOutput) Properties() AlertProcessingRulePropertiesResponseOutput {
	return o.ApplyT(func(v *AlertProcessingRuleByName) AlertProcessingRulePropertiesResponseOutput { return v.Properties }).(AlertProcessingRulePropertiesResponseOutput)
}

// Alert processing rule system data.
func (o AlertProcessingRuleByNameOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AlertProcessingRuleByName) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o AlertProcessingRuleByNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlertProcessingRuleByName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o AlertProcessingRuleByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertProcessingRuleByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertProcessingRuleByNameOutput{})
}
