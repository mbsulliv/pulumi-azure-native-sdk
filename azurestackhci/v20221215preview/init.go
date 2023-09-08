// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221215preview

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
	case "azure-native:azurestackhci/v20221215preview:ArcSetting":
		r = &ArcSetting{}
	case "azure-native:azurestackhci/v20221215preview:Cluster":
		r = &Cluster{}
	case "azure-native:azurestackhci/v20221215preview:Extension":
		r = &Extension{}
	case "azure-native:azurestackhci/v20221215preview:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:azurestackhci/v20221215preview:GuestAgent":
		r = &GuestAgent{}
	case "azure-native:azurestackhci/v20221215preview:HybridIdentityMetadatum":
		r = &HybridIdentityMetadatum{}
	case "azure-native:azurestackhci/v20221215preview:MachineExtension":
		r = &MachineExtension{}
	case "azure-native:azurestackhci/v20221215preview:MarketplaceGalleryImage":
		r = &MarketplaceGalleryImage{}
	case "azure-native:azurestackhci/v20221215preview:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:azurestackhci/v20221215preview:StorageContainer":
		r = &StorageContainer{}
	case "azure-native:azurestackhci/v20221215preview:Update":
		r = &Update{}
	case "azure-native:azurestackhci/v20221215preview:UpdateRun":
		r = &UpdateRun{}
	case "azure-native:azurestackhci/v20221215preview:UpdateSummary":
		r = &UpdateSummary{}
	case "azure-native:azurestackhci/v20221215preview:VirtualHardDisk":
		r = &VirtualHardDisk{}
	case "azure-native:azurestackhci/v20221215preview:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:azurestackhci/v20221215preview:VirtualNetwork":
		r = &VirtualNetwork{}
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
		"azurestackhci/v20221215preview",
		&module{version},
	)
}