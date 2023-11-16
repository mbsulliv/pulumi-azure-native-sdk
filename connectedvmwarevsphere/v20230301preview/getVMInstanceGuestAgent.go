// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements GuestAgent GET method.
func LookupVMInstanceGuestAgent(ctx *pulumi.Context, args *LookupVMInstanceGuestAgentArgs, opts ...pulumi.InvokeOption) (*LookupVMInstanceGuestAgentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVMInstanceGuestAgentResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20230301preview:getVMInstanceGuestAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVMInstanceGuestAgentArgs struct {
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri string `pulumi:"resourceUri"`
}

// Defines the GuestAgent.
type LookupVMInstanceGuestAgentResult struct {
	// Username / Password Credentials to provision guest agent.
	Credentials *GuestCredentialResponse `pulumi:"credentials"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig *HttpProxyConfigurationResponse `pulumi:"httpProxyConfig"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource id of the private link scope this machine is assigned to, if any.
	PrivateLinkScopeResourceId *string `pulumi:"privateLinkScopeResourceId"`
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// Gets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets the guest agent status.
	Status string `pulumi:"status"`
	// The resource status information.
	Statuses []ResourceStatusResponse `pulumi:"statuses"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid string `pulumi:"uuid"`
}

func LookupVMInstanceGuestAgentOutput(ctx *pulumi.Context, args LookupVMInstanceGuestAgentOutputArgs, opts ...pulumi.InvokeOption) LookupVMInstanceGuestAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVMInstanceGuestAgentResult, error) {
			args := v.(LookupVMInstanceGuestAgentArgs)
			r, err := LookupVMInstanceGuestAgent(ctx, &args, opts...)
			var s LookupVMInstanceGuestAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVMInstanceGuestAgentResultOutput)
}

type LookupVMInstanceGuestAgentOutputArgs struct {
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupVMInstanceGuestAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVMInstanceGuestAgentArgs)(nil)).Elem()
}

// Defines the GuestAgent.
type LookupVMInstanceGuestAgentResultOutput struct{ *pulumi.OutputState }

func (LookupVMInstanceGuestAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVMInstanceGuestAgentResult)(nil)).Elem()
}

func (o LookupVMInstanceGuestAgentResultOutput) ToLookupVMInstanceGuestAgentResultOutput() LookupVMInstanceGuestAgentResultOutput {
	return o
}

func (o LookupVMInstanceGuestAgentResultOutput) ToLookupVMInstanceGuestAgentResultOutputWithContext(ctx context.Context) LookupVMInstanceGuestAgentResultOutput {
	return o
}

// Username / Password Credentials to provision guest agent.
func (o LookupVMInstanceGuestAgentResultOutput) Credentials() GuestCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) *GuestCredentialResponse { return v.Credentials }).(GuestCredentialResponsePtrOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupVMInstanceGuestAgentResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// HTTP Proxy configuration for the VM.
func (o LookupVMInstanceGuestAgentResultOutput) HttpProxyConfig() HttpProxyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) *HttpProxyConfigurationResponse { return v.HttpProxyConfig }).(HttpProxyConfigurationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupVMInstanceGuestAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVMInstanceGuestAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource id of the private link scope this machine is assigned to, if any.
func (o LookupVMInstanceGuestAgentResultOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) *string { return v.PrivateLinkScopeResourceId }).(pulumi.StringPtrOutput)
}

// Gets or sets the guest agent provisioning action.
func (o LookupVMInstanceGuestAgentResultOutput) ProvisioningAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) *string { return v.ProvisioningAction }).(pulumi.StringPtrOutput)
}

// Gets the provisioning state.
func (o LookupVMInstanceGuestAgentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the guest agent status.
func (o LookupVMInstanceGuestAgentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) string { return v.Status }).(pulumi.StringOutput)
}

// The resource status information.
func (o LookupVMInstanceGuestAgentResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVMInstanceGuestAgentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVMInstanceGuestAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupVMInstanceGuestAgentResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMInstanceGuestAgentResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVMInstanceGuestAgentResultOutput{})
}
