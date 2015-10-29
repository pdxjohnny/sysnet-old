package commands

var ConfigOptions = map[string]interface{}{
	"discovery": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": 54634,
			"help":  "Port to bind to",
		},
		"cert": map[string]interface{}{
			"value": "keys/discovery/cert.pem",
			"help":  "Certificate for https server",
		},
		"key": map[string]interface{}{
			"value": "keys/discovery/key.pem",
			"help":  "Key for https server",
		},
		"password": map[string]interface{}{
			"value": "password",
			"help":  "Password to become a node",
		},
	},
}
