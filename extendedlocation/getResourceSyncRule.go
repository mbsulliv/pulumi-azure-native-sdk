// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extendedlocation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the resourceSyncRule with a specified resource group, subscription id Custom Location resource name and Resource Sync Rule name.
// Azure REST API version: 2021-08-31-preview.
func LookupResourceSyncRule(ctx *pulumi.Context, args *LookupResourceSyncRuleArgs, opts ...pulumi.InvokeOption) (*LookupResourceSyncRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceSyncRuleResult
	err := ctx.Invoke("azure-native:extendedlocation:getResourceSyncRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceSyncRuleArgs struct {
	// Resource Sync Rule name.
	ChildResourceName string `pulumi:"childResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Custom Locations name.
	ResourceName string `pulumi:"resourceName"`
}

// Resource Sync Rules definition.
type LookupResourceSyncRuleResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Priority represents a priority of the Resource Sync Rule
	Priority *int `pulumi:"priority"`
	// Provisioning State for the Resource Sync Rule.
	ProvisioningState string `pulumi:"provisioningState"`
	// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
	Selector *ResourceSyncRulePropertiesResponseSelector `pulumi:"selector"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
	TargetResourceGroup *string `pulumi:"targetResourceGroup"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupResourceSyncRuleOutput(ctx *pulumi.Context, args LookupResourceSyncRuleOutputArgs, opts ...pulumi.InvokeOption) LookupResourceSyncRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceSyncRuleResult, error) {
			args := v.(LookupResourceSyncRuleArgs)
			r, err := LookupResourceSyncRule(ctx, &args, opts...)
			var s LookupResourceSyncRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceSyncRuleResultOutput)
}

type LookupResourceSyncRuleOutputArgs struct {
	// Resource Sync Rule name.
	ChildResourceName pulumi.StringInput `pulumi:"childResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Custom Locations name.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupResourceSyncRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceSyncRuleArgs)(nil)).Elem()
}

// Resource Sync Rules definition.
type LookupResourceSyncRuleResultOutput struct{ *pulumi.OutputState }

func (LookupResourceSyncRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceSyncRuleResult)(nil)).Elem()
}

func (o LookupResourceSyncRuleResultOutput) ToLookupResourceSyncRuleResultOutput() LookupResourceSyncRuleResultOutput {
	return o
}

func (o LookupResourceSyncRuleResultOutput) ToLookupResourceSyncRuleResultOutputWithContext(ctx context.Context) LookupResourceSyncRuleResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupResourceSyncRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupResourceSyncRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupResourceSyncRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Priority represents a priority of the Resource Sync Rule
func (o LookupResourceSyncRuleResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

// Provisioning State for the Resource Sync Rule.
func (o LookupResourceSyncRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
func (o LookupResourceSyncRuleResultOutput) Selector() ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) *ResourceSyncRulePropertiesResponseSelector { return v.Selector }).(ResourceSyncRulePropertiesResponseSelectorPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o LookupResourceSyncRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupResourceSyncRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
func (o LookupResourceSyncRuleResultOutput) TargetResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) *string { return v.TargetResourceGroup }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupResourceSyncRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceSyncRuleResultOutput{})
}
