// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get private endpoint connection properties
func LookupIotDpsResourcePrivateEndpointConnection(ctx *pulumi.Context, args *LookupIotDpsResourcePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupIotDpsResourcePrivateEndpointConnectionResult, error) {
	var rv LookupIotDpsResourcePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:devices/v20230301preview:getIotDpsResourcePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotDpsResourcePrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the provisioning service.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the provisioning service.
	ResourceName string `pulumi:"resourceName"`
}

// The private endpoint connection of a provisioning service
type LookupIotDpsResourcePrivateEndpointConnectionResult struct {
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupIotDpsResourcePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupIotDpsResourcePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupIotDpsResourcePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotDpsResourcePrivateEndpointConnectionResult, error) {
			args := v.(LookupIotDpsResourcePrivateEndpointConnectionArgs)
			r, err := LookupIotDpsResourcePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupIotDpsResourcePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotDpsResourcePrivateEndpointConnectionResultOutput)
}

type LookupIotDpsResourcePrivateEndpointConnectionOutputArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the provisioning service.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the provisioning service.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupIotDpsResourcePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourcePrivateEndpointConnectionArgs)(nil)).Elem()
}

// The private endpoint connection of a provisioning service
type LookupIotDpsResourcePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupIotDpsResourcePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourcePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) ToLookupIotDpsResourcePrivateEndpointConnectionResultOutput() LookupIotDpsResourcePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) ToLookupIotDpsResourcePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupIotDpsResourcePrivateEndpointConnectionResultOutput {
	return o
}

// The resource identifier.
func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of a private endpoint connection
func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type.
func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotDpsResourcePrivateEndpointConnectionResultOutput{})
}
