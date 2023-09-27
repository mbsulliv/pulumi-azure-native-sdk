// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

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
	case "azure-native:containerregistry:AgentPool":
		r = &AgentPool{}
	case "azure-native:containerregistry:Archife":
		r = &Archife{}
	case "azure-native:containerregistry:ArchiveVersion":
		r = &ArchiveVersion{}
	case "azure-native:containerregistry:CacheRule":
		r = &CacheRule{}
	case "azure-native:containerregistry:ConnectedRegistry":
		r = &ConnectedRegistry{}
	case "azure-native:containerregistry:CredentialSet":
		r = &CredentialSet{}
	case "azure-native:containerregistry:ExportPipeline":
		r = &ExportPipeline{}
	case "azure-native:containerregistry:ImportPipeline":
		r = &ImportPipeline{}
	case "azure-native:containerregistry:PipelineRun":
		r = &PipelineRun{}
	case "azure-native:containerregistry:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:containerregistry:Registry":
		r = &Registry{}
	case "azure-native:containerregistry:Replication":
		r = &Replication{}
	case "azure-native:containerregistry:ScopeMap":
		r = &ScopeMap{}
	case "azure-native:containerregistry:Task":
		r = &Task{}
	case "azure-native:containerregistry:TaskRun":
		r = &TaskRun{}
	case "azure-native:containerregistry:Token":
		r = &Token{}
	case "azure-native:containerregistry:Webhook":
		r = &Webhook{}
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
		"containerregistry",
		&module{version},
	)
}
