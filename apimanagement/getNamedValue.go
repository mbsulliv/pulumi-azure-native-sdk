// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the named value specified by its identifier.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
func LookupNamedValue(ctx *pulumi.Context, args *LookupNamedValueArgs, opts ...pulumi.InvokeOption) (*LookupNamedValueResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamedValueResult
	err := ctx.Invoke("azure-native:apimanagement:getNamedValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamedValueArgs struct {
	// Identifier of the NamedValue.
	NamedValueId string `pulumi:"namedValueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// NamedValue details.
type LookupNamedValueResult struct {
	// Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// KeyVault location details of the namedValue.
	KeyVault *KeyVaultContractPropertiesResponse `pulumi:"keyVault"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Determines whether the value is a secret and should be encrypted or not. Default value is false.
	Secret *bool `pulumi:"secret"`
	// Optional tags that when provided can be used to filter the NamedValue list.
	Tags []string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	Value *string `pulumi:"value"`
}

func LookupNamedValueOutput(ctx *pulumi.Context, args LookupNamedValueOutputArgs, opts ...pulumi.InvokeOption) LookupNamedValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamedValueResult, error) {
			args := v.(LookupNamedValueArgs)
			r, err := LookupNamedValue(ctx, &args, opts...)
			var s LookupNamedValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamedValueResultOutput)
}

type LookupNamedValueOutputArgs struct {
	// Identifier of the NamedValue.
	NamedValueId pulumi.StringInput `pulumi:"namedValueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupNamedValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamedValueArgs)(nil)).Elem()
}

// NamedValue details.
type LookupNamedValueResultOutput struct{ *pulumi.OutputState }

func (LookupNamedValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamedValueResult)(nil)).Elem()
}

func (o LookupNamedValueResultOutput) ToLookupNamedValueResultOutput() LookupNamedValueResultOutput {
	return o
}

func (o LookupNamedValueResultOutput) ToLookupNamedValueResultOutputWithContext(ctx context.Context) LookupNamedValueResultOutput {
	return o
}

// Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.
func (o LookupNamedValueResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNamedValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.Id }).(pulumi.StringOutput)
}

// KeyVault location details of the namedValue.
func (o LookupNamedValueResultOutput) KeyVault() KeyVaultContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupNamedValueResult) *KeyVaultContractPropertiesResponse { return v.KeyVault }).(KeyVaultContractPropertiesResponsePtrOutput)
}

// The name of the resource
func (o LookupNamedValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.Name }).(pulumi.StringOutput)
}

// Determines whether the value is a secret and should be encrypted or not. Default value is false.
func (o LookupNamedValueResultOutput) Secret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamedValueResult) *bool { return v.Secret }).(pulumi.BoolPtrOutput)
}

// Optional tags that when provided can be used to filter the NamedValue list.
func (o LookupNamedValueResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNamedValueResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNamedValueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.Type }).(pulumi.StringOutput)
}

// Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o LookupNamedValueResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamedValueResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamedValueResultOutput{})
}
