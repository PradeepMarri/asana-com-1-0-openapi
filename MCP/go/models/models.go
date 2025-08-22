package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CustomFieldRequest represents the CustomFieldRequest schema from the OpenAPI specification
type CustomFieldRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the custom field.
	Display_value string `json:"display_value,omitempty"` // A string representation for the value of the custom field. Integrations that don't require the underlying type should use this field to read values. Using this field will future-proof an app against new custom field types.
	Enum_value interface{} `json:"enum_value,omitempty"`
	Date_value map[string]interface{} `json:"date_value,omitempty"` // *Conditional*. Only relevant for custom fields of type `date`. This object reflects the chosen date (and optionally, time) value of a `date` custom field. If no date is selected, the value of `date_value` will be `null`.
	Enum_options []EnumOption `json:"enum_options,omitempty"` // *Conditional*. Only relevant for custom fields of type `enum`. This array specifies the possible values which an `enum` custom field can adopt. To modify the enum options, refer to [working with enum options](/docs/create-an-enum-option).
	Number_value float64 `json:"number_value,omitempty"` // *Conditional*. This number is the value of a `number` custom field.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The type of the custom field. Must be one of the given values.
	Text_value string `json:"text_value,omitempty"` // *Conditional*. This string is the value of a `text` custom field.
	TypeField string `json:"type,omitempty"` // *Deprecated: new integrations should prefer the resource_subtype field.* The type of the custom field. Must be one of the given values.
	Enabled bool `json:"enabled,omitempty"` // *Conditional*. Determines if the custom field is enabled or not.
	Multi_enum_values []EnumOption `json:"multi_enum_values,omitempty"` // *Conditional*. Only relevant for custom fields of type `multi_enum`. This object is the chosen values of a `multi_enum` custom field.
	Custom_label_position string `json:"custom_label_position,omitempty"` // Only relevant for custom fields with `custom` format. This depicts where to place the custom label. This will be null if the `format` is not `custom`.
	Enum_options []EnumOption `json:"enum_options,omitempty"` // *Conditional*. Only relevant for custom fields of type `enum`. This array specifies the possible values which an `enum` custom field can adopt. To modify the enum options, refer to [working with enum options](/docs/create-an-enum-option).
	Is_global_to_workspace bool `json:"is_global_to_workspace,omitempty"` // This flag describes whether this custom field is available to every container in the workspace. Before project-specific custom fields, this field was always true.
	Precision int `json:"precision,omitempty"` // Only relevant for custom fields of type ‘Number’. This field dictates the number of places after the decimal to round to, i.e. 0 is integer values, 1 rounds to the nearest tenth, and so on. Must be between 0 and 6, inclusive. For percentage format, this may be unintuitive, as a value of 0.25 has a precision of 0, while a value of 0.251 has a precision of 1. This is due to 0.25 being displayed as 25%. The identifier format will always have a precision of 0.
	Currency_code string `json:"currency_code,omitempty"` // ISO 4217 currency code to format this custom field. This will be null if the `format` is not `currency`.
	Format string `json:"format,omitempty"` // The format of this custom field.
	Asana_created_field string `json:"asana_created_field,omitempty"` // *Conditional*. A unique identifier to associate this field with the template source of truth.
	Custom_label string `json:"custom_label,omitempty"` // This is the string that appears next to the custom field value. This will be null if the `format` is not `custom`.
	Description string `json:"description,omitempty"` // [Opt In](/docs/input-output-options). The description of the custom field.
	Has_notifications_enabled bool `json:"has_notifications_enabled,omitempty"` // *Conditional*. This flag describes whether a follower of a task with this field should receive inbox notifications from changes to this field.
	People_value []string `json:"people_value,omitempty"` // *Conditional*. Only relevant for custom fields of type `people`. This array of user GIDs reflects the users to be written to a `people` custom field. Note that *write* operations will replace existing users (if any) in the custom field with the users specified in this array.
	Workspace string `json:"workspace"` // *Create-Only* The workspace to create a custom field in.
}

// SectionTaskInsertRequest represents the SectionTaskInsertRequest schema from the OpenAPI specification
type SectionTaskInsertRequest struct {
	Task string `json:"task"` // The task to add to this section.
	Insert_after string `json:"insert_after,omitempty"` // An existing task within this section after which the added task should be inserted. Cannot be provided together with insert_before.
	Insert_before string `json:"insert_before,omitempty"` // An existing task within this section before which the added task should be inserted. Cannot be provided together with insert_after.
}

// WorkspaceRemoveUserRequest represents the WorkspaceRemoveUserRequest schema from the OpenAPI specification
type WorkspaceRemoveUserRequest struct {
	User string `json:"user,omitempty"` // A string identifying a user. This can either be the string "me", an email, or the gid of a user.
}

// TaskCompact represents the TaskCompact schema from the OpenAPI specification
type TaskCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the task.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The resource_subtype `milestone` represent a single moment in time. This means tasks with this subtype cannot have a start_date.
}

// PortfolioMembershipCompact represents the PortfolioMembershipCompact schema from the OpenAPI specification
type PortfolioMembershipCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Portfolio PortfolioCompact `json:"portfolio,omitempty"`
	User UserCompact `json:"user,omitempty"`
}

// JobCompact represents the JobCompact schema from the OpenAPI specification
type JobCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	New_project_template ProjectTemplateCompact `json:"new_project_template,omitempty"`
	New_task TaskCompact `json:"new_task,omitempty"`
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Status string `json:"status,omitempty"` // The current status of this job. The value is one of: `not_started`, `in_progress`, `succeeded`, or `failed`.
	New_project ProjectCompact `json:"new_project,omitempty"`
}

// StoryBase represents the StoryBase schema from the OpenAPI specification
type StoryBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Sticker_name string `json:"sticker_name,omitempty"` // The name of the sticker in this story. `null` if there is no sticker.
	Text string `json:"text,omitempty"` // The plain text of the comment to add. Cannot be used with html_text.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). HTML formatted text for a comment. This will not include the name of the creator.
	Is_pinned bool `json:"is_pinned,omitempty"` // *Conditional*. Whether the story should be pinned on the resource.
}

// PortfolioRequest represents the PortfolioRequest schema from the OpenAPI specification
type PortfolioRequest struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Name string `json:"name,omitempty"` // The name of the portfolio.
	Color string `json:"color,omitempty"` // Color of the portfolio.
	Members []string `json:"members,omitempty"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
	Public bool `json:"public,omitempty"` // True if the portfolio is public to its workspace members.
	Workspace string `json:"workspace,omitempty"` // Gid of an object.
}

// TaskSetParentRequest represents the TaskSetParentRequest schema from the OpenAPI specification
type TaskSetParentRequest struct {
	Insert_after string `json:"insert_after,omitempty"` // A subtask of the parent to insert the task after, or `null` to insert at the beginning of the list.
	Insert_before string `json:"insert_before,omitempty"` // A subtask of the parent to insert the task before, or `null` to insert at the end of the list.
	Parent string `json:"parent"` // The new parent of the task, or `null` for no parent.
}

// StatusUpdateRequest represents the StatusUpdateRequest schema from the OpenAPI specification
type StatusUpdateRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The `resource_subtype`s for `status` objects represent the type of their parent.
	Title string `json:"title,omitempty"` // The title of the status update.
	Status_type string `json:"status_type"` // The type associated with the status update. This represents the current state of the object this object is on.
	Text string `json:"text"` // The text content of the status update.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). The text content of the status update with formatting as HTML.
	Parent interface{} `json:"parent"`
}

// ProjectSaveAsTemplateRequest represents the ProjectSaveAsTemplateRequest schema from the OpenAPI specification
type ProjectSaveAsTemplateRequest struct {
	Name string `json:"name"` // The name of the new project template.
	Public bool `json:"public"` // Sets the project template to public to its team.
	Team string `json:"team,omitempty"` // Sets the team of the new project template. If the project exists in an organization, specify team and not workspace.
	Workspace string `json:"workspace,omitempty"` // Sets the workspace of the new project template. Only specify workspace if the project exists in a workspace.
}

// AuditLogEventContext represents the AuditLogEventContext schema from the OpenAPI specification
type AuditLogEventContext struct {
	Oauth_app_name string `json:"oauth_app_name,omitempty"` // The name of the OAuth App that initiated the event. Only present if the `api_authentication_method` is `oauth`.
	User_agent string `json:"user_agent,omitempty"` // The user agent of the client that initiated the event, if applicable.
	Api_authentication_method string `json:"api_authentication_method,omitempty"` // The authentication method used in the context of an API request. Only present if the `context_type` is `api`. Can be one of `cookie`, `oauth`, `personal_access_token`, or `service_account`.
	Client_ip_address string `json:"client_ip_address,omitempty"` // The IP address of the client that initiated the event, if applicable.
	Context_type string `json:"context_type,omitempty"` // The type of context. Can be one of `web`, `desktop`, `mobile`, `asana_support`, `asana`, `email`, or `api`.
}

// ProjectBriefRequest represents the ProjectBriefRequest schema from the OpenAPI specification
type ProjectBriefRequest struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Html_text string `json:"html_text,omitempty"` // HTML formatted text for the project brief.
	Title string `json:"title,omitempty"` // The title of the project brief.
	Text string `json:"text,omitempty"` // The plain text of the project brief. When writing to a project brief, you can specify either `html_text` (preferred) or `text`, but not both.
}

// EmptyResponse represents the EmptyResponse schema from the OpenAPI specification
type EmptyResponse struct {
}

// TaskBase represents the TaskBase schema from the OpenAPI specification
type TaskBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the task.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The resource_subtype `milestone` represent a single moment in time. This means tasks with this subtype cannot have a start_date.
	External map[string]interface{} `json:"external,omitempty"` // *OAuth Required*. *Conditional*. This field is returned only if external values are set or included by using [Opt In] (/docs/input-output-options). The external field allows you to store app-specific metadata on tasks, including a gid that can be used to retrieve tasks and a data blob that can store app-specific character strings. Note that you will need to authenticate with Oauth to access or modify this data. Once an external gid is set, you can use the notation `external:custom_gid` to reference your object anywhere in the API where you may use the original object gid. See the page on Custom External Data for more details.
	Due_at string `json:"due_at,omitempty"` // The UTC date and time on which this task is due, or null if the task has no due time. This takes an ISO 8601 date string in UTC and should not be used together with `due_on`.
	Name string `json:"name,omitempty"` // Name of the task. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Actual_time_minutes float64 `json:"actual_time_minutes,omitempty"` // This value represents the sum of all the Time Tracking entries in the Actual Time field on a given Task. It is represented as a nullable long value.
	Dependents []AsanaResource `json:"dependents,omitempty"` // [Opt In](/docs/input-output-options). Array of resources referencing tasks that depend on this task. The objects contain only the ID of the dependent.
	Start_on string `json:"start_on,omitempty"` // The day on which work begins for the task , or null if the task has no start date. This takes a date with `YYYY-MM-DD` format and should not be used together with `start_at`. *Note: `due_on` or `due_at` must be present in the request when setting or unsetting the `start_on` parameter.*
	Liked bool `json:"liked,omitempty"` // True if the task is liked by the authorized user, false if not.
	Num_likes int `json:"num_likes,omitempty"` // The number of users who have liked this task.
	Completed_at string `json:"completed_at,omitempty"` // The time at which this task was completed, or null if the task is incomplete.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the task (i.e. its description).
	Assignee_status string `json:"assignee_status,omitempty"` // *Deprecated* Scheduling status of this task for the user it is assigned to. This field can only be set if the assignee is non-null. Setting this field to "inbox" or "upcoming" inserts it at the top of the section, while the other options will insert at the bottom.
	Html_notes string `json:"html_notes,omitempty"` // [Opt In](/docs/input-output-options). The notes of the text with formatting as HTML.
	Likes []Like `json:"likes,omitempty"` // Array of likes for users who have liked this task.
	Dependencies []AsanaResource `json:"dependencies,omitempty"` // [Opt In](/docs/input-output-options). Array of resources referencing tasks that this task depends on. The objects contain only the gid of the dependency.
	Hearted bool `json:"hearted,omitempty"` // *Deprecated - please use liked instead* True if the task is hearted by the authorized user, false if not.
	Is_rendered_as_separator bool `json:"is_rendered_as_separator,omitempty"` // [Opt In](/docs/input-output-options). In some contexts tasks can be rendered as a visual separator; for instance, subtasks can appear similar to [sections](/docs/asana-sections) without being true `section` objects. If a `task` object is rendered this way in any context it will have the property `is_rendered_as_separator` set to `true`.
	Due_on string `json:"due_on,omitempty"` // The localized date on which this task is due, or null if the task has no due date. This takes a date with `YYYY-MM-DD` format and should not be used together with `due_at`.
	Approval_status string `json:"approval_status,omitempty"` // *Conditional* Reflects the approval status of this task. This field is kept in sync with `completed`, meaning `pending` translates to false while `approved`, `rejected`, and `changes_requested` translate to true. If you set completed to true, this field will be set to `approved`.
	Completed_by UserCompact `json:"completed_by,omitempty"`
	Completed bool `json:"completed,omitempty"` // True if the task is currently marked complete, false if not.
	Memberships []map[string]interface{} `json:"memberships,omitempty"` // *Create-only*. Array of projects this task is associated with and the section it is in. At task creation time, this array can be used to add the task to specific sections. After task creation, these associations can be modified using the `addProject` and `removeProject` endpoints. Note that over time, more types of memberships may be added to this property.
	Num_hearts int `json:"num_hearts,omitempty"` // *Deprecated - please use likes instead* The number of users who have hearted this task.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Num_subtasks int `json:"num_subtasks,omitempty"` // [Opt In](/docs/input-output-options). The number of subtasks on this task.
	Hearts []Like `json:"hearts,omitempty"` // *Deprecated - please use likes instead* Array of likes for users who have hearted this task.
	Modified_at string `json:"modified_at,omitempty"` // The time at which this task was last modified. *Note: This does not currently reflect any changes in associations such as projects or comments that may have been added or removed from the task.*
	Start_at string `json:"start_at,omitempty"` // Date and time on which work begins for the task, or null if the task has no start time. This takes an ISO 8601 date string in UTC and should not be used together with `start_on`. *Note: `due_at` must be present in the request when setting or unsetting the `start_at` parameter.*
}

// ProjectCompact represents the ProjectCompact schema from the OpenAPI specification
type ProjectCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the project. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
}

// AddFollowersRequest represents the AddFollowersRequest schema from the OpenAPI specification
type AddFollowersRequest struct {
	Followers string `json:"followers"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
}

