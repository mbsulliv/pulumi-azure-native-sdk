// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230630

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a shared access policy by name from an IoT hub. For more information, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-security.
func ListIotHubResourceKeysForKeyName(ctx *pulumi.Context, args *ListIotHubResourceKeysForKeyNameArgs, opts ...pulumi.InvokeOption) (*ListIotHubResourceKeysForKeyNameResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListIotHubResourceKeysForKeyNameResult
	err := ctx.Invoke("azure-native:devices/v20230630:listIotHubResourceKeysForKeyName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotHubResourceKeysForKeyNameArgs struct {
	// The name of the shared access policy.
	KeyName string `pulumi:"keyName"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The properties of an IoT hub shared access policy.
type ListIotHubResourceKeysForKeyNameResult struct {
	// The name of the shared access policy.
	KeyName string `pulumi:"keyName"`
	// The primary key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The permissions assigned to the shared access policy.
	Rights string `pulumi:"rights"`
	// The secondary key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListIotHubResourceKeysForKeyNameOutput(ctx *pulumi.Context, args ListIotHubResourceKeysForKeyNameOutputArgs, opts ...pulumi.InvokeOption) ListIotHubResourceKeysForKeyNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIotHubResourceKeysForKeyNameResult, error) {
			args := v.(ListIotHubResourceKeysForKeyNameArgs)
			r, err := ListIotHubResourceKeysForKeyName(ctx, &args, opts...)
			var s ListIotHubResourceKeysForKeyNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIotHubResourceKeysForKeyNameResultOutput)
}

type ListIotHubResourceKeysForKeyNameOutputArgs struct {
	// The name of the shared access policy.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListIotHubResourceKeysForKeyNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysForKeyNameArgs)(nil)).Elem()
}

// The properties of an IoT hub shared access policy.
type ListIotHubResourceKeysForKeyNameResultOutput struct{ *pulumi.OutputState }

func (ListIotHubResourceKeysForKeyNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysForKeyNameResult)(nil)).Elem()
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) ToListIotHubResourceKeysForKeyNameResultOutput() ListIotHubResourceKeysForKeyNameResultOutput {
	return o
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) ToListIotHubResourceKeysForKeyNameResultOutputWithContext(ctx context.Context) ListIotHubResourceKeysForKeyNameResultOutput {
	return o
}

// The name of the shared access policy.
func (o ListIotHubResourceKeysForKeyNameResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// The primary key.
func (o ListIotHubResourceKeysForKeyNameResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// The permissions assigned to the shared access policy.
func (o ListIotHubResourceKeysForKeyNameResultOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) string { return v.Rights }).(pulumi.StringOutput)
}

// The secondary key.
func (o ListIotHubResourceKeysForKeyNameResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotHubResourceKeysForKeyNameResultOutput{})
}
