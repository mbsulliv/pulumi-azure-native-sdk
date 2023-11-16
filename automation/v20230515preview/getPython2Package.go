// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230515preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the python 2 package identified by package name.
func LookupPython2Package(ctx *pulumi.Context, args *LookupPython2PackageArgs, opts ...pulumi.InvokeOption) (*LookupPython2PackageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPython2PackageResult
	err := ctx.Invoke("azure-native:automation/v20230515preview:getPython2Package", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPython2PackageArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The python package name.
	PackageName string `pulumi:"packageName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the module type.
type LookupPython2PackageResult struct {
	// Gets or sets the activity count of the module.
	ActivityCount *int `pulumi:"activityCount"`
	// Gets or sets the contentLink of the module.
	ContentLink *ContentLinkResponse `pulumi:"contentLink"`
	// Gets or sets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets or sets the error info of the module.
	Error *ModuleErrorInfoResponse `pulumi:"error"`
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Gets or sets type of module, if its composite or not.
	IsComposite *bool `pulumi:"isComposite"`
	// Gets or sets the isGlobal flag of the module.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the provisioning state of the module.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets the size in bytes of the module.
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Gets or sets the version of the module.
	Version *string `pulumi:"version"`
}

func LookupPython2PackageOutput(ctx *pulumi.Context, args LookupPython2PackageOutputArgs, opts ...pulumi.InvokeOption) LookupPython2PackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPython2PackageResult, error) {
			args := v.(LookupPython2PackageArgs)
			r, err := LookupPython2Package(ctx, &args, opts...)
			var s LookupPython2PackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPython2PackageResultOutput)
}

type LookupPython2PackageOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The python package name.
	PackageName pulumi.StringInput `pulumi:"packageName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPython2PackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPython2PackageArgs)(nil)).Elem()
}

// Definition of the module type.
type LookupPython2PackageResultOutput struct{ *pulumi.OutputState }

func (LookupPython2PackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPython2PackageResult)(nil)).Elem()
}

func (o LookupPython2PackageResultOutput) ToLookupPython2PackageResultOutput() LookupPython2PackageResultOutput {
	return o
}

func (o LookupPython2PackageResultOutput) ToLookupPython2PackageResultOutputWithContext(ctx context.Context) LookupPython2PackageResultOutput {
	return o
}

// Gets or sets the activity count of the module.
func (o LookupPython2PackageResultOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *int { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

// Gets or sets the contentLink of the module.
func (o LookupPython2PackageResultOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *ContentLinkResponse { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

// Gets or sets the creation time.
func (o LookupPython2PackageResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o LookupPython2PackageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the error info of the module.
func (o LookupPython2PackageResultOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *ModuleErrorInfoResponse { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

// Gets or sets the etag of the resource.
func (o LookupPython2PackageResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupPython2PackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets type of module, if its composite or not.
func (o LookupPython2PackageResultOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *bool { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

// Gets or sets the isGlobal flag of the module.
func (o LookupPython2PackageResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o LookupPython2PackageResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupPython2PackageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPython2PackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state of the module.
func (o LookupPython2PackageResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets the size in bytes of the module.
func (o LookupPython2PackageResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPython2PackageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPython2PackageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPython2PackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the version of the module.
func (o LookupPython2PackageResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPython2PackageResultOutput{})
}
