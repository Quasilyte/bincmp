// Code generated by "stringer -type=SymbolType"; DO NOT EDIT

package nm

import "fmt"

const _SymbolType_name = "SymbolTypeUnknownSymbolTypeBSSSymbolTypeGlobalBSSSymbolTypeDataSymbolTypeGlobalDataSymbolTypeTextSymbolTypeGlobalTextSymbolTypeReadOnlyDataSymbolTypeGlobalReadOnlyData"

var _SymbolType_index = [...]uint8{0, 17, 30, 49, 63, 83, 97, 117, 139, 167}

func (i SymbolType) String() string {
	if i >= SymbolType(len(_SymbolType_index)-1) {
		return fmt.Sprintf("SymbolType(%d)", i)
	}
	return _SymbolType_name[_SymbolType_index[i]:_SymbolType_index[i+1]]
}
