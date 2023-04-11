package doc

import (
	"db-doc/model"
	"db-doc/util"
	"fmt"
	"path"
	"strings"
)

// createOnlineDoc create _siderbar.md
func createOnlineDoc(docPath string, dbInfo model.DbInfo, tables []model.Table) {
    
	var sidebar []string
	var readme []string
	// sidebar = append(sidebar, "* [数据库文档](README.md)")
	readme = append(readme, fmt.Sprintf("# %s 数据库文档", dbInfo.DbName))
	// 生成基础信息
	readme = append(readme, "### 基础信息")
	readme = append(readme, "| 数据库名称 | 版本 | 字符集 | 排序规则 |")
	readme = append(readme, "| ---- | ---- | ---- | ---- |")
	readme = append(readme, fmt.Sprintf("| %s | %s | %s | %s |", dbInfo.DbName, dbInfo.Version, dbInfo.Charset, dbInfo.Collation))
	for i := range tables {
		sidebar = append(sidebar, fmt.Sprintf("* [%s(%s)](%s.md)", tables[i].TableName, tables[i].TableComment, tables[i].TableName))
		var tableMd []string
		tableMd = append(tableMd, fmt.Sprintf("# %s(%s)", tables[i].TableName, tables[i].TableComment))
		tableMd = append(tableMd, "| 列名 | 类型 | KEY | 可否为空 | 默认值 | 注释 |")
		tableMd = append(tableMd, "| ---- | ---- | ---- | ---- | ---- | ----  |")
		// create table.md
		cols := tables[i].ColList
		for j := range cols {
			tableMd = append(tableMd, fmt.Sprintf("| %s | %s | %s | %s | %s | %s |",
				cols[j].ColName, cols[j].ColType, cols[j].ColKey, cols[j].IsNullable, cols[j].ColDefault, cols[j].ColComment))
		}
		tableStr := strings.Join(tableMd, "\r\n")
		util.WriteToFile(path.Join(docPath, tables[i].TableName+".md"), tableStr)
	}
	// create readme.md
	readmeStr := strings.Join(readme, "\r\n")
	util.WriteToFile(path.Join(docPath, "README.md"), readmeStr)
	// create _sidebar.md
	sidebarStr := strings.Join(sidebar, "\r\n")
	util.WriteToFile(path.Join(docPath, "_sidebar.md"), sidebarStr)
	// create index.html
	util.WriteToFile(path.Join(docPath, "index.html"), docsifyHTML)
	// create .nojekyll
	util.WriteToFile(path.Join(docPath, ".nojekyll"), "")
	fmt.Println("doc generate successfully!")
}

