// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const EssentialContactsContactAssetType string = "essentialcontacts.googleapis.com/Contact"

func resourceConverterEssentialContactsContact() ResourceConverter {
	return ResourceConverter{
		AssetType: EssentialContactsContactAssetType,
		Convert:   GetEssentialContactsContactCaiObject,
	}
}

func GetEssentialContactsContactCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//essentialcontacts.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetEssentialContactsContactApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: EssentialContactsContactAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/essentialcontacts/v1/rest",
				DiscoveryName:        "Contact",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetEssentialContactsContactApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	emailProp, err := expandEssentialContactsContactEmail(d.Get("email"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("email"); !tpgresource.IsEmptyValue(reflect.ValueOf(emailProp)) && (ok || !reflect.DeepEqual(v, emailProp)) {
		obj["email"] = emailProp
	}
	notificationCategorySubscriptionsProp, err := expandEssentialContactsContactNotificationCategorySubscriptions(d.Get("notification_category_subscriptions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_category_subscriptions"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationCategorySubscriptionsProp)) && (ok || !reflect.DeepEqual(v, notificationCategorySubscriptionsProp)) {
		obj["notificationCategorySubscriptions"] = notificationCategorySubscriptionsProp
	}
	languageTagProp, err := expandEssentialContactsContactLanguageTag(d.Get("language_tag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("language_tag"); !tpgresource.IsEmptyValue(reflect.ValueOf(languageTagProp)) && (ok || !reflect.DeepEqual(v, languageTagProp)) {
		obj["languageTag"] = languageTagProp
	}

	return obj, nil
}

func expandEssentialContactsContactEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEssentialContactsContactNotificationCategorySubscriptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEssentialContactsContactLanguageTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
