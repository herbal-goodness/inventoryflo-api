package shopify

import (
	"encoding/json"
	"fmt"
	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"
	"net/http"
)

func GetResources(resource string) (map[string]interface{}, error) {
	var url string

	switch resource {
	case "products":
		url = getBaseUrl() + "products.json"
	case "orders":
		url = getBaseUrl() + "orders.json"
	default:
		return nil, fmt.Errorf("%s is not a supported resource", resource)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resources map[string]interface{}
	switch resource {
	case "products":
		var products model.ShopifyProducts
		err = json.NewDecoder(resp.Body).Decode(&products)
		if err != nil {
			return nil, err
		}
		resources = map[string]interface{}{
			"products": products.Products,
		}
	case "orders":
		var orders model.ShopifyOrders
		err = json.NewDecoder(resp.Body).Decode(&orders)
		if err != nil {
			return nil, err
		}
		resources = map[string]interface{}{
			"orders": orders.Orders,
		}
	}
	return resources, nil
}

func getBaseUrl() string {
	apiKey := config.Get("shopifyKey")
	password := config.Get("shopifyPass")
	baseUrl := config.Get("shopifyUrl")
	return fmt.Sprintf("https://%s:%s@%s", apiKey, password, baseUrl)
}
