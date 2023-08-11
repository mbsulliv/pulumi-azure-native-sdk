// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This method gets the unencrypted secrets related to the job.
func ListJobCredentials(ctx *pulumi.Context, args *ListJobCredentialsArgs, opts ...pulumi.InvokeOption) (*ListJobCredentialsResult, error) {
	var rv ListJobCredentialsResult
	err := ctx.Invoke("azure-native:databox/v20230301:listJobCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListJobCredentialsArgs struct {
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	JobName string `pulumi:"jobName"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of unencrypted credentials for accessing device.
type ListJobCredentialsResult struct {
	// Link for the next set of unencrypted credentials.
	NextLink *string `pulumi:"nextLink"`
	// List of unencrypted credentials.
	Value []UnencryptedCredentialsResponse `pulumi:"value"`
}

func ListJobCredentialsOutput(ctx *pulumi.Context, args ListJobCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListJobCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListJobCredentialsResult, error) {
			args := v.(ListJobCredentialsArgs)
			r, err := ListJobCredentials(ctx, &args, opts...)
			var s ListJobCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListJobCredentialsResultOutput)
}

type ListJobCredentialsOutputArgs struct {
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListJobCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobCredentialsArgs)(nil)).Elem()
}

// List of unencrypted credentials for accessing device.
type ListJobCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListJobCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobCredentialsResult)(nil)).Elem()
}

func (o ListJobCredentialsResultOutput) ToListJobCredentialsResultOutput() ListJobCredentialsResultOutput {
	return o
}

func (o ListJobCredentialsResultOutput) ToListJobCredentialsResultOutputWithContext(ctx context.Context) ListJobCredentialsResultOutput {
	return o
}

// Link for the next set of unencrypted credentials.
func (o ListJobCredentialsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListJobCredentialsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of unencrypted credentials.
func (o ListJobCredentialsResultOutput) Value() UnencryptedCredentialsResponseArrayOutput {
	return o.ApplyT(func(v ListJobCredentialsResult) []UnencryptedCredentialsResponse { return v.Value }).(UnencryptedCredentialsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListJobCredentialsResultOutput{})
}
