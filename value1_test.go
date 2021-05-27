
package v8go

func TestNonPrimitiveType(t *testing.T){
	arrayUint8 := make([]uint8, 2)
	arrayUint8[0] = 120
	arrayUint8[1] = 130
	v, err := NewValue(iso, arrayUint8)
			if err != nil {
				fmt.Println("error creating new val ", err)
				return val
			}

}