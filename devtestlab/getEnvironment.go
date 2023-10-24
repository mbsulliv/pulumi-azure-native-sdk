// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get environment.
// Azure REST API version: 2018-09-15.
//
// Other available API versions: 2016-05-15.
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:devtestlab:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	// Specify the $expand query. Example: 'properties($select=deploymentProperties)'
	Expand *string `pulumi:"expand"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the environment.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the user profile.
	UserName string `pulumi:"userName"`
}

// An environment, which is essentially an ARM template deployment.
type LookupEnvironmentResult struct {
	// The display name of the Azure Resource Manager template that produced the environment.
	ArmTemplateDisplayName *string `pulumi:"armTemplateDisplayName"`
	// The creator of the environment.
	CreatedByUser string `pulumi:"createdByUser"`
	// The deployment properties of the environment.
	DeploymentProperties *EnvironmentDeploymentPropertiesResponse `pulumi:"deploymentProperties"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The identifier of the resource group containing the environment's resources.
	ResourceGroupId string `pulumi:"resourceGroupId"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
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
	// Specify the $expand query. Example: 'properties($select=deploymentProperties)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the environment.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the user profile.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

// An environment, which is essentially an ARM template deployment.
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

// The display name of the Azure Resource Manager template that produced the environment.
func (o LookupEnvironmentResultOutput) ArmTemplateDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.ArmTemplateDisplayName }).(pulumi.StringPtrOutput)
}

// The creator of the environment.
func (o LookupEnvironmentResultOutput) CreatedByUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.CreatedByUser }).(pulumi.StringOutput)
}

// The deployment properties of the environment.
func (o LookupEnvironmentResultOutput) DeploymentProperties() EnvironmentDeploymentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentDeploymentPropertiesResponse {
		return v.DeploymentProperties
	}).(EnvironmentDeploymentPropertiesResponsePtrOutput)
}

// The identifier of the resource.
func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource.
func (o LookupEnvironmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o LookupEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The identifier of the resource group containing the environment's resources.
func (o LookupEnvironmentResultOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o LookupEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupEnvironmentResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
