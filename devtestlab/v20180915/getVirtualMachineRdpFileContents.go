// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a string that represents the contents of the RDP file for the virtual machine
func GetVirtualMachineRdpFileContents(ctx *pulumi.Context, args *GetVirtualMachineRdpFileContentsArgs, opts ...pulumi.InvokeOption) (*GetVirtualMachineRdpFileContentsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualMachineRdpFileContentsResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getVirtualMachineRdpFileContents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualMachineRdpFileContentsArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the virtual machine.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a .rdp file
type GetVirtualMachineRdpFileContentsResult struct {
	// The contents of the .rdp file
	Contents *string `pulumi:"contents"`
}

func GetVirtualMachineRdpFileContentsOutput(ctx *pulumi.Context, args GetVirtualMachineRdpFileContentsOutputArgs, opts ...pulumi.InvokeOption) GetVirtualMachineRdpFileContentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualMachineRdpFileContentsResult, error) {
			args := v.(GetVirtualMachineRdpFileContentsArgs)
			r, err := GetVirtualMachineRdpFileContents(ctx, &args, opts...)
			var s GetVirtualMachineRdpFileContentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualMachineRdpFileContentsResultOutput)
}

type GetVirtualMachineRdpFileContentsOutputArgs struct {
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the virtual machine.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetVirtualMachineRdpFileContentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMachineRdpFileContentsArgs)(nil)).Elem()
}

// Represents a .rdp file
type GetVirtualMachineRdpFileContentsResultOutput struct{ *pulumi.OutputState }

func (GetVirtualMachineRdpFileContentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMachineRdpFileContentsResult)(nil)).Elem()
}

func (o GetVirtualMachineRdpFileContentsResultOutput) ToGetVirtualMachineRdpFileContentsResultOutput() GetVirtualMachineRdpFileContentsResultOutput {
	return o
}

func (o GetVirtualMachineRdpFileContentsResultOutput) ToGetVirtualMachineRdpFileContentsResultOutputWithContext(ctx context.Context) GetVirtualMachineRdpFileContentsResultOutput {
	return o
}

// The contents of the .rdp file
func (o GetVirtualMachineRdpFileContentsResultOutput) Contents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineRdpFileContentsResult) *string { return v.Contents }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualMachineRdpFileContentsResultOutput{})
}
