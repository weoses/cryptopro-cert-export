package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/weoses/cryptopro-cert-export/registry2"

	"golang.org/x/sys/windows/registry"
	"golang.org/x/text/encoding/charmap"
)

// HKEY_LOCAL_MACHINE\S-1-5-21-2408434269-2496968396-1739759940-6064\Keys\
func main() {

	container := ""
	userSid := ""
	outputFolder := "."

	flag.StringVar(&container, "container", "", "Container name to export")
	flag.StringVar(&userSid, "sid", "", "User sid")
	flag.StringVar(&outputFolder, "out", ".", "Output folder")

	flag.Parse()

	log.Printf("Arg: container     = '%s'\n", container)
	log.Printf("Arg: userSid       = '%s'\n", userSid)
	log.Printf("Arg: outputFolder  = '%s'\n", outputFolder)

	log.Println("Opening cryptopro root (WOW64)...")
	cproot, err := registry2.HKLMPath(
		registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS,
		`SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users`)
	if err != nil {
		log.Println("Opening cryptopro root (x64)...")
		cproot, err = registry2.HKLMPath(
			registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS,
			`SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users`)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer cproot.Close()

	log.Println("Reading users nodes...")
	userRoots := cproot.SubKeys(
		registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS,
		func(s string) bool {
			log.Printf("Found user %s ...\n", s)
			return userSid == "" || userSid == s
		},
		func(s string, err error) {
			log.Printf("Error open subkey for user %s  - %s...\n", s, err.Error())
		},
	)

	log.Printf("User nodes count - %d \n", len(userRoots))

	defer func() {
		for _, rk := range userRoots {
			rk.Close()
		}
	}()

	for _, userRoot := range userRoots {
		log.Println("Reading containers...")
		keyRoot, err := userRoot.Resolve("Keys")
		if err != nil {
			log.Printf("Cant read user Keys dir - %s / Keys", userRoot.GetKeyFullPath())
			continue
		}
		defer keyRoot.Close()

		containerRoots := keyRoot.SubKeys(
			registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS,
			func(s string) bool {
				log.Printf("Found Container %s ...\n", s)
				return container == "" || container == s
			},
			func(s string, err error) {
				log.Printf("Error open subkey for Container %s  - %s...\n", s, err.Error())
			},
		)
		defer func() {
			for _, rk := range containerRoots {
				rk.Close()
			}
		}()

		prefixes := map[string]int{}

		for _, containerRoot := range containerRoots {
			log.Printf("Full container name - %s", containerRoot.GetKeyFullPath())

			keyName := containerRoot.GetKeyName()
			keys, err := containerRoot.ValueNames()

			if err != nil {
				log.Printf("No values in %s - %s\n", keyName, err.Error())
				continue
			}

			folder := ""

			for _, char := range []rune(keyName)[0:8] {
				valByte, _ := charmap.Windows1251.EncodeRune(char)

				// its a magic....
				if valByte > 127 {
					valByte = valByte - 20
					valByte = valByte % 26
					valByte = valByte + 97
				}

				if valByte < 35 {
					valByte = valByte % 26
					valByte = valByte + 97
				}

				valConverted := charmap.Windows1251.DecodeByte(valByte)
				folder += string(valConverted)
			}

			log.Printf("Encoded Container name - %s \n", folder)

			v := prefixes[folder]
			folder = fmt.Sprintf("%s.%03d", folder, v)
			v++
			prefixes[folder] = v
			folder = filepath.Join(outputFolder, folder)

			os.MkdirAll(folder, 0755)

			for _, name := range keys {
				data, err := containerRoot.GetBinaryValue(name)
				if err != nil {
					log.Printf("Cant read %s from reg - %s\n", name, err.Error())
					continue
				}
				filename := filepath.Join(folder, name)
				file, err := os.Create(filename)
				if err != nil {
					log.Printf("Cant write %s file - %s\n", name, err.Error())
					continue
				}
				defer file.Close()

				log.Printf("Saving %s file\n", filename)

				io.Copy(file, bytes.NewBuffer(data))
			}
		}
	}
}
