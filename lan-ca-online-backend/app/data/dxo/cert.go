package dxo

import (
	"encoding/json"

	"github.com/starter-go/base/lang"
)

// Fingerprint 表示证书或密钥的指纹
type Fingerprint lang.Hex

////////////////////////////////////////////////////////////////////////////////

// CertificateState 是一个字符串, 表示证书的当前状态
type CertificateState string

func (cs CertificateState) String() string {
	return string(cs)
}

const (
	CertificateStateInit      = "initialled" // 初始状态
	CertificateStateSubmitted = "submitted"  // 已经提交申请
	CertificateStateSigned    = "signed"     // 已经签署
	CertificateStateDeployed  = "deployed"   // 已经部署
	CertificateStateExpired   = "expired"    // 已经过期
)

////////////////////////////////////////////////////////////////////////////////

type CertificateUserInfo struct {
	CN    string `json:"cn"`    // Common-Name
	C     string `json:"c"`     // Country
	O     string `json:"o"`     // Organization
	OU    string `json:"ou"`    // Organization-Unit
	L     string `json:"l"`     // Locality (city)
	ST    string `json:"st"`    // State|Province
	Email string `json:"email"` // Email-Address
}

// JSON 函数把该对象转换成 JSON-string 形式
func (inst *CertificateUserInfo) JSON() CertificateUserInfoJSON {
	if inst == nil {
		return "{}"
	}
	bin, err := json.Marshal(inst)
	if err != nil {
		return "{}"
	}
	return CertificateUserInfoJSON(bin)
}

////////////////////////////////////////////////////////////////////////////////

// CertificateUserInfoJSON 是一个字符串, 表示以 JSON 形式存储的 CertificateUserInfo
type CertificateUserInfoJSON string

// Info 函数把该字符串解析成 CertificateUserInfo 对象
func (str CertificateUserInfoJSON) Info() *CertificateUserInfo {
	info := &CertificateUserInfo{}
	bin := []byte(str)
	json.Unmarshal(bin, info)
	return info
}

////////////////////////////////////////////////////////////////////////////////
