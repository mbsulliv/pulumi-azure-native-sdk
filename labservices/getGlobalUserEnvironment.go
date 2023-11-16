// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the virtual machine details
// Azure REST API version: 2018-10-15.
func GetGlobalUserEnvironment(ctx *pulumi.Context, args *GetGlobalUserEnvironmentArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserEnvironmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetGlobalUserEnvironmentResult
	err := ctx.Invoke("azure-native:labservices:getGlobalUserEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserEnvironmentArgs struct {
	// The resourceId of the environment
	EnvironmentId string `pulumi:"environmentId"`
	// Specify the $expand query. Example: 'properties($expand=environment)'
	Expand *string `pulumi:"expand"`
	// The name of the user.
	UserName string `pulumi:"userName"`
}

// Represents the environments details
type GetGlobalUserEnvironmentResult struct {
	// Details of the environment
	Environment EnvironmentDetailsResponse `pulumi:"environment"`
}

func GetGlobalUserEnvironmentOutput(ctx *pulumi.Context, args GetGlobalUserEnvironmentOutputArgs, opts ...pulumi.InvokeOption) GetGlobalUserEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGlobalUserEnvironmentResult, error) {
			args := v.(GetGlobalUserEnvironmentArgs)
			r, err := GetGlobalUserEnvironment(ctx, &args, opts...)
			var s GetGlobalUserEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGlobalUserEnvironmentResultOutput)
}

type GetGlobalUserEnvironmentOutputArgs struct {
	// The resourceId of the environment
	EnvironmentId pulumi.StringInput `pulumi:"environmentId"`
	// Specify the $expand query. Example: 'properties($expand=environment)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the user.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (GetGlobalUserEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserEnvironmentArgs)(nil)).Elem()
}

// Represents the environments details
type GetGlobalUserEnvironmentResultOutput struct{ *pulumi.OutputState }

func (GetGlobalUserEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserEnvironmentResult)(nil)).Elem()
}

func (o GetGlobalUserEnvironmentResultOutput) ToGetGlobalUserEnvironmentResultOutput() GetGlobalUserEnvironmentResultOutput {
	return o
}

func (o GetGlobalUserEnvironmentResultOutput) ToGetGlobalUserEnvironmentResultOutputWithContext(ctx context.Context) GetGlobalUserEnvironmentResultOutput {
	return o
}

// Details of the environment
func (o GetGlobalUserEnvironmentResultOutput) Environment() EnvironmentDetailsResponseOutput {
	return o.ApplyT(func(v GetGlobalUserEnvironmentResult) EnvironmentDetailsResponse { return v.Environment }).(EnvironmentDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGlobalUserEnvironmentResultOutput{})
}
