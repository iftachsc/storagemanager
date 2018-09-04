package zadara_test

import (
	
	"github.com/iftachsc/storagemanager/zadara"
)

const (
	MockVolume1 = zadara.Volume{
		Name: "volume-00000001",
		PoolName: "pool-00000001",
		Encryption:  "NO"
	}
	
	RootVolumeResponseJson = `{
		"response": {
		  "status": 0,
		  "volumes": [
			{
			  "name": "volume-00000001",
			  "display_name": "volume",
			  "cg_display_name": "volume",
			  "cg_name": "cg-00000001",
			  "cg_user_created": "NO",
			  "pool_display_name": "poolrenamed",
			  "pool_name": "pool-00000001",
			  "status": "Available",
			  "virtual_capacity": 200,
			  "server_name": null,
			  "data_type": "BLOCK",
			  "thin": "YES",
			  "encryption": "NO",
			  "encryption_key_size": "0",
			  "allocated_capacity": 0.014,
			  "data_copies_capacity": 0.0,
			  "access_type": null,
			  "read_only": null,
			  "export_name": null,
			  "nfs_root_squash": "YES",
			  "mount_sync": "NO",
			  "atime_update": "NO",
			  "read_ahead_kb": "16",
			  "target": null,
			  "lun": null,
			  "cache": null,
			  "smb_acl": null,
			  "smb_guest": null,
			  "smb_only": null,
			  "smb_browseable": "YES",
			  "smb_hidden_files": null,
			  "smb_hide_unreadable": "NO",
			  "smb_hide_unwriteable": "NO",
			  "smb_hide_dot_files": "NO",
			  "smb_encryption_mode": "off",
			  "smb_enable_oplocks": "YES",
			  "read_iops_limit": "0",
			  "read_mbps_limit": "0",
			  "write_iops_limit": "0",
			  "write_mbps_limit": "0",
			  "uquota_state": "on",
			  "gquota_state": "off",
			  "pquota_state": "off",
			  "created_at": "2013-09-26 00:42:30 UTC",
			  "modified_at": "2013-09-26 00:42:31 UTC"
			},
			{
			  "name": "volume-00000002",
			  "display_name": "share",
			  "cg_display_name": "share",
			  "cg_name": "cg-00000002",
			  "cg_user_created": "NO",
			  "pool_display_name": "poolrenamed",
			  "pool_name": "pool-00000001",
			  "status": "Available",
			  "virtual_capacity": 200,
			  "server_name": null,
			  "data_type": "File-System",
			  "thin": "NO",
			  "encryption": "NO",
			  "encryption_key_size": "0",
			  "allocated_capacity": 0.014,
			  "data_copies_capacity": 0.0,
			  "access_type": null,
			  "has_backup": "YES",
			  "has_mirror": "YES",
			  "has_snapshots": "YES",
			  "read_only": null,
			  "export_name": "/export/share",
			  "nfs_export_path": "192.168.1.1:/export/share",
			  "smb_export_path": "\\\\192.168.1.1\\share",
			  "nfs_root_squash": "YES",
			  "mount_sync": "YES",
			  "atime_update": "NO",
			  "read_ahead_kb": "16",
			  "target": null,
			  "lun": null,
			  "cache": null,
			  "smb_acl": "YES",
			  "smb_guest": "YES",
			  "smb_only": "NO",
			  "smb_browseable": "YES",
			  "smb_hidden_files": null,
			  "smb_hide_unreadable": "NO",
			  "smb_hide_unwriteable": "NO",
			  "smb_hide_dot_files": "NO",
			  "smb_encryption_mode": "off",
			  "smb_enable_oplocks": "YES",
			  "smb_store_dos_attributes": "NO",
			  "read_iops_limit": "0",
			  "read_mbps_limit": "0",
			  "write_iops_limit": "0",
			  "write_mbps_limit": "0",
			  "uquota_state": "on",
			  "gquota_state": "off",
			  "pquota_state": "off",
			  "created_at": "2013-09-26 00:52:00 UTC",
			  "modified_at": "2013-09-26 00:52:05 UTC"
			}
		  ],
		  "count": 2
		}
	  }`
)
