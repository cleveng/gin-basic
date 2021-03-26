package config

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
	"regexp"
	"strings"
)

var Validate *validator.Validate
var trans ut.Translator

//func init() {
//	zh_CN := zhongwen.New()
//	uni := ut.New(zh_CN)
//	trans, _ := uni.GetTranslator("zh")
//	Validate = validator.New()
//	zh_translations.RegisterDefaultTranslations(Validate, trans)
//}

/**
验证规则
required ：必填
email：验证字符串是email格式；例：“email”
url：这将验证字符串值包含有效的网址;例：“url”
max：字符串最大长度；例：“max=20”
min:字符串最小长度；例：“min=6”
excludesall:不能包含特殊字符；例：“excludesall=0x2C”//注意这里用十六进制表示。
len：字符长度必须等于n，或者数组、切片、map的len值为n，即包含的项目数；例：“len=6”
eq：数字等于n，或者或者数组、切片、map的len值为n，即包含的项目数；例：“eq=6”
ne：数字不等于n，或者或者数组、切片、map的len值不等于为n，即包含的项目数不为n，其和eq相反；例：“ne=6”
gt：数字大于n，或者或者数组、切片、map的len值大于n，即包含的项目数大于n；例：“gt=6”
gte：数字大于或等于n，或者或者数组、切片、map的len值大于或等于n，即包含的项目数大于或等于n；例：“gte=6”
lt：数字小于n，或者或者数组、切片、map的len值小于n，即包含的项目数小于n；例：“lt=6”
lte：数字小于或等于n，或者或者数组、切片、map的len值小于或等于n，即包含的项目数小于或等于n；例：“lte=6”

=====
跨字段验证
eqfield=Field: 必须等于 Field 的值；
nefield=Field: 必须不等于 Field 的值；
gtfield=Field: 必须大于 Field 的值；
gtefield=Field: 必须大于等于 Field 的值；
ltfield=Field: 必须小于 Field 的值；
ltefield=Field: 必须小于等于 Field 的值；
eqcsfield=Other.Field: 必须等于 struct Other 中 Field 的值；
necsfield=Other.Field: 必须不等于 struct Other 中 Field 的值；
gtcsfield=Other.Field: 必须大于 struct Other 中 Field 的值；
gtecsfield=Other.Field: 必须大于等于 struct Other 中 Field 的值；
ltcsfield=Other.Field: 必须小于 struct Other 中 Field 的值；
ltecsfield=Other.Field: 必须小于等于 struct Other 中 Field 的值；

Passwd   string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
Repasswd   string `form:"repasswd" json:"repasswd" validate:"required,max=20,min=6,eqfield=Passwd"`

*/


func InitValidate() {
	zhCN := zh.New()
	uni := ut.New(zhCN)
	trans, _ = uni.GetTranslator("zh")
	Validate = validator.New()
	Validate.RegisterTranslation("mobile", trans, verifyTel, translatorTel)
	// 注册 RegisterTagNameFunc
	Validate.RegisterTagNameFunc(comment)
	Validate.RegisterValidation("mobile", isMobilePhone) // 注册tag验证器
	zh_trans.RegisterDefaultTranslations(Validate, trans)
}

func comment(field reflect.StructField) string {
	return field.Tag.Get("comment")
}

// 校验是否是手机号码
func isMobilePhone(field validator.FieldLevel) bool {
	tel := field.Field().String()
	if len(tel) != 11 {
		return false
	}
	isMobile, _ := regexp.MatchString(`^((\+|00)86)?1([3568][0-9]|4[579]|6[67]|7[01235678]|9[012356789])[0-9]{8}$`, tel)
	return isMobile
}

func verifyTel(ut ut.Translator) error {
	if err := ut.Add("mobile", "{0}格式不正确~", false); err != nil {
		return err
	}
	return nil
}
func translatorTel(ut ut.Translator, fe validator.FieldError) string{
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("警告: 翻译字段错误: %#v", fe)
		return fe.(error).Error()
	}
	return t
}



func Translate(errs validator.ValidationErrors) string {
	var errList []string
	err_key := []string{"Zh_CN","简体","Zh_Word","原文字"}	// 只能治标 不能治本。后期需要调整
	for _, e := range errs {
		re := strings.NewReplacer(err_key...)
		content := re.Replace(e.Translate(trans))
		errList = append(errList, content)
	}
	//[验证码为必填字段 邮箱为必填字段]
	return strings.Join(errList, "|")	// 字符串拼接
}
