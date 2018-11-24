// 購物資料庫系統：TUI Inferface
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

// stdin: 接收 os.Stdin 的 Reader
var stdin = bufio.NewReader(os.Stdin)

// action(): 檢查使用者的訊息
func action(db *PurchaseDB, inputd []string) {
  switch inputd[0] {
    case "exit": // 退出程式
      os.Exit(0)
      
    case "usage": // 顯示此說明
      fmt.Println(usagetxt) // -> libpurchase_str.go
      
    case "print": // 顯示購物明細資料
      fmt.Println((*db).Print(FormatTemplate, ProductFormatTemplate))
      
    case "dump":  // 將資料庫輸出成 JSON 格式供使用
      fmt.Println("以下的內容就是未來可透過 load 指令載入的 JSON 資料：")
      fmt.Println((*db).Dump(false))
      
    case "load": // 從 JSON 格式中讀取資料庫內容
      jsond := input("請輸入之前從 dump 指令輸出的 JSON 資料：")
      (*db).Load(jsond)
      
    case "exist": // 確認某購物資料庫中是否有某產品
      var productname string
      if len(inputd) == 2 {
        productname = inputd[1]
      } else {
        productname = input("請輸入產品名稱：")
      }
      
      if (*db).IsExist(productname) {
        fmt.Printf("產品「%s」存在\n", productname)
      } else {
        fmt.Printf("產品「%s」不存在\n", productname)
      }
    
    case "comment":
      var com string
      if len(inputd) == 2 {
        com = inputd[1]
      } else {
        com = input("請輸入備註：")
      }
      
      (*db).SetComment(com)
      fmt.Printf("備註「%s」加入或編輯成功。\n", com)
    
    case "remove": // 從購物資料庫中移除某產品
      var prod string
      if len(inputd) == 2 {
        prod = inputd[1]
      } else {
        prod = input("請輸入產品名稱：")
      }
      
      (*db).RemoveProduct(prod)
      if prod == "*" {
        fmt.Println("已刪除所有產品。")
      } else {
        fmt.Printf("產品「%s」刪除成功。\n", prod)
      }

    case "add": // 增加一商品到購物資料庫中
      var prod string
      var pric float64
      var coun int
      var err error
      
      if len(inputd) == 4 {
        prod = inputd[1]
        pric, err = strconv.ParseFloat(inputd[2], 64)
        errh(err)
        coun, err = strconv.Atoi(inputd[3])
        errh(err)
      } else {
        prod = input("請輸入產品名稱：")
        pric, err = strconv.ParseFloat(input("請輸入價格："), 64)
        errh(err)
        coun, err = strconv.Atoi(input("請輸入數量："))
        errh(err)
      }
      
      (*db).AddProduct(prod, pric, coun)
      fmt.Printf("產品「%s」加入成功。\n", prod)

    default:
      fmt.Fprintf(os.Stderr, "您輸入的選項「%s」不能用！\n", inputd[0])
  }
  fmt.Println()
}

// errh(): 負責處理 error 訊息
func errh(err error) {
  if err != nil {
    panic(err)
  }
}

// input(): 負責輸入與輸出
func input(prompt string) string {
  fmt.Print(prompt)
  strdata, err := stdin.ReadString('\n')
  errh(err)
  strdata = strings.Replace(strdata, "\n", "", -1)
  return strdata
}
  
// [主程式] TUI(): 只須呼叫此程式即可進入 TUI 模式。
// 傳入 &PurchaseDB{} 應即可使用。
func TUI(db *PurchaseDB) {
  var userInput string
  var inputData []string
  fmt.Println(welcomeMsg) // -> libpurchase_str.go
  for {
    userInput = input("請輸入想做的功能：")
    inputData = strings.Split(userInput, " ")
    action(db, inputData)
  }
}
