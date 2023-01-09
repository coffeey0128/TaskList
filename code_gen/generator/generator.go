package generator

import (
	"TaskList/config"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"gorm.io/gorm"
)

const (
	PREVENTCOLUMN_ID          = "id"
	PREVENT_COLUMN_CREATED_AT = "created_at"
	PREVENT_COLUMN_UPDATED_AT = "updated_at"
)

var (
	BasePath        = config.GetBasePath()
	REPOSITORY      = BasePath + "/code_gen/templates/repository_gorm.tmpl"
	REPOSITORY_TEST = BasePath + "/code_gen/templates/repository_gorm_test.tmpl"
	SERVICE         = BasePath + "/code_gen/templates/service.tmpl"
	SERVICE_TEST    = BasePath + "/code_gen/templates/service_test.tmpl"
	MODEL           = BasePath + "/code_gen/templates/model.tmpl"
	API             = BasePath + "/code_gen/templates/api.tmpl"
	API_TEST        = BasePath + "/code_gen/templates/api_test.tmpl"
	Route           = BasePath + "/code_gen/templates/router.tmpl"
	APIREQ          = BasePath + "/code_gen/templates/apireq.tmpl"
	APIRES          = BasePath + "/code_gen/templates/apires.tmpl"
)

var preventColumns = map[string]bool{
	PREVENTCOLUMN_ID:          true,
	PREVENT_COLUMN_CREATED_AT: true,
	PREVENT_COLUMN_UPDATED_AT: true,
}

// map for converting mysql type to golang types
var typeForMysqlToGo = map[string]string{
	"int":                "int64",
	"integer":            "int64",
	"tinyint":            "int64",
	"smallint":           "int64",
	"mediumint":          "int64",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int64",
	"smallint unsigned":  "int64",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
	"json":               "string",
}

type Generator struct {
	StructName        string
	ProjectName       string
	TableName         string
	InternalPath      string
	ModelPackageName  string
	GenRepo           bool
	GenService        bool
	GenModel          bool
	GenApi            bool
	GenRoute          bool
	Db                *gorm.DB `json:"-"`
	Templates         map[string]*template.Template
	TableColumns      []column
	APIRequestColumns []column
	WireInfo          WireInfo
	APIRoute          string
	TimeField         bool
	DisableField      bool
}

type WireInfo struct {
	StructName          string
	VariableName        string
	ConstructorSrvName  string
	ConstructorRepoName string
	TableName           string
	ImportPath          string
}

