// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of a domain.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// EventGrid Domain.
type LookupDomainResult struct {
	// This Boolean is used to specify the creation mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, creation of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is null or set to true, Event Grid is responsible of automatically creating the domain topic when the first event subscription is
	// created at the scope of the domain topic. If this property is set to false, then creating the first event subscription will require creating a domain topic
	// by the user. The self-management mode can be used if the user wants full control of when the domain topic is created, while auto-managed mode provides the
	// flexibility to perform less operations and manage fewer resources by the user. Also, note that in auto-managed creation mode, user is allowed to create the
	// domain topic on demand if needed.
	AutoCreateTopicWithFirstSubscription *bool `pulumi:"autoCreateTopicWithFirstSubscription"`
	// This Boolean is used to specify the deletion mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, deletion of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is set to true, Event Grid is responsible of automatically deleting the domain topic when the last event subscription at the scope
	// of the domain topic is deleted. If this property is set to false, then the user needs to manually delete the domain topic when it is no longer needed
	// (e.g., when last event subscription is deleted and the resource needs to be cleaned up). The self-management mode can be used if the user wants full
	// control of when the domain topic needs to be deleted, while auto-managed mode provides the flexibility to perform less operations and manage fewer
	// resources by the user.
	AutoDeleteTopicWithLastSubscription *bool `pulumi:"autoDeleteTopicWithLastSubscription"`
	// Data Residency Boundary of the resource.
	DataResidencyBoundary *string `pulumi:"dataResidencyBoundary"`
	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the domain.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Endpoint for the Event Grid Domain Resource which is used for publishing the events.
	Endpoint string `pulumi:"endpoint"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Identity information for the Event Grid Domain resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleResponse `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the Event Grid Domain Resource.
	InputSchema *string `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Metric resource id for the Event Grid Domain Resource.
	MetricResourceId string `pulumi:"metricResourceId"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the Event Grid Domain Resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The system metadata relating to the Event Grid Domain resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDomainResult
func (val *LookupDomainResult) Defaults() *LookupDomainResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AutoCreateTopicWithFirstSubscription == nil {
		autoCreateTopicWithFirstSubscription_ := true
		tmp.AutoCreateTopicWithFirstSubscription = &autoCreateTopicWithFirstSubscription_
	}
	if tmp.AutoDeleteTopicWithLastSubscription == nil {
		autoDeleteTopicWithLastSubscription_ := true
		tmp.AutoDeleteTopicWithLastSubscription = &autoDeleteTopicWithLastSubscription_
	}
	if tmp.DisableLocalAuth == nil {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	if tmp.InputSchema == nil {
		inputSchema_ := "EventGridSchema"
		tmp.InputSchema = &inputSchema_
	}
	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
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
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

// EventGrid Domain.
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

// This Boolean is used to specify the creation mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
// In this context, creation of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
// When this property is null or set to true, Event Grid is responsible of automatically creating the domain topic when the first event subscription is
// created at the scope of the domain topic. If this property is set to false, then creating the first event subscription will require creating a domain topic
// by the user. The self-management mode can be used if the user wants full control of when the domain topic is created, while auto-managed mode provides the
// flexibility to perform less operations and manage fewer resources by the user. Also, note that in auto-managed creation mode, user is allowed to create the
// domain topic on demand if needed.
func (o LookupDomainResultOutput) AutoCreateTopicWithFirstSubscription() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.AutoCreateTopicWithFirstSubscription }).(pulumi.BoolPtrOutput)
}

// This Boolean is used to specify the deletion mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
// In this context, deletion of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
// When this property is set to true, Event Grid is responsible of automatically deleting the domain topic when the last event subscription at the scope
// of the domain topic is deleted. If this property is set to false, then the user needs to manually delete the domain topic when it is no longer needed
// (e.g., when last event subscription is deleted and the resource needs to be cleaned up). The self-management mode can be used if the user wants full
// control of when the domain topic needs to be deleted, while auto-managed mode provides the flexibility to perform less operations and manage fewer
// resources by the user.
func (o LookupDomainResultOutput) AutoDeleteTopicWithLastSubscription() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.AutoDeleteTopicWithLastSubscription }).(pulumi.BoolPtrOutput)
}

// Data Residency Boundary of the resource.
func (o LookupDomainResultOutput) DataResidencyBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DataResidencyBoundary }).(pulumi.StringPtrOutput)
}

// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the domain.
func (o LookupDomainResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Endpoint for the Event Grid Domain Resource which is used for publishing the events.
func (o LookupDomainResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Fully qualified identifier of the resource.
func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity information for the Event Grid Domain resource.
func (o LookupDomainResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
func (o LookupDomainResultOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []InboundIpRuleResponse { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

// This determines the format that Event Grid should expect for incoming events published to the Event Grid Domain Resource.
func (o LookupDomainResultOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.InputSchema }).(pulumi.StringPtrOutput)
}

// Information about the InputSchemaMapping which specified the info about mapping event payload.
func (o LookupDomainResultOutput) InputSchemaMapping() JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *JsonInputSchemaMappingResponse { return v.InputSchemaMapping }).(JsonInputSchemaMappingResponsePtrOutput)
}

// Location of the resource.
func (o LookupDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

// Metric resource id for the Event Grid Domain Resource.
func (o LookupDomainResultOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.MetricResourceId }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections.
func (o LookupDomainResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the Event Grid Domain Resource.
func (o LookupDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
func (o LookupDomainResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The system metadata relating to the Event Grid Domain resource.
func (o LookupDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o LookupDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o LookupDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
