// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alertsmanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Action rule object containing target scope, conditions and suppression logic
// Azure REST API version: 2019-05-05-preview. Prior API version in Azure Native 1.x: 2019-05-05-preview
type ActionRuleByName struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// action rule properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewActionRuleByName registers a new resource with the given unique name, arguments, and options.
func NewActionRuleByName(ctx *pulumi.Context,
	name string, args *ActionRuleByNameArgs, opts ...pulumi.ResourceOption) (*ActionRuleByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20181102privatepreview:ActionRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190505preview:ActionRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20210808:ActionRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20210808preview:ActionRuleByName"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20230501preview:ActionRuleByName"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ActionRuleByName
	err := ctx.RegisterResource("azure-native:alertsmanagement:ActionRuleByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionRuleByName gets an existing ActionRuleByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionRuleByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionRuleByNameState, opts ...pulumi.ResourceOption) (*ActionRuleByName, error) {
	var resource ActionRuleByName
	err := ctx.ReadResource("azure-native:alertsmanagement:ActionRuleByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionRuleByName resources.
type actionRuleByNameState struct {
}

type ActionRuleByNameState struct {
}

func (ActionRuleByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleByNameState)(nil)).Elem()
}

type actionRuleByNameArgs struct {
	// The name of action rule that needs to be created/updated
	ActionRuleName *string `pulumi:"actionRuleName"`
	// Resource location
	Location *string `pulumi:"location"`
	// action rule properties
	Properties interface{} `pulumi:"properties"`
	// Resource group name where the resource is created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ActionRuleByName resource.
type ActionRuleByNameArgs struct {
	// The name of action rule that needs to be created/updated
	ActionRuleName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// action rule properties
	Properties pulumi.Input
	// Resource group name where the resource is created.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ActionRuleByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleByNameArgs)(nil)).Elem()
}

type ActionRuleByNameInput interface {
	pulumi.Input

	ToActionRuleByNameOutput() ActionRuleByNameOutput
	ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput
}

func (*ActionRuleByName) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleByName)(nil)).Elem()
}

func (i *ActionRuleByName) ToActionRuleByNameOutput() ActionRuleByNameOutput {
	return i.ToActionRuleByNameOutputWithContext(context.Background())
}

func (i *ActionRuleByName) ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleByNameOutput)
}

func (i *ActionRuleByName) ToOutput(ctx context.Context) pulumix.Output[*ActionRuleByName] {
	return pulumix.Output[*ActionRuleByName]{
		OutputState: i.ToActionRuleByNameOutputWithContext(ctx).OutputState,
	}
}

type ActionRuleByNameOutput struct{ *pulumi.OutputState }

func (ActionRuleByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleByName)(nil)).Elem()
}

func (o ActionRuleByNameOutput) ToActionRuleByNameOutput() ActionRuleByNameOutput {
	return o
}

func (o ActionRuleByNameOutput) ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput {
	return o
}

func (o ActionRuleByNameOutput) ToOutput(ctx context.Context) pulumix.Output[*ActionRuleByName] {
	return pulumix.Output[*ActionRuleByName]{
		OutputState: o.OutputState,
	}
}

// Resource location
func (o ActionRuleByNameOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionRuleByName) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o ActionRuleByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionRuleByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// action rule properties
func (o ActionRuleByNameOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ActionRuleByName) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags
func (o ActionRuleByNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ActionRuleByName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o ActionRuleByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionRuleByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionRuleByNameOutput{})
}
