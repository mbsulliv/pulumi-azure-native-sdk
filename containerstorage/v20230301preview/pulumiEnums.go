// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

// Reclaim Policy, Delete or Retain
type ReclaimPolicy string

const (
	ReclaimPolicyDelete = ReclaimPolicy("Delete")
	ReclaimPolicyRetain = ReclaimPolicy("Retain")
)

// Indicates how the volumes created from the snapshot should be attached
type VolumeMode string

const (
	VolumeModeFilesystem = VolumeMode("Filesystem")
	VolumeModeRaw        = VolumeMode("Raw")
)

func init() {
}
