// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an Internet Gateway Rule resource.
func LookupInternetGatewayRule(ctx *pulumi.Context, args *LookupInternetGatewayRuleArgs, opts ...pulumi.InvokeOption) (*LookupInternetGatewayRuleResult, error) {
	var rv LookupInternetGatewayRuleResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230615:getInternetGatewayRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInternetGatewayRuleArgs struct {
	// Name of the Internet Gateway rule.
	InternetGatewayRuleName string `pulumi:"internetGatewayRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Internet Gateway Rule resource definition.
type LookupInternetGatewayRuleResult struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// List of Internet Gateway resource Id.
	InternetGatewayIds []string `pulumi:"internetGatewayIds"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Rules for the InternetGateways
	RuleProperties RulePropertiesResponse `pulumi:"ruleProperties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupInternetGatewayRuleOutput(ctx *pulumi.Context, args LookupInternetGatewayRuleOutputArgs, opts ...pulumi.InvokeOption) LookupInternetGatewayRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInternetGatewayRuleResult, error) {
			args := v.(LookupInternetGatewayRuleArgs)
			r, err := LookupInternetGatewayRule(ctx, &args, opts...)
			var s LookupInternetGatewayRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInternetGatewayRuleResultOutput)
}

type LookupInternetGatewayRuleOutputArgs struct {
	// Name of the Internet Gateway rule.
	InternetGatewayRuleName pulumi.StringInput `pulumi:"internetGatewayRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInternetGatewayRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetGatewayRuleArgs)(nil)).Elem()
}

// The Internet Gateway Rule resource definition.
type LookupInternetGatewayRuleResultOutput struct{ *pulumi.OutputState }

func (LookupInternetGatewayRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetGatewayRuleResult)(nil)).Elem()
}

func (o LookupInternetGatewayRuleResultOutput) ToLookupInternetGatewayRuleResultOutput() LookupInternetGatewayRuleResultOutput {
	return o
}

func (o LookupInternetGatewayRuleResultOutput) ToLookupInternetGatewayRuleResultOutputWithContext(ctx context.Context) LookupInternetGatewayRuleResultOutput {
	return o
}

// Switch configuration description.
func (o LookupInternetGatewayRuleResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupInternetGatewayRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of Internet Gateway resource Id.
func (o LookupInternetGatewayRuleResultOutput) InternetGatewayIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) []string { return v.InternetGatewayIds }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o LookupInternetGatewayRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupInternetGatewayRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupInternetGatewayRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Rules for the InternetGateways
func (o LookupInternetGatewayRuleResultOutput) RuleProperties() RulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) RulePropertiesResponse { return v.RuleProperties }).(RulePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupInternetGatewayRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupInternetGatewayRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupInternetGatewayRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetGatewayRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternetGatewayRuleResultOutput{})
}
