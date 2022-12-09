// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service Registry resource
func LookupServiceRegistry(ctx *pulumi.Context, args *LookupServiceRegistryArgs, opts ...pulumi.InvokeOption) (*LookupServiceRegistryResult, error) {
	var rv LookupServiceRegistryResult
	err := ctx.Invoke("azure-native:appplatform/v20221201:getServiceRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceRegistryArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// The name of Service Registry.
	ServiceRegistryName string `pulumi:"serviceRegistryName"`
}

// Service Registry resource
type LookupServiceRegistryResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Service Registry properties payload
	Properties ServiceRegistryPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupServiceRegistryOutput(ctx *pulumi.Context, args LookupServiceRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupServiceRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceRegistryResult, error) {
			args := v.(LookupServiceRegistryArgs)
			r, err := LookupServiceRegistry(ctx, &args, opts...)
			var s LookupServiceRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceRegistryResultOutput)
}

type LookupServiceRegistryOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The name of Service Registry.
	ServiceRegistryName pulumi.StringInput `pulumi:"serviceRegistryName"`
}

func (LookupServiceRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRegistryArgs)(nil)).Elem()
}

// Service Registry resource
type LookupServiceRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupServiceRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRegistryResult)(nil)).Elem()
}

func (o LookupServiceRegistryResultOutput) ToLookupServiceRegistryResultOutput() LookupServiceRegistryResultOutput {
	return o
}

func (o LookupServiceRegistryResultOutput) ToLookupServiceRegistryResultOutputWithContext(ctx context.Context) LookupServiceRegistryResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupServiceRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupServiceRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Service Registry properties payload
func (o LookupServiceRegistryResultOutput) Properties() ServiceRegistryPropertiesResponseOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) ServiceRegistryPropertiesResponse { return v.Properties }).(ServiceRegistryPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupServiceRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupServiceRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceRegistryResultOutput{})
}
