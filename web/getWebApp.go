// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the details of a web, mobile, or API app.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2016-08-01, 2018-11-01, 2020-10-01, 2023-01-01.
func LookupWebApp(ctx *pulumi.Context, args *LookupWebAppArgs, opts ...pulumi.InvokeOption) (*LookupWebAppResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppResult
	err := ctx.Invoke("azure-native:web:getWebApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebAppArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A web app, a mobile app backend, or an API app.
type LookupWebAppResult struct {
	// Management information availability state for the app.
	AvailabilityState string `pulumi:"availabilityState"`
	// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// client certificate authentication comma-separated exclusion paths
	ClientCertExclusionPaths *string `pulumi:"clientCertExclusionPaths"`
	// This composes with ClientCertEnabled setting.
	// - ClientCertEnabled: false means ClientCert is ignored.
	// - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
	// - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
	ClientCertMode *string `pulumi:"clientCertMode"`
	// Size of the function container.
	ContainerSize *int `pulumi:"containerSize"`
	// Unique identifier that verifies the custom domains assigned to the app. Customer will add this id to a txt record for verification.
	CustomDomainVerificationId *string `pulumi:"customDomainVerificationId"`
	// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// Default hostname of the app. Read-only.
	DefaultHostName string `pulumi:"defaultHostName"`
	// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
	Enabled *bool `pulumi:"enabled"`
	// Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
	// the app is not served on those hostnames.
	EnabledHostNames []string `pulumi:"enabledHostNames"`
	// Extended Location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates []HostNameSslStateResponse `pulumi:"hostNameSslStates"`
	// Hostnames associated with the app.
	HostNames []string `pulumi:"hostNames"`
	// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	//  If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled *bool `pulumi:"hostNamesDisabled"`
	// App Service Environment to use for the app.
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// Hyper-V sandbox.
	HyperV *bool `pulumi:"hyperV"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Managed service identity.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Specifies an operation id if this site has a pending operation.
	InProgressOperationId string `pulumi:"inProgressOperationId"`
	// <code>true</code> if the app is a default container; otherwise, <code>false</code>.
	IsDefaultContainer bool `pulumi:"isDefaultContainer"`
	// Obsolete: Hyper-V sandbox.
	IsXenon *bool `pulumi:"isXenon"`
	// Identity to use for Key Vault Reference authentication.
	KeyVaultReferenceIdentity *string `pulumi:"keyVaultReferenceIdentity"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Last time the app was modified, in UTC. Read-only.
	LastModifiedTimeUtc string `pulumi:"lastModifiedTimeUtc"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Azure Resource Manager ID of the customer's selected Managed Environment on which to host this app. This must be of the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.App/managedEnvironments/{managedEnvironmentName}
	ManagedEnvironmentId *string `pulumi:"managedEnvironmentId"`
	// Maximum number of workers.
	// This only applies to Functions container.
	MaxNumberOfWorkers int `pulumi:"maxNumberOfWorkers"`
	// Resource Name.
	Name string `pulumi:"name"`
	// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
	OutboundIpAddresses string `pulumi:"outboundIpAddresses"`
	// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants except dataComponent. Read-only.
	PossibleOutboundIpAddresses string `pulumi:"possibleOutboundIpAddresses"`
	// Property to allow or block all public traffic. Allowed Values: 'Enabled', 'Disabled' or an empty string.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Site redundancy mode
	RedundancyMode *string `pulumi:"redundancyMode"`
	// Name of the repository site.
	RepositorySiteName string `pulumi:"repositorySiteName"`
	// <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved *bool `pulumi:"reserved"`
	// Name of the resource group the app belongs to. Read-only.
	ResourceGroup string `pulumi:"resourceGroup"`
	// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
	ScmSiteAlsoStopped *bool `pulumi:"scmSiteAlsoStopped"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId *string `pulumi:"serverFarmId"`
	// Configuration of the app.
	SiteConfig *SiteConfigResponse `pulumi:"siteConfig"`
	// Status of the last deployment slot swap operation.
	SlotSwapStatus SlotSwapStatusResponse `pulumi:"slotSwapStatus"`
	// Current state of the app.
	State string `pulumi:"state"`
	// Checks if Customer provided storage account is required
	StorageAccountRequired *bool `pulumi:"storageAccountRequired"`
	// App suspended till in case memory-time quota is exceeded.
	SuspendedTill string `pulumi:"suspendedTill"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies which deployment slot this app will swap into. Read-only.
	TargetSwapSlot string `pulumi:"targetSwapSlot"`
	// Azure Traffic Manager hostnames associated with the app. Read-only.
	TrafficManagerHostNames []string `pulumi:"trafficManagerHostNames"`
	// Resource type.
	Type string `pulumi:"type"`
	// State indicating whether the app has exceeded its quota usage. Read-only.
	UsageState string `pulumi:"usageState"`
	// Azure Resource Manager ID of the Virtual network and subnet to be joined by Regional VNET Integration.
	// This must be of the form /subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
	// To enable accessing content over virtual network
	VnetContentShareEnabled *bool `pulumi:"vnetContentShareEnabled"`
	// To enable pulling image over Virtual Network
	VnetImagePullEnabled *bool `pulumi:"vnetImagePullEnabled"`
	// Virtual Network Route All enabled. This causes all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied.
	VnetRouteAllEnabled *bool `pulumi:"vnetRouteAllEnabled"`
}

// Defaults sets the appropriate defaults for LookupWebAppResult
func (val *LookupWebAppResult) Defaults() *LookupWebAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.HyperV == nil {
		hyperV_ := false
		tmp.HyperV = &hyperV_
	}
	if tmp.IsXenon == nil {
		isXenon_ := false
		tmp.IsXenon = &isXenon_
	}
	if tmp.Reserved == nil {
		reserved_ := false
		tmp.Reserved = &reserved_
	}
	if tmp.ScmSiteAlsoStopped == nil {
		scmSiteAlsoStopped_ := false
		tmp.ScmSiteAlsoStopped = &scmSiteAlsoStopped_
	}
	tmp.SiteConfig = tmp.SiteConfig.Defaults()

	return &tmp
}

func LookupWebAppOutput(ctx *pulumi.Context, args LookupWebAppOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppResult, error) {
			args := v.(LookupWebAppArgs)
			r, err := LookupWebApp(ctx, &args, opts...)
			var s LookupWebAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppResultOutput)
}

type LookupWebAppOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppArgs)(nil)).Elem()
}

// A web app, a mobile app backend, or an API app.
type LookupWebAppResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppResult)(nil)).Elem()
}

func (o LookupWebAppResultOutput) ToLookupWebAppResultOutput() LookupWebAppResultOutput {
	return o
}

func (o LookupWebAppResultOutput) ToLookupWebAppResultOutputWithContext(ctx context.Context) LookupWebAppResultOutput {
	return o
}

// Management information availability state for the app.
func (o LookupWebAppResultOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.AvailabilityState }).(pulumi.StringOutput)
}

// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
func (o LookupWebAppResultOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
func (o LookupWebAppResultOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

// client certificate authentication comma-separated exclusion paths
func (o LookupWebAppResultOutput) ClientCertExclusionPaths() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.ClientCertExclusionPaths }).(pulumi.StringPtrOutput)
}

// This composes with ClientCertEnabled setting.
// - ClientCertEnabled: false means ClientCert is ignored.
// - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
// - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
func (o LookupWebAppResultOutput) ClientCertMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.ClientCertMode }).(pulumi.StringPtrOutput)
}

// Size of the function container.
func (o LookupWebAppResultOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *int { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

// Unique identifier that verifies the custom domains assigned to the app. Customer will add this id to a txt record for verification.
func (o LookupWebAppResultOutput) CustomDomainVerificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.CustomDomainVerificationId }).(pulumi.StringPtrOutput)
}

// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
func (o LookupWebAppResultOutput) DailyMemoryTimeQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *int { return v.DailyMemoryTimeQuota }).(pulumi.IntPtrOutput)
}

// Default hostname of the app. Read-only.
func (o LookupWebAppResultOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.DefaultHostName }).(pulumi.StringOutput)
}

// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
func (o LookupWebAppResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
// the app is not served on those hostnames.
func (o LookupWebAppResultOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []string { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

// Extended Location.
func (o LookupWebAppResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
func (o LookupWebAppResultOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []HostNameSslStateResponse { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

// Hostnames associated with the app.
func (o LookupWebAppResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
//
//	If <code>true</code>, the app is only accessible via API management process.
func (o LookupWebAppResultOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

// App Service Environment to use for the app.
func (o LookupWebAppResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
// http requests
func (o LookupWebAppResultOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

// Hyper-V sandbox.
func (o LookupWebAppResultOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.HyperV }).(pulumi.BoolPtrOutput)
}

// Resource Id.
func (o LookupWebAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity.
func (o LookupWebAppResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Specifies an operation id if this site has a pending operation.
func (o LookupWebAppResultOutput) InProgressOperationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.InProgressOperationId }).(pulumi.StringOutput)
}

// <code>true</code> if the app is a default container; otherwise, <code>false</code>.
func (o LookupWebAppResultOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppResult) bool { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

// Obsolete: Hyper-V sandbox.
func (o LookupWebAppResultOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

// Identity to use for Key Vault Reference authentication.
func (o LookupWebAppResultOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o LookupWebAppResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Last time the app was modified, in UTC. Read-only.
func (o LookupWebAppResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// Resource Location.
func (o LookupWebAppResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure Resource Manager ID of the customer's selected Managed Environment on which to host this app. This must be of the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.App/managedEnvironments/{managedEnvironmentName}
func (o LookupWebAppResultOutput) ManagedEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.ManagedEnvironmentId }).(pulumi.StringPtrOutput)
}

// Maximum number of workers.
// This only applies to Functions container.
func (o LookupWebAppResultOutput) MaxNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebAppResult) int { return v.MaxNumberOfWorkers }).(pulumi.IntOutput)
}

// Resource Name.
func (o LookupWebAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
func (o LookupWebAppResultOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants except dataComponent. Read-only.
func (o LookupWebAppResultOutput) PossibleOutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.PossibleOutboundIpAddresses }).(pulumi.StringOutput)
}

// Property to allow or block all public traffic. Allowed Values: 'Enabled', 'Disabled' or an empty string.
func (o LookupWebAppResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Site redundancy mode
func (o LookupWebAppResultOutput) RedundancyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.RedundancyMode }).(pulumi.StringPtrOutput)
}

// Name of the repository site.
func (o LookupWebAppResultOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.RepositorySiteName }).(pulumi.StringOutput)
}

// <code>true</code> if reserved; otherwise, <code>false</code>.
func (o LookupWebAppResultOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.Reserved }).(pulumi.BoolPtrOutput)
}

// Name of the resource group the app belongs to. Read-only.
func (o LookupWebAppResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
func (o LookupWebAppResultOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
func (o LookupWebAppResultOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

// Configuration of the app.
func (o LookupWebAppResultOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *SiteConfigResponse { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

// Status of the last deployment slot swap operation.
func (o LookupWebAppResultOutput) SlotSwapStatus() SlotSwapStatusResponseOutput {
	return o.ApplyT(func(v LookupWebAppResult) SlotSwapStatusResponse { return v.SlotSwapStatus }).(SlotSwapStatusResponseOutput)
}

// Current state of the app.
func (o LookupWebAppResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.State }).(pulumi.StringOutput)
}

// Checks if Customer provided storage account is required
func (o LookupWebAppResultOutput) StorageAccountRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.StorageAccountRequired }).(pulumi.BoolPtrOutput)
}

// App suspended till in case memory-time quota is exceeded.
func (o LookupWebAppResultOutput) SuspendedTill() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.SuspendedTill }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupWebAppResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies which deployment slot this app will swap into. Read-only.
func (o LookupWebAppResultOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

// Azure Traffic Manager hostnames associated with the app. Read-only.
func (o LookupWebAppResultOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []string { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

// Resource type.
func (o LookupWebAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Type }).(pulumi.StringOutput)
}

// State indicating whether the app has exceeded its quota usage. Read-only.
func (o LookupWebAppResultOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.UsageState }).(pulumi.StringOutput)
}

// Azure Resource Manager ID of the Virtual network and subnet to be joined by Regional VNET Integration.
// This must be of the form /subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}
func (o LookupWebAppResultOutput) VirtualNetworkSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.VirtualNetworkSubnetId }).(pulumi.StringPtrOutput)
}

// To enable accessing content over virtual network
func (o LookupWebAppResultOutput) VnetContentShareEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.VnetContentShareEnabled }).(pulumi.BoolPtrOutput)
}

// To enable pulling image over Virtual Network
func (o LookupWebAppResultOutput) VnetImagePullEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.VnetImagePullEnabled }).(pulumi.BoolPtrOutput)
}

// Virtual Network Route All enabled. This causes all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied.
func (o LookupWebAppResultOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.VnetRouteAllEnabled }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppResultOutput{})
}
