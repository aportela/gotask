package utils

import "database/sql"

func Int64Ptr(n sql.NullInt64) *int64 {
	if n.Valid {
		return &n.Int64
	}
	return nil
}

func StrPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func NullableInt64ToSQL(i *int64) interface{} {
	if i != nil {
		return i
	}
	return nil
}

func NullableStringToSQL(s *string) interface{} {
	if s != nil {
		return s
	}
	return nil
}
