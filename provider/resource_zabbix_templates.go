package provider

import (
	"github.com/dainis/zabbix"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceZabbixTemplates() *schema.Resource {
	return &schema.Resource{
		Create: resourceZabbixTemplatesCreate,
		Read:   resourceZabbixTemplatesRead,
		Update: resourceZabbixTemplatesUpdate,
		Delete: resourceZabbixTemplatesDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Template.",
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
		},
	}
}

func resourceZabbixTemplatesCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	hostGroup := zabbix.HostGroup{
		Name: d.Get("name").(string),
	}
	groups := zabbix.Templates{template}

	err := api.TemplatesCreate(groups)
	if err != nil {
		return err
	}

	groupId := groups[0].GroupId

	log.Printf("Created Template, id is %s", groupId)

	d.Set("group_id", groupId)
	d.SetId(groupId)

	return nil
}

func resourceZabbixTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	log.Printf("Will read templates with id %s", d.Id())

	group, err := api.TemplatesGetById(d.Id())

	if err != nil {
		return err
	}

	d.Set("name", group.Name)

	return nil
}

func resourceZabbixTemplatesUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	hostGroup := zabbix.Templates{
		Name:    d.Get("name").(string),
		GroupId: d.Id(),
	}

	return api.TemplatesUpdate(zabbix.Templates{template})
}

func resourceZabbixTemplatesDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	return api.TemplatesDeleteByIds([]string{d.Id()})
}
