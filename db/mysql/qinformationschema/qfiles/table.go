// Code generated by Lingo for table information_schema.FILES - DO NOT EDIT

package qfiles

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QFiles {
	return newQFiles(alias)
}

func New() QFiles {
	return newQFiles("")
}

func newQFiles(alias string) QFiles {
	q := QFiles{_alias: alias}
	q.fileId = path.NewInt64Path(q, "FILE_ID")
	q.fileName = path.NewStringPath(q, "FILE_NAME")
	q.fileType = path.NewStringPath(q, "FILE_TYPE")
	q.tablespaceName = path.NewStringPath(q, "TABLESPACE_NAME")
	q.tableCatalog = path.NewStringPath(q, "TABLE_CATALOG")
	q.tableSchema = NewUnknownPathType(q, "TABLE_SCHEMA")
	q.tableName = NewUnknownPathType(q, "TABLE_NAME")
	q.logfileGroupName = path.NewStringPath(q, "LOGFILE_GROUP_NAME")
	q.logfileGroupNumber = path.NewInt64Path(q, "LOGFILE_GROUP_NUMBER")
	q.engine = path.NewStringPath(q, "ENGINE")
	q.fulltextKeys = NewUnknownPathType(q, "FULLTEXT_KEYS")
	q.deletedRows = NewUnknownPathType(q, "DELETED_ROWS")
	q.updateCount = NewUnknownPathType(q, "UPDATE_COUNT")
	q.freeExtents = path.NewInt64Path(q, "FREE_EXTENTS")
	q.totalExtents = path.NewInt64Path(q, "TOTAL_EXTENTS")
	q.extentSize = path.NewInt64Path(q, "EXTENT_SIZE")
	q.initialSize = path.NewInt64Path(q, "INITIAL_SIZE")
	q.maximumSize = path.NewInt64Path(q, "MAXIMUM_SIZE")
	q.autoextendSize = path.NewInt64Path(q, "AUTOEXTEND_SIZE")
	q.creationTime = NewUnknownPathType(q, "CREATION_TIME")
	q.lastUpdateTime = NewUnknownPathType(q, "LAST_UPDATE_TIME")
	q.lastAccessTime = NewUnknownPathType(q, "LAST_ACCESS_TIME")
	q.recoverTime = NewUnknownPathType(q, "RECOVER_TIME")
	q.transactionCounter = NewUnknownPathType(q, "TRANSACTION_COUNTER")
	q.version = path.NewInt64Path(q, "VERSION")
	q.rowFormat = path.NewStringPath(q, "ROW_FORMAT")
	q.tableRows = NewUnknownPathType(q, "TABLE_ROWS")
	q.avgRowLength = NewUnknownPathType(q, "AVG_ROW_LENGTH")
	q.dataLength = NewUnknownPathType(q, "DATA_LENGTH")
	q.maxDataLength = NewUnknownPathType(q, "MAX_DATA_LENGTH")
	q.indexLength = NewUnknownPathType(q, "INDEX_LENGTH")
	q.dataFree = path.NewInt64Path(q, "DATA_FREE")
	q.createTime = NewUnknownPathType(q, "CREATE_TIME")
	q.updateTime = NewUnknownPathType(q, "UPDATE_TIME")
	q.checkTime = NewUnknownPathType(q, "CHECK_TIME")
	q.checksum = NewUnknownPathType(q, "CHECKSUM")
	q.status = path.NewStringPath(q, "STATUS")
	q.extra = path.NewStringPath(q, "EXTRA")
	return q
}

type QFiles struct {
	_alias             string
	fileId             path.Int64Path
	fileName           path.StringPath
	fileType           path.StringPath
	tablespaceName     path.StringPath
	tableCatalog       path.StringPath
	tableSchema        UnknownPathType
	tableName          UnknownPathType
	logfileGroupName   path.StringPath
	logfileGroupNumber path.Int64Path
	engine             path.StringPath
	fulltextKeys       UnknownPathType
	deletedRows        UnknownPathType
	updateCount        UnknownPathType
	freeExtents        path.Int64Path
	totalExtents       path.Int64Path
	extentSize         path.Int64Path
	initialSize        path.Int64Path
	maximumSize        path.Int64Path
	autoextendSize     path.Int64Path
	creationTime       UnknownPathType
	lastUpdateTime     UnknownPathType
	lastAccessTime     UnknownPathType
	recoverTime        UnknownPathType
	transactionCounter UnknownPathType
	version            path.Int64Path
	rowFormat          path.StringPath
	tableRows          UnknownPathType
	avgRowLength       UnknownPathType
	dataLength         UnknownPathType
	maxDataLength      UnknownPathType
	indexLength        UnknownPathType
	dataFree           path.Int64Path
	createTime         UnknownPathType
	updateTime         UnknownPathType
	checkTime          UnknownPathType
	checksum           UnknownPathType
	status             path.StringPath
	extra              path.StringPath
}

