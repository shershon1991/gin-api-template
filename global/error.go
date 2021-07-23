/**
 * @Description 全局捕获错误
 **/
package global

import "fmt"

// 捕获错误
func CatchError()  {
	err := recover()
	if err != nil {
		// 捕获错误
		fmt.Println(fmt.Sprintf("运行失败: %s",err))
		if GvaLogger != nil {
			GvaLogger.Error(fmt.Sprintf("运行失败: %s",err))
		}
	}
}
