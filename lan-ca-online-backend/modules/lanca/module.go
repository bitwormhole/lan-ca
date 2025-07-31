package lanca

import (
	lancaonlinebackend "github.com/bitwormhole/lan-ca/backend"
	"github.com/bitwormhole/lan-ca/backend/gen/main4lanca"
	"github.com/bitwormhole/lan-ca/backend/gen/test4lanca"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security-gin-gorm/modules/securitygingorm"
	"github.com/starter-go/security/modules/security"
)

// Module  ...
func Module() application.Module {
	mb := lancaonlinebackend.NewMainModule()
	mb.Components(main4lanca.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())
	mb.Depend(securitygingorm.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := lancaonlinebackend.NewTestModule()
	mb.Components(test4lanca.ExportComponents)
	return mb.Create()
}
