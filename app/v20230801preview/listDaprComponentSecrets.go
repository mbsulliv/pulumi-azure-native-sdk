// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Dapr component Secrets Collection for ListSecrets Action.
func ListDaprComponentSecrets(ctx *pulumi.Context, args *ListDaprComponentSecretsArgs, opts ...pulumi.InvokeOption) (*ListDaprComponentSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDaprComponentSecretsResult
	err := ctx.Invoke("azure-native:app/v20230801preview:listDaprComponentSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDaprComponentSecretsArgs struct {
	// Name of the Dapr Component.
	ComponentName string `pulumi:"componentName"`
	// Name of the Managed Environment.
	EnvironmentName string `pulumi:"environmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Dapr component Secrets Collection for ListSecrets Action.
type ListDaprComponentSecretsResult struct {
	// Collection of secrets used by a Dapr component
	Value []DaprSecretResponse `pulumi:"value"`
}

func ListDaprComponentSecretsOutput(ctx *pulumi.Context, args ListDaprComponentSecretsOutputArgs, opts ...pulumi.InvokeOption) ListDaprComponentSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDaprComponentSecretsResult, error) {
			args := v.(ListDaprComponentSecretsArgs)
			r, err := ListDaprComponentSecrets(ctx, &args, opts...)
			var s ListDaprComponentSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDaprComponentSecretsResultOutput)
}

type ListDaprComponentSecretsOutputArgs struct {
	// Name of the Dapr Component.
	ComponentName pulumi.StringInput `pulumi:"componentName"`
	// Name of the Managed Environment.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDaprComponentSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDaprComponentSecretsArgs)(nil)).Elem()
}

// Dapr component Secrets Collection for ListSecrets Action.
type ListDaprComponentSecretsResultOutput struct{ *pulumi.OutputState }

func (ListDaprComponentSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDaprComponentSecretsResult)(nil)).Elem()
}

func (o ListDaprComponentSecretsResultOutput) ToListDaprComponentSecretsResultOutput() ListDaprComponentSecretsResultOutput {
	return o
}

func (o ListDaprComponentSecretsResultOutput) ToListDaprComponentSecretsResultOutputWithContext(ctx context.Context) ListDaprComponentSecretsResultOutput {
	return o
}

// Collection of secrets used by a Dapr component
func (o ListDaprComponentSecretsResultOutput) Value() DaprSecretResponseArrayOutput {
	return o.ApplyT(func(v ListDaprComponentSecretsResult) []DaprSecretResponse { return v.Value }).(DaprSecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDaprComponentSecretsResultOutput{})
}
