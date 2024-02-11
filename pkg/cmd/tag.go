package iemanja

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	iemanjaclient "github.com/ugsto/iemanja-cli/pkg/iemanja_client"
)

func ListTags(client *iemanjaclient.APIClient, limit, offset int) {
	tags, err := client.ListTags(limit, offset)
	if err != nil {
		log.Fatalf("Error listing tags: %v", err)
	}

	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Printf("Total Tags: %d\n\n", tags.Total)

	if err := w.Write([]string{"ID", "Name"}); err != nil {
		log.Fatalln("error writing header to csv:", err)
	}

	for _, tag := range tags.Tags {
		record := []string{tag.ID, tag.Name}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
}

func CreateTag(client *iemanjaclient.APIClient, name string) {
	tag := iemanjaclient.NewTagRequest{Name: name}
	response, err := client.CreateTag(tag)
	if err != nil {
		log.Fatalf("Error creating tag: %v", err)
	}
	fmt.Printf("Tag created successfully:\n\nID: %s,\nName: %s\n", response.Tag.ID, response.Tag.Name)
}

func GetTag(client *iemanjaclient.APIClient, id string) {
	response, err := client.GetTag(id)
	if err != nil {
		log.Fatalf("Error getting tag: %v", err)
	}
	fmt.Printf("Tag retrieved successfully:\n\nID: %s,\nName: %s\n", response.Tag.ID, response.Tag.Name)
}

func UpdateTag(client *iemanjaclient.APIClient, id, name string) {
	tag := iemanjaclient.NewTagRequest{Name: name}
	response, err := client.UpdateTag(id, tag)
	if err != nil {
		log.Fatalf("Error updating tag: %v", err)
	}
	fmt.Printf("Tag updated successfully:\n\nID: %s,\nName: %s\n", response.Tag.ID, response.Tag.Name)
}

func DeleteTag(client *iemanjaclient.APIClient, id string) {
	err := client.DeleteTag(id)
	if err != nil {
		log.Fatalf("Error deleting tag: %v", err)
	}
	fmt.Printf("Tag deleted successfully:\n\nID: %s\n", id)
}
