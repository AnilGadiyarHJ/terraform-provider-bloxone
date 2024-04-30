package anycast_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/infobloxopen/bloxone-go-client/anycast"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccOnPremAnycastHostResource_basic(t *testing.T) {
	var resourceName = "bloxone_anycast_host.test"
	var v anycast.ProtoOnpremHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_disappears(t *testing.T) {
	resourceName := "bloxone_anycast_host.test"
	var v anycast.ProtoOnpremHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnPremAnycastHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOnPremAnycastHostBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					testAccCheckOnPremAnycastHostDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOnPremAnycastHostResource_enableRouting(t *testing.T) {
	var resourceName = "bloxone_anycast_host.test"
	var v anycast.ProtoOnpremHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostEnableRoutingBGP("BGP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.asn", "6500"),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.holddown_secs", "180"),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.neighbors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.neighbors.0.asn", "6501"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostEnableRoutingOSPF("OSPF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.area_type", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.area", "10.10.0.1"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.authentication_type", "Clear"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.interface", "eth0"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.authentication_key", "YXV0aGV"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.hello_interval", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_BGP(t *testing.T) {
	var resourceName = "bloxone_anycast_host.test"
	var v anycast.ProtoOnpremHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostBGP("BGP", 6500, 180),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.asn", "6500"),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.holddown_secs", "180"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostBGP("BGP", 6601, 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.asn", "6601"),
					resource.TestCheckResourceAttr(resourceName, "config_bgp.holddown_secs", "200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_OSPF(t *testing.T) {
	var resourceName = "bloxone_anycast_host.test"
	var v anycast.ProtoOnpremHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostOSPF("OSPF", "STANDARD", "10.10.0.1", "Clear", "eth0", 10, 40, 5, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.area_type", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.area", "10.10.0.1"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.authentication_type", "Clear"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.interface", "eth0"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.authentication_key", "YXV0aGV"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.hello_interval", "10"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.dead_interval", "40"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.retransmit_interval", "5"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.transmit_delay", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostOSPF("OSPF", "NSSA", "10.10.0.2", "MD5", "ens160", 20, 50, 10, 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.area_type", "NSSA"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.area", "10.10.0.2"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.authentication_type", "MD5"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.interface", "ens160"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.authentication_key", "YXV0aGV"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.hello_interval", "20"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.dead_interval", "50"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.retransmit_interval", "10"),
					resource.TestCheckResourceAttr(resourceName, "config_ospf.transmit_delay", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckOnPremAnycastHostExists(ctx context.Context, resourceName string, v *anycast.ProtoOnpremHost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing ID: %v", err)
		}
		apiRes, _, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			OnPremAnycastManagerGetOnpremHost(ctx, id). //OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResults() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResults()
		return nil
	}
}

func testAccCheckOnPremAnycastHostDestroy(ctx context.Context, v *anycast.ProtoOnpremHost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			OnPremAnycastManagerGetOnpremHost(ctx, *v.Id). //OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckOnPremAnycastHostDisappears(ctx context.Context, v *anycast.ProtoOnpremHost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, _, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			OnPremAnycastManagerDeleteOnpremHost(ctx, *v.Id). //testAccCheckOnPremAnycastHostDisappears
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithAnycastConfig() string {
	return fmt.Sprintf(`
data "bloxone_infra_hosts" "anycast_hosts" {
  filters = {
    display_name = "anycast_real"
  }
}

resource "bloxone_anycast_config" "test_onprem_hosts" {
  anycast_ip_address = "10.10.0.1"
  name               = "Anycast_Config_Example"
  service            = "DNS"

}
`)
}

func testAccOnPremAnycastHostBasicConfig() string {
	config := fmt.Sprintf(`
resource "bloxone_anycast_host" "test" {
  id = one(data.bloxone_infra_hosts.anycast_hosts.results).legacy_id

 anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.test_onprem_hosts.name
    }
  ]
}
`)
	return strings.Join([]string{testAccBaseWithAnycastConfig(), config}, "")
}

func testAccOnPremAnycastHostEnableRoutingBGP(routing_protocols string) string {
	config := fmt.Sprintf(`
resource "bloxone_anycast_host" "test" {
  id = one(data.bloxone_infra_hosts.anycast_hosts.results).legacy_id

 anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.test_onprem_hosts.name
      routing_protocols   = ["%s"]
    }
  ]

 config_bgp = {
     asn       = "6500"
     holddown_secs = 180
     neighbors       = [
       {
        asn       = "6501"
        ip_address = "172.28.4.198"
       }
     ]
   }
 }
`, routing_protocols)
	return strings.Join([]string{testAccBaseWithAnycastConfig(), config}, "")
}

func testAccOnPremAnycastHostEnableRoutingOSPF(routing_protocols string) string {
	config := fmt.Sprintf(`
resource "bloxone_anycast_host" "test" {
  id = one(data.bloxone_infra_hosts.anycast_hosts.results).legacy_id

 anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.test_onprem_hosts.name
      routing_protocols   = ["%s"]
    }
  ]

  config_ospf = {
    area_type       = "STANDARD"
    area            = "10.10.0.1"
    authentication_type = "Clear"
    interface       = "eth0"
    authentication_key = "YXV0aGV"
    hello_interval = 10
    dead_interval = 40
    retransmit_interval = 5
    transmit_delay = 1
 }
}
`, routing_protocols)
	return strings.Join([]string{testAccBaseWithAnycastConfig(), config}, "")
}

func testAccOnPremAnycastHostBGP(routing_protocols string, asn, holddown_secs int64) string {
	config := fmt.Sprintf(`
resource "bloxone_anycast_host" "test" {
  id = one(data.bloxone_infra_hosts.anycast_hosts.results).legacy_id

 anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.test_onprem_hosts.name
      routing_protocols   = ["%s"]
    }
  ]

 config_bgp = {
     asn       = "%d"
     holddown_secs = "%d"
     neighbors       = [
       {
        asn       = "6501"
        ip_address = "172.28.4.198"
       }
     ]
   }
 }
`, routing_protocols, asn, holddown_secs)
	return strings.Join([]string{testAccBaseWithAnycastConfig(), config}, "")
}

func testAccOnPremAnycastHostOSPF(routing_protocols, area_type, area, authentication_type, ospfInterface string, hello_interval, dead_interval, retransmit_interval, transmit_delay int64) string {
	config := fmt.Sprintf(`
resource "bloxone_anycast_host" "test" {
  id = one(data.bloxone_infra_hosts.anycast_hosts.results).legacy_id

 anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.test_onprem_hosts.name
      routing_protocols   = ["%s"]
    }
  ]

  config_ospf = {
    area_type       = "%s"
    area            = "%s"
    authentication_type = "%s"
    interface       = "%s"
    authentication_key = "YXV0aGV"
	authentication_key_id = "1"
    hello_interval = "%d"
    dead_interval = "%d"
    retransmit_interval ="%d"
    transmit_delay = "%d"
 }
}
`, routing_protocols, area_type, area, authentication_type, ospfInterface, hello_interval, dead_interval, retransmit_interval, transmit_delay)
	return strings.Join([]string{testAccBaseWithAnycastConfig(), config}, "")
}
