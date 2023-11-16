// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of a storage account’s Table service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
func LookupTableServiceProperties(ctx *pulumi.Context, args *LookupTableServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupTableServicePropertiesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTableServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20230101:getTableServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableServicePropertiesArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Table Service within the specified storage account. Table Service Name must be 'default'
	TableServiceName string `pulumi:"tableServiceName"`
}

// The properties of a storage account’s Table service.
type LookupTableServicePropertiesResult struct {
	// Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Table service.
	Cors *CorsRulesResponse `pulumi:"cors"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTableServicePropertiesOutput(ctx *pulumi.Context, args LookupTableServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupTableServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableServicePropertiesResult, error) {
			args := v.(LookupTableServicePropertiesArgs)
			r, err := LookupTableServiceProperties(ctx, &args, opts...)
			var s LookupTableServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableServicePropertiesResultOutput)
}

type LookupTableServicePropertiesOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Table Service within the specified storage account. Table Service Name must be 'default'
	TableServiceName pulumi.StringInput `pulumi:"tableServiceName"`
}

func (LookupTableServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableServicePropertiesArgs)(nil)).Elem()
}

// The properties of a storage account’s Table service.
type LookupTableServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupTableServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableServicePropertiesResult)(nil)).Elem()
}

func (o LookupTableServicePropertiesResultOutput) ToLookupTableServicePropertiesResultOutput() LookupTableServicePropertiesResultOutput {
	return o
}

func (o LookupTableServicePropertiesResultOutput) ToLookupTableServicePropertiesResultOutputWithContext(ctx context.Context) LookupTableServicePropertiesResultOutput {
	return o
}

// Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Table service.
func (o LookupTableServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTableServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTableServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTableServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableServicePropertiesResultOutput{})
}
