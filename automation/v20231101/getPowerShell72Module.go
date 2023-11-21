// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the module identified by module name.
func LookupPowerShell72Module(ctx *pulumi.Context, args *LookupPowerShell72ModuleArgs, opts ...pulumi.InvokeOption) (*LookupPowerShell72ModuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPowerShell72ModuleResult
	err := ctx.Invoke("azure-native:automation/v20231101:getPowerShell72Module", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPowerShell72ModuleArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of module.
	ModuleName string `pulumi:"moduleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the module type.
type LookupPowerShell72ModuleResult struct {
	// Gets the activity count of the module.
	ActivityCount *int `pulumi:"activityCount"`
	// Gets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets the error info of the module.
	Error *ModuleErrorInfoResponse `pulumi:"error"`
	// Gets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// Gets type of module, if its composite or not.
	IsComposite *bool `pulumi:"isComposite"`
	// Gets the isGlobal flag of the module.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the provisioning state of the module.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets the size in bytes of the module.
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Gets the version of the module.
	Version *string `pulumi:"version"`
}

func LookupPowerShell72ModuleOutput(ctx *pulumi.Context, args LookupPowerShell72ModuleOutputArgs, opts ...pulumi.InvokeOption) LookupPowerShell72ModuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPowerShell72ModuleResult, error) {
			args := v.(LookupPowerShell72ModuleArgs)
			r, err := LookupPowerShell72Module(ctx, &args, opts...)
			var s LookupPowerShell72ModuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPowerShell72ModuleResultOutput)
}

type LookupPowerShell72ModuleOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The name of module.
	ModuleName pulumi.StringInput `pulumi:"moduleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPowerShell72ModuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPowerShell72ModuleArgs)(nil)).Elem()
}

// Definition of the module type.
type LookupPowerShell72ModuleResultOutput struct{ *pulumi.OutputState }

func (LookupPowerShell72ModuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPowerShell72ModuleResult)(nil)).Elem()
}

func (o LookupPowerShell72ModuleResultOutput) ToLookupPowerShell72ModuleResultOutput() LookupPowerShell72ModuleResultOutput {
	return o
}

func (o LookupPowerShell72ModuleResultOutput) ToLookupPowerShell72ModuleResultOutputWithContext(ctx context.Context) LookupPowerShell72ModuleResultOutput {
	return o
}

// Gets the activity count of the module.
func (o LookupPowerShell72ModuleResultOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *int { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

// Gets the creation time.
func (o LookupPowerShell72ModuleResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o LookupPowerShell72ModuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets the error info of the module.
func (o LookupPowerShell72ModuleResultOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *ModuleErrorInfoResponse { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

// Gets the etag of the resource.
func (o LookupPowerShell72ModuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource Id for the resource
func (o LookupPowerShell72ModuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets type of module, if its composite or not.
func (o LookupPowerShell72ModuleResultOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *bool { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

// Gets the isGlobal flag of the module.
func (o LookupPowerShell72ModuleResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

// Gets the last modified time.
func (o LookupPowerShell72ModuleResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The Azure Region where the resource lives
func (o LookupPowerShell72ModuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupPowerShell72ModuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the module.
func (o LookupPowerShell72ModuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets the size in bytes of the module.
func (o LookupPowerShell72ModuleResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

// Resource tags.
func (o LookupPowerShell72ModuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupPowerShell72ModuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets the version of the module.
func (o LookupPowerShell72ModuleResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPowerShell72ModuleResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPowerShell72ModuleResultOutput{})
}