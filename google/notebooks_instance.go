// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

import "reflect"

func GetNotebooksInstanceCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//notebooks.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetNotebooksInstanceApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "notebooks.googleapis.com/Instance",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/notebooks/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetNotebooksInstanceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	machineTypeProp, err := expandNotebooksInstanceMachineType(d.Get("machine_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("machine_type"); !isEmptyValue(reflect.ValueOf(machineTypeProp)) && (ok || !reflect.DeepEqual(v, machineTypeProp)) {
		obj["machineType"] = machineTypeProp
	}
	postStartupScriptProp, err := expandNotebooksInstancePostStartupScript(d.Get("post_startup_script"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("post_startup_script"); !isEmptyValue(reflect.ValueOf(postStartupScriptProp)) && (ok || !reflect.DeepEqual(v, postStartupScriptProp)) {
		obj["postStartupScript"] = postStartupScriptProp
	}
	instanceOwnersProp, err := expandNotebooksInstanceInstanceOwners(d.Get("instance_owners"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance_owners"); !isEmptyValue(reflect.ValueOf(instanceOwnersProp)) && (ok || !reflect.DeepEqual(v, instanceOwnersProp)) {
		obj["instanceOwners"] = instanceOwnersProp
	}
	serviceAccountProp, err := expandNotebooksInstanceServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account"); !isEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	acceleratorConfigProp, err := expandNotebooksInstanceAcceleratorConfig(d.Get("accelerator_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("accelerator_config"); !isEmptyValue(reflect.ValueOf(acceleratorConfigProp)) && (ok || !reflect.DeepEqual(v, acceleratorConfigProp)) {
		obj["acceleratorConfig"] = acceleratorConfigProp
	}
	installGpuDriverProp, err := expandNotebooksInstanceInstallGpuDriver(d.Get("install_gpu_driver"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("install_gpu_driver"); !isEmptyValue(reflect.ValueOf(installGpuDriverProp)) && (ok || !reflect.DeepEqual(v, installGpuDriverProp)) {
		obj["installGpuDriver"] = installGpuDriverProp
	}
	customGpuDriverPathProp, err := expandNotebooksInstanceCustomGpuDriverPath(d.Get("custom_gpu_driver_path"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_gpu_driver_path"); !isEmptyValue(reflect.ValueOf(customGpuDriverPathProp)) && (ok || !reflect.DeepEqual(v, customGpuDriverPathProp)) {
		obj["customGpuDriverPath"] = customGpuDriverPathProp
	}
	bootDiskTypeProp, err := expandNotebooksInstanceBootDiskType(d.Get("boot_disk_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("boot_disk_type"); !isEmptyValue(reflect.ValueOf(bootDiskTypeProp)) && (ok || !reflect.DeepEqual(v, bootDiskTypeProp)) {
		obj["bootDiskType"] = bootDiskTypeProp
	}
	bootDiskSizeGbProp, err := expandNotebooksInstanceBootDiskSizeGb(d.Get("boot_disk_size_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("boot_disk_size_gb"); !isEmptyValue(reflect.ValueOf(bootDiskSizeGbProp)) && (ok || !reflect.DeepEqual(v, bootDiskSizeGbProp)) {
		obj["bootDiskSizeGb"] = bootDiskSizeGbProp
	}
	dataDiskTypeProp, err := expandNotebooksInstanceDataDiskType(d.Get("data_disk_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_disk_type"); !isEmptyValue(reflect.ValueOf(dataDiskTypeProp)) && (ok || !reflect.DeepEqual(v, dataDiskTypeProp)) {
		obj["dataDiskType"] = dataDiskTypeProp
	}
	dataDiskSizeGbProp, err := expandNotebooksInstanceDataDiskSizeGb(d.Get("data_disk_size_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_disk_size_gb"); !isEmptyValue(reflect.ValueOf(dataDiskSizeGbProp)) && (ok || !reflect.DeepEqual(v, dataDiskSizeGbProp)) {
		obj["dataDiskSizeGb"] = dataDiskSizeGbProp
	}
	noRemoveDataDiskProp, err := expandNotebooksInstanceNoRemoveDataDisk(d.Get("no_remove_data_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_remove_data_disk"); !isEmptyValue(reflect.ValueOf(noRemoveDataDiskProp)) && (ok || !reflect.DeepEqual(v, noRemoveDataDiskProp)) {
		obj["noRemoveDataDisk"] = noRemoveDataDiskProp
	}
	diskEncryptionProp, err := expandNotebooksInstanceDiskEncryption(d.Get("disk_encryption"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_encryption"); !isEmptyValue(reflect.ValueOf(diskEncryptionProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionProp)) {
		obj["diskEncryption"] = diskEncryptionProp
	}
	kmsKeyProp, err := expandNotebooksInstanceKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key"); !isEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}
	noPublicIpProp, err := expandNotebooksInstanceNoPublicIp(d.Get("no_public_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_public_ip"); !isEmptyValue(reflect.ValueOf(noPublicIpProp)) && (ok || !reflect.DeepEqual(v, noPublicIpProp)) {
		obj["noPublicIp"] = noPublicIpProp
	}
	noProxyAccessProp, err := expandNotebooksInstanceNoProxyAccess(d.Get("no_proxy_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_proxy_access"); !isEmptyValue(reflect.ValueOf(noProxyAccessProp)) && (ok || !reflect.DeepEqual(v, noProxyAccessProp)) {
		obj["noProxyAccess"] = noProxyAccessProp
	}
	networkProp, err := expandNotebooksInstanceNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetProp, err := expandNotebooksInstanceSubnet(d.Get("subnet"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnet"); !isEmptyValue(reflect.ValueOf(subnetProp)) && (ok || !reflect.DeepEqual(v, subnetProp)) {
		obj["subnet"] = subnetProp
	}
	labelsProp, err := expandNotebooksInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	metadataProp, err := expandNotebooksInstanceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	vmImageProp, err := expandNotebooksInstanceVmImage(d.Get("vm_image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vm_image"); !isEmptyValue(reflect.ValueOf(vmImageProp)) && (ok || !reflect.DeepEqual(v, vmImageProp)) {
		obj["vmImage"] = vmImageProp
	}
	containerImageProp, err := expandNotebooksInstanceContainerImage(d.Get("container_image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("container_image"); !isEmptyValue(reflect.ValueOf(containerImageProp)) && (ok || !reflect.DeepEqual(v, containerImageProp)) {
		obj["containerImage"] = containerImageProp
	}

	return obj, nil
}

func expandNotebooksInstanceMachineType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstancePostStartupScript(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceInstanceOwners(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceServiceAccount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceAcceleratorConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandNotebooksInstanceAcceleratorConfigType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedCoreCount, err := expandNotebooksInstanceAcceleratorConfigCoreCount(original["core_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCoreCount); val.IsValid() && !isEmptyValue(val) {
		transformed["coreCount"] = transformedCoreCount
	}

	return transformed, nil
}

func expandNotebooksInstanceAcceleratorConfigType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceAcceleratorConfigCoreCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceInstallGpuDriver(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceCustomGpuDriverPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceBootDiskType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceBootDiskSizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDataDiskType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDataDiskSizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoRemoveDataDisk(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDiskEncryption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceKmsKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoPublicIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoProxyAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceSubnet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNotebooksInstanceMetadata(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNotebooksInstanceVmImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProject, err := expandNotebooksInstanceVmImageProject(original["project"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProject); val.IsValid() && !isEmptyValue(val) {
		transformed["project"] = transformedProject
	}

	transformedImageFamily, err := expandNotebooksInstanceVmImageImageFamily(original["image_family"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageFamily); val.IsValid() && !isEmptyValue(val) {
		transformed["imageFamily"] = transformedImageFamily
	}

	transformedImageName, err := expandNotebooksInstanceVmImageImageName(original["image_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageName); val.IsValid() && !isEmptyValue(val) {
		transformed["imageName"] = transformedImageName
	}

	return transformed, nil
}

func expandNotebooksInstanceVmImageProject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceVmImageImageFamily(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceVmImageImageName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceContainerImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepository, err := expandNotebooksInstanceContainerImageRepository(original["repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepository); val.IsValid() && !isEmptyValue(val) {
		transformed["repository"] = transformedRepository
	}

	transformedTag, err := expandNotebooksInstanceContainerImageTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandNotebooksInstanceContainerImageRepository(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceContainerImageTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}