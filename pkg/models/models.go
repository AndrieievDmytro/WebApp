package models

//Hold data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string  `json:"stringmap"`
	IntMap    map[string]int     `json:"intmap"`
	FloatMap  map[string]float32 `json:"floatmap"`
	Data      map[string]interface{}
	Flash     string
	Warning   string
	Error     string
}
