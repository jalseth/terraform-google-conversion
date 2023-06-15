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

package kms

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const KMSKeyRingAssetType string = "cloudkms.googleapis.com/KeyRing"

func ResourceConverterKMSKeyRing() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: KMSKeyRingAssetType,
		Convert:   GetKMSKeyRingCaiObject,
	}
}

func GetKMSKeyRingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//cloudkms.googleapis.com/projects/{{project}}/locations/{{location}}/keyRings/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetKMSKeyRingApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: KMSKeyRingAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "KeyRing",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetKMSKeyRingApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandKMSKeyRingName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationProp, err := expandKMSKeyRingLocation(d.Get("location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}

	return resourceKMSKeyRingEncoder(d, config, obj)
}

func resourceKMSKeyRingEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func expandKMSKeyRingName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSKeyRingLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}