// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response for POST requests that return single SharedAccessAuthorizationRule.
func LookupNotificationHubAuthorizationRule(ctx *pulumi.Context, args *LookupNotificationHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationHubAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20230101preview:getNotificationHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubAuthorizationRuleArgs struct {
	// Authorization Rule Name
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Notification Hub name
	NotificationHubName string `pulumi:"notificationHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for POST requests that return single SharedAccessAuthorizationRule.
type LookupNotificationHubAuthorizationRuleResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Deprecated - only for compatibility.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// SharedAccessAuthorizationRule properties.
	Properties SharedAccessAuthorizationRulePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Deprecated - only for compatibility.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNotificationHubAuthorizationRuleOutput(ctx *pulumi.Context, args LookupNotificationHubAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationHubAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationHubAuthorizationRuleResult, error) {
			args := v.(LookupNotificationHubAuthorizationRuleArgs)
			r, err := LookupNotificationHubAuthorizationRule(ctx, &args, opts...)
			var s LookupNotificationHubAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationHubAuthorizationRuleResultOutput)
}

type LookupNotificationHubAuthorizationRuleOutputArgs struct {
	// Authorization Rule Name
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Notification Hub name
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationHubAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubAuthorizationRuleArgs)(nil)).Elem()
}

// Response for POST requests that return single SharedAccessAuthorizationRule.
type LookupNotificationHubAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationHubAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) ToLookupNotificationHubAuthorizationRuleResultOutput() LookupNotificationHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) ToLookupNotificationHubAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupNotificationHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNotificationHubAuthorizationRuleResult] {
	return pulumix.Output[LookupNotificationHubAuthorizationRuleResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNotificationHubAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Deprecated - only for compatibility.
func (o LookupNotificationHubAuthorizationRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupNotificationHubAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// SharedAccessAuthorizationRule properties.
func (o LookupNotificationHubAuthorizationRuleResultOutput) Properties() SharedAccessAuthorizationRulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) SharedAccessAuthorizationRulePropertiesResponse {
		return v.Properties
	}).(SharedAccessAuthorizationRulePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNotificationHubAuthorizationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Deprecated - only for compatibility.
func (o LookupNotificationHubAuthorizationRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNotificationHubAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationHubAuthorizationRuleResultOutput{})
}
