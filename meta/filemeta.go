package meta

import "sort"

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
