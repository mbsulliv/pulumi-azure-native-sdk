// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List keys of APM sensitive properties.
func ListApmSecretKeys(ctx *pulumi.Context, args *ListApmSecretKeysArgs, opts ...pulumi.InvokeOption) (*ListApmSecretKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListApmSecretKeysResult
	err := ctx.Invoke("azure-native:appplatform/v20230901preview:listApmSecretKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApmSecretKeysArgs struct {
	// The name of the APM
	ApmName string `pulumi:"apmName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Keys of APM sensitive properties
type ListApmSecretKeysResult struct {
	// Collection of the keys for the APM sensitive properties
	Value []string `pulumi:"value"`
}

func ListApmSecretKeysOutput(ctx *pulumi.Context, args ListApmSecretKeysOutputArgs, opts ...pulumi.InvokeOption) ListApmSecretKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApmSecretKeysResult, error) {
			args := v.(ListApmSecretKeysArgs)
			r, err := ListApmSecretKeys(ctx, &args, opts...)
			var s ListApmSecretKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApmSecretKeysResultOutput)
}

type ListApmSecretKeysOutputArgs struct {
	// The name of the APM
	ApmName pulumi.StringInput `pulumi:"apmName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListApmSecretKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApmSecretKeysArgs)(nil)).Elem()
}

// Keys of APM sensitive properties
type ListApmSecretKeysResultOutput struct{ *pulumi.OutputState }

func (ListApmSecretKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApmSecretKeysResult)(nil)).Elem()
}

func (o ListApmSecretKeysResultOutput) ToListApmSecretKeysResultOutput() ListApmSecretKeysResultOutput {
	return o
}

func (o ListApmSecretKeysResultOutput) ToListApmSecretKeysResultOutputWithContext(ctx context.Context) ListApmSecretKeysResultOutput {
	return o
}

func (o ListApmSecretKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListApmSecretKeysResult] {
	return pulumix.Output[ListApmSecretKeysResult]{
		OutputState: o.OutputState,
	}
}

// Collection of the keys for the APM sensitive properties
func (o ListApmSecretKeysResultOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListApmSecretKeysResult) []string { return v.Value }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApmSecretKeysResultOutput{})
}
