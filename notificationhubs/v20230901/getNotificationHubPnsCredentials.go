// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description of a NotificationHub PNS Credentials. This is a response of the POST requests that return namespace or hubs
// PNS credentials.
func GetNotificationHubPnsCredentials(ctx *pulumi.Context, args *GetNotificationHubPnsCredentialsArgs, opts ...pulumi.InvokeOption) (*GetNotificationHubPnsCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetNotificationHubPnsCredentialsResult
	err := ctx.Invoke("azure-native:notificationhubs/v20230901:getNotificationHubPnsCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetNotificationHubPnsCredentialsArgs struct {
	// Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Notification Hub name
	NotificationHubName string `pulumi:"notificationHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a NotificationHub PNS Credentials. This is a response of the POST requests that return namespace or hubs
// PNS credentials.
type GetNotificationHubPnsCredentialsResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Deprecated - only for compatibility.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Collection of Notification Hub or Notification Hub Namespace PNS credentials.
	Properties PnsCredentialsResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Deprecated - only for compatibility.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func GetNotificationHubPnsCredentialsOutput(ctx *pulumi.Context, args GetNotificationHubPnsCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetNotificationHubPnsCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNotificationHubPnsCredentialsResult, error) {
			args := v.(GetNotificationHubPnsCredentialsArgs)
			r, err := GetNotificationHubPnsCredentials(ctx, &args, opts...)
			var s GetNotificationHubPnsCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNotificationHubPnsCredentialsResultOutput)
}

type GetNotificationHubPnsCredentialsOutputArgs struct {
	// Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Notification Hub name
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetNotificationHubPnsCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationHubPnsCredentialsArgs)(nil)).Elem()
}

// Description of a NotificationHub PNS Credentials. This is a response of the POST requests that return namespace or hubs
// PNS credentials.
type GetNotificationHubPnsCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetNotificationHubPnsCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationHubPnsCredentialsResult)(nil)).Elem()
}

func (o GetNotificationHubPnsCredentialsResultOutput) ToGetNotificationHubPnsCredentialsResultOutput() GetNotificationHubPnsCredentialsResultOutput {
	return o
}

func (o GetNotificationHubPnsCredentialsResultOutput) ToGetNotificationHubPnsCredentialsResultOutputWithContext(ctx context.Context) GetNotificationHubPnsCredentialsResultOutput {
	return o
}

func (o GetNotificationHubPnsCredentialsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetNotificationHubPnsCredentialsResult] {
	return pulumix.Output[GetNotificationHubPnsCredentialsResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o GetNotificationHubPnsCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Deprecated - only for compatibility.
func (o GetNotificationHubPnsCredentialsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o GetNotificationHubPnsCredentialsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Collection of Notification Hub or Notification Hub Namespace PNS credentials.
func (o GetNotificationHubPnsCredentialsResultOutput) Properties() PnsCredentialsResponseOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) PnsCredentialsResponse { return v.Properties }).(PnsCredentialsResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GetNotificationHubPnsCredentialsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Deprecated - only for compatibility.
func (o GetNotificationHubPnsCredentialsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GetNotificationHubPnsCredentialsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNotificationHubPnsCredentialsResultOutput{})
}
