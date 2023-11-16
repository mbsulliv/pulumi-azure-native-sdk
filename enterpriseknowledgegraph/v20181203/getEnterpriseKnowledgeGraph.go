// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181203

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a EnterpriseKnowledgeGraph service specified by the parameters.
func LookupEnterpriseKnowledgeGraph(ctx *pulumi.Context, args *LookupEnterpriseKnowledgeGraphArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseKnowledgeGraphResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnterpriseKnowledgeGraphResult
	err := ctx.Invoke("azure-native:enterpriseknowledgegraph/v20181203:getEnterpriseKnowledgeGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseKnowledgeGraphArgs struct {
	// The name of the EnterpriseKnowledgeGraph resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the EnterpriseKnowledgeGraph resource.
	ResourceName string `pulumi:"resourceName"`
}

// EnterpriseKnowledgeGraph resource definition
type LookupEnterpriseKnowledgeGraphResult struct {
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties EnterpriseKnowledgeGraphPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}

func LookupEnterpriseKnowledgeGraphOutput(ctx *pulumi.Context, args LookupEnterpriseKnowledgeGraphOutputArgs, opts ...pulumi.InvokeOption) LookupEnterpriseKnowledgeGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnterpriseKnowledgeGraphResult, error) {
			args := v.(LookupEnterpriseKnowledgeGraphArgs)
			r, err := LookupEnterpriseKnowledgeGraph(ctx, &args, opts...)
			var s LookupEnterpriseKnowledgeGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnterpriseKnowledgeGraphResultOutput)
}

type LookupEnterpriseKnowledgeGraphOutputArgs struct {
	// The name of the EnterpriseKnowledgeGraph resource group in the user subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the EnterpriseKnowledgeGraph resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupEnterpriseKnowledgeGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseKnowledgeGraphArgs)(nil)).Elem()
}

// EnterpriseKnowledgeGraph resource definition
type LookupEnterpriseKnowledgeGraphResultOutput struct{ *pulumi.OutputState }

func (LookupEnterpriseKnowledgeGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseKnowledgeGraphResult)(nil)).Elem()
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) ToLookupEnterpriseKnowledgeGraphResultOutput() LookupEnterpriseKnowledgeGraphResultOutput {
	return o
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) ToLookupEnterpriseKnowledgeGraphResultOutputWithContext(ctx context.Context) LookupEnterpriseKnowledgeGraphResultOutput {
	return o
}

// Specifies the resource ID.
func (o LookupEnterpriseKnowledgeGraphResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the location of the resource.
func (o LookupEnterpriseKnowledgeGraphResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o LookupEnterpriseKnowledgeGraphResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to EnterpriseKnowledgeGraph resource
func (o LookupEnterpriseKnowledgeGraphResultOutput) Properties() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) EnterpriseKnowledgeGraphPropertiesResponse {
		return v.Properties
	}).(EnterpriseKnowledgeGraphPropertiesResponseOutput)
}

// Gets or sets the SKU of the resource.
func (o LookupEnterpriseKnowledgeGraphResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o LookupEnterpriseKnowledgeGraphResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o LookupEnterpriseKnowledgeGraphResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnterpriseKnowledgeGraphResultOutput{})
}
