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

// Returns details of the environment.
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:apicenter/v20240301:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	// The name of the environment.
	EnvironmentName string `pulumi:"environmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Environment entity.
type LookupEnvironmentResult struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties interface{} `pulumi:"customProperties"`
	Description      *string     `pulumi:"description"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Environment kind.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name       string              `pulumi:"name"`
	Onboarding *OnboardingResponse `pulumi:"onboarding"`
	// Server information of the environment.
	Server *EnvironmentServerResponse `pulumi:"server"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Environment title.
	Title string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	// The name of the environment.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

// Environment entity.
type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEnvironmentResult] {
	return pulumix.Output[LookupEnvironmentResult]{
		OutputState: o.OutputState,
	}
}

// The custom metadata defined for API catalog entities.
func (o LookupEnvironmentResultOutput) CustomProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) interface{} { return v.CustomProperties }).(pulumi.AnyOutput)
}

func (o LookupEnvironmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Environment kind.
func (o LookupEnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Onboarding() OnboardingResponsePtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *OnboardingResponse { return v.Onboarding }).(OnboardingResponsePtrOutput)
}

// Server information of the environment.
func (o LookupEnvironmentResultOutput) Server() EnvironmentServerResponsePtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentServerResponse { return v.Server }).(EnvironmentServerResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEnvironmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Environment title.
func (o LookupEnvironmentResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
