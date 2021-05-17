/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/example/nisira/services"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Crear una Api completa",
	Long: `Crea una api desde el manetenedor hasta la implementacion`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		/*MapVar := make(map[string]string)
		MapVar["float64"]	= "habilitado"*/
		if flags, _ := cmd.Flags().GetBool("struct"); flags {
			/*fmt.Println("imprimiendo estructura...")
			services.VarSrv.CreateMantenedor("tipos Movimiento", MapVar)*/
		}else{
			services.VarSrv.CrearMantenedor(args[0])
		/*Main(args)
		services.VarSrv.CreateMantenedor("tipos Movimiento", MapVar)*/
		}
		/*fmt.Println("Mantenedor \"" + strings.Title(args[0]) + "\" creado satisfactoriamente!")*/
	},
}

func init() {
	apiCmd.Flags().BoolP("struct", "s", false, "Crear solo estructura de carpetas y archivos")
	makeCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*func Main(args []string){

	reader := bufio.NewReader(os.Stdin)
	MapAttr := make(map[string]string)
	for {
		fmt.Print("NUEVO ATRIBUTO\n")
		fmt.Print("Nombre de atributo\n")
		fmt.Print("-> ")
		nameVar, _ := reader.ReadString('\n')
		// convert CRLF to LF
		nameVar = strings.Replace(nameVar, "\r\n", "", -1)

		if strings.Compare("end", nameVar) == 0 {
			fmt.Println(MapAttr)
			break
		}

		fmt.Print("Tipo de variable | Ingrese el numero del tipo de variable\n")
		for valor, tipoVar := range utils.ListarVarTipos() {
			fmt.Println(valor ,"->", tipoVar)
		}
		fmt.Print("-> ")
		typeVar, _ := reader.ReadString('\n')
		// convert CRLF to LF
		typeVar = strings.Replace(typeVar, "\r\n", "", -1)
		for key, value := range utils.ListarVarTipos(){
			if textVal,_ := strconv.Atoi(typeVar); key == textVal{
				MapAttr[nameVar] = value
				break
			}
		}
	}
	services.VarSrv.CreateMantenedor(args[0],MapAttr)
}*/