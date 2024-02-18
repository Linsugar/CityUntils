package Untils

//
//import (
//	"fmt"
//
//)
//
//type ImageOcrInter interface {
//	ImageOcrToS() (data any)
//}
//
//type ImageConf struct {
//}
//
//func (i *ImageConf) ImageOcrToS() (data any) {
//	client := gosseract.NewClient()
//	defer client.Close()
//	client.SetImage("image.jpg") // 替换为您要识别的图片路径
//	text, err := client.Text()
//	if err != nil {
//		fmt.Println("OCR error:", err)
//		return
//	}
//	fmt.Println("识别结果:", text)
//	return text
//}
