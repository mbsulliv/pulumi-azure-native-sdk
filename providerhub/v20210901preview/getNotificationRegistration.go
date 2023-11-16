// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the notification registration details.
func LookupNotificationRegistration(ctx *pulumi.Context, args *LookupNotificationRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupNotificationRegistrationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationRegistrationResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getNotificationRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationRegistrationArgs struct {
	// The notification registration.
	NotificationRegistrationName string `pulumi:"notificationRegistrationName"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
}

// The notification registration definition.
type LookupNotificationRegistrationResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name       string                                     `pulumi:"name"`
	Properties NotificationRegistrationResponseProperties `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNotificationRegistrationOutput(ctx *pulumi.Context, args LookupNotificationRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationRegistrationResult, error) {
			args := v.(LookupNotificationRegistrationArgs)
			r, err := LookupNotificationRegistration(ctx, &args, opts...)
			var s LookupNotificationRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationRegistrationResultOutput)
}

type LookupNotificationRegistrationOutputArgs struct {
	// The notification registration.
	NotificationRegistrationName pulumi.StringInput `pulumi:"notificationRegistrationName"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
}

func (LookupNotificationRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationRegistrationArgs)(nil)).Elem()
}

// The notification registration definition.
type LookupNotificationRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationRegistrationResult)(nil)).Elem()
}

func (o LookupNotificationRegistrationResultOutput) ToLookupNotificationRegistrationResultOutput() LookupNotificationRegistrationResultOutput {
	return o
}

func (o LookupNotificationRegistrationResultOutput) ToLookupNotificationRegistrationResultOutputWithContext(ctx context.Context) LookupNotificationRegistrationResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNotificationRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNotificationRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNotificationRegistrationResultOutput) Properties() NotificationRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) NotificationRegistrationResponseProperties {
		return v.Properties
	}).(NotificationRegistrationResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupNotificationRegistrationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNotificationRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationRegistrationResultOutput{})
}
