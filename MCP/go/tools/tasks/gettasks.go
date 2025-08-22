package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/asana/mcp-server/config"
	"github.com/asana/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GettasksHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["assignee"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("assignee=%v", val))
		}
		if val, ok := args["project"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project=%v", val))
		}
		if val, ok := args["section"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("section=%v", val))
		}
		if val, ok := args["workspace"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("workspace=%v", val))
		}
		if val, ok := args["completed_since"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("completed_since=%v", val))
		}
		if val, ok := args["modified_since"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified_since=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("offset=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/tasks%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGettasksTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_tasks",
		mcp.WithDescription("Get multiple tasks"),
		mcp.WithNumber("limit", mcp.Description("Results per page.\nThe number of objects to return per page. The value must be between 1 and 100.")),
		mcp.WithString("offset", mcp.Description("Offset token.\nAn offset to the next page returned by the API. A pagination request will return an offset token, which can be used as an input parameter to the next request. If an offset is not passed in, the API will return the first page of results.\n'Note: You can only pass in an offset that was returned to you via a previously paginated request.'")),
		mcp.WithString("assignee", mcp.Description("The assignee to filter tasks on. If searching for unassigned tasks, assignee.any = null can be specified.\n*Note: If you specify `assignee`, you must also specify the `workspace` to filter on.*")),
		mcp.WithString("project", mcp.Description("The project to filter tasks on.")),
		mcp.WithString("section", mcp.Description("The section to filter tasks on.\n*Note: Currently, this is only supported in board views.*")),
		mcp.WithString("workspace", mcp.Description("The workspace to filter tasks on.\n*Note: If you specify `workspace`, you must also specify the `assignee` to filter on.*")),
		mcp.WithString("completed_since", mcp.Description("Only return tasks that are either incomplete or that have been completed since this time.")),
		mcp.WithString("modified_since", mcp.Description("Only return tasks that have been modified since the given time.\n\n*Note: A task is considered “modified” if any of its properties\nchange, or associations between it and other objects are modified\n(e.g.  a task being added to a project). A task is not considered\nmodified just because another object it is associated with (e.g. a\nsubtask) is modified. Actions that count as modifying the task\ninclude assigning, renaming, completing, and adding stories.*")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GettasksHandler(cfg),
	}
}
