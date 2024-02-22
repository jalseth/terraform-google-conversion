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

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecuritySecurityProfileGroupAssetType string = "networksecurity.googleapis.com/SecurityProfileGroup"

func ResourceConverterNetworkSecuritySecurityProfileGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecuritySecurityProfileGroupAssetType,
		Convert:   GetNetworkSecuritySecurityProfileGroupCaiObject,
	}
}

func GetNetworkSecuritySecurityProfileGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecuritySecurityProfileGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecuritySecurityProfileGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "SecurityProfileGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecuritySecurityProfileGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecuritySecurityProfileGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	threatPreventionProfileProp, err := expandNetworkSecuritySecurityProfileGroupThreatPreventionProfile(d.Get("threat_prevention_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("threat_prevention_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(threatPreventionProfileProp)) && (ok || !reflect.DeepEqual(v, threatPreventionProfileProp)) {
		obj["threatPreventionProfile"] = threatPreventionProfileProp
	}
	labelsProp, err := expandNetworkSecuritySecurityProfileGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkSecuritySecurityProfileGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileGroupThreatPreventionProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}