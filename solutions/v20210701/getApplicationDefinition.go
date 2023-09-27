// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the managed application definition.
func LookupApplicationDefinition(ctx *pulumi.Context, args *LookupApplicationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationDefinitionResult
	err := ctx.Invoke("azure-native:solutions/v20210701:getApplicationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationDefinitionArgs struct {
	// The name of the managed application definition.
	ApplicationDefinitionName string `pulumi:"applicationDefinitionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about managed application definition.
type LookupApplicationDefinitionResult struct {
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts []ApplicationDefinitionArtifactResponse `pulumi:"artifacts"`
	// The managed application provider authorizations.
	Authorizations []ApplicationAuthorizationResponse `pulumi:"authorizations"`
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition interface{} `pulumi:"createUiDefinition"`
	// The managed application deployment policy.
	DeploymentPolicy *ApplicationDeploymentPolicyResponse `pulumi:"deploymentPolicy"`
	// The managed application definition description.
	Description *string `pulumi:"description"`
	// The managed application definition display name.
	DisplayName *string `pulumi:"displayName"`
	// Resource ID
	Id string `pulumi:"id"`
	// A value indicating whether the package is enabled or not.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Resource location
	Location *string `pulumi:"location"`
	// The managed application lock level.
	LockLevel string `pulumi:"lockLevel"`
	// The managed application locking policy.
	LockingPolicy *ApplicationPackageLockingPolicyDefinitionResponse `pulumi:"lockingPolicy"`
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate interface{} `pulumi:"mainTemplate"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The managed application management policy that determines publisher's access to the managed resource group.
	ManagementPolicy *ApplicationManagementPolicyResponse `pulumi:"managementPolicy"`
	// Resource name
	Name string `pulumi:"name"`
	// The managed application notification policy.
	NotificationPolicy *ApplicationNotificationPolicyResponse `pulumi:"notificationPolicy"`
	// The managed application definition package file Uri. Use this element
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The managed application provider policies.
	Policies []ApplicationPolicyResponse `pulumi:"policies"`
	// The SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// The storage account id for bring your own storage scenario.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupApplicationDefinitionOutput(ctx *pulumi.Context, args LookupApplicationDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationDefinitionResult, error) {
			args := v.(LookupApplicationDefinitionArgs)
			r, err := LookupApplicationDefinition(ctx, &args, opts...)
			var s LookupApplicationDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationDefinitionResultOutput)
}

type LookupApplicationDefinitionOutputArgs struct {
	// The name of the managed application definition.
	ApplicationDefinitionName pulumi.StringInput `pulumi:"applicationDefinitionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationDefinitionArgs)(nil)).Elem()
}

// Information about managed application definition.
type LookupApplicationDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationDefinitionResult)(nil)).Elem()
}

func (o LookupApplicationDefinitionResultOutput) ToLookupApplicationDefinitionResultOutput() LookupApplicationDefinitionResultOutput {
	return o
}

func (o LookupApplicationDefinitionResultOutput) ToLookupApplicationDefinitionResultOutputWithContext(ctx context.Context) LookupApplicationDefinitionResultOutput {
	return o
}

func (o LookupApplicationDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApplicationDefinitionResult] {
	return pulumix.Output[LookupApplicationDefinitionResult]{
		OutputState: o.OutputState,
	}
}

// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
func (o LookupApplicationDefinitionResultOutput) Artifacts() ApplicationDefinitionArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) []ApplicationDefinitionArtifactResponse { return v.Artifacts }).(ApplicationDefinitionArtifactResponseArrayOutput)
}

// The managed application provider authorizations.
func (o LookupApplicationDefinitionResultOutput) Authorizations() ApplicationAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) []ApplicationAuthorizationResponse { return v.Authorizations }).(ApplicationAuthorizationResponseArrayOutput)
}

// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
func (o LookupApplicationDefinitionResultOutput) CreateUiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) interface{} { return v.CreateUiDefinition }).(pulumi.AnyOutput)
}

// The managed application deployment policy.
func (o LookupApplicationDefinitionResultOutput) DeploymentPolicy() ApplicationDeploymentPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *ApplicationDeploymentPolicyResponse {
		return v.DeploymentPolicy
	}).(ApplicationDeploymentPolicyResponsePtrOutput)
}

// The managed application definition description.
func (o LookupApplicationDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The managed application definition display name.
func (o LookupApplicationDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource ID
func (o LookupApplicationDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// A value indicating whether the package is enabled or not.
func (o LookupApplicationDefinitionResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Resource location
func (o LookupApplicationDefinitionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The managed application lock level.
func (o LookupApplicationDefinitionResultOutput) LockLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.LockLevel }).(pulumi.StringOutput)
}

// The managed application locking policy.
func (o LookupApplicationDefinitionResultOutput) LockingPolicy() ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *ApplicationPackageLockingPolicyDefinitionResponse {
		return v.LockingPolicy
	}).(ApplicationPackageLockingPolicyDefinitionResponsePtrOutput)
}

// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
func (o LookupApplicationDefinitionResultOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) interface{} { return v.MainTemplate }).(pulumi.AnyOutput)
}

// ID of the resource that manages this resource.
func (o LookupApplicationDefinitionResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The managed application management policy that determines publisher's access to the managed resource group.
func (o LookupApplicationDefinitionResultOutput) ManagementPolicy() ApplicationManagementPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *ApplicationManagementPolicyResponse {
		return v.ManagementPolicy
	}).(ApplicationManagementPolicyResponsePtrOutput)
}

// Resource name
func (o LookupApplicationDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The managed application notification policy.
func (o LookupApplicationDefinitionResultOutput) NotificationPolicy() ApplicationNotificationPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *ApplicationNotificationPolicyResponse {
		return v.NotificationPolicy
	}).(ApplicationNotificationPolicyResponsePtrOutput)
}

// The managed application definition package file Uri. Use this element
func (o LookupApplicationDefinitionResultOutput) PackageFileUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.PackageFileUri }).(pulumi.StringPtrOutput)
}

// The managed application provider policies.
func (o LookupApplicationDefinitionResultOutput) Policies() ApplicationPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) []ApplicationPolicyResponse { return v.Policies }).(ApplicationPolicyResponseArrayOutput)
}

// The SKU of the resource.
func (o LookupApplicationDefinitionResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The storage account id for bring your own storage scenario.
func (o LookupApplicationDefinitionResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApplicationDefinitionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupApplicationDefinitionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupApplicationDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationDefinitionResultOutput{})
}
