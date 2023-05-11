package errcode

// 只登录没有选择企业的token 前缀
var Prefix_token = "$temp$_"

// 系统错误
const LGERROR = "LOGIN_9999"

// 验证码错误
const LGERR_CAPERR = "LOGIN_001"

// 账号和密码不能为空
const LGERR_ACCNULL = "LOGIN_002"

// 账户不存在
const LGERR_NOACCOUNT = "LOGIN_003"

// 账户无效
const LGERR_ACCDISCARD = "LOGIN_004"

// 锁定账户
const LGERR_ACCLOCK = "LOGIN_005"

// 冻结账户
const LGERR_ACCFROZE = "LOGIN_006"

// 用户已过期
const LGERR_ACCEXPIRE = "LOGIN_007"

// 密码错误
const LGERR_ACCPASS = "LOGIN_008"

// 当前账号没有选择公司
const LGERR_ACCNOCOMP = "LOGIN_009"

// 非法请求
const LGERR_INVALID = "LOGIN_010"
