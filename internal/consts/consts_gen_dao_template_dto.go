package consts

const TemplateGenDaoDtoContent = `
// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

{TplPackageImports}

// {TplTableNameCamelCase}ForDao is the golang structure of table {TplTableName} for DAO operations like Where/Data.
{TplStructDefine}
`
