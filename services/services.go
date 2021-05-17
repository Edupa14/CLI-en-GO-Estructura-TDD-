package services

import (
	"github.com/example/nisira/services/json"
	"github.com/example/nisira/utils"
	"strings"
)

var VarSrv *Service

type Service struct {
}

func (s *Service) CreateMapDefault(fileName string, json json.Json) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%NOMBREMANTENEDOR%"] = utils.ConvertNameTitle(fileName)
	MapForReplace["%NOMBREPAQUETEMANTENEDOR%"] = utils.ConvertNameAttr(fileName)
	MapForReplace["%RUTAMANTENEDOR%"] = "ns-" + utils.ConvertirNombreARuta(json.Servicio) + "-service/api/" + utils.ConvertirNombreARuta(json.Tipo) + "/" + utils.ConvertirNombreARuta(fileName)
	MapForReplace["%VARIABLESJSON%"] = utils.JsonAAtributos(json)
	MapForReplace["%VARIABLESENTIDAD%"] = utils.JsonAEntidad(json)
	MapForReplace["%VARIABLESVO%"] = utils.ModeloAVO(json)
	return MapForReplace
}

func (s *Service) ObtenerRutasPlantillas() map[string]string {
	MapPlatilla := make(map[string]string)
	MapPlatilla["service"] = "application/service"
	MapPlatilla["usecase_app"] = "application/usecase"
	MapPlatilla["request"] = "domain/dto"
	MapPlatilla["entity"] = "domain/entity"
	MapPlatilla["mapping"] = "domain/mapping"
	MapPlatilla["model"] = "domain/model"
	MapPlatilla["repository"] = "domain/repository"
	MapPlatilla["usecase_dom"] = "domain/usecase"
	MapPlatilla["response"] = "domain/valueobject"
	MapPlatilla["controller"] = "infrastructure/delivery/http/controller"
	MapPlatilla["routes"] = "infrastructure/delivery/http/routes"
	MapPlatilla["mssql"] = "infrastructure/persistence/mssql_repository"
	MapPlatilla["mnt"] = ""

	return MapPlatilla
}

func (s *Service) CrearMantenedorBack(rutaJson string) {

	rutaBasePlantillas := "C:/Proyectos/Nisira-CLI/templates/backend/"
	objJson := json.GetJson(rutaJson)
	nomMantenedor := objJson.NomMantenedor
	MapForReplace := s.CreateMapDefault(nomMantenedor, objJson)
	nomMantenedor = utils.ConvertirNombreARuta(nomMantenedor)
	routeApi := "api/mantenedores/" + nomMantenedor
	if objJson.Tipo == "movimientos" {
		routeApi = "api/movimientos/" + nomMantenedor
	}
	utils.CreateFolder(routeApi)
	for fileName, routeFile := range s.ObtenerRutasPlantillas() {

		utils.CreateFolder(routeApi + "/" + routeFile)

		newFileName := strings.Split(fileName, "_")[0]

		routeTemplate := ""
		newRouteFile := ""

		if strings.Compare(routeFile, "") == 1 {
			routeTemplate = rutaBasePlantillas + routeFile + "/mantenedor.stub"
			newRouteFile = routeApi + "/" + routeFile + "/" + nomMantenedor + "." + newFileName + ".go"
		} else {
			routeTemplate = rutaBasePlantillas + "mantenedor.stub"
			newRouteFile = routeApi + "/" + nomMantenedor + "." + newFileName + ".go"
		}
		utils.NewFileforTemplate(newRouteFile, routeTemplate, MapForReplace)
	}
}