// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package policyinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An attestation resource.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2021-01-01.
type AttestationAtResource struct {
	pulumi.CustomResourceState

	// The time the evidence was assessed
	AssessmentDate pulumi.StringPtrOutput `pulumi:"assessmentDate"`
	// Comments describing why this attestation was created.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// The compliance state that should be set on the resource.
	ComplianceState pulumi.StringPtrOutput `pulumi:"complianceState"`
	// The evidence supporting the compliance state set in this attestation.
	Evidence AttestationEvidenceResponseArrayOutput `pulumi:"evidence"`
	// The time the compliance state should expire.
	ExpiresOn pulumi.StringPtrOutput `pulumi:"expiresOn"`
	// The time the compliance state was last changed in this attestation.
	LastComplianceStateChangeAt pulumi.StringOutput `pulumi:"lastComplianceStateChangeAt"`
	// Additional metadata for this attestation
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The person responsible for setting the state of the resource. This value is typically an Azure Active Directory object ID.
	Owner pulumi.StringPtrOutput `pulumi:"owner"`
	// The resource ID of the policy assignment that the attestation is setting the state for.
	PolicyAssignmentId pulumi.StringOutput `pulumi:"policyAssignmentId"`
	// The policy definition reference ID from a policy set definition that the attestation is setting the state for. If the policy assignment assigns a policy set definition the attestation can choose a definition within the set definition with this property or omit this and set the state for the entire set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrOutput `pulumi:"policyDefinitionReferenceId"`
	// The status of the attestation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAttestationAtResource registers a new resource with the given unique name, arguments, and options.
func NewAttestationAtResource(ctx *pulumi.Context,
	name string, args *AttestationAtResourceArgs, opts ...pulumi.ResourceOption) (*AttestationAtResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyAssignmentId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyAssignmentId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights/v20210101:AttestationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20220901:AttestationAtResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AttestationAtResource
	err := ctx.RegisterResource("azure-native:policyinsights:AttestationAtResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestationAtResource gets an existing AttestationAtResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestationAtResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestationAtResourceState, opts ...pulumi.ResourceOption) (*AttestationAtResource, error) {
	var resource AttestationAtResource
	err := ctx.ReadResource("azure-native:policyinsights:AttestationAtResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestationAtResource resources.
type attestationAtResourceState struct {
}

type AttestationAtResourceState struct {
}

func (AttestationAtResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtResourceState)(nil)).Elem()
}

type attestationAtResourceArgs struct {
	// The time the evidence was assessed
	AssessmentDate *string `pulumi:"assessmentDate"`
	// The name of the attestation.
	AttestationName *string `pulumi:"attestationName"`
	// Comments describing why this attestation was created.
	Comments *string `pulumi:"comments"`
	// The compliance state that should be set on the resource.
	ComplianceState *string `pulumi:"complianceState"`
	// The evidence supporting the compliance state set in this attestation.
	Evidence []AttestationEvidence `pulumi:"evidence"`
	// The time the compliance state should expire.
	ExpiresOn *string `pulumi:"expiresOn"`
	// Additional metadata for this attestation
	Metadata interface{} `pulumi:"metadata"`
	// The person responsible for setting the state of the resource. This value is typically an Azure Active Directory object ID.
	Owner *string `pulumi:"owner"`
	// The resource ID of the policy assignment that the attestation is setting the state for.
	PolicyAssignmentId string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID from a policy set definition that the attestation is setting the state for. If the policy assignment assigns a policy set definition the attestation can choose a definition within the set definition with this property or omit this and set the state for the entire set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// Resource ID.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a AttestationAtResource resource.
type AttestationAtResourceArgs struct {
	// The time the evidence was assessed
	AssessmentDate pulumi.StringPtrInput
	// The name of the attestation.
	AttestationName pulumi.StringPtrInput
	// Comments describing why this attestation was created.
	Comments pulumi.StringPtrInput
	// The compliance state that should be set on the resource.
	ComplianceState pulumi.StringPtrInput
	// The evidence supporting the compliance state set in this attestation.
	Evidence AttestationEvidenceArrayInput
	// The time the compliance state should expire.
	ExpiresOn pulumi.StringPtrInput
	// Additional metadata for this attestation
	Metadata pulumi.Input
	// The person responsible for setting the state of the resource. This value is typically an Azure Active Directory object ID.
	Owner pulumi.StringPtrInput
	// The resource ID of the policy assignment that the attestation is setting the state for.
	PolicyAssignmentId pulumi.StringInput
	// The policy definition reference ID from a policy set definition that the attestation is setting the state for. If the policy assignment assigns a policy set definition the attestation can choose a definition within the set definition with this property or omit this and set the state for the entire set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// Resource ID.
	ResourceId pulumi.StringInput
}

func (AttestationAtResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtResourceArgs)(nil)).Elem()
}

type AttestationAtResourceInput interface {
	pulumi.Input

	ToAttestationAtResourceOutput() AttestationAtResourceOutput
	ToAttestationAtResourceOutputWithContext(ctx context.Context) AttestationAtResourceOutput
}

func (*AttestationAtResource) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtResource)(nil)).Elem()
}

