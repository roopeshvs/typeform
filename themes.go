package gotypeform

import (
	"bytes"
	"encoding/json"
)

type Themes struct {
	TotalItems int     `json:"total_items"`
	PageCount  int     `json:"page_count"`
	Items      []Theme `json:"items"`
}

type Theme struct {
	Background           *Background  `json:"background,omitempty"`
	Colors               *Colors      `json:"colors,omitempty"`
	Fields               *ThemeFields `json:"fields,omitempty"`
	Font                 string       `json:"font,omitempty"`
	HasTransparentButton bool         `json:"has_transparent_button"`
	ID                   string       `json:"id,omitempty"`
	Name                 string       `json:"name"`
	Screens              *Screens     `json:"screens,omitempty"`
	Visibility           string       `json:"visibility,omitempty"`
}
type Background struct {
	Brightness float64 `json:"brightness"`
	Href       string  `json:"href"`
	Layout     string  `json:"layout"`
}
type Colors struct {
	Answer     string `json:"answer"`
	Background string `json:"background"`
	Button     string `json:"button"`
	Question   string `json:"question"`
}
type ThemeFields struct {
	Alignment string `json:"alignment"`
	FontSize  string `json:"font_size"`
}
type Screens struct {
	Alignment string `json:"alignment"`
	FontSize  string `json:"font_size"`
}
type ThemeQueryParams struct {
	Page     int `url:"page,omitempty"`
	PageSize int `url:"page_size,omitempty"`
}

func (t *Typeform) GetTheme(id string) (*Theme, error) {
	contents, err := t.buildAndExecRequest("GET", themesUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	var theme *Theme
	err = json.Unmarshal(contents, &theme)
	if err != nil {
		return nil, err
	}
	return theme, nil
}

func (t *Typeform) CreateTheme(themedata Theme) (*Theme, error) {
	requestBody := &themedata
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	contents, err := t.buildAndExecRequest("POST", themesUrl, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	var theme *Theme
	err = json.Unmarshal(contents, &theme)
	if err != nil {
		return nil, err
	}
	return theme, nil
}

func (t *Typeform) DeleteTheme(id string) error {
	_, err := t.buildAndExecRequest("DELETE", themesUrl+"/"+id, nil)
	return err
}

func (t *Typeform) UpdateTheme(id string, theme Theme) error {
	requestBody := &theme
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	_, err = t.buildAndExecRequest("PUT", themesUrl+"/"+id, bytes.NewBuffer(json_data))
	return err
}

func (t *Typeform) GetAllThemes(params ThemeQueryParams) (*Themes, error) {
	requestBody := &params
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	contents, err := t.buildAndExecRequest("GET", themesUrl, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	var themes *Themes
	err = json.Unmarshal(contents, &themes)
	if err != nil {
		return nil, err
	}
	return themes, nil
}
