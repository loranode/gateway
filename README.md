# gateway

REST & webhook gateway for a Meshtastic node.

gateway connects directly to a physical Meshtastic node over its TCP client API,
keeps a local view of the mesh — nodes, messages and channels — in SQLite, and
exposes it over a REST API. It can send text messages into the mesh and delivers
webhook events when nodes appear or change and when new messages arrive.

---

REST и webhook шлюз для ноды Meshtastic.

gateway подключается напрямую к физической ноде Meshtastic по её TCP client API,
хранит локальное состояние сети — ноды, сообщения и каналы — в SQLite и отдаёт
его через REST API. Сервис умеет отправлять текстовые сообщения в mesh и
рассылает webhook-события при появлении и изменении нод и при получении новых
сообщений.
