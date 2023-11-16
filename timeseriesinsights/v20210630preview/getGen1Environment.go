// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the environment with the specified name in the specified subscription and resource group.
func LookupGen1Environment(ctx *pulumi.Context, args *LookupGen1EnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupGen1EnvironmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGen1EnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getGen1Environment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGen1EnvironmentArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// Setting $expand=status will include the status of the internal services of the environment in the Time Series Insights service.
	Expand *string `pulumi:"expand"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. Gen1 environments have data retention limits.
type LookupGen1EnvironmentResult struct {
	// The time the resource was created.
	CreationTime string `pulumi:"creationTime"`
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn string `pulumi:"dataAccessFqdn"`
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId string `pulumi:"dataAccessId"`
	// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	DataRetentionTime string `pulumi:"dataRetentionTime"`
	// Resource Id
	Id string `pulumi:"id"`
	// The kind of the environment.
	// Expected value is 'Gen1'.
	Kind string `pulumi:"kind"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The list of event properties which will be used to partition data in the environment. Currently, only a single partition key property is supported.
	PartitionKeyProperties []TimeSeriesIdPropertyResponse `pulumi:"partitionKeyProperties"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
	Sku SkuResponse `pulumi:"sku"`
	// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
	Status EnvironmentStatusResponse `pulumi:"status"`
	// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
	StorageLimitExceededBehavior *string `pulumi:"storageLimitExceededBehavior"`
	// Indicates whether an environment supports Encryption at Rest with Customer Managed Key.
	SupportsCustomerManagedKey bool `pulumi:"supportsCustomerManagedKey"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupGen1EnvironmentOutput(ctx *pulumi.Context, args LookupGen1EnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupGen1EnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGen1EnvironmentResult, error) {
			args := v.(LookupGen1EnvironmentArgs)
			r, err := LookupGen1Environment(ctx, &args, opts...)
			var s LookupGen1EnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGen1EnvironmentResultOutput)
}

type LookupGen1EnvironmentOutputArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Setting $expand=status will include the status of the internal services of the environment in the Time Series Insights service.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGen1EnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen1EnvironmentArgs)(nil)).Elem()
}

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. Gen1 environments have data retention limits.
type LookupGen1EnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupGen1EnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen1EnvironmentResult)(nil)).Elem()
}

func (o LookupGen1EnvironmentResultOutput) ToLookupGen1EnvironmentResultOutput() LookupGen1EnvironmentResultOutput {
	return o
}

func (o LookupGen1EnvironmentResultOutput) ToLookupGen1EnvironmentResultOutputWithContext(ctx context.Context) LookupGen1EnvironmentResultOutput {
	return o
}

// The time the resource was created.
func (o LookupGen1EnvironmentResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
func (o LookupGen1EnvironmentResultOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
func (o LookupGen1EnvironmentResultOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.DataAccessId }).(pulumi.StringOutput)
}

// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
func (o LookupGen1EnvironmentResultOutput) DataRetentionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.DataRetentionTime }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupGen1EnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the environment.
// Expected value is 'Gen1'.
func (o LookupGen1EnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location
func (o LookupGen1EnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupGen1EnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of event properties which will be used to partition data in the environment. Currently, only a single partition key property is supported.
func (o LookupGen1EnvironmentResultOutput) PartitionKeyProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) []TimeSeriesIdPropertyResponse { return v.PartitionKeyProperties }).(TimeSeriesIdPropertyResponseArrayOutput)
}

// Provisioning state of the resource.
func (o LookupGen1EnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
func (o LookupGen1EnvironmentResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
func (o LookupGen1EnvironmentResultOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) EnvironmentStatusResponse { return v.Status }).(EnvironmentStatusResponseOutput)
}

// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
func (o LookupGen1EnvironmentResultOutput) StorageLimitExceededBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) *string { return v.StorageLimitExceededBehavior }).(pulumi.StringPtrOutput)
}

// Indicates whether an environment supports Encryption at Rest with Customer Managed Key.
func (o LookupGen1EnvironmentResultOutput) SupportsCustomerManagedKey() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) bool { return v.SupportsCustomerManagedKey }).(pulumi.BoolOutput)
}

// Resource tags
func (o LookupGen1EnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupGen1EnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGen1EnvironmentResultOutput{})
}
