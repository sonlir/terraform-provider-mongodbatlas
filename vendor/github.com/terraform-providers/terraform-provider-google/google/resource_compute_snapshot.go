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

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/customdiff"
	"github.com/hashicorp/terraform/helper/schema"
	compute "google.golang.org/api/compute/v1"
)

func customDiffComputeSnapshotSnapshotEncryptionKeys(diff *schema.ResourceDiff, meta interface{}) error {
	oldConvenience, newConvenience := diff.GetChange("snapshot_encryption_key_raw")
	oldNewField, newNewField := diff.GetChange("snapshot_encryption_key.0.raw_key")

	if newConvenience != "" && newNewField != "" {
		return fmt.Errorf("can't use snapshot_encryption_key_raw and snapshot_encryption_key.0.raw_key at the same time." +
			"If you're removing snapshot_encryption_key.0.raw_key, set the value to \"\" instead. This is due to limitations in Terraform.")
	}

	// Either field (convenience or new) has a value
	// and then has another different value, so we ForceNew.
	// We need to handle _EVERY_ ForceNew case in this diff
	if oldConvenience != "" && newConvenience != "" && oldConvenience != newConvenience {
		return diff.ForceNew("snapshot_encryption_key_raw")
	}

	if oldNewField != "" && newNewField != "" && oldNewField != newNewField {
		return diff.ForceNew("snapshot_encryption_key.0.raw_key")
	}

	// Our resource isn't using either field, then uses one;
	// ForceNew on whichever one is now using it.
	if (oldConvenience == "" && oldNewField == "" && newConvenience != "") || (oldConvenience == "" && oldNewField == "" && newNewField != "") {
		if oldConvenience == "" && newConvenience != "" {
			return diff.ForceNew("snapshot_encryption_key_raw")
		} else {
			return diff.ForceNew("snapshot_encryption_key.0.raw_key")
		}
	}

	// convenience no longer used
	if oldConvenience != "" && newConvenience == "" {
		if newNewField == "" {
			// convenience is being nulled, and the new field is empty as well
			// we've stopped using the field altogether
			return diff.ForceNew("snapshot_encryption_key_raw")
		} else if oldConvenience != newNewField {
			// convenience is being nulled, and the new field has a new value
			// so we ForceNew on either field
			return diff.ForceNew("snapshot_encryption_key_raw")
		} else {
			// If we reach it here, we're using the same value in the new field as we had in the convenience field
		}
	}

	// new no longer used
	// note that it will remain _set_ because of how Computed fields work
	// unset fields will have their values kept in state as a non-zero value
	if oldNewField != "" && newNewField == "" {
		if newConvenience == "" {
			// new field is being nulled, and the convenience field is empty as well
			// we've stopped using the field altogether
			return diff.ForceNew("snapshot_encryption_key.0.raw_key")
		} else if oldNewField != newConvenience {
			// new is being nulled, and the convenience field has a new value
			// so we ForceNew on either field

			// This stops a really opaque diffs don't match during apply error. Without this, wee see
			// a diff from the old state -> new state with a ForceNew at plan time (as expected!)
			// But during apply time the entire nested object is nil in old state unexpectedly.
			// So we just force the diff to match more by nilling it here, which is unclear why it
			// works, and probably a worse UX with some real ugly diff, but also makes the tests pass.
			// Computed nested fields are hard.
			err := diff.SetNew("snapshot_encryption_key", nil)
			if err != nil {
				return err
			}

			return diff.ForceNew("snapshot_encryption_key.0.raw_key")
		} else {
			// If we reach it here, we're using the same value in the convenience field as we had in the new field
		}
	}

	return nil
}

