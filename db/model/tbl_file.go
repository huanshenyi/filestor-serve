package model


//CREATE TABLE IF not exists `tbl_file` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`file_sha1` char(40) NOT NULL DEFAULT '' COMMENT '文件hash',
//`file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
//`file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
//`file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
//`create_at` datetime default NOW() COMMENT '创建日期',
//`update_at` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
//`status` int(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等状态)',
//`ext1` int(11) DEFAULT '0' COMMENT '备用字段1',
//`ext2` text COMMENT '备用字段2',
//PRIMARY KEY (`id`),
//UNIQUE KEY `idx_file_hash` (`file_sha1`),
//KEY `idx_status` (`status`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;

type TblFile struct {
	Id int64 `xorm: "pk autoincr int(64)" form:"id" json:"id"`
	FileSha1 string `json:"file_sha1" xorm: "VARCHAR(40) notnull default'' comment('文件hash')"`
	FileName string `json:"file_name" xorm: "VARCHAR(256) notnull default '' comment('文件名')"`
	FileSize int64 `json:"file_size" xorm: "BIGINT(20) default '0' comment'文件大小'"`
	FileAddr string `json:"file_addr" xorm: "VARCHAR(1024) notnull default '' comment '文件存储位置'"`
	CreateAt int `json:"create_at" xorm: "DATETIME created comment '创建日期'"`
	UpdateAt int `json:"update_at" xorm: "DATETIME created "`
}
