package organization

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/v2/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	strings "strings"
	time "time"
)

type SysOrganizationORM struct {
	Children          []*SysOrganizationORM `gorm:"foreignKey:SysOrganizationId;references:Id"`
	CreateBy          string
	CreatedAt         *time.Time
	Depth             int32
	Hierarchy         string
	Id                int32 `gorm:"type:integer;primaryKey;autoIncrement"`
	Name              string
	ParentId          int32
	Remarks           string
	Sort              int32
	SysOrganizationId *int32
	UpdateBy          string
	UpdatedAt         *time.Time
}

// TableName overrides the default tablename generated by GORM
func (SysOrganizationORM) TableName() string {
	return "sys_organizations"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *SysOrganization) ToORM(ctx context.Context) (SysOrganizationORM, error) {
	to := SysOrganizationORM{}
	var err error
	if prehook, ok := interface{}(m).(SysOrganizationWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.ParentId = m.ParentId
	to.Hierarchy = m.Hierarchy
	to.Depth = m.Depth
	to.Remarks = m.Remarks
	to.Sort = m.Sort
	to.CreateBy = m.CreateBy
	to.UpdateBy = m.UpdateBy
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
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
	if posthook, ok := interface{}(m).(SysOrganizationWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SysOrganizationORM) ToPB(ctx context.Context) (SysOrganization, error) {
	to := SysOrganization{}
	var err error
	if prehook, ok := interface{}(m).(SysOrganizationWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.ParentId = m.ParentId
	to.Hierarchy = m.Hierarchy
	to.Depth = m.Depth
	to.Remarks = m.Remarks
	to.Sort = m.Sort
	to.CreateBy = m.CreateBy
	to.UpdateBy = m.UpdateBy
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
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
	if posthook, ok := interface{}(m).(SysOrganizationWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type SysOrganization the arg will be the target, the caller the one being converted from

// SysOrganizationBeforeToORM called before default ToORM code
type SysOrganizationWithBeforeToORM interface {
	BeforeToORM(context.Context, *SysOrganizationORM) error
}

// SysOrganizationAfterToORM called after default ToORM code
type SysOrganizationWithAfterToORM interface {
	AfterToORM(context.Context, *SysOrganizationORM) error
}

// SysOrganizationBeforeToPB called before default ToPB code
type SysOrganizationWithBeforeToPB interface {
	BeforeToPB(context.Context, *SysOrganization) error
}

// SysOrganizationAfterToPB called after default ToPB code
type SysOrganizationWithAfterToPB interface {
	AfterToPB(context.Context, *SysOrganization) error
}

// DefaultCreateSysOrganization executes a basic gorm create call
func DefaultCreateSysOrganization(ctx context.Context, in *SysOrganization, db *gorm.DB) (*SysOrganization, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeCreate_); ok {
		if ormObj, db, err = hook.BeforeCreate_(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type SysOrganizationORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB, SysOrganizationORM) (SysOrganizationORM, *gorm.DB, error)
}
type SysOrganizationORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadSysOrganization(ctx context.Context, in *SysOrganization, db *gorm.DB) (*SysOrganization, error) {
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
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := SysOrganizationORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(SysOrganizationORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type SysOrganizationORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteSysOrganization(ctx context.Context, in *SysOrganization, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeDelete_); ok {
		if ormObj, db, err = hook.BeforeDelete_(ctx, db, ormObj); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&SysOrganizationORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type SysOrganizationORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB, SysOrganizationORM) (SysOrganizationORM, *gorm.DB, error)
}
type SysOrganizationORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteSysOrganizationSet(ctx context.Context, in []*SysOrganization, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&SysOrganizationORM{})).(SysOrganizationORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&SysOrganizationORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&SysOrganizationORM{})).(SysOrganizationORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type SysOrganizationORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*SysOrganization, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*SysOrganization, *gorm.DB) error
}

// DefaultStrictUpdateSysOrganization clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateSysOrganization(ctx context.Context, in *SysOrganization, db *gorm.DB) (*SysOrganization, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateSysOrganization")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &SysOrganizationORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeStrictUpdateCleanup); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateCleanup(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	filterChildren := SysOrganizationORM{}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	filterChildren.SysOrganizationId = new(int32)
	*filterChildren.SysOrganizationId = ormObj.Id
	if err = db.Where(filterChildren).Delete(SysOrganizationORM{}).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeStrictUpdateSave); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateSave(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithAfterStrictUpdateSave); ok {
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

type SysOrganizationORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB, SysOrganizationORM) (SysOrganizationORM, *gorm.DB, error)
}
type SysOrganizationORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB, SysOrganizationORM) (SysOrganizationORM, *gorm.DB, error)
}
type SysOrganizationORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchSysOrganization executes a basic gorm update call with patch behavior
func DefaultPatchSysOrganization(ctx context.Context, in *SysOrganization, updateMask *field_mask.FieldMask, db *gorm.DB) (*SysOrganization, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj SysOrganization
	var err error
	if hook, ok := interface{}(&pbObj).(SysOrganizationWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadSysOrganization(ctx, &SysOrganization{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(SysOrganizationWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskSysOrganization(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(SysOrganizationWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateSysOrganization(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(SysOrganizationWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type SysOrganizationWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *SysOrganization, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *SysOrganization, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *SysOrganization, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *SysOrganization, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetSysOrganization executes a bulk gorm update call with patch behavior
func DefaultPatchSetSysOrganization(ctx context.Context, objects []*SysOrganization, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*SysOrganization, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*SysOrganization, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchSysOrganization(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskSysOrganization patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskSysOrganization(ctx context.Context, patchee *SysOrganization, patcher *SysOrganization, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*SysOrganization, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"ParentId" {
			patchee.ParentId = patcher.ParentId
			continue
		}
		if f == prefix+"Hierarchy" {
			patchee.Hierarchy = patcher.Hierarchy
			continue
		}
		if f == prefix+"Depth" {
			patchee.Depth = patcher.Depth
			continue
		}
		if f == prefix+"Remarks" {
			patchee.Remarks = patcher.Remarks
			continue
		}
		if f == prefix+"Sort" {
			patchee.Sort = patcher.Sort
			continue
		}
		if f == prefix+"CreateBy" {
			patchee.CreateBy = patcher.CreateBy
			continue
		}
		if f == prefix+"UpdateBy" {
			patchee.UpdateBy = patcher.UpdateBy
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
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

// DefaultListSysOrganization executes a gorm list call
func DefaultListSysOrganization(ctx context.Context, db *gorm.DB) ([]*SysOrganization, error) {
	in := SysOrganization{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []SysOrganizationORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysOrganizationORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*SysOrganization{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type SysOrganizationORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysOrganizationORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]SysOrganizationORM) error
}
type OrganizationDefaultServer struct {
	DB *gorm.DB
}

// CreateOrganization ...
func (m *OrganizationDefaultServer) CreateOrganization(ctx context.Context, in *SysOrganization) (*SysOrganization, error) {
	db := m.DB
	if custom, ok := interface{}(in).(OrganizationSysOrganizationWithBeforeCreateOrganization); ok {
		var err error
		if db, err = custom.BeforeCreateOrganization(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateSysOrganization(ctx, in, db)
	if err != nil {
		return nil, err
	}
	out := res
	if custom, ok := interface{}(in).(OrganizationSysOrganizationWithAfterCreateOrganization); ok {
		var err error
		if err = custom.AfterCreateOrganization(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// OrganizationSysOrganizationWithBeforeCreateOrganization called before DefaultCreateOrganizationSysOrganization in the default CreateOrganization handler
type OrganizationSysOrganizationWithBeforeCreateOrganization interface {
	BeforeCreateOrganization(context.Context, *gorm.DB) (*gorm.DB, error)
}

// OrganizationSysOrganizationWithAfterCreateOrganization called before DefaultCreateOrganizationSysOrganization in the default CreateOrganization handler
type OrganizationSysOrganizationWithAfterCreateOrganization interface {
	AfterCreateOrganization(context.Context, *SysOrganization, *gorm.DB) error
}

// QueryOrganization ...
func (m *OrganizationDefaultServer) QueryOrganization(ctx context.Context, in *SysOrganization) (*SysOrganization, error) {
	out := &SysOrganization{}
	return out, nil
}

// UpdateOrganization ...
func (m *OrganizationDefaultServer) UpdateOrganization(ctx context.Context, in *SysOrganization) (*SysOrganization, error) {
	var err error
	var res *SysOrganization
	db := m.DB
	if custom, ok := interface{}(in).(OrganizationSysOrganizationWithBeforeUpdateOrganization); ok {
		var err error
		if db, err = custom.BeforeUpdateOrganization(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err = DefaultStrictUpdateSysOrganization(ctx, in, db)
	if err != nil {
		return nil, err
	}
	out := res
	if custom, ok := interface{}(in).(OrganizationSysOrganizationWithAfterUpdateOrganization); ok {
		var err error
		if err = custom.AfterUpdateOrganization(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// OrganizationSysOrganizationWithBeforeUpdateOrganization called before DefaultUpdateOrganizationSysOrganization in the default UpdateOrganization handler
type OrganizationSysOrganizationWithBeforeUpdateOrganization interface {
	BeforeUpdateOrganization(context.Context, *gorm.DB) (*gorm.DB, error)
}

// OrganizationSysOrganizationWithAfterUpdateOrganization called before DefaultUpdateOrganizationSysOrganization in the default UpdateOrganization handler
type OrganizationSysOrganizationWithAfterUpdateOrganization interface {
	AfterUpdateOrganization(context.Context, *SysOrganization, *gorm.DB) error
}

// DeleteOrganization ...
func (m *OrganizationDefaultServer) DeleteOrganization(ctx context.Context, in *SysOrganization) (*SysOrganization, error) {
	out := &SysOrganization{}
	return out, nil
}

// ListOrganization ...
func (m *OrganizationDefaultServer) ListOrganization(ctx context.Context, in *SysOrganization) (*ListSysOrganization, error) {
	db := m.DB
	if custom, ok := interface{}(in).(OrganizationSysOrganizationWithBeforeListOrganization); ok {
		var err error
		if db, err = custom.BeforeListOrganization(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListSysOrganization(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListSysOrganization{Results: res}
	if custom, ok := interface{}(in).(OrganizationSysOrganizationWithAfterListOrganization); ok {
		var err error
		if err = custom.AfterListOrganization(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// OrganizationSysOrganizationWithBeforeListOrganization called before DefaultListOrganizationSysOrganization in the default ListOrganization handler
type OrganizationSysOrganizationWithBeforeListOrganization interface {
	BeforeListOrganization(context.Context, *gorm.DB) (*gorm.DB, error)
}

// OrganizationSysOrganizationWithAfterListOrganization called before DefaultListOrganizationSysOrganization in the default ListOrganization handler
type OrganizationSysOrganizationWithAfterListOrganization interface {
	AfterListOrganization(context.Context, *ListSysOrganization, *gorm.DB) error
}

// OrganizationTree ...
func (m *OrganizationDefaultServer) OrganizationTree(ctx context.Context, in *SysOrganization) (*ListSysOrganization, error) {
	out := &ListSysOrganization{}
	return out, nil
}