func customDiffComputeSnapshotSourceDiskEncryptionKeys(diff *schema.ResourceDiff, meta interface{}) error {
	oldConvenience, newConvenience := diff.GetChange("source_disk_encryption_key_raw")
	oldNewField, newNewField := diff.GetChange("source_disk_encryption_key.0.raw_key")

	// Either field has a value and then has another value
	// We need to handle _EVERY_ ForceNew case in this diff
	if oldConvenience != "" && newConvenience != "" && oldConvenience != newConvenience {
		return diff.ForceNew("source_disk_encryption_key_raw")
	}

	if oldNewField != "" && newNewField != "" && oldNewField != newNewField {
		return diff.ForceNew("source_disk_encryption_key.0.raw_key")
	}

	// Our resource isn't using either field, then uses one;
	// ForceNew on whichever one is now using it.
	if (oldConvenience == "" && oldNewField == "" && newConvenience != "") || (oldConvenience == "" && oldNewField == "" && newNewField != "") {
		if oldConvenience == "" && newConvenience != "" {
			return diff.ForceNew("source_disk_encryption_key_raw")
		} else {
			return diff.ForceNew("source_disk_encryption_key.0.raw_key")
		}
	}

	// convenience no longer used
	if oldConvenience != "" && newConvenience == "" {
		if newNewField == "" {
			// convenience is being nulled, and the new field is empty as well
			// we've stopped using the field altogether
			return diff.ForceNew("source_disk_encryption_key_raw")
		} else if oldConvenience != newNewField {
			// convenience is being nulled, and the new field has a new value
			// so we ForceNew on either field
			return diff.ForceNew("source_disk_encryption_key_raw")
		} else {
			// If we reach it here, we're using the same value in the new field as we had in the convenience field
		}
	}

	// new no longer used
	if oldNewField != "" && newNewField == "" {
		if newConvenience == "" {
			// new field is being nulled, and the convenience field is empty as well
			// we've stopped using the field altogether
			return diff.ForceNew("source_disk_encryption_key.0.raw_key")
		} else if newConvenience != oldNewField {
			// new is being nulled, and the convenience field has a new value
			// so we ForceNew on either field
			return diff.ForceNew("source_disk_encryption_key.0.raw_key")
		} else {
			// If we reach it here, we're using the same value in the convenience field as we had in the new field
		}
	}

	return nil
}

func resourceComputeSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSnapshotCreate,
		Read:   resourceComputeSnapshotRead,
		Update: resourceComputeSnapshotUpdate,
		Delete: resourceComputeSnapshotDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSnapshotImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(300 * time.Second),
			Update: schema.DefaultTimeout(300 * time.Second),
			Delete: schema.DefaultTimeout(300 * time.Second),
		},
		CustomizeDiff: customdiff.All(
			customDiffComputeSnapshotSnapshotEncryptionKeys,
			customDiffComputeSnapshotSourceDiskEncryptionKeys,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_disk": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_encryption_key": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"raw_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"source_disk_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"raw_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
					},
				},
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disk_size_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"licenses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"snapshot_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"storage_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"source_disk_link": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"snapshot_encryption_key_raw": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: "Use snapshot_encryption_key.raw_key instead.",
			},

			"snapshot_encryption_key_sha256": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "Use snapshot_encryption_key.sha256 instead.",
			},

			"source_disk_encryption_key_raw": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: "Use source_disk_encryption_key.raw_key instead.",
			},

			"source_disk_encryption_key_sha256": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "Use source_disk_encryption_key.sha256 instead.",
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeSnapshotName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandComputeSnapshotLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	sourceDiskProp, err := expandComputeSnapshotSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_disk"); !isEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	zoneProp, err := expandComputeSnapshotZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	snapshotEncryptionKeyProp, err := expandComputeSnapshotSnapshotEncryptionKey(d.Get("snapshot_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("snapshot_encryption_key"); !isEmptyValue(reflect.ValueOf(snapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, snapshotEncryptionKeyProp)) {
		obj["snapshotEncryptionKey"] = snapshotEncryptionKeyProp
	}
	sourceDiskEncryptionKeyProp, err := expandComputeSnapshotSourceDiskEncryptionKey(d.Get("source_disk_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_disk_encryption_key"); !isEmptyValue(reflect.ValueOf(sourceDiskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceDiskEncryptionKeyProp)) {
		obj["sourceDiskEncryptionKey"] = sourceDiskEncryptionKeyProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/disks/{{source_disk}}/createSnapshot")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Snapshot: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Snapshot: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating Snapshot",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Snapshot: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Snapshot %q: %#v", d.Id(), res)

	return resourceComputeSnapshotRead(d, meta)
}

func resourceComputeSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeSnapshot %q", d.Id()))
	}

	res, err = resourceComputeSnapshotDecoder(d, meta, res)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeSnapshotCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("snapshot_id", flattenComputeSnapshotSnapshot_id(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("disk_size_gb", flattenComputeSnapshotDiskSizeGb(res["diskSizeGb"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("name", flattenComputeSnapshotName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("description", flattenComputeSnapshotDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("storage_bytes", flattenComputeSnapshotStorageBytes(res["storageBytes"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("licenses", flattenComputeSnapshotLicenses(res["licenses"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("labels", flattenComputeSnapshotLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeSnapshotLabelFingerprint(res["labelFingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("source_disk", flattenComputeSnapshotSourceDisk(res["sourceDisk"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("snapshot_encryption_key", flattenComputeSnapshotSnapshotEncryptionKey(res["snapshotEncryptionKey"], d)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}

	return nil
}

func resourceComputeSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	if d.HasChange("labels") || d.HasChange("label_fingerprint") {
		obj := make(map[string]interface{})
		labelsProp, err := expandComputeSnapshotLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}
		labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/snapshots/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err := sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating Snapshot %q: %s", d.Id(), err)
		}

		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating Snapshot",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("labels")
		d.SetPartial("label_fingerprint")
	}

	d.Partial(false)

	return resourceComputeSnapshotRead(d, meta)
}

func resourceComputeSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Snapshot %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "Snapshot")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting Snapshot",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Snapshot %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeSnapshotImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/global/snapshots/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSnapshotCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSnapshotSnapshot_id(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeSnapshotDiskSizeGb(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeSnapshotName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSnapshotDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSnapshotStorageBytes(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeSnapshotLicenses(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeSnapshotLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSnapshotLabelFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSnapshotSourceDisk(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeSnapshotSnapshotEncryptionKey(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeSnapshotSnapshotEncryptionKeyRawKey(original["rawKey"], d)
	transformed["sha256"] =
		flattenComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d)
	return []interface{}{transformed}
}
func flattenComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d *schema.ResourceData) interface{} {
	return d.Get("snapshot_encryption_key.0.raw_key")
}

func flattenComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeSnapshotName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeSnapshotLabelFingerprint(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDisk(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseZonalFieldValue("disks", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_disk: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotZone(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotSnapshotEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, 1)
	if len(l) == 1 && l[0].(map[string]interface{})["raw_key"] != "" {
		// There is a value
		outMap := make(map[string]interface{})
		outMap["rawKey"] = l[0].(map[string]interface{})["raw_key"]
		req = append(req, outMap)
	} else {
		// Check alternative setting?
		if altV, ok := d.GetOk("snapshot_encryption_key_raw"); ok && altV != "" {
			outMap := make(map[string]interface{})
			outMap["rawKey"] = altV
			req = append(req, outMap)
		}
	}
	return req, nil
}

func expandComputeSnapshotSourceDiskEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, 1)
	if len(l) == 1 {
		// There is a value
		outMap := make(map[string]interface{})
		outMap["rawKey"] = l[0].(map[string]interface{})["raw_key"]
		req = append(req, outMap)
	} else {
		// Check alternative setting?
		if altV, ok := d.GetOk("source_disk_encryption_key_raw"); ok && altV != "" {
			outMap := make(map[string]interface{})
			outMap["rawKey"] = altV
			req = append(req, outMap)
		}
	}
	return req, nil
}

func resourceComputeSnapshotDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	d.Set("source_disk_link", ConvertSelfLinkToV1(res["sourceDisk"].(string)))
	if snapshotEncryptionKey := res["snapshotEncryptionKey"]; snapshotEncryptionKey != nil {
		d.Set("snapshot_encryption_key_sha256", snapshotEncryptionKey.((map[string]interface{}))["sha256"])
	}

	return res, nil
}