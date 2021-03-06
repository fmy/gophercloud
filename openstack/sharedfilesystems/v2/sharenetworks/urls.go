package sharenetworks

import "github.com/gophercloud/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("share-networks")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("share-networks", id)
}
