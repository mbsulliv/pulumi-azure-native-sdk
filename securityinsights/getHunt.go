// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a hunt, without relations and comments.
// Azure REST API version: 2023-06-01-preview.
func LookupHunt(ctx *pulumi.Context, args *LookupHuntArgs, opts ...pulumi.InvokeOption) (*LookupHuntResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHuntResult
	err := ctx.Invoke("azure-native:securityinsights:getHunt", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupHuntArgs struct {
	// The hunt id (GUID)
	HuntId string `pulumi:"huntId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Hunt in Azure Security Insights.
type LookupHuntResult struct {
	// A list of mitre attack tactics the hunt is associated with
	AttackTactics []string `pulumi:"attackTactics"`
	// A list of a mitre attack techniques the hunt is associated with
	AttackTechniques []string `pulumi:"attackTechniques"`
	// The description of the hunt
	Description string `pulumi:"description"`
	// The display name of the hunt
	DisplayName string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The hypothesis status of the hunt.
	HypothesisStatus *string `pulumi:"hypothesisStatus"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// List of labels relevant to this hunt
	Labels []string `pulumi:"labels"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Describes a user that the hunt is assigned to
	Owner *HuntOwnerResponse `pulumi:"owner"`
	// The status of the hunt.
	Status *string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupHuntResult
func (val *LookupHuntResult) Defaults() *LookupHuntResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.HypothesisStatus == nil {
		hypothesisStatus_ := "Unknown"
		tmp.HypothesisStatus = &hypothesisStatus_
	}
	if tmp.Status == nil {
		status_ := "New"
		tmp.Status = &status_
	}
	return &tmp
}

func LookupHuntOutput(ctx *pulumi.Context, args LookupHuntOutputArgs, opts ...pulumi.InvokeOption) LookupHuntResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHuntResult, error) {
			args := v.(LookupHuntArgs)
			r, err := LookupHunt(ctx, &args, opts...)
			var s LookupHuntResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHuntResultOutput)
}

type LookupHuntOutputArgs struct {
	// The hunt id (GUID)
	HuntId pulumi.StringInput `pulumi:"huntId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupHuntOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHuntArgs)(nil)).Elem()
}

// Represents a Hunt in Azure Security Insights.
type LookupHuntResultOutput struct{ *pulumi.OutputState }

func (LookupHuntResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHuntResult)(nil)).Elem()
}

func (o LookupHuntResultOutput) ToLookupHuntResultOutput() LookupHuntResultOutput {
	return o
}

func (o LookupHuntResultOutput) ToLookupHuntResultOutputWithContext(ctx context.Context) LookupHuntResultOutput {
	return o
}

func (o LookupHuntResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupHuntResult] {
	return pulumix.Output[LookupHuntResult]{
		OutputState: o.OutputState,
	}
}

// A list of mitre attack tactics the hunt is associated with
func (o LookupHuntResultOutput) AttackTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHuntResult) []string { return v.AttackTactics }).(pulumi.StringArrayOutput)
}

// A list of a mitre attack techniques the hunt is associated with
func (o LookupHuntResultOutput) AttackTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHuntResult) []string { return v.AttackTechniques }).(pulumi.StringArrayOutput)
}

// The description of the hunt
func (o LookupHuntResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name of the hunt
func (o LookupHuntResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o LookupHuntResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHuntResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The hypothesis status of the hunt.
func (o LookupHuntResultOutput) HypothesisStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHuntResult) *string { return v.HypothesisStatus }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHuntResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of labels relevant to this hunt
func (o LookupHuntResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHuntResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupHuntResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes a user that the hunt is assigned to
func (o LookupHuntResultOutput) Owner() HuntOwnerResponsePtrOutput {
	return o.ApplyT(func(v LookupHuntResult) *HuntOwnerResponse { return v.Owner }).(HuntOwnerResponsePtrOutput)
}

// The status of the hunt.
func (o LookupHuntResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHuntResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupHuntResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHuntResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHuntResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHuntResultOutput{})
}
