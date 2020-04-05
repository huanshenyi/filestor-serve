package meta

import (
	"filestor-serve/db"
	"sort"
)

// ファイル要素構造
type FileMeta struct {
	FileSha1 string //ファイルの唯一のシンボル(ID)
	FileName string
	FileSize int64
	Location string //保存先
	UploadAt string //時間
}

var fileMetas map[string]FileMeta

func init()  {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta: 新規追加，更新ファイルの元データ
func UpdateFileMeta(fmeta FileMeta){
    fileMetas[fmeta.FileSha1] = fmeta
}

// UpdateFileMetaDB:ファイルのメタデータの更新をmysqlへ保存
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return db.OnFileUploadFinished(
		fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFileMetaDB: mysqlからファイルデータを取得
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tfile, err := db.GetFileMeta(fileSha1)
	if err != nil{
		return FileMeta{}, nil
	}
	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil
}

// GetFileMeta:sha1を通して、ファイルの元データを取得
func GetFileMeta(fileSha1 string)FileMeta {
	return fileMetas[fileSha1]
}

// 複数のfileMetaを並べて返す
func GetLastFileMetas(count int) []FileMeta {
	var fMetaArray []FileMeta
	for _, v := range fileMetas{
		fMetaArray = append(fMetaArray, v)
	}
	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}

//指定されたファイルを削除
func RemoveFileMeta(fileSha1 string){
	delete(fileMetas, fileSha1)
}
