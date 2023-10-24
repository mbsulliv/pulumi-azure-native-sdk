// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get object metadata from an integration runtime
// Azure REST API version: 2021-06-01.
//
// Other available API versions: 2021-06-01-preview.
func GetIntegrationRuntimeObjectMetadatum(ctx *pulumi.Context, args *GetIntegrationRuntimeObjectMetadatumArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetIntegrationRuntimeObjectMetadatumResult
	err := ctx.Invoke("azure-native:synapse:getIntegrationRuntimeObjectMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeObjectMetadatumArgs struct {
	// Integration runtime name
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath *string `pulumi:"metadataPath"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// A list of SSIS object metadata.
type GetIntegrationRuntimeObjectMetadatumResult struct {
	// The link to the next page of results, if any remaining results exist.
	NextLink *string `pulumi:"nextLink"`
	// List of SSIS object metadata.
	Value []interface{} `pulumi:"value"`
}

func GetIntegrationRuntimeObjectMetadatumOutput(ctx *pulumi.Context, args GetIntegrationRuntimeObjectMetadatumOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeObjectMetadatumResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIntegrationRuntimeObjectMetadatumResult, error) {
			args := v.(GetIntegrationRuntimeObjectMetadatumArgs)
			r, err := GetIntegrationRuntimeObjectMetadatum(ctx, &args, opts...)
			var s GetIntegrationRuntimeObjectMetadatumResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIntegrationRuntimeObjectMetadatumResultOutput)
}

type GetIntegrationRuntimeObjectMetadatumOutputArgs struct {
	// Integration runtime name
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath pulumi.StringPtrInput `pulumi:"metadataPath"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetIntegrationRuntimeObjectMetadatumOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumArgs)(nil)).Elem()
}

// A list of SSIS object metadata.
type GetIntegrationRuntimeObjectMetadatumResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeObjectMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToGetIntegrationRuntimeObjectMetadatumResultOutput() GetIntegrationRuntimeObjectMetadatumResultOutput {
	return o
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToGetIntegrationRuntimeObjectMetadatumResultOutputWithContext(ctx context.Context) GetIntegrationRuntimeObjectMetadatumResultOutput {
	return o
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetIntegrationRuntimeObjectMetadatumResult] {
	return pulumix.Output[GetIntegrationRuntimeObjectMetadatumResult]{
		OutputState: o.OutputState,
	}
}

// The link to the next page of results, if any remaining results exist.
func (o GetIntegrationRuntimeObjectMetadatumResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeObjectMetadatumResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of SSIS object metadata.
func (o GetIntegrationRuntimeObjectMetadatumResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeObjectMetadatumResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIntegrationRuntimeObjectMetadatumResultOutput{})
}
