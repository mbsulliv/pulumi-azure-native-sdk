// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the Application Live  and its properties.
func LookupDevToolPortal(ctx *pulumi.Context, args *LookupDevToolPortalArgs, opts ...pulumi.InvokeOption) (*LookupDevToolPortalResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDevToolPortalResult
	err := ctx.Invoke("azure-native:appplatform/v20231101preview:getDevToolPortal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDevToolPortalArgs struct {
	// The name of Dev Tool Portal.
	DevToolPortalName string `pulumi:"devToolPortalName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Dev Tool Portal resource
type LookupDevToolPortalResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Dev Tool Portal properties payload
	Properties DevToolPortalPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDevToolPortalResult
func (val *LookupDevToolPortalResult) Defaults() *LookupDevToolPortalResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupDevToolPortalOutput(ctx *pulumi.Context, args LookupDevToolPortalOutputArgs, opts ...pulumi.InvokeOption) LookupDevToolPortalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevToolPortalResult, error) {
			args := v.(LookupDevToolPortalArgs)
			r, err := LookupDevToolPortal(ctx, &args, opts...)
			var s LookupDevToolPortalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevToolPortalResultOutput)
}

type LookupDevToolPortalOutputArgs struct {
	// The name of Dev Tool Portal.
	DevToolPortalName pulumi.StringInput `pulumi:"devToolPortalName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDevToolPortalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevToolPortalArgs)(nil)).Elem()
}

// Dev Tool Portal resource
type LookupDevToolPortalResultOutput struct{ *pulumi.OutputState }

func (LookupDevToolPortalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevToolPortalResult)(nil)).Elem()
}

func (o LookupDevToolPortalResultOutput) ToLookupDevToolPortalResultOutput() LookupDevToolPortalResultOutput {
	return o
}

func (o LookupDevToolPortalResultOutput) ToLookupDevToolPortalResultOutputWithContext(ctx context.Context) LookupDevToolPortalResultOutput {
	return o
}

func (o LookupDevToolPortalResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDevToolPortalResult] {
	return pulumix.Output[LookupDevToolPortalResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupDevToolPortalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupDevToolPortalResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) string { return v.Name }).(pulumi.StringOutput)
}

// Dev Tool Portal properties payload
func (o LookupDevToolPortalResultOutput) Properties() DevToolPortalPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) DevToolPortalPropertiesResponse { return v.Properties }).(DevToolPortalPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDevToolPortalResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupDevToolPortalResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevToolPortalResultOutput{})
}
