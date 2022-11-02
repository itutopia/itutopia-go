package test

import "testing"

/**
+
-
*
/
==
+=
-=
*=
/=
>
<
>=
>=
& &&
| ||
!=
<<
>>
&^

*/

func TestArithmetic(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	// 数组顺序不一样,返回false
	t.Log(a == b)
	// 数组长度不同,不可以比较
	//t.Log(a == c)
	// 数组顺序完成相同,返回true
	t.Log(a == d)
}
