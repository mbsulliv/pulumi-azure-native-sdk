// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get formula.
func LookupFormula(ctx *pulumi.Context, args *LookupFormulaArgs, opts ...pulumi.InvokeOption) (*LookupFormulaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFormulaResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getFormula", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFormulaArgs struct {
	// Specify the $expand query. Example: 'properties($select=description)'
	Expand *string `pulumi:"expand"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the formula.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A formula for creating a VM, specifying an image base and other parameters
type LookupFormulaResult struct {
	// The author of the formula.
	Author string `pulumi:"author"`
	// The creation date of the formula.
	CreationDate string `pulumi:"creationDate"`
	// The description of the formula.
	Description *string `pulumi:"description"`
	// The content of the formula.
	FormulaContent *LabVirtualMachineCreationParameterResponse `pulumi:"formulaContent"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The OS type of the formula.
	OsType *string `pulumi:"osType"`
	// The provisioning status of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
	// Information about a VM from which a formula is to be created.
	Vm *FormulaPropertiesFromVmResponse `pulumi:"vm"`
}

// Defaults sets the appropriate defaults for LookupFormulaResult
func (val *LookupFormulaResult) Defaults() *LookupFormulaResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FormulaContent = tmp.FormulaContent.Defaults()

	return &tmp
}

func LookupFormulaOutput(ctx *pulumi.Context, args LookupFormulaOutputArgs, opts ...pulumi.InvokeOption) LookupFormulaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFormulaResult, error) {
			args := v.(LookupFormulaArgs)
			r, err := LookupFormula(ctx, &args, opts...)
			var s LookupFormulaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFormulaResultOutput)
}

type LookupFormulaOutputArgs struct {
	// Specify the $expand query. Example: 'properties($select=description)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the formula.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFormulaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFormulaArgs)(nil)).Elem()
}

// A formula for creating a VM, specifying an image base and other parameters
type LookupFormulaResultOutput struct{ *pulumi.OutputState }

func (LookupFormulaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFormulaResult)(nil)).Elem()
}

func (o LookupFormulaResultOutput) ToLookupFormulaResultOutput() LookupFormulaResultOutput {
	return o
}

func (o LookupFormulaResultOutput) ToLookupFormulaResultOutputWithContext(ctx context.Context) LookupFormulaResultOutput {
	return o
}

// The author of the formula.
func (o LookupFormulaResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.Author }).(pulumi.StringOutput)
}

// The creation date of the formula.
func (o LookupFormulaResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The description of the formula.
func (o LookupFormulaResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The content of the formula.
func (o LookupFormulaResultOutput) FormulaContent() LabVirtualMachineCreationParameterResponsePtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *LabVirtualMachineCreationParameterResponse { return v.FormulaContent }).(LabVirtualMachineCreationParameterResponsePtrOutput)
}

// The identifier of the resource.
func (o LookupFormulaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource.
func (o LookupFormulaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupFormulaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.Name }).(pulumi.StringOutput)
}

// The OS type of the formula.
func (o LookupFormulaResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o LookupFormulaResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o LookupFormulaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFormulaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupFormulaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupFormulaResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// Information about a VM from which a formula is to be created.
func (o LookupFormulaResultOutput) Vm() FormulaPropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *FormulaPropertiesFromVmResponse { return v.Vm }).(FormulaPropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFormulaResultOutput{})
}
