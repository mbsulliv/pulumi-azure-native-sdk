// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieve a hybrid runbook worker group.
// Azure REST API version: 2022-08-08.
func LookupHybridRunbookWorkerGroup(ctx *pulumi.Context, args *LookupHybridRunbookWorkerGroupArgs, opts ...pulumi.InvokeOption) (*LookupHybridRunbookWorkerGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHybridRunbookWorkerGroupResult
	err := ctx.Invoke("azure-native:automation:getHybridRunbookWorkerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridRunbookWorkerGroupArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The hybrid runbook worker group name
	HybridRunbookWorkerGroupName string `pulumi:"hybridRunbookWorkerGroupName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of hybrid runbook worker group.
type LookupHybridRunbookWorkerGroupResult struct {
	// Sets the credential of a worker group.
	Credential *RunAsCredentialAssociationPropertyResponse `pulumi:"credential"`
	// Type of the HybridWorkerGroup.
	GroupType *string `pulumi:"groupType"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Resource system metadata.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupHybridRunbookWorkerGroupOutput(ctx *pulumi.Context, args LookupHybridRunbookWorkerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupHybridRunbookWorkerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridRunbookWorkerGroupResult, error) {
			args := v.(LookupHybridRunbookWorkerGroupArgs)
			r, err := LookupHybridRunbookWorkerGroup(ctx, &args, opts...)
			var s LookupHybridRunbookWorkerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridRunbookWorkerGroupResultOutput)
}

type LookupHybridRunbookWorkerGroupOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The hybrid runbook worker group name
	HybridRunbookWorkerGroupName pulumi.StringInput `pulumi:"hybridRunbookWorkerGroupName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHybridRunbookWorkerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridRunbookWorkerGroupArgs)(nil)).Elem()
}

// Definition of hybrid runbook worker group.
type LookupHybridRunbookWorkerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupHybridRunbookWorkerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridRunbookWorkerGroupResult)(nil)).Elem()
}

func (o LookupHybridRunbookWorkerGroupResultOutput) ToLookupHybridRunbookWorkerGroupResultOutput() LookupHybridRunbookWorkerGroupResultOutput {
	return o
}

func (o LookupHybridRunbookWorkerGroupResultOutput) ToLookupHybridRunbookWorkerGroupResultOutputWithContext(ctx context.Context) LookupHybridRunbookWorkerGroupResultOutput {
	return o
}

func (o LookupHybridRunbookWorkerGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupHybridRunbookWorkerGroupResult] {
	return pulumix.Output[LookupHybridRunbookWorkerGroupResult]{
		OutputState: o.OutputState,
	}
}

// Sets the credential of a worker group.
func (o LookupHybridRunbookWorkerGroupResultOutput) Credential() RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) *RunAsCredentialAssociationPropertyResponse {
		return v.Credential
	}).(RunAsCredentialAssociationPropertyResponsePtrOutput)
}

// Type of the HybridWorkerGroup.
func (o LookupHybridRunbookWorkerGroupResultOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource Id for the resource
func (o LookupHybridRunbookWorkerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupHybridRunbookWorkerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource system metadata.
func (o LookupHybridRunbookWorkerGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupHybridRunbookWorkerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridRunbookWorkerGroupResultOutput{})
}
