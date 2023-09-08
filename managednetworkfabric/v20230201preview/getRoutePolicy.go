// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements Route Policy GET method.
func LookupRoutePolicy(ctx *pulumi.Context, args *LookupRoutePolicyArgs, opts ...pulumi.InvokeOption) (*LookupRoutePolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRoutePolicyResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230201preview:getRoutePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoutePolicyArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Route Policy
	RoutePolicyName string `pulumi:"routePolicyName"`
}

// The RoutePolicy resource definition.
type LookupRoutePolicyResult struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Route Policy statements.
	Statements []RoutePolicyStatementPropertiesResponse `pulumi:"statements"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupRoutePolicyOutput(ctx *pulumi.Context, args LookupRoutePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupRoutePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutePolicyResult, error) {
			args := v.(LookupRoutePolicyArgs)
			r, err := LookupRoutePolicy(ctx, &args, opts...)
			var s LookupRoutePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoutePolicyResultOutput)
}

type LookupRoutePolicyOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Route Policy
	RoutePolicyName pulumi.StringInput `pulumi:"routePolicyName"`
}

func (LookupRoutePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutePolicyArgs)(nil)).Elem()
}

// The RoutePolicy resource definition.
type LookupRoutePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupRoutePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutePolicyResult)(nil)).Elem()
}

func (o LookupRoutePolicyResultOutput) ToLookupRoutePolicyResultOutput() LookupRoutePolicyResultOutput {
	return o
}

func (o LookupRoutePolicyResultOutput) ToLookupRoutePolicyResultOutputWithContext(ctx context.Context) LookupRoutePolicyResultOutput {
	return o
}

// Switch configuration description.
func (o LookupRoutePolicyResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRoutePolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupRoutePolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRoutePolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o LookupRoutePolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Route Policy statements.
func (o LookupRoutePolicyResultOutput) Statements() RoutePolicyStatementPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) []RoutePolicyStatementPropertiesResponse { return v.Statements }).(RoutePolicyStatementPropertiesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRoutePolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupRoutePolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRoutePolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutePolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutePolicyResultOutput{})
}