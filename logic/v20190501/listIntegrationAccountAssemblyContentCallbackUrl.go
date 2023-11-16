// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the content callback url for an integration account assembly.
func ListIntegrationAccountAssemblyContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountAssemblyContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountAssemblyContentCallbackUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListIntegrationAccountAssemblyContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20190501:listIntegrationAccountAssemblyContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountAssemblyContentCallbackUrlArgs struct {
	// The assembly artifact name.
	AssemblyArtifactName string `pulumi:"assemblyArtifactName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The workflow trigger callback URL.
type ListIntegrationAccountAssemblyContentCallbackUrlResult struct {
	// Gets the workflow trigger callback URL base path.
	BasePath string `pulumi:"basePath"`
	// Gets the workflow trigger callback URL HTTP method.
	Method string `pulumi:"method"`
	// Gets the workflow trigger callback URL query parameters.
	Queries *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	// Gets the workflow trigger callback URL relative path.
	RelativePath string `pulumi:"relativePath"`
	// Gets the workflow trigger callback URL relative path parameters.
	RelativePathParameters []string `pulumi:"relativePathParameters"`
	// Gets the workflow trigger callback URL.
	Value string `pulumi:"value"`
}

func ListIntegrationAccountAssemblyContentCallbackUrlOutput(ctx *pulumi.Context, args ListIntegrationAccountAssemblyContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountAssemblyContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountAssemblyContentCallbackUrlResult, error) {
			args := v.(ListIntegrationAccountAssemblyContentCallbackUrlArgs)
			r, err := ListIntegrationAccountAssemblyContentCallbackUrl(ctx, &args, opts...)
			var s ListIntegrationAccountAssemblyContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountAssemblyContentCallbackUrlResultOutput)
}

type ListIntegrationAccountAssemblyContentCallbackUrlOutputArgs struct {
	// The assembly artifact name.
	AssemblyArtifactName pulumi.StringInput `pulumi:"assemblyArtifactName"`
	// The integration account name.
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListIntegrationAccountAssemblyContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountAssemblyContentCallbackUrlArgs)(nil)).Elem()
}

// The workflow trigger callback URL.
type ListIntegrationAccountAssemblyContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountAssemblyContentCallbackUrlResult)(nil)).Elem()
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) ToListIntegrationAccountAssemblyContentCallbackUrlResultOutput() ListIntegrationAccountAssemblyContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) ToListIntegrationAccountAssemblyContentCallbackUrlResultOutputWithContext(ctx context.Context) ListIntegrationAccountAssemblyContentCallbackUrlResultOutput {
	return o
}

// Gets the workflow trigger callback URL base path.
func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

// Gets the workflow trigger callback URL HTTP method.
func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

// Gets the workflow trigger callback URL query parameters.
func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

// Gets the workflow trigger callback URL relative path.
func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

// Gets the workflow trigger callback URL relative path parameters.
func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) []string {
		return v.RelativePathParameters
	}).(pulumi.StringArrayOutput)
}

// Gets the workflow trigger callback URL.
func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountAssemblyContentCallbackUrlResultOutput{})
}
