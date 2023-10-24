// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the config server and its properties.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview.
func LookupConfigServer(ctx *pulumi.Context, args *LookupConfigServerArgs, opts ...pulumi.InvokeOption) (*LookupConfigServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigServerResult
	err := ctx.Invoke("azure-native:appplatform:getConfigServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigServerArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Config Server resource
type LookupConfigServerResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the Config Server resource
	Properties ConfigServerPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupConfigServerOutput(ctx *pulumi.Context, args LookupConfigServerOutputArgs, opts ...pulumi.InvokeOption) LookupConfigServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigServerResult, error) {
			args := v.(LookupConfigServerArgs)
			r, err := LookupConfigServer(ctx, &args, opts...)
			var s LookupConfigServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigServerResultOutput)
}

type LookupConfigServerOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupConfigServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigServerArgs)(nil)).Elem()
}

// Config Server resource
type LookupConfigServerResultOutput struct{ *pulumi.OutputState }

func (LookupConfigServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigServerResult)(nil)).Elem()
}

func (o LookupConfigServerResultOutput) ToLookupConfigServerResultOutput() LookupConfigServerResultOutput {
	return o
}

func (o LookupConfigServerResultOutput) ToLookupConfigServerResultOutputWithContext(ctx context.Context) LookupConfigServerResultOutput {
	return o
}

func (o LookupConfigServerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConfigServerResult] {
	return pulumix.Output[LookupConfigServerResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupConfigServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupConfigServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Config Server resource
func (o LookupConfigServerResultOutput) Properties() ConfigServerPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigServerResult) ConfigServerPropertiesResponse { return v.Properties }).(ConfigServerPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupConfigServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupConfigServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigServerResultOutput{})
}
