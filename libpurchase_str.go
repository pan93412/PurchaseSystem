// 購物資料庫系統：Inferface Text
package main

/* 購物資料庫: TUI Interface Text */

// 購物資料庫用法
const usagetxt = ` ~~ 購物資料庫 [ 用法 ] ~~
add (產品名稱) (價格，必須為數字，可有小數點) (數量)
  增加一商品到購物資料庫中。
  若只輸入 add，則會用互動方式詢問。
  
remove (產品名稱)
  從購物資料庫中移除某產品。
  若產品名稱是 *，則移除所有產品。
  若只輸入 remove，則會用互動方式詢問。
  
comment (備註)
  增加或編輯備註。
  若只輸入 comment，則會用互動方式詢問。
  
exist (產品名稱)
  確認某購物資料庫中是否有某產品。
  若只輸入 exist，則會用互動方式詢問。

dump
  將資料庫輸出成 JSON 格式供使用。
  
load
  從 JSON 格式中讀取資料庫內容。
  會使用互動方式請求 JSON 資料。
  
print
  顯示購物明細資料。
  將使用 config.go 中的範本顯示明細資料。
  若想要自訂，請手動修改 config.go 中的範本。
  
usage
  顯示此說明。
  
exit
  退出程式。`

// 購物資料庫初始訊息
const welcomeMsg = `歡迎使用購物資料庫的 TUI 工具！

  -> 輸入 usage 顯示說明
  -> 輸入  exit 退出程式

使用愉快！Have fun!`
