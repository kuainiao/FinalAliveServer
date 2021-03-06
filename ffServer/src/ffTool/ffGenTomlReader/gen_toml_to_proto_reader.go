package main

import (
	"ffCommon/util"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var fmtTransPackage = `package main

import (
	"ffCommon/log/log"
	"ffCommon/util"
	"path/filepath"
	{ImportSort}

	proto "github.com/golang/protobuf/proto"
)
`

var fmtTransInit = `
func init() {
	allTrans = append(allTrans, trans{FileName})
}
`

var fmtTransFuncMain = `
func trans{FileName}() {
	message := &{FileName}{}
%v
	pbBuf := proto.NewBuffer(make([]byte, 0, 1024*10))
	if err := pbBuf.Marshal(message); err != nil {
		log.RunLogger.Printf("trans{FileName} err[%%v]", err)
		return
	}

	util.WriteFile(filepath.Join("ProtoBuf", "Client", "bytes", toml{FileName}.Name()+".bytes"), pbBuf.Bytes())
}
`

var fmtTransStructMap = `
	// {StructName}
	{MapKeyInt32Commet}{StructName}Keys := make([]{MapKeyInt32}, 0, len(toml{FileName}.{StructName})) // 必须使用64位机器
	{MapKeyInt64Commet}{StructName}Keys := make([]{MapKeyInt64}, 0, len(toml{FileName}.{StructName})) // 必须使用64位机器
	{MapKeyStringCommet}{StructName}Keys := make([]{MapKeyString}, 0, len(toml{FileName}.{StructName})) // 必须使用64位机器
	for key := range toml{FileName}.{StructName} {
		{MapKeyInt32Commet}{StructName}Keys = append({StructName}Keys, {MapKeyInt32}(key))
		{MapKeyInt64Commet}{StructName}Keys = append({StructName}Keys, {MapKeyInt64}(key))
		{MapKeyStringCommet}{StructName}Keys = append({StructName}Keys, {MapKeyString}(key))
	}
	{MapKeyInt32Commet}sort.Ints({StructName}Keys)
	{MapKeyInt64Commet}sort.Ints({StructName}Keys)
	{MapKeyStringCommet}sort.Strings({StructName}Keys)

	message.{StructName}Key = make([]{KeyType}, len(toml{FileName}.{StructName}))
	message.{StructName}Value = make([]*{FileName}_St{StructName}, len(toml{FileName}.{StructName}))
	for k, key := range {StructName}Keys {
		{MapKeyInt32Commet}i := {KeyType}(key)
		{MapKeyInt64Commet}i := {KeyType}(key)
		{MapKeyStringCommet}i := {KeyType}(key)
		v := toml{FileName}.{StructName}[i]

		message.{StructName}Key[k] = i
		message.{StructName}Value[k] = &{FileName}_St{StructName}{%v
		}
		%v
	}
`

var fmtTransStructStruct = `
	// {StructName}
	message.{StructName} = &{FileName}_St{StructName}{%v
	}
	%v
`

var fmtTransStructList = `
	// {StructName}
	message.{StructName} = make([]*{FileName}_St{StructName}, len(toml{FileName}.{StructName}))
	for k, v := range toml{FileName}.{StructName} {
		message.{StructName}[k] = &{FileName}_St{StructName}{%v
		}
		%v
	}
`

var fmtTransMemberBasic = "\n			{ProtoVar}: {GoDataVar}.{GoVar},"
var fmtTransMemberGrammar = "\n			{ProtoVar}: transGrammar({GoDataVar}.{GoVar}),"

var fmtTransMemberEnum = "\n		message.{StructName}{MapValue}[k].{ProtoVar} = {ProtoType}({GoDataVar}.{GoVar})"

var fmtTransMemberEnumArray = `
		message.{StructName}{MapValue}[k].{ProtoVar} = make({ProtoType}, len({GoDataVar}.{GoVar}), len({GoDataVar}.{GoVar}))
		for xx, yy := range {GoDataVar}.{GoVar} {
			message.{StructName}{MapValue}[k].{ProtoVar}[xx] = {ProtoVarContentType}(yy)
		}`

// 正则表达式说明
// http://www.cnblogs.com/golove/p/3269099.html
var regexpStruct = regexp.MustCompile(`type\s+([\w]+)\s+struct\s+{\n(?s)(.+?)\}`)
var regexpStructVar = regexp.MustCompile(`\s*([\w]+)\s+([\w\[\]\.\*]+)`) // 不捕获map

// 结构体定义
type structDef struct {
	name      string   // 结构体自身的定义
	vars      []string // 结构体的成员变量的名称
	types     []string // 结构体的成员变量的类型
	lowerVars []string // 结构体的成员变量的小写名称
}

// 文件内的所有结构体定义
type fileStructDef struct {
	name string
	defs []*structDef
}

// 解析文件内的结构体定义
func getFileDef(content string, filename string) *fileStructDef {
	// 捕获所有的结构体定义
	result1 := regexpStruct.FindAllStringSubmatch(content, -1)
	fileDef := &fileStructDef{
		name: filename,
		defs: make([]*structDef, len(result1), len(result1)),
	}

	for i, one := range result1 {
		// 结构体名称
		nameStruct := one[1]
		// 结构体成员
		allStructVars := one[2]

		// 捕获成员定义
		result2 := regexpStructVar.FindAllStringSubmatch(allStructVars, -1)

		structDef := &structDef{
			name:      nameStruct,
			vars:      make([]string, 0, len(result2)),
			lowerVars: make([]string, 0, len(result2)),
			types:     make([]string, 0, len(result2)),
		}

		//
		for _, two := range result2 {
			varName := two[1]
			varType := two[2]
			if filename != nameStruct && strings.HasPrefix(varType, "map[") {
				continue
			}

			structDef.vars = append(structDef.vars, varName)
			structDef.lowerVars = append(structDef.lowerVars, strings.ToLower(varName))
			structDef.types = append(structDef.types, varType)
		}

		fileDef.defs[i] = structDef
	}

	return fileDef
}

// 读取文件
func fileContent(fullpath string) (filename string, content string) {
	// 获取文件名, 不含扩展
	filename = filepath.Base(fullpath)
	fileext := filepath.Ext(fullpath)
	filename = filename[0 : len(filename)-len(fileext)]

	// 读取文件内容
	bytes, _ := util.ReadFile(fullpath)
	return filename, string(bytes)
}

// 生成转换代码
func genTransCode(saveFullDir string, protoFileDef, tomlFileDef *fileStructDef) {

	getStructDef := func(fileDef *fileStructDef, key string) *structDef {
		for _, structDef := range fileDef.defs {
			if key == strings.ToLower(structDef.name) {
				return structDef
			}
		}
		return nil
	}

	tomlMainStructDef := getStructDef(tomlFileDef, strings.ToLower(tomlFileDef.name))

	result := ""
	ImportSort := ""

	mainContent := strings.Replace(fmtTransFuncMain, "{FileName}", tomlFileDef.name, -1)

	allStructs := ""
	for _, tomlDef := range tomlFileDef.defs {
		if strings.ToLower(tomlDef.name) == strings.ToLower(tomlFileDef.name) {
			continue
		}

		// 工作簿在表格主类的成员类型
		mainStructVarType, mainStructVarTypeMapKey := "struct", ""
		for j, name := range tomlMainStructDef.vars {
			if name == tomlDef.name {
				if strings.HasPrefix(tomlMainStructDef.types[j], "map[") {
					mainStructVarType = "map"

					regexpMapKey := regexp.MustCompile(`\[([\w]+)\]`)
					result := regexpMapKey.FindAllStringSubmatch(tomlMainStructDef.types[j], -1)
					mainStructVarTypeMapKey = result[0][1]
				} else if strings.HasPrefix(tomlMainStructDef.types[j], "[]") {
					mainStructVarType = "list"
				}
				break
			}
		}

		protoStructDef := getStructDef(protoFileDef, strings.ToLower(tomlFileDef.name)+"_st"+strings.ToLower(tomlDef.name))

		// 来源go数据变量名称
		GoDataVar := "v"
		if mainStructVarType == "struct" {
			GoDataVar = "toml" + tomlFileDef.name + "." + tomlDef.name
		}

		// 基本成员变量以及Grammar
		hasEnumVarIndex := make([]int, 0, 1)
		members := ""
		for j := 0; j < len(tomlDef.vars); j++ {
			member := ""
			if tomlDef.types[j] != "ffGrammar.Grammar" {
				if tomlDef.types[j] != protoStructDef.types[j] {
					hasEnumVarIndex = append(hasEnumVarIndex, j)
					continue
				}

				member = fmtTransMemberBasic
			} else {
				member = fmtTransMemberGrammar
			}

			member = strings.Replace(member, "{ProtoVar}", protoStructDef.vars[j], -1)
			member = strings.Replace(member, "{GoDataVar}", GoDataVar, -1)
			member = strings.Replace(member, "{GoVar}", tomlDef.vars[j], -1)
			members += member
		}

		// 枚举成员
		enumMembers := ""
		if len(hasEnumVarIndex) > 0 {
			for _, j := range hasEnumVarIndex {
				member := ""
				ProtoVarContentType := protoStructDef.types[j]

				if strings.HasPrefix(tomlDef.types[j], "[]") {
					member = fmtTransMemberEnumArray
					ProtoVarContentType = protoStructDef.types[j][2:]
				} else if strings.HasPrefix(tomlDef.types[j], "map[") {
					// member = fmtTransMemberEnumArray
				} else {
					member = fmtTransMemberEnum
				}

				MapValue := ""
				if mainStructVarType == "map" {
					MapValue = "Value"
				}

				member = strings.Replace(member, "{StructName}", tomlDef.name, -1)
				member = strings.Replace(member, "{ProtoVar}", protoStructDef.vars[j], -1)
				member = strings.Replace(member, "{ProtoType}", protoStructDef.types[j], -1)
				member = strings.Replace(member, "{GoDataVar}", GoDataVar, -1)
				member = strings.Replace(member, "{GoVar}", tomlDef.vars[j], -1)
				member = strings.Replace(member, "{MapValue}", MapValue, -1)
				member = strings.Replace(member, "{ProtoVarContentType}", ProtoVarContentType, -1)

				enumMembers += member
			}
		}

		var structs string
		if mainStructVarType == "map" {
			ImportSort = `"sort"`
			structs = strings.Replace(fmtTransStructMap, "{FileName}", tomlFileDef.name, -1)
			structs = strings.Replace(structs, "{KeyType}", mainStructVarTypeMapKey, -1)

			MapKeyCommet := map[string]string{
				"int32":  "//",
				"int64":  "//",
				"string": "//",
			}
			MapKeyCommet[mainStructVarTypeMapKey] = ""

			MapKey := map[string]string{
				"int32":  "int",
				"int64":  "int",
				"string": "string",
			}

			structs = strings.Replace(structs, "{MapKeyInt32Commet}", MapKeyCommet["int32"], -1)
			structs = strings.Replace(structs, "{MapKeyInt64Commet}", MapKeyCommet["int64"], -1)
			structs = strings.Replace(structs, "{MapKeyStringCommet}", MapKeyCommet["string"], -1)
			structs = strings.Replace(structs, "{MapKeyInt32}", MapKey["int32"], -1)
			structs = strings.Replace(structs, "{MapKeyInt64}", MapKey["int64"], -1)
			structs = strings.Replace(structs, "{MapKeyString}", MapKey["string"], -1)

		} else if mainStructVarType == "list" {
			structs = strings.Replace(fmtTransStructList, "{FileName}", tomlFileDef.name, -1)
		} else {
			structs = strings.Replace(fmtTransStructStruct, "{FileName}", tomlFileDef.name, -1)
		}
		structs = strings.Replace(structs, "{StructName}", tomlDef.name, -1)
		structs = fmt.Sprintf(structs, members, enumMembers)

		allStructs += structs
	}
	mainContent = fmt.Sprintf(mainContent, allStructs)

	result += strings.Replace(fmtTransPackage, "{ImportSort}", ImportSort, -1)
	result += mainContent
	result += strings.Replace(fmtTransInit, "{FileName}", tomlFileDef.name, -1)

	util.WriteFile(filepath.Join(saveFullDir, "trans"+tomlFileDef.name+".go"), []byte(result))
}

// 转换
func transGoToProto(saveFullDir string, protoFilePath string, goFullPathFiles []string, packageName string) {
	r := regexp.MustCompile(`^trans[\w]*\.go`)

	// 移除之前的转换文件
	util.Walk(saveFullDir, func(info os.FileInfo) error {
		if info.IsDir() {
			return nil
		}

		if len(r.FindString(info.Name())) != 0 {
			util.RemoveFile(filepath.Join(saveFullDir, info.Name()))
		}

		return nil
	})

	// Proto的go代码
	var protoFileDef *fileStructDef
	{
		filename, content := fileContent(protoFilePath)
		protoFileDef = getFileDef(string(content), filename)

		fmt.Printf("%v:\n", filename)
		for _, v := range protoFileDef.defs {
			fmt.Printf("protoFileDef %v:%v\n%q\n%q\n\n", v.name, len(v.vars), v.vars, v.types)
		}
		fmt.Printf("\n\n")
	}

	// 读取toml的go代码
	for _, fullpath := range goFullPathFiles {
		filename, content := fileContent(fullpath)
		tomlFileDef := getFileDef(string(content), filename)

		fmt.Printf("%v:\n", filename)
		for _, v := range tomlFileDef.defs {
			fmt.Printf("tomlFileDef %v:%v\n%q\n%q\n\n", v.name, len(v.vars), v.vars, v.types)
		}
		fmt.Printf("\n\n")

		genTransCode(saveFullDir, protoFileDef, tomlFileDef)
	}
}
