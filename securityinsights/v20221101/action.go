// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Action for alert rule.
type Action struct {
	pulumi.CustomResourceState

	// Etag of the action.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
	LogicAppResourceId pulumi.StringOutput `pulumi:"logicAppResourceId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The name of the logic app's workflow.
	WorkflowId pulumi.StringPtrOutput `pulumi:"workflowId"`
}

// NewAction registers a new resource with the given unique name, arguments, and options.
func NewAction(ctx *pulumi.Context,
	name string, args *ActionArgs, opts ...pulumi.ResourceOption) (*Action, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogicAppResourceId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	if args.TriggerUri == nil {
		return nil, errors.New("invalid value for required argument 'TriggerUri'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:Action"),
		},
	})
	opts = append(opts, aliases)
	var resource Action
	err := ctx.RegisterResource("azure-native:securityinsights/v20221101:Action", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAction gets an existing Action resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionState, opts ...pulumi.ResourceOption) (*Action, error) {
	var resource Action
	err := ctx.ReadResource("azure-native:securityinsights/v20221101:Action", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Action resources.
type actionState struct {
}

type ActionState struct {
}

func (ActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionState)(nil)).Elem()
}

type actionArgs struct {
	// Action ID
	ActionId *string `pulumi:"actionId"`
	// Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
	LogicAppResourceId string `pulumi:"logicAppResourceId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId string `pulumi:"ruleId"`
	// Logic App Callback URL for this specific workflow.
	TriggerUri string `pulumi:"triggerUri"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Action resource.
type ActionArgs struct {
	// Action ID
	ActionId pulumi.StringPtrInput
	// Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
	LogicAppResourceId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringInput
	// Logic App Callback URL for this specific workflow.
	TriggerUri pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionArgs)(nil)).Elem()
}

type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(ctx context.Context) ActionOutput
}

func (*Action) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (i *Action) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i *Action) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

// Etag of the action.
func (o ActionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Action) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
func (o ActionOutput) LogicAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.LogicAppResourceId }).(pulumi.StringOutput)
}

// The name of the resource
func (o ActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ActionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Action) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The name of the logic app's workflow.
func (o ActionOutput) WorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Action) pulumi.StringPtrOutput { return v.WorkflowId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
}
