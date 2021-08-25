package likou

func Find(arrayA, arrayB []int)(result []int)  {
	a,b:=0,0
	for a<=len(arrayA)-1 && b<=len(arrayB)-1{
		if arrayA[a]==arrayB[b]{
			result = append(result,arrayA[a])
			a++
			b++
			continue
		}
		if arrayA[a]<arrayB[b]{
			a++
		}else{
			b++
		}
	}

	return result
}
