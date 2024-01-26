// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a container on the  Data Box Edge/Gateway device.
func LookupContainer(ctx *pulumi.Context, args *LookupContainerArgs, opts ...pulumi.InvokeOption) (*LookupContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerResult
	err := ctx.Invoke("azure-native:databoxedge/v20231201:getContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerArgs struct {
	// The container Name
	ContainerName string `pulumi:"containerName"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Storage Account Name
	StorageAccountName string `pulumi:"storageAccountName"`
}

// Represents a container on the  Data Box Edge/Gateway device.
type LookupContainerResult struct {
	// Current status of the container.
	ContainerStatus string `pulumi:"containerStatus"`
	// The UTC time when container got created.
	CreatedDateTime string `pulumi:"createdDateTime"`
	// DataFormat for Container
	DataFormat string `pulumi:"dataFormat"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The object name.
	Name string `pulumi:"name"`
	// Details of the refresh job on this container.
	RefreshDetails RefreshDetailsResponse `pulumi:"refreshDetails"`
	// Metadata pertaining to creation and last modification of Container
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupContainerOutput(ctx *pulumi.Context, args LookupContainerOutputArgs, opts ...pulumi.InvokeOption) LookupContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerResult, error) {
			args := v.(LookupContainerArgs)
			r, err := LookupContainer(ctx, &args, opts...)
			var s LookupContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerResultOutput)
}

type LookupContainerOutputArgs struct {
	// The container Name
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Storage Account Name
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

func (LookupContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerArgs)(nil)).Elem()
}

// Represents a container on the  Data Box Edge/Gateway device.
type LookupContainerResultOutput struct{ *pulumi.OutputState }

func (LookupContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerResult)(nil)).Elem()
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutput() LookupContainerResultOutput {
	return o
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutputWithContext(ctx context.Context) LookupContainerResultOutput {
	return o
}

// Current status of the container.
func (o LookupContainerResultOutput) ContainerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.ContainerStatus }).(pulumi.StringOutput)
}

// The UTC time when container got created.
func (o LookupContainerResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// DataFormat for Container
func (o LookupContainerResultOutput) DataFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.DataFormat }).(pulumi.StringOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The object name.
func (o LookupContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Details of the refresh job on this container.
func (o LookupContainerResultOutput) RefreshDetails() RefreshDetailsResponseOutput {
	return o.ApplyT(func(v LookupContainerResult) RefreshDetailsResponse { return v.RefreshDetails }).(RefreshDetailsResponseOutput)
}

// Metadata pertaining to creation and last modification of Container
func (o LookupContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerResultOutput{})
}
