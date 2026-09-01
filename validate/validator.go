package validate

import (
	"errors"
	"net/http"
	"strings"

	enLang "github.com/go-playground/locales/en"
	zhLang "github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"golang.org/x/text/language"
)

const (
	langZh = "zh"
	langEn = "en"
)

var supportLang = map[string]string{
	"zh":    langZh,
	"zh-CN": langZh,
	"en":    langEn,
	"en-US": langEn,
}

// Validator 参数校验器
type Validator struct {
	validator *validator.Validate
	trans     map[string]ut.Translator
}

// New 创建参数校验器
func New() (*Validator, error) {
	en := enLang.New()
	zh := zhLang.New()

	uni := ut.New(zh, en, zh)

	enTrans, _ := uni.GetTranslator(langEn)
	zhTrans, _ := uni.GetTranslator(langZh)

	v := &Validator{
		validator: validator.New(),
		trans: map[string]ut.Translator{
			langEn: enTrans,
			langZh: zhTrans,
		},
	}

	if err := enTranslations.RegisterDefaultTranslations(v.validator, enTrans); err != nil {
		return nil, err
	}

	if err := zhTranslations.RegisterDefaultTranslations(v.validator, zhTrans); err != nil {
		return nil, err
	}

	return v, nil
}

// Validate 校验请求参数
func (v *Validator) Validate(r *http.Request, data any) error {
	if err := v.validator.Struct(data); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		lang := parseAcceptLanguage(r.Header.Get("Accept-Language"))
		trans := v.trans[lang]

		messages := make([]string, 0, len(validationErrors))
		for _, validationError := range validationErrors {
			messages = append(messages, validationError.Translate(trans))
		}

		return errors.New(strings.Join(messages, " "))
	}

	return nil
}

// RegisterValidation 注册自定义校验规则
func (v *Validator) RegisterValidation(tag string, fn validator.Func) error {
	return v.validator.RegisterValidation(tag, fn)
}

// RegisterValidationTranslation 注册自定义校验规则翻译
func (v *Validator) RegisterValidationTranslation(
	tag string,
	lang string,
	registerFn validator.RegisterTranslationsFunc,
	translationFn validator.TranslationFunc,
) error {
	trans, ok := v.trans[lang]
	if !ok {
		return errors.New("不支持的校验语言")
	}

	return v.validator.RegisterTranslation(
		tag,
		trans,
		registerFn,
		translationFn,
	)
}

// parseAcceptLanguage 解析请求语言
func parseAcceptLanguage(value string) string {
	if strings.TrimSpace(value) == "" {
		return langZh
	}

	tags, _, err := language.ParseAcceptLanguage(value)
	if err != nil {
		return langZh
	}

	for _, tag := range tags {
		if lang, ok := supportLang[tag.String()]; ok {
			return lang
		}
	}

	return langZh
}
