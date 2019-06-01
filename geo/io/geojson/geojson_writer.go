package geojson

import (
	"encoding/json"
	"github.com/Tony-Kwan/go-geo/geo"
)

type GeojsonWriter struct{}

type GeometryInfo struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type Geojson struct {
	Type       string                 `json:"type"`
	Geometry   GeometryInfo           `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

func (w GeojsonWriter) EncodePolygon(polygon geo.Polygon, properties map[string]interface{}) string {
	geojson := Geojson{Type: "Feature", Properties: properties, Geometry: GeometryInfo{Type: "Polygon"}}
	coord := make([][]float64, 0)
	for _, point := range polygon.GetShell() {
		coord = append(coord, []float64{point.X(), point.Y()})
	}
	geojson.Geometry.Coordinates = [][][]float64{coord}
	buf, _ := json.MarshalIndent(&geojson, "", "\t")
	return string(buf)
}
