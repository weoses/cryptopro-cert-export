package registry2

import (
	"io"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

type regKeyImpl struct {
	key     registry.Key
	keyPath string
	access  uint32
}

type RegKey interface {
	ResolveAccess(access uint32, path ...string) (RegKey, error)
	Resolve(path ...string) (RegKey, error)
	SubKeysNames() ([]string, error)
	SubKeys(
		access uint32,
		nameFilterCallback func(string) bool,
		errCallback func(string, error)) []RegKey

	ValueNames() ([]string, error)

	GetStringValue(name string) (string, error)
	GetBinaryValue(name string) ([]byte, error)

	GetKeyFullPath() string
	GetKeyName() string
	Close()
}

func (it *regKeyImpl) Close() {
	it.key.Close()
}

func (it *regKeyImpl) Resolve(path ...string) (RegKey, error) {
	return it.ResolveAccess(it.access, path...)
}

func (it *regKeyImpl) ResolveAccess(access uint32, path ...string) (RegKey, error) {
	childPath := filepath.Join(path...)
	key, err := registry.OpenKey(it.key, childPath, access)
	if err != nil {
		return nil, err
	}
	return &regKeyImpl{
		key,
		filepath.Join(it.keyPath, childPath),
		access,
	}, nil
}

func (it *regKeyImpl) SubKeysNames() ([]string, error) {
	subkeys := make([]string, 0)
	var errEnd error
	it.SubKeys(
		0,
		func(name string) bool {
			subkeys = append(subkeys, name)
			return false
		},
		func(s string, err error) {
			if s == "." {
				errEnd = err
			}
		},
	)

	return subkeys, errEnd
}

func (it *regKeyImpl) SubKeys(
	access uint32,
	nameFilterCallback func(string) bool,
	errCallback func(string, error)) []RegKey {
	retValue := make([]RegKey, 0)

	subs, err := it.key.ReadSubKeyNames(500)
	if err != nil && err != io.EOF {
		errCallback(".", err)
		return make([]RegKey, 0)
	}

	for _, subkey := range subs {
		if !nameFilterCallback(subkey) {
			continue
		}

		newKey, err := it.ResolveAccess(access, subkey)

		if err != nil {
			errCallback(subkey, err)
			continue
		}
		retValue = append(retValue, newKey)
	}

	return retValue
}

func (it *regKeyImpl) GetStringValue(name string) (string, error) {
	val, _, err := it.key.GetStringValue(name)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (it *regKeyImpl) GetBinaryValue(name string) ([]byte, error) {
	val, _, err := it.key.GetBinaryValue(name)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (it *regKeyImpl) ValueNames() ([]string, error) {
	valueNames, err := it.key.ReadValueNames(500)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return valueNames, nil
}

func (it *regKeyImpl) GetKeyFullPath() string {
	return it.keyPath
}

func (it *regKeyImpl) GetKeyName() string {
	return filepath.Base(it.keyPath)
}

func HKLM(access uint32) (RegKey, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "\\", access)
	if err != nil {
		return nil, err
	}

	return &regKeyImpl{
		key,
		"HKLM",
		access,
	}, nil
}

func HKLMPath(access uint32, path string) (RegKey, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, access)
	if err != nil {
		return nil, err
	}

	return &regKeyImpl{
		key,
		filepath.Join("HKLM", path),
		access,
	}, nil
}
