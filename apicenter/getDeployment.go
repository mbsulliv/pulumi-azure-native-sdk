// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apicenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns details of the API deployment.
// Azure REST API version: 2024-03-01.
func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDeploymentResult
	err := ctx.Invoke("azure-native:apicenter:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentArgs struct {
	// The name of the API.
	ApiName string `pulumi:"apiName"`
	// The name of the API deployment.
	DeploymentName string `pulumi:"deploymentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// API deployment entity.
type LookupDeploymentResult struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties interface{} `pulumi:"customProperties"`
	// API center-scoped definition resource ID.
	DefinitionId *string `pulumi:"definitionId"`
	// Description of the deployment.
	Description *string `pulumi:"description"`
	// API center-scoped environment resource ID.
	EnvironmentId *string `pulumi:"environmentId"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name   string                    `pulumi:"name"`
	Server *DeploymentServerResponse `pulumi:"server"`
	// State of API deployment.
	State *string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// API deployment title
	Title *string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDeploymentOutput(ctx *pulumi.Context, args LookupDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentResult, error) {
			args := v.(LookupDeploymentArgs)
			r, err := LookupDeployment(ctx, &args, opts...)
			var s LookupDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentResultOutput)
}

type LookupDeploymentOutputArgs struct {
	// The name of the API.
	ApiName pulumi.StringInput `pulumi:"apiName"`
	// The name of the API deployment.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArgs)(nil)).Elem()
}

// API deployment entity.
type LookupDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentResult)(nil)).Elem()
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutput() LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutputWithContext(ctx context.Context) LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDeploymentResult] {
	return pulumix.Output[LookupDeploymentResult]{
		OutputState: o.OutputState,
	}
}

// The custom metadata defined for API catalog entities.
func (o LookupDeploymentResultOutput) CustomProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDeploymentResult) interface{} { return v.CustomProperties }).(pulumi.AnyOutput)
}

// API center-scoped definition resource ID.
func (o LookupDeploymentResultOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *string { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Description of the deployment.
func (o LookupDeploymentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// API center-scoped environment resource ID.
func (o LookupDeploymentResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentResultOutput) Server() DeploymentServerResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *DeploymentServerResponse { return v.Server }).(DeploymentServerResponsePtrOutput)
}

// State of API deployment.
func (o LookupDeploymentResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// API deployment title
func (o LookupDeploymentResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentResultOutput{})
}