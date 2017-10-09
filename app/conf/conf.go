package conf

const ()

/*
　　　　　　　　　　　　　　　　　　　　　　　　　　　　　 　 -―─- ､
　　　　　　　　　　　　　　　　　　　　　　　　　　 ,　'´ ／ヽ/レ'＾＼｀¨　￣｀ヽ
　　　　　　　　　　　　　　　　　　　　　　　 　 ／/＾∨　　　　　　　＼ ⌒＼ ＼
　　　　　　　　　　　　　　 　 　 　 　 　 　 ／／　　　　 　 　 　 　 　 ＼　　＼ ＼
　　　　　　　　　　　　　　　　　　　 ／￣7／:/　　　　　　ｌ　　　　　　　 ＼　　＼ ヽ
　　　　　　　　 　 　 　 　 　 　 　 / 　 ／　./: 　 ,':.../!:.　 |:::l:{ :..　　　　　　＼　　Ｖﾊ
　　　　　　　　　　　　　　　　　　/　／　.:::/:'::. 　!:::/ |::..　|:::l:::ヽ::..　＼　　　　ヽ　 Vj __
　　　　　　　　　　　　　　　　 　 |.:/,ｲ...::::::ｌ::ｌ::::.　|:::|　ｌ::::.. |::.l:::::::.＼ :::..＼::...　　∨::Y　 ｀＼
　　　　　　　　　　　　　　　　 　 l// :!.::/:::|::ｌ::::.　|:::l　 !::::: |::.{ヽ::{＼ ヽ､::: ＼::...　.l::::|⌒ヽ　 l
　　　　　　　　　　　　　　　　　 //:::::|:::l:::::|::l ::::.､|l:_|　 ＼_;{ ヽ ＼　＼　＼:_丶‐::.|.::|:::::::::l　j
　　　　　　 　 　 　 　 　 　 　 /ｲ ::::::|:::l:::::|::ｌ :::::::|ヽ｢二ﾆﾍ:ヽ　　　｀ﾆ二下＼:ヽ:::|.::! :::::::|　{
　　　　　　　　　　　　　　　　 { !| :::::::l:/!::::|:ﾊ::::::::V仟ｱてヽ＼ 　 　 仟アてヽ乂:::|/ ::::::: |　l
　　　　　　 　 　 　 　 　 　 　 Ⅵ :::::::{:∧::∨{＼_::ヽ∨少'_　　　　　　∨少'_/／リ:::: 　　|　|
　　　　　　　　　　　　　　　　　`ﾄ､::::::ヽ:∧:Ⅵ:::: ハ ／／／　　　　／／／イ:::::::|! :::: 　 |　l
　　　　　　　　　　　　　　　　　 l　::::::::::/::∧＼:::::小、　　　　　'　　　　　　 小 ::: |ｌ :::::.　 ∨
　　　　　　　　　　　　　　 　 　 i　::::::::/::/　 ヽ:ヽ:::|:::l＼　　　 (⌒)　　　／/l| l:::: |.| :::::　　{
　　　　　　　　　　　 　 　 　 　 l　:::::/::/　　 /|::::: :|:::|ハ{＞　､　 _ ,　イﾍＶ /' |::::.| ! :::..
　　　　　　　　　　　　　　　　　l　::::::::/:!　　 {ﾊ::::::::|::{　ヾ:{ F＝=v==＝7}:}/　 !::: ｌ人:::::.　　ヽ
　　　　　　　　　　　　　　　　 　 :::::::/:::ｌ　　 ヽﾍ:::　ﾄ ＼　rK （(db)）　 h<_　 /:::./　 |＼:::..　　＼
　　　　　　　　　　　　　　　 /　:::::::/ ｌ:::!　, -‐- ヽ:::ｌ／￣ ｀　‐-|ｌ|-‐　´ 　`/:::./. ___j_:::.＼:::..　　＼
　　　　　　　 　 　 　 　 　 /　:::::::/　j:/／　　　　 ＼{　 　 　 　 |ｌ|　　　　 /:::./ /　　　｀＼＼::. 　　＼
　　　　　　　　　　　　　　/　:::::::/　 ／ 　 　 　 　 l　|＼　 ＿　 |ｌ| ＿　／; イ /　　　　　　＼＼::. 　　＼
　　　　　　　 　 　 　 　 /　::::::/　 / 　 　 　 　 　 :l　|　　 {__rx／{＼ ｀}}￣:/ /　　 　 　 　 　 ＼＼::. 　　＼
.　　　　　　 　 　 　 　 /　::::::/ 　 {　　　　　 　 　 ∧ヽ ／ _,＞!:::}::::}＼　 / ∧　　 　 　 　 　　　}　＼::. 　　＼
　　　　　　　　　　 　 /　::::::/　　 }ヽ　　　　　　　 l /∨　 　 ／}ーく 　 ＼,∧ヽヽ 　 　 　 　 　,/　 　 ＼::. 　　＼
　　 　 　 　 　 　 　 /　::::::/　 　 ｛＼　　 　 　 　 l l　ｌ{　　　二 }::::::{　　　 　 l　l│　　　　　　/　　　　　 ＼::. 　　＼
　　　　　　　　　 　/　::::::/　　　 人　＼＿＿＿/l l　j }　　　ヽ/:::::::} 　 　 　 l　l│　　　　 ／　　　 　 　 　 ＼::. 　　＼
　　　　　 　 　 　 /　::::::/　　　 <　 ヽ＿＿_＿ ハ_{　{ !　　　 l::::::::::ｌ　　　 　 l　l人＿＿／/　 　 　 　 　 　 　 ＼::. 　　＼
　　　　　　　　　/　::::::/　 　 　 /　　　　ー─/￣}厂{j　　　 /| ::::::::|　　　　孑/　ヽー‐　ｲ　　　　　　　　　　　　 ＼::. 　　ヽ
　　 　 　 　 　 /　::::::/　 　 　 /　　　　　　∠二ﾆﾍ　 ＼ ／{│::::::::|　ﾆ二 /´l|　＼__,／│　　　　　　 　 　 　 　 　 ＼:: 　　＼
*/
var SADR = "\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\x2d\u2015\u2500\x2d\x20\uff64\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\x2c\u3000\x27\xb4\x20\uff0f\u30fd\x2f\u30ec\x27\uff3e\uff3c\uff40\xa8\u3000\uffe3\uff40\u30fd\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\uff0f\x2f\uff3e\u2228\u3000\u3000\u3000\u3000\u3000\u3000\u3000\uff3c\x20\u2312\uff3c\x20\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\uff0f\uff0f\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\uff3c\u3000\u3000\uff3c\x20\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\uff0f\uffe37\uff0f\x3a\x2f\u3000\u3000\u3000\u3000\u3000\u3000\uff4c\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\uff3c\u3000\u3000\uff3c\x20\u30fd\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\x20\u3000\x20\uff0f\u3000\x2e\x2f\x3a\x20\u3000\x20\x2c\x27\x3a\x2e\x2e\x2e\x2f\x21\x3a\x2e\u3000\x20\x7c\x3a\x3a\x3al\x3a\x7b\x20\x3a\x2e\x2e\u3000\u3000\u3000\u3000\u3000\u3000\uff3c\u3000\u3000\uff36\uff8a\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x2f\u3000\uff0f\u3000\x2e\x3a\x3a\x3a\x2f\x3a\x27\x3a\x3a\x2e\x20\u3000\x21\x3a\x3a\x3a\x2f\x20\x7c\x3a\x3a\x2e\x2e\u3000\x7c\x3a\x3a\x3al\x3a\x3a\x3a\u30fd\x3a\x3a\x2e\x2e\u3000\uff3c\u3000\u3000\u3000\u3000\u30fd\u3000\x20Vj\x20__\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\x7c\x2e\x3a\x2f\x2c\uff72\x2e\x2e\x2e\x3a\x3a\x3a\x3a\x3a\x3a\uff4c\x3a\x3a\uff4c\x3a\x3a\x3a\x3a\x2e\u3000\x7c\x3a\x3a\x3a\x7c\u3000\uff4c\x3a\x3a\x3a\x3a\x2e\x2e\x20\x7c\x3a\x3a\x2el\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2e\uff3c\x20\x3a\x3a\x3a\x2e\x2e\uff3c\x3a\x3a\x2e\x2e\x2e\u3000\u3000\u2228\x3a\x3aY\u3000\x20\uff40\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20l\x2f\x2f\x20\x3a\x21\x2e\x3a\x3a\x2f\x3a\x3a\x3a\x7c\x3a\x3a\uff4c\x3a\x3a\x3a\x3a\x2e\u3000\x7c\x3a\x3a\x3al\u3000\x20\x21\x3a\x3a\x3a\x3a\x3a\x20\x7c\x3a\x3a\x2e\x7b\u30fd\x3a\x3a\x7b\uff3c\x20\u30fd\uff64\x3a\x3a\x3a\x20\uff3c\x3a\x3a\x2e\x2e\x2e\u3000\x2el\x3a\x3a\x3a\x3a\x7c\u2312\u30fd\u3000\x20l\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\x2f\x2f\x3a\x3a\x3a\x3a\x3a\x7c\x3a\x3a\x3al\x3a\x3a\x3a\x3a\x3a\x7c\x3a\x3al\x20\x3a\x3a\x3a\x3a\x2e\uff64\x7cl\x3a_\x7c\u3000\x20\uff3c_\x3b\x7b\x20\u30fd\x20\uff3c\u3000\uff3c\u3000\uff3c\x3a_\u4e36\u2010\x3a\x3a\x2e\x7c\x2e\x3a\x3a\x7c\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3al\u3000j\x0a\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\uff72\x20\x3a\x3a\x3a\x3a\x3a\x3a\x7c\x3a\x3a\x3al\x3a\x3a\x3a\x3a\x3a\x7c\x3a\x3a\uff4c\x20\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7c\u30fd\uff62\u4e8c\uff86\uff8d\x3a\u30fd\u3000\u3000\u3000\uff40\uff86\u4e8c\u4e0b\uff3c\x3a\u30fd\x3a\x3a\x3a\x7c\x2e\x3a\x3a\x21\x20\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7c\u3000\x7b\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\x7b\x20\x21\x7c\x20\x3a\x3a\x3a\x3a\x3a\x3a\x3al\x3a\x2f\x21\x3a\x3a\x3a\x3a\x7c\x3a\uff8a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3aV\u4edf\uff71\u3066\u30fd\uff3c\x20\u3000\x20\u3000\x20\u4edf\u30a2\u3066\u30fd\u4e42\x3a\x3a\x3a\x7c\x2f\x20\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x20\x7c\u3000l\x0a\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u2165\x20\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7b\x3a\u2227\x3a\x3a\u2228\x7b\uff3c_\x3a\x3a\u30fd\u2228\u5c11\x27_\u3000\u3000\u3000\u3000\u3000\u3000\u2228\u5c11\x27_\x2f\uff0f\u30ea\x3a\x3a\x3a\x3a\x20\u3000\u3000\x7c\u3000\x7c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x60\uff84\uff64\x3a\x3a\x3a\x3a\x3a\x3a\u30fd\x3a\u2227\x3a\u2165\x3a\x3a\x3a\x3a\x20\u30cf\x20\uff0f\uff0f\uff0f\u3000\u3000\u3000\u3000\uff0f\uff0f\uff0f\u30a4\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7c\x21\x20\x3a\x3a\x3a\x3a\x20\u3000\x20\x7c\u3000l\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20l\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2f\x3a\x3a\u2227\uff3c\x3a\x3a\x3a\x3a\x3a\u5c0f\u3001\u3000\u3000\u3000\u3000\u3000\x27\u3000\u3000\u3000\u3000\u3000\u3000\x20\u5c0f\x20\x3a\x3a\x3a\x20\x7c\uff4c\x20\x3a\x3a\x3a\x3a\x3a\x2e\u3000\x20\u2228\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20i\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2f\x3a\x3a\x2f\u3000\x20\u30fd\x3a\u30fd\x3a\x3a\x3a\x7c\x3a\x3a\x3al\uff3c\u3000\u3000\u3000\x20\x28\u2312\x29\u3000\u3000\u3000\uff0f\x2fl\x7c\x20l\x3a\x3a\x3a\x3a\x20\x7c\x2e\x7c\x20\x3a\x3a\x3a\x3a\x3a\u3000\u3000\x7b\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20l\u3000\x3a\x3a\x3a\x3a\x3a\x2f\x3a\x3a\x2f\u3000\u3000\x20\x2f\x7c\x3a\x3a\x3a\x3a\x3a\x20\x3a\x7c\x3a\x3a\x3a\x7c\u30cf\x7b\uff1e\u3000\uff64\u3000\x20_\x20\x2c\u3000\u30a4\uff8d\uff36\x20\x2f\x27\x20\x7c\x3a\x3a\x3a\x3a\x2e\x7c\x20\x21\x20\x3a\x3a\x3a\x2e\x2e\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000l\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2f\x3a\x21\u3000\u3000\x20\x7b\uff8a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7c\x3a\x3a\x7b\u3000\u30fe\x3a\x7b\x20F\uff1d\x3dv\x3d\x3d\uff1d7\x7d\x3a\x7d\x2f\u3000\x20\x21\x3a\x3a\x3a\x20\uff4c\u4eba\x3a\x3a\x3a\x3a\x3a\x2e\u3000\u3000\u30fd\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2f\x3a\x3a\x3a\uff4c\u3000\u3000\x20\u30fd\uff8d\x3a\x3a\x3a\u3000\uff84\x20\uff3c\u3000rK\x20\uff08\x28db\x29\uff09\u3000\x20h\x3c_\u3000\x20\x2f\x3a\x3a\x3a\x2e\x2f\u3000\x20\x7c\uff3c\x3a\x3a\x3a\x2e\x2e\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2f\x20\uff4c\x3a\x3a\x3a\x21\u3000\x2c\x20\x2d\u2010\x2d\x20\u30fd\x3a\x3a\x3a\uff4c\uff0f\uffe3\x20\uff40\u3000\u2010\x2d\x7c\uff4c\x7c\x2d\u2010\u3000\xb4\x20\u3000\x60\x2f\x3a\x3a\x3a\x2e\x2f\x2e\x20___j_\x3a\x3a\x3a\x2e\uff3c\x3a\x3a\x3a\x2e\x2e\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000j\x3a\x2f\uff0f\u3000\u3000\u3000\u3000\x20\uff3c\x7b\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x7c\uff4c\x7c\u3000\u3000\u3000\u3000\x20\x2f\x3a\x3a\x3a\x2e\x2f\x20\x2f\u3000\u3000\u3000\uff40\uff3c\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\x20\uff0f\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20l\u3000\x7c\uff3c\u3000\x20\uff3f\u3000\x20\x7c\uff4c\x7c\x20\uff3f\u3000\uff0f\x3b\x20\u30a4\x20\x2f\u3000\u3000\u3000\u3000\u3000\u3000\uff3c\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\x20\x2f\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x3al\u3000\x7c\u3000\u3000\x20\x7b__rx\uff0f\x7b\uff3c\x20\uff40\x7d\x7d\uffe3\x3a\x2f\x20\x2f\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\uff3c\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\x2e\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\x20\u3000\x20\x7b\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u2227\u30fd\x20\uff0f\x20_\x2c\uff1e\x21\x3a\x3a\x3a\x7d\x3a\x3a\x3a\x3a\x7d\uff3c\u3000\x20\x2f\x20\u2227\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\u3000\u3000\x7d\u3000\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\u3000\x20\x7d\u30fd\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20l\x20\x2f\u2228\u3000\x20\u3000\x20\uff0f\x7d\u30fc\u304f\x20\u3000\x20\uff3c\x2c\u2227\u30fd\u30fd\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x2c\x2f\u3000\x20\u3000\x20\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\x20\u3000\x20\uff5b\uff3c\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20l\x20l\u3000\uff4c\x7b\u3000\u3000\u3000\u4e8c\x20\x7d\x3a\x3a\x3a\x3a\x3a\x3a\x7b\u3000\u3000\u3000\x20\u3000\x20l\u3000l\u2502\u3000\u3000\u3000\u3000\u3000\u3000\x2f\u3000\u3000\u3000\u3000\u3000\x20\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\u3000\u3000\x20\u4eba\u3000\uff3c\uff3f\uff3f\uff3f\x2fl\x20l\u3000j\x20\x7d\u3000\u3000\u3000\u30fd\x2f\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7d\x20\u3000\x20\u3000\x20\u3000\x20l\u3000l\u2502\u3000\u3000\u3000\u3000\x20\uff0f\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\u3000\u3000\x20\x3c\u3000\x20\u30fd\uff3f\uff3f_\uff3f\x20\u30cf_\x7b\u3000\x7b\x20\x21\u3000\u3000\u3000\x20l\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\uff4c\u3000\u3000\u3000\x20\u3000\x20l\u3000l\u4eba\uff3f\uff3f\uff0f\x2f\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\uff3c\x3a\x3a\x2e\x20\u3000\u3000\uff3c\x0a\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\u3000\u3000\u3000\u30fc\u2500\x2f\uffe3\x7d\u5382\x7bj\u3000\u3000\u3000\x20\x2f\x7c\x20\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7c\u3000\u3000\u3000\u3000\u5b51\x2f\u3000\u30fd\u30fc\u2010\u3000\uff72\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\x20\uff3c\x3a\x3a\x2e\x20\u3000\u3000\u30fd\x0a\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\x3a\x3a\x3a\x3a\x3a\x3a\x2f\u3000\x20\u3000\x20\u3000\x20\x2f\u3000\u3000\u3000\u3000\u3000\u3000\u2220\u4e8c\uff86\uff8d\u3000\x20\uff3c\x20\uff0f\x7b\u2502\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x3a\x7c\u3000\uff86\u4e8c\x20\x2f\xb4l\x7c\u3000\uff3c__\x2c\uff0f\u2502\u3000\u3000\u3000\u3000\u3000\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\u3000\x20\uff3c\x3a\x3a\x20\u3000\u3000\uff3c\x0a"

type DataPacket struct {
	Name    string // パケット名
	Idx     float64
	Signals map[string]float64
}