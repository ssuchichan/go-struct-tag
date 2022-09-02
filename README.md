# go-struct-tag
## 结构体标签
```go
type MyStruct struct {
	FieldName FieldType `key1:"value1;value2" key2:"value"`
}
```
## 作用
* 文档说明
* 开发过程中的特殊标记
  * ORM（Object-relational mapping）对象关系映射
  * 数据验证：form表单
  * 序列化与反序列化：结构体与JSON字符串互转

## 好处
* 把常用的字段验证抽取出来，代码复用度高
* 简化客户代码编写难度，是用户更关注其他业务的处理，一定程度上解耦合

