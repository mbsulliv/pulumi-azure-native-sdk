// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401

// The allowed type DNS record types for this profile.
type AllowedEndpointRecordType string

const (
	AllowedEndpointRecordTypeDomainName  = AllowedEndpointRecordType("DomainName")
	AllowedEndpointRecordTypeIPv4Address = AllowedEndpointRecordType("IPv4Address")
	AllowedEndpointRecordTypeIPv6Address = AllowedEndpointRecordType("IPv6Address")
	AllowedEndpointRecordTypeAny         = AllowedEndpointRecordType("Any")
)

// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method.
type AlwaysServe string

const (
	AlwaysServeEnabled  = AlwaysServe("Enabled")
	AlwaysServeDisabled = AlwaysServe("Disabled")
)

// The monitoring status of the endpoint.
type EndpointMonitorStatus string

const (
	EndpointMonitorStatusCheckingEndpoint = EndpointMonitorStatus("CheckingEndpoint")
	EndpointMonitorStatusOnline           = EndpointMonitorStatus("Online")
	EndpointMonitorStatusDegraded         = EndpointMonitorStatus("Degraded")
	EndpointMonitorStatusDisabled         = EndpointMonitorStatus("Disabled")
	EndpointMonitorStatusInactive         = EndpointMonitorStatus("Inactive")
	EndpointMonitorStatusStopped          = EndpointMonitorStatus("Stopped")
	EndpointMonitorStatusUnmonitored      = EndpointMonitorStatus("Unmonitored")
)

// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
type EndpointStatus string

const (
	EndpointStatusEnabled  = EndpointStatus("Enabled")
	EndpointStatusDisabled = EndpointStatus("Disabled")
)

// The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health.
type MonitorProtocol string

const (
	MonitorProtocolHTTP  = MonitorProtocol("HTTP")
	MonitorProtocolHTTPS = MonitorProtocol("HTTPS")
	MonitorProtocolTCP   = MonitorProtocol("TCP")
)

// The profile-level monitoring status of the Traffic Manager profile.
type ProfileMonitorStatus string

const (
	ProfileMonitorStatusCheckingEndpoints = ProfileMonitorStatus("CheckingEndpoints")
	ProfileMonitorStatusOnline            = ProfileMonitorStatus("Online")
	ProfileMonitorStatusDegraded          = ProfileMonitorStatus("Degraded")
	ProfileMonitorStatusDisabled          = ProfileMonitorStatus("Disabled")
	ProfileMonitorStatusInactive          = ProfileMonitorStatus("Inactive")
)

// The status of the Traffic Manager profile.
type ProfileStatus string

const (
	ProfileStatusEnabled  = ProfileStatus("Enabled")
	ProfileStatusDisabled = ProfileStatus("Disabled")
)

// The traffic routing method of the Traffic Manager profile.
type TrafficRoutingMethod string

const (
	TrafficRoutingMethodPerformance = TrafficRoutingMethod("Performance")
	TrafficRoutingMethodPriority    = TrafficRoutingMethod("Priority")
	TrafficRoutingMethodWeighted    = TrafficRoutingMethod("Weighted")
	TrafficRoutingMethodGeographic  = TrafficRoutingMethod("Geographic")
	TrafficRoutingMethodMultiValue  = TrafficRoutingMethod("MultiValue")
	TrafficRoutingMethodSubnet      = TrafficRoutingMethod("Subnet")
)

// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
type TrafficViewEnrollmentStatus string

const (
	TrafficViewEnrollmentStatusEnabled  = TrafficViewEnrollmentStatus("Enabled")
	TrafficViewEnrollmentStatusDisabled = TrafficViewEnrollmentStatus("Disabled")
)

func init() {
}