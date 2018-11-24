// 購物資料庫系統：CLI Lib
package main

import (
  "time"
  "fmt"
  "math/rand"
  "encoding/json"
)

// 購物資料庫
type PurchaseDB struct {
  ProductList []Product    // 購買產品列表
  UID         string       // 使用者購買唯一碼
  Time        time.Time    // 購買時間
  Comment     string       // 備註
}

// 商品格式
type Product struct {
  Name string    // 產品名稱
  Price float64  // 產品價格
  Count int      // 購買數
}

// IsExist(): 檢查商品是否在購物資料庫中已經存在
func (db PurchaseDB) IsExist(name string) bool {
  for _, data := range db.ProductList {
    if data.Name == name {
      return true
    }
  }
  return false
}

// SetComment(): 設定備註
func (db *PurchaseDB) SetComment(comment string) {
  (*db).Comment = comment
  (*db).UpdDB()
}

// AddProduct(): 增加商品到購物資料庫
func (db *PurchaseDB) AddProduct(name string, price float64, count int) {
  var prod Product
  if (*db).IsExist(name) {
    return
  }
  prod.Name = name
  prod.Price = price
  prod.Count = count
  (*db).ProductList = append((*db).ProductList, prod)
  (*db).UpdDB()
}

// RemoveProduct(): 從購物資料庫移除一個商品
// 若 name 為 *，則刪除所有產品。
func (db *PurchaseDB) RemoveProduct(name string) {
  var tmpdb []Product
  for _, data := range db.ProductList {
    if data.Name == name || name == "*" {
      continue
    }
    tmpdb = append(tmpdb, data)
  }
  (*db).ProductList = tmpdb
  (*db).UpdDB()
}

// Dump(): 輸出目前的購物資料庫成 JSON 格式
// indent (bool): 是否將輸出內容縮排？
// 縮排後的內容將較不便於載入，除非此 JSON 資料
// 只是要用來觀看，否則建議不使用縮排功能。
func (db PurchaseDB) Dump(indent bool) string {
  var jsonData []byte
  var err error
  
  if indent {
    jsonData, err = json.MarshalIndent(db, "", "  ")
  } else {
    jsonData, err = json.Marshal(db)
  }
  if err != nil {
    panic(err)
  }
  return string(jsonData)
}

// Load(): 從 Dump() 產生的 JSON 資料中載入進購物資料庫
func (db *PurchaseDB) Load(data string) {
  err := json.Unmarshal([]byte(data), db)
  if err != nil {
    panic(err)
  }
}

// UpdDB(): 更新購物資料庫中的 UID 和 Time
// 您不需要手動呼叫。
func (db *PurchaseDB) UpdDB() {
  currentTime := time.Now()
  rand.Seed(currentTime.Unix())
  (*db).UID = currentTime.Format("20060102150405") + fmt.Sprintf("%02d", rand.Intn(23) + 1)
  (*db).Time = currentTime
}

// Print(): 顯示目前購物資料庫的購物明細。
// Format 請參閱 config.go 中的 FormatTemplate 常數說明部份。
// ProductFormat 請參閱 config.go 中的 ProductFormatTemplate 常數說明部份。
// 若僅想要使用預設格式，可直接使用以上兩個範本。
func (db PurchaseDB) Print(Format string, ProductFormat string) string {
  var ProductDetails string
  var FinalList      string
  var TotalPrice     float64
  var Time           string = db.Time.Format(TimeFormat) // TimeFormat 可在 config.go 修改
  
  // ProductDetails
  for _, data := range db.ProductList {
    ProductDetails += fmt.Sprintf(ProductFormat, data.Name, data.Price, data.Count)
    TotalPrice += data.Price * float64(data.Count)
  }
  
  // FinalList
  FinalList = fmt.Sprintf(Format, db.UID, Time, ProductDetails, TotalPrice, db.Comment)
  return FinalList
}
