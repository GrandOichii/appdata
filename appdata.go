package appdata

import (
	"os"
	"path"
)

type AppDataManager struct {
	ApplicationPath string
}

func CreateAppDataManager(appPath string) (AppDataManager, error) {
	result := AppDataManager{ApplicationPath: appPath}

	err := result.createDirectory()
	if err != nil {
		return AppDataManager{}, err
	}
	return result, nil
}

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

func (m AppDataManager) GetFiles() []string {
	return []string{}
}

func (m AppDataManager) getDirectory() string {
	return path.Join(appDataPath, m.ApplicationPath)
}

func (m AppDataManager) ConcatPath(file string) string {
	return path.Join(m.getDirectory(), file)
}

func (m AppDataManager) ReadFile(file string) ([]byte, error) {
	contents, err := os.ReadFile(m.ConcatPath(file))
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (m AppDataManager) WriteToFile(file string, data []byte) error {
	return os.WriteFile(m.ConcatPath(file), data, 0755)
}

func (m AppDataManager) FileExists(file string) (bool, error) {
	_, err := os.Stat(m.ConcatPath(file))
	if os.IsNotExist(err) {
		// file doesn't exist
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m AppDataManager) CreateFolder(folderPath string) error {
	err := os.Mkdir(m.ConcatPath(folderPath), 0755)
	if err != nil {
		return err
	}
	return nil
}
