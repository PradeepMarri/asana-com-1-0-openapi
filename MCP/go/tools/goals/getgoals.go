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

func GetgoalsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["portfolio"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("portfolio=%v", val))
		}
		if val, ok := args["project"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project=%v", val))
		}
		if val, ok := args["is_workspace_level"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_workspace_level=%v", val))
		}
		if val, ok := args["team"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("team=%v", val))
		}
		if val, ok := args["workspace"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("workspace=%v", val))
		}
		if val, ok := args["time_periods"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("time_periods=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("offset=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/goals%s", cfg.BaseURL, queryString)
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

func CreateGetgoalsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_goals",
		mcp.WithDescription("Get goals"),
		mcp.WithString("portfolio", mcp.Description("Globally unique identifier for supporting portfolio.")),
		mcp.WithString("project", mcp.Description("Globally unique identifier for supporting project.")),
		mcp.WithBoolean("is_workspace_level", mcp.Description("Filter to goals with is_workspace_level set to query value. Must be used with the workspace parameter.")),
		mcp.WithString("team", mcp.Description("Globally unique identifier for the team.")),
		mcp.WithString("workspace", mcp.Description("Globally unique identifier for the workspace.")),
		mcp.WithArray("time_periods", mcp.Description("Globally unique identifiers for the time periods.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetgoalsHandler(cfg),
	}
}
