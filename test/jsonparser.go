package main

import (
	"encoding/json"
	"fmt"
)

func UnmarshalShape(data []byte) (Shape, error) {
	var shapeData map[string]json.RawMessage
	err := json.Unmarshal(data, &shapeData)
	if err != nil {
		return nil, err
	}

	shapeTypeData, ok := shapeData["type"]
	if !ok {
		return nil, fmt.Errorf("shape type not found")
	}

	var shapeType string
	err = json.Unmarshal(shapeTypeData, &shapeType)
	if err != nil {
		return nil, err
	}

	switch shapeType {
	case "circle":
		var circle Circle
		err := json.Unmarshal(data, &circle)
		if err != nil {
			return nil, err
		}
		return &circle, nil
	case "triangle":
		var triangle Triangle
		err := json.Unmarshal(data, &triangle)
		if err != nil {
			return nil, err
		}
		return &triangle, nil
	case "square":
		var square Square
		err := json.Unmarshal(data, &square)
		if err != nil {
			return nil, err
		}
		return &square, nil
	case "polygon":
		var polygon Polygon
		err := json.Unmarshal(data, &polygon)
		if err != nil {
			return nil, err
		}
		return &polygon, nil
	case "line":
		var line Line
		err := json.Unmarshal(data, &line)
		if err != nil {
			return nil, err
		}
		return &line, nil
	case "ellipse":
		var ellipse Ellipse
		err := json.Unmarshal(data, &ellipse)
		if err != nil {
			return nil, err
		}
		return &ellipse, nil
	default:
		return nil, fmt.Errorf("unsupported shape type: %s", shapeType)
	}
}
