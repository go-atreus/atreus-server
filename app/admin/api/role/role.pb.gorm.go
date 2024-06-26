package role

import (
	context "context"
	fmt "fmt"
	menu "github.com/go-atreus/atreus-server/app/admin/api/menu"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	gorm "gorm.io/gorm"
)

type SysRoleORM struct {
	Code       string
	CreateTime string
	Deleted    string
	Id         int32              `gorm:"type:integer;primaryKey;autoIncrement"`
	Menu       []*menu.SysMenuORM `gorm:"foreignKey:Id;references:Id;many2many:sys_role_menu;joinForeignKey:SysRoleId;joinReferences:SysMenuId"`
	Name       string
	Remarks    string
	ScopeType  int32
	Type       int32
	UpdateTime string
}

// TableName overrides the default tablename generated by GORM
func (SysRoleORM) TableName() string {
	return "sys_roles"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *SysRole) ToORM(ctx context.Context) (SysRoleORM, error) {
	to := SysRoleORM{}
	var err error
	if prehook, ok := interface{}(m).(SysRoleWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.Code = m.Code
	to.Remarks = m.Remarks
	to.Type = m.Type
	to.Deleted = m.Deleted
	to.CreateTime = m.CreateTime
	to.UpdateTime = m.UpdateTime
	to.ScopeType = m.ScopeType
	for _, v := range m.Menu {
		if v != nil {
			if tempMenu, cErr := v.ToORM(ctx); cErr == nil {
				to.Menu = append(to.Menu, &tempMenu)
			} else {
				return to, cErr
			}
		} else {
			to.Menu = append(to.Menu, nil)
		}
	}
	if posthook, ok := interface{}(m).(SysRoleWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SysRoleORM) ToPB(ctx context.Context) (SysRole, error) {
	to := SysRole{}
	var err error
	if prehook, ok := interface{}(m).(SysRoleWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.Code = m.Code
	to.Remarks = m.Remarks
	to.Type = m.Type
	to.Deleted = m.Deleted
	to.CreateTime = m.CreateTime
	to.UpdateTime = m.UpdateTime
	to.ScopeType = m.ScopeType
	for _, v := range m.Menu {
		if v != nil {
			if tempMenu, cErr := v.ToPB(ctx); cErr == nil {
				to.Menu = append(to.Menu, &tempMenu)
			} else {
				return to, cErr
			}
		} else {
			to.Menu = append(to.Menu, nil)
		}
	}
	if posthook, ok := interface{}(m).(SysRoleWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type SysRole the arg will be the target, the caller the one being converted from

// SysRoleBeforeToORM called before default ToORM code
type SysRoleWithBeforeToORM interface {
	BeforeToORM(context.Context, *SysRoleORM) error
}

// SysRoleAfterToORM called after default ToORM code
type SysRoleWithAfterToORM interface {
	AfterToORM(context.Context, *SysRoleORM) error
}

// SysRoleBeforeToPB called before default ToPB code
type SysRoleWithBeforeToPB interface {
	BeforeToPB(context.Context, *SysRole) error
}

// SysRoleAfterToPB called after default ToPB code
type SysRoleWithAfterToPB interface {
	AfterToPB(context.Context, *SysRole) error
}

// DefaultCreateSysRole executes a basic gorm create call
func DefaultCreateSysRole(ctx context.Context, in *SysRole, db *gorm.DB) (*SysRole, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeCreate_); ok {
		if ormObj, db, err = hook.BeforeCreate_(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type SysRoleORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB, SysRoleORM) (SysRoleORM, *gorm.DB, error)
}
type SysRoleORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadSysRole(ctx context.Context, in *SysRole, db *gorm.DB) (*SysRole, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := SysRoleORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(SysRoleORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type SysRoleORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysRoleORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysRoleORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteSysRole(ctx context.Context, in *SysRole, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeDelete_); ok {
		if ormObj, db, err = hook.BeforeDelete_(ctx, db, ormObj); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&SysRoleORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type SysRoleORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB, SysRoleORM) (SysRoleORM, *gorm.DB, error)
}
type SysRoleORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteSysRoleSet(ctx context.Context, in []*SysRole, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []int32{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&SysRoleORM{})).(SysRoleORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&SysRoleORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&SysRoleORM{})).(SysRoleORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type SysRoleORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*SysRole, *gorm.DB) (*gorm.DB, error)
}
type SysRoleORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*SysRole, *gorm.DB) error
}

// DefaultStrictUpdateSysRole clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateSysRole(ctx context.Context, in *SysRole, db *gorm.DB) (*SysRole, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateSysRole")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &SysRoleORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeStrictUpdateCleanup); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateCleanup(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Model(&ormObj).Association("Menu").Replace(ormObj.Menu); err != nil {
		return nil, err
	}
	ormObj.Menu = nil
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeStrictUpdateSave); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateSave(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type SysRoleORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB, SysRoleORM) (SysRoleORM, *gorm.DB, error)
}
type SysRoleORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB, SysRoleORM) (SysRoleORM, *gorm.DB, error)
}
type SysRoleORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchSysRole executes a basic gorm update call with patch behavior
func DefaultPatchSysRole(ctx context.Context, in *SysRole, updateMask *field_mask.FieldMask, db *gorm.DB) (*SysRole, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj SysRole
	var err error
	if hook, ok := interface{}(&pbObj).(SysRoleWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadSysRole(ctx, &SysRole{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(SysRoleWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskSysRole(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(SysRoleWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateSysRole(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(SysRoleWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type SysRoleWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *SysRole, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysRoleWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *SysRole, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysRoleWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *SysRole, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysRoleWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *SysRole, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetSysRole executes a bulk gorm update call with patch behavior
func DefaultPatchSetSysRole(ctx context.Context, objects []*SysRole, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*SysRole, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*SysRole, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchSysRole(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskSysRole patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskSysRole(ctx context.Context, patchee *SysRole, patcher *SysRole, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*SysRole, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Code" {
			patchee.Code = patcher.Code
			continue
		}
		if f == prefix+"Remarks" {
			patchee.Remarks = patcher.Remarks
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Deleted" {
			patchee.Deleted = patcher.Deleted
			continue
		}
		if f == prefix+"CreateTime" {
			patchee.CreateTime = patcher.CreateTime
			continue
		}
		if f == prefix+"UpdateTime" {
			patchee.UpdateTime = patcher.UpdateTime
			continue
		}
		if f == prefix+"ScopeType" {
			patchee.ScopeType = patcher.ScopeType
			continue
		}
		if f == prefix+"Menu" {
			patchee.Menu = patcher.Menu
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListSysRole executes a gorm list call
func DefaultListSysRole(ctx context.Context, db *gorm.DB) ([]*SysRole, error) {
	in := SysRole{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []SysRoleORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysRoleORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*SysRole{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type SysRoleORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysRoleORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysRoleORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]SysRoleORM) error
}
type RoleDefaultServer struct {
	DB *gorm.DB
}

// CreateSysRole ...
func (m *RoleDefaultServer) CreateSysRole(ctx context.Context, in *SysRole) (*SysRole, error) {
	db := m.DB
	if custom, ok := interface{}(in).(RoleSysRoleWithBeforeCreateSysRole); ok {
		var err error
		if db, err = custom.BeforeCreateSysRole(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateSysRole(ctx, in, db)
	if err != nil {
		return nil, err
	}
	out := res
	if custom, ok := interface{}(in).(RoleSysRoleWithAfterCreateSysRole); ok {
		var err error
		if err = custom.AfterCreateSysRole(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RoleSysRoleWithBeforeCreateSysRole called before DefaultCreateSysRoleSysRole in the default CreateSysRole handler
type RoleSysRoleWithBeforeCreateSysRole interface {
	BeforeCreateSysRole(context.Context, *gorm.DB) (*gorm.DB, error)
}

// RoleSysRoleWithAfterCreateSysRole called before DefaultCreateSysRoleSysRole in the default CreateSysRole handler
type RoleSysRoleWithAfterCreateSysRole interface {
	AfterCreateSysRole(context.Context, *SysRole, *gorm.DB) error
}

// UpdateRole ...
func (m *RoleDefaultServer) UpdateRole(ctx context.Context, in *SysRole) (*SysRole, error) {
	var err error
	var res *SysRole
	db := m.DB
	if custom, ok := interface{}(in).(RoleSysRoleWithBeforeUpdateRole); ok {
		var err error
		if db, err = custom.BeforeUpdateRole(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err = DefaultStrictUpdateSysRole(ctx, in, db)
	if err != nil {
		return nil, err
	}
	out := res
	if custom, ok := interface{}(in).(RoleSysRoleWithAfterUpdateRole); ok {
		var err error
		if err = custom.AfterUpdateRole(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RoleSysRoleWithBeforeUpdateRole called before DefaultUpdateRoleSysRole in the default UpdateRole handler
type RoleSysRoleWithBeforeUpdateRole interface {
	BeforeUpdateRole(context.Context, *gorm.DB) (*gorm.DB, error)
}

// RoleSysRoleWithAfterUpdateRole called before DefaultUpdateRoleSysRole in the default UpdateRole handler
type RoleSysRoleWithAfterUpdateRole interface {
	AfterUpdateRole(context.Context, *SysRole, *gorm.DB) error
}

// DeleteRole ...
func (m *RoleDefaultServer) DeleteRole(ctx context.Context, in *SysRole) (*emptypb.Empty, error) {
	out := &emptypb.Empty{}
	return out, nil
}

// ListRole ...
func (m *RoleDefaultServer) ListRole(ctx context.Context, in *emptypb.Empty) (*ListRoleResp, error) {
	db := m.DB
	if custom, ok := interface{}(in).(RoleSysRoleWithBeforeListRole); ok {
		var err error
		if db, err = custom.BeforeListRole(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListSysRole(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListRoleResp{Results: res}
	if custom, ok := interface{}(in).(RoleSysRoleWithAfterListRole); ok {
		var err error
		if err = custom.AfterListRole(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RoleSysRoleWithBeforeListRole called before DefaultListRoleSysRole in the default ListRole handler
type RoleSysRoleWithBeforeListRole interface {
	BeforeListRole(context.Context, *gorm.DB) (*gorm.DB, error)
}

// RoleSysRoleWithAfterListRole called before DefaultListRoleSysRole in the default ListRole handler
type RoleSysRoleWithAfterListRole interface {
	AfterListRole(context.Context, *ListRoleResp, *gorm.DB) error
}
