// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a job agent.
func LookupJobAgent(ctx *pulumi.Context, args *LookupJobAgentArgs, opts ...pulumi.InvokeOption) (*LookupJobAgentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobAgentResult
	err := ctx.Invoke("azure-native:sql/v20230801preview:getJobAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobAgentArgs struct {
	// The name of the job agent to be retrieved.
	JobAgentName string `pulumi:"jobAgentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// An Azure SQL job agent.
type LookupJobAgentResult struct {
	// Resource ID of the database to store job metadata in.
	DatabaseId string `pulumi:"databaseId"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The identity of the job agent.
	Identity *JobAgentIdentityResponse `pulumi:"identity"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The name and tier of the SKU.
	Sku *SkuResponse `pulumi:"sku"`
	// The state of the job agent.
	State string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupJobAgentOutput(ctx *pulumi.Context, args LookupJobAgentOutputArgs, opts ...pulumi.InvokeOption) LookupJobAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobAgentResult, error) {
			args := v.(LookupJobAgentArgs)
			r, err := LookupJobAgent(ctx, &args, opts...)
			var s LookupJobAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobAgentResultOutput)
}

type LookupJobAgentOutputArgs struct {
	// The name of the job agent to be retrieved.
	JobAgentName pulumi.StringInput `pulumi:"jobAgentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupJobAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobAgentArgs)(nil)).Elem()
}

// An Azure SQL job agent.
type LookupJobAgentResultOutput struct{ *pulumi.OutputState }

func (LookupJobAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobAgentResult)(nil)).Elem()
}

func (o LookupJobAgentResultOutput) ToLookupJobAgentResultOutput() LookupJobAgentResultOutput {
	return o
}

func (o LookupJobAgentResultOutput) ToLookupJobAgentResultOutputWithContext(ctx context.Context) LookupJobAgentResultOutput {
	return o
}

// Resource ID of the database to store job metadata in.
func (o LookupJobAgentResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupJobAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the job agent.
func (o LookupJobAgentResultOutput) Identity() JobAgentIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupJobAgentResult) *JobAgentIdentityResponse { return v.Identity }).(JobAgentIdentityResponsePtrOutput)
}

// Resource location.
func (o LookupJobAgentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupJobAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name and tier of the SKU.
func (o LookupJobAgentResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupJobAgentResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The state of the job agent.
func (o LookupJobAgentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.State }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupJobAgentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobAgentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupJobAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobAgentResultOutput{})
}
