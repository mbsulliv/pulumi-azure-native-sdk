// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the API Operation specified by its identifier.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2016-07-07, 2016-10-10, 2022-09-01-preview, 2023-03-01-preview.
func LookupApiOperation(ctx *pulumi.Context, args *LookupApiOperationArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiOperationResult
	err := ctx.Invoke("azure-native:apimanagement:getApiOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiOperationArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// API Operation details.
type LookupApiOperationResult struct {
	// Description of the operation. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Operation Name.
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Method string `pulumi:"method"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Operation Policies
	Policies *string `pulumi:"policies"`
	// An entity containing request details.
	Request *RequestContractResponse `pulumi:"request"`
	// Array of Operation responses.
	Responses []ResponseContractResponse `pulumi:"responses"`
	// Collection of URL template parameters.
	TemplateParameters []ParameterContractResponse `pulumi:"templateParameters"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	UrlTemplate string `pulumi:"urlTemplate"`
}

func LookupApiOperationOutput(ctx *pulumi.Context, args LookupApiOperationOutputArgs, opts ...pulumi.InvokeOption) LookupApiOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiOperationResult, error) {
			args := v.(LookupApiOperationArgs)
			r, err := LookupApiOperation(ctx, &args, opts...)
			var s LookupApiOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiOperationResultOutput)
}

type LookupApiOperationOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput `pulumi:"operationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationArgs)(nil)).Elem()
}

// API Operation details.
type LookupApiOperationResultOutput struct{ *pulumi.OutputState }

func (LookupApiOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationResult)(nil)).Elem()
}

func (o LookupApiOperationResultOutput) ToLookupApiOperationResultOutput() LookupApiOperationResultOutput {
	return o
}

func (o LookupApiOperationResultOutput) ToLookupApiOperationResultOutputWithContext(ctx context.Context) LookupApiOperationResultOutput {
	return o
}

func (o LookupApiOperationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApiOperationResult] {
	return pulumix.Output[LookupApiOperationResult]{
		OutputState: o.OutputState,
	}
}

// Description of the operation. May include HTML formatting tags.
func (o LookupApiOperationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiOperationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Operation Name.
func (o LookupApiOperationResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
func (o LookupApiOperationResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Method }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operation Policies
func (o LookupApiOperationResultOutput) Policies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiOperationResult) *string { return v.Policies }).(pulumi.StringPtrOutput)
}

// An entity containing request details.
func (o LookupApiOperationResultOutput) Request() RequestContractResponsePtrOutput {
	return o.ApplyT(func(v LookupApiOperationResult) *RequestContractResponse { return v.Request }).(RequestContractResponsePtrOutput)
}

// Array of Operation responses.
func (o LookupApiOperationResultOutput) Responses() ResponseContractResponseArrayOutput {
	return o.ApplyT(func(v LookupApiOperationResult) []ResponseContractResponse { return v.Responses }).(ResponseContractResponseArrayOutput)
}

// Collection of URL template parameters.
func (o LookupApiOperationResultOutput) TemplateParameters() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v LookupApiOperationResult) []ParameterContractResponse { return v.TemplateParameters }).(ParameterContractResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
func (o LookupApiOperationResultOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiOperationResultOutput{})
}
