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
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const WorkflowsWorkflowAssetType string = "workflows.googleapis.com/Workflow"

func resourceConverterWorkflowsWorkflow() ResourceConverter {
	return ResourceConverter{
		AssetType: WorkflowsWorkflowAssetType,
		Convert:   GetWorkflowsWorkflowCaiObject,
	}
}

func GetWorkflowsWorkflowCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//workflows.googleapis.com/projects/{{project}}/locations/{{region}}/workflows/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetWorkflowsWorkflowApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: WorkflowsWorkflowAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/workflows/v1/rest",
				DiscoveryName:        "Workflow",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetWorkflowsWorkflowApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandWorkflowsWorkflowName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandWorkflowsWorkflowDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandWorkflowsWorkflowLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	serviceAccountProp, err := expandWorkflowsWorkflowServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	sourceContentsProp, err := expandWorkflowsWorkflowSourceContents(d.Get("source_contents"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_contents"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceContentsProp)) && (ok || !reflect.DeepEqual(v, sourceContentsProp)) {
		obj["sourceContents"] = sourceContentsProp
	}
	cryptoKeyNameProp, err := expandWorkflowsWorkflowCryptoKeyName(d.Get("crypto_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("crypto_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(cryptoKeyNameProp)) && (ok || !reflect.DeepEqual(v, cryptoKeyNameProp)) {
		obj["cryptoKeyName"] = cryptoKeyNameProp
	}

	return resourceWorkflowsWorkflowEncoder(d, config, obj)
}

func resourceWorkflowsWorkflowEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	var ResName string
	if v, ok := d.GetOk("name"); ok {
		ResName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		ResName = resource.PrefixedUniqueId(v.(string))
	} else {
		ResName = resource.UniqueId()
	}

	if err := d.Set("name", ResName); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return obj, nil
}

func expandWorkflowsWorkflowName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkflowsWorkflowServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowSourceContents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowCryptoKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
