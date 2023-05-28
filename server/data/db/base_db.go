package db

import (
	"errors"

	"gorm.io/gorm"
)

// AllowedModels defines allowed models to be used in the db
type AllowedModels interface {
	GetId() uint
}

// BaseDB implements CRUDRepo for the user model
type BaseDB[T AllowedModels] struct {
	db   *gorm.DB
	zero T
}

// newBaseDB returns a new BaseDB instance for the required type,
// it's not an exported function because it's not a safe one,
// since it might cause lots of db connections and that's wack!
func newBaseDB[T AllowedModels](db *gorm.DB) *BaseDB[T] {
	return &BaseDB[T]{db: db}
}

// Add creates a new record of the given object, and returns an occurring error
// the new object is a pointer, so it updates the object's id after creation
func (b *BaseDB[T]) Add(obj *T) error {
	if obj == nil {
		return errors.New("object's pointer is nil")
	}

	err := b.db.
		Model(new(T)).
		Create(obj).
		Error

	if err != nil {
		return err
	}

	return nil
}

// AddMany is same as Add but for numerous objects
func (b *BaseDB[T]) AddMany(objs []*T) error {
	if len(objs) == 0 {
		return errors.New("slice is nil or empty")
	}

	err := b.db.
		Model(new([]T)).
		Create(&objs).
		Error

	if err != nil {
		return err
	}

	return nil
}

// Exists checks the existence of the given record's id
func (b *BaseDB[T]) Exists(id uint) bool {
	if id == 0 { // better to check this before, fetching eh?
		return false
	}
	_, err := b.Get(id)
	return err == nil
}

// Get retrieves the object which has the given id
func (b *BaseDB[T]) Get(id uint) (T, error) {
	var obj T

	err := b.db.
		Model(new(T)).
		First(&obj, "id = ?", id).
		Error

	if err != nil {
		return b.zero, err
	}

	return obj, nil
}

// GetByConds is the extended version of Get,
// which uses a given search condition and retrieves every record with the given condition
func (b *BaseDB[T]) GetByConds(conds ...any) ([]T, error) {
	if !checkConds(conds...) {
		return nil, errors.New("invalid conditions")
	}

	var foundRecords []T

	err := b.db.
		Model(new(T)).
		Find(&foundRecords, conds...).
		Error

	if err != nil || len(foundRecords) == 0 {
		return nil, errors.New("no records were found")
	}

	return foundRecords, nil
}

// GetAll retrieves all the records of the given model
func (b *BaseDB[T]) GetAll() ([]T, error) {
	return b.GetByConds("id != ?", 0)
}

// Count returns the number of records of the given model
func (b *BaseDB[T]) Count() int64 {
	var count int64

	err := b.db.
		Model(new(T)).
		Count(&count).
		Error

	if err != nil {
		return 0
	}

	return count
}

// Update updates the given object/s based on the given condition
// the updated object is a pointer, so it changes the values in it as well,
// and gives it its id(in case searching condition weren't using id)
func (b *BaseDB[T]) Update(obj *T, conds ...any) error {
	if obj == nil {
		return errors.New("object's pointer is nil")
	}

	if len(conds) > 1 {
		if !checkConds(conds...) {
			return errors.New("invalid conditions")
		}
	} else {
		conds = []any{"id = ?", (*obj).GetId()}
	}

	_, err := b.GetByConds(conds...)
	if err != nil {
		return err
	}

	// if reflect.DeepEqual(*obj, foundRecords[0]) {
	// 	return errors.ErrCouldntUpdate.New("new values are the same as the old ones " + byBaseDB + ".Update")
	// }

	err = b.db.
		Model(new(T)).
		Where(conds[0], conds[1:]...).
		Updates(obj).
		Error

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the given object/s based on the given object
func (b *BaseDB[T]) Delete(conds ...any) error {
	obj, err := b.GetByConds(conds...)
	if err != nil {
		return err
	}

	err = b.db.
		Model(new(T)).
		Delete(&obj, conds...).
		Error

	if err != nil {
		return err
	}

	return nil
}

// GetDB well, ding...
func (b *BaseDB[T]) GetDB() *gorm.DB {
	return b.db
}
