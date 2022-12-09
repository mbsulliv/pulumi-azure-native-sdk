// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Domain service.
//
// Deprecated: Version 2020-01-01 will be removed in v2 of the provider.
type DomainService struct {
	pulumi.CustomResourceState

	// Deployment Id
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// Domain Configuration Type
	DomainConfigurationType pulumi.StringPtrOutput `pulumi:"domainConfigurationType"`
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings DomainSecuritySettingsResponsePtrOutput `pulumi:"domainSecuritySettings"`
	// Resource etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync pulumi.StringPtrOutput `pulumi:"filteredSync"`
	// Secure LDAP Settings
	LdapsSettings LdapsSettingsResponsePtrOutput `pulumi:"ldapsSettings"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Migration Properties
	MigrationProperties MigrationPropertiesResponseOutput `pulumi:"migrationProperties"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification Settings
	NotificationSettings NotificationSettingsResponsePtrOutput `pulumi:"notificationSettings"`
	// the current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of ReplicaSets
	ReplicaSets ReplicaSetResponseArrayOutput `pulumi:"replicaSets"`
	// Resource Forest Settings
	ResourceForestSettings ResourceForestSettingsResponsePtrOutput `pulumi:"resourceForestSettings"`
	// Sku Type
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// SyncOwner ReplicaSet Id
	SyncOwner pulumi.StringOutput `pulumi:"syncOwner"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure Active Directory Tenant Id
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Data Model Version
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewDomainService registers a new resource with the given unique name, arguments, and options.
func NewDomainService(ctx *pulumi.Context,
	name string, args *DomainServiceArgs, opts ...pulumi.ResourceOption) (*DomainService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DomainSecuritySettings != nil {
		args.DomainSecuritySettings = args.DomainSecuritySettings.ToDomainSecuritySettingsPtrOutput().ApplyT(func(v *DomainSecuritySettings) *DomainSecuritySettings { return v.Defaults() }).(DomainSecuritySettingsPtrOutput)
	}
	if args.LdapsSettings != nil {
		args.LdapsSettings = args.LdapsSettings.ToLdapsSettingsPtrOutput().ApplyT(func(v *LdapsSettings) *LdapsSettings { return v.Defaults() }).(LdapsSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aad:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170101:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170601:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210301:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210501:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20220901:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20221201:DomainService"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainService
	err := ctx.RegisterResource("azure-native:aad/v20200101:DomainService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainService gets an existing DomainService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainServiceState, opts ...pulumi.ResourceOption) (*DomainService, error) {
	var resource DomainService
	err := ctx.ReadResource("azure-native:aad/v20200101:DomainService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainService resources.
type domainServiceState struct {
}

type DomainServiceState struct {
}

func (DomainServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceState)(nil)).Elem()
}

type domainServiceArgs struct {
	// Domain Configuration Type
	DomainConfigurationType *string `pulumi:"domainConfigurationType"`
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettings `pulumi:"domainSecuritySettings"`
	// The name of the domain service.
	DomainServiceName *string `pulumi:"domainServiceName"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync *string `pulumi:"filteredSync"`
	// Secure LDAP Settings
	LdapsSettings *LdapsSettings `pulumi:"ldapsSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// Notification Settings
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	// List of ReplicaSets
	ReplicaSets []ReplicaSet `pulumi:"replicaSets"`
	// Resource Forest Settings
	ResourceForestSettings *ResourceForestSettings `pulumi:"resourceForestSettings"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku Type
	Sku *string `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DomainService resource.
type DomainServiceArgs struct {
	// Domain Configuration Type
	DomainConfigurationType pulumi.StringPtrInput
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName pulumi.StringPtrInput
	// DomainSecurity Settings
	DomainSecuritySettings DomainSecuritySettingsPtrInput
	// The name of the domain service.
	DomainServiceName pulumi.StringPtrInput
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync pulumi.StringPtrInput
	// Secure LDAP Settings
	LdapsSettings LdapsSettingsPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Notification Settings
	NotificationSettings NotificationSettingsPtrInput
	// List of ReplicaSets
	ReplicaSets ReplicaSetArrayInput
	// Resource Forest Settings
	ResourceForestSettings ResourceForestSettingsPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sku Type
	Sku pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DomainServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceArgs)(nil)).Elem()
}

type DomainServiceInput interface {
	pulumi.Input

	ToDomainServiceOutput() DomainServiceOutput
	ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput
}

func (*DomainService) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainService)(nil)).Elem()
}

func (i *DomainService) ToDomainServiceOutput() DomainServiceOutput {
	return i.ToDomainServiceOutputWithContext(context.Background())
}

func (i *DomainService) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainServiceOutput)
}

type DomainServiceOutput struct{ *pulumi.OutputState }

func (DomainServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainService)(nil)).Elem()
}

func (o DomainServiceOutput) ToDomainServiceOutput() DomainServiceOutput {
	return o
}

func (o DomainServiceOutput) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return o
}

// Deployment Id
func (o DomainServiceOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// Domain Configuration Type
func (o DomainServiceOutput) DomainConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.DomainConfigurationType }).(pulumi.StringPtrOutput)
}

// The name of the Azure domain that the user would like to deploy Domain Services to.
func (o DomainServiceOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// DomainSecurity Settings
func (o DomainServiceOutput) DomainSecuritySettings() DomainSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) DomainSecuritySettingsResponsePtrOutput { return v.DomainSecuritySettings }).(DomainSecuritySettingsResponsePtrOutput)
}

// Resource etag
func (o DomainServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Enabled or Disabled flag to turn on Group-based filtered sync
func (o DomainServiceOutput) FilteredSync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.FilteredSync }).(pulumi.StringPtrOutput)
}

// Secure LDAP Settings
func (o DomainServiceOutput) LdapsSettings() LdapsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) LdapsSettingsResponsePtrOutput { return v.LdapsSettings }).(LdapsSettingsResponsePtrOutput)
}

// Resource location
func (o DomainServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Migration Properties
func (o DomainServiceOutput) MigrationProperties() MigrationPropertiesResponseOutput {
	return o.ApplyT(func(v *DomainService) MigrationPropertiesResponseOutput { return v.MigrationProperties }).(MigrationPropertiesResponseOutput)
}

// Resource name
func (o DomainServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Notification Settings
func (o DomainServiceOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) NotificationSettingsResponsePtrOutput { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

// the current deployment or provisioning state, which only appears in the response.
func (o DomainServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of ReplicaSets
func (o DomainServiceOutput) ReplicaSets() ReplicaSetResponseArrayOutput {
	return o.ApplyT(func(v *DomainService) ReplicaSetResponseArrayOutput { return v.ReplicaSets }).(ReplicaSetResponseArrayOutput)
}

// Resource Forest Settings
func (o DomainServiceOutput) ResourceForestSettings() ResourceForestSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) ResourceForestSettingsResponsePtrOutput { return v.ResourceForestSettings }).(ResourceForestSettingsResponsePtrOutput)
}

// Sku Type
func (o DomainServiceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

// SyncOwner ReplicaSet Id
func (o DomainServiceOutput) SyncOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.SyncOwner }).(pulumi.StringOutput)
}

// Resource tags
func (o DomainServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure Active Directory Tenant Id
func (o DomainServiceOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type
func (o DomainServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Data Model Version
func (o DomainServiceOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *DomainService) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainServiceOutput{})
}
