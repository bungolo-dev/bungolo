package bungolow

import (
	"os"
	"path"
)

var environment *EnvironmentConfig

type EnvironmentConfig struct {
	BaseDirectory   string
	ImgDirectory    string
	PluginDirectory string
}

func NewEnvironment(baseDirectory string, imgDirectory string) (Environment, error) {
	env := &EnvironmentConfig{
		BaseDirectory:   "C://Temp/Bungolow",
		ImgDirectory:    "C://Temp/Images",
		PluginDirectory: "C://Temp/Plugins",
	}

	environment = env

	baseErr := os.Mkdir(env.BaseDirectory, os.ModeAppend)
	if baseErr != nil {
		return nil, baseErr
	}

	imgErr := os.Mkdir(env.ImgDirectory, 0755)
	if imgErr != nil {
		return nil, imgErr
	}

	plErr := os.Mkdir(env.PluginDirectory, 0755)
	if plErr != nil {
		return nil, plErr
	}

	return env, nil
}

func GetEnvironment() Environment {
	return environment
}

type Environment interface {
	BaseDir() string
	ImgDir() string

	GetFilePath(filename string) string
	GetImagePath(filename string) string
}

func (env *EnvironmentConfig) GetFilePath(filename string) string {
	return path.Join(env.BaseDirectory, filename)
}

func (env *EnvironmentConfig) GetImagePath(filename string) string {
	return path.Join(env.ImgDirectory, filename)
}

func (env *EnvironmentConfig) BaseDir() string {
	return env.BaseDirectory
}

func (env *EnvironmentConfig) ImgDir() string {
	return env.ImgDirectory
}
