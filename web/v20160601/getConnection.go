// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a specific connection
func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectionResult
	err := ctx.Invoke("azure-native:web/v20160601:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	// Connection name
	ConnectionName string `pulumi:"connectionName"`
	// The resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// API connection
type LookupConnectionResult struct {
	// Resource ETag
	Etag *string `pulumi:"etag"`
	// Resource id
	Id string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name       string                                    `pulumi:"name"`
	Properties ApiConnectionDefinitionResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupConnectionOutput(ctx *pulumi.Context, args LookupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionResult, error) {
			args := v.(LookupConnectionArgs)
			r, err := LookupConnection(ctx, &args, opts...)
			var s LookupConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionResultOutput)
}

type LookupConnectionOutputArgs struct {
	// Connection name
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}

// API connection
type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectionResult] {
	return pulumix.Output[LookupConnectionResult]{
		OutputState: o.OutputState,
	}
}

// Resource ETag
func (o LookupConnectionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource id
func (o LookupConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) Properties() ApiConnectionDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectionResult) ApiConnectionDefinitionResponseProperties { return v.Properties }).(ApiConnectionDefinitionResponsePropertiesOutput)
}

// Resource tags
func (o LookupConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