type column struct {
	ColumnName    string
	Type          string
	Nullable      string
	TableName     string
	ColumnComment string
	Tag           string
	ModelTag      string
	ModelField    string
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Table(t string) *Generator {
	g.TableName = t
	return g
}

func (g *Generator) DB(db *gorm.DB) *Generator {
	g.Db = db
	return g
}

// Function for fetching schema definition of passed table
func (g *Generator) GetColumns() (err error) {
	// sql
	var sqlStr = `SELECT
					COLUMN_NAME as column_name,
				  	DATA_TYPE as type,
					IS_NULLABLE as nullable,
					TABLE_NAME as table_name,
					COLUMN_COMMENT as column_comment
				FROM information_schema.COLUMNS 
				WHERE table_schema = DATABASE()`
	sqlStr += fmt.Sprintf(`AND TABLE_NAME = "%s"`, g.TableName)

	cols := []column{}
	result := g.Db.Raw(sqlStr).Scan(&cols)
	if result.Error != nil {
		return err
	}
	g.getModelTag(cols)
	g.TableColumns = cols
	g.StructName = g.ToUpper(g.TableName)
	g.WireInfo = WireInfo{
		TableName:           g.TableName,
		VariableName:        strings.ToLower(g.StructName[0:1]) + g.StructName[1:],
		StructName:          g.StructName,
		ConstructorSrvName:  fmt.Sprintf("New%sService", g.StructName),
		ConstructorRepoName: fmt.Sprintf("New%sRepo", g.StructName),
		ImportPath:          fmt.Sprintf("%s/internal/%s%s", g.ProjectName, g.InternalPath, g.TableName),
	}
	g.getAPIRoute()

	return nil
}

func (g *Generator) CreateFolder() {
	internalPath := BasePath + "/internal/" + g.InternalPath + g.TableName
	modelPath := BasePath + "/models/"
	routePath := BasePath + "/routes/"
	apiPath := BasePath + "/api/"
	apiTestPath := BasePath + "/api_test/"
	apireqPath := BasePath + "/models/apireq/"
	apiresPath := BasePath + "/models/apires/"
	paths := []string{internalPath, modelPath, routePath, apiPath, apiTestPath, apireqPath, apiresPath}
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (g *Generator) ToUpper(s string) string {
	CamelCaseTableName := ""
	tableNames := strings.Split(s, "_")
	for _, value := range tableNames {
		CamelCaseTableName += strings.ToUpper(value[0:1]) + value[1:]
	}
	return CamelCaseTableName
}

func (g *Generator) getModelTag(cols []column) {
	apiRequestColumns := make([]column, 0)
	for k, col := range cols {
		cols[k].Type = typeForMysqlToGo[col.Type]
		if cols[k].Type == "time.Time" {
			g.TimeField = true
		}
		cols[k].Tag = col.ColumnName
		cols[k].ModelField = g.ToUpper(col.ColumnName)
		if cols[k].ModelField == "IsDisable" {
			g.DisableField = true
		}
	}
	for _, col := range cols {
		if _, ok := preventColumns[col.ColumnName]; !ok {
			apiRequestColumns = append(apiRequestColumns, col)
		}
	}
	g.APIRequestColumns = apiRequestColumns
}

func (g *Generator) getAPIRoute() {
	g.APIRoute = strings.ReplaceAll(g.TableName, "_", "-") + "s"
}

func (g *Generator) LoadTemplates() {
	internalPath := fmt.Sprintf("%s/internal/%s%s", BasePath, g.InternalPath, g.TableName)
	g.Templates = make(map[string]*template.Template, 0)
	if g.GenRepo {
		tmpl := template.Must(template.ParseFiles(REPOSITORY))
		repoPath := internalPath + "/repository.go"
		g.Templates[repoPath] = tmpl
		testTmpl := template.Must(template.ParseFiles(REPOSITORY_TEST))
		repoTestPath := internalPath + "/repository_test.go"
		g.Templates[repoTestPath] = testTmpl
	}
	if g.GenService {
		tmpl := template.Must(template.ParseFiles(SERVICE))
		srvPath := internalPath + "/service.go"
		g.Templates[srvPath] = tmpl
		testTmpl := template.Must(template.ParseFiles(SERVICE_TEST))
		srvTestPath := internalPath + "/service_test.go"
		g.Templates[srvTestPath] = testTmpl
	}

	if g.GenModel {
		tmpl := template.Must(template.ParseFiles(MODEL))
		modelPath := BasePath + "/models/" + g.TableName + ".go"
		g.Templates[modelPath] = tmpl
		tmpl = template.Must(template.ParseFiles(APIREQ))
		apireqPath := BasePath + "/models/apireq/" + g.TableName + ".go"
		g.Templates[apireqPath] = tmpl
		tmpl = template.Must(template.ParseFiles(APIRES))
		apiresPath := BasePath + "/models/apires/" + g.TableName + ".go"
		g.Templates[apiresPath] = tmpl
	}
	if g.GenApi {
		tmpl := template.Must(template.ParseFiles(API))
		apiPath := BasePath + "/api/" + g.TableName + "_api.go"
		g.Templates[apiPath] = tmpl
		testTmpl := template.Must(template.ParseFiles(API_TEST))
		apiTestPath := BasePath + "/api_test/" + g.TableName + "_test.go"
		g.Templates[apiTestPath] = testTmpl
	}
	if g.GenRoute {
		tmpl := template.Must(template.ParseFiles(Route))
		modelPath := BasePath + "/routes/" + g.TableName + ".go"
		g.Templates[modelPath] = tmpl
	}
}

func (g *Generator) WriteWire() {
	// append data to wire.go
	wirePath := BasePath + "/api/wire.go"
	f, err := os.OpenFile(wirePath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1
	fileStr := []string{}
	for scanner.Scan() {
		fileStr = append(fileStr, scanner.Text())
		if strings.Contains(scanner.Text(), "import") {
			fileStr = append(fileStr, fmt.Sprintf(`	%s "%s"`, g.WireInfo.VariableName, g.WireInfo.ImportPath))
		}
		line++
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		panic(err)
	}
	srvSet := fmt.Sprintf(`var %sServiceSet = wire.NewSet(%s.%s, %s.%s)`, g.WireInfo.StructName,
		g.WireInfo.VariableName,
		g.WireInfo.ConstructorSrvName,
		g.WireInfo.VariableName,
		g.WireInfo.ConstructorRepoName)
	buildFunc := fmt.Sprintf(`func Build%sSrv() %s.Service {
	wire.Build(%sServiceSet, gormSet)
	return &%s.%sService{}
}`,
		g.WireInfo.StructName,
		g.WireInfo.VariableName,
		g.WireInfo.StructName,
		g.WireInfo.VariableName,
		g.WireInfo.StructName)
	fileStr = append(fileStr, srvSet)
	fileStr = append(fileStr, buildFunc)
	s := ""
	for _, str := range fileStr {
		s += str
		s += "\n"
	}
	f.Truncate(0)
	f.Seek(0, 0)
	f.WriteString(s)
}

func (g *Generator) WriteMakefile() {
	// append data to Makefile
	filePath := BasePath + "/Makefile"
	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1
	fileStr := []string{}
	for scanner.Scan() {
		fileStr = append(fileStr, scanner.Text())
		if strings.Contains(scanner.Text(), "gen-mock:") {
			mockPath := g.InternalPath + g.TableName
			fileStr = append(fileStr, fmt.Sprintf(`	mockgen -source=./internal/%s/repository.go -destination=./mock/%s/repository.go -package=mock_%s`, mockPath, mockPath, g.TableName))
		}
		line++
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		panic(err)
	}
	s := ""
	for _, str := range fileStr {
		s += str
		s += "\n"
	}
	f.Truncate(0)
	f.Seek(0, 0)
	f.WriteString(s)
}

func (g *Generator) WriteRoute() {
	// append data to Makefile
	filePath := BasePath + "/routes/route.go"
	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1
	fileStr := []string{}
	for scanner.Scan() {
		fileStr = append(fileStr, scanner.Text())
		if strings.Contains(scanner.Text(), "generate by code-gen") {
			fileStr = append(fileStr, fmt.Sprintf(`	%sV1(r)`, g.StructName))
		}
		line++
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		panic(err)
	}
	s := ""
	for _, str := range fileStr {
		s += str
		s += "\n"
	}
	f.Truncate(0)
	f.Seek(0, 0)
	f.WriteString(s)
}

func (g *Generator) Execute() error {
	var data map[string]interface{}
	b, err := json.Marshal(g)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	g.CreateFolder()
	for path, template := range g.Templates {
		var buf bytes.Buffer
		err = template.Execute(&buf, data)
		if err != nil {
			return err
		}
		err = os.WriteFile(path, buf.Bytes(), 0644)
		if err != nil {
			return err
		}
	}
	g.WriteWire()
	g.WriteMakefile()
	g.WriteRoute()
	return nil
}
