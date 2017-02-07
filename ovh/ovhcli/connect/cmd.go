package connect

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userConfigPath  = "/.ovh.conf" // prefixed with homeDir
	localConfigPath = "./ovh.conf"
)

// get the configuration file path or through an error
func getConfigPath() (path string, err error) {
	_, err = os.Stat(localConfigPath)
	if err == nil {
		path = localConfigPath
		return
	}

	usr, err := user.Current()
	if err == nil {
		currentUserConfigPath := filepath.Join(usr.HomeDir, userConfigPath)
		_, err = os.Stat(currentUserConfigPath)
		if err == nil {
			path = currentUserConfigPath
			return
		}
		return
	}
	return
}

func init() {
	flags := Cmd.Flags()

	flags.String("save-profile", "", "store consumer key in a profile")
	viper.BindPFlag("save_profile", flags.Lookup("save-profile"))

	flags.String("use-profile", "", "use a store profile")
	viper.BindPFlag("use_profile", flags.Lookup("use-profile"))

	flags.String("remove-profile", "", "delete a store profile")
	viper.BindPFlag("remove_profile", flags.Lookup("remove-profile"))

	flags.Bool("list-profiles", false, "list all profiles")
	viper.BindPFlag("list_profiles", flags.Lookup("list-profiles"))
}

func helper(consumerKey string) error {
	return fmt.Errorf("Set environment variable OVH_CONSUMER_KEY=%s\n", consumerKey)
}

func writeConsumerKeyFile(filename string, consumerKey string) (err error) {
	var cfg *ini.File
	var section *ini.Section
	var endpoint string

	if cfg, err = ini.Load(filename); err != nil {
		return errors.New("Cannot load file " + filename)
	}

	if defaultSection, errSection := cfg.GetSection("default"); errSection == nil {
		endpointKey, errKey := defaultSection.GetKey("endpoint")
		if errKey != nil {
			return errors.New("Cannot read endpoint from configuration")
		}
		endpoint = endpointKey.String()
	} else {
		return errors.New("Cannot read default section")
	}

	if section, err = cfg.GetSection(endpoint); err != nil {
		if section, err = cfg.NewSection(endpoint); err != nil {
			return errors.New("Cannot create section " + endpoint)
		}
	}

	if section.NewKey("consumer_key", consumerKey); err != nil {
		return errors.New("Cannot create key 'consumer_key'")
	}

	if err = cfg.SaveTo(filename); err != nil {
		return errors.New("Cannot save to " + filename)
	}
	return
}

func writeConsumerKey(consumerKey string) (err error) {
	if path, err := getConfigPath(); err == nil {
		if err = writeConsumerKeyFile(path, consumerKey); err != nil {
			return helper(consumerKey)
		}
	}

	return
}

func saveConsumerKey(consumerKey string, profileToSave string) error {
	if iniPath, err := getConfigPath(); err == nil {
		var cfg *ini.File
		var section *ini.Section

		if cfg, err = ini.Load(iniPath); err != nil {
			return errors.New("Cannot load file " + iniPath)
		}

		if section, err = cfg.GetSection("profiles"); err != nil {
			if section, err = cfg.NewSection("profiles"); err != nil {
				return errors.New("Cannot create section [profiles]")
			}
		}

		if section.NewKey(profileToSave, consumerKey); err != nil {
			return errors.New("Cannot create profile")
		}

		if err = cfg.SaveTo(iniPath); err != nil {
			return errors.New("Cannot save to " + iniPath)
		}
	}
	return nil
}

func getProfile(profileToUse string) (consumerKey string, err error) {
	var iniPath string
	var profileSection *ini.Section
	if iniPath, err = getConfigPath(); err == nil {
		var cfg *ini.File

		cfg, err = ini.Load(iniPath)
		if err != nil {
			err = errors.New("Cannot load file " + iniPath)
			return
		}

		profileSection, err = cfg.GetSection("profiles")
		if err == nil {
			consumer, errKey := profileSection.GetKey(profileToUse)
			if errKey != nil {
				err = errors.New("Cannot read endpoint from configuration")
				return
			}
			consumerKey = consumer.String()
		}
	}
	return
}

func delProfile(profileToDel string) (err error) {
	var iniPath string
	var profileSection *ini.Section
	if iniPath, err = getConfigPath(); err == nil {
		var cfg *ini.File

		cfg, err = ini.Load(iniPath)
		if err != nil {
			err = errors.New("Cannot load file " + iniPath)
			return
		}

		profileSection, err = cfg.GetSection("profiles")
		if err == nil {
			profileSection.DeleteKey(profileToDel)
			err = cfg.SaveTo(iniPath)
			if err != nil {
				return errors.New("Cannot save to " + iniPath)
			}
		}
	}
	return
}

func listProfiles() error {
	hasProfiles := false
	if iniPath, err := getConfigPath(); err == nil {
		var cfg *ini.File

		if cfg, err = ini.Load(iniPath); err != nil {
			return errors.New("Cannot load file " + iniPath)
		}

		if profileSection, errSection := cfg.GetSection("profiles"); errSection == nil {
			profileKeys := profileSection.Keys()
			hasProfiles = len(profileSection.Keys()) > 0
			if hasProfiles {
				fmt.Printf("Profiles :\n")
			}
			for _, key := range profileKeys {
				fmt.Printf(" - %s\n", key.Name())
			}

		}

		if !hasProfiles {
			fmt.Printf("No profile\n")
		}

	}

	return nil
}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "connect",
	Short: "Domain commands: ovhcli connect",
	Long:  `Domain commands: ovhcli connect`,
	Run: func(cmd *cobra.Command, args []string) {

		profileToSave := viper.GetString("save_profile")
		profileToUse := viper.GetString("use_profile")
		profileToDel := viper.GetString("remove_profile")

		if profileToDel != "" {
			if err := delProfile(profileToDel); err != nil {
				fmt.Errorf("Cannot remove %s\n", profileToDel)
			}
			fmt.Printf("Profile %s deleted\n", profileToDel)
			return
		}

		if viper.GetBool("list_profiles") {
			listProfiles()
			return
		}

		if profileToUse == "" {
			c, err := ovh.NewDefaultClient()
			common.Check(err)

			ckReq := c.NewCkRequest()

			// Allow GET method on /me
			ckReq.AddRules(ovh.ReadWrite, "/*")

			response, err := ckReq.Do()
			common.Check(err)

			// Print the validation URL
			fmt.Printf("Please visit %s to complete your login\n", response.ValidationURL)

			// set consumer key
			if err = writeConsumerKey(response.ConsumerKey); err != nil {
				common.Check(err)
			}

			if profileToSave != "" {
				saveConsumerKey(response.ConsumerKey, profileToSave)
			}
		} else {
			if consumerKey, err := getProfile(profileToUse); err == nil {
				writeConsumerKey(consumerKey)
				fmt.Printf("Switched to profile %s\n", profileToUse)
			} else {
				fmt.Errorf("Profile %s not found\n", profileToUse)
			}
		}

	},
}
