// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package integrationspaces

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The list business process development artifacts action.
// Azure REST API version: 2023-11-14-preview.
func ListApplicationBusinessProcessDevelopmentArtifacts(ctx *pulumi.Context, args *ListApplicationBusinessProcessDevelopmentArtifactsArgs, opts ...pulumi.InvokeOption) (*ListApplicationBusinessProcessDevelopmentArtifactsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListApplicationBusinessProcessDevelopmentArtifactsResult
	err := ctx.Invoke("azure-native:integrationspaces:listApplicationBusinessProcessDevelopmentArtifacts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplicationBusinessProcessDevelopmentArtifactsArgs struct {
	// The name of the Application
	ApplicationName string `pulumi:"applicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the space
	SpaceName string `pulumi:"spaceName"`
}

// The business process development artifact get collection response.
type ListApplicationBusinessProcessDevelopmentArtifactsResult struct {
	// The list of the business process development artifact.
	Value []SaveOrGetBusinessProcessDevelopmentArtifactResponseResponse `pulumi:"value"`
}

func ListApplicationBusinessProcessDevelopmentArtifactsOutput(ctx *pulumi.Context, args ListApplicationBusinessProcessDevelopmentArtifactsOutputArgs, opts ...pulumi.InvokeOption) ListApplicationBusinessProcessDevelopmentArtifactsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplicationBusinessProcessDevelopmentArtifactsResult, error) {
			args := v.(ListApplicationBusinessProcessDevelopmentArtifactsArgs)
			r, err := ListApplicationBusinessProcessDevelopmentArtifacts(ctx, &args, opts...)
			var s ListApplicationBusinessProcessDevelopmentArtifactsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplicationBusinessProcessDevelopmentArtifactsResultOutput)
}

type ListApplicationBusinessProcessDevelopmentArtifactsOutputArgs struct {
	// The name of the Application
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the space
	SpaceName pulumi.StringInput `pulumi:"spaceName"`
}

func (ListApplicationBusinessProcessDevelopmentArtifactsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplicationBusinessProcessDevelopmentArtifactsArgs)(nil)).Elem()
}

// The business process development artifact get collection response.
type ListApplicationBusinessProcessDevelopmentArtifactsResultOutput struct{ *pulumi.OutputState }

func (ListApplicationBusinessProcessDevelopmentArtifactsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplicationBusinessProcessDevelopmentArtifactsResult)(nil)).Elem()
}

func (o ListApplicationBusinessProcessDevelopmentArtifactsResultOutput) ToListApplicationBusinessProcessDevelopmentArtifactsResultOutput() ListApplicationBusinessProcessDevelopmentArtifactsResultOutput {
	return o
}

func (o ListApplicationBusinessProcessDevelopmentArtifactsResultOutput) ToListApplicationBusinessProcessDevelopmentArtifactsResultOutputWithContext(ctx context.Context) ListApplicationBusinessProcessDevelopmentArtifactsResultOutput {
	return o
}

func (o ListApplicationBusinessProcessDevelopmentArtifactsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListApplicationBusinessProcessDevelopmentArtifactsResult] {
	return pulumix.Output[ListApplicationBusinessProcessDevelopmentArtifactsResult]{
		OutputState: o.OutputState,
	}
}

// The list of the business process development artifact.
func (o ListApplicationBusinessProcessDevelopmentArtifactsResultOutput) Value() SaveOrGetBusinessProcessDevelopmentArtifactResponseResponseArrayOutput {
	return o.ApplyT(func(v ListApplicationBusinessProcessDevelopmentArtifactsResult) []SaveOrGetBusinessProcessDevelopmentArtifactResponseResponse {
		return v.Value
	}).(SaveOrGetBusinessProcessDevelopmentArtifactResponseResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplicationBusinessProcessDevelopmentArtifactsResultOutput{})
}