package wechat

import "github.com/pkg/errors"

const (
	WxCorpId         = "ww797fe1e6262986c1"
	WxToken          = "Q9a9XwRVj54lWitrCR9Lom7Q3mk"
	WxEncodingAESKey = "CHTVkHtIwggHRSFXzDpKoxtuikMVHFxlUqcXyIAx4PS"
)

func Verify(msgSign, msgTs, msgNonce, msgEchoStr string) (string, error) {
	wxCpt := NewWXBizMsgCrypt(WxToken, WxEncodingAESKey, WxCorpId, XmlType)

	echoStr, cryptErr := wxCpt.VerifyURL(msgSign, msgTs, msgNonce, msgEchoStr)
	if cryptErr != nil {
		return string(echoStr), errors.Errorf("%s(%d)", cryptErr.ErrMsg, cryptErr.ErrCode)
	}
	return string(echoStr), nil
}
