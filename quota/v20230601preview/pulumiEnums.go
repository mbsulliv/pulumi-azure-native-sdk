// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

// GroupingId type. It is a required property. More types of groupIds can be supported in future. MGID is already in the URI, so it's not needed.'
type GroupingIdType string

const (
	GroupingIdTypeServiceTreeId = GroupingIdType("ServiceTreeId")
	GroupingIdTypeBillingId     = GroupingIdType("BillingId")
)

func init() {
}