package fl

import "os"

type Flag struct {
	Cmd        string
	Value       *bool  
	Description string
}


var Flags []Flag

func Parse() {
	args := os.Args[1:]


	for _, arg := range args {

		for i := range Flags {
			if arg == Flags[i].Cmd {

				*Flags[i].Value = true
			}
		}
	}
}

func Bool(cmd string, defaultValue bool, description string) *bool {
	
	val := defaultValue
	
	
	newFlag := Flag{
		Cmd:        cmd,
		Value:       &val, 
		Description: description,
	}
	Flags = append(Flags, newFlag)

	
	return &val
}

