package dynadot

import (
	"api/model"
	"fmt"
)

type CreateContactResponse struct {
	CreateContactResponse struct {
		Error                string `json:",omitempty"`
		ResponseCode         string
		CreateContactContent struct {
			ContactId string
		}
		Status string
	}
}

func CreateContact(contact model.ContactInput) (*CreateContactResponse, error) {
	var resp CreateContactResponse
	err := client.sendParams("create_contact", &resp, map[string]string{
		"name":     contact.Name,
		"email":    contact.Email,
		"phonenum": contact.PhoneNumber,
		"phonecc":  contact.PhoneCC,
		"address1": contact.Address,
		"city":     contact.City,
		"zip":      contact.Zip,
		"country":  contact.Country,
	})
	if err != nil {
		return nil, err
	}

	if resp.CreateContactResponse.ResponseCode == "-1" {
		return nil, fmt.Errorf("received error from dynadot api: %s", resp.CreateContactResponse.Error)
	}

	return &resp, nil
}

type DeleteContactResponse struct {
	DeleteContactResponse struct {
		Error        string `json:",omitempty"`
		ResponseCode string
		Status       string
	}
}

func DeleteContact(id string) error {
	var resp DeleteContactResponse
	err := client.sendParams("delete_contact", &resp, map[string]string{"contact_id": id})
	if err != nil {
		return err
	}

	if resp.DeleteContactResponse.ResponseCode == "-1" {
		return fmt.Errorf("received error from dynadot api: %s", resp.DeleteContactResponse.Error)
	}

	return nil
}
