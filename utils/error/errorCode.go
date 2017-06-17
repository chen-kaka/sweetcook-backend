package error

const CODE_NAME = "retcode"  // 返回CODE字段名称
const MSG_NAME = "msg"  // 返回MSG字段名称
const DATA_NAME = "data"  // 返回DATA字段名称

const SUCCESS = 0  //成功

const ERROR_LOGIN_FAILED = 10001  //登录失败
const ERROR_PARAM_ERROR = 10002  //参数错误
const ERROR_VALIDATE_ERROR = 10003  //校验错误

const ERROR_SERVICE_ERROR = 20001  //服务错误
const ERROR_INTERFACE_ERROR = 20002  //接口错误
const ERROR_NO_DATA_ERROR = 20003  //无数据错误
const ERROR_DATABASE_ERROR = 20004  //数据库异常