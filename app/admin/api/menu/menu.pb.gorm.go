package menu

import (
	context "context"
	fmt "fmt"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	gorm "gorm.io/gorm"
)

type SysMenuORM struct {
	Children   []*SysMenuORM `gorm:"foreignKey:SysMenuId;references:Id"`
	CreateTime string
	Hidden     int32
	I18NTitle  string
	Icon       string
	Id         int32 `gorm:"type:integer;primaryKey;autoIncrement"`
	KeepAlive  int32
	ParentId   int32
	Path       string
	Permission string
	Remarks    string
	Sort       int32
	SysMenuId  *int32
	TargetType int32
	Title      string
	Type       int32
	UpdateTime string
	Uri        string
}

// TableName overrides the default tablename generated by GORM
func (SysMenuORM) TableName() string {
	return "sys_menus"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *SysMenu) ToORM(ctx context.Context) (SysMenuORM, error) {
	to := SysMenuORM{}
	var err error
	if prehook, ok := interface{}(m).(SysMenuWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ParentId = m.ParentId
	to.Title = m.Title
	to.I18NTitle = m.I18NTitle
	to.Icon = m.Icon
	to.Permission = m.Permission
	to.Path = m.Path
	to.TargetType = m.TargetType
	to.Uri = m.Uri
	to.Sort = m.Sort
	to.KeepAlive = m.KeepAlive
	to.Hidden = m.Hidden
	to.Type = m.Type
	to.Remarks = m.Remarks
	to.CreateTime = m.CreateTime
	to.UpdateTime = m.UpdateTime
	for _, v := range m.Children {
		if v != nil {
			if tempChildren, cErr := v.ToORM(ctx); cErr == nil {
				to.Children = append(to.Children, &tempChildren)
			} else {
				return to, cErr
			}
		} else {
			to.Children = append(to.Children, nil)
		}
	}
	if posthook, ok := interface{}(m).(SysMenuWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SysMenuORM) ToPB(ctx context.Context) (SysMenu, error) {
	to := SysMenu{}
	var err error
	if prehook, ok := interface{}(m).(SysMenuWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ParentId = m.ParentId
	to.Title = m.Title
	to.I18NTitle = m.I18NTitle
	to.Icon = m.Icon
	to.Permission = m.Permission
	to.Path = m.Path
	to.TargetType = m.TargetType
	to.Uri = m.Uri
	to.Sort = m.Sort
	to.KeepAlive = m.KeepAlive
	to.Hidden = m.Hidden
	to.Type = m.Type
	to.Remarks = m.Remarks
	to.CreateTime = m.CreateTime
	to.UpdateTime = m.UpdateTime
	for _, v := range m.Children {
		if v != nil {
			if tempChildren, cErr := v.ToPB(ctx); cErr == nil {
				to.Children = append(to.Children, &tempChildren)
			} else {
				return to, cErr
			}
		} else {
			to.Children = append(to.Children, nil)
		}
	}
	if posthook, ok := interface{}(m).(SysMenuWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type SysMenu the arg will be the target, the caller the one being converted from

// SysMenuBeforeToORM called before default ToORM code
type SysMenuWithBeforeToORM interface {
	BeforeToORM(context.Context, *SysMenuORM) error
}

// SysMenuAfterToORM called after default ToORM code
type SysMenuWithAfterToORM interface {
	AfterToORM(context.Context, *SysMenuORM) error
}

// SysMenuBeforeToPB called before default ToPB code
type SysMenuWithBeforeToPB interface {
	BeforeToPB(context.Context, *SysMenu) error
}

// SysMenuAfterToPB called after default ToPB code
type SysMenuWithAfterToPB interface {
	AfterToPB(context.Context, *SysMenu) error
}

// DefaultCreateSysMenu executes a basic gorm create call
func DefaultCreateSysMenu(ctx context.Context, in *SysMenu, db *gorm.DB) (*SysMenu, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeCreate_); ok {
		if ormObj, db, err = hook.BeforeCreate_(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type SysMenuORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB, SysMenuORM) (SysMenuORM, *gorm.DB, error)
}
type SysMenuORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadSysMenu(ctx context.Context, in *SysMenu, db *gorm.DB) (*SysMenu, error) {
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
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := SysMenuORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(SysMenuORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type SysMenuORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysMenuORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysMenuORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteSysMenu(ctx context.Context, in *SysMenu, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeDelete_); ok {
		if ormObj, db, err = hook.BeforeDelete_(ctx, db, ormObj); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&SysMenuORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type SysMenuORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB, SysMenuORM) (SysMenuORM, *gorm.DB, error)
}
type SysMenuORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteSysMenuSet(ctx context.Context, in []*SysMenu, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&SysMenuORM{})).(SysMenuORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&SysMenuORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&SysMenuORM{})).(SysMenuORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type SysMenuORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*SysMenu, *gorm.DB) (*gorm.DB, error)
}
type SysMenuORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*SysMenu, *gorm.DB) error
}

// DefaultStrictUpdateSysMenu clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateSysMenu(ctx context.Context, in *SysMenu, db *gorm.DB) (*SysMenu, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateSysMenu")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &SysMenuORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeStrictUpdateCleanup); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateCleanup(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	filterChildren := SysMenuORM{}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	filterChildren.SysMenuId = new(int32)
	*filterChildren.SysMenuId = ormObj.Id
	if err = db.Where(filterChildren).Delete(SysMenuORM{}).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeStrictUpdateSave); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateSave(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithAfterStrictUpdateSave); ok {
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

type SysMenuORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB, SysMenuORM) (SysMenuORM, *gorm.DB, error)
}
type SysMenuORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB, SysMenuORM) (SysMenuORM, *gorm.DB, error)
}
type SysMenuORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchSysMenu executes a basic gorm update call with patch behavior
func DefaultPatchSysMenu(ctx context.Context, in *SysMenu, updateMask *field_mask.FieldMask, db *gorm.DB) (*SysMenu, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj SysMenu
	var err error
	if hook, ok := interface{}(&pbObj).(SysMenuWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadSysMenu(ctx, &SysMenu{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(SysMenuWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskSysMenu(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(SysMenuWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateSysMenu(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(SysMenuWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type SysMenuWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *SysMenu, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysMenuWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *SysMenu, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysMenuWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *SysMenu, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysMenuWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *SysMenu, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetSysMenu executes a bulk gorm update call with patch behavior
func DefaultPatchSetSysMenu(ctx context.Context, objects []*SysMenu, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*SysMenu, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*SysMenu, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchSysMenu(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskSysMenu patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskSysMenu(ctx context.Context, patchee *SysMenu, patcher *SysMenu, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*SysMenu, error) {
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
		if f == prefix+"ParentId" {
			patchee.ParentId = patcher.ParentId
			continue
		}
		if f == prefix+"Title" {
			patchee.Title = patcher.Title
			continue
		}
		if f == prefix+"I18NTitle" {
			patchee.I18NTitle = patcher.I18NTitle
			continue
		}
		if f == prefix+"Icon" {
			patchee.Icon = patcher.Icon
			continue
		}
		if f == prefix+"Permission" {
			patchee.Permission = patcher.Permission
			continue
		}
		if f == prefix+"Path" {
			patchee.Path = patcher.Path
			continue
		}
		if f == prefix+"TargetType" {
			patchee.TargetType = patcher.TargetType
			continue
		}
		if f == prefix+"Uri" {
			patchee.Uri = patcher.Uri
			continue
		}
		if f == prefix+"Sort" {
			patchee.Sort = patcher.Sort
			continue
		}
		if f == prefix+"KeepAlive" {
			patchee.KeepAlive = patcher.KeepAlive
			continue
		}
		if f == prefix+"Hidden" {
			patchee.Hidden = patcher.Hidden
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Remarks" {
			patchee.Remarks = patcher.Remarks
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
		if f == prefix+"Children" {
			patchee.Children = patcher.Children
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListSysMenu executes a gorm list call
func DefaultListSysMenu(ctx context.Context, db *gorm.DB) ([]*SysMenu, error) {
	in := SysMenu{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []SysMenuORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysMenuORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*SysMenu{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type SysMenuORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysMenuORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysMenuORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]SysMenuORM) error
}
type MenuDefaultServer struct {
	DB *gorm.DB
}

// ListSysMenu ...
func (m *MenuDefaultServer) ListSysMenu(ctx context.Context, in *GetMenuReq) (*ListSysMenuResp, error) {
	db := m.DB
	if custom, ok := interface{}(in).(MenuSysMenuWithBeforeListSysMenu); ok {
		var err error
		if db, err = custom.BeforeListSysMenu(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListSysMenu(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListSysMenuResp{Results: res}
	if custom, ok := interface{}(in).(MenuSysMenuWithAfterListSysMenu); ok {
		var err error
		if err = custom.AfterListSysMenu(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// MenuSysMenuWithBeforeListSysMenu called before DefaultListSysMenuSysMenu in the default ListSysMenu handler
type MenuSysMenuWithBeforeListSysMenu interface {
	BeforeListSysMenu(context.Context, *gorm.DB) (*gorm.DB, error)
}

// MenuSysMenuWithAfterListSysMenu called before DefaultListSysMenuSysMenu in the default ListSysMenu handler
type MenuSysMenuWithAfterListSysMenu interface {
	AfterListSysMenu(context.Context, *ListSysMenuResp, *gorm.DB) error
}

// CreateSysMenu ...
func (m *MenuDefaultServer) CreateSysMenu(ctx context.Context, in *SysMenu) (*SysMenu, error) {
	db := m.DB
	if custom, ok := interface{}(in).(MenuSysMenuWithBeforeCreateSysMenu); ok {
		var err error
		if db, err = custom.BeforeCreateSysMenu(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateSysMenu(ctx, in, db)
	if err != nil {
		return nil, err
	}
	out := res
	if custom, ok := interface{}(in).(MenuSysMenuWithAfterCreateSysMenu); ok {
		var err error
		if err = custom.AfterCreateSysMenu(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// MenuSysMenuWithBeforeCreateSysMenu called before DefaultCreateSysMenuSysMenu in the default CreateSysMenu handler
type MenuSysMenuWithBeforeCreateSysMenu interface {
	BeforeCreateSysMenu(context.Context, *gorm.DB) (*gorm.DB, error)
}

// MenuSysMenuWithAfterCreateSysMenu called before DefaultCreateSysMenuSysMenu in the default CreateSysMenu handler
type MenuSysMenuWithAfterCreateSysMenu interface {
	AfterCreateSysMenu(context.Context, *SysMenu, *gorm.DB) error
}
