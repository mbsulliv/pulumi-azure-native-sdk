// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an azure databricks accessConnector.
func LookupAccessConnector(ctx *pulumi.Context, args *LookupAccessConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAccessConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessConnectorResult
	err := ctx.Invoke("azure-native:databricks/v20220401preview:getAccessConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessConnectorArgs struct {
	// The name of the azure databricks accessConnector.
	ConnectorName string `pulumi:"connectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about azure databricks accessConnector.
type LookupAccessConnectorResult struct {
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityDataResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Databricks accessConnector properties
	Properties AccessConnectorPropertiesResponse `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

func LookupAccessConnectorOutput(ctx *pulumi.Context, args LookupAccessConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAccessConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessConnectorResult, error) {
			args := v.(LookupAccessConnectorArgs)
			r, err := LookupAccessConnector(ctx, &args, opts...)
			var s LookupAccessConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessConnectorResultOutput)
}

type LookupAccessConnectorOutputArgs struct {
	// The name of the azure databricks accessConnector.
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessConnectorArgs)(nil)).Elem()
}

// Information about azure databricks accessConnector.
type LookupAccessConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAccessConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessConnectorResult)(nil)).Elem()
}

func (o LookupAccessConnectorResultOutput) ToLookupAccessConnectorResultOutput() LookupAccessConnectorResultOutput {
	return o
}

func (o LookupAccessConnectorResultOutput) ToLookupAccessConnectorResultOutputWithContext(ctx context.Context) LookupAccessConnectorResultOutput {
	return o
}

func (o LookupAccessConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccessConnectorResult] {
	return pulumix.Output[LookupAccessConnectorResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAccessConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupAccessConnectorResultOutput) Identity() IdentityDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) *IdentityDataResponse { return v.Identity }).(IdentityDataResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupAccessConnectorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAccessConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Databricks accessConnector properties
func (o LookupAccessConnectorResultOutput) Properties() AccessConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) AccessConnectorPropertiesResponse { return v.Properties }).(AccessConnectorPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o LookupAccessConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAccessConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupAccessConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessConnectorResultOutput{})
}
