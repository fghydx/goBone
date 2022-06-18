package GLUtils

import (
	"database/sql"
	"fmt"
	"strings"
)

//生成类型
func RowsToStruct(row sql.Rows,structName string) (result string) {
	columnstype,_ := row.ColumnTypes()
	strb := strings.Builder{}
	strb.WriteString("type "+structName+" struct{\n")
	for i, _ := range columnstype {
		tmp := columnstype[i].Name()

		switch {
		case columnstype[i].DatabaseTypeName()=="DECIMAL" :
			strb.WriteString("	"+strings.ToUpper(tmp[:1])+tmp[1:]+"	float64"+
				fmt.Sprintf("	`col:\"%s,omitempty\" json:\"%s,omitempty\"` \n",tmp,strings.ToLower(tmp)))
		default:
			strb.WriteString("	"+strings.ToUpper(tmp[:1])+tmp[1:]+"	"+columnstype[i].ScanType().Name() +
				fmt.Sprintf("	`col:\"%s,omitempty\" json:\"%s,omitempty\"` \n",tmp,strings.ToLower(tmp)))
		}
	}
	strb.WriteString("}")
	return strb.String()
}
