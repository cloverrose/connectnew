package connectnew

import (
	"flag"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const doc = "connectnew enforces the use of constructor functions (connect.NewRequest/connect.NewResponse) instead of direct struct initialization."

// Analyzer checks if &connect.Request or &connect.Response are used.
var Analyzer = &analysis.Analyzer{
	Name:     "connectnew",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{},
	Flags:    *flag.NewFlagSet("connectnew", flag.ExitOnError),
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// ユニリ演算式をチェック
			if unaryExpr, ok := n.(*ast.UnaryExpr); ok {
				// アドレス演算子 `&` を検出
				if unaryExpr.Op == token.AND {
					// コンポジットリテラルをチェック
					if compositeLit, ok := unaryExpr.X.(*ast.CompositeLit); ok {
						// 型アサーション
						if inst, ok := compositeLit.Type.(*ast.IndexExpr); ok {
							// セレクタ式かどうかをチェック
							if selectorExpr, ok := inst.X.(*ast.SelectorExpr); ok {
								// セレクタが `connect.Request` または `connect.Response` かを確認
								if ident, ok := selectorExpr.X.(*ast.Ident); ok {
									if ident.Name == "connect" {
										if selectorExpr.Sel.Name == "Request" {
											pass.Reportf(unaryExpr.Pos(), "use of &connect.Request[T]{} detected. Use connect.NewRequest() instead.")
										}
										if selectorExpr.Sel.Name == "Response" {
											pass.Reportf(unaryExpr.Pos(), "use of &connect.Response[T]{} detected. Use connect.NewResponse() instead.")
										}
									}
								}
							}
						}
					}
				}
			}

			return true
		})
	}

	return nil, nil
}
