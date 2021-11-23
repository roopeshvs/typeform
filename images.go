package gotypeform

import (
	"bytes"
	"encoding/json"
)

type Images struct {
	FileName string `json:"file_name"`
	ID       string `json:"id"`
	Src      string `json:"src"`
}

type Image struct {
	ID        string `json:"id"`
	Src       string `json:"src"`
	FileName  string `json:"file_name"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	MediaType string `json:"media_type"`
	HasAlpha  bool   `json:"has_alpha"`
	AvgColor  string `json:"avg_color"`
}

type CreateImageRequestBody struct {
	FileName string `json:"file_name"`
	Image    string `json:"image,omitempty"`
	Url      string `json:"url,omitempty"`
}

func (t *Typeform) GetImage(id string) (*Image, error) {
	contents, err := t.buildAndExecRequest("GET", imagesUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	var image *Image
	err = json.Unmarshal(contents, &image)
	if err != nil {
		return nil, err
	}
	return image, err
}

func (t *Typeform) GetAllImages() (*[]Images, error) {
	contents, err := t.buildAndExecRequest("GET", imagesUrl, nil)
	if err != nil {
		return nil, err
	}
	var images *[]Images
	err = json.Unmarshal(contents, &images)
	if err != nil {
		return nil, err
	}
	return images, err
}

func (t *Typeform) CreateImage(img CreateImageRequestBody) (*Image, error) {
	requestBody := &img
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	contents, err := t.buildAndExecRequest("POST", imagesUrl, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	var image *Image
	err = json.Unmarshal(contents, &image)
	if err != nil {
		return nil, err
	}
	return image, err
}

func (t *Typeform) DeleteImage(id string) error {
	_, err := t.buildAndExecRequest("DELETE", imagesUrl+"/"+id, nil)
	return err
}
