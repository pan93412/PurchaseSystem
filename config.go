// 購物資料庫系統：設定
package main

// Print() 的範本：FormatTemplate
// 第一個 %s：此購物明細的唯一識別碼
// 第二個 %s：購買的時間，格式為 年-月-日 上/下午 12制時:分:秒
// 第三個 %s：參考 ProductFormatTemplate，購買產品明細
// 第四個 %.2f：前面的 `.2` 是必須的。顯示總計價格。
// 第五個 %s：顯示備註
const FormatTemplate = `唯一代碼：%s
購買時間：%s
=== 購買之產品 ===
產品		價格		數量
%s
總價：%.2f$
備註：%s
===============
`

// Print() 的範本：ProductFormatTemplate
// 第一個 %s：產品名稱
// 第二個 %.2f：產品價格
// 第三個 %d：產品數量
// 可參考 FormatTemplate。
const ProductFormatTemplate = "%s       %.2f$    %d個\n"

// 顯示用時間的格式
// Go 時間格式參考：
// 2006: 年
// 01:   月
// 02:   日
// 03:   12 小時制時
// 15:   24 小時制時
// 04:   月
// 05:   日
// PM:   12 小時制才比較有用；顯示 AM/PM
const TimeFormat = "2006-01-02 03:04:05 PM"
