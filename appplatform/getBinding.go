// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Binding and its properties.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview.
func LookupBinding(ctx *pulumi.Context, args *LookupBindingArgs, opts ...pulumi.InvokeOption) (*LookupBindingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBindingResult
	err := ctx.Invoke("azure-native:appplatform:getBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBindingArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Binding resource.
	BindingName string `pulumi:"bindingName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Binding resource payload
type LookupBindingResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the Binding resource
	Properties BindingResourcePropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupBindingOutput(ctx *pulumi.Context, args LookupBindingOutputArgs, opts ...pulumi.InvokeOption) LookupBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBindingResult, error) {
			args := v.(LookupBindingArgs)
			r, err := LookupBinding(ctx, &args, opts...)
			var s LookupBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBindingResultOutput)
}

type LookupBindingOutputArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the Binding resource.
	BindingName pulumi.StringInput `pulumi:"bindingName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBindingArgs)(nil)).Elem()
}

// Binding resource payload
type LookupBindingResultOutput struct{ *pulumi.OutputState }

func (LookupBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBindingResult)(nil)).Elem()
}

func (o LookupBindingResultOutput) ToLookupBindingResultOutput() LookupBindingResultOutput {
	return o
}

func (o LookupBindingResultOutput) ToLookupBindingResultOutputWithContext(ctx context.Context) LookupBindingResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupBindingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Binding resource
func (o LookupBindingResultOutput) Properties() BindingResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupBindingResult) BindingResourcePropertiesResponse { return v.Properties }).(BindingResourcePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupBindingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBindingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupBindingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBindingResultOutput{})
}
