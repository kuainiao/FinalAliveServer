package ffDef

// DBQueryCallback 数据库操作结果回调
//	result: 为空时, 表明sql执行失败
type DBQueryCallback func(result IDBQueryResult)

// IDBQueryRequest 数据库操作请求
type IDBQueryRequest interface {
	// IsValid 请求是否还有效
	IsValid() bool

	// Cancel 请求者主动取消，执行此操作后，请求者不能再持有此实例
	// 由于数据库操作是异步进行的，一旦在执行Cancel操作时，数据库操作正在锁定输入参数状态，将导致此操作被阻塞，直到数据库操作结束
	Cancel()

	// Query 请求者请求执行数据库操作
	Query()
}

// IDBQueryResult 数据库操作结果
type IDBQueryResult interface {
	// //----------------------------------------------------------------------------------------
	// // these method is public

	// SQL return the sql string
	SQL() string

	// SQLResult return the sql excute result.
	// 仅表明sql语句本身的执行结果, 不代表达成了逻辑上的预期.
	SQLResult() error

	// // Close used for Query, and is ignore in Exec Query,
	// Close() error
	// //----------------------------------------------------------------------------------------

	//----------------------------------------------------------------------------------------
	// these method is wrap of sql.Result

	// LastInsertId returns the integer generated by the database
	// in response to a command. Typically this will be from an
	// "auto increment" column when inserting a new row. Not all
	// databases support this feature, and the syntax of such
	// statements varies.
	LastInsertId() (int64, error)

	// RowsAffected returns the number of rows affected by an
	// update, insert, or delete. Not every database or database
	// driver may support this.
	// 注意: 如update操作, 如果库内的列的值与要更新到的值一致, 则返回的影响到的行数为0
	RowsAffected() (int64, error)
	//----------------------------------------------------------------------------------------

	//----------------------------------------------------------------------------------------
	// these method is wrap of sql.Rows

	// Next prepares the next result row for reading with the Scan method. It
	// returns true on success, or false if there is no next result row or an error
	// happened while preparing it. Err should be consulted to distinguish between
	// the two cases.
	//
	// Every call to Scan, even the first one, must be preceded by a call to Next.
	Next() bool

	// Err returns the error, if any, that was encountered during iteration.
	// Err may be called after an explicit or implicit Close.
	Err() error

	// Columns returns the column names.
	// Columns returns an error if the rows are closed, or if the rows
	// are from QueryRow and there was a deferred error.
	Columns() ([]string, error)

	// Scan copies the columns in the current row into the values pointed
	// at by dest. The number of values in dest must be the same as the
	// number of columns in Rows.
	//
	// Scan converts columns read from the database into the following
	// common Go types and special types provided by the sql package:
	//
	//    *string
	//    *[]byte
	//    *int, *int8, *int16, *int32, *int64
	//    *uint, *uint8, *uint16, *uint32, *uint64
	//    *bool
	//    *float32, *float64
	//    *interface{}
	//    *RawBytes
	//    any type implementing Scanner (see Scanner docs)
	//
	// In the most simple case, if the type of the value from the source
	// column is an integer, bool or string type T and dest is of type *T,
	// Scan simply assigns the value through the pointer.
	//
	// Scan also converts between string and numeric types, as long as no
	// information would be lost. While Scan stringifies all numbers
	// scanned from numeric database columns into *string, scans into
	// numeric types are checked for overflow. For example, a float64 with
	// value 300 or a string with value "300" can scan into a uint16, but
	// not into a uint8, though float64(255) or "255" can scan into a
	// uint8. One exception is that scans of some float64 numbers to
	// strings may lose information when stringifying. In general, scan
	// floating point columns into *float64.
	//
	// If a dest argument has type *[]byte, Scan saves in that argument a
	// copy of the corresponding data. The copy is owned by the caller and
	// can be modified and held indefinitely. The copy can be avoided by
	// using an argument of type *RawBytes instead; see the documentation
	// for RawBytes for restrictions on its use.
	//
	// If an argument has type *interface{}, Scan copies the value
	// provided by the underlying driver without conversion. When scanning
	// from a source value of type []byte to *interface{}, a copy of the
	// slice is made and the caller owns the result.
	//
	// Source values of type time.Time may be scanned into values of type
	// *time.Time, *interface{}, *string, or *[]byte. When converting to
	// the latter two, time.Format3339Nano is used.
	//
	// Source values of type bool may be scanned into types *bool,
	// *interface{}, *string, *[]byte, or *RawBytes.
	//
	// For scanning into *bool, the source may be true, false, 1, 0, or
	// string inputs parseable by strconv.ParseBool.
	Scan(dest ...interface{}) error

	// Close closes the Rows, preventing further enumeration. If Next returns
	// false, the Rows are closed automatically and it will suffice to check the
	// result of Err. Close is idempotent and does not affect the result of Err.
	// Close() error
	//----------------------------------------------------------------------------------------
}