// SectionCompact represents the SectionCompact schema from the OpenAPI specification
type SectionCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the section (i.e. the text displayed as the section header).
}

// EnumOptionBase represents the EnumOptionBase schema from the OpenAPI specification
type EnumOptionBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the enum option.
	Color string `json:"color,omitempty"` // The color of the enum option. Defaults to ‘none’.
	Enabled bool `json:"enabled,omitempty"` // Whether or not the enum option is a selectable value for the custom field.
}

// TimePeriodBase represents the TimePeriodBase schema from the OpenAPI specification
type TimePeriodBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Period string `json:"period,omitempty"` // The cadence and index of the time period. The value is one of: `FY`, `H1`, `H2`, `Q1`, `Q2`, `Q3`, or `Q4`.
	Start_on string `json:"start_on,omitempty"` // The localized start date of the time period in `YYYY-MM-DD` format.
	Display_name string `json:"display_name,omitempty"` // A string representing the cadence code and the fiscal year.
	End_on string `json:"end_on,omitempty"` // The localized end date of the time period in `YYYY-MM-DD` format.
	Parent TimePeriodCompact `json:"parent,omitempty"`
}

// TagResponse represents the TagResponse schema from the OpenAPI specification
type TagResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the tag. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Color string `json:"color,omitempty"` // Color of the tag.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the tag (i.e. its description).
	Workspace WorkspaceCompact `json:"workspace,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Followers []UserCompact `json:"followers,omitempty"` // Array of users following this tag.
	Permalink_url string `json:"permalink_url,omitempty"` // A url that points directly to the object within Asana.
}

// UserTaskListResponse represents the UserTaskListResponse schema from the OpenAPI specification
type UserTaskListResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the user task list.
	Owner interface{} `json:"owner,omitempty"` // The owner of the user task list, i.e. the person whose My Tasks is represented by this resource.
	Workspace interface{} `json:"workspace,omitempty"` // The workspace in which the user task list is located.
}

// OrganizationExportRequest represents the OrganizationExportRequest schema from the OpenAPI specification
type OrganizationExportRequest struct {
	Organization string `json:"organization,omitempty"` // Globally unique identifier for the workspace or organization.
}

// GoalRelationshipResponse represents the GoalRelationshipResponse schema from the OpenAPI specification
type GoalRelationshipResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Supporting_resource interface{} `json:"supporting_resource,omitempty"`
	Contribution_weight float64 `json:"contribution_weight,omitempty"` // The weight that the supporting resource's progress contributes to the supported goal's progress. This can only be 0 or 1.
	Supported_goal interface{} `json:"supported_goal,omitempty"`
}

// StatusUpdateResponse represents the StatusUpdateResponse schema from the OpenAPI specification
type StatusUpdateResponse struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The `resource_subtype`s for `status` objects represent the type of their parent.
	Title string `json:"title,omitempty"` // The title of the status update.
	Text string `json:"text"` // The text content of the status update.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). The text content of the status update with formatting as HTML.
	Status_type string `json:"status_type"` // The type associated with the status update. This represents the current state of the object this object is on.
	Modified_at string `json:"modified_at,omitempty"` // The time at which this project status was last modified. *Note: This does not currently reflect any changes in associations such as comments that may have been added or removed from the status.*
	Num_hearts int `json:"num_hearts,omitempty"` // *Deprecated - please use likes instead* The number of users who have hearted this status.
	Num_likes int `json:"num_likes,omitempty"` // The number of users who have liked this status.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Hearts []Like `json:"hearts,omitempty"` // *Deprecated - please use likes instead* Array of likes for users who have hearted this status.
	Liked bool `json:"liked,omitempty"` // True if the status is liked by the authorized user, false if not.
	Likes []Like `json:"likes,omitempty"` // Array of likes for users who have liked this status.
	Created_by UserCompact `json:"created_by,omitempty"`
	Parent interface{} `json:"parent,omitempty"`
	Author UserCompact `json:"author,omitempty"`
	Hearted bool `json:"hearted,omitempty"` // *Deprecated - please use liked instead* True if the status is hearted by the authorized user, false if not.
}

// OrganizationExportBase represents the OrganizationExportBase schema from the OpenAPI specification
type OrganizationExportBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Organization WorkspaceCompact `json:"organization,omitempty"`
	State string `json:"state,omitempty"` // The current state of the export.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Download_url string `json:"download_url,omitempty"` // Download this URL to retreive the full export of the organization in JSON format. It will be compressed in a gzip (.gz) container. *Note: May be null if the export is still in progress or failed. If present, this URL may only be valid for 1 hour from the time of retrieval. You should avoid persisting this URL somewhere and rather refresh on demand to ensure you do not keep stale URLs.*
}

// CustomFieldBase represents the CustomFieldBase schema from the OpenAPI specification
type CustomFieldBase struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The type of the custom field. Must be one of the given values.
	Text_value string `json:"text_value,omitempty"` // *Conditional*. This string is the value of a `text` custom field.
	TypeField string `json:"type,omitempty"` // *Deprecated: new integrations should prefer the resource_subtype field.* The type of the custom field. Must be one of the given values.
	Enabled bool `json:"enabled,omitempty"` // *Conditional*. Determines if the custom field is enabled or not.
	Multi_enum_values []EnumOption `json:"multi_enum_values,omitempty"` // *Conditional*. Only relevant for custom fields of type `multi_enum`. This object is the chosen values of a `multi_enum` custom field.
	Name string `json:"name,omitempty"` // The name of the custom field.
	Display_value string `json:"display_value,omitempty"` // A string representation for the value of the custom field. Integrations that don't require the underlying type should use this field to read values. Using this field will future-proof an app against new custom field types.
	Enum_value interface{} `json:"enum_value,omitempty"`
	Date_value map[string]interface{} `json:"date_value,omitempty"` // *Conditional*. Only relevant for custom fields of type `date`. This object reflects the chosen date (and optionally, time) value of a `date` custom field. If no date is selected, the value of `date_value` will be `null`.
	Enum_options []EnumOption `json:"enum_options,omitempty"` // *Conditional*. Only relevant for custom fields of type `enum`. This array specifies the possible values which an `enum` custom field can adopt. To modify the enum options, refer to [working with enum options](/docs/create-an-enum-option).
	Number_value float64 `json:"number_value,omitempty"` // *Conditional*. This number is the value of a `number` custom field.
	Asana_created_field string `json:"asana_created_field,omitempty"` // *Conditional*. A unique identifier to associate this field with the template source of truth.
	Custom_label string `json:"custom_label,omitempty"` // This is the string that appears next to the custom field value. This will be null if the `format` is not `custom`.
	Description string `json:"description,omitempty"` // [Opt In](/docs/input-output-options). The description of the custom field.
	Has_notifications_enabled bool `json:"has_notifications_enabled,omitempty"` // *Conditional*. This flag describes whether a follower of a task with this field should receive inbox notifications from changes to this field.
	Custom_label_position string `json:"custom_label_position,omitempty"` // Only relevant for custom fields with `custom` format. This depicts where to place the custom label. This will be null if the `format` is not `custom`.
	Enum_options []EnumOption `json:"enum_options,omitempty"` // *Conditional*. Only relevant for custom fields of type `enum`. This array specifies the possible values which an `enum` custom field can adopt. To modify the enum options, refer to [working with enum options](/docs/create-an-enum-option).
	Is_global_to_workspace bool `json:"is_global_to_workspace,omitempty"` // This flag describes whether this custom field is available to every container in the workspace. Before project-specific custom fields, this field was always true.
	Precision int `json:"precision,omitempty"` // Only relevant for custom fields of type ‘Number’. This field dictates the number of places after the decimal to round to, i.e. 0 is integer values, 1 rounds to the nearest tenth, and so on. Must be between 0 and 6, inclusive. For percentage format, this may be unintuitive, as a value of 0.25 has a precision of 0, while a value of 0.251 has a precision of 1. This is due to 0.25 being displayed as 25%. The identifier format will always have a precision of 0.
	Currency_code string `json:"currency_code,omitempty"` // ISO 4217 currency code to format this custom field. This will be null if the `format` is not `currency`.
	Format string `json:"format,omitempty"` // The format of this custom field.
}

// OrganizationExportCompact represents the OrganizationExportCompact schema from the OpenAPI specification
type OrganizationExportCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Download_url string `json:"download_url,omitempty"` // Download this URL to retreive the full export of the organization in JSON format. It will be compressed in a gzip (.gz) container. *Note: May be null if the export is still in progress or failed. If present, this URL may only be valid for 1 hour from the time of retrieval. You should avoid persisting this URL somewhere and rather refresh on demand to ensure you do not keep stale URLs.*
	Organization WorkspaceCompact `json:"organization,omitempty"`
	State string `json:"state,omitempty"` // The current state of the export.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
}

// UserRequest represents the UserRequest schema from the OpenAPI specification
type UserRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // *Read-only except when same user as requester*. The user’s name.
}

// ProjectResponse represents the ProjectResponse schema from the OpenAPI specification
type ProjectResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the project. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Custom_field_settings []CustomFieldSettingResponse `json:"custom_field_settings,omitempty"` // Array of Custom Field Settings (in compact form).
	Html_notes string `json:"html_notes,omitempty"` // [Opt In](/docs/input-output-options). The notes of the project with formatting as HTML.
	Public bool `json:"public,omitempty"` // True if the project is public to its team.
	Modified_at string `json:"modified_at,omitempty"` // The time at which this project was last modified. *Note: This does not currently reflect any changes in associations such as tasks or comments that may have been added or removed from the project.*
	Archived bool `json:"archived,omitempty"` // True if the project is archived, false if not. Archived projects do not show in the UI by default and may be treated differently for queries.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Due_date string `json:"due_date,omitempty"` // *Deprecated: new integrations should prefer the `due_on` field.*
	Is_template bool `json:"is_template,omitempty"` // [Opt In](/docs/input-output-options). *Deprecated - please use a project template endpoint instead (more in [this forum post](https://forum.asana.com/t/a-new-api-for-project-templates/156432)).* Determines if the project is a template.
	Color string `json:"color,omitempty"` // Color of the project.
	Due_on string `json:"due_on,omitempty"` // The day on which this project is due. This takes a date with format YYYY-MM-DD.
	Current_status interface{} `json:"current_status,omitempty"` // *Deprecated: new integrations should prefer the `current_status_update` resource.*
	Current_status_update interface{} `json:"current_status_update,omitempty"` // The latest `status_update` posted to this project.
	Members []UserCompact `json:"members,omitempty"` // Array of users who are members of this project.
	Start_on string `json:"start_on,omitempty"` // The day on which work for this project begins, or null if the project has no start date. This takes a date with `YYYY-MM-DD` format. *Note: `due_on` or `due_at` must be present in the request when setting or unsetting the `start_on` parameter. Additionally, `start_on` and `due_on` cannot be the same date.*
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the project (ie., its description).
	Workspace interface{} `json:"workspace,omitempty"`
	Default_view string `json:"default_view,omitempty"` // The default view (list, board, calendar, or timeline) of a project.
	Team interface{} `json:"team,omitempty"`
	Completed bool `json:"completed,omitempty"` // True if the project is currently marked complete, false if not.
	Completed_at string `json:"completed_at,omitempty"` // The time at which this project was completed, or null if the project is not completed.
	Created_from_template interface{} `json:"created_from_template,omitempty"`
	Owner interface{} `json:"owner,omitempty"` // The current owner of the project, may be null.
	Permalink_url string `json:"permalink_url,omitempty"` // A url that points directly to the object within Asana.
	Project_brief interface{} `json:"project_brief,omitempty"`
	Followers []UserCompact `json:"followers,omitempty"` // Array of users following this project. Followers are a subset of members who have opted in to receive "tasks added" notifications for a project.
	Icon string `json:"icon,omitempty"` // The icon for a project.
	Completed_by UserCompact `json:"completed_by,omitempty"`
	Custom_fields []CustomFieldCompact `json:"custom_fields,omitempty"` // Array of Custom Fields.
}

// GoalRelationshipRequest represents the GoalRelationshipRequest schema from the OpenAPI specification
type GoalRelationshipRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Contribution_weight float64 `json:"contribution_weight,omitempty"` // The weight that the supporting resource's progress contributes to the supported goal's progress. This can only be 0 or 1.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Supporting_resource interface{} `json:"supporting_resource,omitempty"`
	Supported_goal interface{} `json:"supported_goal,omitempty"`
}

// AuditLogEventActor represents the AuditLogEventActor schema from the OpenAPI specification
type AuditLogEventActor struct {
	Actor_type string `json:"actor_type,omitempty"` // The type of actor. Can be one of `user`, `asana`, `asana_support`, `anonymous`, or `external_administrator`.
	Email string `json:"email,omitempty"` // The email of the actor, if it is a user.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the actor, if it is a user.
	Name string `json:"name,omitempty"` // The name of the actor, if it is a user.
}

// TeamRemoveUserRequest represents the TeamRemoveUserRequest schema from the OpenAPI specification
type TeamRemoveUserRequest struct {
	User string `json:"user,omitempty"` // A string identifying a user. This can either be the string "me", an email, or the gid of a user.
}

// WebhookRequest represents the WebhookRequest schema from the OpenAPI specification
type WebhookRequest struct {
	Filters []interface{} `json:"filters,omitempty"` // An array of WebhookFilter objects to specify a whitelist of filters to apply to events from this webhook. If a webhook event passes any of the filters the event will be delivered; otherwise no event will be sent to the receiving server.
	Resource string `json:"resource"` // A resource ID to subscribe to. Many Asana resources are valid to create webhooks on, but higher-level resources require filters.
	Target string `json:"target"` // The URL to receive the HTTP POST. The full URL will be used to deliver events from this webhook (including parameters) which allows encoding of application-specific state when the webhook is created.
}

// PortfolioMembershipBase represents the PortfolioMembershipBase schema from the OpenAPI specification
type PortfolioMembershipBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Portfolio PortfolioCompact `json:"portfolio,omitempty"`
	User UserCompact `json:"user,omitempty"`
}

// PortfolioMembershipResponse represents the PortfolioMembershipResponse schema from the OpenAPI specification
type PortfolioMembershipResponse struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Portfolio PortfolioCompact `json:"portfolio,omitempty"`
	User UserCompact `json:"user,omitempty"`
}

// WorkspaceBase represents the WorkspaceBase schema from the OpenAPI specification
type WorkspaceBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the workspace.
}

// ProjectStatusCompact represents the ProjectStatusCompact schema from the OpenAPI specification
type ProjectStatusCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Title string `json:"title,omitempty"` // The title of the project status update.
}

// GoalAddSupportingWorkRequest represents the GoalAddSupportingWorkRequest schema from the OpenAPI specification
type GoalAddSupportingWorkRequest struct {
	Supporting_work string `json:"supporting_work"` // The project/portfolio gid to add as supporting work for a goal
}

// WorkspaceRequest represents the WorkspaceRequest schema from the OpenAPI specification
type WorkspaceRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the workspace.
}

// TaskRemoveProjectRequest represents the TaskRemoveProjectRequest schema from the OpenAPI specification
type TaskRemoveProjectRequest struct {
	Project string `json:"project"` // The project to remove the task from.
}

// PortfolioBase represents the PortfolioBase schema from the OpenAPI specification
type PortfolioBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the portfolio.
	Color string `json:"color,omitempty"` // Color of the portfolio.
}

// ProjectTemplateBase represents the ProjectTemplateBase schema from the OpenAPI specification
type ProjectTemplateBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the project template.
	Html_description string `json:"html_description,omitempty"` // The description of the project template with formatting as HTML.
	Owner interface{} `json:"owner,omitempty"` // The current owner of the project template, may be null.
	Public bool `json:"public,omitempty"` // True if the project template is public to its team.
	Requested_dates []DateVariableCompact `json:"requested_dates,omitempty"` // Array of date variables in this project template. Calendar dates must be provided for these variables when instantiating a project.
	Team interface{} `json:"team,omitempty"`
	Color string `json:"color,omitempty"` // Color of the project template.
	Description string `json:"description,omitempty"` // Free-form textual information associated with the project template
}

// TagRequest represents the TagRequest schema from the OpenAPI specification
type TagRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the tag. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Color string `json:"color,omitempty"` // Color of the tag.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the tag (i.e. its description).
	Followers []string `json:"followers,omitempty"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
	Workspace string `json:"workspace,omitempty"` // Gid of an object.
}

