// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the ISV access token for a SaaS resource.
// Azure REST API version: 2018-03-01-beta.
func ListSaasResourceAccessToken(ctx *pulumi.Context, args *ListSaasResourceAccessTokenArgs, opts ...pulumi.InvokeOption) (*ListSaasResourceAccessTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSaasResourceAccessTokenResult
	err := ctx.Invoke("azure-native:saas:listSaasResourceAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSaasResourceAccessTokenArgs struct {
	// The Saas resource ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
	ResourceId string `pulumi:"resourceId"`
}

// the ISV access token result response.
type ListSaasResourceAccessTokenResult struct {
	// The Publisher Offer Base Uri
	PublisherOfferBaseUri *string `pulumi:"publisherOfferBaseUri"`
	// The generated token
	Token *string `pulumi:"token"`
}

func ListSaasResourceAccessTokenOutput(ctx *pulumi.Context, args ListSaasResourceAccessTokenOutputArgs, opts ...pulumi.InvokeOption) ListSaasResourceAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSaasResourceAccessTokenResult, error) {
			args := v.(ListSaasResourceAccessTokenArgs)
			r, err := ListSaasResourceAccessToken(ctx, &args, opts...)
			var s ListSaasResourceAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSaasResourceAccessTokenResultOutput)
}

type ListSaasResourceAccessTokenOutputArgs struct {
	// The Saas resource ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
}

func (ListSaasResourceAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSaasResourceAccessTokenArgs)(nil)).Elem()
}

// the ISV access token result response.
type ListSaasResourceAccessTokenResultOutput struct{ *pulumi.OutputState }

func (ListSaasResourceAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSaasResourceAccessTokenResult)(nil)).Elem()
}

func (o ListSaasResourceAccessTokenResultOutput) ToListSaasResourceAccessTokenResultOutput() ListSaasResourceAccessTokenResultOutput {
	return o
}

func (o ListSaasResourceAccessTokenResultOutput) ToListSaasResourceAccessTokenResultOutputWithContext(ctx context.Context) ListSaasResourceAccessTokenResultOutput {
	return o
}

// The Publisher Offer Base Uri
func (o ListSaasResourceAccessTokenResultOutput) PublisherOfferBaseUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSaasResourceAccessTokenResult) *string { return v.PublisherOfferBaseUri }).(pulumi.StringPtrOutput)
}

// The generated token
func (o ListSaasResourceAccessTokenResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSaasResourceAccessTokenResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSaasResourceAccessTokenResultOutput{})
}
