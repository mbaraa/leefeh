package config

import (
	"net"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	if len(MachineAddress()) == 0 {
		os.Setenv("MACHINE_IP", getMachineIP())
	}
}

func PortNumber() string     { return os.Getenv("PORT") }
func DBUser() string         { return os.Getenv("DB_USER") }
func DBPassword() string     { return os.Getenv("DB_PASSWORD") }
func DBHost() string         { return os.Getenv("DB_HOST") }
func AllowedClients() string { return os.Getenv("ALLOWED_CLIENTS") }
func MachineAddress() string { return os.Getenv("MACHINE_IP") }
func JWTSecret() []byte      { return []byte(os.Getenv("JWT_SECRET")) }
func Development() bool      { return os.Getenv("DEVELOPMENT") == "true" }

func getMachineIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:"+PortNumber())
	if err != nil {
		panic(err)
	}

	err = conn.Close()
	if err != nil {
		panic(err)
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
