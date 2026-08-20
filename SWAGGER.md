## CallbackService
CallbackService manages webhook subscriptions.

### [DELETE] /api/v1/callback
**DeleteCallback unsubscribes a webhook URL.**

#### Request Body

DeleteCallbackRequest unregisters a webhook URL.

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [DeleteCallbackRequest](#deletecallbackrequest-schema)<br> |

#### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |
| default |  |

### [GET] /api/v1/callback
**ListCallbacks returns every registered webhook subscriptions.**

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | ListCallbacksResponse carries the registered callbacks. | **application/json**: [ListCallbacksResponse](#listcallbacksresponse-schema)<br> |
| default |  |  |

### [POST] /api/v1/callback
**RegisterCallback subscribes a webhook URL to mesh events.**

#### Request Body

RegisterCallbackRequest registers a webhook URL.

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [RegisterCallbackRequest](#registercallbackrequest-schema)<br> |

#### Responses

| Code | Description |
| ---- | ----------- |
| 201 |  |
| default |  |

---
## MeshService
MeshService exposes the connected Meshtastic node.

### [GET] /api/v1/channels
**ListChannels returns every configured channel on the radio.**

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Channel is one configured channel on the connected radio. | **application/json**: [ListChannelsResponse.Channels](#listchannelsresponsechannels-schema)<br> |
| default |  |  |

### [GET] /api/v1/channels/{index}/messages
**ListChannelMessages returns messages received on one channel.**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| index | path | channel index (path) | Yes | integer |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | ListMessagesResponse carries the messages for a node. | **application/json**: [ListMessagesResponse](#listmessagesresponse-schema)<br> |
| default |  |  |

### [GET] /api/v1/nodes
**ListNodes returns a compact summary of every node the radio can see.**

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | NodeSummary is the compact node entry returned by the list endpoint. | **application/json**: [ListNodesResponse.Nodes](#listnodesresponsenodes-schema)<br> |
| default |  |  |

### [GET] /api/v1/nodes/{num}
**GetNode returns all stored information for one node by its number.**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| num | path | node number (path) | Yes | integer |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Node is one full entry in the mesh node database. | **application/json**: [Node](#node-schema)<br> |
| default |  |  |

### [GET] /api/v1/nodes/{num}/messages
**ListNodeMessages returns direct messages received from one node.**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| num | path | sender node number (path) | Yes | integer |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Message is one received text message with all stored metadata. | **application/json**: [ListMessagesResponse.Messages](#listmessagesresponsemessages-schema)<br> |
| default |  |  |

### [POST] /api/v1/nodes/{num}/messages
**SendNodeMessage transmits a text message to one node.**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| num | path | recipient node number (path) | Yes | integer |

#### Request Body

SendMessageRequest is a text message to transmit to one node.

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [SendMessageRequest](#sendmessagerequest-schema)<br> |

#### Responses

| Code | Description |
| ---- | ----------- |
| 202 |  |
| default |  |

---
### Schemas

#### DeleteCallbackRequest Schema

DeleteCallbackRequest unregisters a webhook URL.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| url | string | webhook URL to remove | Yes |

#### ListCallbacksResponse Schema

ListCallbacksResponse carries the registered callbacks.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| urls | [ string ] |  | Yes |

#### ListChannelsResponse.Channels Schema

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ListChannelsResponse.Channels | [ { **"index"**: integer, **"name"**: string, **"role"**: string } ] |  |  |

#### ListMessagesResponse Schema

ListMessagesResponse carries the messages for a node.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ { **"channel"**: integer, **"from"**: integer, **"hopsAway"**: integer, **"id"**: integer, **"rssi"**: integer, **"rxTime"**: string (int64), **"snr"**: number, **"text"**: string, **"to"**: integer } ] |  | Yes |

#### ListMessagesResponse.Messages Schema

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ListMessagesResponse.Messages | [ { **"channel"**: integer, **"from"**: integer, **"hopsAway"**: integer, **"id"**: integer, **"rssi"**: integer, **"rxTime"**: string (int64), **"snr"**: number, **"text"**: string, **"to"**: integer } ] |  |  |

#### ListNodesResponse.Nodes Schema

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ListNodesResponse.Nodes | [ { **"longName"**: string, **"num"**: integer, **"shortName"**: string } ] |  |  |

#### Node Schema

Node is one full entry in the mesh node database.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| altitude | integer | altitude, metres; absent if unknown | No |
| battery | integer | battery level, percent; absent if unknown | No |
| hopsAway | integer | hops away from us; absent if unknown | No |
| hwModel | string | hardware model | Yes |
| isFavorite | boolean | true if marked favorite on the radio | Yes |
| lastHeard | string (int64) | last time a packet was heard, Unix seconds | Yes |
| latitude | number | latitude, degrees; absent if unknown | No |
| longName | string | human-readable long name | Yes |
| longitude | number | longitude, degrees; absent if unknown | No |
| num | integer | node number (primary key) | Yes |
| role | string | device role | Yes |
| rssi | integer | received signal strength of the last packet, dBm | Yes |
| shortName | string | short name / call sign | Yes |
| snr | number | signal-to-noise ratio of the last packet, dB | Yes |
| updatedAt | string (int64) | last time this record changed, Unix seconds | Yes |
| viaMqtt | boolean | true if last seen over MQTT rather than LoRa | Yes |
| voltage | number | battery voltage, volts; absent if unknown | No |

#### RegisterCallbackRequest Schema

RegisterCallbackRequest registers a webhook URL.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| url | string | webhook URL to register | Yes |

#### SendMessageRequest Schema

SendMessageRequest is a text message to transmit to one node.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| channel | integer | channel index (0 by default) | No |
| text | string | message body text | Yes |
