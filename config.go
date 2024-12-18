package main

type AppConfig struct {
	Common   CommonConfig   `toml:"common"`
	Yescrypt YescryptConfig `toml:"yescrypt"`
}

type CommonConfig struct {
	Port       int `toml:"port"`
	SaltLength int `toml:"salt_length"`
}

type YescryptConfig struct {
	YescryptNLog2 int `toml:"yescrypt_nlog2"`
	YescryptR     int `toml:"yescrypt_r"`
}
