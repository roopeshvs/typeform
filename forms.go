package gotypeform

import (
	"bytes"
	"encoding/json"
)

type Forms struct {
	TotalItems int         `json:"total_items"`
	PageCount  int         `json:"page_count"`
	Items      []FormItems `json:"items"`
}
type FormItems struct {
	Links         FormLinks `json:"_links"`
	ID            string    `json:"id"`
	LastUpdatedAt string    `json:"last_updated_at"`
	Self          Self      `json:"self"`
	Theme         ThemeHref `json:"theme"`
	Title         string    `json:"title"`
}
type FormLinks struct {
	Display string `json:"display"`
}
type Form struct {
	CuiSettings     *CuiSettings      `json:"cui_settings,omitempty"`
	Fields          *Fields           `json:"fields,omitempty"`
	Hidden          *[]string         `json:"hidden,omitempty"`
	Logic           *[]Logic          `json:"logic,omitempty"`
	Settings        *Settings         `json:"settings,omitempty"`
	ThankyouScreens *[]ThankyouScreen `json:"thankyou_screens,omitempty"`
	Theme           *ThemeHref        `json:"theme,omitempty"`
	Title           string            `json:"title"`
	Type            string            `json:"type"`
	Variables       *Variables        `json:"variables,omitempty"`
	WelcomeScreens  *[]WelcomeScreen  `json:"welcome_screens,omitempty"`
	Workspace       *WorkspaceHref    `json:"workspace,omitempty"`
}
type Fields struct {
	Attachment  Attachment       `json:"attachment,omitempty"`
	FieldType   string           `json:"field_type,omitempty"`
	ID          string           `json:"id,omitempty"`
	Layout      Layout           `json:"layout,omitempty"`
	Name        string           `json:"name,omitempty"`
	Options     Options          `json:"options,omitempty"`
	Properties  FieldsProperties `json:"properties,omitempty"`
	Ref         string           `json:"ref,omitempty"`
	Required    bool             `json:"required,omitempty"`
	Title       string           `json:"title"`
	Type        string           `json:"type"`
	Validations Validations      `json:"validations,omitempty"`
}
type FieldsProperties struct {
	AllowMultipleSelection bool    `json:"allow_multiple_selection"`
	AllowOtherChoice       bool    `json:"allow_other_choice"`
	AlphabeticalOrder      bool    `json:"alphabetical_order"`
	ButtonText             string  `json:"button_text"`
	Choices                Choices `json:"choices"`
	Currency               string  `json:"currency"`
	DefaultCountryCode     string  `json:"default_country_code"`
	Description            string  `json:"description"`
	Fields                 *Fields `json:"fields"`
	HideMarks              bool    `json:"hide_marks"`
	Labels                 Labels  `json:"labels"`
	Price                  Price   `json:"price"`
	Randomize              bool    `json:"randomize"`
	Separator              string  `json:"separator"`
	Shape                  string  `json:"shape"`
	ShowButton             bool    `json:"show_button"`
	ShowLabels             bool    `json:"show_labels"`
	StartAtOne             bool    `json:"start_at_one"`
	Steps                  int     `json:"steps"`
	Structure              string  `json:"structure"`
	Supersized             bool    `json:"supersized"`
	VerticalAlignment      bool    `json:"vertical_alignment"`
}
type Options struct {
	Label string `json:"label,omitempty"`
}
type Choices struct {
	Label      string            `json:"label,omitempty"`
	Ref        string            `json:"ref,omitempty"`
	Attachment ChoicesAttachment `json:"attachment,omitempty"`
}
type ChoicesAttachment struct {
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`
}
type Labels struct {
	Center string `json:"center,omitempty"`
	Left   string `json:"left,omitempty"`
	Right  string `json:"right,omitempty"`
}
type Price struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Validations struct {
	MaxLength    int  `json:"max_length"`
	MaxSelection int  `json:"max_selection"`
	MaxValue     int  `json:"max_value"`
	MinSelection int  `json:"min_selection"`
	MinValue     int  `json:"min_value"`
	Required     bool `json:"required"`
}

// ToDo : Fields in Form
type CuiSettings struct {
	Avatar                    string `json:"avatar,omitempty"`
	IsTypingEmulationDisabled bool   `json:"is_typing_emulation_disabled,omitempty"`
	TypingEmulationSpeed      string `json:"typing_emulation_speed,omitempty"`
}
type Vars struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Condition struct {
	Op   string `json:"op"`
	Vars []Vars `json:"vars"`
}
type To struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Target struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Value struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}
type Details struct {
	To     To     `json:"to"`
	Target Target `json:"target"`
	Value  Value  `json:"value"`
}
type Actions struct {
	Action    string    `json:"action"`
	Condition Condition `json:"condition"`
	Details   Details   `json:"details,omitempty"`
}
type Logic struct {
	Actions []Actions `json:"actions"`
	Ref     string    `json:"ref,omitempty"`
	Type    string    `json:"type"`
}
type ImageHref struct {
	Href string `json:"href"`
}
type Meta struct {
	AllowIndexing bool      `json:"allow_indexing,omitempty"`
	Description   string    `json:"description,omitempty"`
	Image         ImageHref `json:"image,omitempty"`
	Title         string    `json:"title,omitempty"`
}
type Respondent struct {
	Enabled   bool     `json:"enabled,omitempty"`
	Message   string   `json:"message,omitempty"`
	Recipient string   `json:"recipient,omitempty"`
	ReplyTo   []string `json:"reply_to,omitempty"`
	Subject   string   `json:"subject,omitempty"`
}
type NotificationsSelf struct {
	Enabled    bool     `json:"enabled,omitempty"`
	Message    string   `json:"message,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	ReplyTo    string   `json:"reply_to,omitempty"`
	Subject    string   `json:"subject,omitempty"`
}
type Notifications struct {
	Respondent Respondent        `json:"respondent,omitempty"`
	Self       NotificationsSelf `json:"self,omitempty"`
}
type Settings struct {
	FacebookPixel          string        `json:"facebook_pixel,omitempty"`
	GoogleAnalytics        string        `json:"google_analytics,omitempty"`
	GoogleTagManager       string        `json:"google_tag_manager,omitempty"`
	HideNavigation         bool          `json:"hide_navigation,omitempty"`
	IsPublic               bool          `json:"is_public,omitempty"`
	Language               string        `json:"language,omitempty"`
	Meta                   Meta          `json:"meta,omitempty"`
	Notifications          Notifications `json:"notifications,omitempty"`
	ProgressBar            string        `json:"progress_bar,omitempty"`
	RedirectAfterSubmitURL string        `json:"redirect_after_submit_url,omitempty"`
	ShowProgressBar        bool          `json:"show_progress_bar,omitempty"`
	ShowTimeToComplete     bool          `json:"show_time_to_complete,omitempty"`
	ShowTypeformBranding   bool          `json:"show_typeform_branding,omitempty"`
}
type Attachment struct {
	Href       string               `json:"href,omitempty"`
	Type       string               `json:"type,omitempty"`
	Scale      int                  `json:"scale,omitempty"`
	Properties AttachmentProperties `json:"properties,omitempty"`
}
type AttachmentProperties struct {
	Description string `json:"description,omitempty"`
}
type ThankYouScreenProperties struct {
	ButtonMode  string `json:"button_mode,omitempty"`
	ButtonText  string `json:"button_text,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	ShareIcons  bool   `json:"share_icons,omitempty"`
	ShowButton  bool   `json:"show_button,omitempty"`
}
type ThankyouScreen struct {
	Attachment Attachment               `json:"attachment,omitempty"`
	Properties ThankYouScreenProperties `json:"properties,omitempty"`
	Ref        string                   `json:"ref,omitempty"`
	Title      string                   `json:"title"`
	Layout     Layout                   `json:"layout,omitempty"`
}
type ThemeHref struct {
	Href string `json:"href,omitempty"`
}
type Variables struct {
	Age   int     `json:"age,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
	Score int     `json:"score,omitempty"`
}
type Layout struct {
	Attachment Attachment `json:"attachment,omitempty"`
	Placement  string     `json:"placement,omitempty"`
	Type       string     `json:"type,omitempty"`
}
type WelcomeScreenProperties struct {
	ButtonText  string `json:"button_text,omitempty"`
	Description string `json:"description,omitempty"`
	ShowButton  bool   `json:"show_button,omitempty"`
}
type WelcomeScreen struct {
	Layout     Layout                  `json:"layout,omitempty"`
	Properties WelcomeScreenProperties `json:"properties,omitempty"`
	Ref        string                  `json:"ref,omitempty"`
	Title      string                  `json:"title"`
	Attachment Attachment              `json:"attachment,omitempty"`
}
type WorkspaceHref struct {
	Href string `json:"href,omitempty"`
}
type FormQueryParams struct {
	Search      string `url:"search,omitempty"`
	Page        int    `url:"page,omitempty"`
	PageSize    int    `url:"page_size,omitempty"`
	WorkspaceID string `url:"workspace_id,omitempty"`
}
type CustomFormMessages struct {
	LabelButtonHintDefault           string `json:"label.buttonHint.default,omitempty"`
	LabelButtonHintLongtext          string `json:"label.buttonHint.longtext,omitempty"`
	LabelWarningConnection           string `json:"label.warning.connection,omitempty"`
	LabelButtonNoAnswerDefault       string `json:"label.buttonNoAnswer.default,omitempty"`
	LabelWarningCorrection           string `json:"label.warning.correction,omitempty"`
	BlockPaymentCardNameTitle        string `json:"block.payment.cardNameTitle,omitempty"`
	BlockPaymentCardNumberTitle      string `json:"block.payment.cardNumberTitle,omitempty"`
	BlockPaymentCvcDescription       string `json:"block.payment.cvcDescription,omitempty"`
	BlockPaymentCvcNumberTitle       string `json:"block.payment.cvcNumberTitle,omitempty"`
	BlockShortTextPlaceholder        string `json:"block.shortText.placeholder,omitempty"`
	LabelErrorEmailAddress           string `json:"label.error.emailAddress,omitempty"`
	LabelErrorExpiryMonthTitle       string `json:"label.error.expiryMonthTitle,omitempty"`
	LabelErrorExpiryYearTitle        string `json:"label.error.expiryYearTitle,omitempty"`
	LabelWarningFallbackAlert        string `json:"label.warning.fallbackAlert,omitempty"`
	BlockFileUploadChoose            string `json:"block.fileUpload.choose,omitempty"`
	BlockFileUploadDrag              string `json:"block.fileUpload.drag,omitempty"`
	BlockFileUploadUploadingProgress string `json:"block.fileUpload.uploadingProgress,omitempty"`
	LabelErrorSizeLimit              string `json:"label.error.sizeLimit,omitempty"`
	LabelWarningFormUnavailable      string `json:"label.warning.formUnavailable,omitempty"`
	LabelErrorIncompleteForm         string `json:"label.error.incompleteForm,omitempty"`
	LabelHintKey                     string `json:"label.hint.key,omitempty"`
	BlockLegalReject                 string `json:"block.legal.reject,omitempty"`
	BlockLegalAccept                 string `json:"block.legal.accept,omitempty"`
	LabelErrorMaxValue               string `json:"label.error.maxValue,omitempty"`
	LabelErrorMaxLength              string `json:"label.error.maxLength,omitempty"`
	LabelErrorMinValue               string `json:"label.error.minValue,omitempty"`
	LabelErrorRange                  string `json:"label.error.range,omitempty"`
	BlockMultipleChoiceHint          string `json:"block.multipleChoice.hint,omitempty"`
	LabelErrorMustEnter              string `json:"label.error.mustEnter,omitempty"`
	LabelErrorMustSelect             string `json:"label.error.mustSelect,omitempty"`
	LabelNoShortcut                  string `json:"label.no.shortcut,omitempty"`
	LabelNoDefault                   string `json:"label.no.default,omitempty"`
	BlockDropdownHint                string `json:"block.dropdown.hint,omitempty"`
	BlockMultipleChoiceOther         string `json:"block.multipleChoice.other,omitempty"`
	LabelProgressPercent             string `json:"label.progress.percent,omitempty"`
	LabelProgressProportion          string `json:"label.progress.proportion,omitempty"`
	LabelErrorRequired               string `json:"label.error.required,omitempty"`
	LabelPreview                     string `json:"label.preview,omitempty"`
	LabelButtonReview                string `json:"label.button.review,omitempty"`
	LabelErrorServer                 string `json:"label.error.server,omitempty"`
	LabelActionShare                 string `json:"label.action.share,omitempty"`
	LabelButtonSubmit                string `json:"label.button.submit,omitempty"`
	LabelWarningSuccess              string `json:"label.warning.success,omitempty"`
	LabelButtonOk                    string `json:"label.button.ok,omitempty"`
	LabelErrorMustAccept             string `json:"label.error.mustAccept,omitempty"`
	BlockLongtextHint                string `json:"block.longtext.hint,omitempty"`
	BlockDropdownPlaceholder         string `json:"block.dropdown.placeholder,omitempty"`
	BlockDropdownPlaceholderTouch    string `json:"block.dropdown.placeholderTouch,omitempty"`
	LabelErrorURL                    string `json:"label.error.url,omitempty"`
	LabelYesShortcut                 string `json:"label.yes.shortcut,omitempty"`
	LabelYesDefault                  string `json:"label.yes.default,omitempty"`
}

