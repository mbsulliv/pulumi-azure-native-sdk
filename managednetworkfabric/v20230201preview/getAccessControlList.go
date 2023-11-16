// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements Access Control List GET method.
func LookupAccessControlList(ctx *pulumi.Context, args *LookupAccessControlListArgs, opts ...pulumi.InvokeOption) (*LookupAccessControlListResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessControlListResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230201preview:getAccessControlList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessControlListArgs struct {
	// Name of the Access Control List
	AccessControlListName string `pulumi:"accessControlListName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The AccessControlList resource definition.
type LookupAccessControlListResult struct {
	// IP address family. Example: ipv4 | ipv6.
	AddressFamily string `pulumi:"addressFamily"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Access Control List conditions.
	Conditions []AccessControlListConditionPropertiesResponse `pulumi:"conditions"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
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

func LookupAccessControlListOutput(ctx *pulumi.Context, args LookupAccessControlListOutputArgs, opts ...pulumi.InvokeOption) LookupAccessControlListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessControlListResult, error) {
			args := v.(LookupAccessControlListArgs)
			r, err := LookupAccessControlList(ctx, &args, opts...)
			var s LookupAccessControlListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessControlListResultOutput)
}

type LookupAccessControlListOutputArgs struct {
	// Name of the Access Control List
	AccessControlListName pulumi.StringInput `pulumi:"accessControlListName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessControlListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessControlListArgs)(nil)).Elem()
}

// The AccessControlList resource definition.
type LookupAccessControlListResultOutput struct{ *pulumi.OutputState }

func (LookupAccessControlListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessControlListResult)(nil)).Elem()
}

func (o LookupAccessControlListResultOutput) ToLookupAccessControlListResultOutput() LookupAccessControlListResultOutput {
	return o
}

func (o LookupAccessControlListResultOutput) ToLookupAccessControlListResultOutputWithContext(ctx context.Context) LookupAccessControlListResultOutput {
	return o
}

// IP address family. Example: ipv4 | ipv6.
func (o LookupAccessControlListResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupAccessControlListResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Access Control List conditions.
func (o LookupAccessControlListResultOutput) Conditions() AccessControlListConditionPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) []AccessControlListConditionPropertiesResponse {
		return v.Conditions
	}).(AccessControlListConditionPropertiesResponseArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAccessControlListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupAccessControlListResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAccessControlListResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o LookupAccessControlListResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAccessControlListResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAccessControlListResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAccessControlListResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlListResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessControlListResultOutput{})
}
