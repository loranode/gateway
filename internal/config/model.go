package config

// Config holds every runtime setting. Values come from the environment with the
// defaults declared here, so the service runs with zero configuration against a
// node reachable at meshtastic.local:4403.
type Config struct {
	// NodeAddr is the host:port of the Meshtastic node's TCP API. A bare host
	// gets the default :4403 port appended in Load.
	NodeAddr string `env:"NODE_ADDR" envDefault:"meshtastic.local:4403"`
	// HTTPAddr is the listen address of the REST server.
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8080"`
	// DBPath is the SQLite database file path.
	DBPath string `env:"DB_PATH" envDefault:"data.db"`
}
