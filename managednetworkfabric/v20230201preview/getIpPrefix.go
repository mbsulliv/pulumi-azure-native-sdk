// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements IP Prefix GET method.
func LookupIpPrefix(ctx *pulumi.Context, args *LookupIpPrefixArgs, opts ...pulumi.InvokeOption) (*LookupIpPrefixResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIpPrefixResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230201preview:getIpPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpPrefixArgs struct {
	// Name of the IP Prefix
	IpPrefixName string `pulumi:"ipPrefixName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The IPPrefix resource definition.
type LookupIpPrefixResult struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// IpPrefix contains the list of IP PrefixRules objects.
	IpPrefixRules []IpPrefixPropertiesResponseIpPrefixRules `pulumi:"ipPrefixRules"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIpPrefixOutput(ctx *pulumi.Context, args LookupIpPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupIpPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpPrefixResult, error) {
			args := v.(LookupIpPrefixArgs)
			r, err := LookupIpPrefix(ctx, &args, opts...)
			var s LookupIpPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpPrefixResultOutput)
}

type LookupIpPrefixOutputArgs struct {
	// Name of the IP Prefix
	IpPrefixName pulumi.StringInput `pulumi:"ipPrefixName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIpPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpPrefixArgs)(nil)).Elem()
}

// The IPPrefix resource definition.
type LookupIpPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupIpPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpPrefixResult)(nil)).Elem()
}

func (o LookupIpPrefixResultOutput) ToLookupIpPrefixResultOutput() LookupIpPrefixResultOutput {
	return o
}

func (o LookupIpPrefixResultOutput) ToLookupIpPrefixResultOutputWithContext(ctx context.Context) LookupIpPrefixResultOutput {
	return o
}

// Switch configuration description.
func (o LookupIpPrefixResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupIpPrefixResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) string { return v.Id }).(pulumi.StringOutput)
}

// IpPrefix contains the list of IP PrefixRules objects.
func (o LookupIpPrefixResultOutput) IpPrefixRules() IpPrefixPropertiesResponseIpPrefixRulesArrayOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) []IpPrefixPropertiesResponseIpPrefixRules { return v.IpPrefixRules }).(IpPrefixPropertiesResponseIpPrefixRulesArrayOutput)
}

// The geo-location where the resource lives
func (o LookupIpPrefixResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIpPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o LookupIpPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIpPrefixResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupIpPrefixResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIpPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpPrefixResultOutput{})
}
