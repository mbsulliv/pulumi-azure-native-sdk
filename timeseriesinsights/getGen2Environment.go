// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the environment with the specified name in the specified subscription and resource group.
// Azure REST API version: 2020-05-15.
func LookupGen2Environment(ctx *pulumi.Context, args *LookupGen2EnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupGen2EnvironmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGen2EnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getGen2Environment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGen2EnvironmentArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// Setting $expand=status will include the status of the internal services of the environment in the Time Series Insights service.
	Expand *string `pulumi:"expand"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. Gen2 environments do not have set data retention limits.
type LookupGen2EnvironmentResult struct {
	// The time the resource was created.
	CreationTime string `pulumi:"creationTime"`
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn string `pulumi:"dataAccessFqdn"`
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId string `pulumi:"dataAccessId"`
	// Resource Id
	Id string `pulumi:"id"`
	// The kind of the environment.
	// Expected value is 'Gen2'.
	Kind string `pulumi:"kind"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
	Sku SkuResponse `pulumi:"sku"`
	// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
	Status EnvironmentStatusResponse `pulumi:"status"`
	// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
	StorageConfiguration Gen2StorageConfigurationOutputResponse `pulumi:"storageConfiguration"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The list of event properties which will be used to define the environment's time series id.
	TimeSeriesIdProperties []TimeSeriesIdPropertyResponse `pulumi:"timeSeriesIdProperties"`
	// Resource type
	Type string `pulumi:"type"`
	// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
	WarmStoreConfiguration *WarmStoreConfigurationPropertiesResponse `pulumi:"warmStoreConfiguration"`
}

func LookupGen2EnvironmentOutput(ctx *pulumi.Context, args LookupGen2EnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupGen2EnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGen2EnvironmentResult, error) {
			args := v.(LookupGen2EnvironmentArgs)
			r, err := LookupGen2Environment(ctx, &args, opts...)
			var s LookupGen2EnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGen2EnvironmentResultOutput)
}

type LookupGen2EnvironmentOutputArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Setting $expand=status will include the status of the internal services of the environment in the Time Series Insights service.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGen2EnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen2EnvironmentArgs)(nil)).Elem()
}

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. Gen2 environments do not have set data retention limits.
type LookupGen2EnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupGen2EnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen2EnvironmentResult)(nil)).Elem()
}

func (o LookupGen2EnvironmentResultOutput) ToLookupGen2EnvironmentResultOutput() LookupGen2EnvironmentResultOutput {
	return o
}

func (o LookupGen2EnvironmentResultOutput) ToLookupGen2EnvironmentResultOutputWithContext(ctx context.Context) LookupGen2EnvironmentResultOutput {
	return o
}

func (o LookupGen2EnvironmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGen2EnvironmentResult] {
	return pulumix.Output[LookupGen2EnvironmentResult]{
		OutputState: o.OutputState,
	}
}

// The time the resource was created.
func (o LookupGen2EnvironmentResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
func (o LookupGen2EnvironmentResultOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
func (o LookupGen2EnvironmentResultOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.DataAccessId }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupGen2EnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the environment.
// Expected value is 'Gen2'.
func (o LookupGen2EnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location
func (o LookupGen2EnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupGen2EnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupGen2EnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
func (o LookupGen2EnvironmentResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
func (o LookupGen2EnvironmentResultOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) EnvironmentStatusResponse { return v.Status }).(EnvironmentStatusResponseOutput)
}

// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
func (o LookupGen2EnvironmentResultOutput) StorageConfiguration() Gen2StorageConfigurationOutputResponseOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) Gen2StorageConfigurationOutputResponse {
		return v.StorageConfiguration
	}).(Gen2StorageConfigurationOutputResponseOutput)
}

// Resource tags
func (o LookupGen2EnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The list of event properties which will be used to define the environment's time series id.
func (o LookupGen2EnvironmentResultOutput) TimeSeriesIdProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) []TimeSeriesIdPropertyResponse { return v.TimeSeriesIdProperties }).(TimeSeriesIdPropertyResponseArrayOutput)
}

// Resource type
func (o LookupGen2EnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
func (o LookupGen2EnvironmentResultOutput) WarmStoreConfiguration() WarmStoreConfigurationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) *WarmStoreConfigurationPropertiesResponse {
		return v.WarmStoreConfiguration
	}).(WarmStoreConfigurationPropertiesResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGen2EnvironmentResultOutput{})
}
