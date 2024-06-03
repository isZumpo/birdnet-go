//go:build docker

package birdnet

import (
	"fmt"
	"os"
)

// TensorFlow Lite model data.
var modelData []byte

// TensorFlow Lite meta model data.
var metaModelData []byte

// Labels in zip format.
var labelsZip []byte

func init() {
	modelDataValue, err := os.ReadFile("/birdnet_models/BirdNET_GLOBAL_6K_V2.4_Model_FP32.tflite")
	if err != nil {
		fmt.Println(err)
		return
	}

	modelData = modelDataValue

	metaModelDataValue, err := os.ReadFile("/birdnet_models/BirdNET_GLOBAL_6K_V2.4_MData_Model_V2_FP16.tflite")
	if err != nil {
		fmt.Println(err)
		return
	}

	metaModelData = metaModelDataValue

	labelsZipValue, err := os.ReadFile("/birdnet_models/labels.zip")
	if err != nil {
		fmt.Println(err)
		return
	}

	labelsZip = labelsZipValue
}
