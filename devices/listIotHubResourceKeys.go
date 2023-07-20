// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the security metadata for an IoT hub. For more information, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-security.
// Azure REST API version: 2022-04-30-preview.
func ListIotHubResourceKeys(ctx *pulumi.Context, args *ListIotHubResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotHubResourceKeysResult, error) {
	var rv ListIotHubResourceKeysResult
	err := ctx.Invoke("azure-native:devices:listIotHubResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotHubResourceKeysArgs struct {
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The list of shared access policies with a next link.
type ListIotHubResourceKeysResult struct {
	// The next link.
	NextLink string `pulumi:"nextLink"`
	// The list of shared access policies.
	Value []SharedAccessSignatureAuthorizationRuleResponse `pulumi:"value"`
}

func ListIotHubResourceKeysOutput(ctx *pulumi.Context, args ListIotHubResourceKeysOutputArgs, opts ...pulumi.InvokeOption) ListIotHubResourceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIotHubResourceKeysResult, error) {
			args := v.(ListIotHubResourceKeysArgs)
			r, err := ListIotHubResourceKeys(ctx, &args, opts...)
			var s ListIotHubResourceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIotHubResourceKeysResultOutput)
}

type ListIotHubResourceKeysOutputArgs struct {
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListIotHubResourceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysArgs)(nil)).Elem()
}

// The list of shared access policies with a next link.
type ListIotHubResourceKeysResultOutput struct{ *pulumi.OutputState }

func (ListIotHubResourceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysResult)(nil)).Elem()
}

func (o ListIotHubResourceKeysResultOutput) ToListIotHubResourceKeysResultOutput() ListIotHubResourceKeysResultOutput {
	return o
}

func (o ListIotHubResourceKeysResultOutput) ToListIotHubResourceKeysResultOutputWithContext(ctx context.Context) ListIotHubResourceKeysResultOutput {
	return o
}

// The next link.
func (o ListIotHubResourceKeysResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// The list of shared access policies.
func (o ListIotHubResourceKeysResultOutput) Value() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysResult) []SharedAccessSignatureAuthorizationRuleResponse { return v.Value }).(SharedAccessSignatureAuthorizationRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotHubResourceKeysResultOutput{})
}
