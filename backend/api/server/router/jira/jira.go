package jira

import (
	"context"
	"net/http"

	"github.com/redhat-appstudio/quality-studio/pkg/utils/httputils"
)

// Jira godoc
// @Summary Jira API Info
// @Description returns a list of jira issues which contain the label appstudio-e2e-tests-known-issues
// @Tags Jira API Info
// @Produce json
// @Router /jira/bugs/e2e [get]
// @Success 200 {object} []jira.Issue
func (s *jiraRouter) listE2EBugsKnown(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	issues := s.Jira.GetIssueByJQLQuery(`project in (STONE, DEVHAS, SRVKP, GITOPSRVCE, HACBS) AND status not in (Closed) AND labels = ci-fail`)

	return httputils.WriteJSON(w, http.StatusOK, issues)
}
