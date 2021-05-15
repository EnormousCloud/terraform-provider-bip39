package core

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tyler-smith/go-bip39"
)

func resourceRandomEntropy() *schema.Resource {
	return &schema.Resource{
		Create: onRandomEntropyCreate,
		Read:   onRandomEntropyRead,
		Update: onRandomEntropyUpdate,
		Delete: onRandomEntropyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"bit_size": {
				Type:        schema.TypeInt,
				Required:    true,
				Default:     128,
				Description: "bit size for the entropy has to be a multiple 32 and be within the inclusive range of (128, 256).",
			},
			"mnemonic": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "list of human-readable words of the seed",
			},
			"entropy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hexadecimal entropy representation",
			},
		},
	}
}

func onRandomEntropyCreate(d *schema.ResourceData, m interface{}) error {
	bitSize := d.Get("bit_size").(int)
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return err
	}
	d.Set("entropy", fmt.Sprintf("%x", entropy))

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}
	d.Set("mnemonic", mnemonic)
	return nil
}

func onRandomEntropyRead(d *schema.ResourceData, m interface{}) error {
	// no actions here. only creation does something
	return nil
}

func onRandomEntropyUpdate(d *schema.ResourceData, m interface{}) error {
	// no actions here. only creation does something
	return nil
}

func onRandomEntropyDelete(d *schema.ResourceData, m interface{}) error {
	// no actions here. only creation does something
	return nil
}
