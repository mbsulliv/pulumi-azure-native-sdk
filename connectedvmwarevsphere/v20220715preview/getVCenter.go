// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220715preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Implements vCenter GET method.
func LookupVCenter(ctx *pulumi.Context, args *LookupVCenterArgs, opts ...pulumi.InvokeOption) (*LookupVCenterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVCenterResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220715preview:getVCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVCenterArgs struct {
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the vCenter.
	VcenterName string `pulumi:"vcenterName"`
}

// Defines the vCenter.
type LookupVCenterResult struct {
	// Gets or sets the connection status to the vCenter.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// Username / Password Credentials to connect to vcenter.
	Credentials *VICredentialResponse `pulumi:"credentials"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Gets or sets the FQDN/IPAddress of the vCenter.
	Fqdn string `pulumi:"fqdn"`
	// Gets or sets the Id.
	Id string `pulumi:"id"`
	// Gets or sets the instance UUID of the vCenter.
	InstanceUuid string `pulumi:"instanceUuid"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Gets or sets the port of the vCenter.
	Port *int `pulumi:"port"`
	// Gets or sets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource status information.
	Statuses []ResourceStatusResponse `pulumi:"statuses"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid string `pulumi:"uuid"`
	// Gets or sets the version of the vCenter.
	Version string `pulumi:"version"`
}

func LookupVCenterOutput(ctx *pulumi.Context, args LookupVCenterOutputArgs, opts ...pulumi.InvokeOption) LookupVCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVCenterResult, error) {
			args := v.(LookupVCenterArgs)
			r, err := LookupVCenter(ctx, &args, opts...)
			var s LookupVCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVCenterResultOutput)
}

type LookupVCenterOutputArgs struct {
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the vCenter.
	VcenterName pulumi.StringInput `pulumi:"vcenterName"`
}

func (LookupVCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVCenterArgs)(nil)).Elem()
}

// Defines the vCenter.
type LookupVCenterResultOutput struct{ *pulumi.OutputState }

func (LookupVCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVCenterResult)(nil)).Elem()
}

func (o LookupVCenterResultOutput) ToLookupVCenterResultOutput() LookupVCenterResultOutput {
	return o
}

func (o LookupVCenterResultOutput) ToLookupVCenterResultOutputWithContext(ctx context.Context) LookupVCenterResultOutput {
	return o
}

func (o LookupVCenterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVCenterResult] {
	return pulumix.Output[LookupVCenterResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the connection status to the vCenter.
func (o LookupVCenterResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// Username / Password Credentials to connect to vcenter.
func (o LookupVCenterResultOutput) Credentials() VICredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *VICredentialResponse { return v.Credentials }).(VICredentialResponsePtrOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupVCenterResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the extended location.
func (o LookupVCenterResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the FQDN/IPAddress of the vCenter.
func (o LookupVCenterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Gets or sets the Id.
func (o LookupVCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the instance UUID of the vCenter.
func (o LookupVCenterResultOutput) InstanceUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.InstanceUuid }).(pulumi.StringOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupVCenterResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o LookupVCenterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the name.
func (o LookupVCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the port of the vCenter.
func (o LookupVCenterResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// Gets or sets the provisioning state.
func (o LookupVCenterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o LookupVCenterResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVCenterResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o LookupVCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o LookupVCenterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVCenterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o LookupVCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupVCenterResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the version of the vCenter.
func (o LookupVCenterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVCenterResultOutput{})
}
