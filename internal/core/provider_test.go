package core

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testProviderSimple = `
resource "bip39_entropy" "wallets" {
	bit_size = 128
}
`

var testBip39Providers map[string]terraform.ResourceProvider
var testBip39Provider *schema.Provider

func init() {
	testBip39Provider = Provider().(*schema.Provider)
	testBip39Providers = map[string]terraform.ResourceProvider{
		"bip39": testBip39Provider,
	}
}

func TestBip39Provider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
func TestBip39Provider_ResourceEntropy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testBip39Providers,
		Steps: []resource.TestStep{
			{
				Config: testProviderSimple,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"bip39_entropy.wallets", "mnemonic"),
					resource.TestCheckResourceAttrSet(
						"bip39_entropy.wallets", "entropy"),
				),
			},
		},
	})
}
