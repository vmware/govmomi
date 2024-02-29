package simulator

import (
	"time"

	kmip "github.com/smira/go-kmip"
)

// CreateKey sends a new key creation request to server
func (c *client) CreateKey() (serverResp interface{}, err error) {
	var resp interface{}

	RequestPayload :=
		kmip.CreateRequest{
			ObjectType: kmip.OBJECT_TYPE_SYMMETRIC_KEY,
			TemplateAttribute: kmip.TemplateAttribute{
				Attributes: []kmip.Attribute{
					{
						Name:  kmip.ATTRIBUTE_NAME_CRYPTOGRAPHIC_ALGORITHM,
						Value: kmip.CRYPTO_AES,
					},
					{
						Name:  kmip.ATTRIBUTE_NAME_CRYPTOGRAPHIC_LENGTH,
						Value: int32(128),
					},
					{
						Name:  kmip.ATTRIBUTE_NAME_CRYPTOGRAPHIC_USAGE_MASK,
						Value: int32(12),
					},
					{
						Name:  kmip.ATTRIBUTE_NAME_INITIAL_DATE,
						Value: time.Unix(12345, 0),
					},
					// {
					// 	Name: kmip.ATTRIBUTE_NAME_NAME,
					// 	Value: kmip.Name{
					// 		Tag:   kmip.NAME_TYPE,
					// 		Value: "some field",
					// 		Type:  kmip.Enum(kmip.NAME_TYPE),
					// 	},
					// },
				},
			},
		}
	resp, err = c.kclient.Send(kmip.OPERATION_CREATE, RequestPayload)

	return resp, err
}

// GetKey retrieves the key from keyvault
func (c *client) GetKey(id string) (serverResp interface{}, err error) {
	var resp interface{}
	RequestPayload :=
		kmip.GetRequest{
			UniqueIdentifier: id,
		}
	resp, err = c.kclient.Send(kmip.OPERATION_GET, RequestPayload)
	return resp, err
}

// GetKeyAttributes retrieves the key attributes from keyvault
func (c *client) GetKeyAttributes(id string) (serverResp interface{}, err error) {
	var resp interface{}
	RequestPayload :=
		kmip.GetAttributeListRequest{
			UniqueIdentifier: id,
		}
	resp, err = c.kclient.Send(kmip.OPERATION_GET_ATTRIBUTES, RequestPayload)
	return resp, err
}

// DeleteKey deletes the key from keyvault
func (c *client) DeleteKey(id string) (serverResp interface{}, err error) {
	var resp interface{}
	RequestPayload :=
		kmip.DestroyRequest{
			UniqueIdentifier: id,
		}
	resp, err = c.kclient.Send(kmip.OPERATION_DESTROY, RequestPayload)
	return resp, err
}
