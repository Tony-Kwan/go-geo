package geojson

import (
	"encoding/json"
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"math/rand"
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

func (w GeojsonWriter) EncodePolygons(ps []geo.Polygon) string {
	geojson := map[string]interface{}{
		"type": "FeatureCollection",
	}
	features := make([]map[string]interface{}, 0)
	for _, p := range ps {
		coord := make([][]float64, 0)
		for _, point := range p.GetShell() {
			coord = append(coord, []float64{point.X(), point.Y()})
		}
		feature := map[string]interface{}{
			"type": "Feature",
			"properties": map[string]interface{}{
				"stroke":         "#555555",
				"stroke-width":   1,
				"stroke-opacity": 1,
				"fill":           randomColor(),
				"fill-opacity":   0.5,
			},
			"geometry": map[string]interface{}{
				"type":        "Polygon",
				"coordinates": [][][]float64{coord},
			},
		}
		features = append(features, feature)
	}
	geojson["features"] = features
	buf, _ := json.MarshalIndent(&geojson, "", "\t")
	return string(buf)
}

func randomColor() string {
	return fmt.Sprintf("#%x%x%x", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}
