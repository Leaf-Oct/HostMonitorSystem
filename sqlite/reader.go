package sqlite

import "database/sql"

func QueryTableAll(table_name string) *数据库中的表 {
	var rows *sql.Rows
	var err error
	switch table_name {
	case "xiajia":
		rows, err = host_db.Query(查询下架表所有信息的语句)
	case "jiangpei":
		rows, err = host_db.Query(查询降配表所有信息的语句)
	default:
		return nil
	}
	// rows, err := host_db.Query(查询表所有信息的语句, table_name)
	if err != nil {
		println("查询" + table_name + "表所有信息失败")
		println(err.Error())
		return nil
	}
	defer rows.Close()
	var result []服务器数据库条目
	columns, err := rows.Columns()
	if err != nil {
		println("获取" + table_name + "列信息失败")
		println(err.Error())
		return nil
	}
	for rows.Next() {
		ptrs := make([]interface{}, len(columns))
		vals := make([]string, len(columns))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			println("读取行信息失败")
			println(err.Error())
			return nil
		}
		result = append(result, 服务器数据库条目{vals[0], vals[1], vals[2], vals[3], vals[4], vals[5], vals[6], vals[7]})
	}
	return &数据库中的表{result}
}
