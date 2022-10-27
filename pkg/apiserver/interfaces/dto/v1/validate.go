package apis

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

const (
	LanguageChinese = "zh"
	LanguageEnglish = "en"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	zh_trans ut.Translator
	en_trans ut.Translator
)

func init() {
	validate = validator.New()
	zh := zh.New()
	en := en.New()
	uni = ut.New(en, zh)

	zh_trans, _ = uni.GetTranslator(LanguageChinese)
	en_trans, _ = uni.GetTranslator(LanguageEnglish)
	zh_translations.RegisterDefaultTranslations(validate, zh_trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Validate(ctx context.Context, i interface{}) map[string]string {
	err := validate.StructCtx(ctx, i)
	errs := err.(validator.ValidationErrors)
	if len(errs) == 0 {
		return nil
	}
	var trans validator.ValidationErrorsTranslations
	v := ctx.Value(options.LanguageKey)
	if v != nil && v.(string) == LanguageChinese {
		trans = errs.Translate(zh_trans)
	} else {
		trans = errs.Translate(en_trans)
	}
	return trans
}
