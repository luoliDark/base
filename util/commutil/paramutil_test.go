package commutil

import (
	"fmt"
	"testing"
)

/**
 * @author : Yix
 * @date 2022-06-04
 * @desc:
 **/

func TestGetWhereSqlByParamMap(t *testing.T) {

}

func TestReplaceSpecialCharTheEmpty(t *testing.T) {
	teststr := `		
所属本部	
		
`
	fmt.Println(ReplaceSpecialCharTheEmpty(teststr))
}
