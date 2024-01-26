// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthcareapis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified Iot Connector FHIR destination.
// Azure REST API version: 2023-02-28.
//
// Other available API versions: 2023-09-06, 2023-11-01, 2023-12-01.
func LookupIotConnectorFhirDestination(ctx *pulumi.Context, args *LookupIotConnectorFhirDestinationArgs, opts ...pulumi.InvokeOption) (*LookupIotConnectorFhirDestinationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIotConnectorFhirDestinationResult
	err := ctx.Invoke("azure-native:healthcareapis:getIotConnectorFhirDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotConnectorFhirDestinationArgs struct {
	// The name of IoT Connector FHIR destination resource.
	FhirDestinationName string `pulumi:"fhirDestinationName"`
	// The name of IoT Connector resource.
	IotConnectorName string `pulumi:"iotConnectorName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// IoT Connector FHIR destination definition.
type LookupIotConnectorFhirDestinationResult struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// FHIR Mappings
	FhirMapping IotMappingPropertiesResponse `pulumi:"fhirMapping"`
	// Fully qualified resource id of the FHIR service to connect to.
	FhirServiceResourceId string `pulumi:"fhirServiceResourceId"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// Determines how resource identity is resolved on the destination.
	ResourceIdentityResolutionType string `pulumi:"resourceIdentityResolutionType"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupIotConnectorFhirDestinationOutput(ctx *pulumi.Context, args LookupIotConnectorFhirDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupIotConnectorFhirDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotConnectorFhirDestinationResult, error) {
			args := v.(LookupIotConnectorFhirDestinationArgs)
			r, err := LookupIotConnectorFhirDestination(ctx, &args, opts...)
			var s LookupIotConnectorFhirDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotConnectorFhirDestinationResultOutput)
}

type LookupIotConnectorFhirDestinationOutputArgs struct {
	// The name of IoT Connector FHIR destination resource.
	FhirDestinationName pulumi.StringInput `pulumi:"fhirDestinationName"`
	// The name of IoT Connector resource.
	IotConnectorName pulumi.StringInput `pulumi:"iotConnectorName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIotConnectorFhirDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorFhirDestinationArgs)(nil)).Elem()
}

// IoT Connector FHIR destination definition.
type LookupIotConnectorFhirDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupIotConnectorFhirDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorFhirDestinationResult)(nil)).Elem()
}

func (o LookupIotConnectorFhirDestinationResultOutput) ToLookupIotConnectorFhirDestinationResultOutput() LookupIotConnectorFhirDestinationResultOutput {
	return o
}

func (o LookupIotConnectorFhirDestinationResultOutput) ToLookupIotConnectorFhirDestinationResultOutputWithContext(ctx context.Context) LookupIotConnectorFhirDestinationResultOutput {
	return o
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupIotConnectorFhirDestinationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// FHIR Mappings
func (o LookupIotConnectorFhirDestinationResultOutput) FhirMapping() IotMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) IotMappingPropertiesResponse { return v.FhirMapping }).(IotMappingPropertiesResponseOutput)
}

// Fully qualified resource id of the FHIR service to connect to.
func (o LookupIotConnectorFhirDestinationResultOutput) FhirServiceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.FhirServiceResourceId }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupIotConnectorFhirDestinationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource location.
func (o LookupIotConnectorFhirDestinationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupIotConnectorFhirDestinationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Determines how resource identity is resolved on the destination.
func (o LookupIotConnectorFhirDestinationResultOutput) ResourceIdentityResolutionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.ResourceIdentityResolutionType }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupIotConnectorFhirDestinationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type.
func (o LookupIotConnectorFhirDestinationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotConnectorFhirDestinationResultOutput{})
}
