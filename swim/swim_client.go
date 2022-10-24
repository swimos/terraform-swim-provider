package swim

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type SwimClient struct {
	url string
}

//Todo add deadlines https://github.com/gorilla/websocket/blob/af47554f343b4675b30172ac301638d350db34a5/examples/chat/client.go#L91
//Todo add error for lane not found
//Todo refactor duplications

// ------------------- Value Downlink Operations -------------------
func (client SwimClient) GetValueDownlink(node string, lane string) (string, diag.Diagnostics) {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	defer conn.Close()

	value := ""

	if err != nil {
		return value, diag.FromErr(err)
	}

	message := fmt.Sprintf("@sync(node: %q, lane: %q)", node, lane)
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		return value, diag.FromErr(err)
	}

	for {
		_, resp, err := conn.ReadMessage()
		response := string(resp)

		if err != nil {
			return value, diag.FromErr(err)
		} else {
			if strings.HasPrefix(response, "@event") {
				re := regexp.MustCompile(`^@event\(.*\)(.*?)$`)
				match := re.FindStringSubmatch(response)
				value = string(match[1])
			} else if strings.HasPrefix(response, "@synced") {
				return value, nil
			}
		}
	}
}

func (client SwimClient) SetValueDownlink(node string, lane string, value string) diag.Diagnostics {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	defer conn.Close()

	if err != nil {
		return diag.FromErr(err)
	}

	message := fmt.Sprintf("@command(node: %q, lane: %q)%q", node, lane, value)
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func (client SwimClient) ClearValueDownlink(node string, lane string) diag.Diagnostics {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	defer conn.Close()

	if err != nil {
		return diag.FromErr(err)
	}

	message := fmt.Sprintf("@command(node: %q, lane: %q)", node, lane)
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// ------------------- Map Downlink Operations -------------------
func (client SwimClient) GetMapDownlink(node string, lane string) (map[string]string, diag.Diagnostics) {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	defer conn.Close()

	items := make(map[string]string)

	if err != nil {
		return items, diag.FromErr(err)
	}

	message := fmt.Sprintf("@sync(node: %q, lane: %q)", node, lane)
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		return items, diag.FromErr(err)
	}

	for {
		_, resp, err := conn.ReadMessage()
		response := string(resp)

		if err != nil {
			return items, diag.FromErr(err)
		} else {
			if strings.HasPrefix(response, "@event") {
				re := regexp.MustCompile(`^@event\(.*\)@update\(key:(.*?)\)(.*?)$`)
				match := re.FindStringSubmatch(response)
				items[match[1]] = match[2]
			} else if strings.HasPrefix(response, "@synced") {
				return items, nil
			}
		}
	}
}

func (client SwimClient) SetMapDownlink(node string, lane string, items map[string]interface{}) diag.Diagnostics {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	defer conn.Close()

	if err != nil {
		return diag.FromErr(err)
	}

	message := fmt.Sprintf("@command(node: %q, lane: %q)@clear", node, lane)
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		return diag.FromErr(err)
	}

	for key, value := range items {
		message := fmt.Sprintf("@command(node: %q, lane: %q)@update(key: %q)%q", node, lane, key, value)
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))

		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func (client SwimClient) ClearMapDownlink(node string, lane string) diag.Diagnostics {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	defer conn.Close()

	if err != nil {
		return diag.FromErr(err)
	}

	message := fmt.Sprintf("@command(node: %q, lane: %q)@clear", node, lane)
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
