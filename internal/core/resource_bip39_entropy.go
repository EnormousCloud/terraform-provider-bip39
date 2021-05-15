package core

import (
	"fmt"
	"provider/internal/randstring"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tyler-smith/go-bip39"
)

func resourceBip39Entropy() *schema.Resource {
	return &schema.Resource{
		Create: onBip39EntropyCreate,
		Read:   onBip39EntropyRead,
		Update: onBip39EntropyUpdate,
		Delete: onBip39EntropyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"bit_size": {
				Type:        schema.TypeInt,
				Required:    true,
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

func onBip39EntropyCreate(d *schema.ResourceData, m interface{}) error {
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
	d.SetId(randstring.New(16))
	return nil
}

func onBip39EntropyRead(d *schema.ResourceData, m interface{}) error {
	// no actions here. only creation does something
	return nil
}

func onBip39EntropyUpdate(d *schema.ResourceData, m interface{}) error {
	// no actions here. only creation does something
	return nil
}

func onBip39EntropyDelete(d *schema.ResourceData, m interface{}) error {
	// no actions here. only creation does something
	return nil
}
