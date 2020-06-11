package auth

import (
	"github.com/fatedier/frp/models/msg"
	"github.com/vaughan0/go-ini"
)

type oauth2Config struct {
	// Token specifies the authorization token used to create keys to be sent
	// to the server. The server must have a matching token for authorization
	// to succeed.  By default, this value is "".
	Token string `json:"token"`
}

type Auth2SetterVerifier struct {
	baseConfig
}

func (a Auth2SetterVerifier) VerifyLogin(*msg.Login) error {
	/*	if util.GetAuthKey(auth.token, loginMsg.Timestamp) != loginMsg.PrivilegeKey {
			return fmt.Errorf("token in login doesn't match token from configuration")
		}
		return nil*/
	panic("implement me")
}

func (a Auth2SetterVerifier) VerifyPing(*msg.Ping) error {
	/*	if !auth.AuthenticateHeartBeats {
			return nil
		}

		if util.GetAuthKey(auth.token, pingMsg.Timestamp) != pingMsg.PrivilegeKey {
			return fmt.Errorf("token in heartbeat doesn't match token from configuration")
		}
		return nil*/
	panic("implement me")
}

func (a Auth2SetterVerifier) VerifyNewWorkConn(*msg.NewWorkConn) error {
	/*	if !auth.AuthenticateNewWorkConns {
			return nil
		}

		if util.GetAuthKey(auth.token, newWorkConnMsg.Timestamp) != newWorkConnMsg.PrivilegeKey {
			return fmt.Errorf("token in NewWorkConn doesn't match token from configuration")
		}
		return nil*/
	panic("implement me")
}

func NewOAuth2Verifie(baseCfg baseConfig, cfg oauth2ServerConfig) *Auth2SetterVerifier {
	return &Auth2SetterVerifier{
		baseConfig: baseCfg,
	}
}

type oauth2ServerConfig struct {
	Oauth2Url string `json:"oauth2_url"`
}

func unmarshalAuth2ConfFromIni(conf ini.File) oauth2ServerConfig {
	var (
		tmpStr string
		ok     bool
	)

	cfg := getDefaultAuth2Conf()

	if tmpStr, ok = conf.Get("common", "oauth2_url"); ok {
		cfg.Oauth2Url = tmpStr
	}

	return cfg
}
func getDefaultAuth2Conf() oauth2ServerConfig {
	return oauth2ServerConfig{
		Oauth2Url: "",
	}
}