// TimePeriodResponse represents the TimePeriodResponse schema from the OpenAPI specification
type TimePeriodResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Period string `json:"period,omitempty"` // The cadence and index of the time period. The value is one of: `FY`, `H1`, `H2`, `Q1`, `Q2`, `Q3`, or `Q4`.
	Start_on string `json:"start_on,omitempty"` // The localized start date of the time period in `YYYY-MM-DD` format.
	Display_name string `json:"display_name,omitempty"` // A string representing the cadence code and the fiscal year.
	End_on string `json:"end_on,omitempty"` // The localized end date of the time period in `YYYY-MM-DD` format.
	Parent TimePeriodCompact `json:"parent,omitempty"`
}

// BatchRequest represents the BatchRequest schema from the OpenAPI specification
type BatchRequest struct {
	Actions []BatchRequestAction `json:"actions,omitempty"`
}

// AuditLogEvent represents the AuditLogEvent schema from the OpenAPI specification
type AuditLogEvent struct {
	Resource interface{} `json:"resource,omitempty"`
	Actor interface{} `json:"actor,omitempty"`
	Context interface{} `json:"context,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The time the event was created.
	Details interface{} `json:"details,omitempty"`
	Event_category string `json:"event_category,omitempty"` // The category that this `event_type` belongs to.
	Event_type string `json:"event_type,omitempty"` // The type of the event.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the `AuditLogEvent`, as a string.
}

// AuditLogEventResource represents the AuditLogEventResource schema from the OpenAPI specification
type AuditLogEventResource struct {
	Email string `json:"email,omitempty"` // The email of the resource, if applicable.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource.
	Name string `json:"name,omitempty"` // The name of the resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of resource. Most resources will not have a subtype.
	Resource_type string `json:"resource_type,omitempty"` // The type of resource.
}

// TaskDuplicateRequest represents the TaskDuplicateRequest schema from the OpenAPI specification
type TaskDuplicateRequest struct {
	Include string `json:"include,omitempty"` // The fields that will be duplicated to the new task.
	Name string `json:"name,omitempty"` // The name of the new task.
}

// GoalResponse represents the GoalResponse schema from the OpenAPI specification
type GoalResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Start_on string `json:"start_on,omitempty"` // The day on which work for this goal begins, or null if the goal has no start date. This takes a date with `YYYY-MM-DD` format, and cannot be set unless there is an accompanying due date.
	Status string `json:"status,omitempty"` // The current status of this goal. When the goal is open, its status can be `green`, `yellow`, and `red` to reflect "On Track", "At Risk", and "Off Track", respectively. When the goal is closed, the value can be `missed`, `achieved`, `partial`, or `dropped`. *Note* you can only write to this property if `metric` is set.
	Due_on string `json:"due_on,omitempty"` // The localized day on which this goal is due. This takes a date with format `YYYY-MM-DD`.
	Html_notes string `json:"html_notes,omitempty"` // The notes of the goal with formatting as HTML.
	Is_workspace_level bool `json:"is_workspace_level,omitempty"` // *Conditional*. This property is only present when the `workspace` provided is an organization. Whether the goal belongs to the `workspace` (and is listed as part of the workspace’s goals) or not. If it isn’t a workspace-level goal, it is a team-level goal, and is associated with the goal’s team.
	Liked bool `json:"liked,omitempty"` // True if the goal is liked by the authorized user, false if not.
	Name string `json:"name,omitempty"` // The name of the goal.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the goal (i.e. its description).
	Current_status_update interface{} `json:"current_status_update,omitempty"` // The latest `status_update` posted to this goal.
	Owner interface{} `json:"owner,omitempty"`
	Workspace interface{} `json:"workspace,omitempty"`
	Followers []UserCompact `json:"followers,omitempty"` // Array of users who are members of this goal.
	Time_period interface{} `json:"time_period,omitempty"`
	Likes []Like `json:"likes,omitempty"` // Array of likes for users who have liked this goal.
	Metric interface{} `json:"metric,omitempty"`
	Num_likes int `json:"num_likes,omitempty"` // The number of users who have liked this goal.
	Team interface{} `json:"team,omitempty"` // *Conditional*. This property is only present when the `workspace` provided is an organization.
}

// StoryResponseDates represents the StoryResponseDates schema from the OpenAPI specification
type StoryResponseDates struct {
	Due_at string `json:"due_at,omitempty"` // The UTC date and time on which this task is due, or null if the task has no due time. This takes an ISO 8601 date string in UTC and should not be used together with `due_on`.
	Due_on string `json:"due_on,omitempty"` // The localized day on which this goal is due. This takes a date with format `YYYY-MM-DD`.
	Start_on string `json:"start_on,omitempty"` // The day on which work for this goal begins, or null if the goal has no start date. This takes a date with `YYYY-MM-DD` format, and cannot be set unless there is an accompanying due date.
}

// OrganizationExportResponse represents the OrganizationExportResponse schema from the OpenAPI specification
type OrganizationExportResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Download_url string `json:"download_url,omitempty"` // Download this URL to retreive the full export of the organization in JSON format. It will be compressed in a gzip (.gz) container. *Note: May be null if the export is still in progress or failed. If present, this URL may only be valid for 1 hour from the time of retrieval. You should avoid persisting this URL somewhere and rather refresh on demand to ensure you do not keep stale URLs.*
	Organization WorkspaceCompact `json:"organization,omitempty"`
	State string `json:"state,omitempty"` // The current state of the export.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
}

// PortfolioRemoveItemRequest represents the PortfolioRemoveItemRequest schema from the OpenAPI specification
type PortfolioRemoveItemRequest struct {
	Item string `json:"item"` // The item to remove from the portfolio.
}

// CustomFieldSettingResponse represents the CustomFieldSettingResponse schema from the OpenAPI specification
type CustomFieldSettingResponse struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Project interface{} `json:"project,omitempty"`
	Custom_field interface{} `json:"custom_field,omitempty"`
	Is_important bool `json:"is_important,omitempty"` // `is_important` is used in the Asana web application to determine if this custom field is displayed in the list/grid view of a project or portfolio.
	Parent interface{} `json:"parent,omitempty"`
}

// Like represents the Like schema from the OpenAPI specification
type Like struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the object, as a string.
	User UserCompact `json:"user,omitempty"`
}

// TaskRemoveTagRequest represents the TaskRemoveTagRequest schema from the OpenAPI specification
type TaskRemoveTagRequest struct {
	Tag string `json:"tag"` // The tag to remove from the task.
}

// GoalMetricCurrentValueRequest represents the GoalMetricCurrentValueRequest schema from the OpenAPI specification
type GoalMetricCurrentValueRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Current_number_value float64 `json:"current_number_value,omitempty"` // *Conditional*. This number is the current value of a goal metric of type number.
}

// TaskRequest represents the TaskRequest schema from the OpenAPI specification
type TaskRequest struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Name string `json:"name,omitempty"` // The name of the task.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The resource_subtype `milestone` represent a single moment in time. This means tasks with this subtype cannot have a start_date.
	Due_on string `json:"due_on,omitempty"` // The localized date on which this task is due, or null if the task has no due date. This takes a date with `YYYY-MM-DD` format and should not be used together with `due_at`.
	Approval_status string `json:"approval_status,omitempty"` // *Conditional* Reflects the approval status of this task. This field is kept in sync with `completed`, meaning `pending` translates to false while `approved`, `rejected`, and `changes_requested` translate to true. If you set completed to true, this field will be set to `approved`.
	Completed_by UserCompact `json:"completed_by,omitempty"`
	Completed bool `json:"completed,omitempty"` // True if the task is currently marked complete, false if not.
	Memberships []map[string]interface{} `json:"memberships,omitempty"` // *Create-only*. Array of projects this task is associated with and the section it is in. At task creation time, this array can be used to add the task to specific sections. After task creation, these associations can be modified using the `addProject` and `removeProject` endpoints. Note that over time, more types of memberships may be added to this property.
	Num_hearts int `json:"num_hearts,omitempty"` // *Deprecated - please use likes instead* The number of users who have hearted this task.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Num_subtasks int `json:"num_subtasks,omitempty"` // [Opt In](/docs/input-output-options). The number of subtasks on this task.
	Hearts []Like `json:"hearts,omitempty"` // *Deprecated - please use likes instead* Array of likes for users who have hearted this task.
	Modified_at string `json:"modified_at,omitempty"` // The time at which this task was last modified. *Note: This does not currently reflect any changes in associations such as projects or comments that may have been added or removed from the task.*
	Start_at string `json:"start_at,omitempty"` // Date and time on which work begins for the task, or null if the task has no start time. This takes an ISO 8601 date string in UTC and should not be used together with `start_on`. *Note: `due_at` must be present in the request when setting or unsetting the `start_at` parameter.*
	External map[string]interface{} `json:"external,omitempty"` // *OAuth Required*. *Conditional*. This field is returned only if external values are set or included by using [Opt In] (/docs/input-output-options). The external field allows you to store app-specific metadata on tasks, including a gid that can be used to retrieve tasks and a data blob that can store app-specific character strings. Note that you will need to authenticate with Oauth to access or modify this data. Once an external gid is set, you can use the notation `external:custom_gid` to reference your object anywhere in the API where you may use the original object gid. See the page on Custom External Data for more details.
	Due_at string `json:"due_at,omitempty"` // The UTC date and time on which this task is due, or null if the task has no due time. This takes an ISO 8601 date string in UTC and should not be used together with `due_on`.
	Name string `json:"name,omitempty"` // Name of the task. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Actual_time_minutes float64 `json:"actual_time_minutes,omitempty"` // This value represents the sum of all the Time Tracking entries in the Actual Time field on a given Task. It is represented as a nullable long value.
	Dependents []AsanaResource `json:"dependents,omitempty"` // [Opt In](/docs/input-output-options). Array of resources referencing tasks that depend on this task. The objects contain only the ID of the dependent.
	Start_on string `json:"start_on,omitempty"` // The day on which work begins for the task , or null if the task has no start date. This takes a date with `YYYY-MM-DD` format and should not be used together with `start_at`. *Note: `due_on` or `due_at` must be present in the request when setting or unsetting the `start_on` parameter.*
	Liked bool `json:"liked,omitempty"` // True if the task is liked by the authorized user, false if not.
	Num_likes int `json:"num_likes,omitempty"` // The number of users who have liked this task.
	Completed_at string `json:"completed_at,omitempty"` // The time at which this task was completed, or null if the task is incomplete.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the task (i.e. its description).
	Assignee_status string `json:"assignee_status,omitempty"` // *Deprecated* Scheduling status of this task for the user it is assigned to. This field can only be set if the assignee is non-null. Setting this field to "inbox" or "upcoming" inserts it at the top of the section, while the other options will insert at the bottom.
	Html_notes string `json:"html_notes,omitempty"` // [Opt In](/docs/input-output-options). The notes of the text with formatting as HTML.
	Likes []Like `json:"likes,omitempty"` // Array of likes for users who have liked this task.
	Dependencies []AsanaResource `json:"dependencies,omitempty"` // [Opt In](/docs/input-output-options). Array of resources referencing tasks that this task depends on. The objects contain only the gid of the dependency.
	Hearted bool `json:"hearted,omitempty"` // *Deprecated - please use liked instead* True if the task is hearted by the authorized user, false if not.
	Is_rendered_as_separator bool `json:"is_rendered_as_separator,omitempty"` // [Opt In](/docs/input-output-options). In some contexts tasks can be rendered as a visual separator; for instance, subtasks can appear similar to [sections](/docs/asana-sections) without being true `section` objects. If a `task` object is rendered this way in any context it will have the property `is_rendered_as_separator` set to `true`.
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"` // An object where each key is a Custom Field GID and each value is an enum GID, string, number, object, or array.
	Followers []string `json:"followers,omitempty"` // *Create-Only* An array of strings identifying users. These can either be the string "me", an email, or the gid of a user. In order to change followers on an existing task use `addFollowers` and `removeFollowers`.
	Parent string `json:"parent,omitempty"` // Gid of a task.
	Projects []string `json:"projects,omitempty"` // *Create-Only* Array of project gids. In order to change projects on an existing task use `addProject` and `removeProject`.
	Tags []string `json:"tags,omitempty"` // *Create-Only* Array of tag gids. In order to change tags on an existing task use `addTag` and `removeTag`.
	Workspace string `json:"workspace,omitempty"` // Gid of a workspace.
	Assignee string `json:"assignee,omitempty"` // Gid of a user.
	Assignee_section string `json:"assignee_section,omitempty"` // The *assignee section* is a subdivision of a project that groups tasks together in the assignee's "My Tasks" list. It can either be a header above a list of tasks in a list view or a column in a board view of "My Tasks." The `assignee_section` property will be returned in the response only if the request was sent by the user who is the assignee of the task. Note that you can only write to `assignee_section` with the gid of an existing section visible in the user's "My Tasks" list.
}

