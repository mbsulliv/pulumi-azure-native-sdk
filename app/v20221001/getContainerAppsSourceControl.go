// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App SourceControl.
func LookupContainerAppsSourceControl(ctx *pulumi.Context, args *LookupContainerAppsSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppsSourceControlResult, error) {
	var rv LookupContainerAppsSourceControlResult
	err := ctx.Invoke("azure-native:app/v20221001:getContainerAppsSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerAppsSourceControlArgs struct {
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Container App SourceControl.
	SourceControlName string `pulumi:"sourceControlName"`
}

// Container App SourceControl.
type LookupContainerAppsSourceControlResult struct {
	// The branch which will trigger the auto deployment
	Branch *string `pulumi:"branch"`
	// Container App Revision Template with all possible settings and the
	// defaults if user did not provide them. The defaults are populated
	// as they were at the creation time
	GithubActionConfiguration *GithubActionConfigurationResponse `pulumi:"githubActionConfiguration"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Current provisioning State of the operation
	OperationState string `pulumi:"operationState"`
	// The repo url which will be integrated to ContainerApp.
	RepoUrl *string `pulumi:"repoUrl"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupContainerAppsSourceControlOutput(ctx *pulumi.Context, args LookupContainerAppsSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupContainerAppsSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerAppsSourceControlResult, error) {
			args := v.(LookupContainerAppsSourceControlArgs)
			r, err := LookupContainerAppsSourceControl(ctx, &args, opts...)
			var s LookupContainerAppsSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerAppsSourceControlResultOutput)
}

type LookupContainerAppsSourceControlOutputArgs struct {
	// Name of the Container App.
	ContainerAppName pulumi.StringInput `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Container App SourceControl.
	SourceControlName pulumi.StringInput `pulumi:"sourceControlName"`
}

func (LookupContainerAppsSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsSourceControlArgs)(nil)).Elem()
}

// Container App SourceControl.
type LookupContainerAppsSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupContainerAppsSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsSourceControlResult)(nil)).Elem()
}

func (o LookupContainerAppsSourceControlResultOutput) ToLookupContainerAppsSourceControlResultOutput() LookupContainerAppsSourceControlResultOutput {
	return o
}

func (o LookupContainerAppsSourceControlResultOutput) ToLookupContainerAppsSourceControlResultOutputWithContext(ctx context.Context) LookupContainerAppsSourceControlResultOutput {
	return o
}

// The branch which will trigger the auto deployment
func (o LookupContainerAppsSourceControlResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

// Container App Revision Template with all possible settings and the
// defaults if user did not provide them. The defaults are populated
// as they were at the creation time
func (o LookupContainerAppsSourceControlResultOutput) GithubActionConfiguration() GithubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) *GithubActionConfigurationResponse {
		return v.GithubActionConfiguration
	}).(GithubActionConfigurationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupContainerAppsSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupContainerAppsSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

// Current provisioning State of the operation
func (o LookupContainerAppsSourceControlResultOutput) OperationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.OperationState }).(pulumi.StringOutput)
}

// The repo url which will be integrated to ContainerApp.
func (o LookupContainerAppsSourceControlResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupContainerAppsSourceControlResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupContainerAppsSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerAppsSourceControlResultOutput{})
}
