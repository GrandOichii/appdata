package appdata

import (
	"os"
	"path"
)

// The manager of the application data
type AppDataManager struct {
	ApplicationPath string
}

// Creates an AppDataManager object
func CreateAppDataManager(appPath string) (*AppDataManager, error) {
	result := AppDataManager{ApplicationPath: appPath}

	err := result.createDirectory()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Creates the application data directory
func (m AppDataManager) createDirectory() error {
	path := m.getDirectory()
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Mkdir(path, 0755)
	}
	if err != nil {
		return err
	}
	return nil
}

// Returns the current directory
func (m AppDataManager) getDirectory() string {
	return path.Join(appDataPath, m.ApplicationPath)
}

// Returns the path to the file
func (m AppDataManager) PathTo(file string) string {
	return path.Join(m.getDirectory(), file)
}

// Return the contents of the file
func (m AppDataManager) ReadFile(file string) ([]byte, error) {
	contents, err := os.ReadFile(m.PathTo(file))
	if err != nil {
		return nil, err
	}
	return contents, nil
}

// Writes data to the file
func (m AppDataManager) WriteToFile(file string, data []byte) error {
	return os.WriteFile(m.PathTo(file), data, 0755)
}

// Returns true if the file exists
func (m AppDataManager) FileExists(file string) (bool, error) {
	_, err := os.Stat(m.PathTo(file))
	if os.IsNotExist(err) {
		// file doesn't exist
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// Creates a folder in the application data directory
func (m AppDataManager) CreateFolder(folderPath string) error {
	err := os.Mkdir(m.PathTo(folderPath), 0755)
	if err != nil {
		return err
	}
	return nil
}
