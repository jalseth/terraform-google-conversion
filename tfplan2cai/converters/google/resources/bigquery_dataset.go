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
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const datasetIdRegexp = `[0-9A-Za-z_]+`

func validateDatasetId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(datasetIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_)", k))
	}

	if len(value) > 1024 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be greater than 1,024 characters", k))
	}

	return
}

func validateDefaultTableExpirationMs(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value < 3600000 {
		errors = append(errors, fmt.Errorf("%q cannot be shorter than 3600000 milliseconds (one hour)", k))
	}

	return
}

const BigQueryDatasetAssetType string = "bigquery.googleapis.com/Dataset"

func resourceConverterBigQueryDataset() ResourceConverter {
	return ResourceConverter{
		AssetType: BigQueryDatasetAssetType,
		Convert:   GetBigQueryDatasetCaiObject,
	}
}

func GetBigQueryDatasetCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//bigquery.googleapis.com/projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetBigQueryDatasetApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: BigQueryDatasetAssetType,
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigquery/v2/rest",
				DiscoveryName:        "Dataset",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetBigQueryDatasetApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	maxTimeTravelHoursProp, err := expandBigQueryDatasetMaxTimeTravelHours(d.Get("max_time_travel_hours"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("max_time_travel_hours"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxTimeTravelHoursProp)) && (ok || !reflect.DeepEqual(v, maxTimeTravelHoursProp)) {
		obj["maxTimeTravelHours"] = maxTimeTravelHoursProp
	}
	accessProp, err := expandBigQueryDatasetAccess(d.Get("access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("access"); !tpgresource.IsEmptyValue(reflect.ValueOf(accessProp)) && (ok || !reflect.DeepEqual(v, accessProp)) {
		obj["access"] = accessProp
	}
	datasetReferenceProp, err := expandBigQueryDatasetDatasetReference(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dataset_reference"); !tpgresource.IsEmptyValue(reflect.ValueOf(datasetReferenceProp)) && (ok || !reflect.DeepEqual(v, datasetReferenceProp)) {
		obj["datasetReference"] = datasetReferenceProp
	}
	defaultTableExpirationMsProp, err := expandBigQueryDatasetDefaultTableExpirationMs(d.Get("default_table_expiration_ms"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_table_expiration_ms"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultTableExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultTableExpirationMsProp)) {
		obj["defaultTableExpirationMs"] = defaultTableExpirationMsProp
	}
	defaultPartitionExpirationMsProp, err := expandBigQueryDatasetDefaultPartitionExpirationMs(d.Get("default_partition_expiration_ms"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_partition_expiration_ms"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultPartitionExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultPartitionExpirationMsProp)) {
		obj["defaultPartitionExpirationMs"] = defaultPartitionExpirationMsProp
	}
	descriptionProp, err := expandBigQueryDatasetDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	friendlyNameProp, err := expandBigQueryDatasetFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("friendly_name"); ok || !reflect.DeepEqual(v, friendlyNameProp) {
		obj["friendlyName"] = friendlyNameProp
	}
	labelsProp, err := expandBigQueryDatasetLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	locationProp, err := expandBigQueryDatasetLocation(d.Get("location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	defaultEncryptionConfigurationProp, err := expandBigQueryDatasetDefaultEncryptionConfiguration(d.Get("default_encryption_configuration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_encryption_configuration"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultEncryptionConfigurationProp)) && (ok || !reflect.DeepEqual(v, defaultEncryptionConfigurationProp)) {
		obj["defaultEncryptionConfiguration"] = defaultEncryptionConfigurationProp
	}
	isCaseInsensitiveProp, err := expandBigQueryDatasetIsCaseInsensitive(d.Get("is_case_insensitive"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_case_insensitive"); !tpgresource.IsEmptyValue(reflect.ValueOf(isCaseInsensitiveProp)) && (ok || !reflect.DeepEqual(v, isCaseInsensitiveProp)) {
		obj["isCaseInsensitive"] = isCaseInsensitiveProp
	}
	defaultCollationProp, err := expandBigQueryDatasetDefaultCollation(d.Get("default_collation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_collation"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultCollationProp)) && (ok || !reflect.DeepEqual(v, defaultCollationProp)) {
		obj["defaultCollation"] = defaultCollationProp
	}

	return obj, nil
}

func expandBigQueryDatasetMaxTimeTravelHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandBigQueryDatasetAccessDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedGroupByEmail, err := expandBigQueryDatasetAccessGroupByEmail(original["group_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroupByEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["groupByEmail"] = transformedGroupByEmail
		}

		transformedRole, err := expandBigQueryDatasetAccessRole(original["role"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRole); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["role"] = transformedRole
		}

		transformedSpecialGroup, err := expandBigQueryDatasetAccessSpecialGroup(original["special_group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSpecialGroup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["specialGroup"] = transformedSpecialGroup
		}

		transformedUserByEmail, err := expandBigQueryDatasetAccessUserByEmail(original["user_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUserByEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["userByEmail"] = transformedUserByEmail
		}

		transformedView, err := expandBigQueryDatasetAccessView(original["view"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedView); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["view"] = transformedView
		}

		transformedDataset, err := expandBigQueryDatasetAccessDataset(original["dataset"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDataset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["dataset"] = transformedDataset
		}

		transformedRoutine, err := expandBigQueryDatasetAccessRoutine(original["routine"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRoutine); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["routine"] = transformedRoutine
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBigQueryDatasetAccessDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessGroupByEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessSpecialGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessUserByEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessView(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessViewDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessViewProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedTableId, err := expandBigQueryDatasetAccessViewTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessViewDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewTableId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDataset, err := expandBigQueryDatasetAccessDatasetDataset(original["dataset"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataset"] = transformedDataset
	}

	transformedTargetTypes, err := expandBigQueryDatasetAccessDatasetTargetTypes(original["target_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetTypes"] = transformedTargetTypes
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessDatasetDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessDatasetDatasetDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessDatasetDatasetProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessDatasetDatasetDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessDatasetDatasetProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessDatasetTargetTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRoutine(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessRoutineDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessRoutineProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedRoutineId, err := expandBigQueryDatasetAccessRoutineRoutineId(original["routine_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRoutineId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["routineId"] = transformedRoutineId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessRoutineDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRoutineProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRoutineRoutineId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDatasetReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedDatasetId, err := expandBigQueryDatasetDatasetReferenceDatasetId(d.Get("dataset_id"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	return transformed, nil
}

func expandBigQueryDatasetDatasetReferenceDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultTableExpirationMs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultPartitionExpirationMs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetFriendlyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBigQueryDatasetLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultEncryptionConfiguration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetIsCaseInsensitive(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultCollation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
