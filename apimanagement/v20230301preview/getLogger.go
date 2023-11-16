// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the logger specified by its identifier.
func LookupLogger(ctx *pulumi.Context, args *LookupLoggerArgs, opts ...pulumi.InvokeOption) (*LookupLoggerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLoggerResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getLogger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggerArgs struct {
	// Logger identifier. Must be unique in the API Management service instance.
	LoggerId string `pulumi:"loggerId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Logger details.
type LookupLoggerResult struct {
	// The name and SendRule connection string of the event hub for azureEventHub logger.
	// Instrumentation key for applicationInsights logger.
	Credentials map[string]string `pulumi:"credentials"`
	// Logger description.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Whether records are buffered in the logger before publishing. Default is assumed to be true.
	IsBuffered *bool `pulumi:"isBuffered"`
	// Logger type.
	LoggerType string `pulumi:"loggerType"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
	ResourceId *string `pulumi:"resourceId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupLoggerOutput(ctx *pulumi.Context, args LookupLoggerOutputArgs, opts ...pulumi.InvokeOption) LookupLoggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoggerResult, error) {
			args := v.(LookupLoggerArgs)
			r, err := LookupLogger(ctx, &args, opts...)
			var s LookupLoggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoggerResultOutput)
}

type LookupLoggerOutputArgs struct {
	// Logger identifier. Must be unique in the API Management service instance.
	LoggerId pulumi.StringInput `pulumi:"loggerId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupLoggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerArgs)(nil)).Elem()
}

// Logger details.
type LookupLoggerResultOutput struct{ *pulumi.OutputState }

func (LookupLoggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerResult)(nil)).Elem()
}

func (o LookupLoggerResultOutput) ToLookupLoggerResultOutput() LookupLoggerResultOutput {
	return o
}

func (o LookupLoggerResultOutput) ToLookupLoggerResultOutputWithContext(ctx context.Context) LookupLoggerResultOutput {
	return o
}

// The name and SendRule connection string of the event hub for azureEventHub logger.
// Instrumentation key for applicationInsights logger.
func (o LookupLoggerResultOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLoggerResult) map[string]string { return v.Credentials }).(pulumi.StringMapOutput)
}

// Logger description.
func (o LookupLoggerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLoggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether records are buffered in the logger before publishing. Default is assumed to be true.
func (o LookupLoggerResultOutput) IsBuffered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoggerResult) *bool { return v.IsBuffered }).(pulumi.BoolPtrOutput)
}

// Logger type.
func (o LookupLoggerResultOutput) LoggerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.LoggerType }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLoggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
func (o LookupLoggerResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLoggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoggerResultOutput{})
}
