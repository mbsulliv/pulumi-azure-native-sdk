// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Get a deployment by its ID for an app, or a deployment slot.
func LookupWebAppDeployment(ctx *pulumi.Context, args *LookupWebAppDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDeploymentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppDeploymentResult
	err := ctx.Invoke("azure-native:web/v20220901:getWebAppDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDeploymentArgs struct {
	// Deployment ID.
	Id string `pulumi:"id"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// User credentials used for publishing activity.
type LookupWebAppDeploymentResult struct {
	// True if deployment is currently active, false if completed and null if not started.
	Active *bool `pulumi:"active"`
	// Who authored the deployment.
	Author *string `pulumi:"author"`
	// Author email.
	AuthorEmail *string `pulumi:"authorEmail"`
	// Who performed the deployment.
	Deployer *string `pulumi:"deployer"`
	// Details on deployment.
	Details *string `pulumi:"details"`
	// End time.
	EndTime *string `pulumi:"endTime"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Details about deployment status.
	Message *string `pulumi:"message"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Start time.
	StartTime *string `pulumi:"startTime"`
	// Deployment status.
	Status *int `pulumi:"status"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppDeploymentOutput(ctx *pulumi.Context, args LookupWebAppDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDeploymentResult, error) {
			args := v.(LookupWebAppDeploymentArgs)
			r, err := LookupWebAppDeployment(ctx, &args, opts...)
			var s LookupWebAppDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDeploymentResultOutput)
}

type LookupWebAppDeploymentOutputArgs struct {
	// Deployment ID.
	Id pulumi.StringInput `pulumi:"id"`
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentArgs)(nil)).Elem()
}

// User credentials used for publishing activity.
type LookupWebAppDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentResult)(nil)).Elem()
}

func (o LookupWebAppDeploymentResultOutput) ToLookupWebAppDeploymentResultOutput() LookupWebAppDeploymentResultOutput {
	return o
}

func (o LookupWebAppDeploymentResultOutput) ToLookupWebAppDeploymentResultOutputWithContext(ctx context.Context) LookupWebAppDeploymentResultOutput {
	return o
}

func (o LookupWebAppDeploymentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppDeploymentResult] {
	return pulumix.Output[LookupWebAppDeploymentResult]{
		OutputState: o.OutputState,
	}
}

// True if deployment is currently active, false if completed and null if not started.
func (o LookupWebAppDeploymentResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

// Who authored the deployment.
func (o LookupWebAppDeploymentResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

// Author email.
func (o LookupWebAppDeploymentResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

// Who performed the deployment.
func (o LookupWebAppDeploymentResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

// Details on deployment.
func (o LookupWebAppDeploymentResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

// End time.
func (o LookupWebAppDeploymentResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Details about deployment status.
func (o LookupWebAppDeploymentResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Start time.
func (o LookupWebAppDeploymentResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Deployment status.
func (o LookupWebAppDeploymentResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupWebAppDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDeploymentResultOutput{})
}
