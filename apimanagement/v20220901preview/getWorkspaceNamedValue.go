// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the named value specified by its identifier.
func LookupWorkspaceNamedValue(ctx *pulumi.Context, args *LookupWorkspaceNamedValueArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceNamedValueResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceNamedValueResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getWorkspaceNamedValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceNamedValueArgs struct {
	// Identifier of the NamedValue.
	NamedValueId string `pulumi:"namedValueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// NamedValue details.
type LookupWorkspaceNamedValueResult struct {
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

func LookupWorkspaceNamedValueOutput(ctx *pulumi.Context, args LookupWorkspaceNamedValueOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceNamedValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceNamedValueResult, error) {
			args := v.(LookupWorkspaceNamedValueArgs)
			r, err := LookupWorkspaceNamedValue(ctx, &args, opts...)
			var s LookupWorkspaceNamedValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceNamedValueResultOutput)
}

type LookupWorkspaceNamedValueOutputArgs struct {
	// Identifier of the NamedValue.
	NamedValueId pulumi.StringInput `pulumi:"namedValueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceNamedValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceNamedValueArgs)(nil)).Elem()
}

// NamedValue details.
type LookupWorkspaceNamedValueResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceNamedValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceNamedValueResult)(nil)).Elem()
}

func (o LookupWorkspaceNamedValueResultOutput) ToLookupWorkspaceNamedValueResultOutput() LookupWorkspaceNamedValueResultOutput {
	return o
}

func (o LookupWorkspaceNamedValueResultOutput) ToLookupWorkspaceNamedValueResultOutputWithContext(ctx context.Context) LookupWorkspaceNamedValueResultOutput {
	return o
}

// Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.
func (o LookupWorkspaceNamedValueResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceNamedValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) string { return v.Id }).(pulumi.StringOutput)
}

// KeyVault location details of the namedValue.
func (o LookupWorkspaceNamedValueResultOutput) KeyVault() KeyVaultContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) *KeyVaultContractPropertiesResponse { return v.KeyVault }).(KeyVaultContractPropertiesResponsePtrOutput)
}

// The name of the resource
func (o LookupWorkspaceNamedValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) string { return v.Name }).(pulumi.StringOutput)
}

// Determines whether the value is a secret and should be encrypted or not. Default value is false.
func (o LookupWorkspaceNamedValueResultOutput) Secret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) *bool { return v.Secret }).(pulumi.BoolPtrOutput)
}

// Optional tags that when provided can be used to filter the NamedValue list.
func (o LookupWorkspaceNamedValueResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceNamedValueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) string { return v.Type }).(pulumi.StringOutput)
}

// Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o LookupWorkspaceNamedValueResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceNamedValueResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceNamedValueResultOutput{})
}