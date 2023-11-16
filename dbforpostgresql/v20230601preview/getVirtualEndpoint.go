// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a virtual endpoint.
func LookupVirtualEndpoint(ctx *pulumi.Context, args *LookupVirtualEndpointArgs, opts ...pulumi.InvokeOption) (*LookupVirtualEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualEndpointResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20230601preview:getVirtualEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualEndpointArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual endpoint.
	VirtualEndpointName string `pulumi:"virtualEndpointName"`
}

// Represents a virtual endpoint for a server.
type LookupVirtualEndpointResult struct {
	// The endpoint type for the virtual endpoint.
	EndpointType *string `pulumi:"endpointType"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// List of members for a virtual endpoint
	Members []string `pulumi:"members"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// List of virtual endpoints for a server
	VirtualEndpoints []string `pulumi:"virtualEndpoints"`
}

func LookupVirtualEndpointOutput(ctx *pulumi.Context, args LookupVirtualEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualEndpointResult, error) {
			args := v.(LookupVirtualEndpointArgs)
			r, err := LookupVirtualEndpoint(ctx, &args, opts...)
			var s LookupVirtualEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualEndpointResultOutput)
}

type LookupVirtualEndpointOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of the virtual endpoint.
	VirtualEndpointName pulumi.StringInput `pulumi:"virtualEndpointName"`
}

func (LookupVirtualEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualEndpointArgs)(nil)).Elem()
}

// Represents a virtual endpoint for a server.
type LookupVirtualEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualEndpointResult)(nil)).Elem()
}

func (o LookupVirtualEndpointResultOutput) ToLookupVirtualEndpointResultOutput() LookupVirtualEndpointResultOutput {
	return o
}

func (o LookupVirtualEndpointResultOutput) ToLookupVirtualEndpointResultOutputWithContext(ctx context.Context) LookupVirtualEndpointResultOutput {
	return o
}

// The endpoint type for the virtual endpoint.
func (o LookupVirtualEndpointResultOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualEndpointResult) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupVirtualEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of members for a virtual endpoint
func (o LookupVirtualEndpointResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualEndpointResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupVirtualEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVirtualEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVirtualEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of virtual endpoints for a server
func (o LookupVirtualEndpointResultOutput) VirtualEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualEndpointResult) []string { return v.VirtualEndpoints }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualEndpointResultOutput{})
}
