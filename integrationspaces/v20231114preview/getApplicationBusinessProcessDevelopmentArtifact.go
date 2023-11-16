// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231114preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The get business process development artifact action.
func GetApplicationBusinessProcessDevelopmentArtifact(ctx *pulumi.Context, args *GetApplicationBusinessProcessDevelopmentArtifactArgs, opts ...pulumi.InvokeOption) (*GetApplicationBusinessProcessDevelopmentArtifactResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationBusinessProcessDevelopmentArtifactResult
	err := ctx.Invoke("azure-native:integrationspaces/v20231114preview:getApplicationBusinessProcessDevelopmentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApplicationBusinessProcessDevelopmentArtifactArgs struct {
	// The name of the Application
	ApplicationName string `pulumi:"applicationName"`
	// The name of the business process development artifact.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the space
	SpaceName string `pulumi:"spaceName"`
}

// The business process development artifact save or get response.
type GetApplicationBusinessProcessDevelopmentArtifactResult struct {
	// The name of the business process development artifact.
	Name string `pulumi:"name"`
	// The properties of the business process development artifact.
	Properties BusinessProcessDevelopmentArtifactPropertiesResponse `pulumi:"properties"`
	// The system data of the business process development artifact.
	SystemData SystemDataResponse `pulumi:"systemData"`
}

func GetApplicationBusinessProcessDevelopmentArtifactOutput(ctx *pulumi.Context, args GetApplicationBusinessProcessDevelopmentArtifactOutputArgs, opts ...pulumi.InvokeOption) GetApplicationBusinessProcessDevelopmentArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationBusinessProcessDevelopmentArtifactResult, error) {
			args := v.(GetApplicationBusinessProcessDevelopmentArtifactArgs)
			r, err := GetApplicationBusinessProcessDevelopmentArtifact(ctx, &args, opts...)
			var s GetApplicationBusinessProcessDevelopmentArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationBusinessProcessDevelopmentArtifactResultOutput)
}

type GetApplicationBusinessProcessDevelopmentArtifactOutputArgs struct {
	// The name of the Application
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	// The name of the business process development artifact.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the space
	SpaceName pulumi.StringInput `pulumi:"spaceName"`
}

func (GetApplicationBusinessProcessDevelopmentArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationBusinessProcessDevelopmentArtifactArgs)(nil)).Elem()
}

// The business process development artifact save or get response.
type GetApplicationBusinessProcessDevelopmentArtifactResultOutput struct{ *pulumi.OutputState }

func (GetApplicationBusinessProcessDevelopmentArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationBusinessProcessDevelopmentArtifactResult)(nil)).Elem()
}

func (o GetApplicationBusinessProcessDevelopmentArtifactResultOutput) ToGetApplicationBusinessProcessDevelopmentArtifactResultOutput() GetApplicationBusinessProcessDevelopmentArtifactResultOutput {
	return o
}

func (o GetApplicationBusinessProcessDevelopmentArtifactResultOutput) ToGetApplicationBusinessProcessDevelopmentArtifactResultOutputWithContext(ctx context.Context) GetApplicationBusinessProcessDevelopmentArtifactResultOutput {
	return o
}

// The name of the business process development artifact.
func (o GetApplicationBusinessProcessDevelopmentArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationBusinessProcessDevelopmentArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the business process development artifact.
func (o GetApplicationBusinessProcessDevelopmentArtifactResultOutput) Properties() BusinessProcessDevelopmentArtifactPropertiesResponseOutput {
	return o.ApplyT(func(v GetApplicationBusinessProcessDevelopmentArtifactResult) BusinessProcessDevelopmentArtifactPropertiesResponse {
		return v.Properties
	}).(BusinessProcessDevelopmentArtifactPropertiesResponseOutput)
}

// The system data of the business process development artifact.
func (o GetApplicationBusinessProcessDevelopmentArtifactResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetApplicationBusinessProcessDevelopmentArtifactResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationBusinessProcessDevelopmentArtifactResultOutput{})
}
