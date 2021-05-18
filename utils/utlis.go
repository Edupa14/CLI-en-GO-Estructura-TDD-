package utils

import (
	"github.com/example/nisira/services/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
valida si un archivo existe
*/
func IsFileExist(dir string) bool {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return true
	}
	return false
}
func CreateFolder(directorio string) {

	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.MkdirAll(directorio, 0777)
		//MkdirAll
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}

//revisar
func CreateDirAll(directorio string) {

	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.MkdirAll(directorio, 0755)

		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}

// Crea un archivo
func CreateFile(rutaDestino string, data string) {
	err := ioutil.WriteFile(rutaDestino, []byte(data), 0755)
	if err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		//	fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}

//	lee un archivo y retorna un string de lo leido
//"path/filepath" templateName, _ := filepath.Abs(templateName)
func ReadFile(templateName string) string {
	data, _ := ioutil.ReadFile(templateName)
	return string(data)
}

//Copia un archivo usando operaciones del sistema
func Copy(srcFileDir string, destFileDir string) {
	srcFile, err := os.Open(srcFileDir)
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create(destFileDir) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	err = destFile.Sync()
	check(err)
}

//lee un archivo y luego lo copia a otro
func ReadAndCopy(srcFileDir string, destFileDir string) {

	b, err := ioutil.ReadFile(srcFileDir)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(destFileDir, b, 0644)
	if err != nil {
		panic(err)
	}

}

// modifica un archivo buscando algo
func ChancarFile(templateName string, MapForReplace map[string]string) {
	data := ReplaceTextInFile(templateName, MapForReplace)
	CreateFile(templateName, data)
}

// Crea un archivo nuevo partiendo de un una plantilla y un arreglo de opciones a remplazar
func NewFileforTemplate(newName string, templateName string, MapForReplace map[string]string) {
	data := ReplaceTextInFile(templateName, MapForReplace)
	CreateFile(newName, data)
}

//remplaza info en un archivo luego lo pasa a una variable
func ReplaceTextInFile(templateName string, MapForReplace map[string]string) string {
	input := ReadFile(templateName)
	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}

//añade al final del archivo un string
func AppEndToFile(destFileDir string, data string) {
	b, err := os.OpenFile(destFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer b.Close()
	if _, err = b.WriteString(data); err != nil {
		panic(err)
	}
}

//falta probar....... agrega muchas lineas nuevas
func AppArrayEndToFile(destFileDir string, datas []string) {
	for _, data := range datas {
		AppEndToFile(destFileDir, data)
	}

}

/*------------------------------------*/
//nombre de mantendor a ruta
func ConvertirNombreARuta(namePath string) string {
	newNamePath := strings.ToLower(strings.Join(strings.Split(namePath, " "), "-"))
	return newNamePath
}
func ConvertNameTitle(attrTitle string) string {
	return strings.Join(strings.Split(strings.Title(attrTitle), " "), "")
}

//
func ConvertTableTitle(json json.Json) string {
	TablaTitulo := ""
	if TrimQuotes(json.Tipo) == "movimientos" {
		TablaTitulo = "TR" + strings.Join(strings.Split(strings.ToUpper(TrimQuotes(json.NomMantenedor)), " "), "_")
	} else if TrimQuotes(json.Tipo) == "mantenedores" {
		TablaTitulo = "TM" + strings.Join(strings.Split(strings.ToUpper(TrimQuotes(json.NomMantenedor)), " "), "_")
	}
	return TablaTitulo
}

//
func ConvertTitleMayus(attrTitle string) string {
	return strings.Join(strings.Split(strings.ToUpper(attrTitle), " "), "")
}

// Convertir los atributos del Json
func ConvertAtributoTitle(attrTitle string) string {
	return strings.Join(ConvertStringToTitle(strings.Split(strings.ToLower(attrTitle), "_")), "")
}

func ConvertNameAttr(attrTitle string) string {
	return strings.Join(strings.Split(strings.ToLower(attrTitle), " "), "_")
}

/*func ListarVarTipos() map[int]string {
	MapVar := make(map[int]string)
	MapVar[0] = "int64"
	MapVar[1] = "float64"
	MapVar[2] = "bool"
	MapVar[3] = "string"
	return MapVar
}*/
func JsonAAtributos(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "fecha_guardado" || TrimQuotes(value.Nombre) == "transaction_uid" {
		} else if TrimQuotes(value.Tipo) == "datetime" {
			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + "\t string \t`json:\"" + TrimQuotes(ConvertNameAttr(value.Nombre)) + "\"`\n\t"
		} else {
			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + "\t" + TrimQuotes(value.Tipo) + "\t`json:\"" + TrimQuotes(ConvertNameAttr(value.Nombre)) + "\"`\n\t"
		}
	}
	return msg
}

// Convertir Atributos a una entidad de una tabla
func JsonAEntidad(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "transaction_uid" {
			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + "\t mssql.UniqueIdentifier \t`db:\"" + TrimQuotes(ConvertNameAttr(value.Nombre)) + "\"`\n\t"
		} else if TrimQuotes(value.Tipo) == "datetime" {
			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + "\t time.Time \t`db:\"" + TrimQuotes(ConvertNameAttr(value.Nombre)) + "\"`\n\t"
		} else {
			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + "\t" + TrimQuotes(value.Tipo) + "\t`db:\"" + TrimQuotes(ConvertNameAttr(value.Nombre)) + "\"`\n\t"
		}

	}
	return msg
}

