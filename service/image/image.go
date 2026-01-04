package image

import (
	"io"
	"log"
	"mime/multipart"
	"shc-ai-demo/common/image"
)

func RecognizeImage(file *multipart.FileHeader) (string, error) {

	modelPath := "/home/byte/develop/model/mobilenetv2-7.onnx"
	labelPath := "/home/byte/develop/model/imagenet_classes.txt"
	inputH, inputW := 224, 224

	recognizer, err := image.NewImageRecognizer(modelPath, labelPath, inputH, inputW)
	if err != nil {
		log.Println("NewImageRecognizer fail err is : ", err)
		return "", err
	}
	defer recognizer.Close()

	src, err := file.Open()
	if err != nil {
		log.Println("file open fail err is : ", err)
		return "", err
	}
	defer src.Close()

	buf, err := io.ReadAll(src)
	if err != nil {
		log.Println("io.ReadAll fail err is : ", err)
		return "", err
	}

	return recognizer.PredictFromBuffer(buf)
}
