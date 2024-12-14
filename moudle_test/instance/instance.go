package main
import "fmt"
type Reader interface{
	read()
}
type Writer interface{
  write()
}
// 读写器
type ReadWriter struct{
}
func (r *ReadWriter) read()  {
	fmt.Println("read")
}
func (r *ReadWriter) write()  {
	fmt.Println("write")
}
func main() {
  rw := &ReadWriter{}

  var w Writer
  // 向下转型
  w = rw
  w.write()
  var r Reader
  // 类型断言
  r = w.(Reader)
  r.read()
}
