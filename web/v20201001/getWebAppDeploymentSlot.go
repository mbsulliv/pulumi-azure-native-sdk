// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a deployment by its ID for an app, or a deployment slot.
func LookupWebAppDeploymentSlot(ctx *pulumi.Context, args *LookupWebAppDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDeploymentSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppDeploymentSlotResult
	err := ctx.Invoke("azure-native:web/v20201001:getWebAppDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDeploymentSlotArgs struct {
	// Deployment ID.
	Id string `pulumi:"id"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API gets a deployment for the production slot.
	Slot string `pulumi:"slot"`
}

// User credentials used for publishing activity.
type LookupWebAppDeploymentSlotResult struct {
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
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppDeploymentSlotOutput(ctx *pulumi.Context, args LookupWebAppDeploymentSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDeploymentSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDeploymentSlotResult, error) {
			args := v.(LookupWebAppDeploymentSlotArgs)
			r, err := LookupWebAppDeploymentSlot(ctx, &args, opts...)
			var s LookupWebAppDeploymentSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDeploymentSlotResultOutput)
}

type LookupWebAppDeploymentSlotOutputArgs struct {
	// Deployment ID.
	Id pulumi.StringInput `pulumi:"id"`
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API gets a deployment for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppDeploymentSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentSlotArgs)(nil)).Elem()
}

// User credentials used for publishing activity.
type LookupWebAppDeploymentSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDeploymentSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentSlotResult)(nil)).Elem()
}

func (o LookupWebAppDeploymentSlotResultOutput) ToLookupWebAppDeploymentSlotResultOutput() LookupWebAppDeploymentSlotResultOutput {
	return o
}

func (o LookupWebAppDeploymentSlotResultOutput) ToLookupWebAppDeploymentSlotResultOutputWithContext(ctx context.Context) LookupWebAppDeploymentSlotResultOutput {
	return o
}

// True if deployment is currently active, false if completed and null if not started.
func (o LookupWebAppDeploymentSlotResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

// Who authored the deployment.
func (o LookupWebAppDeploymentSlotResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

// Author email.
func (o LookupWebAppDeploymentSlotResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

// Who performed the deployment.
func (o LookupWebAppDeploymentSlotResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

// Details on deployment.
func (o LookupWebAppDeploymentSlotResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

// End time.
func (o LookupWebAppDeploymentSlotResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppDeploymentSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppDeploymentSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Details about deployment status.
func (o LookupWebAppDeploymentSlotResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppDeploymentSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Start time.
func (o LookupWebAppDeploymentSlotResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Deployment status.
func (o LookupWebAppDeploymentSlotResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupWebAppDeploymentSlotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupWebAppDeploymentSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDeploymentSlotResultOutput{})
}
