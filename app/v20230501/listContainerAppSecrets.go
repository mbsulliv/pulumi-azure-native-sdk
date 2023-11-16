// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App Secrets Collection ARM resource.
func ListContainerAppSecrets(ctx *pulumi.Context, args *ListContainerAppSecretsArgs, opts ...pulumi.InvokeOption) (*ListContainerAppSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListContainerAppSecretsResult
	err := ctx.Invoke("azure-native:app/v20230501:listContainerAppSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListContainerAppSecretsArgs struct {
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Container App Secrets Collection ARM resource.
type ListContainerAppSecretsResult struct {
	// Collection of resources.
	Value []ContainerAppSecretResponse `pulumi:"value"`
}

func ListContainerAppSecretsOutput(ctx *pulumi.Context, args ListContainerAppSecretsOutputArgs, opts ...pulumi.InvokeOption) ListContainerAppSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListContainerAppSecretsResult, error) {
			args := v.(ListContainerAppSecretsArgs)
			r, err := ListContainerAppSecrets(ctx, &args, opts...)
			var s ListContainerAppSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListContainerAppSecretsResultOutput)
}

type ListContainerAppSecretsOutputArgs struct {
	// Name of the Container App.
	ContainerAppName pulumi.StringInput `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListContainerAppSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppSecretsArgs)(nil)).Elem()
}

// Container App Secrets Collection ARM resource.
type ListContainerAppSecretsResultOutput struct{ *pulumi.OutputState }

func (ListContainerAppSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppSecretsResult)(nil)).Elem()
}

func (o ListContainerAppSecretsResultOutput) ToListContainerAppSecretsResultOutput() ListContainerAppSecretsResultOutput {
	return o
}

func (o ListContainerAppSecretsResultOutput) ToListContainerAppSecretsResultOutputWithContext(ctx context.Context) ListContainerAppSecretsResultOutput {
	return o
}

// Collection of resources.
func (o ListContainerAppSecretsResultOutput) Value() ContainerAppSecretResponseArrayOutput {
	return o.ApplyT(func(v ListContainerAppSecretsResult) []ContainerAppSecretResponse { return v.Value }).(ContainerAppSecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListContainerAppSecretsResultOutput{})
}
