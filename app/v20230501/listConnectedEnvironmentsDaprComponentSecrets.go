// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Dapr component Secrets Collection for ListSecrets Action.
func ListConnectedEnvironmentsDaprComponentSecrets(ctx *pulumi.Context, args *ListConnectedEnvironmentsDaprComponentSecretsArgs, opts ...pulumi.InvokeOption) (*ListConnectedEnvironmentsDaprComponentSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListConnectedEnvironmentsDaprComponentSecretsResult
	err := ctx.Invoke("azure-native:app/v20230501:listConnectedEnvironmentsDaprComponentSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedEnvironmentsDaprComponentSecretsArgs struct {
	// Name of the Dapr Component.
	ComponentName string `pulumi:"componentName"`
	// Name of the connected environment.
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Dapr component Secrets Collection for ListSecrets Action.
type ListConnectedEnvironmentsDaprComponentSecretsResult struct {
	// Collection of secrets used by a Dapr component
	Value []DaprSecretResponse `pulumi:"value"`
}

func ListConnectedEnvironmentsDaprComponentSecretsOutput(ctx *pulumi.Context, args ListConnectedEnvironmentsDaprComponentSecretsOutputArgs, opts ...pulumi.InvokeOption) ListConnectedEnvironmentsDaprComponentSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectedEnvironmentsDaprComponentSecretsResult, error) {
			args := v.(ListConnectedEnvironmentsDaprComponentSecretsArgs)
			r, err := ListConnectedEnvironmentsDaprComponentSecrets(ctx, &args, opts...)
			var s ListConnectedEnvironmentsDaprComponentSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectedEnvironmentsDaprComponentSecretsResultOutput)
}

type ListConnectedEnvironmentsDaprComponentSecretsOutputArgs struct {
	// Name of the Dapr Component.
	ComponentName pulumi.StringInput `pulumi:"componentName"`
	// Name of the connected environment.
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListConnectedEnvironmentsDaprComponentSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedEnvironmentsDaprComponentSecretsArgs)(nil)).Elem()
}

// Dapr component Secrets Collection for ListSecrets Action.
type ListConnectedEnvironmentsDaprComponentSecretsResultOutput struct{ *pulumi.OutputState }

func (ListConnectedEnvironmentsDaprComponentSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedEnvironmentsDaprComponentSecretsResult)(nil)).Elem()
}

func (o ListConnectedEnvironmentsDaprComponentSecretsResultOutput) ToListConnectedEnvironmentsDaprComponentSecretsResultOutput() ListConnectedEnvironmentsDaprComponentSecretsResultOutput {
	return o
}

func (o ListConnectedEnvironmentsDaprComponentSecretsResultOutput) ToListConnectedEnvironmentsDaprComponentSecretsResultOutputWithContext(ctx context.Context) ListConnectedEnvironmentsDaprComponentSecretsResultOutput {
	return o
}

func (o ListConnectedEnvironmentsDaprComponentSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListConnectedEnvironmentsDaprComponentSecretsResult] {
	return pulumix.Output[ListConnectedEnvironmentsDaprComponentSecretsResult]{
		OutputState: o.OutputState,
	}
}

// Collection of secrets used by a Dapr component
func (o ListConnectedEnvironmentsDaprComponentSecretsResultOutput) Value() DaprSecretResponseArrayOutput {
	return o.ApplyT(func(v ListConnectedEnvironmentsDaprComponentSecretsResult) []DaprSecretResponse { return v.Value }).(DaprSecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectedEnvironmentsDaprComponentSecretsResultOutput{})
}
