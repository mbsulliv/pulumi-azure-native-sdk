// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements IP Extended Community GET method.
func LookupIpExtendedCommunity(ctx *pulumi.Context, args *LookupIpExtendedCommunityArgs, opts ...pulumi.InvokeOption) (*LookupIpExtendedCommunityResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIpExtendedCommunityResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230615:getIpExtendedCommunity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpExtendedCommunityArgs struct {
	// Name of the IP Extended Community.
	IpExtendedCommunityName string `pulumi:"ipExtendedCommunityName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The IP Extended Community resource definition.
type LookupIpExtendedCommunityResult struct {
	// Administrative state of the resource.
	AdministrativeState string `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState string `pulumi:"configurationState"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// List of IP Extended Community Rules.
	IpExtendedCommunityRules []IpExtendedCommunityRuleResponse `pulumi:"ipExtendedCommunityRules"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIpExtendedCommunityOutput(ctx *pulumi.Context, args LookupIpExtendedCommunityOutputArgs, opts ...pulumi.InvokeOption) LookupIpExtendedCommunityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpExtendedCommunityResult, error) {
			args := v.(LookupIpExtendedCommunityArgs)
			r, err := LookupIpExtendedCommunity(ctx, &args, opts...)
			var s LookupIpExtendedCommunityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpExtendedCommunityResultOutput)
}

type LookupIpExtendedCommunityOutputArgs struct {
	// Name of the IP Extended Community.
	IpExtendedCommunityName pulumi.StringInput `pulumi:"ipExtendedCommunityName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIpExtendedCommunityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpExtendedCommunityArgs)(nil)).Elem()
}

// The IP Extended Community resource definition.
type LookupIpExtendedCommunityResultOutput struct{ *pulumi.OutputState }

func (LookupIpExtendedCommunityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpExtendedCommunityResult)(nil)).Elem()
}

func (o LookupIpExtendedCommunityResultOutput) ToLookupIpExtendedCommunityResultOutput() LookupIpExtendedCommunityResultOutput {
	return o
}

func (o LookupIpExtendedCommunityResultOutput) ToLookupIpExtendedCommunityResultOutputWithContext(ctx context.Context) LookupIpExtendedCommunityResultOutput {
	return o
}

// Administrative state of the resource.
func (o LookupIpExtendedCommunityResultOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) string { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupIpExtendedCommunityResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o LookupIpExtendedCommunityResultOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) string { return v.ConfigurationState }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupIpExtendedCommunityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of IP Extended Community Rules.
func (o LookupIpExtendedCommunityResultOutput) IpExtendedCommunityRules() IpExtendedCommunityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) []IpExtendedCommunityRuleResponse {
		return v.IpExtendedCommunityRules
	}).(IpExtendedCommunityRuleResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupIpExtendedCommunityResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIpExtendedCommunityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupIpExtendedCommunityResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIpExtendedCommunityResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupIpExtendedCommunityResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIpExtendedCommunityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpExtendedCommunityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpExtendedCommunityResultOutput{})
}