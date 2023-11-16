// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get service runner.
func LookupServiceRunner(ctx *pulumi.Context, args *LookupServiceRunnerArgs, opts ...pulumi.InvokeOption) (*LookupServiceRunnerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceRunnerResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getServiceRunner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceRunnerArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the service runner.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A container for a managed identity to execute DevTest lab services.
type LookupServiceRunnerResult struct {
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupServiceRunnerOutput(ctx *pulumi.Context, args LookupServiceRunnerOutputArgs, opts ...pulumi.InvokeOption) LookupServiceRunnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceRunnerResult, error) {
			args := v.(LookupServiceRunnerArgs)
			r, err := LookupServiceRunner(ctx, &args, opts...)
			var s LookupServiceRunnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceRunnerResultOutput)
}

type LookupServiceRunnerOutputArgs struct {
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the service runner.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServiceRunnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRunnerArgs)(nil)).Elem()
}

// A container for a managed identity to execute DevTest lab services.
type LookupServiceRunnerResultOutput struct{ *pulumi.OutputState }

func (LookupServiceRunnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRunnerResult)(nil)).Elem()
}

func (o LookupServiceRunnerResultOutput) ToLookupServiceRunnerResultOutput() LookupServiceRunnerResultOutput {
	return o
}

func (o LookupServiceRunnerResultOutput) ToLookupServiceRunnerResultOutputWithContext(ctx context.Context) LookupServiceRunnerResultOutput {
	return o
}

// The identifier of the resource.
func (o LookupServiceRunnerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupServiceRunnerResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The location of the resource.
func (o LookupServiceRunnerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupServiceRunnerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o LookupServiceRunnerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupServiceRunnerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceRunnerResultOutput{})
}
