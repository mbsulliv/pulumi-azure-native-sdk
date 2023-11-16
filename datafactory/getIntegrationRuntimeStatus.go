// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets detailed status information for an integration runtime.
// Azure REST API version: 2018-06-01.
func GetIntegrationRuntimeStatus(ctx *pulumi.Context, args *GetIntegrationRuntimeStatusArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeStatusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetIntegrationRuntimeStatusResult
	err := ctx.Invoke("azure-native:datafactory:getIntegrationRuntimeStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeStatusArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Integration runtime status response.
type GetIntegrationRuntimeStatusResult struct {
	// The integration runtime name.
	Name string `pulumi:"name"`
	// Integration runtime properties.
	Properties interface{} `pulumi:"properties"`
}

func GetIntegrationRuntimeStatusOutput(ctx *pulumi.Context, args GetIntegrationRuntimeStatusOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIntegrationRuntimeStatusResult, error) {
			args := v.(GetIntegrationRuntimeStatusArgs)
			r, err := GetIntegrationRuntimeStatus(ctx, &args, opts...)
			var s GetIntegrationRuntimeStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIntegrationRuntimeStatusResultOutput)
}

type GetIntegrationRuntimeStatusOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetIntegrationRuntimeStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeStatusArgs)(nil)).Elem()
}

// Integration runtime status response.
type GetIntegrationRuntimeStatusResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeStatusResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeStatusResultOutput) ToGetIntegrationRuntimeStatusResultOutput() GetIntegrationRuntimeStatusResultOutput {
	return o
}

func (o GetIntegrationRuntimeStatusResultOutput) ToGetIntegrationRuntimeStatusResultOutputWithContext(ctx context.Context) GetIntegrationRuntimeStatusResultOutput {
	return o
}

// The integration runtime name.
func (o GetIntegrationRuntimeStatusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeStatusResult) string { return v.Name }).(pulumi.StringOutput)
}

// Integration runtime properties.
func (o GetIntegrationRuntimeStatusResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeStatusResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIntegrationRuntimeStatusResultOutput{})
}
