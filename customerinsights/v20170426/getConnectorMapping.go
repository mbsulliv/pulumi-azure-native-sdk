// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a connector mapping in the connector.
func LookupConnectorMapping(ctx *pulumi.Context, args *LookupConnectorMappingArgs, opts ...pulumi.InvokeOption) (*LookupConnectorMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorMappingResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getConnectorMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorMappingArgs struct {
	// The name of the connector.
	ConnectorName string `pulumi:"connectorName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the connector mapping.
	MappingName string `pulumi:"mappingName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connector mapping resource format.
type LookupConnectorMappingResult struct {
	// The connector mapping name
	ConnectorMappingName string `pulumi:"connectorMappingName"`
	// The connector name.
	ConnectorName string `pulumi:"connectorName"`
	// Type of connector.
	ConnectorType *string `pulumi:"connectorType"`
	// The created time.
	Created string `pulumi:"created"`
	// The DataFormat ID.
	DataFormatId string `pulumi:"dataFormatId"`
	// The description of the connector mapping.
	Description *string `pulumi:"description"`
	// Display name for the connector mapping.
	DisplayName *string `pulumi:"displayName"`
	// Defines which entity type the file should map to.
	EntityType string `pulumi:"entityType"`
	// The mapping entity name.
	EntityTypeName string `pulumi:"entityTypeName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The last modified time.
	LastModified string `pulumi:"lastModified"`
	// The properties of the mapping.
	MappingProperties ConnectorMappingPropertiesResponse `pulumi:"mappingProperties"`
	// Resource name.
	Name string `pulumi:"name"`
	// The next run time based on customer's settings.
	NextRunTime string `pulumi:"nextRunTime"`
	// The RunId.
	RunId string `pulumi:"runId"`
	// State of connector mapping.
	State string `pulumi:"state"`
	// The hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupConnectorMappingOutput(ctx *pulumi.Context, args LookupConnectorMappingOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorMappingResult, error) {
			args := v.(LookupConnectorMappingArgs)
			r, err := LookupConnectorMapping(ctx, &args, opts...)
			var s LookupConnectorMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorMappingResultOutput)
}

type LookupConnectorMappingOutputArgs struct {
	// The name of the connector.
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
	// The name of the hub.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// The name of the connector mapping.
	MappingName pulumi.StringInput `pulumi:"mappingName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectorMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorMappingArgs)(nil)).Elem()
}

// The connector mapping resource format.
type LookupConnectorMappingResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorMappingResult)(nil)).Elem()
}

func (o LookupConnectorMappingResultOutput) ToLookupConnectorMappingResultOutput() LookupConnectorMappingResultOutput {
	return o
}

func (o LookupConnectorMappingResultOutput) ToLookupConnectorMappingResultOutputWithContext(ctx context.Context) LookupConnectorMappingResultOutput {
	return o
}

// The connector mapping name
func (o LookupConnectorMappingResultOutput) ConnectorMappingName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.ConnectorMappingName }).(pulumi.StringOutput)
}

// The connector name.
func (o LookupConnectorMappingResultOutput) ConnectorName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.ConnectorName }).(pulumi.StringOutput)
}

// Type of connector.
func (o LookupConnectorMappingResultOutput) ConnectorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) *string { return v.ConnectorType }).(pulumi.StringPtrOutput)
}

// The created time.
func (o LookupConnectorMappingResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Created }).(pulumi.StringOutput)
}

// The DataFormat ID.
func (o LookupConnectorMappingResultOutput) DataFormatId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.DataFormatId }).(pulumi.StringOutput)
}

// The description of the connector mapping.
func (o LookupConnectorMappingResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name for the connector mapping.
func (o LookupConnectorMappingResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Defines which entity type the file should map to.
func (o LookupConnectorMappingResultOutput) EntityType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.EntityType }).(pulumi.StringOutput)
}

// The mapping entity name.
func (o LookupConnectorMappingResultOutput) EntityTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.EntityTypeName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupConnectorMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last modified time.
func (o LookupConnectorMappingResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The properties of the mapping.
func (o LookupConnectorMappingResultOutput) MappingProperties() ConnectorMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) ConnectorMappingPropertiesResponse { return v.MappingProperties }).(ConnectorMappingPropertiesResponseOutput)
}

// Resource name.
func (o LookupConnectorMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The next run time based on customer's settings.
func (o LookupConnectorMappingResultOutput) NextRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.NextRunTime }).(pulumi.StringOutput)
}

// The RunId.
func (o LookupConnectorMappingResultOutput) RunId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.RunId }).(pulumi.StringOutput)
}

// State of connector mapping.
func (o LookupConnectorMappingResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.State }).(pulumi.StringOutput)
}

// The hub name.
func (o LookupConnectorMappingResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupConnectorMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorMappingResultOutput{})
}
