// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package views

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
)

func (l *gCurd) webEditTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["formItem"] = l.generateWebEditFormItem(ctx, in)
	data["script"] = l.generateWebEditScript(ctx, in)
	return
}

func (l *gCurd) generateWebEditFormItem(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	for k, field := range in.masterFields {
		if !field.IsEdit {
			continue
		}

		if field.Index == consts.GenCodesIndexPK {
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-input placeholder=\"请输入%s\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)
			component        string
		)

		switch field.FormMode {
		case FormModeInput:
			component = defaultComponent

		case FormModeInputNumber:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-input-number placeholder=\"请输入%s\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)

		case FormModeInputTextarea:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-input type=\"textarea\" placeholder=\"%s\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)

		case FormModeInputEditor:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <Editor style=\"height: 450px\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeInputDynamic:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-dynamic-input\n            v-model:value=\"params.%s\"\n            preset=\"pair\"\n            key-placeholder=\"键名\"\n            value-placeholder=\"键值\"\n          />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeDate:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <DatePicker v-model:formValue=\"params.%s\" type=\"date\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		//case FormModeDateRange:  // 必须要有两个字段，后面优化下

		case FormModeTime:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <DatePicker v-model:formValue=\"params.%s\" type=\"datetime\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		//case FormModeTimeRange: // 必须要有两个字段，后面优化下

		case FormModeRadio:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-radio-group v-model:value=\"params.%s\" name=\"%s\">\n            <n-radio-button\n              v-for=\"%s in options.%s\"\n              :key=\"%s.value\"\n              :value=\"%s.value\"\n              :label=\"%s.label\"\n            />\n          </n-radio-group>\n        </n-form-item>", field.Dc, field.TsName, field.TsName, field.TsName, field.TsName, in.options.dictMap[field.TsName], field.TsName, field.TsName, field.TsName)

		case FormModeCheckbox:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-checkbox-group v-model:value=\"params.%s\">\n            <n-space>\n              <n-checkbox\n                v-for=\"item in options.%s\"\n                :key=\"item.value\"\n                :value=\"item.value\"\n                :label=\"item.label\"\n              />\n            </n-space>\n          </n-checkbox-group>\n        </n-form-item>", field.Dc, field.TsName, field.TsName, in.options.dictMap[field.TsName])

		case FormModeSelect:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-select v-model:value=\"params.%s\" :options=\"options.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName, in.options.dictMap[field.TsName])

		case FormModeSelectMultiple:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-select multiple v-model:value=\"params.%s\" :options=\"options.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName, in.options.dictMap[field.TsName])

		case FormModeUploadImage:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <UploadImage :maxNumber=\"1\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeUploadImages:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <UploadImage :maxNumber=\"10\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeUploadFile:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <UploadFile :maxNumber=\"1\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeUploadFiles:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <UploadFile :maxNumber=\"10\" v-model:value=\"params.%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeSwitch:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-switch v-model:value=\"params.%s\"\n        />\n        </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeRate:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-rate allow-half :default-value=\"params.%s\" :on-update:value=\"update%s\" />\n        </n-form-item>", field.Dc, field.TsName, field.TsName, field.GoName)

		default:
			component = defaultComponent
		}

		if len(in.masterFields) == k {
			buffer.WriteString("        " + component)
		} else {
			buffer.WriteString("        " + component + "\n\n")
		}

	}
	return buffer.String()
}

func (l *gCurd) generateWebEditScript(ctx context.Context, in *CurdPreviewInput) g.Map {
	var (
		data         = make(g.Map)
		importBuffer = bytes.NewBuffer(nil)
		setupBuffer  = bytes.NewBuffer(nil)
	)

	if in.options.Step.HasMaxSort {
		importBuffer.WriteString("  import { onMounted, ref, computed, watch } from 'vue';\n")
		importBuffer.WriteString("  import { Edit, MaxSort } from '@/api/" + gstr.LcFirst(in.In.VarName) + "';\n")
		setupBuffer.WriteString("  watch(\n    () => params.value,\n    (value) => {\n      if (value.id === 0) {\n        MaxSort().then((res) => {\n          params.value.sort = res.sort;\n        });\n      }\n    }\n  );\n\n")
	} else {
		importBuffer.WriteString("  import { onMounted, ref, computed } from 'vue';\n")
		importBuffer.WriteString("  import { Edit } from '@/api/" + gstr.LcFirst(in.In.VarName) + "';\n")
	}

	for _, field := range in.masterFields {
		if !field.IsEdit {
			continue
		}
		switch field.FormMode {
		case FormModeDate, FormModeDateRange, FormModeTime, FormModeTimeRange:
			if !gstr.Contains(importBuffer.String(), `import DatePicker`) {
				importBuffer.WriteString("  import DatePicker from '@/components/DatePicker/datePicker.vue';\n")
			}
		case FormModeInputEditor:
			if !gstr.Contains(importBuffer.String(), `import Editor`) {
				importBuffer.WriteString("  import Editor from '@/components/Editor/editor.vue';\n")
			}
		case FormModeUploadImage, FormModeUploadImages:
			if !gstr.Contains(importBuffer.String(), `import UploadImage`) {
				importBuffer.WriteString("  import UploadImage from '@/components/Upload/uploadImage.vue';\n")
			}
		case FormModeUploadFile, FormModeUploadFiles:
			if !gstr.Contains(importBuffer.String(), `import UploadFile`) {
				importBuffer.WriteString("  import UploadFile from '@/components/Upload/uploadFile.vue';\n")
			}
		case FormModeRate:
			setupBuffer.WriteString(fmt.Sprintf("  function update%s(num) {\n    params.value.%s = num;\n  }\n", field.GoName, field.TsName))
		}

	}

	data["import"] = importBuffer.String()
	data["setup"] = setupBuffer.String()

	return data
}