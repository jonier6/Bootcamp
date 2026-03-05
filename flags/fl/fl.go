package fl

import "os"

type Flag struct {
	cmd        string
	value       *bool  
	description string
}


var Flags []Flag

func Parse() {
	args := os.Args[1:]


	for _, arg := range args {

		for i := range Flags {
			if arg == Flags[i].cmd {

				*Flags[i].value = true
			}
		}
	}
}

func Bool(cmd string, defaultValue bool, description string) *bool {
	
	val := defaultValue
	
	
	newFlag := Flag{
		cmd:        cmd,
		value:       &val, 
		description: description,
	}
	Flags = append(Flags, newFlag)

	
	return &val
}

