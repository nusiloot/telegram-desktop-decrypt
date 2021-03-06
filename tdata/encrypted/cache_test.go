package encrypted

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/atilaromero/telegram-desktop-decrypt/tdata"
)

const hexCache = "54444624ff4d0f00000000205103d8223659d1c68467b0f14a3b4bb8ad16e832eaad688c280d46ffe9ca803263243b89e392ced3aa364d33b7fe7efd"

func ExampleToCache() {
	cache0, _ := hex.DecodeString(hexCache)
	rawtdf, err := tdata.ReadRawTDF(bytes.NewReader(cache0))
	if err != nil {
		rawtdf.Print(false)
	}
	ecache, err := ReadECache(rawtdf)
	if err != nil {
		fmt.Println("error converting to cache:", err)
	}
	fmt.Println(hex.EncodeToString(ecache.Encrypted))
	//Output:
	// 5103d8223659d1c68467b0f14a3b4bb8ad16e832eaad688c280d46ffe9ca8032
}
