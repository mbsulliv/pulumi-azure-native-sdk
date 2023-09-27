// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an integration account.
// Azure REST API version: 2019-05-01.
func LookupIntegrationAccount(ctx *pulumi.Context, args *LookupIntegrationAccountArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationAccountResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The integration account.
type LookupIntegrationAccountResult struct {
	// The resource id.
	Id string `pulumi:"id"`
	// The integration service environment.
	IntegrationServiceEnvironment *ResourceReferenceResponse `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The sku.
	Sku *IntegrationAccountSkuResponse `pulumi:"sku"`
	// The workflow state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}

func LookupIntegrationAccountOutput(ctx *pulumi.Context, args LookupIntegrationAccountOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountResult, error) {
			args := v.(LookupIntegrationAccountArgs)
			r, err := LookupIntegrationAccount(ctx, &args, opts...)
			var s LookupIntegrationAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountResultOutput)
}

type LookupIntegrationAccountOutputArgs struct {
	// The integration account name.
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountArgs)(nil)).Elem()
}

// The integration account.
type LookupIntegrationAccountResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountResult)(nil)).Elem()
}

func (o LookupIntegrationAccountResultOutput) ToLookupIntegrationAccountResultOutput() LookupIntegrationAccountResultOutput {
	return o
}

func (o LookupIntegrationAccountResultOutput) ToLookupIntegrationAccountResultOutputWithContext(ctx context.Context) LookupIntegrationAccountResultOutput {
	return o
}

func (o LookupIntegrationAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIntegrationAccountResult] {
	return pulumix.Output[LookupIntegrationAccountResult]{
		OutputState: o.OutputState,
	}
}

// The resource id.
func (o LookupIntegrationAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The integration service environment.
func (o LookupIntegrationAccountResultOutput) IntegrationServiceEnvironment() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) *ResourceReferenceResponse {
		return v.IntegrationServiceEnvironment
	}).(ResourceReferenceResponsePtrOutput)
}

// The resource location.
func (o LookupIntegrationAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o LookupIntegrationAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The sku.
func (o LookupIntegrationAccountResultOutput) Sku() IntegrationAccountSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) *IntegrationAccountSkuResponse { return v.Sku }).(IntegrationAccountSkuResponsePtrOutput)
}

// The workflow state.
func (o LookupIntegrationAccountResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o LookupIntegrationAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o LookupIntegrationAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountResultOutput{})
}
