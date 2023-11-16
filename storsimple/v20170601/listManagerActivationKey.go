// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the activation key of the manager.
func ListManagerActivationKey(ctx *pulumi.Context, args *ListManagerActivationKeyArgs, opts ...pulumi.InvokeOption) (*ListManagerActivationKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagerActivationKeyResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:listManagerActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagerActivationKeyArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The key.
type ListManagerActivationKeyResult struct {
	// The activation key for the device.
	ActivationKey string `pulumi:"activationKey"`
}

func ListManagerActivationKeyOutput(ctx *pulumi.Context, args ListManagerActivationKeyOutputArgs, opts ...pulumi.InvokeOption) ListManagerActivationKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagerActivationKeyResult, error) {
			args := v.(ListManagerActivationKeyArgs)
			r, err := ListManagerActivationKey(ctx, &args, opts...)
			var s ListManagerActivationKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagerActivationKeyResultOutput)
}

type ListManagerActivationKeyOutputArgs struct {
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListManagerActivationKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerActivationKeyArgs)(nil)).Elem()
}

// The key.
type ListManagerActivationKeyResultOutput struct{ *pulumi.OutputState }

func (ListManagerActivationKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerActivationKeyResult)(nil)).Elem()
}

func (o ListManagerActivationKeyResultOutput) ToListManagerActivationKeyResultOutput() ListManagerActivationKeyResultOutput {
	return o
}

func (o ListManagerActivationKeyResultOutput) ToListManagerActivationKeyResultOutputWithContext(ctx context.Context) ListManagerActivationKeyResultOutput {
	return o
}

// The activation key for the device.
func (o ListManagerActivationKeyResultOutput) ActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagerActivationKeyResult) string { return v.ActivationKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagerActivationKeyResultOutput{})
}
