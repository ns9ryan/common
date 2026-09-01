package casbin

import (
	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	"github.com/suyuan32/simple-admin-common/plugins/casbin/entadapter"
	"github.com/zeromicro/go-zero/core/logx"
)

const defaultModelText = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act
`

// CasbinConf Casbin 配置
type CasbinConf struct {
	ModelText string `json:",optional,env=CASBIN_MODEL_TEXT"` // Casbin 模型配置
}

// NewCasbin 创建 Casbin 权限执行器
func (c CasbinConf) NewCasbin(dbType, dsn string) (*casbin.Enforcer, error) {
	// 创建 Casbin 数据库适配器
	adapter, err := entadapter.NewAdapter(dbType, dsn)
	logx.Must(err)

	// 使用自定义模型，为空时使用默认模型
	modelText := c.ModelText
	if modelText == "" {
		modelText = defaultModelText
	}

	// 创建 Casbin 模型
	casbinModel, err := model.NewModelFromString(modelText)
	logx.Must(err)

	// 创建 Casbin 权限执行器
	enforcer, err := casbin.NewEnforcer(casbinModel, adapter)
	logx.Must(err)

	// 加载 Casbin 权限策略
	err = enforcer.LoadPolicy()

	return enforcer, nil
}
