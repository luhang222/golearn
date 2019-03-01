package main

func main() {
	//b,_ := base64.StdEncoding.DecodeString("OEQzRTRBQUFBRDZGODZERRxKaz47hZMM31qtGuLlxIsDJplJR7lymVfIomdgGxxF4YPl9LviJ9/Ws3wFU+CmaM5BPB18v7XofaJOCQsPGu26HLEECei2HjNg1ld8oyd6")
	//fmt.Println(string(b))
	//a := []string{"111","222"}
	//t(&a)
	//t(&a)
	//tt(a)
	//fmt.Println(a)
}

func t(a *[]string) {

	*a = append(*a, "hello")
}

func tt(a []string) {
	a = append(a, "hello")
}
