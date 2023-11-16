// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the action of alert rule.
func LookupAction(ctx *pulumi.Context, args *LookupActionArgs, opts ...pulumi.InvokeOption) (*LookupActionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupActionResult
	err := ctx.Invoke("azure-native:securityinsights/v20230801preview:getAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActionArgs struct {
	// Action ID
	ActionId string `pulumi:"actionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Action for alert rule.
type LookupActionResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
	LogicAppResourceId string `pulumi:"logicAppResourceId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The name of the logic app's workflow.
	WorkflowId *string `pulumi:"workflowId"`
}

func LookupActionOutput(ctx *pulumi.Context, args LookupActionOutputArgs, opts ...pulumi.InvokeOption) LookupActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActionResult, error) {
			args := v.(LookupActionArgs)
			r, err := LookupAction(ctx, &args, opts...)
			var s LookupActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActionResultOutput)
}

type LookupActionOutputArgs struct {
	// Action ID
	ActionId pulumi.StringInput `pulumi:"actionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionArgs)(nil)).Elem()
}

// Action for alert rule.
type LookupActionResultOutput struct{ *pulumi.OutputState }

func (LookupActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionResult)(nil)).Elem()
}

func (o LookupActionResultOutput) ToLookupActionResultOutput() LookupActionResultOutput {
	return o
}

func (o LookupActionResultOutput) ToLookupActionResultOutputWithContext(ctx context.Context) LookupActionResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupActionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupActionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
func (o LookupActionResultOutput) LogicAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.LogicAppResourceId }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupActionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupActionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupActionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupActionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The name of the logic app's workflow.
func (o LookupActionResultOutput) WorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActionResult) *string { return v.WorkflowId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionResultOutput{})
}