// CustomFieldSettingBase represents the CustomFieldSettingBase schema from the OpenAPI specification
type CustomFieldSettingBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
}

// CustomFieldSettingCompact represents the CustomFieldSettingCompact schema from the OpenAPI specification
type CustomFieldSettingCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
}

// WorkspaceMembershipCompact represents the WorkspaceMembershipCompact schema from the OpenAPI specification
type WorkspaceMembershipCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	User UserCompact `json:"user,omitempty"`
	Workspace WorkspaceCompact `json:"workspace,omitempty"`
}

// StatusUpdateBase represents the StatusUpdateBase schema from the OpenAPI specification
type StatusUpdateBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The `resource_subtype`s for `status` objects represent the type of their parent.
	Title string `json:"title,omitempty"` // The title of the status update.
	Status_type string `json:"status_type"` // The type associated with the status update. This represents the current state of the object this object is on.
	Text string `json:"text"` // The text content of the status update.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). The text content of the status update with formatting as HTML.
}

// PortfolioAddItemRequest represents the PortfolioAddItemRequest schema from the OpenAPI specification
type PortfolioAddItemRequest struct {
	Insert_after string `json:"insert_after,omitempty"` // An id of an item in this portfolio. The new item will be added after the one specified here. `insert_before` and `insert_after` parameters cannot both be specified.
	Insert_before string `json:"insert_before,omitempty"` // An id of an item in this portfolio. The new item will be added before the one specified here. `insert_before` and `insert_after` parameters cannot both be specified.
	Item string `json:"item"` // The item to add to the portfolio.
}

// GoalRemoveSupportingRelationshipRequest represents the GoalRemoveSupportingRelationshipRequest schema from the OpenAPI specification
type GoalRemoveSupportingRelationshipRequest struct {
	Supporting_resource string `json:"supporting_resource"` // The gid of the supporting resource to remove from the parent goal. Must be the gid of a goal, project, or portfolio.
}

// WebhookCompact represents the WebhookCompact schema from the OpenAPI specification
type WebhookCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Target string `json:"target,omitempty"` // The URL to receive the HTTP POST.
	Active bool `json:"active,omitempty"` // If true, the webhook will send events - if false it is considered inactive and will not generate events.
	Resource AsanaNamedResource `json:"resource,omitempty"`
}

// AttachmentCompact represents the AttachmentCompact schema from the OpenAPI specification
type AttachmentCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the file.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The service hosting the attachment. Valid values are `asana`, `dropbox`, `gdrive`, `onedrive`, `box`, `vimeo`, and `external`.
}

// ProjectTemplateCompact represents the ProjectTemplateCompact schema from the OpenAPI specification
type ProjectTemplateCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the project template.
}

// EnumOptionRequest represents the EnumOptionRequest schema from the OpenAPI specification
type EnumOptionRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Enabled bool `json:"enabled,omitempty"` // Whether or not the enum option is a selectable value for the custom field.
	Name string `json:"name,omitempty"` // The name of the enum option.
	Color string `json:"color,omitempty"` // The color of the enum option. Defaults to ‘none’.
	Insert_before string `json:"insert_before,omitempty"` // An existing enum option within this custom field before which the new enum option should be inserted. Cannot be provided together with after_enum_option.
	Insert_after string `json:"insert_after,omitempty"` // An existing enum option within this custom field after which the new enum option should be inserted. Cannot be provided together with before_enum_option.
}

// EnumOptionInsertRequest represents the EnumOptionInsertRequest schema from the OpenAPI specification
type EnumOptionInsertRequest struct {
	After_enum_option string `json:"after_enum_option,omitempty"` // An existing enum option within this custom field after which the new enum option should be inserted. Cannot be provided together with before_enum_option.
	Before_enum_option string `json:"before_enum_option,omitempty"` // An existing enum option within this custom field before which the new enum option should be inserted. Cannot be provided together with after_enum_option.
	Enum_option string `json:"enum_option"` // The gid of the enum option to relocate.
}

// GoalMembershipResponse represents the GoalMembershipResponse schema from the OpenAPI specification
type GoalMembershipResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Is_commenter bool `json:"is_commenter,omitempty"` // Describes if the user is comment only in goal.
	Is_editor bool `json:"is_editor,omitempty"` // Describes if the user is editor in goal.
	Member UserCompact `json:"member,omitempty"`
}

// UserTaskListRequest represents the UserTaskListRequest schema from the OpenAPI specification
type UserTaskListRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Workspace interface{} `json:"workspace,omitempty"` // The workspace in which the user task list is located.
	Name string `json:"name,omitempty"` // The name of the user task list.
	Owner interface{} `json:"owner,omitempty"` // The owner of the user task list, i.e. the person whose My Tasks is represented by this resource.
}

// ProjectStatusBase represents the ProjectStatusBase schema from the OpenAPI specification
type ProjectStatusBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Title string `json:"title,omitempty"` // The title of the project status update.
	Color string `json:"color"` // The color associated with the status update.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). The text content of the status update with formatting as HTML.
	Text string `json:"text"` // The text content of the status update.
}

// TeamRequest represents the TeamRequest schema from the OpenAPI specification
type TeamRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the team.
	Html_description string `json:"html_description,omitempty"` // The description of the team with formatting as HTML.
	Organization string `json:"organization,omitempty"` // The organization/workspace the team belongs to. This must be the same organization you are in and cannot be changed once set.
	Visibility string `json:"visibility,omitempty"` // The visibility of the team to users in the same organization
	Description string `json:"description,omitempty"` // The description of the team.
}

// WorkspaceResponse represents the WorkspaceResponse schema from the OpenAPI specification
type WorkspaceResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the workspace.
	Email_domains []string `json:"email_domains,omitempty"` // The email domains that are associated with this workspace.
	Is_organization bool `json:"is_organization,omitempty"` // Whether the workspace is an *organization*.
}

// TaskResponse represents the TaskResponse schema from the OpenAPI specification
type TaskResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the task.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The resource_subtype `milestone` represent a single moment in time. This means tasks with this subtype cannot have a start_date.
	External map[string]interface{} `json:"external,omitempty"` // *OAuth Required*. *Conditional*. This field is returned only if external values are set or included by using [Opt In] (/docs/input-output-options). The external field allows you to store app-specific metadata on tasks, including a gid that can be used to retrieve tasks and a data blob that can store app-specific character strings. Note that you will need to authenticate with Oauth to access or modify this data. Once an external gid is set, you can use the notation `external:custom_gid` to reference your object anywhere in the API where you may use the original object gid. See the page on Custom External Data for more details.
	Due_at string `json:"due_at,omitempty"` // The UTC date and time on which this task is due, or null if the task has no due time. This takes an ISO 8601 date string in UTC and should not be used together with `due_on`.
	Name string `json:"name,omitempty"` // Name of the task. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Actual_time_minutes float64 `json:"actual_time_minutes,omitempty"` // This value represents the sum of all the Time Tracking entries in the Actual Time field on a given Task. It is represented as a nullable long value.
	Dependents []AsanaResource `json:"dependents,omitempty"` // [Opt In](/docs/input-output-options). Array of resources referencing tasks that depend on this task. The objects contain only the ID of the dependent.
	Start_on string `json:"start_on,omitempty"` // The day on which work begins for the task , or null if the task has no start date. This takes a date with `YYYY-MM-DD` format and should not be used together with `start_at`. *Note: `due_on` or `due_at` must be present in the request when setting or unsetting the `start_on` parameter.*
	Liked bool `json:"liked,omitempty"` // True if the task is liked by the authorized user, false if not.
	Num_likes int `json:"num_likes,omitempty"` // The number of users who have liked this task.
	Completed_at string `json:"completed_at,omitempty"` // The time at which this task was completed, or null if the task is incomplete.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the task (i.e. its description).
	Assignee_status string `json:"assignee_status,omitempty"` // *Deprecated* Scheduling status of this task for the user it is assigned to. This field can only be set if the assignee is non-null. Setting this field to "inbox" or "upcoming" inserts it at the top of the section, while the other options will insert at the bottom.
	Html_notes string `json:"html_notes,omitempty"` // [Opt In](/docs/input-output-options). The notes of the text with formatting as HTML.
	Likes []Like `json:"likes,omitempty"` // Array of likes for users who have liked this task.
	Dependencies []AsanaResource `json:"dependencies,omitempty"` // [Opt In](/docs/input-output-options). Array of resources referencing tasks that this task depends on. The objects contain only the gid of the dependency.
	Hearted bool `json:"hearted,omitempty"` // *Deprecated - please use liked instead* True if the task is hearted by the authorized user, false if not.
	Is_rendered_as_separator bool `json:"is_rendered_as_separator,omitempty"` // [Opt In](/docs/input-output-options). In some contexts tasks can be rendered as a visual separator; for instance, subtasks can appear similar to [sections](/docs/asana-sections) without being true `section` objects. If a `task` object is rendered this way in any context it will have the property `is_rendered_as_separator` set to `true`.
	Due_on string `json:"due_on,omitempty"` // The localized date on which this task is due, or null if the task has no due date. This takes a date with `YYYY-MM-DD` format and should not be used together with `due_at`.
	Approval_status string `json:"approval_status,omitempty"` // *Conditional* Reflects the approval status of this task. This field is kept in sync with `completed`, meaning `pending` translates to false while `approved`, `rejected`, and `changes_requested` translate to true. If you set completed to true, this field will be set to `approved`.
	Completed_by UserCompact `json:"completed_by,omitempty"`
	Completed bool `json:"completed,omitempty"` // True if the task is currently marked complete, false if not.
	Memberships []map[string]interface{} `json:"memberships,omitempty"` // *Create-only*. Array of projects this task is associated with and the section it is in. At task creation time, this array can be used to add the task to specific sections. After task creation, these associations can be modified using the `addProject` and `removeProject` endpoints. Note that over time, more types of memberships may be added to this property.
	Num_hearts int `json:"num_hearts,omitempty"` // *Deprecated - please use likes instead* The number of users who have hearted this task.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Num_subtasks int `json:"num_subtasks,omitempty"` // [Opt In](/docs/input-output-options). The number of subtasks on this task.
	Hearts []Like `json:"hearts,omitempty"` // *Deprecated - please use likes instead* Array of likes for users who have hearted this task.
	Modified_at string `json:"modified_at,omitempty"` // The time at which this task was last modified. *Note: This does not currently reflect any changes in associations such as projects or comments that may have been added or removed from the task.*
	Start_at string `json:"start_at,omitempty"` // Date and time on which work begins for the task, or null if the task has no start time. This takes an ISO 8601 date string in UTC and should not be used together with `start_on`. *Note: `due_at` must be present in the request when setting or unsetting the `start_at` parameter.*
	Assignee interface{} `json:"assignee,omitempty"`
	Tags []TagCompact `json:"tags,omitempty"` // Array of tags associated with this task. In order to change tags on an existing task use `addTag` and `removeTag`.
	Workspace interface{} `json:"workspace,omitempty"`
	Assignee_section interface{} `json:"assignee_section,omitempty"`
	Parent interface{} `json:"parent,omitempty"`
	Projects []ProjectCompact `json:"projects,omitempty"` // *Create-only.* Array of projects this task is associated with. At task creation time, this array can be used to add the task to many projects at once. After task creation, these associations can be modified using the addProject and removeProject endpoints.
	Custom_fields []CustomFieldResponse `json:"custom_fields,omitempty"` // Array of custom field values applied to the task. These represent the custom field values recorded on this project for a particular custom field. For example, these custom field values will contain an `enum_value` property for custom fields of type `enum`, a `text_value` property for custom fields of type `text`, and so on. Please note that the `gid` returned on each custom field value *is identical* to the `gid` of the custom field, which allows referencing the custom field metadata through the `/custom_fields/custom_field-gid` endpoint.
	Followers []UserCompact `json:"followers,omitempty"` // Array of users following this task.
	Permalink_url string `json:"permalink_url,omitempty"` // A url that points directly to the object within Asana.
}

