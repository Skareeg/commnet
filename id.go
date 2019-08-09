package comms

import (
	"fmt"
	"github.com/mediocregopher/radix"
)

func GetID(redis radix.Pool) int {
	var newID int
	getIDErr := redis.Do(radix.Cmd(&newID, "INCR", "gg.gid"))
	if(getIDErr != nil) {
		fmt.Printf("Creating new global ID: %s\n", getIDErr.Error())
		newID = 0
		setIDErr := redis.Do(radix.Cmd(nil, "SET", "gg.gid", "0"))
		if(setIDErr != nil) {
			fmt.Printf("CANNOT SET ID!\n")
			panic(setIDErr)
		}
		fmt.Printf("Generated.\n")
	}
	return newID
}