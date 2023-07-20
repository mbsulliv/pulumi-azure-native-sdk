// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a DNS zone. Retrieves the zone properties, but not the record sets within the zone.
// Azure REST API version: 2018-05-01.
func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	var rv LookupZoneResult
	err := ctx.Invoke("azure-native:network:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupZoneArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName string `pulumi:"zoneName"`
}

// Describes a DNS zone.
type LookupZoneResult struct {
	// The etag of the zone.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets float64 `pulumi:"maxNumberOfRecordSets"`
	// The maximum number of records per record set that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordsPerRecordSet float64 `pulumi:"maxNumberOfRecordsPerRecordSet"`
	// Resource name.
	Name string `pulumi:"name"`
	// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers []string `pulumi:"nameServers"`
	// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets float64 `pulumi:"numberOfRecordSets"`
	// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
	RegistrationVirtualNetworks []SubResourceResponse `pulumi:"registrationVirtualNetworks"`
	// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
	ResolutionVirtualNetworks []SubResourceResponse `pulumi:"resolutionVirtualNetworks"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The type of this DNS zone (Public or Private).
	ZoneType *string `pulumi:"zoneType"`
}

// Defaults sets the appropriate defaults for LookupZoneResult
func (val *LookupZoneResult) Defaults() *LookupZoneResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.ZoneType == nil {
		zoneType_ := "Public"
		tmp.ZoneType = &zoneType_
	}
	return &tmp
}

func LookupZoneOutput(ctx *pulumi.Context, args LookupZoneOutputArgs, opts ...pulumi.InvokeOption) LookupZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZoneResult, error) {
			args := v.(LookupZoneArgs)
			r, err := LookupZone(ctx, &args, opts...)
			var s LookupZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupZoneResultOutput)
}

type LookupZoneOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName pulumi.StringInput `pulumi:"zoneName"`
}

func (LookupZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneArgs)(nil)).Elem()
}

// Describes a DNS zone.
type LookupZoneResultOutput struct{ *pulumi.OutputState }

func (LookupZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneResult)(nil)).Elem()
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutput() LookupZoneResultOutput {
	return o
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutputWithContext(ctx context.Context) LookupZoneResultOutput {
	return o
}

// The etag of the zone.
func (o LookupZoneResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZoneResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupZoneResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Location }).(pulumi.StringOutput)
}

// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
func (o LookupZoneResultOutput) MaxNumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v LookupZoneResult) float64 { return v.MaxNumberOfRecordSets }).(pulumi.Float64Output)
}

// The maximum number of records per record set that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
func (o LookupZoneResultOutput) MaxNumberOfRecordsPerRecordSet() pulumi.Float64Output {
	return o.ApplyT(func(v LookupZoneResult) float64 { return v.MaxNumberOfRecordsPerRecordSet }).(pulumi.Float64Output)
}

// Resource name.
func (o LookupZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
func (o LookupZoneResultOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
func (o LookupZoneResultOutput) NumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v LookupZoneResult) float64 { return v.NumberOfRecordSets }).(pulumi.Float64Output)
}

// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
func (o LookupZoneResultOutput) RegistrationVirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []SubResourceResponse { return v.RegistrationVirtualNetworks }).(SubResourceResponseArrayOutput)
}

// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
func (o LookupZoneResultOutput) ResolutionVirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []SubResourceResponse { return v.ResolutionVirtualNetworks }).(SubResourceResponseArrayOutput)
}

// Resource tags.
func (o LookupZoneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupZoneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupZoneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Type }).(pulumi.StringOutput)
}

// The type of this DNS zone (Public or Private).
func (o LookupZoneResultOutput) ZoneType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZoneResult) *string { return v.ZoneType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZoneResultOutput{})
}