func (i *AttestationAtResource) ToAttestationAtResourceOutput() AttestationAtResourceOutput {
	return i.ToAttestationAtResourceOutputWithContext(context.Background())
}

func (i *AttestationAtResource) ToAttestationAtResourceOutputWithContext(ctx context.Context) AttestationAtResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationAtResourceOutput)
}

type AttestationAtResourceOutput struct{ *pulumi.OutputState }

func (AttestationAtResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtResource)(nil)).Elem()
}

func (o AttestationAtResourceOutput) ToAttestationAtResourceOutput() AttestationAtResourceOutput {
	return o
}

func (o AttestationAtResourceOutput) ToAttestationAtResourceOutputWithContext(ctx context.Context) AttestationAtResourceOutput {
	return o
}

// The time the evidence was assessed
func (o AttestationAtResourceOutput) AssessmentDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringPtrOutput { return v.AssessmentDate }).(pulumi.StringPtrOutput)
}

// Comments describing why this attestation was created.
func (o AttestationAtResourceOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// The compliance state that should be set on the resource.
func (o AttestationAtResourceOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringPtrOutput { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

// The evidence supporting the compliance state set in this attestation.
func (o AttestationAtResourceOutput) Evidence() AttestationEvidenceResponseArrayOutput {
	return o.ApplyT(func(v *AttestationAtResource) AttestationEvidenceResponseArrayOutput { return v.Evidence }).(AttestationEvidenceResponseArrayOutput)
}

// The time the compliance state should expire.
func (o AttestationAtResourceOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringPtrOutput { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

// The time the compliance state was last changed in this attestation.
func (o AttestationAtResourceOutput) LastComplianceStateChangeAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringOutput { return v.LastComplianceStateChangeAt }).(pulumi.StringOutput)
}

// Additional metadata for this attestation
func (o AttestationAtResourceOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the resource
func (o AttestationAtResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The person responsible for setting the state of the resource. This value is typically an Azure Active Directory object ID.
func (o AttestationAtResourceOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringPtrOutput { return v.Owner }).(pulumi.StringPtrOutput)
}

// The resource ID of the policy assignment that the attestation is setting the state for.
func (o AttestationAtResourceOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringOutput { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

// The policy definition reference ID from a policy set definition that the attestation is setting the state for. If the policy assignment assigns a policy set definition the attestation can choose a definition within the set definition with this property or omit this and set the state for the entire set definition.
func (o AttestationAtResourceOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

// The status of the attestation.
func (o AttestationAtResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AttestationAtResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttestationAtResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AttestationAtResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestationAtResourceOutput{})
}