// core.Table Functions

func (q QFiles) GetColumns() []core.Column {
	return []core.Column{
		q.fileId,
		q.fileName,
		q.fileType,
		q.tablespaceName,
		q.tableCatalog,
		q.tableSchema,
		q.tableName,
		q.logfileGroupName,
		q.logfileGroupNumber,
		q.engine,
		q.fulltextKeys,
		q.deletedRows,
		q.updateCount,
		q.freeExtents,
		q.totalExtents,
		q.extentSize,
		q.initialSize,
		q.maximumSize,
		q.autoextendSize,
		q.creationTime,
		q.lastUpdateTime,
		q.lastAccessTime,
		q.recoverTime,
		q.transactionCounter,
		q.version,
		q.rowFormat,
		q.tableRows,
		q.avgRowLength,
		q.dataLength,
		q.maxDataLength,
		q.indexLength,
		q.dataFree,
		q.createTime,
		q.updateTime,
		q.checkTime,
		q.checksum,
		q.status,
		q.extra,
	}
}

func (q QFiles) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QFiles) GetAlias() string {
	return q._alias
}

func (q QFiles) GetName() string {
	return "FILES"
}

func (q QFiles) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QFiles) FileId() path.Int64Path {
	return q.fileId
}

func (q QFiles) FileName() path.StringPath {
	return q.fileName
}

func (q QFiles) FileType() path.StringPath {
	return q.fileType
}

func (q QFiles) TablespaceName() path.StringPath {
	return q.tablespaceName
}

func (q QFiles) TableCatalog() path.StringPath {
	return q.tableCatalog
}

func (q QFiles) TableSchema() UnknownPathType {
	return q.tableSchema
}

func (q QFiles) TableName() UnknownPathType {
	return q.tableName
}

func (q QFiles) LogfileGroupName() path.StringPath {
	return q.logfileGroupName
}

func (q QFiles) LogfileGroupNumber() path.Int64Path {
	return q.logfileGroupNumber
}

func (q QFiles) Engine() path.StringPath {
	return q.engine
}

func (q QFiles) FulltextKeys() UnknownPathType {
	return q.fulltextKeys
}

func (q QFiles) DeletedRows() UnknownPathType {
	return q.deletedRows
}

func (q QFiles) UpdateCount() UnknownPathType {
	return q.updateCount
}

func (q QFiles) FreeExtents() path.Int64Path {
	return q.freeExtents
}

func (q QFiles) TotalExtents() path.Int64Path {
	return q.totalExtents
}

func (q QFiles) ExtentSize() path.Int64Path {
	return q.extentSize
}

func (q QFiles) InitialSize() path.Int64Path {
	return q.initialSize
}

func (q QFiles) MaximumSize() path.Int64Path {
	return q.maximumSize
}

func (q QFiles) AutoextendSize() path.Int64Path {
	return q.autoextendSize
}

func (q QFiles) CreationTime() UnknownPathType {
	return q.creationTime
}

func (q QFiles) LastUpdateTime() UnknownPathType {
	return q.lastUpdateTime
}

func (q QFiles) LastAccessTime() UnknownPathType {
	return q.lastAccessTime
}

func (q QFiles) RecoverTime() UnknownPathType {
	return q.recoverTime
}

func (q QFiles) TransactionCounter() UnknownPathType {
	return q.transactionCounter
}

func (q QFiles) Version() path.Int64Path {
	return q.version
}

func (q QFiles) RowFormat() path.StringPath {
	return q.rowFormat
}

func (q QFiles) TableRows() UnknownPathType {
	return q.tableRows
}

func (q QFiles) AvgRowLength() UnknownPathType {
	return q.avgRowLength
}

func (q QFiles) DataLength() UnknownPathType {
	return q.dataLength
}

func (q QFiles) MaxDataLength() UnknownPathType {
	return q.maxDataLength
}

func (q QFiles) IndexLength() UnknownPathType {
	return q.indexLength
}

func (q QFiles) DataFree() path.Int64Path {
	return q.dataFree
}

func (q QFiles) CreateTime() UnknownPathType {
	return q.createTime
}

func (q QFiles) UpdateTime() UnknownPathType {
	return q.updateTime
}

func (q QFiles) CheckTime() UnknownPathType {
	return q.checkTime
}

func (q QFiles) Checksum() UnknownPathType {
	return q.checksum
}

func (q QFiles) Status() path.StringPath {
	return q.status
}

func (q QFiles) Extra() path.StringPath {
	return q.extra
}
