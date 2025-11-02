package config

var (
	BROADCAST_PORT                = 8787
	CHAT_PORT                     = 8989
	BROADCAST_INTERVAL            = 5  // in seconds
	USER_TIMEOUT                  = 15 // in seconds
	DISCOVERY_BUFFER_SIZE         = 1024
	ONLINE_USERS_REFRESH_INTERVAL = 3 // in seconds
	PRESENCE_MESSAGE_PREFIX       = "GOCH_PRESENCE:"
	MAX_USERNAME_LENGTH           = 20
	MIN_USERNAME_LENGTH           = 3
	DISCOVERY_CHANNEL             = "goch_discovery_channel"
	GET_LOCAL_IP_RETRY_INTERVAL   = 2 // in seconds
	GET_LOCAL_IP_MAX_RETRIES      = 5
	TUI_REFRESH_INTERVAL          = 1 // in seconds
	DEFAULT_USERNAME              = "Anonymous"
)
