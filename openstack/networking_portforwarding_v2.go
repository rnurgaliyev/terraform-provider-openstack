package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/portforwarding"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func networkingPortForwardingV2StateRefreshFunc(client *gophercloud.ServiceClient, fipID, pfID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pf, err := portforwarding.Get(client, fipID, pfID).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return pf, "DELETED", nil
			}

			return nil, "", err
		}

		return pf, "ACTIVE", nil
	}
}
