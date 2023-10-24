// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Base definition for datastore secrets.
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01.
func ListDatastoreSecrets(ctx *pulumi.Context, args *ListDatastoreSecretsArgs, opts ...pulumi.InvokeOption) (*ListDatastoreSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDatastoreSecretsResult
	err := ctx.Invoke("azure-native:machinelearningservices:listDatastoreSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatastoreSecretsArgs struct {
	// Datastore name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Base definition for datastore secrets.
type ListDatastoreSecretsResult struct {
	// [Required] Credential type used to authentication with storage.
	SecretsType string `pulumi:"secretsType"`
}

func ListDatastoreSecretsOutput(ctx *pulumi.Context, args ListDatastoreSecretsOutputArgs, opts ...pulumi.InvokeOption) ListDatastoreSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatastoreSecretsResult, error) {
			args := v.(ListDatastoreSecretsArgs)
			r, err := ListDatastoreSecrets(ctx, &args, opts...)
			var s ListDatastoreSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatastoreSecretsResultOutput)
}

type ListDatastoreSecretsOutputArgs struct {
	// Datastore name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListDatastoreSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatastoreSecretsArgs)(nil)).Elem()
}

// Base definition for datastore secrets.
type ListDatastoreSecretsResultOutput struct{ *pulumi.OutputState }

func (ListDatastoreSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatastoreSecretsResult)(nil)).Elem()
}

func (o ListDatastoreSecretsResultOutput) ToListDatastoreSecretsResultOutput() ListDatastoreSecretsResultOutput {
	return o
}

func (o ListDatastoreSecretsResultOutput) ToListDatastoreSecretsResultOutputWithContext(ctx context.Context) ListDatastoreSecretsResultOutput {
	return o
}

func (o ListDatastoreSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListDatastoreSecretsResult] {
	return pulumix.Output[ListDatastoreSecretsResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Credential type used to authentication with storage.
func (o ListDatastoreSecretsResultOutput) SecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatastoreSecretsResult) string { return v.SecretsType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatastoreSecretsResultOutput{})
}
