// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Fetch User API Key from internal database, if it was generated and stored while creating the Elasticsearch Organization.
func GetOrganizationApiKey(ctx *pulumi.Context, args *GetOrganizationApiKeyArgs, opts ...pulumi.InvokeOption) (*GetOrganizationApiKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationApiKeyResult
	err := ctx.Invoke("azure-native:elastic/v20230701preview:getOrganizationApiKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetOrganizationApiKeyArgs struct {
	// The User email Id
	EmailId *string `pulumi:"emailId"`
}

// The User Api Key created for the Organization associated with the User Email Id that was passed in the request
type GetOrganizationApiKeyResult struct {
	Properties UserApiKeyResponsePropertiesResponse `pulumi:"properties"`
}

func GetOrganizationApiKeyOutput(ctx *pulumi.Context, args GetOrganizationApiKeyOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationApiKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrganizationApiKeyResult, error) {
			args := v.(GetOrganizationApiKeyArgs)
			r, err := GetOrganizationApiKey(ctx, &args, opts...)
			var s GetOrganizationApiKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrganizationApiKeyResultOutput)
}

type GetOrganizationApiKeyOutputArgs struct {
	// The User email Id
	EmailId pulumi.StringPtrInput `pulumi:"emailId"`
}

func (GetOrganizationApiKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationApiKeyArgs)(nil)).Elem()
}

// The User Api Key created for the Organization associated with the User Email Id that was passed in the request
type GetOrganizationApiKeyResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationApiKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationApiKeyResult)(nil)).Elem()
}

func (o GetOrganizationApiKeyResultOutput) ToGetOrganizationApiKeyResultOutput() GetOrganizationApiKeyResultOutput {
	return o
}

func (o GetOrganizationApiKeyResultOutput) ToGetOrganizationApiKeyResultOutputWithContext(ctx context.Context) GetOrganizationApiKeyResultOutput {
	return o
}

func (o GetOrganizationApiKeyResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetOrganizationApiKeyResult] {
	return pulumix.Output[GetOrganizationApiKeyResult]{
		OutputState: o.OutputState,
	}
}

func (o GetOrganizationApiKeyResultOutput) Properties() UserApiKeyResponsePropertiesResponseOutput {
	return o.ApplyT(func(v GetOrganizationApiKeyResult) UserApiKeyResponsePropertiesResponse { return v.Properties }).(UserApiKeyResponsePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationApiKeyResultOutput{})
}