// AttachmentBase represents the AttachmentBase schema from the OpenAPI specification
type AttachmentBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the file.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The service hosting the attachment. Valid values are `asana`, `dropbox`, `gdrive`, `onedrive`, `box`, `vimeo`, and `external`.
}

// TeamMembershipCompact represents the TeamMembershipCompact schema from the OpenAPI specification
type TeamMembershipCompact struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Is_guest bool `json:"is_guest,omitempty"` // Describes if the user is a guest in the team.
	Team TeamCompact `json:"team,omitempty"`
	User UserCompact `json:"user,omitempty"`
}

// UserTaskListBase represents the UserTaskListBase schema from the OpenAPI specification
type UserTaskListBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the user task list.
	Owner interface{} `json:"owner,omitempty"` // The owner of the user task list, i.e. the person whose My Tasks is represented by this resource.
	Workspace interface{} `json:"workspace,omitempty"` // The workspace in which the user task list is located.
}

// GoalMetricBase represents the GoalMetricBase schema from the OpenAPI specification
type GoalMetricBase struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Currency_code string `json:"currency_code,omitempty"` // ISO 4217 currency code to format this custom field. This will be null if the `unit` is not `currency`.
	Precision int `json:"precision,omitempty"` // *Conditional*. Only relevant for goal metrics of type ‘Number’. This field dictates the number of places after the decimal to round to, i.e. 0 is integer values, 1 rounds to the nearest tenth, and so on. Must be between 0 and 6, inclusive. For percentage format, this may be unintuitive, as a value of 0.25 has a precision of 0, while a value of 0.251 has a precision of 1. This is due to 0.25 being displayed as 25%.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Current_display_value string `json:"current_display_value,omitempty"` // This string is the current value of a goal metric of type string.
	Current_number_value float64 `json:"current_number_value,omitempty"` // This number is the current value of a goal metric of type number.
	Initial_number_value float64 `json:"initial_number_value,omitempty"` // This number is the start value of a goal metric of type number.
	Unit string `json:"unit,omitempty"` // A supported unit of measure for the goal metric, or none.
	Progress_source string `json:"progress_source,omitempty"` // This field defines how the progress value of a goal metric is being calculated. A goal's progress can be provided manually by the user, calculated automatically from contributing subgoals or projects, or managed by an integration with an external data source, such as Salesforce.
	Target_number_value float64 `json:"target_number_value,omitempty"` // This number is the end value of a goal metric of type number. This number cannot equal `initial_number_value`.
}

// ProjectMembershipResponse represents the ProjectMembershipResponse schema from the OpenAPI specification
type ProjectMembershipResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	User UserCompact `json:"user,omitempty"`
	Project ProjectCompact `json:"project,omitempty"`
	Write_access string `json:"write_access,omitempty"` // Whether the user has full access to the project or has comment-only access.
}

// AsanaResource represents the AsanaResource schema from the OpenAPI specification
type AsanaResource struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
}

// TeamResponse represents the TeamResponse schema from the OpenAPI specification
type TeamResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the team.
	Description string `json:"description,omitempty"` // [Opt In](/docs/input-output-options). The description of the team.
	Html_description string `json:"html_description,omitempty"` // [Opt In](/docs/input-output-options). The description of the team with formatting as HTML.
	Organization interface{} `json:"organization,omitempty"`
	Permalink_url string `json:"permalink_url,omitempty"` // A url that points directly to the object within Asana.
	Visibility string `json:"visibility,omitempty"` // The visibility of the team to users in the same organization
}

// WebhookResponse represents the WebhookResponse schema from the OpenAPI specification
type WebhookResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Target string `json:"target,omitempty"` // The URL to receive the HTTP POST.
	Active bool `json:"active,omitempty"` // If true, the webhook will send events - if false it is considered inactive and will not generate events.
	Resource AsanaNamedResource `json:"resource,omitempty"`
	Filters []interface{} `json:"filters,omitempty"` // Whitelist of filters to apply to events from this webhook. If a webhook event passes any of the filters the event will be delivered; otherwise no event will be sent to the receiving server.
	Last_failure_at string `json:"last_failure_at,omitempty"` // The timestamp when the webhook last received an error when sending an event to the target.
	Last_failure_content string `json:"last_failure_content,omitempty"` // The contents of the last error response sent to the webhook when attempting to deliver events to the target.
	Last_success_at string `json:"last_success_at,omitempty"` // The timestamp when the webhook last successfully sent an event to the target.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
}

// Preview represents the Preview schema from the OpenAPI specification
type Preview struct {
	Fallback string `json:"fallback,omitempty"` // Some fallback text to display if unable to display the full preview.
	Footer string `json:"footer,omitempty"` // Text to display in the footer.
	Header string `json:"header,omitempty"` // Text to display in the header.
	Header_link string `json:"header_link,omitempty"` // Where the header will link to.
	Html_text string `json:"html_text,omitempty"` // HTML formatted text for the body of the preview.
	Text string `json:"text,omitempty"` // Text for the body of the preview.
	Title string `json:"title,omitempty"` // Text to display as the title.
	Title_link string `json:"title_link,omitempty"` // Where to title will link to.
}

// TaskAddTagRequest represents the TaskAddTagRequest schema from the OpenAPI specification
type TaskAddTagRequest struct {
	Tag string `json:"tag"` // The tag to add to the task.
}

// GoalCompact represents the GoalCompact schema from the OpenAPI specification
type GoalCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the goal.
	Owner interface{} `json:"owner,omitempty"`
}

// ModifyDependenciesRequest represents the ModifyDependenciesRequest schema from the OpenAPI specification
type ModifyDependenciesRequest struct {
	Dependencies []string `json:"dependencies,omitempty"` // An array of task gids that a task depends on.
}

// SectionResponse represents the SectionResponse schema from the OpenAPI specification
type SectionResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the section (i.e. the text displayed as the section header).
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Project ProjectCompact `json:"project,omitempty"`
	Projects []ProjectCompact `json:"projects,omitempty"` // *Deprecated - please use project instead*
}

// GoalMembershipBase represents the GoalMembershipBase schema from the OpenAPI specification
type GoalMembershipBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Is_commenter bool `json:"is_commenter,omitempty"` // Describes if the user is comment only in goal.
	Is_editor bool `json:"is_editor,omitempty"` // Describes if the user is editor in goal.
	Member UserCompact `json:"member,omitempty"`
}

// ProjectRequest represents the ProjectRequest schema from the OpenAPI specification
type ProjectRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the project. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Html_notes string `json:"html_notes,omitempty"` // [Opt In](/docs/input-output-options). The notes of the project with formatting as HTML.
	Public bool `json:"public,omitempty"` // True if the project is public to its team.
	Modified_at string `json:"modified_at,omitempty"` // The time at which this project was last modified. *Note: This does not currently reflect any changes in associations such as tasks or comments that may have been added or removed from the project.*
	Archived bool `json:"archived,omitempty"` // True if the project is archived, false if not. Archived projects do not show in the UI by default and may be treated differently for queries.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Due_date string `json:"due_date,omitempty"` // *Deprecated: new integrations should prefer the `due_on` field.*
	Is_template bool `json:"is_template,omitempty"` // [Opt In](/docs/input-output-options). *Deprecated - please use a project template endpoint instead (more in [this forum post](https://forum.asana.com/t/a-new-api-for-project-templates/156432)).* Determines if the project is a template.
	Color string `json:"color,omitempty"` // Color of the project.
	Due_on string `json:"due_on,omitempty"` // The day on which this project is due. This takes a date with format YYYY-MM-DD.
	Current_status interface{} `json:"current_status,omitempty"` // *Deprecated: new integrations should prefer the `current_status_update` resource.*
	Current_status_update interface{} `json:"current_status_update,omitempty"` // The latest `status_update` posted to this project.
	Members []UserCompact `json:"members,omitempty"` // Array of users who are members of this project.
	Start_on string `json:"start_on,omitempty"` // The day on which work for this project begins, or null if the project has no start date. This takes a date with `YYYY-MM-DD` format. *Note: `due_on` or `due_at` must be present in the request when setting or unsetting the `start_on` parameter. Additionally, `start_on` and `due_on` cannot be the same date.*
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the project (ie., its description).
	Workspace interface{} `json:"workspace,omitempty"`
	Default_view string `json:"default_view,omitempty"` // The default view (list, board, calendar, or timeline) of a project.
	Custom_field_settings []CustomFieldSettingResponse `json:"custom_field_settings,omitempty"` // Array of Custom Field Settings (in compact form).
	Team string `json:"team,omitempty"` // The team that this project is shared with.
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"` // An object where each key is a Custom Field GID and each value is an enum GID, string, number, or object.
	Followers string `json:"followers,omitempty"` // *Create-only*. Comma separated string of users. Followers are a subset of members who have opted in to receive "tasks added" notifications for a project.
	Owner string `json:"owner,omitempty"` // The current owner of the project, may be null.
}

// UserBase represents the UserBase schema from the OpenAPI specification
type UserBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // *Read-only except when same user as requester*. The user’s name.
}

// ProjectSectionInsertRequest represents the ProjectSectionInsertRequest schema from the OpenAPI specification
type ProjectSectionInsertRequest struct {
	After_section string `json:"after_section,omitempty"` // Insert the given section immediately after the section specified by this parameter.
	Before_section string `json:"before_section,omitempty"` // Insert the given section immediately before the section specified by this parameter.
	Project string `json:"project"` // The project in which to reorder the given section.
	Section string `json:"section"` // The section to reorder.
}

// UserBaseResponse represents the UserBaseResponse schema from the OpenAPI specification
type UserBaseResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // *Read-only except when same user as requester*. The user’s name.
	Email string `json:"email,omitempty"` // The user's email address.
	Photo map[string]interface{} `json:"photo,omitempty"` // A map of the user’s profile photo in various sizes, or null if no photo is set. Sizes provided are 21, 27, 36, 60, 128, and 1024. All images are in PNG format, except for 1024 (which is in JPEG format).
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Errors []Error `json:"errors,omitempty"`
}

// EnumOption represents the EnumOption schema from the OpenAPI specification
type EnumOption struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the enum option.
	Color string `json:"color,omitempty"` // The color of the enum option. Defaults to ‘none’.
	Enabled bool `json:"enabled,omitempty"` // Whether or not the enum option is a selectable value for the custom field.
}

// GoalRequest represents the GoalRequest schema from the OpenAPI specification
type GoalRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Liked bool `json:"liked,omitempty"` // True if the goal is liked by the authorized user, false if not.
	Name string `json:"name,omitempty"` // The name of the goal.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the goal (i.e. its description).
	Start_on string `json:"start_on,omitempty"` // The day on which work for this goal begins, or null if the goal has no start date. This takes a date with `YYYY-MM-DD` format, and cannot be set unless there is an accompanying due date.
	Status string `json:"status,omitempty"` // The current status of this goal. When the goal is open, its status can be `green`, `yellow`, and `red` to reflect "On Track", "At Risk", and "Off Track", respectively. When the goal is closed, the value can be `missed`, `achieved`, `partial`, or `dropped`. *Note* you can only write to this property if `metric` is set.
	Due_on string `json:"due_on,omitempty"` // The localized day on which this goal is due. This takes a date with format `YYYY-MM-DD`.
	Html_notes string `json:"html_notes,omitempty"` // The notes of the goal with formatting as HTML.
	Is_workspace_level bool `json:"is_workspace_level,omitempty"` // *Conditional*. This property is only present when the `workspace` provided is an organization. Whether the goal belongs to the `workspace` (and is listed as part of the workspace’s goals) or not. If it isn’t a workspace-level goal, it is a team-level goal, and is associated with the goal’s team.
	Time_period string `json:"time_period,omitempty"` // The `gid` of a time period.
	Workspace string `json:"workspace,omitempty"` // The `gid` of a workspace.
	Followers []string `json:"followers,omitempty"`
	Owner string `json:"owner,omitempty"` // The `gid` of a user.
	Team string `json:"team,omitempty"` // *Conditional*. This property is only present when the `workspace` provided is an organization.
}

// StatusUpdateCompact represents the StatusUpdateCompact schema from the OpenAPI specification
type StatusUpdateCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning. The `resource_subtype`s for `status` objects represent the type of their parent.
	Title string `json:"title,omitempty"` // The title of the status update.
}

// TeamMembershipResponse represents the TeamMembershipResponse schema from the OpenAPI specification
type TeamMembershipResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Is_guest bool `json:"is_guest,omitempty"` // Describes if the user is a guest in the team.
	Team TeamCompact `json:"team,omitempty"`
	User UserCompact `json:"user,omitempty"`
}

// ProjectBriefCompact represents the ProjectBriefCompact schema from the OpenAPI specification
type ProjectBriefCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
}

// JobResponse represents the JobResponse schema from the OpenAPI specification
type JobResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Status string `json:"status,omitempty"` // The current status of this job. The value is one of: `not_started`, `in_progress`, `succeeded`, or `failed`.
	New_project ProjectCompact `json:"new_project,omitempty"`
	New_project_template ProjectTemplateCompact `json:"new_project_template,omitempty"`
	New_task TaskCompact `json:"new_task,omitempty"`
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
}

// TeamMembershipBase represents the TeamMembershipBase schema from the OpenAPI specification
type TeamMembershipBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Is_guest bool `json:"is_guest,omitempty"` // Describes if the user is a guest in the team.
	Team TeamCompact `json:"team,omitempty"`
	User UserCompact `json:"user,omitempty"`
}

// AuditLogEventDetails represents the AuditLogEventDetails schema from the OpenAPI specification
type AuditLogEventDetails struct {
}

