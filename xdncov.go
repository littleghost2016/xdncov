package main

func main() {

	studentConfigSlice := CollectConfigs("./configs")

	for _, eachConfig := range studentConfigSlice {
		// go signIn(eachConfig)
		SignIn(eachConfig)
	}

}
