// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List the primary and secondary keys for a provisioning service.
// Azure REST API version: 2022-12-12.
//
// Other available API versions: 2020-09-01-preview, 2023-03-01-preview.
func ListIotDpsResourceKeys(ctx *pulumi.Context, args *ListIotDpsResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotDpsResourceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListIotDpsResourceKeysResult
	err := ctx.Invoke("azure-native:devices:listIotDpsResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotDpsResourceKeysArgs struct {
	// The provisioning service name to get the shared access keys for.
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	// resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of shared access keys.
type ListIotDpsResourceKeysResult struct {
	// The next link.
	NextLink string `pulumi:"nextLink"`
	// The list of shared access policies.
	Value []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse `pulumi:"value"`
}

func ListIotDpsResourceKeysOutput(ctx *pulumi.Context, args ListIotDpsResourceKeysOutputArgs, opts ...pulumi.InvokeOption) ListIotDpsResourceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIotDpsResourceKeysResult, error) {
			args := v.(ListIotDpsResourceKeysArgs)
			r, err := ListIotDpsResourceKeys(ctx, &args, opts...)
			var s ListIotDpsResourceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIotDpsResourceKeysResultOutput)
}

type ListIotDpsResourceKeysOutputArgs struct {
	// The provisioning service name to get the shared access keys for.
	ProvisioningServiceName pulumi.StringInput `pulumi:"provisioningServiceName"`
	// resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListIotDpsResourceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotDpsResourceKeysArgs)(nil)).Elem()
}

// List of shared access keys.
type ListIotDpsResourceKeysResultOutput struct{ *pulumi.OutputState }

func (ListIotDpsResourceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotDpsResourceKeysResult)(nil)).Elem()
}

func (o ListIotDpsResourceKeysResultOutput) ToListIotDpsResourceKeysResultOutput() ListIotDpsResourceKeysResultOutput {
	return o
}

func (o ListIotDpsResourceKeysResultOutput) ToListIotDpsResourceKeysResultOutputWithContext(ctx context.Context) ListIotDpsResourceKeysResultOutput {
	return o
}

// The next link.
func (o ListIotDpsResourceKeysResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// The list of shared access policies.
func (o ListIotDpsResourceKeysResultOutput) Value() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysResult) []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse {
		return v.Value
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotDpsResourceKeysResultOutput{})
}
