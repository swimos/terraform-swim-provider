package swim

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	// Time allowed to write a message.
	writeWait = 5 * time.Second
	// Time allowed to read a message.
	readWait = 5 * time.Second
	// Maximum message size.
	maxMessageSize = 512
)

type SwimClient struct {
	url string
}

//Todo refactor duplications

// ------------------- Value Downlink Operations -------------------
func (client SwimClient) GetValueDownlink(node string, lane string) (string, diag.Diagnostics) {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(readWait))
	conn.SetWriteDeadline(time.Now().Add(writeWait))

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

		if err != nil {
			return value, diag.FromErr(err)
		} else {
			response := string(resp)
			if strings.HasPrefix(response, "@event") {
				re := regexp.MustCompile(`^@event\(.*\)(.*?)$`)
				match := re.FindStringSubmatch(response)
				value = string(match[1])
			} else if strings.HasPrefix(response, "@synced") {
				return value, diag.FromErr(err)
			} else if strings.HasSuffix(response, "@laneNotFound") {
				return value, diag.Errorf("Lane %q on node %q not found", lane, node)
			} else if strings.HasSuffix(response, "@nodeNotFound") {
				return value, diag.Errorf("Node %q not found", node)
			}
		}
	}
}

func (client SwimClient) SetValueDownlink(node string, lane string, value string) diag.Diagnostics {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(readWait))
	conn.SetWriteDeadline(time.Now().Add(writeWait))
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
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(readWait))
	conn.SetWriteDeadline(time.Now().Add(writeWait))
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
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(readWait))
	conn.SetWriteDeadline(time.Now().Add(writeWait))
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

		if err != nil {
			return items, diag.FromErr(err)
		} else {
			response := string(resp)
			if strings.HasPrefix(response, "@event") {
				re := regexp.MustCompile(`^@event\(.*\)@update\(key:(.*?)\)(.*?)$`)
				match := re.FindStringSubmatch(response)
				items[match[1]] = match[2]
			} else if strings.HasPrefix(response, "@synced") {
				return items, nil
			} else if strings.HasSuffix(response, "@laneNotFound") {
				return items, diag.Errorf("Lane %q on node %q not found", lane, node)
			} else if strings.HasSuffix(response, "@nodeNotFound") {
				return items, diag.Errorf("Node %q not found", node)
			}
		}
	}
}

func (client SwimClient) SetMapDownlink(node string, lane string, items map[string]interface{}) diag.Diagnostics {
	conn, _, err := websocket.DefaultDialer.Dial(client.url, nil)
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(readWait))
	conn.SetWriteDeadline(time.Now().Add(writeWait))
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
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(readWait))
	conn.SetWriteDeadline(time.Now().Add(writeWait))
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
