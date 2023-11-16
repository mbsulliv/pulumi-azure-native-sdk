// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the Application Configuration Service and its properties.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview.
func LookupConfigurationService(ctx *pulumi.Context, args *LookupConfigurationServiceArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigurationServiceResult
	err := ctx.Invoke("azure-native:appplatform:getConfigurationService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConfigurationServiceArgs struct {
	// The name of Application Configuration Service.
	ConfigurationServiceName string `pulumi:"configurationServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Application Configuration Service resource
type LookupConfigurationServiceResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Application Configuration Service properties payload
	Properties ConfigurationServicePropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupConfigurationServiceResult
func (val *LookupConfigurationServiceResult) Defaults() *LookupConfigurationServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupConfigurationServiceOutput(ctx *pulumi.Context, args LookupConfigurationServiceOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationServiceResult, error) {
			args := v.(LookupConfigurationServiceArgs)
			r, err := LookupConfigurationService(ctx, &args, opts...)
			var s LookupConfigurationServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationServiceResultOutput)
}

type LookupConfigurationServiceOutputArgs struct {
	// The name of Application Configuration Service.
	ConfigurationServiceName pulumi.StringInput `pulumi:"configurationServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupConfigurationServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationServiceArgs)(nil)).Elem()
}

// Application Configuration Service resource
type LookupConfigurationServiceResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationServiceResult)(nil)).Elem()
}

func (o LookupConfigurationServiceResultOutput) ToLookupConfigurationServiceResultOutput() LookupConfigurationServiceResultOutput {
	return o
}

func (o LookupConfigurationServiceResultOutput) ToLookupConfigurationServiceResultOutputWithContext(ctx context.Context) LookupConfigurationServiceResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupConfigurationServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupConfigurationServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Application Configuration Service properties payload
func (o LookupConfigurationServiceResultOutput) Properties() ConfigurationServicePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) ConfigurationServicePropertiesResponse { return v.Properties }).(ConfigurationServicePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupConfigurationServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupConfigurationServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationServiceResultOutput{})
}
