package i18n

const (
	// 通用结果
	Success = "common.success"
	Failed  = "common.failed"

	// 新增
	CreateSuccess = "common.createSuccess"
	CreateFailed  = "common.createFailed"

	// 修改
	UpdateSuccess = "common.updateSuccess"
	UpdateFailed  = "common.updateFailed"

	// 删除
	DeleteSuccess = "common.deleteSuccess"
	DeleteFailed  = "common.deleteFailed"

	// 通用错误
	TargetNotFound   = "common.targetNotExist"
	DatabaseError    = "common.databaseError"
	RedisError       = "common.redisError"
	CacheError       = "common.cacheError"
	ConstraintError  = "common.constraintError"
	ValidationError  = "common.validationError"
	NotSingularError = "common.notSingularError"
	PermissionDeny   = "common.permissionDeny"

	// 服务状态
	ServiceUnavailable = "common.serviceUnavailable"
	ServiceBusy        = "common.serviceBusy"

	// 初始化
	AlreadyInit = "init.alreadyInit"
	InitRunning = "init.initializeIsRunning"

	// API
	ApiRequestFailed = "sys.api.apiRequestFailed"
)
