package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// config help
func contains(s [10]string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func fix_config() {
	fmt.Println("[*] restoring file...")

	// basic config, later will use the github page of the project to download it
	config := `# these comments will be ignored by the program, WARNING: DO NOT EDIT THE WAY THIS FILE IS SET OUT!

# port the server will be served on, (default: 8080) non-root ports arre recomended :^)
port=8080


# diff templates folder, (default: templates/) can be: /home/user/folders_maby/templates/
folder=templates/

# diff gws script folder,  (default: main.gws) can be: /home/user/folders_maby/main.gws
gws=main.gws
	`

	// delete config
	fmt.Print("[*] removing config.conf... ")
	e := os.Remove("config.conf")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("[OK]")

	// create a new config.conf file
	fmt.Print("[*] creating new config.conf... ")
	_, err := os.Create("config.conf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[OK]")

	// write the config to the file
	fmt.Print("[*] writing data to config.conf... ")
	f, err := os.Create("config.conf")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(config)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("[OK]")
}

func Config_reader() [10]string {
	// data_list will always go up in 2, due to the ways it's set out
	// data_list is set out like this: ["name", "value"] -> kinda like json data but easier to implement
	var data_list [10]string
	num := 0

	for i := 0; i != 1; i -= 1 {
		fmt.Println("[*] reading config file...")

		// check if config.conf file is there
		file, err := os.Open("config.conf")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// oope the file
		scanner := bufio.NewScanner(file)

		// loop through the config file
		for scanner.Scan() {
			// arg found! add the arg to it's slot in the data_list
			// the reason it has it's name as well is so it can be used later on (more dynamic)
			if strings.Contains(scanner.Text(), "=") {
				data_list[num] = strings.Split(scanner.Text(), "=")[0]
				data_list[num+1] = strings.Split(scanner.Text(), "=")[1]
				num += 2
			}
		}

		// final config checks
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		} else if contains(data_list, "") {
			fmt.Println("[!!] config broken! restoring to default...")
			fix_config()
		} else {
			return data_list
		}
	}

	// compile error, need this here
	return data_list

}
