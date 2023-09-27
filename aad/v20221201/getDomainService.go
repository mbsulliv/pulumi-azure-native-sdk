// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Get Domain Service operation retrieves a json representation of the Domain Service.
func LookupDomainService(ctx *pulumi.Context, args *LookupDomainServiceArgs, opts ...pulumi.InvokeOption) (*LookupDomainServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainServiceResult
	err := ctx.Invoke("azure-native:aad/v20221201:getDomainService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainServiceArgs struct {
	// The name of the domain service.
	DomainServiceName string `pulumi:"domainServiceName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Domain service.
type LookupDomainServiceResult struct {
	// Configuration diagnostics data containing latest execution from client.
	ConfigDiagnostics *ConfigDiagnosticsResponse `pulumi:"configDiagnostics"`
	// Deployment Id
	DeploymentId string `pulumi:"deploymentId"`
	// Domain Configuration Type
	DomainConfigurationType *string `pulumi:"domainConfigurationType"`
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettingsResponse `pulumi:"domainSecuritySettings"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync *string `pulumi:"filteredSync"`
	// Resource Id
	Id string `pulumi:"id"`
	// Secure LDAP Settings
	LdapsSettings *LdapsSettingsResponse `pulumi:"ldapsSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// Migration Properties
	MigrationProperties MigrationPropertiesResponse `pulumi:"migrationProperties"`
	// Resource name
	Name string `pulumi:"name"`
	// Notification Settings
	NotificationSettings *NotificationSettingsResponse `pulumi:"notificationSettings"`
	// the current deployment or provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of ReplicaSets
	ReplicaSets []ReplicaSetResponse `pulumi:"replicaSets"`
	// Resource Forest Settings
	ResourceForestSettings *ResourceForestSettingsResponse `pulumi:"resourceForestSettings"`
	// Sku Type
	Sku *string `pulumi:"sku"`
	// The unique sync application id of the Azure AD Domain Services deployment.
	SyncApplicationId string `pulumi:"syncApplicationId"`
	// SyncOwner ReplicaSet Id
	SyncOwner string `pulumi:"syncOwner"`
	// All or CloudOnly, All users in AAD are synced to AAD DS domain or only users actively syncing in the cloud
	SyncScope *string `pulumi:"syncScope"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure Active Directory Tenant Id
	TenantId string `pulumi:"tenantId"`
	// Resource type
	Type string `pulumi:"type"`
	// Data Model Version
	Version int `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupDomainServiceResult
func (val *LookupDomainServiceResult) Defaults() *LookupDomainServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DomainSecuritySettings = tmp.DomainSecuritySettings.Defaults()

	tmp.LdapsSettings = tmp.LdapsSettings.Defaults()

	if tmp.SyncScope == nil {
		syncScope_ := "All"
		tmp.SyncScope = &syncScope_
	}
	return &tmp
}

func LookupDomainServiceOutput(ctx *pulumi.Context, args LookupDomainServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDomainServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainServiceResult, error) {
			args := v.(LookupDomainServiceArgs)
			r, err := LookupDomainService(ctx, &args, opts...)
			var s LookupDomainServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainServiceResultOutput)
}

type LookupDomainServiceOutputArgs struct {
	// The name of the domain service.
	DomainServiceName pulumi.StringInput `pulumi:"domainServiceName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainServiceArgs)(nil)).Elem()
}

// Domain service.
type LookupDomainServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDomainServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainServiceResult)(nil)).Elem()
}

func (o LookupDomainServiceResultOutput) ToLookupDomainServiceResultOutput() LookupDomainServiceResultOutput {
	return o
}

func (o LookupDomainServiceResultOutput) ToLookupDomainServiceResultOutputWithContext(ctx context.Context) LookupDomainServiceResultOutput {
	return o
}

func (o LookupDomainServiceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDomainServiceResult] {
	return pulumix.Output[LookupDomainServiceResult]{
		OutputState: o.OutputState,
	}
}

// Configuration diagnostics data containing latest execution from client.
func (o LookupDomainServiceResultOutput) ConfigDiagnostics() ConfigDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *ConfigDiagnosticsResponse { return v.ConfigDiagnostics }).(ConfigDiagnosticsResponsePtrOutput)
}

// Deployment Id
func (o LookupDomainServiceResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

// Domain Configuration Type
func (o LookupDomainServiceResultOutput) DomainConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.DomainConfigurationType }).(pulumi.StringPtrOutput)
}

// The name of the Azure domain that the user would like to deploy Domain Services to.
func (o LookupDomainServiceResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

// DomainSecurity Settings
func (o LookupDomainServiceResultOutput) DomainSecuritySettings() DomainSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *DomainSecuritySettingsResponse { return v.DomainSecuritySettings }).(DomainSecuritySettingsResponsePtrOutput)
}

// Resource etag
func (o LookupDomainServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Enabled or Disabled flag to turn on Group-based filtered sync
func (o LookupDomainServiceResultOutput) FilteredSync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.FilteredSync }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupDomainServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Secure LDAP Settings
func (o LookupDomainServiceResultOutput) LdapsSettings() LdapsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *LdapsSettingsResponse { return v.LdapsSettings }).(LdapsSettingsResponsePtrOutput)
}

// Resource location
func (o LookupDomainServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Migration Properties
func (o LookupDomainServiceResultOutput) MigrationProperties() MigrationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) MigrationPropertiesResponse { return v.MigrationProperties }).(MigrationPropertiesResponseOutput)
}

// Resource name
func (o LookupDomainServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Notification Settings
func (o LookupDomainServiceResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

// the current deployment or provisioning state, which only appears in the response.
func (o LookupDomainServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of ReplicaSets
func (o LookupDomainServiceResultOutput) ReplicaSets() ReplicaSetResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) []ReplicaSetResponse { return v.ReplicaSets }).(ReplicaSetResponseArrayOutput)
}

// Resource Forest Settings
func (o LookupDomainServiceResultOutput) ResourceForestSettings() ResourceForestSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *ResourceForestSettingsResponse { return v.ResourceForestSettings }).(ResourceForestSettingsResponsePtrOutput)
}

// Sku Type
func (o LookupDomainServiceResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// The unique sync application id of the Azure AD Domain Services deployment.
func (o LookupDomainServiceResultOutput) SyncApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.SyncApplicationId }).(pulumi.StringOutput)
}

// SyncOwner ReplicaSet Id
func (o LookupDomainServiceResultOutput) SyncOwner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.SyncOwner }).(pulumi.StringOutput)
}

// All or CloudOnly, All users in AAD are synced to AAD DS domain or only users actively syncing in the cloud
func (o LookupDomainServiceResultOutput) SyncScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.SyncScope }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o LookupDomainServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupDomainServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure Active Directory Tenant Id
func (o LookupDomainServiceResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type
func (o LookupDomainServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Data Model Version
func (o LookupDomainServiceResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainServiceResultOutput{})
}
