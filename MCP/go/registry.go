package main

import (
	"github.com/asana/mcp-server/config"
	"github.com/asana/mcp-server/models"
	tools_tags "github.com/asana/mcp-server/tools/tags"
	tools_custom_field_settings "github.com/asana/mcp-server/tools/custom_field_settings"
	tools_projects "github.com/asana/mcp-server/tools/projects"
	tools_tasks "github.com/asana/mcp-server/tools/tasks"
	tools_workspaces "github.com/asana/mcp-server/tools/workspaces"
	tools_team_memberships "github.com/asana/mcp-server/tools/team_memberships"
	tools_teams "github.com/asana/mcp-server/tools/teams"
	tools_sections "github.com/asana/mcp-server/tools/sections"
	tools_project_templates "github.com/asana/mcp-server/tools/project_templates"
	tools_custom_fields "github.com/asana/mcp-server/tools/custom_fields"
	tools_events "github.com/asana/mcp-server/tools/events"
	tools_workspace_memberships "github.com/asana/mcp-server/tools/workspace_memberships"
	tools_goals "github.com/asana/mcp-server/tools/goals"
	tools_goal_relationships "github.com/asana/mcp-server/tools/goal_relationships"
	tools_portfolios "github.com/asana/mcp-server/tools/portfolios"
	tools_project_briefs "github.com/asana/mcp-server/tools/project_briefs"
	tools_project_memberships "github.com/asana/mcp-server/tools/project_memberships"
	tools_users "github.com/asana/mcp-server/tools/users"
	tools_time_periods "github.com/asana/mcp-server/tools/time_periods"
	tools_portfolio_memberships "github.com/asana/mcp-server/tools/portfolio_memberships"
	tools_stories "github.com/asana/mcp-server/tools/stories"
	tools_webhooks "github.com/asana/mcp-server/tools/webhooks"
	tools_organization_exports "github.com/asana/mcp-server/tools/organization_exports"
	tools_user_task_lists "github.com/asana/mcp-server/tools/user_task_lists"
	tools_status_updates "github.com/asana/mcp-server/tools/status_updates"
	tools_attachments "github.com/asana/mcp-server/tools/attachments"
	tools_audit_log_api "github.com/asana/mcp-server/tools/audit_log_api"
	tools_batch_api "github.com/asana/mcp-server/tools/batch_api"
	tools_project_statuses "github.com/asana/mcp-server/tools/project_statuses"
	tools_jobs "github.com/asana/mcp-server/tools/jobs"
	tools_typeahead "github.com/asana/mcp-server/tools/typeahead"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tags.CreateDeletetagTool(cfg),
		tools_tags.CreateGettagTool(cfg),
		tools_tags.CreateUpdatetagTool(cfg),
		tools_custom_field_settings.CreateGetcustomfieldsettingsforportfolioTool(cfg),
		tools_projects.CreateGetprojectsforworkspaceTool(cfg),
		tools_projects.CreateCreateprojectforworkspaceTool(cfg),
		tools_tasks.CreateGettasksforsectionTool(cfg),
		tools_workspaces.CreateGetworkspaceTool(cfg),
		tools_workspaces.CreateUpdateworkspaceTool(cfg),
		tools_team_memberships.CreateGetteammembershipTool(cfg),
		tools_teams.CreateGetteamsforuserTool(cfg),
		tools_sections.CreateInsertsectionforprojectTool(cfg),
		tools_project_templates.CreateGetprojecttemplateTool(cfg),
		tools_custom_fields.CreateCreatecustomfieldTool(cfg),
		tools_events.CreateGeteventsTool(cfg),
		tools_projects.CreateRemovefollowersforprojectTool(cfg),
		tools_tasks.CreateRemovedependenciesfortaskTool(cfg),
		tools_workspace_memberships.CreateGetworkspacemembershipsforuserTool(cfg),
		tools_goals.CreateRemovefollowersTool(cfg),
		tools_projects.CreateDuplicateprojectTool(cfg),
		tools_sections.CreateAddtaskforsectionTool(cfg),
		tools_tasks.CreateAddtagfortaskTool(cfg),
		tools_custom_fields.CreateInsertenumoptionforcustomfieldTool(cfg),
		tools_workspaces.CreateAdduserforworkspaceTool(cfg),
		tools_tasks.CreateGettasksforprojectTool(cfg),
		tools_tags.CreateGettagsforworkspaceTool(cfg),
		tools_tags.CreateCreatetagforworkspaceTool(cfg),
		tools_goal_relationships.CreateAddsupportingrelationshipTool(cfg),
		tools_tasks.CreateAddfollowersfortaskTool(cfg),
		tools_teams.CreateRemoveuserforteamTool(cfg),
		tools_custom_fields.CreateCreateenumoptionforcustomfieldTool(cfg),
		tools_portfolios.CreateRemovecustomfieldsettingforportfolioTool(cfg),
		tools_portfolios.CreateAdditemforportfolioTool(cfg),
		tools_project_briefs.CreateCreateprojectbriefTool(cfg),
		tools_project_memberships.CreateGetprojectmembershipTool(cfg),
		tools_users.CreateGetfavoritesforuserTool(cfg),
		tools_team_memberships.CreateGetteammembershipsforuserTool(cfg),
		tools_projects.CreateAddmembersforprojectTool(cfg),
		tools_tasks.CreateRemoveprojectfortaskTool(cfg),
		tools_project_templates.CreateInstantiateprojectTool(cfg),
		tools_time_periods.CreateGettimeperiodsTool(cfg),
		tools_custom_fields.CreateDeletecustomfieldTool(cfg),
		tools_custom_fields.CreateGetcustomfieldTool(cfg),
		tools_custom_fields.CreateUpdatecustomfieldTool(cfg),
		tools_tasks.CreateGettasksforusertasklistTool(cfg),
		tools_portfolio_memberships.CreateGetportfoliomembershipsTool(cfg),
		tools_portfolio_memberships.CreateGetportfoliomembershipTool(cfg),
		tools_goals.CreateGetgoalsTool(cfg),
		tools_goals.CreateCreategoalTool(cfg),
		tools_team_memberships.CreateGetteammembershipsforteamTool(cfg),
		tools_custom_field_settings.CreateGetcustomfieldsettingsforprojectTool(cfg),
		tools_tasks.CreateRemovetagfortaskTool(cfg),
		tools_tags.CreateGettagsTool(cfg),
		tools_tags.CreateCreatetagTool(cfg),
		tools_stories.CreateUpdatestoryTool(cfg),
		tools_stories.CreateDeletestoryTool(cfg),
		tools_stories.CreateGetstoryTool(cfg),
		tools_project_templates.CreateGetprojecttemplatesforteamTool(cfg),
		tools_users.CreateGetuserTool(cfg),
		tools_projects.CreateAddfollowersforprojectTool(cfg),
		tools_webhooks.CreateGetwebhooksTool(cfg),
		tools_webhooks.CreateCreatewebhookTool(cfg),
		tools_sections.CreateDeletesectionTool(cfg),
		tools_sections.CreateGetsectionTool(cfg),
		tools_sections.CreateUpdatesectionTool(cfg),
		tools_teams.CreateGetteamsforworkspaceTool(cfg),
		tools_portfolios.CreateAddcustomfieldsettingforportfolioTool(cfg),
		tools_projects.CreateDeleteprojectTool(cfg),
		tools_projects.CreateGetprojectTool(cfg),
		tools_projects.CreateUpdateprojectTool(cfg),
		tools_organization_exports.CreateCreateorganizationexportTool(cfg),
		tools_portfolios.CreateGetportfoliosTool(cfg),
		tools_portfolios.CreateCreateportfolioTool(cfg),
		tools_user_task_lists.CreateGetusertasklistforuserTool(cfg),
		tools_projects.CreateGettaskcountsforprojectTool(cfg),
		tools_status_updates.CreateGetstatusTool(cfg),
		tools_status_updates.CreateDeletestatusTool(cfg),
		tools_portfolio_memberships.CreateGetportfoliomembershipsforportfolioTool(cfg),
		tools_attachments.CreateDeleteattachmentTool(cfg),
		tools_attachments.CreateGetattachmentTool(cfg),
		tools_tasks.CreateDuplicatetaskTool(cfg),
		tools_audit_log_api.CreateGetauditlogeventsTool(cfg),
		tools_tasks.CreateGetdependentsfortaskTool(cfg),
		tools_goals.CreateGetparentgoalsforgoalTool(cfg),
		tools_goal_relationships.CreateGetgoalrelationshipTool(cfg),
		tools_goal_relationships.CreateUpdategoalrelationshipTool(cfg),
		tools_tasks.CreateAdddependentsfortaskTool(cfg),
		tools_batch_api.CreateCreatebatchrequestTool(cfg),
		tools_workspace_memberships.CreateGetworkspacemembershipsforworkspaceTool(cfg),
		tools_webhooks.CreateDeletewebhookTool(cfg),
		tools_webhooks.CreateGetwebhookTool(cfg),
		tools_webhooks.CreateUpdatewebhookTool(cfg),
		tools_custom_fields.CreateGetcustomfieldsforworkspaceTool(cfg),
		tools_projects.CreateGetprojectsfortaskTool(cfg),
		tools_portfolios.CreateGetitemsforportfolioTool(cfg),
		tools_project_statuses.CreateDeleteprojectstatusTool(cfg),
		tools_project_statuses.CreateGetprojectstatusTool(cfg),
		tools_portfolios.CreateDeleteportfolioTool(cfg),
		tools_portfolios.CreateGetportfolioTool(cfg),
		tools_portfolios.CreateUpdateportfolioTool(cfg),
		tools_tasks.CreateAdddependenciesfortaskTool(cfg),
		tools_custom_fields.CreateUpdateenumoptionTool(cfg),
		tools_projects.CreateProjectsaveastemplateTool(cfg),
		tools_tasks.CreateGettasksTool(cfg),
		tools_tasks.CreateCreatetaskTool(cfg),
		tools_tasks.CreateRemovefollowerfortaskTool(cfg),
		tools_tasks.CreateGetsubtasksfortaskTool(cfg),
		tools_tasks.CreateCreatesubtaskfortaskTool(cfg),
		tools_workspaces.CreateRemoveuserforworkspaceTool(cfg),
		tools_tasks.CreateSearchtasksforworkspaceTool(cfg),
		tools_projects.CreateRemovemembersforprojectTool(cfg),
		tools_jobs.CreateGetjobTool(cfg),
		tools_goals.CreateUpdategoalmetricTool(cfg),
		tools_project_briefs.CreateDeleteprojectbriefTool(cfg),
		tools_project_briefs.CreateGetprojectbriefTool(cfg),
		tools_project_briefs.CreateUpdateprojectbriefTool(cfg),
		tools_workspace_memberships.CreateGetworkspacemembershipTool(cfg),
		tools_goals.CreateAddfollowersTool(cfg),
		tools_portfolios.CreateRemoveitemforportfolioTool(cfg),
		tools_tasks.CreateGettasksfortagTool(cfg),
		tools_tasks.CreateAddprojectfortaskTool(cfg),
		tools_goals.CreateGetgoalTool(cfg),
		tools_goals.CreateUpdategoalTool(cfg),
		tools_goals.CreateDeletegoalTool(cfg),
		tools_time_periods.CreateGettimeperiodTool(cfg),
		tools_users.CreateGetusersforworkspaceTool(cfg),
		tools_project_memberships.CreateGetprojectmembershipsforprojectTool(cfg),
		tools_portfolios.CreateRemovemembersforportfolioTool(cfg),
		tools_projects.CreateCreateprojectforteamTool(cfg),
		tools_projects.CreateGetprojectsforteamTool(cfg),
		tools_goal_relationships.CreateGetgoalrelationshipsTool(cfg),
		tools_tasks.CreateRemovedependentsfortaskTool(cfg),
		tools_teams.CreateAdduserforteamTool(cfg),
		tools_users.CreateGetusersforteamTool(cfg),
		tools_users.CreateGetusersTool(cfg),
		tools_team_memberships.CreateGetteammembershipsTool(cfg),
		tools_tasks.CreateGetdependenciesfortaskTool(cfg),
		tools_project_statuses.CreateGetprojectstatusesforprojectTool(cfg),
		tools_project_statuses.CreateCreateprojectstatusforprojectTool(cfg),
		tools_goal_relationships.CreateRemovesupportingrelationshipTool(cfg),
		tools_projects.CreateRemovecustomfieldsettingforprojectTool(cfg),
		tools_typeahead.CreateTypeaheadforworkspaceTool(cfg),
		tools_goals.CreateCreategoalmetricTool(cfg),
		tools_projects.CreateAddcustomfieldsettingforprojectTool(cfg),
		tools_status_updates.CreateGetstatusesforobjectTool(cfg),
		tools_status_updates.CreateCreatestatusforobjectTool(cfg),
		tools_workspaces.CreateGetworkspacesTool(cfg),
		tools_teams.CreateCreateteamTool(cfg),
		tools_teams.CreateUpdateteamTool(cfg),
		tools_user_task_lists.CreateGetusertasklistTool(cfg),
		tools_attachments.CreateGetattachmentsforobjectTool(cfg),
		tools_tasks.CreateUpdatetaskTool(cfg),
		tools_tasks.CreateDeletetaskTool(cfg),
		tools_tasks.CreateGettaskTool(cfg),
		tools_tags.CreateGettagsfortaskTool(cfg),
		tools_project_templates.CreateGetprojecttemplatesTool(cfg),
		tools_teams.CreateGetteamTool(cfg),
		tools_portfolios.CreateAddmembersforportfolioTool(cfg),
		tools_sections.CreateGetsectionsforprojectTool(cfg),
		tools_sections.CreateCreatesectionforprojectTool(cfg),
		tools_organization_exports.CreateGetorganizationexportTool(cfg),
		tools_tasks.CreateSetparentfortaskTool(cfg),
		tools_stories.CreateGetstoriesfortaskTool(cfg),
		tools_stories.CreateCreatestoryfortaskTool(cfg),
		tools_projects.CreateGetprojectsTool(cfg),
		tools_projects.CreateCreateprojectTool(cfg),
	}
}