// ProjectBase represents the ProjectBase schema from the OpenAPI specification
type ProjectBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the project. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Custom_field_settings []CustomFieldSettingResponse `json:"custom_field_settings,omitempty"` // Array of Custom Field Settings (in compact form).
	Html_notes string `json:"html_notes,omitempty"` // [Opt In](/docs/input-output-options). The notes of the project with formatting as HTML.
	Public bool `json:"public,omitempty"` // True if the project is public to its team.
	Modified_at string `json:"modified_at,omitempty"` // The time at which this project was last modified. *Note: This does not currently reflect any changes in associations such as tasks or comments that may have been added or removed from the project.*
	Archived bool `json:"archived,omitempty"` // True if the project is archived, false if not. Archived projects do not show in the UI by default and may be treated differently for queries.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Due_date string `json:"due_date,omitempty"` // *Deprecated: new integrations should prefer the `due_on` field.*
	Is_template bool `json:"is_template,omitempty"` // [Opt In](/docs/input-output-options). *Deprecated - please use a project template endpoint instead (more in [this forum post](https://forum.asana.com/t/a-new-api-for-project-templates/156432)).* Determines if the project is a template.
	Color string `json:"color,omitempty"` // Color of the project.
	Due_on string `json:"due_on,omitempty"` // The day on which this project is due. This takes a date with format YYYY-MM-DD.
	Current_status interface{} `json:"current_status,omitempty"` // *Deprecated: new integrations should prefer the `current_status_update` resource.*
	Current_status_update interface{} `json:"current_status_update,omitempty"` // The latest `status_update` posted to this project.
	Members []UserCompact `json:"members,omitempty"` // Array of users who are members of this project.
	Start_on string `json:"start_on,omitempty"` // The day on which work for this project begins, or null if the project has no start date. This takes a date with `YYYY-MM-DD` format. *Note: `due_on` or `due_at` must be present in the request when setting or unsetting the `start_on` parameter. Additionally, `start_on` and `due_on` cannot be the same date.*
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the project (ie., its description).
	Workspace interface{} `json:"workspace,omitempty"`
	Default_view string `json:"default_view,omitempty"` // The default view (list, board, calendar, or timeline) of a project.
}

// PortfolioResponse represents the PortfolioResponse schema from the OpenAPI specification
type PortfolioResponse struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Name string `json:"name,omitempty"` // The name of the portfolio.
	Color string `json:"color,omitempty"` // Color of the portfolio.
	Custom_field_settings []CustomFieldSettingResponse `json:"custom_field_settings,omitempty"` // Array of custom field settings applied to the portfolio.
	Due_on string `json:"due_on,omitempty"` // The localized day on which this portfolio is due. This takes a date with format YYYY-MM-DD.
	Members []UserCompact `json:"members,omitempty"`
	Permalink_url string `json:"permalink_url,omitempty"` // A url that points directly to the object within Asana.
	Start_on string `json:"start_on,omitempty"` // The day on which work for this portfolio begins, or null if the portfolio has no start date. This takes a date with `YYYY-MM-DD` format. *Note: `due_on` must be present in the request when setting or unsetting the `start_on` parameter. Additionally, `start_on` and `due_on` cannot be the same date.*
	Created_by UserCompact `json:"created_by,omitempty"`
	Custom_fields []CustomFieldCompact `json:"custom_fields,omitempty"` // Array of Custom Fields.
	Owner UserCompact `json:"owner,omitempty"`
	Public bool `json:"public,omitempty"` // True if the portfolio is public to its workspace members.
	Workspace interface{} `json:"workspace,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Current_status_update interface{} `json:"current_status_update,omitempty"` // The latest `status_update` posted to this portfolio.
}

// TaskCountResponse represents the TaskCountResponse schema from the OpenAPI specification
type TaskCountResponse struct {
	Num_incomplete_tasks int `json:"num_incomplete_tasks,omitempty"` // The number of incomplete tasks in a project.
	Num_milestones int `json:"num_milestones,omitempty"` // The number of milestones in a project.
	Num_tasks int `json:"num_tasks,omitempty"` // The number of tasks in a project.
	Num_completed_milestones int `json:"num_completed_milestones,omitempty"` // The number of completed milestones in a project.
	Num_completed_tasks int `json:"num_completed_tasks,omitempty"` // The number of completed tasks in a project.
	Num_incomplete_milestones int `json:"num_incomplete_milestones,omitempty"` // The number of incomplete milestones in a project.
}

// SectionBase represents the SectionBase schema from the OpenAPI specification
type SectionBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the section (i.e. the text displayed as the section header).
}

// ProjectStatusResponse represents the ProjectStatusResponse schema from the OpenAPI specification
type ProjectStatusResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Title string `json:"title,omitempty"` // The title of the project status update.
	Color string `json:"color"` // The color associated with the status update.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). The text content of the status update with formatting as HTML.
	Text string `json:"text"` // The text content of the status update.
	Author UserCompact `json:"author,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Created_by UserCompact `json:"created_by,omitempty"`
	Modified_at string `json:"modified_at,omitempty"` // The time at which this project status was last modified. *Note: This does not currently reflect any changes in associations such as comments that may have been added or removed from the project status.*
}

// ProjectTemplateResponse represents the ProjectTemplateResponse schema from the OpenAPI specification
type ProjectTemplateResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the project template.
	Public bool `json:"public,omitempty"` // True if the project template is public to its team.
	Requested_dates []DateVariableCompact `json:"requested_dates,omitempty"` // Array of date variables in this project template. Calendar dates must be provided for these variables when instantiating a project.
	Team interface{} `json:"team,omitempty"`
	Color string `json:"color,omitempty"` // Color of the project template.
	Description string `json:"description,omitempty"` // Free-form textual information associated with the project template
	Html_description string `json:"html_description,omitempty"` // The description of the project template with formatting as HTML.
	Owner interface{} `json:"owner,omitempty"` // The current owner of the project template, may be null.
}

// StoryRequest represents the StoryRequest schema from the OpenAPI specification
type StoryRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Text string `json:"text,omitempty"` // The plain text of the comment to add. Cannot be used with html_text.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). HTML formatted text for a comment. This will not include the name of the creator.
	Is_pinned bool `json:"is_pinned,omitempty"` // *Conditional*. Whether the story should be pinned on the resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Sticker_name string `json:"sticker_name,omitempty"` // The name of the sticker in this story. `null` if there is no sticker.
}

// ProjectBriefResponse represents the ProjectBriefResponse schema from the OpenAPI specification
type ProjectBriefResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Html_text string `json:"html_text,omitempty"` // HTML formatted text for the project brief.
	Title string `json:"title,omitempty"` // The title of the project brief.
	Permalink_url string `json:"permalink_url,omitempty"` // A url that points directly to the object within Asana.
	Project interface{} `json:"project,omitempty"`
	Text string `json:"text,omitempty"` // [Opt In](/docs/input-output-options). The plain text of the project brief.
}

// GoalMembershipCompact represents the GoalMembershipCompact schema from the OpenAPI specification
type GoalMembershipCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Is_commenter bool `json:"is_commenter,omitempty"` // Describes if the user is comment only in goal.
	Is_editor bool `json:"is_editor,omitempty"` // Describes if the user is editor in goal.
	Member UserCompact `json:"member,omitempty"`
}

// JobBase represents the JobBase schema from the OpenAPI specification
type JobBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Status string `json:"status,omitempty"` // The current status of this job. The value is one of: `not_started`, `in_progress`, `succeeded`, or `failed`.
	New_project ProjectCompact `json:"new_project,omitempty"`
	New_project_template ProjectTemplateCompact `json:"new_project_template,omitempty"`
	New_task TaskCompact `json:"new_task,omitempty"`
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
}

// DateVariableRequest represents the DateVariableRequest schema from the OpenAPI specification
type DateVariableRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the date field in the project template. A value of `1` refers to the project start date, while `2` refers to the project due date.
	Value string `json:"value,omitempty"` // The date with which the date variable should be replaced when instantiating a project. This takes a date with `YYYY-MM-DD` format.
}

// GoalAddSupportingRelationshipRequest represents the GoalAddSupportingRelationshipRequest schema from the OpenAPI specification
type GoalAddSupportingRelationshipRequest struct {
	Insert_after string `json:"insert_after,omitempty"` // An id of a subgoal of this parent goal. The new subgoal will be added after the one specified here. `insert_before` and `insert_after` parameters cannot both be specified. Currently only supported when adding a subgoal.
	Insert_before string `json:"insert_before,omitempty"` // An id of a subgoal of this parent goal. The new subgoal will be added before the one specified here. `insert_before` and `insert_after` parameters cannot both be specified. Currently only supported when adding a subgoal.
	Supporting_resource string `json:"supporting_resource"` // The gid of the supporting resource to add to the parent goal. Must be the gid of a goal, project, or portfolio.
	Contribution_weight float64 `json:"contribution_weight,omitempty"` // The weight that the supporting resource's progress will contribute to the supported goal's progress. This can only be 0 or 1.
}

// GoalBase represents the GoalBase schema from the OpenAPI specification
type GoalBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Due_on string `json:"due_on,omitempty"` // The localized day on which this goal is due. This takes a date with format `YYYY-MM-DD`.
	Html_notes string `json:"html_notes,omitempty"` // The notes of the goal with formatting as HTML.
	Is_workspace_level bool `json:"is_workspace_level,omitempty"` // *Conditional*. This property is only present when the `workspace` provided is an organization. Whether the goal belongs to the `workspace` (and is listed as part of the workspace’s goals) or not. If it isn’t a workspace-level goal, it is a team-level goal, and is associated with the goal’s team.
	Liked bool `json:"liked,omitempty"` // True if the goal is liked by the authorized user, false if not.
	Name string `json:"name,omitempty"` // The name of the goal.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the goal (i.e. its description).
	Start_on string `json:"start_on,omitempty"` // The day on which work for this goal begins, or null if the goal has no start date. This takes a date with `YYYY-MM-DD` format, and cannot be set unless there is an accompanying due date.
	Status string `json:"status,omitempty"` // The current status of this goal. When the goal is open, its status can be `green`, `yellow`, and `red` to reflect "On Track", "At Risk", and "Off Track", respectively. When the goal is closed, the value can be `missed`, `achieved`, `partial`, or `dropped`. *Note* you can only write to this property if `metric` is set.
}

// ProjectMembershipBase represents the ProjectMembershipBase schema from the OpenAPI specification
type ProjectMembershipBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	User UserCompact `json:"user,omitempty"`
}

// WorkspaceMembershipResponse represents the WorkspaceMembershipResponse schema from the OpenAPI specification
type WorkspaceMembershipResponse struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	User UserCompact `json:"user,omitempty"`
	Workspace WorkspaceCompact `json:"workspace,omitempty"`
	Is_admin bool `json:"is_admin,omitempty"` // Reflects if this user is an admin of the workspace.
	Is_guest bool `json:"is_guest,omitempty"` // Reflects if this user is a guest of the workspace.
	User_task_list UserTaskListResponse `json:"user_task_list,omitempty"`
	Vacation_dates map[string]interface{} `json:"vacation_dates,omitempty"` // Contains keys `start_on` and `end_on` for the vacation dates for the user in this workspace. If `start_on` is null, the entire `vacation_dates` object will be null. If `end_on` is before today, the entire `vacation_dates` object will be null.
	Is_active bool `json:"is_active,omitempty"` // Reflects if this user still a member of the workspace.
}

// TaskAddProjectRequest represents the TaskAddProjectRequest schema from the OpenAPI specification
type TaskAddProjectRequest struct {
	Section string `json:"section,omitempty"` // A section in the project to insert the task into. The task will be inserted at the bottom of the section.
	Insert_after string `json:"insert_after,omitempty"` // A task in the project to insert the task after, or `null` to insert at the beginning of the list.
	Insert_before string `json:"insert_before,omitempty"` // A task in the project to insert the task before, or `null` to insert at the end of the list.
	Project string `json:"project"` // The project to add the task to.
}

// BatchResponse represents the BatchResponse schema from the OpenAPI specification
type BatchResponse struct {
	Body map[string]interface{} `json:"body,omitempty"` // The JSON body that the invoked endpoint returned.
	Headers map[string]interface{} `json:"headers,omitempty"` // A map of HTTP headers specific to this result. This is primarily used for returning a `Location` header to accompany a `201 Created` result. The parent HTTP response will contain all common headers.
	Status_code int `json:"status_code,omitempty"` // The HTTP status code that the invoked endpoint returned.
}

// AsanaNamedResource represents the AsanaNamedResource schema from the OpenAPI specification
type AsanaNamedResource struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the object.
}

// AddMembersRequest represents the AddMembersRequest schema from the OpenAPI specification
type AddMembersRequest struct {
	Members string `json:"members"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
}

// TimePeriodCompact represents the TimePeriodCompact schema from the OpenAPI specification
type TimePeriodCompact struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	End_on string `json:"end_on,omitempty"` // The localized end date of the time period in `YYYY-MM-DD` format.
	Period string `json:"period,omitempty"` // The cadence and index of the time period. The value is one of: `FY`, `H1`, `H2`, `Q1`, `Q2`, `Q3`, or `Q4`.
	Start_on string `json:"start_on,omitempty"` // The localized start date of the time period in `YYYY-MM-DD` format.
	Display_name string `json:"display_name,omitempty"` // A string representing the cadence code and the fiscal year.
}

// WorkspaceCompact represents the WorkspaceCompact schema from the OpenAPI specification
type WorkspaceCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the workspace.
}

// PortfolioCompact represents the PortfolioCompact schema from the OpenAPI specification
type PortfolioCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the portfolio.
}

// GoalRemoveSubgoalRequest represents the GoalRemoveSubgoalRequest schema from the OpenAPI specification
type GoalRemoveSubgoalRequest struct {
	Subgoal string `json:"subgoal"` // The goal gid to remove as subgoal from the parent goal
}

// GoalRelationshipCompact represents the GoalRelationshipCompact schema from the OpenAPI specification
type GoalRelationshipCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Contribution_weight float64 `json:"contribution_weight,omitempty"` // The weight that the supporting resource's progress contributes to the supported goal's progress. This can only be 0 or 1.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Supporting_resource interface{} `json:"supporting_resource,omitempty"`
}

// TagBase represents the TagBase schema from the OpenAPI specification
type TagBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the tag. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
	Notes string `json:"notes,omitempty"` // Free-form textual information associated with the tag (i.e. its description).
	Color string `json:"color,omitempty"` // Color of the tag.
}

// TaskRemoveFollowersRequest represents the TaskRemoveFollowersRequest schema from the OpenAPI specification
type TaskRemoveFollowersRequest struct {
	Followers []string `json:"followers"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
}

