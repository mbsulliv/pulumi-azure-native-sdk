// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns details of the API version.
func LookupApiVersion(ctx *pulumi.Context, args *LookupApiVersionArgs, opts ...pulumi.InvokeOption) (*LookupApiVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiVersionResult
	err := ctx.Invoke("azure-native:apicenter/v20240301:getApiVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiVersionArgs struct {
	// The name of the API.
	ApiName string `pulumi:"apiName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// The name of the API version.
	VersionName string `pulumi:"versionName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// API version entity.
type LookupApiVersionResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Current lifecycle stage of the API.
	LifecycleStage string `pulumi:"lifecycleStage"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// API version title.
	Title string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupApiVersionOutput(ctx *pulumi.Context, args LookupApiVersionOutputArgs, opts ...pulumi.InvokeOption) LookupApiVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiVersionResult, error) {
			args := v.(LookupApiVersionArgs)
			r, err := LookupApiVersion(ctx, &args, opts...)
			var s LookupApiVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiVersionResultOutput)
}

type LookupApiVersionOutputArgs struct {
	// The name of the API.
	ApiName pulumi.StringInput `pulumi:"apiName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The name of the API version.
	VersionName pulumi.StringInput `pulumi:"versionName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupApiVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiVersionArgs)(nil)).Elem()
}

// API version entity.
type LookupApiVersionResultOutput struct{ *pulumi.OutputState }

func (LookupApiVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiVersionResult)(nil)).Elem()
}

func (o LookupApiVersionResultOutput) ToLookupApiVersionResultOutput() LookupApiVersionResultOutput {
	return o
}

func (o LookupApiVersionResultOutput) ToLookupApiVersionResultOutputWithContext(ctx context.Context) LookupApiVersionResultOutput {
	return o
}

func (o LookupApiVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApiVersionResult] {
	return pulumix.Output[LookupApiVersionResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupApiVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Current lifecycle stage of the API.
func (o LookupApiVersionResultOutput) LifecycleStage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionResult) string { return v.LifecycleStage }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupApiVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApiVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// API version title.
func (o LookupApiVersionResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionResult) string { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiVersionResultOutput{})
}
