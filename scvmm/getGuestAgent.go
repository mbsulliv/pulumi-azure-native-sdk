// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scvmm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements GuestAgent GET method.
// Azure REST API version: 2022-05-21-preview.
//
// Other available API versions: 2023-04-01-preview.
func LookupGuestAgent(ctx *pulumi.Context, args *LookupGuestAgentArgs, opts ...pulumi.InvokeOption) (*LookupGuestAgentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGuestAgentResult
	err := ctx.Invoke("azure-native:scvmm:getGuestAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestAgentArgs struct {
	// Name of the GuestAgent.
	GuestAgentName string `pulumi:"guestAgentName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the vm.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// Defines the GuestAgent.
type LookupGuestAgentResult struct {
	// Username / Password Credentials to provision guest agent.
	Credentials *GuestCredentialResponse `pulumi:"credentials"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig *HttpProxyConfigurationResponse `pulumi:"httpProxyConfig"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// Gets or sets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets the guest agent status.
	Status string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid string `pulumi:"uuid"`
}

func LookupGuestAgentOutput(ctx *pulumi.Context, args LookupGuestAgentOutputArgs, opts ...pulumi.InvokeOption) LookupGuestAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestAgentResult, error) {
			args := v.(LookupGuestAgentArgs)
			r, err := LookupGuestAgent(ctx, &args, opts...)
			var s LookupGuestAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestAgentResultOutput)
}

type LookupGuestAgentOutputArgs struct {
	// Name of the GuestAgent.
	GuestAgentName pulumi.StringInput `pulumi:"guestAgentName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the vm.
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupGuestAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestAgentArgs)(nil)).Elem()
}

// Defines the GuestAgent.
type LookupGuestAgentResultOutput struct{ *pulumi.OutputState }

func (LookupGuestAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestAgentResult)(nil)).Elem()
}

func (o LookupGuestAgentResultOutput) ToLookupGuestAgentResultOutput() LookupGuestAgentResultOutput {
	return o
}

func (o LookupGuestAgentResultOutput) ToLookupGuestAgentResultOutputWithContext(ctx context.Context) LookupGuestAgentResultOutput {
	return o
}

// Username / Password Credentials to provision guest agent.
func (o LookupGuestAgentResultOutput) Credentials() GuestCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) *GuestCredentialResponse { return v.Credentials }).(GuestCredentialResponsePtrOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupGuestAgentResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// HTTP Proxy configuration for the VM.
func (o LookupGuestAgentResultOutput) HttpProxyConfig() HttpProxyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) *HttpProxyConfigurationResponse { return v.HttpProxyConfig }).(HttpProxyConfigurationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGuestAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGuestAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the guest agent provisioning action.
func (o LookupGuestAgentResultOutput) ProvisioningAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) *string { return v.ProvisioningAction }).(pulumi.StringPtrOutput)
}

// Gets or sets the provisioning state.
func (o LookupGuestAgentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the guest agent status.
func (o LookupGuestAgentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupGuestAgentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGuestAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupGuestAgentResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestAgentResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestAgentResultOutput{})
}
