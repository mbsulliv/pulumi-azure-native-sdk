// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Get a domain.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:domainregistration/v20220901:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about a domain.
type LookupDomainResult struct {
	AuthCode *string `pulumi:"authCode"`
	// <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Domain creation timestamp.
	CreatedTime string `pulumi:"createdTime"`
	// Current DNS type
	DnsType *string `pulumi:"dnsType"`
	// Azure DNS Zone to use
	DnsZoneId *string `pulumi:"dnsZoneId"`
	// Reasons why domain is not renewable.
	DomainNotRenewableReasons []string `pulumi:"domainNotRenewableReasons"`
	// Domain expiration timestamp.
	ExpirationTime string `pulumi:"expirationTime"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Timestamp when the domain was renewed last time.
	LastRenewedTime string `pulumi:"lastRenewedTime"`
	// Resource Location.
	Location string `pulumi:"location"`
	// All hostnames derived from the domain and assigned to Azure resources.
	ManagedHostNames []HostNameResponse `pulumi:"managedHostNames"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Name servers.
	NameServers []string `pulumi:"nameServers"`
	// <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
	Privacy *bool `pulumi:"privacy"`
	// Domain provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// <code>true</code> if Azure can assign this domain to App Service apps; otherwise, <code>false</code>. This value will be <code>true</code> if domain registration status is active and
	//  it is hosted on name servers Azure has programmatic access to.
	ReadyForDnsRecordManagement bool `pulumi:"readyForDnsRecordManagement"`
	// Domain registration status.
	RegistrationStatus string `pulumi:"registrationStatus"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Target DNS type (would be used for migration)
	TargetDnsType *string `pulumi:"targetDnsType"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDomainResult
func (val *LookupDomainResult) Defaults() *LookupDomainResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AutoRenew == nil {
		autoRenew_ := true
		tmp.AutoRenew = &autoRenew_
	}
	return &tmp
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	// Name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

// Information about a domain.
type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDomainResult] {
	return pulumix.Output[LookupDomainResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupDomainResultOutput) AuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.AuthCode }).(pulumi.StringPtrOutput)
}

// <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
func (o LookupDomainResultOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// Domain creation timestamp.
func (o LookupDomainResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// Current DNS type
func (o LookupDomainResultOutput) DnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DnsType }).(pulumi.StringPtrOutput)
}

// Azure DNS Zone to use
func (o LookupDomainResultOutput) DnsZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DnsZoneId }).(pulumi.StringPtrOutput)
}

// Reasons why domain is not renewable.
func (o LookupDomainResultOutput) DomainNotRenewableReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []string { return v.DomainNotRenewableReasons }).(pulumi.StringArrayOutput)
}

// Domain expiration timestamp.
func (o LookupDomainResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

// Resource Id.
func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupDomainResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Timestamp when the domain was renewed last time.
func (o LookupDomainResultOutput) LastRenewedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.LastRenewedTime }).(pulumi.StringOutput)
}

// Resource Location.
func (o LookupDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

// All hostnames derived from the domain and assigned to Azure resources.
func (o LookupDomainResultOutput) ManagedHostNames() HostNameResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []HostNameResponse { return v.ManagedHostNames }).(HostNameResponseArrayOutput)
}

// Resource Name.
func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name servers.
func (o LookupDomainResultOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

// <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
func (o LookupDomainResultOutput) Privacy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.Privacy }).(pulumi.BoolPtrOutput)
}

// Domain provisioning state.
func (o LookupDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// <code>true</code> if Azure can assign this domain to App Service apps; otherwise, <code>false</code>. This value will be <code>true</code> if domain registration status is active and
//
//	it is hosted on name servers Azure has programmatic access to.
func (o LookupDomainResultOutput) ReadyForDnsRecordManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainResult) bool { return v.ReadyForDnsRecordManagement }).(pulumi.BoolOutput)
}

// Domain registration status.
func (o LookupDomainResultOutput) RegistrationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.RegistrationStatus }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Target DNS type (would be used for migration)
func (o LookupDomainResultOutput) TargetDnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.TargetDnsType }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
