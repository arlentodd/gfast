// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package wf_run_cache

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type AddData struct {
	RunId          uint   `p:"run_id"           json:"run_id"`           // 缓存run工作的全部流程模板步骤等信息,确保修改流程后工作依然不变
	FormId         uint   `p:"form_id"          json:"form_id"`          //
	FlowId         uint   `p:"flow_id"          json:"flow_id"`          // 流程ID
	RunForm        string `p:"run_form"         json:"run_form"`         // 模板信息
	RunFlow        string `p:"run_flow"         json:"run_flow"`         // 流程信息
	RunFlowProcess string `p:"run_flow_process" json:"run_flow_process"` // 流程步骤信息
	Dateline       uint   `p:"dateline"         json:"dateline"`         //
}

func Add(data *AddData, tx *gdb.TX) error {
	_, err := Model.TX(tx).Save(data)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("保存流程日志信息失败")
	}
	return nil
}
