// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

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
	case "azure-native:eventhub/v20240101:ApplicationGroup":
		r = &ApplicationGroup{}
	case "azure-native:eventhub/v20240101:Cluster":
		r = &Cluster{}
	case "azure-native:eventhub/v20240101:ConsumerGroup":
		r = &ConsumerGroup{}
	case "azure-native:eventhub/v20240101:DisasterRecoveryConfig":
		r = &DisasterRecoveryConfig{}
	case "azure-native:eventhub/v20240101:EventHub":
		r = &EventHub{}
	case "azure-native:eventhub/v20240101:EventHubAuthorizationRule":
		r = &EventHubAuthorizationRule{}
	case "azure-native:eventhub/v20240101:Namespace":
		r = &Namespace{}
	case "azure-native:eventhub/v20240101:NamespaceAuthorizationRule":
		r = &NamespaceAuthorizationRule{}
	case "azure-native:eventhub/v20240101:NamespaceNetworkRuleSet":
		r = &NamespaceNetworkRuleSet{}
	case "azure-native:eventhub/v20240101:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:eventhub/v20240101:SchemaRegistry":
		r = &SchemaRegistry{}
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
		"eventhub/v20240101",
		&module{version},
	)
}
