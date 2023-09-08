// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridconnectivity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details about the service to the resource.
// Azure REST API version: 2023-03-15.
func LookupServiceConfiguration(ctx *pulumi.Context, args *LookupServiceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupServiceConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceConfigurationResult
	err := ctx.Invoke("azure-native:hybridconnectivity:getServiceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceConfigurationArgs struct {
	// The endpoint name.
	EndpointName string `pulumi:"endpointName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri string `pulumi:"resourceUri"`
	// The service name.
	ServiceConfigurationName string `pulumi:"serviceConfigurationName"`
}

// The service configuration details associated with the target resource.
type LookupServiceConfigurationResult struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The port on which service is enabled.
	Port *float64 `pulumi:"port"`
	// The resource provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource Id of the connectivity endpoint (optional).
	ResourceId *string `pulumi:"resourceId"`
	// Name of the service.
	ServiceName string `pulumi:"serviceName"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupServiceConfigurationOutput(ctx *pulumi.Context, args LookupServiceConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupServiceConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceConfigurationResult, error) {
			args := v.(LookupServiceConfigurationArgs)
			r, err := LookupServiceConfiguration(ctx, &args, opts...)
			var s LookupServiceConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceConfigurationResultOutput)
}

type LookupServiceConfigurationOutputArgs struct {
	// The endpoint name.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
	// The service name.
	ServiceConfigurationName pulumi.StringInput `pulumi:"serviceConfigurationName"`
}

func (LookupServiceConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceConfigurationArgs)(nil)).Elem()
}

// The service configuration details associated with the target resource.
type LookupServiceConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupServiceConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceConfigurationResult)(nil)).Elem()
}

func (o LookupServiceConfigurationResultOutput) ToLookupServiceConfigurationResultOutput() LookupServiceConfigurationResultOutput {
	return o
}

func (o LookupServiceConfigurationResultOutput) ToLookupServiceConfigurationResultOutputWithContext(ctx context.Context) LookupServiceConfigurationResultOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o LookupServiceConfigurationResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o LookupServiceConfigurationResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o LookupServiceConfigurationResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupServiceConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The timestamp of resource last modification (UTC)
func (o LookupServiceConfigurationResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o LookupServiceConfigurationResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o LookupServiceConfigurationResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupServiceConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The port on which service is enabled.
func (o LookupServiceConfigurationResultOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *float64 { return v.Port }).(pulumi.Float64PtrOutput)
}

// The resource provisioning state.
func (o LookupServiceConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource Id of the connectivity endpoint (optional).
func (o LookupServiceConfigurationResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Name of the service.
func (o LookupServiceConfigurationResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupServiceConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServiceConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceConfigurationResultOutput{})
}