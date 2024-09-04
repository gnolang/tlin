package lints

import (
	"go/ast"
	"go/token"

	tt "github.com/gnoswap-labs/tlin/internal/types"
)

// DetectUselessBreak detects useless break statements in switch or select statements.
func DetectUselessBreak(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	var issues []tt.Issue
	ast.Inspect(node, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.SwitchStmt:
			for _, stmt := range v.Body.List {
				if caseClause, ok := stmt.(*ast.CaseClause); ok {
					checkUselessBreak(caseClause.Body, filename, fset, &issues)
				}
			}
		case *ast.SelectStmt:
			for _, stmt := range v.Body.List {
				if commClause, ok := stmt.(*ast.CommClause); ok {
					checkUselessBreak(commClause.Body, filename, fset, &issues)
				}
			}
		}
		return true
	})

	return issues, nil
}

func checkUselessBreak(stmts []ast.Stmt, filename string, fset *token.FileSet, issues *[]tt.Issue) {
	if len(stmts) == 0 {
		return
	}

	lastStmt := stmts[len(stmts)-1]
	if breakStmt, ok := lastStmt.(*ast.BranchStmt); ok && breakStmt.Tok == token.BREAK && breakStmt.Label == nil {
		startPos := fset.Position(breakStmt.Pos())
		endPos := fset.Position(breakStmt.End())

		*issues = append(*issues, tt.Issue{
			Rule:     "useless-break",
			Filename: filename,
			Start: tt.UniversalPosition{
				Filename: filename,
				Line:     startPos.Line,
				Column:   startPos.Column,
				Offset:   startPos.Offset,
				Length:   endPos.Offset - startPos.Offset,
			},
			End: tt.UniversalPosition{
				Filename: filename,
				Line:     endPos.Line,
				Column:   endPos.Column,
				Offset:   endPos.Offset,
				Length:   0,
			},
			Message: "useless break statement at the end of case clause",
		})
	}
}
