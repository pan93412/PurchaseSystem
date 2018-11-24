// 購物資料庫: 實做介面

package main

import (
  "os"
  "fmt"
)

var usage = `購物資料庫: 實做介面 [說明]

用法：%s [功能]

[功能]
  可選擇選項：run
  
  run:  用 TUI 文字介面操作購物資料庫。
        這也是建議給一般使用者的選項。
`

func main() {
  if len(os.Args) != 2 {
    fmt.Printf(usage, os.Args[0])
  } else if os.Args[1] == "run" {
    TUI(&PurchaseDB{})
  } else {
    fmt.Printf("未知功能：%s\n", os.Args[1])
  } 
}
