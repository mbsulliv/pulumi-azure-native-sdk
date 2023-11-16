// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Notification Hub Resource.
func LookupNotificationHub(ctx *pulumi.Context, args *LookupNotificationHubArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationHubResult
	err := ctx.Invoke("azure-native:notificationhubs/v20230901:getNotificationHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubArgs struct {
	// Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Notification Hub name
	NotificationHubName string `pulumi:"notificationHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Notification Hub Resource.
type LookupNotificationHubResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// NotificationHub properties.
	Properties NotificationHubPropertiesResponse `pulumi:"properties"`
	// The Sku description for a namespace
	Sku *SkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNotificationHubOutput(ctx *pulumi.Context, args LookupNotificationHubOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationHubResult, error) {
			args := v.(LookupNotificationHubArgs)
			r, err := LookupNotificationHub(ctx, &args, opts...)
			var s LookupNotificationHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationHubResultOutput)
}

type LookupNotificationHubOutputArgs struct {
	// Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Notification Hub name
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubArgs)(nil)).Elem()
}

// Notification Hub Resource.
type LookupNotificationHubResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubResult)(nil)).Elem()
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutput() LookupNotificationHubResultOutput {
	return o
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutputWithContext(ctx context.Context) LookupNotificationHubResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNotificationHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNotificationHubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNotificationHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Name }).(pulumi.StringOutput)
}

// NotificationHub properties.
func (o LookupNotificationHubResultOutput) Properties() NotificationHubPropertiesResponseOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) NotificationHubPropertiesResponse { return v.Properties }).(NotificationHubPropertiesResponseOutput)
}

// The Sku description for a namespace
func (o LookupNotificationHubResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNotificationHubResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNotificationHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNotificationHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationHubResultOutput{})
}
