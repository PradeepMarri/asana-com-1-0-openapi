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

func GetgoalrelationshipsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["opt_pretty"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("opt_pretty=%v", val))
		}
		if val, ok := args["opt_fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("opt_fields=%v", val))
		}
		if val, ok := args["supported_goal"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("supported_goal=%v", val))
		}
		if val, ok := args["resource_subtype"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resource_subtype=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("offset=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/goal_relationships%s", cfg.BaseURL, queryString)
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

func CreateGetgoalrelationshipsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_goal_relationships",
		mcp.WithDescription("Get goal relationships"),
		mcp.WithBoolean("opt_pretty", mcp.Description("Provides “pretty” output.\nProvides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")),
		mcp.WithArray("opt_fields", mcp.Description("Defines fields to return.\nSome requests return *compact* representations of objects in order to conserve resources and complete the request more efficiently. Other times requests return more information than you may need. This option allows you to list the exact set of fields that the API should be sure to return for the objects. The field names should be provided as paths, described below.\nThe id of included objects will always be returned, regardless of the field options.")),
		mcp.WithString("supported_goal", mcp.Required(), mcp.Description("Globally unique identifier for the supported goal in the goal relationship.")),
		mcp.WithString("resource_subtype", mcp.Description("If provided, filter to goal relationships with a given resource_subtype.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetgoalrelationshipsHandler(cfg),
	}
}
