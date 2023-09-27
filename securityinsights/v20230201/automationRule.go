// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type AutomationRule struct {
	pulumi.CustomResourceState

	// The actions to execute when the automation rule is triggered.
	Actions pulumi.ArrayOutput `pulumi:"actions"`
	// Information on the client (user or application) that made some action
	CreatedBy ClientInfoResponseOutput `pulumi:"createdBy"`
	// The time the automation rule was created.
	CreatedTimeUtc pulumi.StringOutput `pulumi:"createdTimeUtc"`
	// The display name of the automation rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Information on the client (user or application) that made some action
	LastModifiedBy ClientInfoResponseOutput `pulumi:"lastModifiedBy"`
	// The last time the automation rule was updated.
	LastModifiedTimeUtc pulumi.StringOutput `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The order of execution of the automation rule.
	Order pulumi.IntOutput `pulumi:"order"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Describes automation rule triggering logic.
	TriggeringLogic AutomationRuleTriggeringLogicResponseOutput `pulumi:"triggeringLogic"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAutomationRule registers a new resource with the given unique name, arguments, and options.
func NewAutomationRule(ctx *pulumi.Context,
	name string, args *AutomationRuleArgs, opts ...pulumi.ResourceOption) (*AutomationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TriggeringLogic == nil {
		return nil, errors.New("invalid value for required argument 'TriggeringLogic'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:AutomationRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:AutomationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AutomationRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20230201:AutomationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationRule gets an existing AutomationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationRuleState, opts ...pulumi.ResourceOption) (*AutomationRule, error) {
	var resource AutomationRule
	err := ctx.ReadResource("azure-native:securityinsights/v20230201:AutomationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationRule resources.
type automationRuleState struct {
}

type AutomationRuleState struct {
}

func (AutomationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleState)(nil)).Elem()
}

type automationRuleArgs struct {
	// The actions to execute when the automation rule is triggered.
	Actions []interface{} `pulumi:"actions"`
	// Automation rule ID
	AutomationRuleId *string `pulumi:"automationRuleId"`
	// The display name of the automation rule.
	DisplayName string `pulumi:"displayName"`
	// The order of execution of the automation rule.
	Order int `pulumi:"order"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes automation rule triggering logic.
	TriggeringLogic AutomationRuleTriggeringLogic `pulumi:"triggeringLogic"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AutomationRule resource.
type AutomationRuleArgs struct {
	// The actions to execute when the automation rule is triggered.
	Actions pulumi.ArrayInput
	// Automation rule ID
	AutomationRuleId pulumi.StringPtrInput
	// The display name of the automation rule.
	DisplayName pulumi.StringInput
	// The order of execution of the automation rule.
	Order pulumi.IntInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Describes automation rule triggering logic.
	TriggeringLogic AutomationRuleTriggeringLogicInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AutomationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleArgs)(nil)).Elem()
}

type AutomationRuleInput interface {
	pulumi.Input

	ToAutomationRuleOutput() AutomationRuleOutput
	ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput
}

func (*AutomationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRule)(nil)).Elem()
}

func (i *AutomationRule) ToAutomationRuleOutput() AutomationRuleOutput {
	return i.ToAutomationRuleOutputWithContext(context.Background())
}

func (i *AutomationRule) ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleOutput)
}

func (i *AutomationRule) ToOutput(ctx context.Context) pulumix.Output[*AutomationRule] {
	return pulumix.Output[*AutomationRule]{
		OutputState: i.ToAutomationRuleOutputWithContext(ctx).OutputState,
	}
}

type AutomationRuleOutput struct{ *pulumi.OutputState }

func (AutomationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRule)(nil)).Elem()
}

func (o AutomationRuleOutput) ToAutomationRuleOutput() AutomationRuleOutput {
	return o
}

func (o AutomationRuleOutput) ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput {
	return o
}

func (o AutomationRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*AutomationRule] {
	return pulumix.Output[*AutomationRule]{
		OutputState: o.OutputState,
	}
}

// The actions to execute when the automation rule is triggered.
func (o AutomationRuleOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.ArrayOutput { return v.Actions }).(pulumi.ArrayOutput)
}

// Information on the client (user or application) that made some action
func (o AutomationRuleOutput) CreatedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v *AutomationRule) ClientInfoResponseOutput { return v.CreatedBy }).(ClientInfoResponseOutput)
}

// The time the automation rule was created.
func (o AutomationRuleOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The display name of the automation rule.
func (o AutomationRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o AutomationRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Information on the client (user or application) that made some action
func (o AutomationRuleOutput) LastModifiedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v *AutomationRule) ClientInfoResponseOutput { return v.LastModifiedBy }).(ClientInfoResponseOutput)
}

// The last time the automation rule was updated.
func (o AutomationRuleOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o AutomationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The order of execution of the automation rule.
func (o AutomationRuleOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AutomationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AutomationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Describes automation rule triggering logic.
func (o AutomationRuleOutput) TriggeringLogic() AutomationRuleTriggeringLogicResponseOutput {
	return o.ApplyT(func(v *AutomationRule) AutomationRuleTriggeringLogicResponseOutput { return v.TriggeringLogic }).(AutomationRuleTriggeringLogicResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AutomationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutomationRuleOutput{})
}
