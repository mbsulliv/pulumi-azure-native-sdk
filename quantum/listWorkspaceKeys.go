// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quantum

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the keys to use with the Quantum APIs. A key is used to authenticate and authorize access to the Quantum REST APIs. Only one key is needed at a time; two are given to provide seamless key regeneration.
// Azure REST API version: 2023-11-13-preview.
func ListWorkspaceKeys(ctx *pulumi.Context, args *ListWorkspaceKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspaceKeysResult
	err := ctx.Invoke("azure-native:quantum:listWorkspaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceKeysArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the quantum workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Result of list Api keys and connection strings.
type ListWorkspaceKeysResult struct {
	// Indicator of enablement of the Quantum workspace Api keys.
	ApiKeyEnabled *bool `pulumi:"apiKeyEnabled"`
	// The connection string of the primary api key.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The quantum workspace primary api key.
	PrimaryKey *ApiKeyResponse `pulumi:"primaryKey"`
	// The connection string of the secondary api key.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The quantum workspace secondary api key.
	SecondaryKey *ApiKeyResponse `pulumi:"secondaryKey"`
}

func ListWorkspaceKeysOutput(ctx *pulumi.Context, args ListWorkspaceKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceKeysResult, error) {
			args := v.(ListWorkspaceKeysArgs)
			r, err := ListWorkspaceKeys(ctx, &args, opts...)
			var s ListWorkspaceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceKeysResultOutput)
}

type ListWorkspaceKeysOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the quantum workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysArgs)(nil)).Elem()
}

// Result of list Api keys and connection strings.
type ListWorkspaceKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysResult)(nil)).Elem()
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutput() ListWorkspaceKeysResultOutput {
	return o
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutputWithContext(ctx context.Context) ListWorkspaceKeysResultOutput {
	return o
}

// Indicator of enablement of the Quantum workspace Api keys.
func (o ListWorkspaceKeysResultOutput) ApiKeyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) *bool { return v.ApiKeyEnabled }).(pulumi.BoolPtrOutput)
}

// The connection string of the primary api key.
func (o ListWorkspaceKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// The quantum workspace primary api key.
func (o ListWorkspaceKeysResultOutput) PrimaryKey() ApiKeyResponsePtrOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) *ApiKeyResponse { return v.PrimaryKey }).(ApiKeyResponsePtrOutput)
}

// The connection string of the secondary api key.
func (o ListWorkspaceKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// The quantum workspace secondary api key.
func (o ListWorkspaceKeysResultOutput) SecondaryKey() ApiKeyResponsePtrOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) *ApiKeyResponse { return v.SecondaryKey }).(ApiKeyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceKeysResultOutput{})
}
