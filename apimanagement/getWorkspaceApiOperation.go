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
// Azure REST API version: 2022-09-01-preview.
func LookupWorkspaceApiOperation(ctx *pulumi.Context, args *LookupWorkspaceApiOperationArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceApiOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceApiOperationResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceApiOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceApiOperationArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// API Operation details.
type LookupWorkspaceApiOperationResult struct {
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

func LookupWorkspaceApiOperationOutput(ctx *pulumi.Context, args LookupWorkspaceApiOperationOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceApiOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceApiOperationResult, error) {
			args := v.(LookupWorkspaceApiOperationArgs)
			r, err := LookupWorkspaceApiOperation(ctx, &args, opts...)
			var s LookupWorkspaceApiOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceApiOperationResultOutput)
}

type LookupWorkspaceApiOperationOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput `pulumi:"operationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceApiOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiOperationArgs)(nil)).Elem()
}

// API Operation details.
type LookupWorkspaceApiOperationResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceApiOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiOperationResult)(nil)).Elem()
}

func (o LookupWorkspaceApiOperationResultOutput) ToLookupWorkspaceApiOperationResultOutput() LookupWorkspaceApiOperationResultOutput {
	return o
}

func (o LookupWorkspaceApiOperationResultOutput) ToLookupWorkspaceApiOperationResultOutputWithContext(ctx context.Context) LookupWorkspaceApiOperationResultOutput {
	return o
}

func (o LookupWorkspaceApiOperationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceApiOperationResult] {
	return pulumix.Output[LookupWorkspaceApiOperationResult]{
		OutputState: o.OutputState,
	}
}

// Description of the operation. May include HTML formatting tags.
func (o LookupWorkspaceApiOperationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Operation Name.
func (o LookupWorkspaceApiOperationResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceApiOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
func (o LookupWorkspaceApiOperationResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) string { return v.Method }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceApiOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operation Policies
func (o LookupWorkspaceApiOperationResultOutput) Policies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) *string { return v.Policies }).(pulumi.StringPtrOutput)
}

// An entity containing request details.
func (o LookupWorkspaceApiOperationResultOutput) Request() RequestContractResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) *RequestContractResponse { return v.Request }).(RequestContractResponsePtrOutput)
}

// Array of Operation responses.
func (o LookupWorkspaceApiOperationResultOutput) Responses() ResponseContractResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) []ResponseContractResponse { return v.Responses }).(ResponseContractResponseArrayOutput)
}

// Collection of URL template parameters.
func (o LookupWorkspaceApiOperationResultOutput) TemplateParameters() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) []ParameterContractResponse { return v.TemplateParameters }).(ParameterContractResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceApiOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
func (o LookupWorkspaceApiOperationResultOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationResult) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceApiOperationResultOutput{})
}
