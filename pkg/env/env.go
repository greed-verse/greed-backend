package env

import (
	"os"
	"strings"
	"sync"
)

type SecretStorer interface {
	ResolveEnv(env string) (string, error)
	DB_URL() string
	APP_ADDRESS() string
	APPLE_TEAM_ID() string
	APPLE_CLIENT_ID() string
	APPLE_KEY_ID() string
	APPLE_PRIVATE_KEY() string
}

type Environment struct {
	database        string
	port            string
	appleTeamId     string
	appleClientId   string
	appleKeyId      string
	applePrivateKey string
}

var env SecretStorer
var once sync.Once

func GetEnv() SecretStorer {
	once.Do(func() {
		environment, err := newEnv()
		if err != nil {
			panic("Couldn't Load environment: " + err.Error())
		}
		env = environment
	})
	return env
}

func newEnv() (SecretStorer, error) {
	var environment *Environment = &Environment{}

	url, err := environment.ResolveEnv("DB_URL")
	if err != nil {
		return nil, err
	}

	port, err := environment.ResolveEnv("APP_ADDRESS")
	if err != nil {
		return nil, err
	}

	// teamId, err := environment.ResolveEnv("APPLE_TEAM_ID")
	// if err != nil {
	// 	return nil, err
	// }
	//
	// clientId, err := environment.ResolveEnv("APPLE_CLIENT_ID")
	// if err != nil {
	// 	return nil, err
	// }
	//
	// keyId, err := environment.ResolveEnv("APPLE_KEY_ID")
	// if err != nil {
	// 	return nil, err
	// }
	//
	// privateKey, err := environment.ResolveEnv("APPLE_PRIVATE_KEY")
	// if err != nil {
	// 	return nil, err
	// }

	environment.database = url
	environment.port = port
	// environment.appleTeamId = teamId
	// environment.appleClientId = clientId
	// environment.appleKeyId = keyId
	// environment.applePrivateKey = privateKey
	return environment, nil
}

func (env *Environment) DB_URL() string {
	return env.database
}

func (env *Environment) APP_ADDRESS() string {
	return env.port
}

func (env *Environment) APPLE_TEAM_ID() string {
	return env.appleTeamId
}

func (env *Environment) APPLE_CLIENT_ID() string {
	return env.appleClientId
}

func (env *Environment) APPLE_KEY_ID() string {
	return env.appleKeyId
}

func (env *Environment) APPLE_PRIVATE_KEY() string {
	return env.applePrivateKey
}

func (e *Environment) ResolveEnv(env string) (string, error) {
	filepath, exists := os.LookupEnv(env + "_FILE")

	if exists {
		content, err := os.ReadFile(filepath)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(content)), nil
	}
	return os.Getenv(env), nil
}
