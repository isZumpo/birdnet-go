//go:build !docker

package birdnet

import (
	_ "embed" // Embedding data directly into the binary.
)

// Embedded TensorFlow Lite model data.
//
//go:embed BirdNET_GLOBAL_6K_V2.4_Model_FP32.tflite
var modelData []byte

// Embedded TensorFlow Lite meta model data.
//
//go:embed BirdNET_GLOBAL_6K_V2.4_MData_Model_V2_FP16.tflite
var metaModelData []byte

// Embedded labels in zip format.
//
//go:embed labels.zip
var labelsZip []byte
