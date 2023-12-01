package g

import (
	"fmt"
	"github.com/ecnu/ecnu-openapi-sdk-go/sdk"
	"path/filepath"
	"strings"
)

func SyncWithConfig() error {
	sdk.InitOAuth2ClientCredentials(Config().OAuth2Config)

	file := Config().OutputFile
	// 获取文件后缀名
	extention := filepath.Ext(file)
	if extention != "" {
		// Remove the leading dot (.) from the extension.
		extention = extention[1:]
	}
	var mode string
	switch extention {
	case "xlsx":
		fallthrough
	case "XLSX":
		mode = "xlsx"
	default:
		mode = "csv"
	}
	rows, err := sdk.SyncToFile(mode, file, Config().APIConfig)
	if err != nil {
		return err
	}
	fmt.Printf("%s：组织机构同步 %d 条数据\n", strings.ToUpper(mode), rows)
	return nil
}
func SyncWithoutConfig(clientId *string, clientSecret *string, output *string, apiPath *string) error {
	cf := sdk.OAuth2Config{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
	}
	sdk.InitOAuth2ClientCredentials(cf)
	file := *output
	api := sdk.APIConfig{
		APIPath:  *apiPath,
		PageSize: 2000,
	}
	api.SetParam("ts", "0")
	// 获取文件后缀名
	extention := filepath.Ext(file)
	if extention != "" {
		// Remove the leading dot (.) from the extension.
		extention = extention[1:]
	}
	var mode string
	switch extention {
	case "xlsx":
		fallthrough
	case "XLSX":
		mode = "xlsx"
	default:
		mode = "csv"
	}
	rows, err := sdk.SyncToFile(mode, file, api)
	if err != nil {
		return err
	}
	fmt.Printf("%s：组织机构同步 %d 条数据\n", strings.ToUpper(mode), rows)
	return nil
}
