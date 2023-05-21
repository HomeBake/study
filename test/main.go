package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fogleman/gg"
)

type Color struct {
	R, G, B float64
}

type Shape interface {
	Draw(dc *gg.Context)
}

func main() {
	// Чтение файла с данными фигур
	filePath := "shapes.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// Декодирование JSON данных в структуры фигур
	var shapes []map[string]json.RawMessage
	err = json.Unmarshal(data, &shapes)
	if err != nil {
		log.Fatal(err)
	}

	// Создание контекста для рисования
	const width = 800
	const height = 600
	dc := gg.NewContext(width, height)

	// Отображение фигур на холсте
	for _, shapeData := range shapes {
		data, err := json.Marshal(shapeData)
		if err != nil {
			log.Fatal(err)
		}
		shape, err := UnmarshalShape(data)
		if err != nil {
			log.Fatal(err)
		}
		shape.Draw(dc)
	}

	// Сохранение изображения
	err = dc.SavePNG("output.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Изображение сохранено в output.png")
}
