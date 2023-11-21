// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231215preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:eventgrid/v20231215preview:CaCertificate":
		r = &CaCertificate{}
	case "azure-native:eventgrid/v20231215preview:Channel":
		r = &Channel{}
	case "azure-native:eventgrid/v20231215preview:Client":
		r = &Client{}
	case "azure-native:eventgrid/v20231215preview:ClientGroup":
		r = &ClientGroup{}
	case "azure-native:eventgrid/v20231215preview:Domain":
		r = &Domain{}
	case "azure-native:eventgrid/v20231215preview:DomainEventSubscription":
		r = &DomainEventSubscription{}
	case "azure-native:eventgrid/v20231215preview:DomainTopic":
		r = &DomainTopic{}
	case "azure-native:eventgrid/v20231215preview:DomainTopicEventSubscription":
		r = &DomainTopicEventSubscription{}
	case "azure-native:eventgrid/v20231215preview:EventSubscription":
		r = &EventSubscription{}
	case "azure-native:eventgrid/v20231215preview:Namespace":
		r = &Namespace{}
	case "azure-native:eventgrid/v20231215preview:NamespaceTopic":
		r = &NamespaceTopic{}
	case "azure-native:eventgrid/v20231215preview:NamespaceTopicEventSubscription":
		r = &NamespaceTopicEventSubscription{}
	case "azure-native:eventgrid/v20231215preview:PartnerConfiguration":
		r = &PartnerConfiguration{}
	case "azure-native:eventgrid/v20231215preview:PartnerDestination":
		r = &PartnerDestination{}
	case "azure-native:eventgrid/v20231215preview:PartnerNamespace":
		r = &PartnerNamespace{}
	case "azure-native:eventgrid/v20231215preview:PartnerRegistration":
		r = &PartnerRegistration{}
	case "azure-native:eventgrid/v20231215preview:PartnerTopic":
		r = &PartnerTopic{}
	case "azure-native:eventgrid/v20231215preview:PartnerTopicEventSubscription":
		r = &PartnerTopicEventSubscription{}
	case "azure-native:eventgrid/v20231215preview:PermissionBinding":
		r = &PermissionBinding{}
	case "azure-native:eventgrid/v20231215preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:eventgrid/v20231215preview:SystemTopic":
		r = &SystemTopic{}
	case "azure-native:eventgrid/v20231215preview:SystemTopicEventSubscription":
		r = &SystemTopicEventSubscription{}
	case "azure-native:eventgrid/v20231215preview:Topic":
		r = &Topic{}
	case "azure-native:eventgrid/v20231215preview:TopicEventSubscription":
		r = &TopicEventSubscription{}
	case "azure-native:eventgrid/v20231215preview:TopicSpace":
		r = &TopicSpace{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"eventgrid/v20231215preview",
		&module{version},
	)
}