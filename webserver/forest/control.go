// Warper over Gorilla Mux package
package forest

import (
	"net/http"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/gorilla/mux"
	"encoding/json"
	"encoding/xml"
	_ "github.com/flosch/pongo2-addons"
	"github.com/flosch/pongo2"
	"github.com/DeepForestTeam/mobiussign/components/config"
)

var default_tpl_path string

func init() {
	log.Info("* Init Forest controller")
	var err error
	default_tpl_path, err = config.GetString("TPL_PATH")
	if err != nil {
		default_tpl_path = "view/"
	}
}

type Controller interface {
	Process(http.ResponseWriter, *http.Request)
	PreRoute() (bool)
	Get()
	Post()
	Put()
	Delete()
	Options()
	RenderTemplate()
}

type Control struct {
	ThisName      string
	Input         *http.Request
	Output        http.ResponseWriter
	Context       Context
	//Custom handlers
	PostHandler   func(*Control)
	GetHandler    func(*Control)
	PutHandler    func(*Control)
	DeleteHandler func(*Control)
	OptionHandler func(*Control)
	//Output settings
	AutoRender    bool
	TemplateName  string
	Template      *pongo2.Template
	Data          interface{}
	HasErrors     bool
}
type Context struct {
	Url      string
	Domain   string
	Protocol string
	Method   string
	UrlVars  map[string]string
	GetVars  map[string]string
	PostVars map[string]string
}

func (this *Control)Process(w http.ResponseWriter, r *http.Request) {
	this.HasErrors = false
	this.Output = w
	this.Input = r
	this.Context.UrlVars = mux.Vars(r)
	return
}
func (this *Control)PreRoute() (bool) {
	return false
}
func (this *Control)Get() {
	if this.GetHandler != nil {
		this.GetHandler(this)
	} else {
		log.Warning("Call default Get method", this.ThisName)
		this.Output.WriteHeader(http.StatusMethodNotAllowed)
		this.Output.Write([]byte("405 Method Not Allowed"))
	}
}
func (this *Control)Post() {
	if this.PostHandler != nil {
		this.PostHandler(this)
	} else {
		log.Warning("Call default Post method", this.ThisName)
		this.Output.WriteHeader(http.StatusMethodNotAllowed)
		this.Output.Write([]byte("405 Method Not Allowed"))
	}
}
func (this *Control)Put() {
	this.Output.WriteHeader(http.StatusMethodNotAllowed)
	this.Output.Write([]byte("405 Method Not Allowed"))
}
func (this *Control)Delete() {
	this.Output.WriteHeader(http.StatusMethodNotAllowed)
	this.Output.Write([]byte("405 Method Not Allowed"))
}
func (this *Control)Options() {
	this.Output.Header().Set("Access-Control-Allow-Origin", "*")
	this.Output.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	this.Output.Header().Set("Access-Control-Allow-Headers", "DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type")
	this.Output.WriteHeader(http.StatusOK)
}

func (this *Control)Error(http_error_code string) {
	//@todo show error page
}

func (this *Control)parseRequestData() (err error) {
	err = nil
	return
}

//Support mehtods
func (this *Control)ServeJSON() {
	this.AutoRender = false
	json_string, err := json.MarshalIndent(this.Data, "", "  ")
	if err != nil {
		log.Critical("Can not marshal JSON:", err)
		this.Output.WriteHeader(http.StatusInternalServerError)
		return
	}
	this.Output.Header().Set("Access-Control-Allow-Origin", "*")
	this.Output.Header().Set("Content-Type", "application/json; charset=utf-8")
	this.Output.Write(json_string)
	return
}
func (this *Control)ServeXML() {
	this.AutoRender = false
	xml_string, err := xml.MarshalIndent(this.Data, "", "  ")
	if err != nil {
		log.Critical("Can not marshal JSON:", err)
		this.Output.WriteHeader(http.StatusInternalServerError)
		return
	}
	this.Output.Header().Set("Access-Control-Allow-Origin", "*")
	this.Output.Header().Set("Content-Type", "application/xml; charset=utf-8")
	this.Output.Write(xml_string)
	return
}
func (this *Control)RenderTemplate() {
	if this.AutoRender {
		if this.Template != nil {

		} else {

		}
	}
	return
}