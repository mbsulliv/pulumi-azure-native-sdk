// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a NotificationHub Resource.
type NotificationHub struct {
	pulumi.CustomResourceState

	// The AdmCredential of the created NotificationHub
	AdmCredential AdmCredentialResponsePtrOutput `pulumi:"admCredential"`
	// The ApnsCredential of the created NotificationHub
	ApnsCredential ApnsCredentialResponsePtrOutput `pulumi:"apnsCredential"`
	// The AuthorizationRules of the created NotificationHub
	AuthorizationRules SharedAccessAuthorizationRulePropertiesResponseArrayOutput `pulumi:"authorizationRules"`
	// The BaiduCredential of the created NotificationHub
	BaiduCredential BaiduCredentialResponsePtrOutput `pulumi:"baiduCredential"`
	// The GcmCredential of the created NotificationHub
	GcmCredential GcmCredentialResponsePtrOutput `pulumi:"gcmCredential"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The MpnsCredential of the created NotificationHub
	MpnsCredential MpnsCredentialResponsePtrOutput `pulumi:"mpnsCredential"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The RegistrationTtl of the created NotificationHub
	RegistrationTtl pulumi.StringPtrOutput `pulumi:"registrationTtl"`
	// The sku of the created namespace
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The WnsCredential of the created NotificationHub
	WnsCredential WnsCredentialResponsePtrOutput `pulumi:"wnsCredential"`
}

// NewNotificationHub registers a new resource with the given unique name, arguments, and options.
func NewNotificationHub(ctx *pulumi.Context,
	name string, args *NotificationHubArgs, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:notificationhubs:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20140901:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20230101preview:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20230901:NotificationHub"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NotificationHub
	err := ctx.RegisterResource("azure-native:notificationhubs/v20170401:NotificationHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationHub gets an existing NotificationHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubState, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	var resource NotificationHub
	err := ctx.ReadResource("azure-native:notificationhubs/v20170401:NotificationHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationHub resources.
type notificationHubState struct {
}

type NotificationHubState struct {
}

func (NotificationHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubState)(nil)).Elem()
}

type notificationHubArgs struct {
	// The AdmCredential of the created NotificationHub
	AdmCredential *AdmCredential `pulumi:"admCredential"`
	// The ApnsCredential of the created NotificationHub
	ApnsCredential *ApnsCredential `pulumi:"apnsCredential"`
	// The AuthorizationRules of the created NotificationHub
	AuthorizationRules []SharedAccessAuthorizationRuleProperties `pulumi:"authorizationRules"`
	// The BaiduCredential of the created NotificationHub
	BaiduCredential *BaiduCredential `pulumi:"baiduCredential"`
	// The GcmCredential of the created NotificationHub
	GcmCredential *GcmCredential `pulumi:"gcmCredential"`
	// Resource location
	Location *string `pulumi:"location"`
	// The MpnsCredential of the created NotificationHub
	MpnsCredential *MpnsCredential `pulumi:"mpnsCredential"`
	// The NotificationHub name.
	Name *string `pulumi:"name"`
	// The namespace name.
	NamespaceName string `pulumi:"namespaceName"`
	// The notification hub name.
	NotificationHubName *string `pulumi:"notificationHubName"`
	// The RegistrationTtl of the created NotificationHub
	RegistrationTtl *string `pulumi:"registrationTtl"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku of the created namespace
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The WnsCredential of the created NotificationHub
	WnsCredential *WnsCredential `pulumi:"wnsCredential"`
}

// The set of arguments for constructing a NotificationHub resource.
type NotificationHubArgs struct {
	// The AdmCredential of the created NotificationHub
	AdmCredential AdmCredentialPtrInput
	// The ApnsCredential of the created NotificationHub
	ApnsCredential ApnsCredentialPtrInput
	// The AuthorizationRules of the created NotificationHub
	AuthorizationRules SharedAccessAuthorizationRulePropertiesArrayInput
	// The BaiduCredential of the created NotificationHub
	BaiduCredential BaiduCredentialPtrInput
	// The GcmCredential of the created NotificationHub
	GcmCredential GcmCredentialPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The MpnsCredential of the created NotificationHub
	MpnsCredential MpnsCredentialPtrInput
	// The NotificationHub name.
	Name pulumi.StringPtrInput
	// The namespace name.
	NamespaceName pulumi.StringInput
	// The notification hub name.
	NotificationHubName pulumi.StringPtrInput
	// The RegistrationTtl of the created NotificationHub
	RegistrationTtl pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The sku of the created namespace
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The WnsCredential of the created NotificationHub
	WnsCredential WnsCredentialPtrInput
}

func (NotificationHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubArgs)(nil)).Elem()
}

type NotificationHubInput interface {
	pulumi.Input

	ToNotificationHubOutput() NotificationHubOutput
	ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput
}

func (*NotificationHub) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHub)(nil)).Elem()
}

func (i *NotificationHub) ToNotificationHubOutput() NotificationHubOutput {
	return i.ToNotificationHubOutputWithContext(context.Background())
}

func (i *NotificationHub) ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubOutput)
}

type NotificationHubOutput struct{ *pulumi.OutputState }

func (NotificationHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHub)(nil)).Elem()
}

func (o NotificationHubOutput) ToNotificationHubOutput() NotificationHubOutput {
	return o
}

func (o NotificationHubOutput) ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput {
	return o
}

// The AdmCredential of the created NotificationHub
func (o NotificationHubOutput) AdmCredential() AdmCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) AdmCredentialResponsePtrOutput { return v.AdmCredential }).(AdmCredentialResponsePtrOutput)
}

// The ApnsCredential of the created NotificationHub
func (o NotificationHubOutput) ApnsCredential() ApnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) ApnsCredentialResponsePtrOutput { return v.ApnsCredential }).(ApnsCredentialResponsePtrOutput)
}

// The AuthorizationRules of the created NotificationHub
func (o NotificationHubOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *NotificationHub) SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesResponseArrayOutput)
}

// The BaiduCredential of the created NotificationHub
func (o NotificationHubOutput) BaiduCredential() BaiduCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) BaiduCredentialResponsePtrOutput { return v.BaiduCredential }).(BaiduCredentialResponsePtrOutput)
}

// The GcmCredential of the created NotificationHub
func (o NotificationHubOutput) GcmCredential() GcmCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) GcmCredentialResponsePtrOutput { return v.GcmCredential }).(GcmCredentialResponsePtrOutput)
}

// Resource location
func (o NotificationHubOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The MpnsCredential of the created NotificationHub
func (o NotificationHubOutput) MpnsCredential() MpnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) MpnsCredentialResponsePtrOutput { return v.MpnsCredential }).(MpnsCredentialResponsePtrOutput)
}

// Resource name
func (o NotificationHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The RegistrationTtl of the created NotificationHub
func (o NotificationHubOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringPtrOutput { return v.RegistrationTtl }).(pulumi.StringPtrOutput)
}

// The sku of the created namespace
func (o NotificationHubOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o NotificationHubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o NotificationHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The WnsCredential of the created NotificationHub
func (o NotificationHubOutput) WnsCredential() WnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) WnsCredentialResponsePtrOutput { return v.WnsCredential }).(WnsCredentialResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationHubOutput{})
}
