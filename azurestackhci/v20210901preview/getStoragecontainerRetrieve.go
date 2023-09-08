// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets storagecontainers by resource name
func LookupStoragecontainerRetrieve(ctx *pulumi.Context, args *LookupStoragecontainerRetrieveArgs, opts ...pulumi.InvokeOption) (*LookupStoragecontainerRetrieveResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStoragecontainerRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getStoragecontainerRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStoragecontainerRetrieveArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	StoragecontainersName string `pulumi:"storagecontainersName"`
}

// The storage container resource definition.
type LookupStoragecontainerRetrieveResult struct {
	// Amount of space available on the disk in MB
	AvailableSizeMB float64 `pulumi:"availableSizeMB"`
	// Total size of the disk in MB
	ContainerSizeMB  float64                                    `pulumi:"containerSizeMB"`
	ExtendedLocation *StoragecontainersResponseExtendedLocation `pulumi:"extendedLocation"`
	// Resource Id
	Id string `pulumi:"id"`
	// The resource location
	Location string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Path of the storage container on the disk
	Path              *string `pulumi:"path"`
	ProvisioningState *string `pulumi:"provisioningState"`
	// name of the object to be used in moc
	ResourceName *string `pulumi:"resourceName"`
	// storageContainerStatus defines the observed state of storagecontainers
	Status StorageContainerStatusResponse `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupStoragecontainerRetrieveOutput(ctx *pulumi.Context, args LookupStoragecontainerRetrieveOutputArgs, opts ...pulumi.InvokeOption) LookupStoragecontainerRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStoragecontainerRetrieveResult, error) {
			args := v.(LookupStoragecontainerRetrieveArgs)
			r, err := LookupStoragecontainerRetrieve(ctx, &args, opts...)
			var s LookupStoragecontainerRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStoragecontainerRetrieveResultOutput)
}

type LookupStoragecontainerRetrieveOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	StoragecontainersName pulumi.StringInput `pulumi:"storagecontainersName"`
}

func (LookupStoragecontainerRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStoragecontainerRetrieveArgs)(nil)).Elem()
}

// The storage container resource definition.
type LookupStoragecontainerRetrieveResultOutput struct{ *pulumi.OutputState }

func (LookupStoragecontainerRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStoragecontainerRetrieveResult)(nil)).Elem()
}

func (o LookupStoragecontainerRetrieveResultOutput) ToLookupStoragecontainerRetrieveResultOutput() LookupStoragecontainerRetrieveResultOutput {
	return o
}

func (o LookupStoragecontainerRetrieveResultOutput) ToLookupStoragecontainerRetrieveResultOutputWithContext(ctx context.Context) LookupStoragecontainerRetrieveResultOutput {
	return o
}

// Amount of space available on the disk in MB
func (o LookupStoragecontainerRetrieveResultOutput) AvailableSizeMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) float64 { return v.AvailableSizeMB }).(pulumi.Float64Output)
}

// Total size of the disk in MB
func (o LookupStoragecontainerRetrieveResultOutput) ContainerSizeMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) float64 { return v.ContainerSizeMB }).(pulumi.Float64Output)
}

func (o LookupStoragecontainerRetrieveResultOutput) ExtendedLocation() StoragecontainersResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) *StoragecontainersResponseExtendedLocation {
		return v.ExtendedLocation
	}).(StoragecontainersResponseExtendedLocationPtrOutput)
}

// Resource Id
func (o LookupStoragecontainerRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource location
func (o LookupStoragecontainerRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o LookupStoragecontainerRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

// Path of the storage container on the disk
func (o LookupStoragecontainerRetrieveResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LookupStoragecontainerRetrieveResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// name of the object to be used in moc
func (o LookupStoragecontainerRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// storageContainerStatus defines the observed state of storagecontainers
func (o LookupStoragecontainerRetrieveResultOutput) Status() StorageContainerStatusResponseOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) StorageContainerStatusResponse { return v.Status }).(StorageContainerStatusResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupStoragecontainerRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupStoragecontainerRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o LookupStoragecontainerRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStoragecontainerRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStoragecontainerRetrieveResultOutput{})
}