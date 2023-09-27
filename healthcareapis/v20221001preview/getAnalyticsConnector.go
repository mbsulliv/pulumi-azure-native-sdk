// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the properties of the specified Analytics Connector.
func LookupAnalyticsConnector(ctx *pulumi.Context, args *LookupAnalyticsConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAnalyticsConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAnalyticsConnectorResult
	err := ctx.Invoke("azure-native:healthcareapis/v20221001preview:getAnalyticsConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnalyticsConnectorArgs struct {
	// The name of Analytics Connector resource.
	AnalyticsConnectorName string `pulumi:"analyticsConnectorName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Analytics Connector definition.
type LookupAnalyticsConnectorResult struct {
	// Data destination configuration for Analytics Connector.
	DataDestinationConfiguration AnalyticsConnectorDataLakeDataDestinationResponse `pulumi:"dataDestinationConfiguration"`
	// Data mapping configuration for Analytics Connector.
	DataMappingConfiguration AnalyticsConnectorFhirToParquetMappingResponse `pulumi:"dataMappingConfiguration"`
	// Data source for Analytics Connector.
	DataSourceConfiguration AnalyticsConnectorFhirServiceDataSourceResponse `pulumi:"dataSourceConfiguration"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityResponseIdentity `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupAnalyticsConnectorOutput(ctx *pulumi.Context, args LookupAnalyticsConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAnalyticsConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnalyticsConnectorResult, error) {
			args := v.(LookupAnalyticsConnectorArgs)
			r, err := LookupAnalyticsConnector(ctx, &args, opts...)
			var s LookupAnalyticsConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnalyticsConnectorResultOutput)
}

type LookupAnalyticsConnectorOutputArgs struct {
	// The name of Analytics Connector resource.
	AnalyticsConnectorName pulumi.StringInput `pulumi:"analyticsConnectorName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAnalyticsConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalyticsConnectorArgs)(nil)).Elem()
}

// Analytics Connector definition.
type LookupAnalyticsConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAnalyticsConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalyticsConnectorResult)(nil)).Elem()
}

func (o LookupAnalyticsConnectorResultOutput) ToLookupAnalyticsConnectorResultOutput() LookupAnalyticsConnectorResultOutput {
	return o
}

func (o LookupAnalyticsConnectorResultOutput) ToLookupAnalyticsConnectorResultOutputWithContext(ctx context.Context) LookupAnalyticsConnectorResultOutput {
	return o
}

func (o LookupAnalyticsConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAnalyticsConnectorResult] {
	return pulumix.Output[LookupAnalyticsConnectorResult]{
		OutputState: o.OutputState,
	}
}

// Data destination configuration for Analytics Connector.
func (o LookupAnalyticsConnectorResultOutput) DataDestinationConfiguration() AnalyticsConnectorDataLakeDataDestinationResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) AnalyticsConnectorDataLakeDataDestinationResponse {
		return v.DataDestinationConfiguration
	}).(AnalyticsConnectorDataLakeDataDestinationResponseOutput)
}

// Data mapping configuration for Analytics Connector.
func (o LookupAnalyticsConnectorResultOutput) DataMappingConfiguration() AnalyticsConnectorFhirToParquetMappingResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) AnalyticsConnectorFhirToParquetMappingResponse {
		return v.DataMappingConfiguration
	}).(AnalyticsConnectorFhirToParquetMappingResponseOutput)
}

// Data source for Analytics Connector.
func (o LookupAnalyticsConnectorResultOutput) DataSourceConfiguration() AnalyticsConnectorFhirServiceDataSourceResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) AnalyticsConnectorFhirServiceDataSourceResponse {
		return v.DataSourceConfiguration
	}).(AnalyticsConnectorFhirServiceDataSourceResponseOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupAnalyticsConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource identifier.
func (o LookupAnalyticsConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o LookupAnalyticsConnectorResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// The resource location.
func (o LookupAnalyticsConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupAnalyticsConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o LookupAnalyticsConnectorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupAnalyticsConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAnalyticsConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupAnalyticsConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnalyticsConnectorResultOutput{})
}
