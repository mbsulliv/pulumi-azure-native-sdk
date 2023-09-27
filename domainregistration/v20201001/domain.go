// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Information about a domain.
type Domain struct {
	pulumi.CustomResourceState

	AuthCode pulumi.StringPtrOutput `pulumi:"authCode"`
	// <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// Domain creation timestamp.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Current DNS type
	DnsType pulumi.StringPtrOutput `pulumi:"dnsType"`
	// Azure DNS Zone to use
	DnsZoneId pulumi.StringPtrOutput `pulumi:"dnsZoneId"`
	// Reasons why domain is not renewable.
	DomainNotRenewableReasons pulumi.StringArrayOutput `pulumi:"domainNotRenewableReasons"`
	// Domain expiration timestamp.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Timestamp when the domain was renewed last time.
	LastRenewedTime pulumi.StringOutput `pulumi:"lastRenewedTime"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// All hostnames derived from the domain and assigned to Azure resources.
	ManagedHostNames HostNameResponseArrayOutput `pulumi:"managedHostNames"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name servers.
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
	Privacy pulumi.BoolPtrOutput `pulumi:"privacy"`
	// Domain provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// <code>true</code> if Azure can assign this domain to App Service apps; otherwise, <code>false</code>. This value will be <code>true</code> if domain registration status is active and
	//  it is hosted on name servers Azure has programmatic access to.
	ReadyForDnsRecordManagement pulumi.BoolOutput `pulumi:"readyForDnsRecordManagement"`
	// Domain registration status.
	RegistrationStatus pulumi.StringOutput `pulumi:"registrationStatus"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Target DNS type (would be used for migration)
	TargetDnsType pulumi.StringPtrOutput `pulumi:"targetDnsType"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Consent == nil {
		return nil, errors.New("invalid value for required argument 'Consent'")
	}
	if args.ContactAdmin == nil {
		return nil, errors.New("invalid value for required argument 'ContactAdmin'")
	}
	if args.ContactBilling == nil {
		return nil, errors.New("invalid value for required argument 'ContactBilling'")
	}
	if args.ContactRegistrant == nil {
		return nil, errors.New("invalid value for required argument 'ContactRegistrant'")
	}
	if args.ContactTech == nil {
		return nil, errors.New("invalid value for required argument 'ContactTech'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AutoRenew == nil {
		args.AutoRenew = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:domainregistration:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20150401:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20180201:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20190801:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20200601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20200901:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20201201:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210101:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210115:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210201:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210301:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20220301:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20220901:Domain"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("azure-native:domainregistration/v20201001:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:domainregistration/v20201001:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	AuthCode *string `pulumi:"authCode"`
	// <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Legal agreement consent.
	Consent DomainPurchaseConsent `pulumi:"consent"`
	// Administrative contact.
	ContactAdmin Contact `pulumi:"contactAdmin"`
	// Billing contact.
	ContactBilling Contact `pulumi:"contactBilling"`
	// Registrant contact.
	ContactRegistrant Contact `pulumi:"contactRegistrant"`
	// Technical contact.
	ContactTech Contact `pulumi:"contactTech"`
	// Current DNS type
	DnsType *DnsType `pulumi:"dnsType"`
	// Azure DNS Zone to use
	DnsZoneId *string `pulumi:"dnsZoneId"`
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
	Privacy *bool `pulumi:"privacy"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Target DNS type (would be used for migration)
	TargetDnsType *DnsType `pulumi:"targetDnsType"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	AuthCode pulumi.StringPtrInput
	// <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
	AutoRenew pulumi.BoolPtrInput
	// Legal agreement consent.
	Consent DomainPurchaseConsentInput
	// Administrative contact.
	ContactAdmin ContactInput
	// Billing contact.
	ContactBilling ContactInput
	// Registrant contact.
	ContactRegistrant ContactInput
	// Technical contact.
	ContactTech ContactInput
	// Current DNS type
	DnsType DnsTypePtrInput
	// Azure DNS Zone to use
	DnsZoneId pulumi.StringPtrInput
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
	Privacy pulumi.BoolPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Target DNS type (would be used for migration)
	TargetDnsType DnsTypePtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

func (i *Domain) ToOutput(ctx context.Context) pulumix.Output[*Domain] {
	return pulumix.Output[*Domain]{
		OutputState: i.ToDomainOutputWithContext(ctx).OutputState,
	}
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) ToOutput(ctx context.Context) pulumix.Output[*Domain] {
	return pulumix.Output[*Domain]{
		OutputState: o.OutputState,
	}
}

func (o DomainOutput) AuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.AuthCode }).(pulumi.StringPtrOutput)
}

// <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
func (o DomainOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// Domain creation timestamp.
func (o DomainOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Current DNS type
func (o DomainOutput) DnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.DnsType }).(pulumi.StringPtrOutput)
}

// Azure DNS Zone to use
func (o DomainOutput) DnsZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.DnsZoneId }).(pulumi.StringPtrOutput)
}

// Reasons why domain is not renewable.
func (o DomainOutput) DomainNotRenewableReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.DomainNotRenewableReasons }).(pulumi.StringArrayOutput)
}

// Domain expiration timestamp.
func (o DomainOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ExpirationTime }).(pulumi.StringOutput)
}

// Kind of resource.
func (o DomainOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Timestamp when the domain was renewed last time.
func (o DomainOutput) LastRenewedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.LastRenewedTime }).(pulumi.StringOutput)
}

// Resource Location.
func (o DomainOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// All hostnames derived from the domain and assigned to Azure resources.
func (o DomainOutput) ManagedHostNames() HostNameResponseArrayOutput {
	return o.ApplyT(func(v *Domain) HostNameResponseArrayOutput { return v.ManagedHostNames }).(HostNameResponseArrayOutput)
}

// Resource Name.
func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name servers.
func (o DomainOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.NameServers }).(pulumi.StringArrayOutput)
}

// <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
func (o DomainOutput) Privacy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.Privacy }).(pulumi.BoolPtrOutput)
}

// Domain provisioning state.
func (o DomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// <code>true</code> if Azure can assign this domain to App Service apps; otherwise, <code>false</code>. This value will be <code>true</code> if domain registration status is active and
//
//	it is hosted on name servers Azure has programmatic access to.
func (o DomainOutput) ReadyForDnsRecordManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolOutput { return v.ReadyForDnsRecordManagement }).(pulumi.BoolOutput)
}

// Domain registration status.
func (o DomainOutput) RegistrationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.RegistrationStatus }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o DomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Domain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Target DNS type (would be used for migration)
func (o DomainOutput) TargetDnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.TargetDnsType }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o DomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
