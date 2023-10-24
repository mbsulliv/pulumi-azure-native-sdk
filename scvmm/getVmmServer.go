// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scvmm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Implements VMMServer GET method.
// Azure REST API version: 2022-05-21-preview.
//
// Other available API versions: 2023-04-01-preview, 2023-10-07.
func LookupVmmServer(ctx *pulumi.Context, args *LookupVmmServerArgs, opts ...pulumi.InvokeOption) (*LookupVmmServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVmmServerResult
	err := ctx.Invoke("azure-native:scvmm:getVmmServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVmmServerArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the VMMServer.
	VmmServerName string `pulumi:"vmmServerName"`
}

// The VmmServers resource definition.
type LookupVmmServerResult struct {
	// Gets or sets the connection status to the vmmServer.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// Credentials to connect to VMMServer.
	Credentials *VMMServerPropertiesResponseCredentials `pulumi:"credentials"`
	// Gets or sets any error message if connection to vmmServer is having any issue.
	ErrorMessage string `pulumi:"errorMessage"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fqdn is the hostname/ip of the vmmServer.
	Fqdn string `pulumi:"fqdn"`
	// Resource Id
	Id string `pulumi:"id"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Port is the port on which the vmmServer is listening.
	Port *int `pulumi:"port"`
	// Gets or sets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource Type
	Type string `pulumi:"type"`
	// Unique ID of vmmServer.
	Uuid string `pulumi:"uuid"`
	// Version is the version of the vmmSever.
	Version string `pulumi:"version"`
}

func LookupVmmServerOutput(ctx *pulumi.Context, args LookupVmmServerOutputArgs, opts ...pulumi.InvokeOption) LookupVmmServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVmmServerResult, error) {
			args := v.(LookupVmmServerArgs)
			r, err := LookupVmmServer(ctx, &args, opts...)
			var s LookupVmmServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVmmServerResultOutput)
}

type LookupVmmServerOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the VMMServer.
	VmmServerName pulumi.StringInput `pulumi:"vmmServerName"`
}

func (LookupVmmServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmmServerArgs)(nil)).Elem()
}

// The VmmServers resource definition.
type LookupVmmServerResultOutput struct{ *pulumi.OutputState }

func (LookupVmmServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmmServerResult)(nil)).Elem()
}

func (o LookupVmmServerResultOutput) ToLookupVmmServerResultOutput() LookupVmmServerResultOutput {
	return o
}

func (o LookupVmmServerResultOutput) ToLookupVmmServerResultOutputWithContext(ctx context.Context) LookupVmmServerResultOutput {
	return o
}

func (o LookupVmmServerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVmmServerResult] {
	return pulumix.Output[LookupVmmServerResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the connection status to the vmmServer.
func (o LookupVmmServerResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// Credentials to connect to VMMServer.
func (o LookupVmmServerResultOutput) Credentials() VMMServerPropertiesResponseCredentialsPtrOutput {
	return o.ApplyT(func(v LookupVmmServerResult) *VMMServerPropertiesResponseCredentials { return v.Credentials }).(VMMServerPropertiesResponseCredentialsPtrOutput)
}

// Gets or sets any error message if connection to vmmServer is having any issue.
func (o LookupVmmServerResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The extended location.
func (o LookupVmmServerResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVmmServerResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fqdn is the hostname/ip of the vmmServer.
func (o LookupVmmServerResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupVmmServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the location.
func (o LookupVmmServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o LookupVmmServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Port is the port on which the vmmServer is listening.
func (o LookupVmmServerResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVmmServerResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// Gets or sets the provisioning state.
func (o LookupVmmServerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system data.
func (o LookupVmmServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVmmServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupVmmServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVmmServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o LookupVmmServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of vmmServer.
func (o LookupVmmServerResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Version is the version of the vmmSever.
func (o LookupVmmServerResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVmmServerResultOutput{})
}