func (t *Typeform) CreateForm(formdata Form) (*Form, error) {
	requestBody := &formdata
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	contents, err := t.buildAndExecRequest("POST", formsUrl, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	var form *Form
	err = json.Unmarshal(contents, &form)
	if err != nil {
		return nil, err
	}
	return form, nil
}

func (t *Typeform) GetForm(id string) (*Form, error) {
	contents, err := t.buildAndExecRequest("GET", formsUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	var form *Form
	err = json.Unmarshal(contents, &form)
	if err != nil {
		return nil, err
	}
	return form, nil
}

func (t *Typeform) GetAllForms(params FormQueryParams) (*Forms, error) {
	contents, err := t.buildAndExecRequest("GET", formsUrl, nil)
	if err != nil {
		return nil, err
	}
	var forms *Forms
	err = json.Unmarshal(contents, &forms)
	if err != nil {
		return nil, err
	}
	return forms, nil
}

func (t *Typeform) UpdateForm(id string, updates []UpdateRequestBody) error {
	json_data, err := json.Marshal(updates)
	if err != nil {
		return err
	}
	_, err = t.buildAndExecRequest("PATCH", formsUrl+"/"+id, bytes.NewBuffer(json_data))
	if err != nil {
		return nil
	}
	return nil
}

func (t *Typeform) DeleteForm(id string) error {
	_, err := t.buildAndExecRequest("DELETE", formsUrl+"/"+id, nil)
	if err != nil {
		return err
	}
	return nil
}

func (t *Typeform) GetCustomFormMessages(id string) (*CustomFormMessages, error) {
	contents, err := t.buildAndExecRequest("GET", formsUrl+"/"+id+"/messages", nil)
	if err != nil {
		return nil, err
	}
	var messages *CustomFormMessages
	err = json.Unmarshal(contents, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (t *Typeform) UpdateCustomFormMessages(id string, messages CustomFormMessages) error {
	json_data, err := json.Marshal(messages)
	if err != nil {
		return err
	}
	_, err = t.buildAndExecRequest("PUT", formsUrl+"/"+id+"/messages", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	return nil
}
