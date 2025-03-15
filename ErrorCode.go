package Error

// 错误代码
const (
	ErrorSuccess = iota

	ErrorLogicBegin = 999 + iota

	// *******************************************
	ErrorNoArguments             // 没有参数
	ErrorNoID                    // 没有ID
	ErrorNoName                  // 没有名字
	ErrorNoMobile                // 没有手机号
	ErrorNoAddress               // 没有地址
	ErrorNoContent               // 没有内容
	ErrorNoCallbackURL           // 没有回调地址
	ErrorNoPassword              // 没有密码
	ErrorNoFromPlatform          // 没有来源平台
	ErrorNoToken                 // 没有Token
	ErrorNoValidCode             // 没有验证码
	ErrorMissingArguments        // 缺少参数
	ErrorInvalidData             // 无效数据
	ErrorInvalidDate             // 无效日期
	ErrorInvalidTimeStamp        // 无效时间戳
	ErrorInvalidService          // 无效服务
	ErrorInvalidServiceName      // 无效服务名
	ErrorInvalidServiceIPAddress // 无效服务地址
	ErrorInvalidServiceAuthority // 无效服务权限
	ErrorInvalidAuthority        // 无效权限
	ErrorInvalidPhoneNumber      // 无效手机号
	ErrorInvalidVeriyCode        // 无效验证码
	ErrorInvalidAccount          // 无效账户
	ErrorInvalidToken            // 无效令牌
	ErrorInvalidArguments        // 无效参数
	ErrorFailedReadFile          // 读取文件失败
	ErrorFailedWriteFile         // 写入文件失败
	ErrorFailedConnectDB         // 连接数据库失败
	ErrorFailedInsertData        // 插入数据失败
	ErrorFailedQueryData         // 请求数据失败
	ErrorFailedUpdateData        // 更新数据失败
	ErrorFailedDeleteData        // 删除数据失败
	ErrorEmptyData               // 数据为空
	ErrorNotEnough               // 数据不足
	ErrorFailedVisitURL          // 访问地址失败
	ErrorFailedParseJSON         // 解析格式失败
	ErrorFailedParseData         // 解析数据失败
	ErrorFailedWriteRedis        // 保存缓存失败
	ErrorFailedReadRedis         // 读取缓存失败
	ErrorFailedSendMail          // 发送邮件失败
	ErrorFailedLogin             // 登录失败
	ErrorAccountIsExist          // 账户已存在
	ErrorAccountNotExist         // 账户不存在
	ErrorServiceIsExist          // 服务已存在
	ErrorServiceNotExist         // 服务不存在
	ErrorServiceAuthorityExist   // 服务权限已存在
	ErrorPasswordNotMatch        // 密码不匹配
	ErrorSignatureError          // 签名失败
	ErrorConfigNoChange          // 配置无改变
	ErrorLoginIsInPostDomain     // Login时在后域
	ErrorRoomFullyBooked         // 房间已满
	ErrorIsFreezing              // 冷却中
	ErrorNotSupport              // 不支持
	ErrorSystemError             // 系统错误
)