// WorkspaceMembershipBase represents the WorkspaceMembershipBase schema from the OpenAPI specification
type WorkspaceMembershipBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	User UserCompact `json:"user,omitempty"`
	Workspace WorkspaceCompact `json:"workspace,omitempty"`
}

// StoryCompact represents the StoryCompact schema from the OpenAPI specification
type StoryCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Text string `json:"text,omitempty"` // *Create-only*. Human-readable text for the story or comment. This will not include the name of the creator. *Note: This is not guaranteed to be stable for a given type of story. For example, text for a reassignment may not always say “assigned to …” as the text for a story can both be edited and change based on the language settings of the user making the request.* Use the `resource_subtype` property to discover the action that created the story.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Created_by UserCompact `json:"created_by,omitempty"`
}

// StoryResponse represents the StoryResponse schema from the OpenAPI specification
type StoryResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Is_pinned bool `json:"is_pinned,omitempty"` // *Conditional*. Whether the story should be pinned on the resource.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Sticker_name string `json:"sticker_name,omitempty"` // The name of the sticker in this story. `null` if there is no sticker.
	Text string `json:"text,omitempty"` // The plain text of the comment to add. Cannot be used with html_text.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). HTML formatted text for a comment. This will not include the name of the creator.
	Source string `json:"source,omitempty"` // The component of the Asana product the user used to trigger the story.
	Num_hearts int `json:"num_hearts,omitempty"` // *Deprecated - please use likes instead* *Conditional*. The number of users who have hearted this story.
	New_approval_status string `json:"new_approval_status,omitempty"` // *Conditional*. The new value of approval status.
	Story StoryCompact `json:"story,omitempty"`
	Old_name string `json:"old_name,omitempty"` // *Conditional*'
	Previews []Preview `json:"previews,omitempty"` // *Conditional*. A collection of previews to be displayed in the story. *Note: This property only exists for comment stories.*
	TypeField string `json:"type,omitempty"`
	Task TaskCompact `json:"task,omitempty"`
	Old_multi_enum_values []EnumOption `json:"old_multi_enum_values,omitempty"` // *Conditional*. The old value of a multi-enum custom field story.
	New_section SectionCompact `json:"new_section,omitempty"`
	Old_people_value []UserCompact `json:"old_people_value,omitempty"` // *Conditional*. The old value of a people custom field story.
	Duplicated_from TaskCompact `json:"duplicated_from,omitempty"`
	Assignee UserCompact `json:"assignee,omitempty"`
	Old_resource_subtype string `json:"old_resource_subtype,omitempty"` // *Conditional*
	Created_by UserCompact `json:"created_by,omitempty"`
	Is_editable bool `json:"is_editable,omitempty"` // *Conditional*. Whether the text of the story can be edited after creation.
	Old_enum_value EnumOption `json:"old_enum_value,omitempty"`
	New_people_value []UserCompact `json:"new_people_value,omitempty"` // *Conditional*. The new value of a people custom field story.
	Num_likes int `json:"num_likes,omitempty"` // *Conditional*. The number of users who have liked this story.
	New_number_value int `json:"new_number_value,omitempty"` // *Conditional*
	Target interface{} `json:"target,omitempty"`
	Project ProjectCompact `json:"project,omitempty"`
	New_date_value interface{} `json:"new_date_value,omitempty"`
	New_multi_enum_values []EnumOption `json:"new_multi_enum_values,omitempty"` // *Conditional*. The new value of a multi-enum custom field story.
	Custom_field CustomFieldCompact `json:"custom_field,omitempty"`
	Hearted bool `json:"hearted,omitempty"` // *Deprecated - please use likes instead* *Conditional*. True if the story is hearted by the authorized user, false if not.
	Old_approval_status string `json:"old_approval_status,omitempty"` // *Conditional*. The old value of approval status.
	New_resource_subtype string `json:"new_resource_subtype,omitempty"` // *Conditional*
	New_enum_value EnumOption `json:"new_enum_value,omitempty"`
	New_dates StoryResponseDates `json:"new_dates,omitempty"` // *Conditional*
	Old_date_value interface{} `json:"old_date_value,omitempty"`
	Hearts []Like `json:"hearts,omitempty"` // *Deprecated - please use likes instead* *Conditional*. Array of likes for users who have hearted this story.
	Old_text_value string `json:"old_text_value,omitempty"` // *Conditional*
	Liked bool `json:"liked,omitempty"` // *Conditional*. True if the story is liked by the authorized user, false if not.
	Old_dates StoryResponseDates `json:"old_dates,omitempty"` // *Conditional*
	Likes []Like `json:"likes,omitempty"` // *Conditional*. Array of likes for users who have liked this story.
	Follower UserCompact `json:"follower,omitempty"`
	New_name string `json:"new_name,omitempty"` // *Conditional*
	Old_section SectionCompact `json:"old_section,omitempty"`
	Old_number_value int `json:"old_number_value,omitempty"` // *Conditional*
	Tag TagCompact `json:"tag,omitempty"`
	Is_edited bool `json:"is_edited,omitempty"` // *Conditional*. Whether the text of the story has been edited after creation.
	Dependency TaskCompact `json:"dependency,omitempty"`
	Duplicate_of TaskCompact `json:"duplicate_of,omitempty"`
	New_text_value string `json:"new_text_value,omitempty"` // *Conditional*
}

// GoalRelationshipBase represents the GoalRelationshipBase schema from the OpenAPI specification
type GoalRelationshipBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Contribution_weight float64 `json:"contribution_weight,omitempty"` // The weight that the supporting resource's progress contributes to the supported goal's progress. This can only be 0 or 1.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
	Supporting_resource interface{} `json:"supporting_resource,omitempty"`
	Supported_goal interface{} `json:"supported_goal,omitempty"`
}

// TeamAddUserRequest represents the TeamAddUserRequest schema from the OpenAPI specification
type TeamAddUserRequest struct {
	User string `json:"user,omitempty"` // A string identifying a user. This can either be the string "me", an email, or the gid of a user.
}

// ProjectTemplateInstantiateProjectRequest represents the ProjectTemplateInstantiateProjectRequest schema from the OpenAPI specification
type ProjectTemplateInstantiateProjectRequest struct {
	Requested_dates []DateVariableRequest `json:"requested_dates,omitempty"` // Array of mappings of date variables to calendar dates.
	Team string `json:"team,omitempty"` // *Conditional*. Sets the team of the new project. If the project template exists in an _organization_, you must specify a value for `team` and not `workspace`.
	Workspace string `json:"workspace,omitempty"` // *Conditional*. Sets the workspace of the new project. If the project template exists in a _workspace_, you must specify a value for `workspace` and not `team`.
	Is_strict bool `json:"is_strict,omitempty"` // *Optional*. If set to `true`, the endpoint returns an "Unprocessable Entity" error if you fail to provide a calendar date value for any date variable. If set to `false`, a default date is used for each unfulfilled date variable (e.g., the current date is used as the Start Date of a project).
	Name string `json:"name"` // The name of the new project.
	Public bool `json:"public"` // Sets the project to public to its team.
}

// WebhookFilter represents the WebhookFilter schema from the OpenAPI specification
type WebhookFilter struct {
	Resource_subtype string `json:"resource_subtype,omitempty"` // The resource subtype of the resource that the filter applies to. This should be set to the same value as is returned on the `resource_subtype` field on the resources themselves.
	Resource_type string `json:"resource_type,omitempty"` // The type of the resource which created the event when modified; for example, to filter to changes on regular tasks this field should be set to `task`.
	Action string `json:"action,omitempty"` // The type of change on the **resource** to pass through the filter. For more information refer to `Event.action` in the [Event](/docs/tocS_Event) schema. This can be one of `changed`, `added`, `removed`, `deleted`, and `undeleted` depending on the nature of what has occurred on the resource.
	Fields []string `json:"fields,omitempty"` // *Conditional.* A whitelist of fields for events which will pass the filter when the resource is changed. These can be any combination of the fields on the resources themselves. This field is only valid for `action` of type `changed`
}

// WorkspaceAddUserRequest represents the WorkspaceAddUserRequest schema from the OpenAPI specification
type WorkspaceAddUserRequest struct {
	User string `json:"user,omitempty"` // A string identifying a user. This can either be the string "me", an email, or the gid of a user.
}

// GoalAddSubgoalRequest represents the GoalAddSubgoalRequest schema from the OpenAPI specification
type GoalAddSubgoalRequest struct {
	Subgoal string `json:"subgoal"` // The goal gid to add as subgoal to a parent goal
	Insert_after string `json:"insert_after,omitempty"` // An id of a subgoal of this parent goal. The new subgoal will be added after the one specified here. `insert_before` and `insert_after` parameters cannot both be specified.
	Insert_before string `json:"insert_before,omitempty"` // An id of a subgoal of this parent goal. The new subgoal will be added before the one specified here. `insert_before` and `insert_after` parameters cannot both be specified.
}

// AddCustomFieldSettingRequest represents the AddCustomFieldSettingRequest schema from the OpenAPI specification
type AddCustomFieldSettingRequest struct {
	Insert_after string `json:"insert_after,omitempty"` // A gid of a Custom Field Setting on this container, after which the new Custom Field Setting will be added. `insert_before` and `insert_after` parameters cannot both be specified.
	Insert_before string `json:"insert_before,omitempty"` // A gid of a Custom Field Setting on this container, before which the new Custom Field Setting will be added. `insert_before` and `insert_after` parameters cannot both be specified.
	Is_important bool `json:"is_important,omitempty"` // Whether this field should be considered important to this container (for instance, to display in the list view of items in the container).
	Custom_field string `json:"custom_field"` // The custom field to associate with this container.
}

// EventResponse represents the EventResponse schema from the OpenAPI specification
type EventResponse struct {
	Created_at string `json:"created_at,omitempty"` // The timestamp when the event occurred.
	Parent interface{} `json:"parent,omitempty"`
	Resource interface{} `json:"resource,omitempty"`
	TypeField string `json:"type,omitempty"` // *Deprecated: Refer to the resource_type of the resource.* The type of the resource that generated the event.
	User interface{} `json:"user,omitempty"`
	Action string `json:"action,omitempty"` // The type of action taken on the **resource** that triggered the event. This can be one of `changed`, `added`, `removed`, `deleted`, or `undeleted` depending on the nature of the event.
	Change map[string]interface{} `json:"change,omitempty"` // Information about the type of change that has occurred. This field is only present when the value of the property `action`, describing the action taken on the **resource**, is `changed`.
}

// ProjectMembershipCompact represents the ProjectMembershipCompact schema from the OpenAPI specification
type ProjectMembershipCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	User UserCompact `json:"user,omitempty"`
}

// GoalMetricRequest represents the GoalMetricRequest schema from the OpenAPI specification
type GoalMetricRequest struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Current_display_value string `json:"current_display_value,omitempty"` // This string is the current value of a goal metric of type string.
	Current_number_value float64 `json:"current_number_value,omitempty"` // This number is the current value of a goal metric of type number.
	Initial_number_value float64 `json:"initial_number_value,omitempty"` // This number is the start value of a goal metric of type number.
	Unit string `json:"unit,omitempty"` // A supported unit of measure for the goal metric, or none.
	Progress_source string `json:"progress_source,omitempty"` // This field defines how the progress value of a goal metric is being calculated. A goal's progress can be provided manually by the user, calculated automatically from contributing subgoals or projects, or managed by an integration with an external data source, such as Salesforce.
	Target_number_value float64 `json:"target_number_value,omitempty"` // This number is the end value of a goal metric of type number. This number cannot equal `initial_number_value`.
	Currency_code string `json:"currency_code,omitempty"` // ISO 4217 currency code to format this custom field. This will be null if the `unit` is not `currency`.
	Precision int `json:"precision,omitempty"` // *Conditional*. Only relevant for goal metrics of type ‘Number’. This field dictates the number of places after the decimal to round to, i.e. 0 is integer values, 1 rounds to the nearest tenth, and so on. Must be between 0 and 6, inclusive. For percentage format, this may be unintuitive, as a value of 0.25 has a precision of 0, while a value of 0.251 has a precision of 1. This is due to 0.25 being displayed as 25%.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The subtype of this resource. Different subtypes retain many of the same fields and behavior, but may render differently in Asana or represent resources with different semantic meaning.
}

// DateVariableCompact represents the DateVariableCompact schema from the OpenAPI specification
type DateVariableCompact struct {
	Description string `json:"description,omitempty"` // The description of what the date variable is used for when instantiating a project.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the date field in the project template. A value of `1` refers to the project start date, while `2` refers to the project due date.
	Name string `json:"name,omitempty"` // The name of the date variable.
}

