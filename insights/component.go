// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Application Insights component definition.
// Azure REST API version: 2020-02-02. Prior API version in Azure Native 1.x: 2015-05-01.
//
// Other available API versions: 2020-02-02-preview.
type Component struct {
	pulumi.CustomResourceState

	// Application Insights Unique ID for your Application.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The unique ID of your application. This field mirrors the 'Name' field and cannot be changed.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Type of application being monitored.
	ApplicationType pulumi.StringOutput `pulumi:"applicationType"`
	// Application Insights component connection string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Creation Date for the Application Insights component, in ISO 8601 format.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Disable IP masking.
	DisableIpMasking pulumi.BoolPtrOutput `pulumi:"disableIpMasking"`
	// Disable Non-AAD based Auth.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// Resource etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
	FlowType pulumi.StringPtrOutput `pulumi:"flowType"`
	// Force users to create their own storage account for profiler and debugger.
	ForceCustomerStorageForProfiler pulumi.BoolPtrOutput `pulumi:"forceCustomerStorageForProfiler"`
	// The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
	HockeyAppId pulumi.StringPtrOutput `pulumi:"hockeyAppId"`
	// Token used to authenticate communications with between Application Insights and HockeyApp.
	HockeyAppToken pulumi.StringOutput `pulumi:"hockeyAppToken"`
	// Purge data immediately after 30 days.
	ImmediatePurgeDataOn30Days pulumi.BoolPtrOutput `pulumi:"immediatePurgeDataOn30Days"`
	// Indicates the flow of the ingestion.
	IngestionMode pulumi.StringPtrOutput `pulumi:"ingestionMode"`
	// Application Insights Instrumentation key. A read-only value that applications can use to identify the destination for all telemetry sent to Azure Application Insights. This value will be supplied upon construction of each new Application Insights component.
	InstrumentationKey pulumi.StringOutput `pulumi:"instrumentationKey"`
	// The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The date which the component got migrated to LA, in ISO 8601 format.
	LaMigrationDate pulumi.StringOutput `pulumi:"laMigrationDate"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// List of linked private link scope resources.
	PrivateLinkScopedResources PrivateLinkScopedResourceResponseArrayOutput `pulumi:"privateLinkScopedResources"`
	// Current state of this component: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The network access type for accessing Application Insights ingestion.
	PublicNetworkAccessForIngestion pulumi.StringPtrOutput `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Application Insights query.
	PublicNetworkAccessForQuery pulumi.StringPtrOutput `pulumi:"publicNetworkAccessForQuery"`
	// Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
	RequestSource pulumi.StringPtrOutput `pulumi:"requestSource"`
	// Retention period in days.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
	SamplingPercentage pulumi.Float64PtrOutput `pulumi:"samplingPercentage"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure Tenant Id.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Resource Id of the log analytics workspace which the data will be ingested to. This property is required to create an application with this API version. Applications from older versions will not have this property.
	WorkspaceResourceId pulumi.StringPtrOutput `pulumi:"workspaceResourceId"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ApplicationType == nil {
		args.ApplicationType = pulumi.String("web")
	}
	if args.FlowType == nil {
		args.FlowType = pulumi.StringPtr("Bluefield")
	}
	if args.IngestionMode == nil {
		args.IngestionMode = pulumi.StringPtr("LogAnalytics")
	}
	if args.RequestSource == nil {
		args.RequestSource = pulumi.StringPtr("rest")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20150501:Component"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180501preview:Component"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200202:Component"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200202preview:Component"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Component
	err := ctx.RegisterResource("azure-native:insights:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponent gets an existing Component resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentState, opts ...pulumi.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("azure-native:insights:Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Component resources.
type componentState struct {
}

type ComponentState struct {
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	// Type of application being monitored.
	ApplicationType string `pulumi:"applicationType"`
	// Disable IP masking.
	DisableIpMasking *bool `pulumi:"disableIpMasking"`
	// Disable Non-AAD based Auth.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
	FlowType *string `pulumi:"flowType"`
	// Force users to create their own storage account for profiler and debugger.
	ForceCustomerStorageForProfiler *bool `pulumi:"forceCustomerStorageForProfiler"`
	// The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
	HockeyAppId *string `pulumi:"hockeyAppId"`
	// Purge data immediately after 30 days.
	ImmediatePurgeDataOn30Days *bool `pulumi:"immediatePurgeDataOn30Days"`
	// Indicates the flow of the ingestion.
	IngestionMode *string `pulumi:"ingestionMode"`
	// The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
	Kind string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// The network access type for accessing Application Insights ingestion.
	PublicNetworkAccessForIngestion *string `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Application Insights query.
	PublicNetworkAccessForQuery *string `pulumi:"publicNetworkAccessForQuery"`
	// Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
	RequestSource *string `pulumi:"requestSource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName *string `pulumi:"resourceName"`
	// Retention period in days.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
	SamplingPercentage *float64 `pulumi:"samplingPercentage"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource Id of the log analytics workspace which the data will be ingested to. This property is required to create an application with this API version. Applications from older versions will not have this property.
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	// Type of application being monitored.
	ApplicationType pulumi.StringInput
	// Disable IP masking.
	DisableIpMasking pulumi.BoolPtrInput
	// Disable Non-AAD based Auth.
	DisableLocalAuth pulumi.BoolPtrInput
	// Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
	FlowType pulumi.StringPtrInput
	// Force users to create their own storage account for profiler and debugger.
	ForceCustomerStorageForProfiler pulumi.BoolPtrInput
	// The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
	HockeyAppId pulumi.StringPtrInput
	// Purge data immediately after 30 days.
	ImmediatePurgeDataOn30Days pulumi.BoolPtrInput
	// Indicates the flow of the ingestion.
	IngestionMode pulumi.StringPtrInput
	// The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
	Kind pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// The network access type for accessing Application Insights ingestion.
	PublicNetworkAccessForIngestion pulumi.StringPtrInput
	// The network access type for accessing Application Insights query.
	PublicNetworkAccessForQuery pulumi.StringPtrInput
	// Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
	RequestSource pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringPtrInput
	// Retention period in days.
	RetentionInDays pulumi.IntPtrInput
	// Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
	SamplingPercentage pulumi.Float64PtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource Id of the log analytics workspace which the data will be ingested to. This property is required to create an application with this API version. Applications from older versions will not have this property.
	WorkspaceResourceId pulumi.StringPtrInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

func (i *Component) ToOutput(ctx context.Context) pulumix.Output[*Component] {
	return pulumix.Output[*Component]{
		OutputState: i.ToComponentOutputWithContext(ctx).OutputState,
	}
}

type ComponentOutput struct{ *pulumi.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

func (o ComponentOutput) ToOutput(ctx context.Context) pulumix.Output[*Component] {
	return pulumix.Output[*Component]{
		OutputState: o.OutputState,
	}
}

// Application Insights Unique ID for your Application.
func (o ComponentOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The unique ID of your application. This field mirrors the 'Name' field and cannot be changed.
func (o ComponentOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// Type of application being monitored.
func (o ComponentOutput) ApplicationType() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.ApplicationType }).(pulumi.StringOutput)
}

// Application Insights component connection string.
func (o ComponentOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// Creation Date for the Application Insights component, in ISO 8601 format.
func (o ComponentOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Disable IP masking.
func (o ComponentOutput) DisableIpMasking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.BoolPtrOutput { return v.DisableIpMasking }).(pulumi.BoolPtrOutput)
}

// Disable Non-AAD based Auth.
func (o ComponentOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Resource etag
func (o ComponentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
func (o ComponentOutput) FlowType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.FlowType }).(pulumi.StringPtrOutput)
}

// Force users to create their own storage account for profiler and debugger.
func (o ComponentOutput) ForceCustomerStorageForProfiler() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.BoolPtrOutput { return v.ForceCustomerStorageForProfiler }).(pulumi.BoolPtrOutput)
}

// The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
func (o ComponentOutput) HockeyAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.HockeyAppId }).(pulumi.StringPtrOutput)
}

// Token used to authenticate communications with between Application Insights and HockeyApp.
func (o ComponentOutput) HockeyAppToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.HockeyAppToken }).(pulumi.StringOutput)
}

// Purge data immediately after 30 days.
func (o ComponentOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.BoolPtrOutput { return v.ImmediatePurgeDataOn30Days }).(pulumi.BoolPtrOutput)
}

// Indicates the flow of the ingestion.
func (o ComponentOutput) IngestionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.IngestionMode }).(pulumi.StringPtrOutput)
}

// Application Insights Instrumentation key. A read-only value that applications can use to identify the destination for all telemetry sent to Azure Application Insights. This value will be supplied upon construction of each new Application Insights component.
func (o ComponentOutput) InstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.InstrumentationKey }).(pulumi.StringOutput)
}

// The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
func (o ComponentOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The date which the component got migrated to LA, in ISO 8601 format.
func (o ComponentOutput) LaMigrationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.LaMigrationDate }).(pulumi.StringOutput)
}

// Resource location
func (o ComponentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o ComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of linked private link scope resources.
func (o ComponentOutput) PrivateLinkScopedResources() PrivateLinkScopedResourceResponseArrayOutput {
	return o.ApplyT(func(v *Component) PrivateLinkScopedResourceResponseArrayOutput { return v.PrivateLinkScopedResources }).(PrivateLinkScopedResourceResponseArrayOutput)
}

// Current state of this component: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
func (o ComponentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The network access type for accessing Application Insights ingestion.
func (o ComponentOutput) PublicNetworkAccessForIngestion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.PublicNetworkAccessForIngestion }).(pulumi.StringPtrOutput)
}

// The network access type for accessing Application Insights query.
func (o ComponentOutput) PublicNetworkAccessForQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.PublicNetworkAccessForQuery }).(pulumi.StringPtrOutput)
}

// Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
func (o ComponentOutput) RequestSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.RequestSource }).(pulumi.StringPtrOutput)
}

// Retention period in days.
func (o ComponentOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.IntPtrOutput { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

// Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
func (o ComponentOutput) SamplingPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Component) pulumi.Float64PtrOutput { return v.SamplingPercentage }).(pulumi.Float64PtrOutput)
}

// Resource tags
func (o ComponentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Component) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure Tenant Id.
func (o ComponentOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Azure resource type
func (o ComponentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Resource Id of the log analytics workspace which the data will be ingested to. This property is required to create an application with this API version. Applications from older versions will not have this property.
func (o ComponentOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ComponentOutput{})
}
