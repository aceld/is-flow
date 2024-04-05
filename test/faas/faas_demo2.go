package faas

import (
	"context"
	"fmt"
	"github.com/aceld/kis-flow/kis"
	"log/slog"
)

// type FaaS func(context.Context, Flow) error

func FuncDemo2Handler(ctx context.Context, flow kis.Flow) error {
	fmt.Println("---> Call funcName2Handler ----")
	fmt.Printf("Params = %+v\n", flow.GetFuncParamAll())

	for index, row := range flow.Input() {
		str := fmt.Sprintf("In FuncName = %s, FuncId = %s, row = %s", flow.GetThisFuncConf().FName, flow.GetThisFunction().GetId(), row)
		fmt.Println(str)

		conn, err := flow.GetConnector()
		if err != nil {
			slog.ErrorContext(ctx, "FuncDemo2Handler(): GetConnector", "err", err.Error())
			return err
		}

		if _, err := conn.Call(ctx, flow, row); err != nil {
			slog.ErrorContext(ctx, "FuncDemo2Handler(): Call", "err", err.Error())
			return err
		}

		// 计算结果数据
		resultStr := fmt.Sprintf("data from funcName[%s], index = %d", flow.GetThisFuncConf().FName, index)

		// 提交结果数据
		_ = flow.CommitRow(resultStr)
	}

	return nil
}