// TaskAddFollowersRequest represents the TaskAddFollowersRequest schema from the OpenAPI specification
type TaskAddFollowersRequest struct {
	Followers []string `json:"followers"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
}

// TagCompact represents the TagCompact schema from the OpenAPI specification
type TagCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // Name of the tag. This is generally a short sentence fragment that fits on a line in the UI for maximum readability. However, it can be longer.
}

// BatchRequestAction represents the BatchRequestAction schema from the OpenAPI specification
type BatchRequestAction struct {
	Relative_path string `json:"relative_path"` // The path of the desired endpoint relative to the API’s base URL. Query parameters are not accepted here; put them in `data` instead.
	Data map[string]interface{} `json:"data,omitempty"` // For `GET` requests, this should be a map of query parameters you would have normally passed in the URL. Options and pagination are not accepted here; put them in `options` instead. For `POST`, `PATCH`, and `PUT` methods, this should be the content you would have normally put in the data field of the body.
	Method string `json:"method"` // The HTTP method you wish to emulate for the action.
	Options map[string]interface{} `json:"options,omitempty"` // Pagination (`limit` and `offset`) and output options (`fields` or `expand`) for the action. “Pretty” JSON output is not an available option on individual actions; if you want pretty output, specify that option on the parent request.
}

// ModifyDependentsRequest represents the ModifyDependentsRequest schema from the OpenAPI specification
type ModifyDependentsRequest struct {
	Dependents []string `json:"dependents,omitempty"` // An array of task gids that are dependents of the given task.
}

// SectionRequest represents the SectionRequest schema from the OpenAPI specification
type SectionRequest struct {
	Insert_after string `json:"insert_after,omitempty"` // An existing section within this project after which the added section should be inserted. Cannot be provided together with insert_before.
	Insert_before string `json:"insert_before,omitempty"` // An existing section within this project before which the added section should be inserted. Cannot be provided together with insert_after.
	Name string `json:"name"` // The text to be displayed as the section name. This cannot be an empty string.
}

// UserResponse represents the UserResponse schema from the OpenAPI specification
type UserResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // *Read-only except when same user as requester*. The user’s name.
	Email string `json:"email,omitempty"` // The user's email address.
	Photo map[string]interface{} `json:"photo,omitempty"` // A map of the user’s profile photo in various sizes, or null if no photo is set. Sizes provided are 21, 27, 36, 60, 128, and 1024. All images are in PNG format, except for 1024 (which is in JPEG format).
	Workspaces []WorkspaceCompact `json:"workspaces,omitempty"` // Workspaces and organizations this user may access. Note\: The API will only return workspaces and organizations that also contain the authenticated user.
}

// CustomFieldCompact represents the CustomFieldCompact schema from the OpenAPI specification
type CustomFieldCompact struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	TypeField string `json:"type,omitempty"` // *Deprecated: new integrations should prefer the resource_subtype field.* The type of the custom field. Must be one of the given values.
	Enabled bool `json:"enabled,omitempty"` // *Conditional*. Determines if the custom field is enabled or not.
	Multi_enum_values []EnumOption `json:"multi_enum_values,omitempty"` // *Conditional*. Only relevant for custom fields of type `multi_enum`. This object is the chosen values of a `multi_enum` custom field.
	Name string `json:"name,omitempty"` // The name of the custom field.
	Display_value string `json:"display_value,omitempty"` // A string representation for the value of the custom field. Integrations that don't require the underlying type should use this field to read values. Using this field will future-proof an app against new custom field types.
	Enum_value interface{} `json:"enum_value,omitempty"`
	Date_value map[string]interface{} `json:"date_value,omitempty"` // *Conditional*. Only relevant for custom fields of type `date`. This object reflects the chosen date (and optionally, time) value of a `date` custom field. If no date is selected, the value of `date_value` will be `null`.
	Enum_options []EnumOption `json:"enum_options,omitempty"` // *Conditional*. Only relevant for custom fields of type `enum`. This array specifies the possible values which an `enum` custom field can adopt. To modify the enum options, refer to [working with enum options](/docs/create-an-enum-option).
	Number_value float64 `json:"number_value,omitempty"` // *Conditional*. This number is the value of a `number` custom field.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The type of the custom field. Must be one of the given values.
	Text_value string `json:"text_value,omitempty"` // *Conditional*. This string is the value of a `text` custom field.
}

// RemoveMembersRequest represents the RemoveMembersRequest schema from the OpenAPI specification
type RemoveMembersRequest struct {
	Members string `json:"members"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
}

// WebhookUpdateRequest represents the WebhookUpdateRequest schema from the OpenAPI specification
type WebhookUpdateRequest struct {
	Filters []interface{} `json:"filters,omitempty"` // An array of WebhookFilter objects to specify a whitelist of filters to apply to events from this webhook. If a webhook event passes any of the filters the event will be delivered; otherwise no event will be sent to the receiving server.
}

// ProjectDuplicateRequest represents the ProjectDuplicateRequest schema from the OpenAPI specification
type ProjectDuplicateRequest struct {
	Include string `json:"include,omitempty"` // The elements that will be duplicated to the new project. Tasks are always included.
	Name string `json:"name"` // The name of the new project.
	Schedule_dates map[string]interface{} `json:"schedule_dates,omitempty"` // A dictionary of options to auto-shift dates. `task_dates` must be included to use this option. Requires either `start_on` or `due_on`, but not both.
	Team string `json:"team,omitempty"` // Sets the team of the new project. If team is not defined, the new project will be in the same team as the the original project.
}

// ProjectStatusRequest represents the ProjectStatusRequest schema from the OpenAPI specification
type ProjectStatusRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Title string `json:"title,omitempty"` // The title of the project status update.
	Text string `json:"text"` // The text content of the status update.
	Color string `json:"color"` // The color associated with the status update.
	Html_text string `json:"html_text,omitempty"` // [Opt In](/docs/input-output-options). The text content of the status update with formatting as HTML.
}

// UserCompact represents the UserCompact schema from the OpenAPI specification
type UserCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // *Read-only except when same user as requester*. The user’s name.
}

// AttachmentResponse represents the AttachmentResponse schema from the OpenAPI specification
type AttachmentResponse struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the file.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The service hosting the attachment. Valid values are `asana`, `dropbox`, `gdrive`, `onedrive`, `box`, `vimeo`, and `external`.
	Parent interface{} `json:"parent,omitempty"`
	Permanent_url string `json:"permanent_url,omitempty"`
	Size int `json:"size,omitempty"` // The size of the attachment in bytes. Only present when the `resource_subtype` is `asana`.
	View_url string `json:"view_url,omitempty"` // The URL where the attachment can be viewed, which may be friendlier to users in a browser than just directing them to a raw file. May be null if no view URL exists for the service.
	Connected_to_app bool `json:"connected_to_app,omitempty"` // Whether the attachment is connected to the app making the request for the purposes of showing an app components widget. Only present when the `resource_subtype` is `external` or `gdrive`.
	Created_at string `json:"created_at,omitempty"` // The time at which this resource was created.
	Download_url string `json:"download_url,omitempty"` // The URL containing the content of the attachment. *Note:* May be null if the attachment is hosted by [Box](https://www.box.com/) and will be null if the attachment is a Video Message hosted by [Vimeo](https://vimeo.com/). If present, this URL may only be valid for two minutes from the time of retrieval. You should avoid persisting this URL somewhere and just refresh it on demand to ensure you do not keep stale URLs.
	Host string `json:"host,omitempty"` // The service hosting the attachment. Valid values are `asana`, `dropbox`, `gdrive`, `box`, and `vimeo`.
}

// WorkspaceMembershipRequest represents the WorkspaceMembershipRequest schema from the OpenAPI specification
type WorkspaceMembershipRequest struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	User UserCompact `json:"user,omitempty"`
	Workspace WorkspaceCompact `json:"workspace,omitempty"`
}

// RemoveCustomFieldSettingRequest represents the RemoveCustomFieldSettingRequest schema from the OpenAPI specification
type RemoveCustomFieldSettingRequest struct {
	Custom_field string `json:"custom_field"` // The custom field to remove from this portfolio.
}

// TeamCompact represents the TeamCompact schema from the OpenAPI specification
type TeamCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the team.
}

// ProjectBriefBase represents the ProjectBriefBase schema from the OpenAPI specification
type ProjectBriefBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Html_text string `json:"html_text,omitempty"` // HTML formatted text for the project brief.
	Title string `json:"title,omitempty"` // The title of the project brief.
}

// RemoveFollowersRequest represents the RemoveFollowersRequest schema from the OpenAPI specification
type RemoveFollowersRequest struct {
	Followers string `json:"followers"` // An array of strings identifying users. These can either be the string "me", an email, or the gid of a user.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Help string `json:"help,omitempty"` // Additional information directing developers to resources on how to address and fix the problem, if available.
	Message string `json:"message,omitempty"` // Message providing more detail about the error that occurred, if available.
	Phrase string `json:"phrase,omitempty"` // *500 errors only*. A unique error phrase which can be used when contacting developer support to help identify the exact occurrence of the problem in Asana’s logs.
}

// AttachmentRequest represents the AttachmentRequest schema from the OpenAPI specification
type AttachmentRequest struct {
	Name string `json:"name,omitempty"` // The name of the external resource being attached. Required for attachments of type `external`.
	Parent string `json:"parent,omitempty"` // Required identifier of the parent task, project, or project_brief, as a string.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The type of the attachment. Must be one of the given values. If not specified, a file attachment of type `asana` will be assumed. Note that if the value of `resource_subtype` is `external`, a `parent`, `name`, and `url` must also be provided.
	Url string `json:"url,omitempty"` // The URL of the external resource being attached. Required for attachments of type `external`.
	Connect_to_app bool `json:"connect_to_app,omitempty"` // *Optional*. Only relevant for external attachments with a parent task. A boolean indicating whether the current app should be connected with the attachment for the purposes of showing an app components widget. Requires the app to have been added to a project the parent task is in.
	File string `json:"file,omitempty"` // Required for `asana` attachments.
}

// TeamBase represents the TeamBase schema from the OpenAPI specification
type TeamBase struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the team.
}

// UserTaskListCompact represents the UserTaskListCompact schema from the OpenAPI specification
type UserTaskListCompact struct {
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Name string `json:"name,omitempty"` // The name of the user task list.
	Owner interface{} `json:"owner,omitempty"` // The owner of the user task list, i.e. the person whose My Tasks is represented by this resource.
	Workspace interface{} `json:"workspace,omitempty"` // The workspace in which the user task list is located.
}

// CustomFieldResponse represents the CustomFieldResponse schema from the OpenAPI specification
type CustomFieldResponse struct {
	Resource_type string `json:"resource_type,omitempty"` // The base type of this resource.
	Gid string `json:"gid,omitempty"` // Globally unique identifier of the resource, as a string.
	Name string `json:"name,omitempty"` // The name of the custom field.
	Display_value string `json:"display_value,omitempty"` // A string representation for the value of the custom field. Integrations that don't require the underlying type should use this field to read values. Using this field will future-proof an app against new custom field types.
	Enum_value interface{} `json:"enum_value,omitempty"`
	Date_value map[string]interface{} `json:"date_value,omitempty"` // *Conditional*. Only relevant for custom fields of type `date`. This object reflects the chosen date (and optionally, time) value of a `date` custom field. If no date is selected, the value of `date_value` will be `null`.
	Enum_options []EnumOption `json:"enum_options,omitempty"` // *Conditional*. Only relevant for custom fields of type `enum`. This array specifies the possible values which an `enum` custom field can adopt. To modify the enum options, refer to [working with enum options](/docs/create-an-enum-option).
	Number_value float64 `json:"number_value,omitempty"` // *Conditional*. This number is the value of a `number` custom field.
	Resource_subtype string `json:"resource_subtype,omitempty"` // The type of the custom field. Must be one of the given values.
	Text_value string `json:"text_value,omitempty"` // *Conditional*. This string is the value of a `text` custom field.
	TypeField string `json:"type,omitempty"` // *Deprecated: new integrations should prefer the resource_subtype field.* The type of the custom field. Must be one of the given values.
	Enabled bool `json:"enabled,omitempty"` // *Conditional*. Determines if the custom field is enabled or not.
	Multi_enum_values []EnumOption `json:"multi_enum_values,omitempty"` // *Conditional*. Only relevant for custom fields of type `multi_enum`. This object is the chosen values of a `multi_enum` custom field.
	Custom_label_position string `json:"custom_label_position,omitempty"` // Only relevant for custom fields with `custom` format. This depicts where to place the custom label. This will be null if the `format` is not `custom`.
	Enum_options []EnumOption `json:"enum_options,omitempty"` // *Conditional*. Only relevant for custom fields of type `enum`. This array specifies the possible values which an `enum` custom field can adopt. To modify the enum options, refer to [working with enum options](/docs/create-an-enum-option).
	Is_global_to_workspace bool `json:"is_global_to_workspace,omitempty"` // This flag describes whether this custom field is available to every container in the workspace. Before project-specific custom fields, this field was always true.
	Precision int `json:"precision,omitempty"` // Only relevant for custom fields of type ‘Number’. This field dictates the number of places after the decimal to round to, i.e. 0 is integer values, 1 rounds to the nearest tenth, and so on. Must be between 0 and 6, inclusive. For percentage format, this may be unintuitive, as a value of 0.25 has a precision of 0, while a value of 0.251 has a precision of 1. This is due to 0.25 being displayed as 25%. The identifier format will always have a precision of 0.
	Currency_code string `json:"currency_code,omitempty"` // ISO 4217 currency code to format this custom field. This will be null if the `format` is not `currency`.
	Format string `json:"format,omitempty"` // The format of this custom field.
	Asana_created_field string `json:"asana_created_field,omitempty"` // *Conditional*. A unique identifier to associate this field with the template source of truth.
	Custom_label string `json:"custom_label,omitempty"` // This is the string that appears next to the custom field value. This will be null if the `format` is not `custom`.
	Description string `json:"description,omitempty"` // [Opt In](/docs/input-output-options). The description of the custom field.
	Has_notifications_enabled bool `json:"has_notifications_enabled,omitempty"` // *Conditional*. This flag describes whether a follower of a task with this field should receive inbox notifications from changes to this field.
	Created_by UserCompact `json:"created_by,omitempty"`
	People_value []UserCompact `json:"people_value,omitempty"` // *Conditional*. Only relevant for custom fields of type `people`. This array of [compact user](/docs/user-compact) objects reflects the values of a `people` custom field.
}
