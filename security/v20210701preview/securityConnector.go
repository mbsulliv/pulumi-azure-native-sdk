// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The security connector resource.
type SecurityConnector struct {
	pulumi.CustomResourceState

	// The multi cloud resource's cloud name.
	CloudName pulumi.StringPtrOutput `pulumi:"cloudName"`
	// Entity tag is used for comparing two or more entities from the same requested resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The multi cloud resource identifier (account id in case of AWS connector).
	HierarchyIdentifier pulumi.StringPtrOutput `pulumi:"hierarchyIdentifier"`
	// Kind of the resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Location where the resource is stored
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// A collection of offerings for the security connector.
	Offerings pulumi.ArrayOutput `pulumi:"offerings"`
	// The multi cloud account's organizational data
	OrganizationalData SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput `pulumi:"organizationalData"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// A list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityConnector registers a new resource with the given unique name, arguments, and options.
func NewSecurityConnector(ctx *pulumi.Context,
	name string, args *SecurityConnectorArgs, opts ...pulumi.ResourceOption) (*SecurityConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:SecurityConnector"),
		},
		{
			Type: pulumi.String("azure-native:security/v20211201preview:SecurityConnector"),
		},
		{
			Type: pulumi.String("azure-native:security/v20220501preview:SecurityConnector"),
		},
		{
			Type: pulumi.String("azure-native:security/v20220801preview:SecurityConnector"),
		},
		{
			Type: pulumi.String("azure-native:security/v20230301preview:SecurityConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecurityConnector
	err := ctx.RegisterResource("azure-native:security/v20210701preview:SecurityConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityConnector gets an existing SecurityConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityConnectorState, opts ...pulumi.ResourceOption) (*SecurityConnector, error) {
	var resource SecurityConnector
	err := ctx.ReadResource("azure-native:security/v20210701preview:SecurityConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityConnector resources.
type securityConnectorState struct {
}

type SecurityConnectorState struct {
}

func (SecurityConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorState)(nil)).Elem()
}

type securityConnectorArgs struct {
	// The multi cloud resource's cloud name.
	CloudName *string `pulumi:"cloudName"`
	// The multi cloud resource identifier (account id in case of AWS connector).
	HierarchyIdentifier *string `pulumi:"hierarchyIdentifier"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// A collection of offerings for the security connector.
	Offerings []interface{} `pulumi:"offerings"`
	// The multi cloud account's organizational data
	OrganizationalData *SecurityConnectorPropertiesOrganizationalData `pulumi:"organizationalData"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName *string `pulumi:"securityConnectorName"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SecurityConnector resource.
type SecurityConnectorArgs struct {
	// The multi cloud resource's cloud name.
	CloudName pulumi.StringPtrInput
	// The multi cloud resource identifier (account id in case of AWS connector).
	HierarchyIdentifier pulumi.StringPtrInput
	// Kind of the resource
	Kind pulumi.StringPtrInput
	// Location where the resource is stored
	Location pulumi.StringPtrInput
	// A collection of offerings for the security connector.
	Offerings pulumi.ArrayInput
	// The multi cloud account's organizational data
	OrganizationalData SecurityConnectorPropertiesOrganizationalDataPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The security connector name.
	SecurityConnectorName pulumi.StringPtrInput
	// A list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
}

func (SecurityConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorArgs)(nil)).Elem()
}

type SecurityConnectorInput interface {
	pulumi.Input

	ToSecurityConnectorOutput() SecurityConnectorOutput
	ToSecurityConnectorOutputWithContext(ctx context.Context) SecurityConnectorOutput
}

func (*SecurityConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnector)(nil)).Elem()
}

func (i *SecurityConnector) ToSecurityConnectorOutput() SecurityConnectorOutput {
	return i.ToSecurityConnectorOutputWithContext(context.Background())
}

func (i *SecurityConnector) ToSecurityConnectorOutputWithContext(ctx context.Context) SecurityConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorOutput)
}

func (i *SecurityConnector) ToOutput(ctx context.Context) pulumix.Output[*SecurityConnector] {
	return pulumix.Output[*SecurityConnector]{
		OutputState: i.ToSecurityConnectorOutputWithContext(ctx).OutputState,
	}
}

type SecurityConnectorOutput struct{ *pulumi.OutputState }

func (SecurityConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnector)(nil)).Elem()
}

func (o SecurityConnectorOutput) ToSecurityConnectorOutput() SecurityConnectorOutput {
	return o
}

func (o SecurityConnectorOutput) ToSecurityConnectorOutputWithContext(ctx context.Context) SecurityConnectorOutput {
	return o
}

func (o SecurityConnectorOutput) ToOutput(ctx context.Context) pulumix.Output[*SecurityConnector] {
	return pulumix.Output[*SecurityConnector]{
		OutputState: o.OutputState,
	}
}

// The multi cloud resource's cloud name.
func (o SecurityConnectorOutput) CloudName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringPtrOutput { return v.CloudName }).(pulumi.StringPtrOutput)
}

// Entity tag is used for comparing two or more entities from the same requested resource.
func (o SecurityConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The multi cloud resource identifier (account id in case of AWS connector).
func (o SecurityConnectorOutput) HierarchyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringPtrOutput { return v.HierarchyIdentifier }).(pulumi.StringPtrOutput)
}

// Kind of the resource
func (o SecurityConnectorOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o SecurityConnectorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o SecurityConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A collection of offerings for the security connector.
func (o SecurityConnectorOutput) Offerings() pulumi.ArrayOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.ArrayOutput { return v.Offerings }).(pulumi.ArrayOutput)
}

// The multi cloud account's organizational data
func (o SecurityConnectorOutput) OrganizationalData() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o.ApplyT(func(v *SecurityConnector) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
		return v.OrganizationalData
	}).(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SecurityConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecurityConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of key value pairs that describe the resource.
func (o SecurityConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o SecurityConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityConnectorOutput{})
}
