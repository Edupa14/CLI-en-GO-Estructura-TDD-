package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Estructura JSON

type Atributos struct {
	Nombre         string `json:"nombre"`
	Tipo           string `json:"tipo"`
	CantCaracteres int    `json:"cant_caracteres"`
}

type Json struct {
	NomMantenedor string      `json:"nom_mantenedor"`
	Atributos     []Atributos `json:"atributos"`
	Tipo          string      `json:"tipo"`
	Servicio      string      `json:"servicio"`
}

//Conversion de String a objeto JSON
func GetJson(dir string) Json {
	raw, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var data Json
	err = json.Unmarshal(raw, &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	/*for _, v := range data.Atributos {
		fmt.Println(strconv.Quote(v.CantCaracteres))
		fmt.Printf("Nombre-> %s | Tipo-> %s | Caracteres-> %d \n", v.Nombre, v.Tipo, v.CantCaracteres)
	}*/
	/*fmt.Printf("%+v", resultados.Atributos[1].Nombre)*/
	return data

	/*buf, err := json.Marshal(resultados)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)*/
	/*json.Unmarshal(raw, &tentacles)
	return tentacles*/
}

/*func (t Json) ToString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}*/

/*func main() {
	s :=
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

}*/
