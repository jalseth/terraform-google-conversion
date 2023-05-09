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

const MLEngineModelAssetType string = "ml.googleapis.com/Model"

func resourceConverterMLEngineModel() ResourceConverter {
	return ResourceConverter{
		AssetType: MLEngineModelAssetType,
		Convert:   GetMLEngineModelCaiObject,
	}
}

func GetMLEngineModelCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//ml.googleapis.com/projects/{{project}}/models/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetMLEngineModelApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: MLEngineModelAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/ml/v1/rest",
				DiscoveryName:        "Model",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetMLEngineModelApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandMLEngineModelName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandMLEngineModelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	defaultVersionProp, err := expandMLEngineModelDefaultVersion(d.Get("default_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultVersionProp)) && (ok || !reflect.DeepEqual(v, defaultVersionProp)) {
		obj["defaultVersion"] = defaultVersionProp
	}
	regionsProp, err := expandMLEngineModelRegions(d.Get("regions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("regions"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionsProp)) && (ok || !reflect.DeepEqual(v, regionsProp)) {
		obj["regions"] = regionsProp
	}
	onlinePredictionLoggingProp, err := expandMLEngineModelOnlinePredictionLogging(d.Get("online_prediction_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("online_prediction_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(onlinePredictionLoggingProp)) && (ok || !reflect.DeepEqual(v, onlinePredictionLoggingProp)) {
		obj["onlinePredictionLogging"] = onlinePredictionLoggingProp
	}
	onlinePredictionConsoleLoggingProp, err := expandMLEngineModelOnlinePredictionConsoleLogging(d.Get("online_prediction_console_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("online_prediction_console_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(onlinePredictionConsoleLoggingProp)) && (ok || !reflect.DeepEqual(v, onlinePredictionConsoleLoggingProp)) {
		obj["onlinePredictionConsoleLogging"] = onlinePredictionConsoleLoggingProp
	}
	labelsProp, err := expandMLEngineModelLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandMLEngineModelName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelDefaultVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandMLEngineModelDefaultVersionName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandMLEngineModelDefaultVersionName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelOnlinePredictionLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelOnlinePredictionConsoleLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
