package format

import (
	"fmt"

	"github.com/xmirrorsecurity/opensca-cli/cmd/detail"
)

// Statis 统计概览信息
func Statis(report Report) string {

	// 组件风险统计 key:0代表组件总数
	depStatic := map[int]int{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0,
	}
	// 记录统计过的组件
	depSet := map[string]bool{}

	// 漏洞风险统计
	vulStatic := map[int]int{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0,
	}
	// 记录统计过的漏洞
	vulSet := map[string]bool{}

	report.DepDetailGraph.ForEach(func(n *detail.DepDetailGraph) bool {

		_, ok := depSet[n.Dep.Key()]
		if ok || n.Name == "" {
			return true
		}
		depSet[n.Dep.Key()] = true

		// 当前组件风险
		risk := 5
		for _, v := range n.Vulnerabilities {
			if !vulSet[v.Id] {
				vulSet[v.Id] = true
				vulStatic[v.SecurityLevelId]++
				vulStatic[0]++
			}
			if v.SecurityLevelId < risk {
				// 组件风险取最高漏洞风险
				risk = v.SecurityLevelId
			}
		}

		depStatic[risk]++
		depStatic[0]++

		return true
	})

	return fmt.Sprintf("\n\nComplete!"+
		"\nComponents:%d C:%d H:%d M:%d L:%d"+
		"\nVulnerabilities:%d C:%d H:%d M:%d L:%d",
		depStatic[0], depStatic[1], depStatic[2], depStatic[3], depStatic[4],
		vulStatic[0], vulStatic[1], vulStatic[2], vulStatic[3], vulStatic[4],
	)
}