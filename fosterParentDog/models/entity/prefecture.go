package entity

//都道府県マスターテーブル用
//47都道府県一覧を保存しておくマスターテーブルである。
type Prefecture struct {
	Id   uint   `gorm:"primary_key							"`                //都道府県ID
	name string `gorm:"type:varchar(10);	unique	not null"` //都道府県名
}