// Convertir Modelo a DTO
func ModeloAVO(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "transaction_uid" || TrimQuotes(value.Nombre) == "fecha_guardado" {
		} else {

			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + ":               \tModelo." + TrimQuotes(ConvertAtributoTitle(value.Nombre)) + ",\n\t"
		}
	}
	return msg
}

// Convertir Modelo a DTO
func DTOAModelo(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "transaction_uid" || TrimQuotes(value.Nombre) == "fecha_guardado" {
		} else {

			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + ":               \tRequest." + TrimQuotes(ConvertAtributoTitle(value.Nombre)) + ",\n\t"
		}
	}
	return msg
}

// Convertir DTO a Modelo
func ModeloAPaginacion(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "transaction_uid" || TrimQuotes(value.Nombre) == "fecha_guardado" {
		} else {

			msg += TrimQuotes(ConvertAtributoTitle(value.Nombre)) + ":               \tModelo[i]." + TrimQuotes(ConvertAtributoTitle(value.Nombre)) + ",\n\t"
		}
	}
	return msg
}

// Convertir Modelo a DTO
func JsonAColumnas(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		msg += TrimQuotes(ConvertNameAttr(value.Nombre)) + ","
	}
	return msg[:len(msg)-1]
}

// Convertir Json a dbColumnas
func JsonADbColumnas(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "id" || TrimQuotes(value.Nombre) == "fecha_guardado" || TrimQuotes(value.Nombre) == "transaction_uid" {
		} else {
			msg += TrimQuotes(ConvertNameAttr(value.Nombre)) + ","
		}
	}
	return msg[:len(msg)-1]
}

// Convertir Json a dbColumnasSP
func JsonADbColumnasSp(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "id" || TrimQuotes(value.Nombre) == "fecha_guardado" || TrimQuotes(value.Nombre) == "transaction_uid" {
		} else if TrimQuotes(value.Tipo) == "string" {
			msg += TrimQuotes(ConvertNameAttr(value.Nombre)) + " VARCHAR(" + strconv.Itoa(value.CantCaracteres) + ")" + ","
		} else if TrimQuotes(value.Tipo) == "bool" {
			msg += TrimQuotes(ConvertNameAttr(value.Nombre)) + " BIT" + ","
		} else if TrimQuotes(value.Tipo) == "int64" {
			msg += TrimQuotes(ConvertNameAttr(value.Nombre)) + " INT" + ","
		} else if TrimQuotes(value.Tipo) == "float64" {
			msg += TrimQuotes(ConvertNameAttr(value.Nombre)) + " NUMERIC(17,2)" + ","
		}
	}
	return msg[:len(msg)-1]
}

// Convertir Json a dbColumnasTarget
func JsonADbColumnasTarget(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "id" || TrimQuotes(value.Nombre) == "fecha_guardado" || TrimQuotes(value.Nombre) == "transaction_uid" {
		} else {
			msg += "[target]." + TrimQuotes(ConvertNameAttr(value.Nombre)) + "= [source]." + TrimQuotes(ConvertNameAttr(value.Nombre)) + ","
		}
	}
	return msg[:len(msg)-1]
}
func FiltroAvanzado(Json json.Json) string {
	msg := ""
	for _, value := range Json.Atributos {
		if TrimQuotes(value.Nombre) == "codigo" || TrimQuotes(value.Nombre) == "nombre" {
			msg += TrimQuotes(ConvertNameAttr(value.Nombre)) + ","
		}
	}

	return msg[:len(msg)-1]
}

//Crear un Mapa a partir de un arreglo JSON
/*func ConvertJSONToMAP(Json json.Json, arr string){
	MapVar := make(map[int][]interface{})
	for i, v := range Json.Atributos {
		MapVar[i] =[]interface{}{strconv.Quote(v.Nombre), strconv.Quote(v.Tipo), v.CantCaracteres }
	}
	fmt.Println(MapVar[0][1])
}*/

// Funcion para eliminar comillas a un string
func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// Convertir un arreglo de strings a otro con la primera letra en mayuscula
func ConvertStringToTitle(arr []string) []string {
	for indice, valor := range arr {
		arr[indice] = strings.Title(valor)
	}
	return arr
}
